// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// CmsTopicDao is the data access object for table cms_topic.
type CmsTopicDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of current DAO.
	columns CmsTopicColumns // columns contains all the column names of Table for convenient usage.
}

// CmsTopicColumns defines and stores column names for table cms_topic.
type CmsTopicColumns struct {
	Id        string //
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
	DeletedAt string // 删除时间
	Title     string // 标题
	Content   string // 内容
	Status    string // 状态 (1正常 2停用)
	Remark    string // 备注
}

// cmsTopicColumns holds the columns for table cms_topic.
var cmsTopicColumns = CmsTopicColumns{
	Id:        "id",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
	Title:     "title",
	Content:   "content",
	Status:    "status",
	Remark:    "remark",
}

// NewCmsTopicDao creates and returns a new DAO object for table data access.
func NewCmsTopicDao() *CmsTopicDao {
	return &CmsTopicDao{
		group:   "default",
		table:   "cms_topic",
		columns: cmsTopicColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *CmsTopicDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *CmsTopicDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *CmsTopicDao) Columns() CmsTopicColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *CmsTopicDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *CmsTopicDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *CmsTopicDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
