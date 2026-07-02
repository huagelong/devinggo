package system

import (
	"context"
	"sort"
	"strings"
	"sync"
	"time"

	"devinggo/modules/system/logic/base"
	"devinggo/modules/system/model/res"
	"devinggo/modules/system/myerror"
	"devinggo/modules/system/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

const (
	dbMonitorSeriesLimit     = 60
	slowQueryThresholdMillis = 1000.0
)

type sDbMonitor struct {
	base.BaseService
	mu     sync.Mutex
	series map[string]*groupSeries
}

type dbSnapshot struct {
	CollectedAt         time.Time
	GroupName           string
	DatabaseName        string
	Version             string
	MaxConnections      int64
	CurrentConnections  int64
	ActiveConnections   int64
	IdleConnections     int64
	WaitingConnections  int64
	XactCommit          int64
	XactRollback        int64
	BlksHit             int64
	BlksRead            int64
	DatabaseSize        int64
	HasPgStatStatements bool
	TotalCalls          int64
	TotalExecMs         float64
	SlowCalls           int64
}

type groupSeries struct {
	Last   *dbSnapshot
	Points []res.DbMonitorTimeseriesPoint
}

func init() {
	service.RegisterDbMonitor(NewDbMonitor())
}

func NewDbMonitor() *sDbMonitor {
	return &sDbMonitor{
		series: make(map[string]*groupSeries),
	}
}

func (s *sDbMonitor) ListGroups(ctx context.Context) ([]res.DbMonitorGroup, error) {
	raw := g.Cfg().MustGet(ctx, "database").Map()
	return filterPostgresGroups(raw), nil
}

func (s *sDbMonitor) GetMonitor(ctx context.Context, groupName string) (*res.DbMonitorPayload, error) {
	snapshot, err := s.collectSnapshot(ctx, groupName)
	if err != nil {
		return nil, err
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	state := s.series[groupName]
	if state == nil {
		state = &groupSeries{
			Points: make([]res.DbMonitorTimeseriesPoint, 0, dbMonitorSeriesLimit),
		}
		s.series[groupName] = state
	}
	if state.Last != nil {
		point := buildPoint(*state.Last, *snapshot)
		state.Points = append(state.Points, point)
		if len(state.Points) > dbMonitorSeriesLimit {
			state.Points = state.Points[len(state.Points)-dbMonitorSeriesLimit:]
		}
	}
	state.Last = snapshot

	hitRate := buildHitRate(*snapshot)
	return &res.DbMonitorPayload{
		Summary: res.DbMonitorSummary{
			ConnectionStatus: "up",
			CurrentConn:      snapshot.CurrentConnections,
			ActiveConn:       snapshot.ActiveConnections,
			DatabaseSize:     snapshot.DatabaseSize,
			HitRate:          hitRate,
			SlowQueryEnabled: snapshot.HasPgStatStatements,
		},
		Overview: res.DbMonitorOverview{
			GroupName:          snapshot.GroupName,
			DatabaseName:       snapshot.DatabaseName,
			Version:            snapshot.Version,
			MaxConnections:     snapshot.MaxConnections,
			CurrentConnections: snapshot.CurrentConnections,
			ActiveConnections:  snapshot.ActiveConnections,
			IdleConnections:    snapshot.IdleConnections,
			WaitingConnections: snapshot.WaitingConnections,
			XactCommit:         snapshot.XactCommit,
			XactRollback:       snapshot.XactRollback,
			HitRate:            hitRate,
			DatabaseSize:       snapshot.DatabaseSize,
		},
		Timeseries:   append([]res.DbMonitorTimeseriesPoint(nil), state.Points...),
		Capabilities: buildCapabilities(snapshot.HasPgStatStatements),
		Alerts:       buildAlerts(snapshot.HasPgStatStatements),
	}, nil
}

func filterPostgresGroups(raw map[string]any) []res.DbMonitorGroup {
	groups := make([]res.DbMonitorGroup, 0)
	for key, value := range raw {
		if key == "logger" {
			continue
		}
		link := strings.ToLower(gconv.String(gconv.Map(value)["link"]))
		if !strings.HasPrefix(link, "pgsql:") {
			continue
		}
		groups = append(groups, res.DbMonitorGroup{
			GroupName: key,
			IsDefault: key == "default",
			DbType:    "postgres",
		})
	}
	sort.Slice(groups, func(i, j int) bool {
		if groups[i].IsDefault != groups[j].IsDefault {
			return groups[i].IsDefault
		}
		return groups[i].GroupName < groups[j].GroupName
	})
	return groups
}

func buildPoint(prev, curr dbSnapshot) res.DbMonitorTimeseriesPoint {
	seconds := curr.CollectedAt.Sub(prev.CollectedAt).Seconds()
	if seconds <= 0 {
		seconds = 1
	}

	totalTxDelta := float64((curr.XactCommit + curr.XactRollback) - (prev.XactCommit + prev.XactRollback))
	rollbackDelta := float64(curr.XactRollback - prev.XactRollback)
	point := res.DbMonitorTimeseriesPoint{
		Time:        curr.CollectedAt.Format("15:04:05"),
		Tps:         totalTxDelta / seconds,
		RollbackTps: rollbackDelta / seconds,
	}

	if curr.HasPgStatStatements && prev.HasPgStatStatements {
		callDelta := curr.TotalCalls - prev.TotalCalls
		execDelta := curr.TotalExecMs - prev.TotalExecMs
		slowDelta := curr.SlowCalls - prev.SlowCalls
		if callDelta > 0 {
			value := execDelta / float64(callDelta)
			point.AvgQueryMs = &value
		}
		if slowDelta < 0 {
			slowDelta = 0
		}
		point.SlowQueryCount = &slowDelta
	}

	return point
}

func buildHitRate(snapshot dbSnapshot) float64 {
	total := snapshot.BlksHit + snapshot.BlksRead
	if total == 0 {
		return 0
	}
	return float64(snapshot.BlksHit) / float64(total) * 100
}

func buildCapabilities(hasPgStatStatements bool) res.DbMonitorCapabilities {
	out := res.DbMonitorCapabilities{
		HasPgStatStatements: hasPgStatStatements,
		AvailableMetrics: []string{
			"tps", "rollbackTps", "hitRate", "connections",
		},
		UnavailableMetrics: []string{},
		SlowQueryThreshold: int64(slowQueryThresholdMillis),
	}
	if hasPgStatStatements {
		out.AvailableMetrics = append(out.AvailableMetrics, "avgQueryMs", "slowQueryCount")
	} else {
		out.UnavailableMetrics = append(out.UnavailableMetrics, "avgQueryMs", "slowQueryCount")
	}
	return out
}

func buildAlerts(hasPgStatStatements bool) []res.DbMonitorAlert {
	if hasPgStatStatements {
		return nil
	}
	return []res.DbMonitorAlert{
		{
			Level:   "warning",
			Message: "pg_stat_statements is unavailable",
		},
	}
}

func (s *sDbMonitor) collectSnapshot(ctx context.Context, groupName string) (*dbSnapshot, error) {
	db := g.DB(groupName)
	if db == nil {
		return nil, myerror.ValidationFailed(ctx, "数据库组不存在")
	}

	row, err := db.GetOne(ctx, `
SELECT
	current_database() AS database_name,
	current_setting('server_version') AS version,
	current_setting('max_connections')::bigint AS max_connections,
	(SELECT pg_database_size(current_database())) AS database_size,
	(SELECT count(*) FROM pg_stat_activity WHERE datname = current_database()) AS current_connections,
	(SELECT count(*) FROM pg_stat_activity WHERE datname = current_database() AND state = 'active') AS active_connections,
	(SELECT count(*) FROM pg_stat_activity WHERE datname = current_database() AND state = 'idle') AS idle_connections,
	(SELECT count(*) FROM pg_stat_activity WHERE datname = current_database() AND wait_event IS NOT NULL) AS waiting_connections,
	xact_commit,
	xact_rollback,
	blks_hit,
	blks_read
FROM pg_stat_database
WHERE datname = current_database()
`)
	if err != nil {
		return nil, err
	}

	snapshot := &dbSnapshot{
		CollectedAt:        time.Now(),
		GroupName:          groupName,
		DatabaseName:       gconv.String(row["database_name"]),
		Version:            gconv.String(row["version"]),
		MaxConnections:     gconv.Int64(row["max_connections"]),
		CurrentConnections: gconv.Int64(row["current_connections"]),
		ActiveConnections:  gconv.Int64(row["active_connections"]),
		IdleConnections:    gconv.Int64(row["idle_connections"]),
		WaitingConnections: gconv.Int64(row["waiting_connections"]),
		XactCommit:         gconv.Int64(row["xact_commit"]),
		XactRollback:       gconv.Int64(row["xact_rollback"]),
		BlksHit:            gconv.Int64(row["blks_hit"]),
		BlksRead:           gconv.Int64(row["blks_read"]),
		DatabaseSize:       gconv.Int64(row["database_size"]),
	}

	hasExt, err := db.GetValue(ctx, `SELECT EXISTS(SELECT 1 FROM pg_extension WHERE extname = 'pg_stat_statements')`)
	if err == nil && gconv.Bool(hasExt) {
		snapshot.HasPgStatStatements = true
		stats, statsErr := db.GetOne(ctx, `
SELECT
	COALESCE(sum(calls), 0) AS total_calls,
	COALESCE(sum(total_exec_time), 0) AS total_exec_ms,
	COALESCE(sum(CASE WHEN mean_exec_time >= ? THEN calls ELSE 0 END), 0) AS slow_calls
FROM pg_stat_statements
`, slowQueryThresholdMillis)
		if statsErr == nil {
			snapshot.TotalCalls = gconv.Int64(stats["total_calls"])
			snapshot.TotalExecMs = gconv.Float64(stats["total_exec_ms"])
			snapshot.SlowCalls = gconv.Int64(stats["slow_calls"])
		}
	}

	return snapshot, nil
}
