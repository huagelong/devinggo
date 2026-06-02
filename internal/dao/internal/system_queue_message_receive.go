// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SystemQueueMessageReceiveDao is the data access object for the table system_queue_message_receive.
type SystemQueueMessageReceiveDao struct {
	table    string                           // table is the underlying table name of the DAO.
	group    string                           // group is the database configuration group name of the current DAO.
	columns  SystemQueueMessageReceiveColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler               // handlers for customized model modification.
}

// SystemQueueMessageReceiveColumns defines and stores column names for the table system_queue_message_receive.
type SystemQueueMessageReceiveColumns struct {
	MessageId  string //
	UserId     string //
	ReadStatus string //
}

// systemQueueMessageReceiveColumns holds the columns for the table system_queue_message_receive.
var systemQueueMessageReceiveColumns = SystemQueueMessageReceiveColumns{
	MessageId:  "message_id",
	UserId:     "user_id",
	ReadStatus: "read_status",
}

// NewSystemQueueMessageReceiveDao creates and returns a new DAO object for table data access.
func NewSystemQueueMessageReceiveDao(handlers ...gdb.ModelHandler) *SystemQueueMessageReceiveDao {
	return &SystemQueueMessageReceiveDao{
		group:    "default",
		table:    "system_queue_message_receive",
		columns:  systemQueueMessageReceiveColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SystemQueueMessageReceiveDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SystemQueueMessageReceiveDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SystemQueueMessageReceiveDao) Columns() SystemQueueMessageReceiveColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SystemQueueMessageReceiveDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SystemQueueMessageReceiveDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *SystemQueueMessageReceiveDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
