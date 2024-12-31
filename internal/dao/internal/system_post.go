// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SystemPostDao is the data access object for table system_post.
type SystemPostDao struct {
	table   string            // table is the underlying table name of the DAO.
	group   string            // group is the database configuration group name of current DAO.
	columns SystemPostColumns // columns contains all the column names of Table for convenient usage.
}

// SystemPostColumns defines and stores column names for table system_post.
type SystemPostColumns struct {
	Id        string // 主键
	Name      string // 岗位名称
	Code      string // 岗位代码
	Sort      string // 排序
	Status    string // 状态 (1正常 2停用)
	CreatedBy string // 创建者
	UpdatedBy string // 更新者
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
	DeletedAt string // 删除时间
	Remark    string // 备注
}

// systemPostColumns holds the columns for table system_post.
var systemPostColumns = SystemPostColumns{
	Id:        "id",
	Name:      "name",
	Code:      "code",
	Sort:      "sort",
	Status:    "status",
	CreatedBy: "created_by",
	UpdatedBy: "updated_by",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
	Remark:    "remark",
}

// NewSystemPostDao creates and returns a new DAO object for table data access.
func NewSystemPostDao() *SystemPostDao {
	return &SystemPostDao{
		group:   "default",
		table:   "system_post",
		columns: systemPostColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SystemPostDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SystemPostDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SystemPostDao) Columns() SystemPostColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SystemPostDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SystemPostDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SystemPostDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
