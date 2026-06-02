// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SystemApiLogDao is the data access object for the table system_api_log.
type SystemApiLogDao struct {
	table    string              // table is the underlying table name of the DAO.
	group    string              // group is the database configuration group name of the current DAO.
	columns  SystemApiLogColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler  // handlers for customized model modification.
}

// SystemApiLogColumns defines and stores column names for the table system_api_log.
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

// systemApiLogColumns holds the columns for the table system_api_log.
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
func NewSystemApiLogDao(handlers ...gdb.ModelHandler) *SystemApiLogDao {
	return &SystemApiLogDao{
		group:    "default",
		table:    "system_api_log",
		columns:  systemApiLogColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SystemApiLogDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SystemApiLogDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SystemApiLogDao) Columns() SystemApiLogColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SystemApiLogDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SystemApiLogDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *SystemApiLogDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
