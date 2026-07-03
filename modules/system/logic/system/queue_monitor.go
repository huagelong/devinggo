package system

import (
	"context"
	"slices"
	"time"

	"devinggo/modules/system/logic/base"
	"devinggo/modules/system/model/page"
	"devinggo/modules/system/model/res"
	"devinggo/modules/system/myerror"
	glob2 "devinggo/modules/system/pkg/worker/glob"
	"devinggo/modules/system/service"

	"github.com/hibiken/asynq"
)

var queueMonitorTaskStates = []string{
	"pending",
	"active",
	"scheduled",
	"retry",
	"archived",
	"completed",
}

type queueMonitorTaskLister func(*asynq.Inspector, string, ...asynq.ListOption) ([]*asynq.TaskInfo, error)

type sQueueMonitor struct {
	base.BaseService
}

func init() {
	service.RegisterQueueMonitor(NewQueueMonitor())
}

func NewQueueMonitor() *sQueueMonitor {
	return &sQueueMonitor{}
}

func (s *sQueueMonitor) ListQueues(ctx context.Context) ([]res.QueueMonitorQueueItem, error) {
	inspector := asynq.NewInspector(glob2.GetConnConfig(ctx))
	defer inspector.Close()

	queues, err := inspector.Queues()
	if err != nil {
		return nil, err
	}

	items := make([]res.QueueMonitorQueueItem, 0, len(queues))
	for _, queue := range queues {
		info, infoErr := inspector.GetQueueInfo(queue)
		if infoErr != nil {
			return nil, infoErr
		}
		items = append(items, mapQueueInfo(info))
	}
	return items, nil
}

func (s *sQueueMonitor) GetOverview(ctx context.Context, queue string) (*res.QueueMonitorOverview, error) {
	inspector := asynq.NewInspector(glob2.GetConnConfig(ctx))
	defer inspector.Close()

	info, err := inspector.GetQueueInfo(queue)
	if err != nil {
		return nil, err
	}
	servers, err := inspector.Servers()
	if err != nil {
		return nil, err
	}

	workers := make([]res.QueueMonitorWorkerItem, 0, len(servers))
	for _, server := range servers {
		workers = append(workers, mapWorkerInfo(server))
	}

	return &res.QueueMonitorOverview{
		Queue:   mapQueueInfo(info),
		Workers: workers,
	}, nil
}

func (s *sQueueMonitor) ListTasks(ctx context.Context, queue string, state string, reqPage *page.PageReq) (items []res.QueueMonitorTaskItem, total int, err error) {
	taskState, err := resolveQueueMonitorTaskState(state)
	if err != nil {
		return nil, 0, err
	}

	inspector := asynq.NewInspector(glob2.GetConnConfig(ctx))
	defer inspector.Close()

	lister, err := getQueueMonitorTaskLister(taskState)
	if err != nil {
		return nil, 0, err
	}

	pageNum, pageSize, _ := page.CalPage(reqPage.Page, reqPage.PageSize)
	listOptions := []asynq.ListOption{
		asynq.PageSize(pageSize),
		asynq.Page(pageNum),
	}
	taskInfos, err := lister(inspector, queue, listOptions...)
	if err != nil {
		return nil, 0, err
	}

	items = make([]res.QueueMonitorTaskItem, 0, len(taskInfos))
	for _, item := range taskInfos {
		items = append(items, mapTaskInfo(item))
	}

	queueInfo, err := inspector.GetQueueInfo(queue)
	if err != nil {
		return nil, 0, err
	}
	total = queueTaskStateCount(queueInfo, taskState)
	return items, total, nil
}

func (s *sQueueMonitor) ListSchedulerEntries(ctx context.Context) ([]res.QueueMonitorSchedulerEntry, error) {
	inspector := asynq.NewInspector(glob2.GetConnConfig(ctx))
	defer inspector.Close()

	entries, err := inspector.SchedulerEntries()
	if err != nil {
		return nil, err
	}

	items := make([]res.QueueMonitorSchedulerEntry, 0, len(entries))
	for _, entry := range entries {
		items = append(items, mapSchedulerEntry(entry))
	}
	return items, nil
}

func resolveQueueMonitorTaskState(state string) (string, error) {
	if slices.Contains(queueMonitorTaskStates, state) {
		return state, nil
	}
	return "", myerror.ValidationFailed(context.Background(), "不支持的任务状态")
}

