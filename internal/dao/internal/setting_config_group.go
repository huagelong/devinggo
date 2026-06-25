// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SettingConfigGroupDao is the data access object for the table setting_config_group.
type SettingConfigGroupDao struct {
	table    string                    // table is the underlying table name of the DAO.
	group    string                    // group is the database configuration group name of the current DAO.
	columns  SettingConfigGroupColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler        // handlers for customized model modification.
}

// SettingConfigGroupColumns defines and stores column names for the table setting_config_group.
type SettingConfigGroupColumns struct {
	Id        string //
	Name      string //
	Code      string //
	CreatedBy string //
	UpdatedBy string //
	CreatedAt string //
	UpdatedAt string //
	Remark    string //
}

// settingConfigGroupColumns holds the columns for the table setting_config_group.
var settingConfigGroupColumns = SettingConfigGroupColumns{
	Id:        "id",
	Name:      "name",
	Code:      "code",
	CreatedBy: "created_by",
	UpdatedBy: "updated_by",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	Remark:    "remark",
}

// NewSettingConfigGroupDao creates and returns a new DAO object for table data access.
func NewSettingConfigGroupDao(handlers ...gdb.ModelHandler) *SettingConfigGroupDao {
	return &SettingConfigGroupDao{
		group:    "default",
		table:    "setting_config_group",
		columns:  settingConfigGroupColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SettingConfigGroupDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SettingConfigGroupDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SettingConfigGroupDao) Columns() SettingConfigGroupColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SettingConfigGroupDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SettingConfigGroupDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *SettingConfigGroupDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
