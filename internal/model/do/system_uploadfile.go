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
	Id          interface{} //
	StorageMode interface{} //
	OriginName  interface{} //
	ObjectName  interface{} //
	Hash        interface{} //
	MimeType    interface{} //
	StoragePath interface{} //
	Suffix      interface{} //
	SizeByte    interface{} //
	SizeInfo    interface{} //
	Url         interface{} //
	CreatedBy   interface{} //
	UpdatedBy   interface{} //
	CreatedAt   *gtime.Time //
	UpdatedAt   *gtime.Time //
	DeletedAt   *gtime.Time //
	Remark      interface{} //
}
