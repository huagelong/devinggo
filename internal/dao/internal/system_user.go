// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SystemUserDao is the data access object for table system_user.
type SystemUserDao struct {
	table   string            // table is the underlying table name of the DAO.
	group   string            // group is the database configuration group name of current DAO.
	columns SystemUserColumns // columns contains all the column names of Table for convenient usage.
}

// SystemUserColumns defines and stores column names for table system_user.
type SystemUserColumns struct {
	Id             string // 用户ID，主键
	Username       string // 用户名
	Password       string // 密码
	UserType       string // 用户类型：(100系统用户)
	Nickname       string // 用户昵称
	Phone          string // 手机
	Email          string // 用户邮箱
	Avatar         string // 用户头像
	Signed         string // 个人签名
	Dashboard      string // 后台首页类型
	Status         string // 状态 (1正常 2停用)
	LoginIp        string // 最后登陆IP
	LoginTime      string // 最后登陆时间
	BackendSetting string // 后台设置数据
	CreatedBy      string // 创建者
	UpdatedBy      string // 更新者
	CreatedAt      string // 创建时间
	UpdatedAt      string // 更新时间
	DeletedAt      string // 删除时间
	Remark         string // 备注
}

// systemUserColumns holds the columns for table system_user.
var systemUserColumns = SystemUserColumns{
	Id:             "id",
	Username:       "username",
	Password:       "password",
	UserType:       "user_type",
	Nickname:       "nickname",
	Phone:          "phone",
	Email:          "email",
	Avatar:         "avatar",
	Signed:         "signed",
	Dashboard:      "dashboard",
	Status:         "status",
	LoginIp:        "login_ip",
	LoginTime:      "login_time",
	BackendSetting: "backend_setting",
	CreatedBy:      "created_by",
	UpdatedBy:      "updated_by",
	CreatedAt:      "created_at",
	UpdatedAt:      "updated_at",
	DeletedAt:      "deleted_at",
	Remark:         "remark",
}

// NewSystemUserDao creates and returns a new DAO object for table data access.
func NewSystemUserDao() *SystemUserDao {
	return &SystemUserDao{
		group:   "default",
		table:   "system_user",
		columns: systemUserColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SystemUserDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SystemUserDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SystemUserDao) Columns() SystemUserColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SystemUserDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SystemUserDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SystemUserDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
