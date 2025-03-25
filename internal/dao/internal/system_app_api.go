// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SystemAppApiDao is the data access object for table system_app_api.
type SystemAppApiDao struct {
	table   string              // table is the underlying table name of the DAO.
	group   string              // group is the database configuration group name of current DAO.
	columns SystemAppApiColumns // columns contains all the column names of Table for convenient usage.
}

// SystemAppApiColumns defines and stores column names for table system_app_api.
type SystemAppApiColumns struct {
	AppId string //
	ApiId string //
}

// systemAppApiColumns holds the columns for table system_app_api.
var systemAppApiColumns = SystemAppApiColumns{
	AppId: "app_id",
	ApiId: "api_id",
}

// NewSystemAppApiDao creates and returns a new DAO object for table data access.
func NewSystemAppApiDao() *SystemAppApiDao {
	return &SystemAppApiDao{
		group:   "default",
		table:   "system_app_api",
		columns: systemAppApiColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SystemAppApiDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SystemAppApiDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SystemAppApiDao) Columns() SystemAppApiColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SystemAppApiDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SystemAppApiDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SystemAppApiDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
