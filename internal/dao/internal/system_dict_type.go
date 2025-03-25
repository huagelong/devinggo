// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SystemDictTypeDao is the data access object for table system_dict_type.
type SystemDictTypeDao struct {
	table   string                // table is the underlying table name of the DAO.
	group   string                // group is the database configuration group name of current DAO.
	columns SystemDictTypeColumns // columns contains all the column names of Table for convenient usage.
}

// SystemDictTypeColumns defines and stores column names for table system_dict_type.
type SystemDictTypeColumns struct {
	Id        string //
	Name      string //
	Code      string //
	Status    string //
	CreatedBy string //
	UpdatedBy string //
	CreatedAt string //
	UpdatedAt string //
	DeletedAt string //
	Remark    string //
}

// systemDictTypeColumns holds the columns for table system_dict_type.
var systemDictTypeColumns = SystemDictTypeColumns{
	Id:        "id",
	Name:      "name",
	Code:      "code",
	Status:    "status",
	CreatedBy: "created_by",
	UpdatedBy: "updated_by",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
	Remark:    "remark",
}

// NewSystemDictTypeDao creates and returns a new DAO object for table data access.
func NewSystemDictTypeDao() *SystemDictTypeDao {
	return &SystemDictTypeDao{
		group:   "default",
		table:   "system_dict_type",
		columns: systemDictTypeColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SystemDictTypeDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SystemDictTypeDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SystemDictTypeDao) Columns() SystemDictTypeColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SystemDictTypeDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SystemDictTypeDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SystemDictTypeDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
