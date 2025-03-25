// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SettingConfigGroupDao is the data access object for table setting_config_group.
type SettingConfigGroupDao struct {
	table   string                    // table is the underlying table name of the DAO.
	group   string                    // group is the database configuration group name of current DAO.
	columns SettingConfigGroupColumns // columns contains all the column names of Table for convenient usage.
}

// SettingConfigGroupColumns defines and stores column names for table setting_config_group.
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

// settingConfigGroupColumns holds the columns for table setting_config_group.
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
func NewSettingConfigGroupDao() *SettingConfigGroupDao {
	return &SettingConfigGroupDao{
		group:   "default",
		table:   "setting_config_group",
		columns: settingConfigGroupColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SettingConfigGroupDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SettingConfigGroupDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SettingConfigGroupDao) Columns() SettingConfigGroupColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SettingConfigGroupDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SettingConfigGroupDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SettingConfigGroupDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
