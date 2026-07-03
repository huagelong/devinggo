package res

type DbMonitorGroup struct {
	GroupName string `json:"groupName"`
	IsDefault bool   `json:"isDefault"`
	DbType    string `json:"dbType"`
}

type DbMonitorSummary struct {
	ConnectionStatus string  `json:"connectionStatus"`
	CurrentConn      int64   `json:"currentConn"`
	ActiveConn       int64   `json:"activeConn"`
	DatabaseSize     int64   `json:"databaseSize"`
	HitRate          float64 `json:"hitRate"`
	SlowQueryEnabled bool    `json:"slowQueryEnabled"`
}

type DbMonitorOverview struct {
	GroupName          string  `json:"groupName"`
	DatabaseName       string  `json:"databaseName"`
	Version            string  `json:"version"`
	MaxConnections     int64   `json:"maxConnections"`
	CurrentConnections int64   `json:"currentConnections"`
	ActiveConnections  int64   `json:"activeConnections"`
	IdleConnections    int64   `json:"idleConnections"`
	WaitingConnections int64   `json:"waitingConnections"`
	XactCommit         int64   `json:"xactCommit"`
	XactRollback       int64   `json:"xactRollback"`
	HitRate            float64 `json:"hitRate"`
	DatabaseSize       int64   `json:"databaseSize"`
}

type DbMonitorTimeseriesPoint struct {
	Time           string   `json:"time"`
	Tps            float64  `json:"tps"`
	RollbackTps    float64  `json:"rollbackTps"`
	AvgQueryMs     *float64 `json:"avgQueryMs"`
	SlowQueryCount *int64   `json:"slowQueryCount"`
}

type DbMonitorCapabilities struct {
	HasPgStatStatements bool     `json:"hasPgStatStatements"`
	AvailableMetrics    []string `json:"availableMetrics"`
	UnavailableMetrics  []string `json:"unavailableMetrics"`
	SlowQueryThreshold  int64    `json:"slowQueryThreshold"`
}

type DbMonitorAlert struct {
	Level   string `json:"level"`
	Message string `json:"message"`
}

type DbMonitorPayload struct {
	Summary      DbMonitorSummary           `json:"summary"`
	Overview     DbMonitorOverview          `json:"overview"`
	Timeseries   []DbMonitorTimeseriesPoint `json:"timeseries"`
	Capabilities DbMonitorCapabilities      `json:"capabilities"`
	Alerts       []DbMonitorAlert           `json:"alerts"`
}
