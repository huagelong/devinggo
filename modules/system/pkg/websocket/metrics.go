package websocket

import (
	"sync"
	"time"
)

const pusherMonitorMetricWindow = 5 * time.Minute

type pusherAppMetricState struct {
	events []time.Time
	errors []time.Time
}

var (
	pusherAppMetrics   = make(map[string]*pusherAppMetricState)
	pusherAppMetricsMu sync.Mutex
)

func RecordAppEvent(appID string) {
	recordAppMetric(appID, false)
}

func RecordAppError(appID string) {
	recordAppMetric(appID, true)
}

func GetRecentAppMetrics(appID string) (eventCount int, errorCount int) {
	if appID == "" {
		return 0, 0
	}
	pusherAppMetricsMu.Lock()
	defer pusherAppMetricsMu.Unlock()

	state := pusherAppMetrics[appID]
	if state == nil {
		return 0, 0
	}

	now := time.Now()
	state.events = compactMetricTimes(state.events, now)
	state.errors = compactMetricTimes(state.errors, now)
	return len(state.events), len(state.errors)
}

func recordAppMetric(appID string, isError bool) {
	if appID == "" {
		return
	}
	pusherAppMetricsMu.Lock()
	defer pusherAppMetricsMu.Unlock()

	state := pusherAppMetrics[appID]
	if state == nil {
		state = &pusherAppMetricState{}
		pusherAppMetrics[appID] = state
	}

	now := time.Now()
	state.events = compactMetricTimes(state.events, now)
	state.errors = compactMetricTimes(state.errors, now)
	if isError {
		state.errors = append(state.errors, now)
		return
	}
	state.events = append(state.events, now)
}

func compactMetricTimes(items []time.Time, now time.Time) []time.Time {
	if len(items) == 0 {
		return items
	}
	threshold := now.Add(-pusherMonitorMetricWindow)
	index := 0
	for index < len(items) && items[index].Before(threshold) {
		index++
	}
	if index == 0 {
		return items
	}
	return append([]time.Time(nil), items[index:]...)
}
