// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// CodeGenButtonsDao is the data access object for the table code_gen_buttons.
type CodeGenButtonsDao struct {
	table    string                // table is the underlying table name of the DAO.
	group    string                // group is the database configuration group name of the current DAO.
	columns  CodeGenButtonsColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler    // handlers for customized model modification.
}

// CodeGenButtonsColumns defines and stores column names for the table code_gen_buttons.
type CodeGenButtonsColumns struct {
	Id            string //
	TableId       string // 所属表ID
	ButtonCode    string // 按钮编码
	ButtonName    string // 按钮名称
	ButtonComment string // 按钮描述
	IsShow        string // 是否显示: 1=显示, 2=隐藏
	Sort          string // 排序
	CreatedAt     string // 创建时间
	UpdatedAt     string // 更新时间
}

// codeGenButtonsColumns holds the columns for the table code_gen_buttons.
var codeGenButtonsColumns = CodeGenButtonsColumns{
	Id:            "id",
	TableId:       "table_id",
	ButtonCode:    "button_code",
	ButtonName:    "button_name",
	ButtonComment: "button_comment",
	IsShow:        "is_show",
	Sort:          "sort",
	CreatedAt:     "created_at",
	UpdatedAt:     "updated_at",
}

// NewCodeGenButtonsDao creates and returns a new DAO object for table data access.
func NewCodeGenButtonsDao(handlers ...gdb.ModelHandler) *CodeGenButtonsDao {
	return &CodeGenButtonsDao{
		group:    "default",
		table:    "code_gen_buttons",
		columns:  codeGenButtonsColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *CodeGenButtonsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *CodeGenButtonsDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *CodeGenButtonsDao) Columns() CodeGenButtonsColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *CodeGenButtonsDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *CodeGenButtonsDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *CodeGenButtonsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
