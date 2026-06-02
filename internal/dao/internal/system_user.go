// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SystemUserDao is the data access object for the table system_user.
type SystemUserDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  SystemUserColumns  // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// SystemUserColumns defines and stores column names for the table system_user.
type SystemUserColumns struct {
	Id             string //
	Username       string //
	Password       string //
	UserType       string //
	Nickname       string //
	Phone          string //
	Email          string //
	Avatar         string //
	Signed         string //
	Dashboard      string //
	Status         string //
	LoginIp        string //
	LoginTime      string //
	BackendSetting string //
	CreatedBy      string //
	UpdatedBy      string //
	CreatedAt      string //
	UpdatedAt      string //
	DeletedAt      string //
	Remark         string //
}

// systemUserColumns holds the columns for the table system_user.
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
func NewSystemUserDao(handlers ...gdb.ModelHandler) *SystemUserDao {
	return &SystemUserDao{
		group:    "default",
		table:    "system_user",
		columns:  systemUserColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SystemUserDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SystemUserDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SystemUserDao) Columns() SystemUserColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SystemUserDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SystemUserDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *SystemUserDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
