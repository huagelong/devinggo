// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SystemUploadfileDao is the data access object for table system_uploadfile.
type SystemUploadfileDao struct {
	table   string                  // table is the underlying table name of the DAO.
	group   string                  // group is the database configuration group name of current DAO.
	columns SystemUploadfileColumns // columns contains all the column names of Table for convenient usage.
}

// SystemUploadfileColumns defines and stores column names for table system_uploadfile.
type SystemUploadfileColumns struct {
	Id          string //
	StorageMode string //
	OriginName  string //
	ObjectName  string //
	Hash        string //
	MimeType    string //
	StoragePath string //
	Suffix      string //
	SizeByte    string //
	SizeInfo    string //
	Url         string //
	CreatedBy   string //
	UpdatedBy   string //
	CreatedAt   string //
	UpdatedAt   string //
	DeletedAt   string //
	Remark      string //
}

// systemUploadfileColumns holds the columns for table system_uploadfile.
var systemUploadfileColumns = SystemUploadfileColumns{
	Id:          "id",
	StorageMode: "storage_mode",
	OriginName:  "origin_name",
	ObjectName:  "object_name",
	Hash:        "hash",
	MimeType:    "mime_type",
	StoragePath: "storage_path",
	Suffix:      "suffix",
	SizeByte:    "size_byte",
	SizeInfo:    "size_info",
	Url:         "url",
	CreatedBy:   "created_by",
	UpdatedBy:   "updated_by",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
	DeletedAt:   "deleted_at",
	Remark:      "remark",
}

// NewSystemUploadfileDao creates and returns a new DAO object for table data access.
func NewSystemUploadfileDao() *SystemUploadfileDao {
	return &SystemUploadfileDao{
		group:   "default",
		table:   "system_uploadfile",
		columns: systemUploadfileColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SystemUploadfileDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SystemUploadfileDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SystemUploadfileDao) Columns() SystemUploadfileColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SystemUploadfileDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SystemUploadfileDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SystemUploadfileDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
