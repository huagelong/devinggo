// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SystemMenuDao is the data access object for table system_menu.
type SystemMenuDao struct {
	table   string            // table is the underlying table name of the DAO.
	group   string            // group is the database configuration group name of current DAO.
	columns SystemMenuColumns // columns contains all the column names of Table for convenient usage.
}

// SystemMenuColumns defines and stores column names for table system_menu.
type SystemMenuColumns struct {
	Id        string // 主键
	ParentId  string // 父ID
	Level     string // 组级集合
	Name      string // 菜单名称
	Code      string // 菜单标识代码
	Icon      string // 菜单图标
	Route     string // 路由地址
	Component string // 组件路径
	Redirect  string // 跳转地址
	IsHidden  string // 是否隐藏 (1是 2否)
	Type      string // 菜单类型, (M菜单 B按钮 L链接 I iframe)
	Status    string // 状态 (1正常 2停用)
	Sort      string // 排序
	CreatedBy string // 创建者
	UpdatedBy string // 更新者
	CreatedAt string //
	UpdatedAt string //
	DeletedAt string // 删除时间
	Remark    string // 备注
}

// systemMenuColumns holds the columns for table system_menu.
var systemMenuColumns = SystemMenuColumns{
	Id:        "id",
	ParentId:  "parent_id",
	Level:     "level",
	Name:      "name",
	Code:      "code",
	Icon:      "icon",
	Route:     "route",
	Component: "component",
	Redirect:  "redirect",
	IsHidden:  "is_hidden",
	Type:      "type",
	Status:    "status",
	Sort:      "sort",
	CreatedBy: "created_by",
	UpdatedBy: "updated_by",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
	Remark:    "remark",
}

// NewSystemMenuDao creates and returns a new DAO object for table data access.
func NewSystemMenuDao() *SystemMenuDao {
	return &SystemMenuDao{
		group:   "default",
		table:   "system_menu",
		columns: systemMenuColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SystemMenuDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SystemMenuDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SystemMenuDao) Columns() SystemMenuColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SystemMenuDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SystemMenuDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SystemMenuDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
