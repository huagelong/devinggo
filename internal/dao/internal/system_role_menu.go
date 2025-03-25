// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SystemRoleMenuDao is the data access object for table system_role_menu.
type SystemRoleMenuDao struct {
	table   string                // table is the underlying table name of the DAO.
	group   string                // group is the database configuration group name of current DAO.
	columns SystemRoleMenuColumns // columns contains all the column names of Table for convenient usage.
}

// SystemRoleMenuColumns defines and stores column names for table system_role_menu.
type SystemRoleMenuColumns struct {
	RoleId string //
	MenuId string //
}

// systemRoleMenuColumns holds the columns for table system_role_menu.
var systemRoleMenuColumns = SystemRoleMenuColumns{
	RoleId: "role_id",
	MenuId: "menu_id",
}

// NewSystemRoleMenuDao creates and returns a new DAO object for table data access.
func NewSystemRoleMenuDao() *SystemRoleMenuDao {
	return &SystemRoleMenuDao{
		group:   "default",
		table:   "system_role_menu",
		columns: systemRoleMenuColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SystemRoleMenuDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SystemRoleMenuDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SystemRoleMenuDao) Columns() SystemRoleMenuColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SystemRoleMenuDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SystemRoleMenuDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SystemRoleMenuDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
