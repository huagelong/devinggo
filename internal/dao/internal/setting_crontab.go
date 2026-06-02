// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SettingCrontabDao is the data access object for the table setting_crontab.
type SettingCrontabDao struct {
	table    string                // table is the underlying table name of the DAO.
	group    string                // group is the database configuration group name of the current DAO.
	columns  SettingCrontabColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler    // handlers for customized model modification.
}

// SettingCrontabColumns defines and stores column names for the table setting_crontab.
type SettingCrontabColumns struct {
	Id        string //
	Name      string //
	Type      string //
	Target    string //
	Parameter string //
	Rule      string //
	Singleton string //
	Status    string //
	CreatedBy string //
	UpdatedBy string //
	CreatedAt string //
	UpdatedAt string //
	Remark    string //
}

// settingCrontabColumns holds the columns for the table setting_crontab.
var settingCrontabColumns = SettingCrontabColumns{
	Id:        "id",
	Name:      "name",
	Type:      "type",
	Target:    "target",
	Parameter: "parameter",
	Rule:      "rule",
	Singleton: "singleton",
	Status:    "status",
	CreatedBy: "created_by",
	UpdatedBy: "updated_by",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	Remark:    "remark",
}

// NewSettingCrontabDao creates and returns a new DAO object for table data access.
func NewSettingCrontabDao(handlers ...gdb.ModelHandler) *SettingCrontabDao {
	return &SettingCrontabDao{
		group:    "default",
		table:    "setting_crontab",
		columns:  settingCrontabColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SettingCrontabDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SettingCrontabDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SettingCrontabDao) Columns() SettingCrontabColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SettingCrontabDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SettingCrontabDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *SettingCrontabDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
