// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SystemUploadfile is the golang structure of table system_uploadfile for DAO operations like Where/Data.
type SystemUploadfile struct {
	g.Meta      `orm:"table:system_uploadfile, do:true"`
	Id          any         //
	StorageMode any         //
	OriginName  any         //
	ObjectName  any         //
	Hash        any         //
	MimeType    any         //
	StoragePath any         //
	Suffix      any         //
	SizeByte    any         //
	SizeInfo    any         //
	Url         any         //
	CreatedBy   any         //
	UpdatedBy   any         //
	CreatedAt   *gtime.Time //
	UpdatedAt   *gtime.Time //
	DeletedAt   *gtime.Time //
	Remark      any         //
}
