// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SystemApiLogDao is the data access object for table system_api_log.
type SystemApiLogDao struct {
	table   string              // table is the underlying table name of the DAO.
	group   string              // group is the database configuration group name of current DAO.
	columns SystemApiLogColumns // columns contains all the column names of Table for convenient usage.
}

// SystemApiLogColumns defines and stores column names for table system_api_log.
type SystemApiLogColumns struct {
	Id           string //
	ApiId        string //
	ApiName      string //
	AccessName   string //
	RequestData  string //
	ResponseCode string //
	ResponseData string //
	Ip           string //
	IpLocation   string //
	AccessTime   string //
	Remark       string //
}

// systemApiLogColumns holds the columns for table system_api_log.
var systemApiLogColumns = SystemApiLogColumns{
	Id:           "id",
	ApiId:        "api_id",
	ApiName:      "api_name",
	AccessName:   "access_name",
	RequestData:  "request_data",
	ResponseCode: "response_code",
	ResponseData: "response_data",
	Ip:           "ip",
	IpLocation:   "ip_location",
	AccessTime:   "access_time",
	Remark:       "remark",
}

// NewSystemApiLogDao creates and returns a new DAO object for table data access.
func NewSystemApiLogDao() *SystemApiLogDao {
	return &SystemApiLogDao{
		group:   "default",
		table:   "system_api_log",
		columns: systemApiLogColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SystemApiLogDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SystemApiLogDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SystemApiLogDao) Columns() SystemApiLogColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SystemApiLogDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SystemApiLogDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SystemApiLogDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
