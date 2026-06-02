// Package req
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package req

import "github.com/gogf/gf/v2/net/ghttp"

type SystemUploadFileSearch struct {
	StorageMode string `json:"storage_mode"`
	OriginName  string `json:"origin_name"`
	StoragePath string `json:"storage_path"`
	MimeType    string `json:"mime_type"`
	MinDate     string `json:"minDate"`
	MaxDate     string `json:"maxDate"`
}

type FileUploadInput struct {
	File        *ghttp.UploadFile // 上传文件对象
	Name        string            // 自定义文件名称
	RandomName  bool              // 是否随机命名文件
	StorageMode int
}

type ChunkUploadInput struct {
	Total       int64             `json:"total" v:"required"`
	Index       int64             `json:"index" v:"required"`
	Hash        string            `json:"hash" v:"required"`
	Ext         string            `json:"ext" v:"required"`
	Type        string            `json:"type" v:"required"`
	Name        string            `json:"name" v:"required"`
	Size        int64             `json:"size" v:"required"`
	File        *ghttp.UploadFile `json:"package" type:"file"   dc:"pls upload chunk file"`
	RandomName  bool
	StorageMode int
}
