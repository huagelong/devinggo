// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SettingGenerateTablesDao is the data access object for table setting_generate_tables.
type SettingGenerateTablesDao struct {
	table   string                       // table is the underlying table name of the DAO.
	group   string                       // group is the database configuration group name of current DAO.
	columns SettingGenerateTablesColumns // columns contains all the column names of Table for convenient usage.
}

// SettingGenerateTablesColumns defines and stores column names for table setting_generate_tables.
type SettingGenerateTablesColumns struct {
	Id            string // 主键
	TableName     string // 表名称
	TableComment  string // 表注释
	ModuleName    string // 所属模块
	Namespace     string // 命名空间
	MenuName      string // 生成菜单名
	BelongMenuId  string // 所属菜单
	PackageName   string // controller,api包名
	Type          string // 生成类型，single 单表CRUD，tree 树表CRUD，parent_sub父子表CRUD
	GenerateType  string // 1 压缩包下载 2 生成到模块
	GenerateMenus string // 生成菜单列表
	BuildMenu     string // 是否构建菜单
	ComponentType string // 组件显示方式
	Options       string // 其他业务选项
	CreatedBy     string // 创建者
	UpdatedBy     string // 更新者
	CreatedAt     string // 创建时间
	UpdatedAt     string // 更新时间
	Remark        string // 备注
	Source        string // db连接群组
}

// settingGenerateTablesColumns holds the columns for table setting_generate_tables.
var settingGenerateTablesColumns = SettingGenerateTablesColumns{
	Id:            "id",
	TableName:     "table_name",
	TableComment:  "table_comment",
	ModuleName:    "module_name",
	Namespace:     "namespace",
	MenuName:      "menu_name",
	BelongMenuId:  "belong_menu_id",
	PackageName:   "package_name",
	Type:          "type",
	GenerateType:  "generate_type",
	GenerateMenus: "generate_menus",
	BuildMenu:     "build_menu",
	ComponentType: "component_type",
	Options:       "options",
	CreatedBy:     "created_by",
	UpdatedBy:     "updated_by",
	CreatedAt:     "created_at",
	UpdatedAt:     "updated_at",
	Remark:        "remark",
	Source:        "source",
}

// NewSettingGenerateTablesDao creates and returns a new DAO object for table data access.
func NewSettingGenerateTablesDao() *SettingGenerateTablesDao {
	return &SettingGenerateTablesDao{
		group:   "default",
		table:   "setting_generate_tables",
		columns: settingGenerateTablesColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SettingGenerateTablesDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SettingGenerateTablesDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SettingGenerateTablesDao) Columns() SettingGenerateTablesColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SettingGenerateTablesDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SettingGenerateTablesDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SettingGenerateTablesDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
