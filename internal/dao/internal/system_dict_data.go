// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SystemDictDataDao is the data access object for table system_dict_data.
type SystemDictDataDao struct {
	table   string                // table is the underlying table name of the DAO.
	group   string                // group is the database configuration group name of current DAO.
	columns SystemDictDataColumns // columns contains all the column names of Table for convenient usage.
}

// SystemDictDataColumns defines and stores column names for table system_dict_data.
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

// systemDictDataColumns holds the columns for table system_dict_data.
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
func NewSystemDictDataDao() *SystemDictDataDao {
	return &SystemDictDataDao{
		group:   "default",
		table:   "system_dict_data",
		columns: systemDictDataColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SystemDictDataDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SystemDictDataDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SystemDictDataDao) Columns() SystemDictDataColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SystemDictDataDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SystemDictDataDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SystemDictDataDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
