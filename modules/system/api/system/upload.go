// Package system
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package system

import (
	"devinggo/modules/system/model"
	"devinggo/modules/system/model/req"
	"devinggo/modules/system/model/res"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type UploadFileReq struct {
	g.Meta `path:"/uploadFile" method:"post" mime:"multipart/form-data" tags:"文件上传" summary:"上传文件" x-permission:"system:uploadFile" `
	model.AuthorHeader
	File *ghttp.UploadFile `json:"file" type:"file"  dc:"pls upload file"`
}

type UploadFileRes struct {
	g.Meta `mime:"application/json"`
	Data   res.SystemUploadFileRes `json:"data"`
}

type UploadImageReq struct {
	g.Meta `path:"/uploadImage" method:"post" mime:"multipart/form-data" tags:"文件上传" summary:"上传图片" x-permission:"system:uploadImage" `
	model.AuthorHeader
	File *ghttp.UploadFile `json:"image" type:"file"  dc:"pls upload image file"`
}

type UploadImageRes struct {
	g.Meta `mime:"application/json"`
	Data   res.SystemUploadFileRes `json:"data"`
}

type ChunkUploadReq struct {
	g.Meta `path:"/chunkUpload" method:"post" mime:"multipart/form-data" tags:"文件上传" summary:"分片上传文件" x-permission:"system:chunkUpload" `
	model.AuthorHeader
	req.ChunkUploadInput
}

type ChunkUploadRes struct {
	g.Meta `mime:"application/json"`
	Data   res.SystemUploadFileRes `json:"data"`
}

type SaveNetworkImageReq struct {
	g.Meta `path:"/saveNetworkImage" method:"post" tags:"文件上传" summary:"保存网络图片" x-permission:"system:saveNetworkImage" `
	model.AuthorHeader
	Url string `json:"url" v:"required|url" dc:"pls input url"`
}

type SaveNetworkImageRes struct {
	g.Meta `mime:"application/json"`
	Data   res.SystemUploadFileRes `json:"data"`
}

type GetFileInfoByIdReq struct {
	g.Meta `path:"/getFileInfoById" method:"get" tags:"文件上传" summary:"通过ID获取文件信息." x-permission:"system:getFileInfoById" `
	model.AuthorHeader
	Id int64 `json:"id" v:"required" dc:"pls input id"`
}

type GetFileInfoByIdRes struct {
	g.Meta `mime:"application/json"`
	Data   res.SystemUploadFile `json:"data"`
}

type GetFileInfoByHashReq struct {
	g.Meta `path:"/getFileInfoByHash" method:"get" tags:"文件上传" summary:"通过ID获取文件信息." x-permission:"system:getFileInfoByHash" `
	model.AuthorHeader
	Hash string `json:"hash" v:"required" dc:"pls input hash"`
}

type GetFileInfoByHashRes struct {
	g.Meta `mime:"application/json"`
	Data   res.SystemUploadFile `json:"data"`
}

type DownloadByIdReq struct {
	g.Meta `path:"/downloadById" method:"get" tags:"文件上传" summary:"根据id下载文件." x-permission:"system:downloadById" `
	model.AuthorHeader
	Id int64 `json:"id" v:"required" dc:"pls input id"`
}

type DownloadByIdRes struct {
	g.Meta `mime:"application/json"`
}

type DownloadByHashReq struct {
	g.Meta `path:"/downloadByHash" method:"get" tags:"文件上传" summary:"根据hash下载文件." x-permission:"system:downloadByHash" `
	model.AuthorHeader
	Hash string `json:"hash" v:"required" dc:"pls input hash"`
}

type DownloadByHashRes struct {
	g.Meta `mime:"application/json"`
}

type ShowFileReq struct {
	g.Meta `path:"/showFile/{Hash}" method:"get" tags:"文件上传" summary:"输出图片、文件." x-exceptAuth:"true" x-exceptLogin:"true" x-permission:"system:showFile" `
	Hash string `json:"hash" v:"required" dc:"pls input hash"`
}

type ShowFileRes struct {
	g.Meta `mime:"application/json"`
}
