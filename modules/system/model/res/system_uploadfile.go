// Package res
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package res

import (
	"github.com/gogf/gf/v2/os/gtime"
)

type SystemUploadFileRes struct {
	StorageMode int    `json:"storage_mode" description:"存储模式 (1 本地 2 阿里云 3 七牛云 4 腾讯云)"` // 存储模式 (1 本地 2 阿里云 3 七牛云 4 腾讯云)
	OriginName  string `json:"origin_name"  description:"原文件名"`                          // 原文件名
	ObjectName  string `json:"object_name"  description:"新文件名"`                          // 新文件名
	Hash        string `json:"hash"         description:"文件hash"`                        // 文件hash
	MimeType    string `json:"mime_type"    description:"资源类型"`                          // 资源类型
	StoragePath string `json:"storage_path" description:"存储目录"`                          // 存储目录
	Suffix      string `json:"suffix"       description:"文件后缀"`                          // 文件后缀
	SizeByte    int64  `json:"size_byte"    description:"字节数"`                           // 字节数
	SizeInfo    string `json:"size_info"    description:"文件大小"`                          // 文件大小
	Url         string `json:"url"          description:"url地址"`                         // url地址
	LocalPath   string `json:"local_path"   description:"本地路径"`                          // 本地路径
	Chunk       int64  `json:"chunk" `
}

type SystemUploadFile struct {
	Id          int64       `json:"id"           description:"主键"`                            // 主键
	StorageMode int         `json:"storage_mode" description:"存储模式 (1 本地 2 阿里云 3 七牛云 4 腾讯云)"` // 存储模式 (1 本地 2 阿里云 3 七牛云 4 腾讯云)
	OriginName  string      `json:"origin_name"  description:"原文件名"`                          // 原文件名
	ObjectName  string      `json:"object_name"  description:"新文件名"`                          // 新文件名
	Hash        string      `json:"hash"         description:"文件hash"`                        // 文件hash
	MimeType    string      `json:"mime_type"    description:"资源类型"`                          // 资源类型
	StoragePath string      `json:"storage_path" description:"存储目录"`                          // 存储目录
	Suffix      string      `json:"suffix"       description:"文件后缀"`                          // 文件后缀
	SizeByte    int64       `json:"size_byte"    description:"字节数"`                           // 字节数
	SizeInfo    string      `json:"size_info"    description:"文件大小"`                          // 文件大小
	Url         string      `json:"url"          description:"url地址"`                         // url地址
	CreatedBy   int64       `json:"created_by"   description:"创建者"`                           // 创建者
	UpdatedBy   int64       `json:"updated_by"   description:"更新者"`                           // 更新者
	CreatedAt   *gtime.Time `json:"created_at"   description:"创建时间"`                          // 创建时间
	UpdatedAt   *gtime.Time `json:"updated_at"   description:"更新时间"`                          // 更新时间
	Remark      string      `json:"remark"       description:"备注"`                            // 备注
}
