package system

import (
	"testing"

	"devinggo/modules/system/model/res"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/stretchr/testify/assert"
)

func TestBuildColumnDetailsFromFields(t *testing.T) {
	fields := map[string]*gdb.TableField{
		"name": {
			Index:   2,
			Name:    "name",
			Type:    "varchar(255)",
			Null:    true,
			Default: "guest",
			Comment: "名称",
		},
		"id": {
			Index:   1,
			Name:    "id",
			Type:    "bigint",
			Null:    false,
			Key:     "PRI",
			Default: nil,
			Comment: "主键",
		},
	}

	assert.Equal(t, []res.DataMaintainColumn{
		{
			Field:        "id",
			Type:         "bigint",
			Nullable:     false,
			Key:          "PRI",
			DefaultValue: "",
			Comment:      "主键",
		},
		{
			Field:        "name",
			Type:         "varchar(255)",
			Nullable:     true,
			Key:          "",
			DefaultValue: "guest",
			Comment:      "名称",
		},
	}, buildColumnDetailsFromFields(fields))
}

func TestBuildColumnDetailsFromFieldsReturnsEmptySlice(t *testing.T) {
	details := buildColumnDetailsFromFields(nil)

	assert.NotNil(t, details)
	assert.Empty(t, details)
}
