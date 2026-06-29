// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SystemDictDataDao is the data access object for the table system_dict_data.
type SystemDictDataDao struct {
	table    string                // table is the underlying table name of the DAO.
	group    string                // group is the database configuration group name of the current DAO.
	columns  SystemDictDataColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler    // handlers for customized model modification.
}

// SystemDictDataColumns defines and stores column names for the table system_dict_data.
type SystemDictDataColumns struct {
	Id        string //
	TypeId    string //
	Label     string //
	Value     string //
	Code      string //
	Sort      string //
	Status    string //
	CreatedBy string //
	UpdatedBy string //
	CreatedAt string //
	UpdatedAt string //
	DeletedAt string //
	Remark    string //
}

// systemDictDataColumns holds the columns for the table system_dict_data.
var systemDictDataColumns = SystemDictDataColumns{
	Id:        "id",
	TypeId:    "type_id",
	Label:     "label",
	Value:     "value",
	Code:      "code",
	Sort:      "sort",
	Status:    "status",
	CreatedBy: "created_by",
	UpdatedBy: "updated_by",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
	Remark:    "remark",
}

// NewSystemDictDataDao creates and returns a new DAO object for table data access.
func NewSystemDictDataDao(handlers ...gdb.ModelHandler) *SystemDictDataDao {
	return &SystemDictDataDao{
		group:    "default",
		table:    "system_dict_data",
		columns:  systemDictDataColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SystemDictDataDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SystemDictDataDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SystemDictDataDao) Columns() SystemDictDataColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SystemDictDataDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SystemDictDataDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *SystemDictDataDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
