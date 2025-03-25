// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SystemUserRoleDao is the data access object for table system_user_role.
type SystemUserRoleDao struct {
	table   string                // table is the underlying table name of the DAO.
	group   string                // group is the database configuration group name of current DAO.
	columns SystemUserRoleColumns // columns contains all the column names of Table for convenient usage.
}

// SystemUserRoleColumns defines and stores column names for table system_user_role.
type SystemUserRoleColumns struct {
	UserId string //
	RoleId string //
}

// systemUserRoleColumns holds the columns for table system_user_role.
var systemUserRoleColumns = SystemUserRoleColumns{
	UserId: "user_id",
	RoleId: "role_id",
}

// NewSystemUserRoleDao creates and returns a new DAO object for table data access.
func NewSystemUserRoleDao() *SystemUserRoleDao {
	return &SystemUserRoleDao{
		group:   "default",
		table:   "system_user_role",
		columns: systemUserRoleColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SystemUserRoleDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SystemUserRoleDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SystemUserRoleDao) Columns() SystemUserRoleColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SystemUserRoleDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SystemUserRoleDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SystemUserRoleDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
