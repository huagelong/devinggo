// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SystemRoleDao is the data access object for the table system_role.
type SystemRoleDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  SystemRoleColumns  // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// SystemRoleColumns defines and stores column names for the table system_role.
type SystemRoleColumns struct {
	Id        string // 主键
	Name      string // 角色名称
	Code      string // 角色代码
	DataScope string // 数据范围（1：全部数据权限 2：自定义数据权限 3：本部门数据权限 4：本部门及以下数据权限 5：本人数据权限）
	Status    string // 状态 (1正常 2停用)
	Sort      string // 排序
	CreatedBy string // 创建者
	UpdatedBy string // 更新者
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
	DeletedAt string // 删除时间
	Remark    string // 备注
}

// systemRoleColumns holds the columns for the table system_role.
var systemRoleColumns = SystemRoleColumns{
	Id:        "id",
	Name:      "name",
	Code:      "code",
	DataScope: "data_scope",
	Status:    "status",
	Sort:      "sort",
	CreatedBy: "created_by",
	UpdatedBy: "updated_by",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
	Remark:    "remark",
}

// NewSystemRoleDao creates and returns a new DAO object for table data access.
func NewSystemRoleDao(handlers ...gdb.ModelHandler) *SystemRoleDao {
	return &SystemRoleDao{
		group:    "default",
		table:    "system_role",
		columns:  systemRoleColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SystemRoleDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SystemRoleDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SystemRoleDao) Columns() SystemRoleColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SystemRoleDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SystemRoleDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *SystemRoleDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
