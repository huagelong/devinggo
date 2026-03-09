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
	Id          any         // 主键
	StorageMode any         // 存储模式 (1 本地 2 阿里云 3 七牛云 4 腾讯云)
	OriginName  any         // 原文件名
	ObjectName  any         // 新文件名
	Hash        any         // 文件hash
	MimeType    any         // 资源类型
	StoragePath any         // 存储目录
	Suffix      any         // 文件后缀
	SizeByte    any         // 字节数
	SizeInfo    any         // 文件大小
	Url         any         // url地址
	CreatedBy   any         // 创建者
	UpdatedBy   any         // 更新者
	CreatedAt   *gtime.Time // 创建时间
	UpdatedAt   *gtime.Time // 更新时间
	DeletedAt   *gtime.Time // 删除时间
	Remark      any         // 备注
}
