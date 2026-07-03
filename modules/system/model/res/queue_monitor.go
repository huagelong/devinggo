package res

type QueueMonitorQueueItem struct {
	Queue        string `json:"queue"`
	MemoryUsage  int64  `json:"memoryUsage"`
	Latency      int64  `json:"latency"`
	Size         int    `json:"size"`
	Groups       int    `json:"groups"`
	Pending      int    `json:"pending"`
	Active       int    `json:"active"`
	Scheduled    int    `json:"scheduled"`
	Retry        int    `json:"retry"`
	Archived     int    `json:"archived"`
	Completed    int    `json:"completed"`
	Aggregating  int    `json:"aggregating"`
	Processed    int    `json:"processed"`
	Failed       int    `json:"failed"`
	ProcessedAll int    `json:"processedTotal"`
	FailedAll    int    `json:"failedTotal"`
	Paused       bool   `json:"paused"`
	Timestamp    string `json:"timestamp"`
}

type QueueMonitorTaskItem struct {
	ID            string `json:"id"`
	Queue         string `json:"queue"`
	Type          string `json:"type"`
	Payload       string `json:"payload"`
	State         string `json:"state"`
	MaxRetry      int    `json:"maxRetry"`
	Retried       int    `json:"retried"`
	LastErr       string `json:"lastErr"`
	LastFailedAt  string `json:"lastFailedAt"`
	Timeout       int64  `json:"timeout"`
	Deadline      string `json:"deadline"`
	Group         string `json:"group"`
	NextProcessAt string `json:"nextProcessAt"`
	IsOrphaned    bool   `json:"isOrphaned"`
	Retention     int64  `json:"retention"`
	CompletedAt   string `json:"completedAt"`
	Result        string `json:"result"`
}

type QueueMonitorSchedulerEntry struct {
	ID       string `json:"id"`
	Spec     string `json:"spec"`
	TaskType string `json:"taskType"`
	Queue    string `json:"queue"`
	Next     string `json:"next"`
	Prev     string `json:"prev"`
}

type QueueMonitorWorkerItem struct {
	ID                string         `json:"id"`
	Host              string         `json:"host"`
	PID               int            `json:"pid"`
	Concurrency       int            `json:"concurrency"`
	Queues            map[string]int `json:"queues"`
	StrictPriority    bool           `json:"strictPriority"`
	Started           string         `json:"started"`
	Status            string         `json:"status"`
	ActiveWorkerCount int            `json:"activeWorkerCount"`
}

type QueueMonitorOverview struct {
	Queue   QueueMonitorQueueItem    `json:"queue"`
	Workers []QueueMonitorWorkerItem `json:"workers"`
}