func getQueueMonitorTaskLister(state string) (queueMonitorTaskLister, error) {
	switch state {
	case "pending":
		return (*asynq.Inspector).ListPendingTasks, nil
	case "active":
		return (*asynq.Inspector).ListActiveTasks, nil
	case "scheduled":
		return (*asynq.Inspector).ListScheduledTasks, nil
	case "retry":
		return (*asynq.Inspector).ListRetryTasks, nil
	case "archived":
		return (*asynq.Inspector).ListArchivedTasks, nil
	case "completed":
		return (*asynq.Inspector).ListCompletedTasks, nil
	default:
		return nil, myerror.ValidationFailed(context.Background(), "不支持的任务状态")
	}
}

func mapQueueInfo(info *asynq.QueueInfo) res.QueueMonitorQueueItem {
	if info == nil {
		return res.QueueMonitorQueueItem{}
	}
	return res.QueueMonitorQueueItem{
		Queue:        info.Queue,
		MemoryUsage:  info.MemoryUsage,
		Latency:      int64(info.Latency / time.Second),
		Size:         info.Size,
		Groups:       info.Groups,
		Pending:      info.Pending,
		Active:       info.Active,
		Scheduled:    info.Scheduled,
		Retry:        info.Retry,
		Archived:     info.Archived,
		Completed:    info.Completed,
		Aggregating:  info.Aggregating,
		Processed:    info.Processed,
		Failed:       info.Failed,
		ProcessedAll: info.ProcessedTotal,
		FailedAll:    info.FailedTotal,
		Paused:       info.Paused,
		Timestamp:    formatQueueMonitorTime(info.Timestamp),
	}
}

func mapTaskInfo(info *asynq.TaskInfo) res.QueueMonitorTaskItem {
	if info == nil {
		return res.QueueMonitorTaskItem{}
	}
	return res.QueueMonitorTaskItem{
		ID:            info.ID,
		Queue:         info.Queue,
		Type:          info.Type,
		Payload:       string(info.Payload),
		State:         info.State.String(),
		MaxRetry:      info.MaxRetry,
		Retried:       info.Retried,
		LastErr:       info.LastErr,
		LastFailedAt:  formatQueueMonitorTime(info.LastFailedAt),
		Timeout:       int64(info.Timeout / time.Second),
		Deadline:      formatQueueMonitorTime(info.Deadline),
		Group:         info.Group,
		NextProcessAt: formatQueueMonitorTime(info.NextProcessAt),
		IsOrphaned:    info.IsOrphaned,
		Retention:     int64(info.Retention / time.Second),
		CompletedAt:   formatQueueMonitorTime(info.CompletedAt),
		Result:        string(info.Result),
	}
}

func mapSchedulerEntry(entry *asynq.SchedulerEntry) res.QueueMonitorSchedulerEntry {
	if entry == nil {
		return res.QueueMonitorSchedulerEntry{}
	}
	taskType := ""
	if entry.Task != nil {
		taskType = entry.Task.Type()
	}
	return res.QueueMonitorSchedulerEntry{
		ID:       entry.ID,
		Spec:     entry.Spec,
		TaskType: taskType,
		Queue:    extractQueueName(entry.Opts),
		Next:     formatQueueMonitorTime(entry.Next),
		Prev:     formatQueueMonitorTime(entry.Prev),
	}
}

func mapWorkerInfo(server *asynq.ServerInfo) res.QueueMonitorWorkerItem {
	if server == nil {
		return res.QueueMonitorWorkerItem{}
	}
	return res.QueueMonitorWorkerItem{
		ID:                server.ID,
		Host:              server.Host,
		PID:               server.PID,
		Concurrency:       server.Concurrency,
		Queues:            server.Queues,
		StrictPriority:    server.StrictPriority,
		Started:           formatQueueMonitorTime(server.Started),
		Status:            server.Status,
		ActiveWorkerCount: len(server.ActiveWorkers),
	}
}

func extractQueueName(opts []asynq.Option) string {
	for _, opt := range opts {
		if opt.Type() == asynq.QueueOpt {
			if queueName, ok := opt.Value().(string); ok && queueName != "" {
				return queueName
			}
		}
	}
	return "default"
}

func queueTaskStateCount(info *asynq.QueueInfo, state string) int {
	if info == nil {
		return 0
	}
	switch state {
	case "pending":
		return info.Pending
	case "active":
		return info.Active
	case "scheduled":
		return info.Scheduled
	case "retry":
		return info.Retry
	case "archived":
		return info.Archived
	case "completed":
		return info.Completed
	default:
		return 0
	}
}

func formatQueueMonitorTime(t time.Time) string {
	if t.IsZero() {
		return ""
	}
	return t.Format(time.RFC3339)
}
