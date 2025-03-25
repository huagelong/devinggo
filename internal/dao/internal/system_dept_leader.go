// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SystemDeptLeaderDao is the data access object for table system_dept_leader.
type SystemDeptLeaderDao struct {
	table   string                  // table is the underlying table name of the DAO.
	group   string                  // group is the database configuration group name of current DAO.
	columns SystemDeptLeaderColumns // columns contains all the column names of Table for convenient usage.
}

// SystemDeptLeaderColumns defines and stores column names for table system_dept_leader.
type SystemDeptLeaderColumns struct {
	DeptId    string //
	UserId    string //
	Username  string //
	CreatedAt string //
}

// systemDeptLeaderColumns holds the columns for table system_dept_leader.
var systemDeptLeaderColumns = SystemDeptLeaderColumns{
	DeptId:    "dept_id",
	UserId:    "user_id",
	Username:  "username",
	CreatedAt: "created_at",
}

// NewSystemDeptLeaderDao creates and returns a new DAO object for table data access.
func NewSystemDeptLeaderDao() *SystemDeptLeaderDao {
	return &SystemDeptLeaderDao{
		group:   "default",
		table:   "system_dept_leader",
		columns: systemDeptLeaderColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SystemDeptLeaderDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SystemDeptLeaderDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SystemDeptLeaderDao) Columns() SystemDeptLeaderColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SystemDeptLeaderDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SystemDeptLeaderDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SystemDeptLeaderDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
