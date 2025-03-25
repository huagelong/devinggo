// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SettingConfigDao is the data access object for table setting_config.
type SettingConfigDao struct {
	table   string               // table is the underlying table name of the DAO.
	group   string               // group is the database configuration group name of current DAO.
	columns SettingConfigColumns // columns contains all the column names of Table for convenient usage.
}

// SettingConfigColumns defines and stores column names for table setting_config.
type SettingConfigColumns struct {
	GroupId          string // 组id
	Key              string // 配置键名
	Value            string // 配置值
	Name             string //
	InputType        string //
	ConfigSelectData string //
	Sort             string //
	Remark           string //
}

// settingConfigColumns holds the columns for table setting_config.
var settingConfigColumns = SettingConfigColumns{
	GroupId:          "group_id",
	Key:              "key",
	Value:            "value",
	Name:             "name",
	InputType:        "input_type",
	ConfigSelectData: "config_select_data",
	Sort:             "sort",
	Remark:           "remark",
}

// NewSettingConfigDao creates and returns a new DAO object for table data access.
func NewSettingConfigDao() *SettingConfigDao {
	return &SettingConfigDao{
		group:   "default",
		table:   "setting_config",
		columns: settingConfigColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SettingConfigDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SettingConfigDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SettingConfigDao) Columns() SettingConfigColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SettingConfigDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SettingConfigDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SettingConfigDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
