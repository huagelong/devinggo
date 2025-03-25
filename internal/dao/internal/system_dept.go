// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SystemDeptDao is the data access object for table system_dept.
type SystemDeptDao struct {
	table   string            // table is the underlying table name of the DAO.
	group   string            // group is the database configuration group name of current DAO.
	columns SystemDeptColumns // columns contains all the column names of Table for convenient usage.
}

// SystemDeptColumns defines and stores column names for table system_dept.
type SystemDeptColumns struct {
	Id        string //
	ParentId  string //
	Level     string //
	Name      string //
	Leader    string //
	Phone     string //
	Status    string //
	Sort      string //
	CreatedBy string //
	UpdatedBy string //
	CreatedAt string //
	UpdatedAt string //
	DeletedAt string //
	Remark    string //
}

// systemDeptColumns holds the columns for table system_dept.
var systemDeptColumns = SystemDeptColumns{
	Id:        "id",
	ParentId:  "parent_id",
	Level:     "level",
	Name:      "name",
	Leader:    "leader",
	Phone:     "phone",
	Status:    "status",
	Sort:      "sort",
	CreatedBy: "created_by",
	UpdatedBy: "updated_by",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
	Remark:    "remark",
}

// NewSystemDeptDao creates and returns a new DAO object for table data access.
func NewSystemDeptDao() *SystemDeptDao {
	return &SystemDeptDao{
		group:   "default",
		table:   "system_dept",
		columns: systemDeptColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SystemDeptDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SystemDeptDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SystemDeptDao) Columns() SystemDeptColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SystemDeptDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SystemDeptDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SystemDeptDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
