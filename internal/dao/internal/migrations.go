// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// MigrationsDao is the data access object for table migrations.
type MigrationsDao struct {
	table   string            // table is the underlying table name of the DAO.
	group   string            // group is the database configuration group name of current DAO.
	columns MigrationsColumns // columns contains all the column names of Table for convenient usage.
}

// MigrationsColumns defines and stores column names for table migrations.
type MigrationsColumns struct {
	Id        string //
	Migration string //
	Batch     string //
}

// migrationsColumns holds the columns for table migrations.
var migrationsColumns = MigrationsColumns{
	Id:        "id",
	Migration: "migration",
	Batch:     "batch",
}

// NewMigrationsDao creates and returns a new DAO object for table data access.
func NewMigrationsDao() *MigrationsDao {
	return &MigrationsDao{
		group:   "default",
		table:   "migrations",
		columns: migrationsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *MigrationsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *MigrationsDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *MigrationsDao) Columns() MigrationsColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *MigrationsDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *MigrationsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *MigrationsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
