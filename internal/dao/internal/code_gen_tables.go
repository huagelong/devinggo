// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// CodeGenTablesDao is the data access object for the table code_gen_tables.
type CodeGenTablesDao struct {
	table    string               // table is the underlying table name of the DAO.
	group    string               // group is the database configuration group name of the current DAO.
	columns  CodeGenTablesColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler   // handlers for customized model modification.
}

// CodeGenTablesColumns defines and stores column names for the table code_gen_tables.
type CodeGenTablesColumns struct {
	Id            string //
	TableName     string // 表名称
	TableComment  string // 表描述
	Remark        string // 备注信息
	ModuleName    string // 所属模块
	BelongMenuId  string // 所属菜单ID
	Type          string // 生成类型: single=单表, tree=树表
	MenuName      string // 菜单名称
	ComponentType string // 组件类型: 1=模态框, 2=抽屉, 3=Tag页
	TplType       string // 模板类型: default
	TreeId        string // 树表主ID字段
	TreeParentId  string // 树表父ID字段
	TreeName      string // 树表显示名称字段
	TagId         string // Tag页ID
	TagName       string // Tag页名称
	TagViewName   string // Tag页显示字段
	GenerateMenus string // 生成的菜单按钮
	Options       string // 扩展配置
	CreatedAt     string // 创建时间
	UpdatedAt     string // 更新时间
	DeletedAt     string // 删除时间
	CreatedBy     string // 创建者ID
	UpdatedBy     string // 更新者ID
	Status        string // 状态: 1=正常, 0=停用
	Sort          string // 排序
}

// codeGenTablesColumns holds the columns for the table code_gen_tables.
var codeGenTablesColumns = CodeGenTablesColumns{
	Id:            "id",
	TableName:     "table_name",
	TableComment:  "table_comment",
	Remark:        "remark",
	ModuleName:    "module_name",
	BelongMenuId:  "belong_menu_id",
	Type:          "type",
	MenuName:      "menu_name",
	ComponentType: "component_type",
	TplType:       "tpl_type",
	TreeId:        "tree_id",
	TreeParentId:  "tree_parent_id",
	TreeName:      "tree_name",
	TagId:         "tag_id",
	TagName:       "tag_name",
	TagViewName:   "tag_view_name",
	GenerateMenus: "generate_menus",
	Options:       "options",
	CreatedAt:     "created_at",
	UpdatedAt:     "updated_at",
	DeletedAt:     "deleted_at",
	CreatedBy:     "created_by",
	UpdatedBy:     "updated_by",
	Status:        "status",
	Sort:          "sort",
}

// NewCodeGenTablesDao creates and returns a new DAO object for table data access.
func NewCodeGenTablesDao(handlers ...gdb.ModelHandler) *CodeGenTablesDao {
	return &CodeGenTablesDao{
		group:    "default",
		table:    "code_gen_tables",
		columns:  codeGenTablesColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *CodeGenTablesDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *CodeGenTablesDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *CodeGenTablesDao) Columns() CodeGenTablesColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *CodeGenTablesDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *CodeGenTablesDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *CodeGenTablesDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
