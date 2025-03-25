// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SystemAppDao is the data access object for table system_app.
type SystemAppDao struct {
	table   string           // table is the underlying table name of the DAO.
	group   string           // group is the database configuration group name of current DAO.
	columns SystemAppColumns // columns contains all the column names of Table for convenient usage.
}

// SystemAppColumns defines and stores column names for table system_app.
type SystemAppColumns struct {
	Id          string //
	GroupId     string //
	AppName     string //
	AppId       string //
	AppSecret   string //
	Status      string //
	Description string //
	CreatedBy   string //
	UpdatedBy   string //
	CreatedAt   string //
	UpdatedAt   string //
	DeletedAt   string //
	Remark      string //
}

// systemAppColumns holds the columns for table system_app.
var systemAppColumns = SystemAppColumns{
	Id:          "id",
	GroupId:     "group_id",
	AppName:     "app_name",
	AppId:       "app_id",
	AppSecret:   "app_secret",
	Status:      "status",
	Description: "description",
	CreatedBy:   "created_by",
	UpdatedBy:   "updated_by",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
	DeletedAt:   "deleted_at",
	Remark:      "remark",
}

// NewSystemAppDao creates and returns a new DAO object for table data access.
func NewSystemAppDao() *SystemAppDao {
	return &SystemAppDao{
		group:   "default",
		table:   "system_app",
		columns: systemAppColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SystemAppDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SystemAppDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SystemAppDao) Columns() SystemAppColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SystemAppDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SystemAppDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SystemAppDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
