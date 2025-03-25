// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SystemApiDao is the data access object for table system_api.
type SystemApiDao struct {
	table   string           // table is the underlying table name of the DAO.
	group   string           // group is the database configuration group name of current DAO.
	columns SystemApiColumns // columns contains all the column names of Table for convenient usage.
}

// SystemApiColumns defines and stores column names for table system_api.
type SystemApiColumns struct {
	Id          string //
	GroupId     string //
	Name        string //
	AccessName  string //
	AuthMode    string //
	RequestMode string //
	Status      string //
	CreatedBy   string //
	UpdatedBy   string //
	CreatedAt   string //
	UpdatedAt   string //
	DeletedAt   string //
	Remark      string //
}

// systemApiColumns holds the columns for table system_api.
var systemApiColumns = SystemApiColumns{
	Id:          "id",
	GroupId:     "group_id",
	Name:        "name",
	AccessName:  "access_name",
	AuthMode:    "auth_mode",
	RequestMode: "request_mode",
	Status:      "status",
	CreatedBy:   "created_by",
	UpdatedBy:   "updated_by",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
	DeletedAt:   "deleted_at",
	Remark:      "remark",
}

// NewSystemApiDao creates and returns a new DAO object for table data access.
func NewSystemApiDao() *SystemApiDao {
	return &SystemApiDao{
		group:   "default",
		table:   "system_api",
		columns: systemApiColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SystemApiDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SystemApiDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SystemApiDao) Columns() SystemApiColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SystemApiDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SystemApiDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SystemApiDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
