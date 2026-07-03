package system

import (
	"testing"
	"time"

	"devinggo/modules/system/model/res"

	"github.com/stretchr/testify/assert"
)

func TestFilterPostgresGroups(t *testing.T) {
	raw := map[string]any{
		"logger": map[string]any{"path": "logs/"},
		"default": map[string]any{
			"link": "pgsql:postgres:pass@tcp(127.0.0.1:5432)/devinggo?sslmode=disable",
		},
		"report": map[string]any{
			"link": "pgsql:postgres:pass@tcp(127.0.0.1:5432)/reporting?sslmode=disable",
		},
		"shadow": map[string]any{
			"link": "mysql:root:pass@tcp(127.0.0.1:3306)/shadow",
		},
	}

	groups := filterPostgresGroups(raw)
	assert.Equal(t, []res.DbMonitorGroup{
		{GroupName: "default", IsDefault: true, DbType: "postgres"},
		{GroupName: "report", IsDefault: false, DbType: "postgres"},
	}, groups)
}

func TestBuildPointAndCapabilities(t *testing.T) {
	prev := dbSnapshot{
		CollectedAt:         time.Unix(0, 0),
		XactCommit:          100,
		XactRollback:        2,
		BlksHit:             900,
		BlksRead:            100,
		HasPgStatStatements: true,
		TotalCalls:          200,
		TotalExecMs:         5000,
		SlowCalls:           10,
	}
	curr := dbSnapshot{
		CollectedAt:         time.Unix(10, 0),
		XactCommit:          170,
		XactRollback:        5,
		BlksHit:             990,
		BlksRead:            110,
		HasPgStatStatements: true,
		TotalCalls:          260,
		TotalExecMs:         6800,
		SlowCalls:           13,
	}

	point := buildPoint(prev, curr)
	assert.InDelta(t, 7.3, point.Tps, 0.1)
	assert.InDelta(t, 0.3, point.RollbackTps, 0.1)
	if assert.NotNil(t, point.AvgQueryMs) {
		assert.InDelta(t, 30, *point.AvgQueryMs, 0.1)
	}
	if assert.NotNil(t, point.SlowQueryCount) {
		assert.Equal(t, int64(3), *point.SlowQueryCount)
	}

	capabilities := buildCapabilities(false)
	assert.False(t, capabilities.HasPgStatStatements)
	assert.Contains(t, capabilities.UnavailableMetrics, "avgQueryMs")
	assert.Contains(t, capabilities.UnavailableMetrics, "slowQueryCount")
}
