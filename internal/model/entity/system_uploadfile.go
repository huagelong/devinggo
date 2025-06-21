// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SystemUploadfile is the golang structure for table system_uploadfile.
type SystemUploadfile struct {
	Id          int64       `json:"id"          orm:"id"           description:"主键"`                            // 主键
	StorageMode int         `json:"storageMode" orm:"storage_mode" description:"存储模式 (1 本地 2 阿里云 3 七牛云 4 腾讯云)"` // 存储模式 (1 本地 2 阿里云 3 七牛云 4 腾讯云)
	OriginName  string      `json:"originName"  orm:"origin_name"  description:"原文件名"`                          // 原文件名
	ObjectName  string      `json:"objectName"  orm:"object_name"  description:"新文件名"`                          // 新文件名
	Hash        string      `json:"hash"        orm:"hash"         description:"文件hash"`                        // 文件hash
	MimeType    string      `json:"mimeType"    orm:"mime_type"    description:"资源类型"`                          // 资源类型
	StoragePath string      `json:"storagePath" orm:"storage_path" description:"存储目录"`                          // 存储目录
	Suffix      string      `json:"suffix"      orm:"suffix"       description:"文件后缀"`                          // 文件后缀
	SizeByte    int64       `json:"sizeByte"    orm:"size_byte"    description:"字节数"`                           // 字节数
	SizeInfo    string      `json:"sizeInfo"    orm:"size_info"    description:"文件大小"`                          // 文件大小
	Url         string      `json:"url"         orm:"url"          description:"url地址"`                         // url地址
	CreatedBy   int64       `json:"createdBy"   orm:"created_by"   description:"创建者"`                           // 创建者
	UpdatedBy   int64       `json:"updatedBy"   orm:"updated_by"   description:"更新者"`                           // 更新者
	CreatedAt   *gtime.Time `json:"createdAt"   orm:"created_at"   description:"创建时间"`                          // 创建时间
	UpdatedAt   *gtime.Time `json:"updatedAt"   orm:"updated_at"   description:"更新时间"`                          // 更新时间
	DeletedAt   *gtime.Time `json:"deletedAt"   orm:"deleted_at"   description:"删除时间"`                          // 删除时间
	Remark      string      `json:"remark"      orm:"remark"       description:"备注"`                            // 备注
}
