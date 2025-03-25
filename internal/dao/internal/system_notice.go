// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SystemNoticeDao is the data access object for table system_notice.
type SystemNoticeDao struct {
	table   string              // table is the underlying table name of the DAO.
	group   string              // group is the database configuration group name of current DAO.
	columns SystemNoticeColumns // columns contains all the column names of Table for convenient usage.
}

// SystemNoticeColumns defines and stores column names for table system_notice.
type SystemNoticeColumns struct {
	Id           string //
	MessageId    string //
	Title        string //
	Type         string //
	Content      string //
	CreatedBy    string //
	UpdatedBy    string //
	CreatedAt    string //
	UpdatedAt    string //
	DeletedAt    string //
	Remark       string //
	ReceiveUsers string //
}

// systemNoticeColumns holds the columns for table system_notice.
var systemNoticeColumns = SystemNoticeColumns{
	Id:           "id",
	MessageId:    "message_id",
	Title:        "title",
	Type:         "type",
	Content:      "content",
	CreatedBy:    "created_by",
	UpdatedBy:    "updated_by",
	CreatedAt:    "created_at",
	UpdatedAt:    "updated_at",
	DeletedAt:    "deleted_at",
	Remark:       "remark",
	ReceiveUsers: "receive_users",
}

// NewSystemNoticeDao creates and returns a new DAO object for table data access.
func NewSystemNoticeDao() *SystemNoticeDao {
	return &SystemNoticeDao{
		group:   "default",
		table:   "system_notice",
		columns: systemNoticeColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SystemNoticeDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SystemNoticeDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SystemNoticeDao) Columns() SystemNoticeColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SystemNoticeDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SystemNoticeDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SystemNoticeDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
