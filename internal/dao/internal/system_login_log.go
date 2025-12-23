// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SystemLoginLogDao is the data access object for the table system_login_log.
type SystemLoginLogDao struct {
	table    string                // table is the underlying table name of the DAO.
	group    string                // group is the database configuration group name of the current DAO.
	columns  SystemLoginLogColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler    // handlers for customized model modification.
}

// SystemLoginLogColumns defines and stores column names for the table system_login_log.
type SystemLoginLogColumns struct {
	Id         string // 主键
	Username   string // 用户名
	Ip         string // 登录IP地址
	IpLocation string // IP所属地
	Os         string // 操作系统
	Browser    string // 浏览器
	Status     string // 登录状态 (1成功 2失败)
	Message    string // 提示消息
	LoginTime  string // 登录时间
	Remark     string // 备注
}

// systemLoginLogColumns holds the columns for the table system_login_log.
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
func NewSystemLoginLogDao(handlers ...gdb.ModelHandler) *SystemLoginLogDao {
	return &SystemLoginLogDao{
		group:    "default",
		table:    "system_login_log",
		columns:  systemLoginLogColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SystemLoginLogDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SystemLoginLogDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SystemLoginLogDao) Columns() SystemLoginLogColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SystemLoginLogDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SystemLoginLogDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *SystemLoginLogDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
