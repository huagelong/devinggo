// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SystemOperLogDao is the data access object for table system_oper_log.
type SystemOperLogDao struct {
	table   string               // table is the underlying table name of the DAO.
	group   string               // group is the database configuration group name of current DAO.
	columns SystemOperLogColumns // columns contains all the column names of Table for convenient usage.
}

// SystemOperLogColumns defines and stores column names for table system_oper_log.
type SystemOperLogColumns struct {
	Id           string // 主键
	Username     string // 用户名
	Method       string // 请求方式
	Router       string // 请求路由
	ServiceName  string // 业务名称
	Ip           string // 请求IP地址
	IpLocation   string // IP所属地
	RequestData  string // 请求数据
	ResponseCode string // 响应状态码
	ResponseData string // 响应数据
	CreatedBy    string // 创建者
	UpdatedBy    string // 更新者
	CreatedAt    string // 创建时间
	UpdatedAt    string // 更新时间
	DeletedAt    string // 删除时间
	Remark       string // 备注
}

// systemOperLogColumns holds the columns for table system_oper_log.
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
func NewSystemOperLogDao() *SystemOperLogDao {
	return &SystemOperLogDao{
		group:   "default",
		table:   "system_oper_log",
		columns: systemOperLogColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SystemOperLogDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SystemOperLogDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SystemOperLogDao) Columns() SystemOperLogColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SystemOperLogDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SystemOperLogDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SystemOperLogDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
