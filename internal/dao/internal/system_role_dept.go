// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SystemRoleDeptDao is the data access object for the table system_role_dept.
type SystemRoleDeptDao struct {
	table    string                // table is the underlying table name of the DAO.
	group    string                // group is the database configuration group name of the current DAO.
	columns  SystemRoleDeptColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler    // handlers for customized model modification.
}

// SystemRoleDeptColumns defines and stores column names for the table system_role_dept.
type SystemRoleDeptColumns struct {
	RoleId string //
	DeptId string //
}

// systemRoleDeptColumns holds the columns for the table system_role_dept.
var systemRoleDeptColumns = SystemRoleDeptColumns{
	RoleId: "role_id",
	DeptId: "dept_id",
}

// NewSystemRoleDeptDao creates and returns a new DAO object for table data access.
func NewSystemRoleDeptDao(handlers ...gdb.ModelHandler) *SystemRoleDeptDao {
	return &SystemRoleDeptDao{
		group:    "default",
		table:    "system_role_dept",
		columns:  systemRoleDeptColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SystemRoleDeptDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SystemRoleDeptDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SystemRoleDeptDao) Columns() SystemRoleDeptColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SystemRoleDeptDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SystemRoleDeptDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *SystemRoleDeptDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
