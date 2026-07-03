package system

import (
	"devinggo/modules/system/model"
	"devinggo/modules/system/model/page"
	"devinggo/modules/system/model/res"

	"github.com/gogf/gf/v2/frame/g"
)

type QueueMonitorQueuesReq struct {
	g.Meta `path:"/queueMonitor/queues" method:"get" tags:"队列监控" summary:"获取队列列表." x-permission:"system:queueMonitor:queues"`
	model.AuthorHeader
}

type QueueMonitorQueuesRes struct {
	g.Meta `mime:"application/json"`
	Data   []res.QueueMonitorQueueItem `json:"data"`
}

type QueueMonitorOverviewReq struct {
	g.Meta `path:"/queueMonitor/overview" method:"get" tags:"队列监控" summary:"获取队列概览." x-permission:"system:queueMonitor:overview"`
	model.AuthorHeader
	Queue string `json:"queue" v:"required#队列名称不能为空"`
}

type QueueMonitorOverviewRes struct {
	g.Meta `mime:"application/json"`
	Data   res.QueueMonitorOverview `json:"data"`
}

type QueueMonitorTasksReq struct {
	g.Meta `path:"/queueMonitor/tasks" method:"get" tags:"队列监控" summary:"获取队列任务列表." x-permission:"system:queueMonitor:tasks"`
	model.AuthorHeader
	page.PageReq
	Queue string `json:"queue" v:"required#队列名称不能为空"`
	State string `json:"state" v:"required#任务状态不能为空"`
}

type QueueMonitorTasksRes struct {
	g.Meta `mime:"application/json"`
	page.PageRes
	Items []res.QueueMonitorTaskItem `json:"items"`
}

type QueueMonitorSchedulerEntriesReq struct {
	g.Meta `path:"/queueMonitor/schedulerEntries" method:"get" tags:"队列监控" summary:"获取周期调度列表." x-permission:"system:queueMonitor:schedulerEntries"`
	model.AuthorHeader
}

type QueueMonitorSchedulerEntriesRes struct {
	g.Meta `mime:"application/json"`
	Data   []res.QueueMonitorSchedulerEntry `json:"data"`
}
