// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SystemDeptLeaderDao is the data access object for the table system_dept_leader.
type SystemDeptLeaderDao struct {
	table    string                  // table is the underlying table name of the DAO.
	group    string                  // group is the database configuration group name of the current DAO.
	columns  SystemDeptLeaderColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler      // handlers for customized model modification.
}

// SystemDeptLeaderColumns defines and stores column names for the table system_dept_leader.
type SystemDeptLeaderColumns struct {
	DeptId    string //
	UserId    string //
	Username  string //
	CreatedAt string //
}

// systemDeptLeaderColumns holds the columns for the table system_dept_leader.
var systemDeptLeaderColumns = SystemDeptLeaderColumns{
	DeptId:    "dept_id",
	UserId:    "user_id",
	Username:  "username",
	CreatedAt: "created_at",
}

// NewSystemDeptLeaderDao creates and returns a new DAO object for table data access.
func NewSystemDeptLeaderDao(handlers ...gdb.ModelHandler) *SystemDeptLeaderDao {
	return &SystemDeptLeaderDao{
		group:    "default",
		table:    "system_dept_leader",
		columns:  systemDeptLeaderColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SystemDeptLeaderDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SystemDeptLeaderDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SystemDeptLeaderDao) Columns() SystemDeptLeaderColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SystemDeptLeaderDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SystemDeptLeaderDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *SystemDeptLeaderDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
