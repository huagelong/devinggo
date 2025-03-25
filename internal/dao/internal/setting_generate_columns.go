// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SettingGenerateColumnsDao is the data access object for table setting_generate_columns.
type SettingGenerateColumnsDao struct {
	table   string                        // table is the underlying table name of the DAO.
	group   string                        // group is the database configuration group name of current DAO.
	columns SettingGenerateColumnsColumns // columns contains all the column names of Table for convenient usage.
}

// SettingGenerateColumnsColumns defines and stores column names for table setting_generate_columns.
type SettingGenerateColumnsColumns struct {
	Id            string //
	TableId       string //
	ColumnName    string //
	ColumnComment string //
	ColumnType    string //
	IsPk          string //
	IsRequired    string //
	IsInsert      string //
	IsEdit        string //
	IsList        string //
	IsQuery       string //
	IsSort        string //
	QueryType     string //
	ViewType      string //
	DictType      string //
	AllowRoles    string //
	Options       string //
	Extra         string //
	Sort          string //
	CreatedBy     string //
	UpdatedBy     string //
	CreatedAt     string //
	UpdatedAt     string //
	Remark        string //
}

// settingGenerateColumnsColumns holds the columns for table setting_generate_columns.
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
func NewSettingGenerateColumnsDao() *SettingGenerateColumnsDao {
	return &SettingGenerateColumnsDao{
		group:   "default",
		table:   "setting_generate_columns",
		columns: settingGenerateColumnsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SettingGenerateColumnsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SettingGenerateColumnsDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SettingGenerateColumnsDao) Columns() SettingGenerateColumnsColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SettingGenerateColumnsDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SettingGenerateColumnsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SettingGenerateColumnsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
