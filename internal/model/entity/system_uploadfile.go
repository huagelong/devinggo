// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SystemUploadfile is the golang structure for table system_uploadfile.
type SystemUploadfile struct {
	Id          int64       `json:"id"          orm:"id"           description:""` //
	StorageMode int         `json:"storageMode" orm:"storage_mode" description:""` //
	OriginName  string      `json:"originName"  orm:"origin_name"  description:""` //
	ObjectName  string      `json:"objectName"  orm:"object_name"  description:""` //
	Hash        string      `json:"hash"        orm:"hash"         description:""` //
	MimeType    string      `json:"mimeType"    orm:"mime_type"    description:""` //
	StoragePath string      `json:"storagePath" orm:"storage_path" description:""` //
	Suffix      string      `json:"suffix"      orm:"suffix"       description:""` //
	SizeByte    int64       `json:"sizeByte"    orm:"size_byte"    description:""` //
	SizeInfo    string      `json:"sizeInfo"    orm:"size_info"    description:""` //
	Url         string      `json:"url"         orm:"url"          description:""` //
	CreatedBy   int64       `json:"createdBy"   orm:"created_by"   description:""` //
	UpdatedBy   int64       `json:"updatedBy"   orm:"updated_by"   description:""` //
	CreatedAt   *gtime.Time `json:"createdAt"   orm:"created_at"   description:""` //
	UpdatedAt   *gtime.Time `json:"updatedAt"   orm:"updated_at"   description:""` //
	DeletedAt   *gtime.Time `json:"deletedAt"   orm:"deleted_at"   description:""` //
	Remark      string      `json:"remark"      orm:"remark"       description:""` //
}
