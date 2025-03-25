// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SettingCrontabLogDao is the data access object for table setting_crontab_log.
type SettingCrontabLogDao struct {
	table   string                   // table is the underlying table name of the DAO.
	group   string                   // group is the database configuration group name of current DAO.
	columns SettingCrontabLogColumns // columns contains all the column names of Table for convenient usage.
}

// SettingCrontabLogColumns defines and stores column names for table setting_crontab_log.
type SettingCrontabLogColumns struct {
	Id            string //
	CrontabId     string //
	Name          string //
	Target        string //
	Parameter     string //
	ExceptionInfo string //
	Status        string //
	CreatedAt     string //
}

// settingCrontabLogColumns holds the columns for table setting_crontab_log.
var settingCrontabLogColumns = SettingCrontabLogColumns{
	Id:            "id",
	CrontabId:     "crontab_id",
	Name:          "name",
	Target:        "target",
	Parameter:     "parameter",
	ExceptionInfo: "exception_info",
	Status:        "status",
	CreatedAt:     "created_at",
}

// NewSettingCrontabLogDao creates and returns a new DAO object for table data access.
func NewSettingCrontabLogDao() *SettingCrontabLogDao {
	return &SettingCrontabLogDao{
		group:   "default",
		table:   "setting_crontab_log",
		columns: settingCrontabLogColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SettingCrontabLogDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SettingCrontabLogDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SettingCrontabLogDao) Columns() SettingCrontabLogColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SettingCrontabLogDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SettingCrontabLogDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SettingCrontabLogDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
