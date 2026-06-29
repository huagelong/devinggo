// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SystemUserDeptDao is the data access object for the table system_user_dept.
type SystemUserDeptDao struct {
	table    string                // table is the underlying table name of the DAO.
	group    string                // group is the database configuration group name of the current DAO.
	columns  SystemUserDeptColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler    // handlers for customized model modification.
}

// SystemUserDeptColumns defines and stores column names for the table system_user_dept.
type SystemUserDeptColumns struct {
	UserId string //
	DeptId string //
}

// systemUserDeptColumns holds the columns for the table system_user_dept.
var systemUserDeptColumns = SystemUserDeptColumns{
	UserId: "user_id",
	DeptId: "dept_id",
}

// NewSystemUserDeptDao creates and returns a new DAO object for table data access.
func NewSystemUserDeptDao(handlers ...gdb.ModelHandler) *SystemUserDeptDao {
	return &SystemUserDeptDao{
		group:    "default",
		table:    "system_user_dept",
		columns:  systemUserDeptColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SystemUserDeptDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SystemUserDeptDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SystemUserDeptDao) Columns() SystemUserDeptColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SystemUserDeptDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SystemUserDeptDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *SystemUserDeptDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
