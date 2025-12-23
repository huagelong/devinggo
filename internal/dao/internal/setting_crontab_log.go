// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SettingCrontabLogDao is the data access object for the table setting_crontab_log.
type SettingCrontabLogDao struct {
	table    string                   // table is the underlying table name of the DAO.
	group    string                   // group is the database configuration group name of the current DAO.
	columns  SettingCrontabLogColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler       // handlers for customized model modification.
}

// SettingCrontabLogColumns defines and stores column names for the table setting_crontab_log.
type SettingCrontabLogColumns struct {
	Id            string // 主键
	CrontabId     string // 任务ID
	Name          string // 任务名称
	Target        string // 任务调用目标字符串
	Parameter     string // 任务调用参数
	ExceptionInfo string // 异常信息
	Status        string // 执行状态 (1成功 2失败)
	CreatedAt     string // 创建时间
}

// settingCrontabLogColumns holds the columns for the table setting_crontab_log.
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
func NewSettingCrontabLogDao(handlers ...gdb.ModelHandler) *SettingCrontabLogDao {
	return &SettingCrontabLogDao{
		group:    "default",
		table:    "setting_crontab_log",
		columns:  settingCrontabLogColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SettingCrontabLogDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SettingCrontabLogDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SettingCrontabLogDao) Columns() SettingCrontabLogColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SettingCrontabLogDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SettingCrontabLogDao) Ctx(ctx context.Context) *gdb.Model {
	model := dao.DB().Model(dao.table)
	for _, handler := range dao.handlers {
		model = handler(model)
	}
	return model.Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *SettingCrontabLogDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
