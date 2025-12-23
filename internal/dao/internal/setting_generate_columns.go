// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SettingGenerateColumnsDao is the data access object for the table setting_generate_columns.
type SettingGenerateColumnsDao struct {
	table    string                        // table is the underlying table name of the DAO.
	group    string                        // group is the database configuration group name of the current DAO.
	columns  SettingGenerateColumnsColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler            // handlers for customized model modification.
}

// SettingGenerateColumnsColumns defines and stores column names for the table setting_generate_columns.
type SettingGenerateColumnsColumns struct {
	Id            string // 主键
	TableId       string // 所属表ID
	ColumnName    string // 字段名称
	ColumnComment string // 字段注释
	ColumnType    string // 字段类型
	IsPk          string // 1 非主键 2 主键
	IsRequired    string // 1 非必填 2 必填
	IsInsert      string // 1 非插入字段 2 插入字段
	IsEdit        string // 1 非编辑字段 2 编辑字段
	IsList        string // 1 非列表显示字段 2 列表显示字段
	IsQuery       string // 1 非查询字段 2 查询字段
	IsSort        string // 1 不排序 2 排序字段
	QueryType     string // 查询方式 eq 等于, neq 不等于, gt 大于, lt 小于, like 范围
	ViewType      string // 页面控件，text, textarea, password, select, checkbox, radio, date, upload, ma-upload（封装的上传控件）
	DictType      string // 字典类型
	AllowRoles    string // 允许查看该字段的角色
	Options       string // 字段其他设置
	Extra         string // 字段扩展信息
	Sort          string // 排序
	CreatedBy     string // 创建者
	UpdatedBy     string // 更新者
	CreatedAt     string // 创建时间
	UpdatedAt     string // 更新时间
	Remark        string // 备注
}

// settingGenerateColumnsColumns holds the columns for the table setting_generate_columns.
var settingGenerateColumnsColumns = SettingGenerateColumnsColumns{
	Id:            "id",
	TableId:       "table_id",
	ColumnName:    "column_name",
	ColumnComment: "column_comment",
	ColumnType:    "column_type",
	IsPk:          "is_pk",
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
	Extra:         "extra",
	Sort:          "sort",
	CreatedBy:     "created_by",
	UpdatedBy:     "updated_by",
	CreatedAt:     "created_at",
	UpdatedAt:     "updated_at",
	Remark:        "remark",
}

// NewSettingGenerateColumnsDao creates and returns a new DAO object for table data access.
func NewSettingGenerateColumnsDao(handlers ...gdb.ModelHandler) *SettingGenerateColumnsDao {
	return &SettingGenerateColumnsDao{
		group:    "default",
		table:    "setting_generate_columns",
		columns:  settingGenerateColumnsColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SettingGenerateColumnsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SettingGenerateColumnsDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SettingGenerateColumnsDao) Columns() SettingGenerateColumnsColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SettingGenerateColumnsDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SettingGenerateColumnsDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *SettingGenerateColumnsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
