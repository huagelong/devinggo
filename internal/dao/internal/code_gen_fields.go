// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// CodeGenFieldsDao is the data access object for the table code_gen_fields.
type CodeGenFieldsDao struct {
	table    string               // table is the underlying table name of the DAO.
	group    string               // group is the database configuration group name of the current DAO.
	columns  CodeGenFieldsColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler   // handlers for customized model modification.
}

// CodeGenFieldsColumns defines and stores column names for the table code_gen_fields.
type CodeGenFieldsColumns struct {
	Id            string //
	TableId       string // 所属表ID
	ColumnName    string // 字段名称
	ColumnComment string // 字段描述
	ColumnType    string // 物理类型
	DataType      string // 数据类型
	IsNullable    string // 是否可空
	Sort          string // 排序
	IsRequired    string // 必填: 1=是, 2=否
	IsInsert      string // 新增: 1=是, 2=否
	IsEdit        string // 编辑: 1=是, 2=否
	IsList        string // 列表显示: 1=是, 2=否
	IsQuery       string // 查询: 1=是, 2=否
	IsSort        string // 排序: 1=是, 2=否
	QueryType     string // 查询方式
	ViewType      string // 页面控件
	DictType      string // 数据字典
	AllowRoles    string // 允许角色
	Options       string // 组件扩展配置
	CreatedAt     string // 创建时间
	UpdatedAt     string // 更新时间
}

// codeGenFieldsColumns holds the columns for the table code_gen_fields.
var codeGenFieldsColumns = CodeGenFieldsColumns{
	Id:            "id",
	TableId:       "table_id",
	ColumnName:    "column_name",
	ColumnComment: "column_comment",
	ColumnType:    "column_type",
	DataType:      "data_type",
	IsNullable:    "is_nullable",
	Sort:          "sort",
	IsRequired:    "is_required",
	IsInsert:      "is_insert",
	IsEdit:        "is_edit",
	IsList:        "is_list",
	IsQuery:       "is_query",
	IsSort:        "is_sort",
	QueryType:     "query_type",
	ViewType:      "view_type",
	DictType:      "dict_type",
	AllowRoles:    "allow_roles",
	Options:       "options",
	CreatedAt:     "created_at",
	UpdatedAt:     "updated_at",
}

// NewCodeGenFieldsDao creates and returns a new DAO object for table data access.
func NewCodeGenFieldsDao(handlers ...gdb.ModelHandler) *CodeGenFieldsDao {
	return &CodeGenFieldsDao{
		group:    "default",
		table:    "code_gen_fields",
		columns:  codeGenFieldsColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *CodeGenFieldsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *CodeGenFieldsDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *CodeGenFieldsDao) Columns() CodeGenFieldsColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *CodeGenFieldsDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *CodeGenFieldsDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *CodeGenFieldsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
