// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SystemUserDeptDao is the data access object for table system_user_dept.
type SystemUserDeptDao struct {
	table   string                // table is the underlying table name of the DAO.
	group   string                // group is the database configuration group name of current DAO.
	columns SystemUserDeptColumns // columns contains all the column names of Table for convenient usage.
}

// SystemUserDeptColumns defines and stores column names for table system_user_dept.
type SystemUserDeptColumns struct {
	UserId string //
	DeptId string //
}

// systemUserDeptColumns holds the columns for table system_user_dept.
var systemUserDeptColumns = SystemUserDeptColumns{
	UserId: "user_id",
	DeptId: "dept_id",
}

// NewSystemUserDeptDao creates and returns a new DAO object for table data access.
func NewSystemUserDeptDao() *SystemUserDeptDao {
	return &SystemUserDeptDao{
		group:   "default",
		table:   "system_user_dept",
		columns: systemUserDeptColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SystemUserDeptDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SystemUserDeptDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SystemUserDeptDao) Columns() SystemUserDeptColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SystemUserDeptDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SystemUserDeptDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SystemUserDeptDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
