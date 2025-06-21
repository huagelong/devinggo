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
	Id          interface{} // 主键
	StorageMode interface{} // 存储模式 (1 本地 2 阿里云 3 七牛云 4 腾讯云)
	OriginName  interface{} // 原文件名
	ObjectName  interface{} // 新文件名
	Hash        interface{} // 文件hash
	MimeType    interface{} // 资源类型
	StoragePath interface{} // 存储目录
	Suffix      interface{} // 文件后缀
	SizeByte    interface{} // 字节数
	SizeInfo    interface{} // 文件大小
	Url         interface{} // url地址
	CreatedBy   interface{} // 创建者
	UpdatedBy   interface{} // 更新者
	CreatedAt   *gtime.Time // 创建时间
	UpdatedAt   *gtime.Time // 更新时间
	DeletedAt   *gtime.Time // 删除时间
	Remark      interface{} // 备注
}
