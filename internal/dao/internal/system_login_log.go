// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SystemLoginLogDao is the data access object for table system_login_log.
type SystemLoginLogDao struct {
	table   string                // table is the underlying table name of the DAO.
	group   string                // group is the database configuration group name of current DAO.
	columns SystemLoginLogColumns // columns contains all the column names of Table for convenient usage.
}

// SystemLoginLogColumns defines and stores column names for table system_login_log.
type SystemLoginLogColumns struct {
	Id         string //
	Username   string //
	Ip         string //
	IpLocation string //
	Os         string //
	Browser    string //
	Status     string //
	Message    string //
	LoginTime  string //
	Remark     string //
}

// systemLoginLogColumns holds the columns for table system_login_log.
var systemLoginLogColumns = SystemLoginLogColumns{
	Id:         "id",
	Username:   "username",
	Ip:         "ip",
	IpLocation: "ip_location",
	Os:         "os",
	Browser:    "browser",
	Status:     "status",
	Message:    "message",
	LoginTime:  "login_time",
	Remark:     "remark",
}

// NewSystemLoginLogDao creates and returns a new DAO object for table data access.
func NewSystemLoginLogDao() *SystemLoginLogDao {
	return &SystemLoginLogDao{
		group:   "default",
		table:   "system_login_log",
		columns: systemLoginLogColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SystemLoginLogDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SystemLoginLogDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SystemLoginLogDao) Columns() SystemLoginLogColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SystemLoginLogDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SystemLoginLogDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SystemLoginLogDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
