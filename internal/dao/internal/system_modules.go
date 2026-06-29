// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SystemModulesDao is the data access object for the table system_modules.
type SystemModulesDao struct {
	table    string               // table is the underlying table name of the DAO.
	group    string               // group is the database configuration group name of the current DAO.
	columns  SystemModulesColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler   // handlers for customized model modification.
}

// SystemModulesColumns defines and stores column names for the table system_modules.
type SystemModulesColumns struct {
	Id          string //
	CreatedAt   string //
	UpdatedAt   string //
	CreatedBy   string //
	UpdatedBy   string //
	Name        string //
	Label       string //
	Description string //
	Installed   string //
	Status      string //
	DeletedAt   string //
}

// systemModulesColumns holds the columns for the table system_modules.
var systemModulesColumns = SystemModulesColumns{
	Id:          "id",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
	CreatedBy:   "created_by",
	UpdatedBy:   "updated_by",
	Name:        "name",
	Label:       "label",
	Description: "description",
	Installed:   "installed",
	Status:      "status",
	DeletedAt:   "deleted_at",
}

// NewSystemModulesDao creates and returns a new DAO object for table data access.
func NewSystemModulesDao(handlers ...gdb.ModelHandler) *SystemModulesDao {
	return &SystemModulesDao{
		group:    "default",
		table:    "system_modules",
		columns:  systemModulesColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SystemModulesDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SystemModulesDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SystemModulesDao) Columns() SystemModulesColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SystemModulesDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SystemModulesDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *SystemModulesDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
