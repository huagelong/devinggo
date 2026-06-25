// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SystemNoticeDao is the data access object for the table system_notice.
type SystemNoticeDao struct {
	table    string              // table is the underlying table name of the DAO.
	group    string              // group is the database configuration group name of the current DAO.
	columns  SystemNoticeColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler  // handlers for customized model modification.
}

// SystemNoticeColumns defines and stores column names for the table system_notice.
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

// systemNoticeColumns holds the columns for the table system_notice.
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
func NewSystemNoticeDao(handlers ...gdb.ModelHandler) *SystemNoticeDao {
	return &SystemNoticeDao{
		group:    "default",
		table:    "system_notice",
		columns:  systemNoticeColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SystemNoticeDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SystemNoticeDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SystemNoticeDao) Columns() SystemNoticeColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SystemNoticeDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SystemNoticeDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *SystemNoticeDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
