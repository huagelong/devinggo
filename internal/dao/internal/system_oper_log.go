// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SystemOperLogDao is the data access object for the table system_oper_log.
type SystemOperLogDao struct {
	table    string               // table is the underlying table name of the DAO.
	group    string               // group is the database configuration group name of the current DAO.
	columns  SystemOperLogColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler   // handlers for customized model modification.
}

// SystemOperLogColumns defines and stores column names for the table system_oper_log.
type SystemOperLogColumns struct {
	Id           string //
	Username     string //
	Method       string //
	Router       string //
	ServiceName  string //
	Ip           string //
	IpLocation   string //
	RequestData  string //
	ResponseCode string //
	ResponseData string //
	CreatedBy    string //
	UpdatedBy    string //
	CreatedAt    string //
	UpdatedAt    string //
	DeletedAt    string //
	Remark       string //
}

// systemOperLogColumns holds the columns for the table system_oper_log.
var systemOperLogColumns = SystemOperLogColumns{
	Id:           "id",
	Username:     "username",
	Method:       "method",
	Router:       "router",
	ServiceName:  "service_name",
	Ip:           "ip",
	IpLocation:   "ip_location",
	RequestData:  "request_data",
	ResponseCode: "response_code",
	ResponseData: "response_data",
	CreatedBy:    "created_by",
	UpdatedBy:    "updated_by",
	CreatedAt:    "created_at",
	UpdatedAt:    "updated_at",
	DeletedAt:    "deleted_at",
	Remark:       "remark",
}

// NewSystemOperLogDao creates and returns a new DAO object for table data access.
func NewSystemOperLogDao(handlers ...gdb.ModelHandler) *SystemOperLogDao {
	return &SystemOperLogDao{
		group:    "default",
		table:    "system_oper_log",
		columns:  systemOperLogColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SystemOperLogDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SystemOperLogDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SystemOperLogDao) Columns() SystemOperLogColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SystemOperLogDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SystemOperLogDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *SystemOperLogDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
