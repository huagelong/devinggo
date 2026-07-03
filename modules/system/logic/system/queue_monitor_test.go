package system

import (
	"testing"
	"time"

	"github.com/hibiken/asynq"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestResolveQueueMonitorTaskState(t *testing.T) {
	cases := []struct {
		name  string
		state string
		want  string
	}{
		{name: "pending", state: "pending", want: "pending"},
		{name: "active", state: "active", want: "active"},
		{name: "scheduled", state: "scheduled", want: "scheduled"},
		{name: "retry", state: "retry", want: "retry"},
		{name: "archived", state: "archived", want: "archived"},
		{name: "completed", state: "completed", want: "completed"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := resolveQueueMonitorTaskState(tc.state)
			require.NoError(t, err)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestResolveQueueMonitorTaskStateRejectsUnsupportedState(t *testing.T) {
	_, err := resolveQueueMonitorTaskState("invalid")
	require.Error(t, err)
}

func TestMapSchedulerEntry(t *testing.T) {
	next := time.Date(2026, 7, 3, 10, 0, 0, 0, time.UTC)
	prev := next.Add(-5 * time.Minute)

	entry := &asynq.SchedulerEntry{
		ID:   "entry-1",
		Spec: "@every 1m",
		Task: asynq.NewTask("system:test", []byte(`{"hello":"world"}`)),
		Opts: []asynq.Option{
			asynq.Queue("critical"),
			asynq.MaxRetry(5),
		},
		Next: next,
		Prev: prev,
	}

	got := mapSchedulerEntry(entry)
	assert.Equal(t, "entry-1", got.ID)
	assert.Equal(t, "@every 1m", got.Spec)
	assert.Equal(t, "system:test", got.TaskType)
	assert.Equal(t, "critical", got.Queue)
	assert.Equal(t, next.Format(time.RFC3339), got.Next)
	assert.Equal(t, prev.Format(time.RFC3339), got.Prev)
}

func TestMapSchedulerEntryDefaultsQueue(t *testing.T) {
	entry := &asynq.SchedulerEntry{
		ID:   "entry-default",
		Spec: "@every 5m",
		Task: asynq.NewTask("system:default", nil),
	}

	got := mapSchedulerEntry(entry)
	assert.Equal(t, "default", got.Queue)
}

func TestMapWorkerInfo(t *testing.T) {
	started := time.Date(2026, 7, 3, 8, 30, 0, 0, time.UTC)
	server := &asynq.ServerInfo{
		ID:          "srv-1",
		Host:        "worker-1",
		PID:         1001,
		Concurrency: 12,
		Queues: map[string]int{
			"critical": 6,
			"default":  3,
		},
		Status:  "active",
		Started: started,
		ActiveWorkers: []*asynq.WorkerInfo{
			{TaskID: "task-1", TaskType: "system:test", Queue: "critical"},
			{TaskID: "task-2", TaskType: "system:retry", Queue: "default"},
		},
	}

	got := mapWorkerInfo(server)
	assert.Equal(t, "srv-1", got.ID)
	assert.Equal(t, "worker-1", got.Host)
	assert.Equal(t, 1001, got.PID)
	assert.Equal(t, 12, got.Concurrency)
	assert.Equal(t, "active", got.Status)
	assert.Equal(t, started.Format(time.RFC3339), got.Started)
	assert.Equal(t, 2, got.ActiveWorkerCount)
	assert.Equal(t, map[string]int{
		"critical": 6,
		"default":  3,
	}, got.Queues)
}
