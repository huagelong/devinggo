// Package system
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package system

import (
	"devinggo/modules/system/model"
	"devinggo/modules/system/model/page"
	"devinggo/modules/system/model/req"
	"devinggo/modules/system/model/res"
	"github.com/gogf/gf/v2/frame/g"
)

type IndexAttachmentReq struct {
	g.Meta `path:"/attachment/index" method:"get" tags:"附件" summary:"列表数据." x-permission:"system:attachment:index" `
	model.AuthorHeader
	model.PageListReq
	req.SystemUploadFileSearch
}

type IndexAttachmentRes struct {
	g.Meta `mime:"application/json"`
	page.PageRes
	Items []res.SystemUploadFile `json:"items"  dc:"list" `
}

type RecycleAttachmentReq struct {
	g.Meta `path:"/attachment/recycle" method:"get" tags:"附件" summary:"回收站列表数据." x-permission:"system:attachment:recycle" `
	model.AuthorHeader
	model.PageListReq
	req.SystemUploadFileSearch
}

type RecycleAttachmentRes struct {
	g.Meta `mime:"application/json"`
	page.PageRes
	Items []res.SystemUploadFile `json:"items"  dc:"list" `
}

type DeleteAttachmentReq struct {
	g.Meta `path:"/attachment/delete" method:"delete" tags:"附件" summary:"单个或批量删除" x-permission:"system:attachment:delete"`
	model.AuthorHeader
	Ids []int64 `json:"ids" dc:"ids" v:"min-length:1#Id不能为空"`
}

type DeleteAttachmentRes struct {
	g.Meta `mime:"application/json"`
}

type RealDeleteAttachmentReq struct {
	g.Meta `path:"/attachment/realDelete" method:"delete" tags:"附件" summary:"单个或批量真实删除 （清空回收站）." x-permission:"system:attachment:realDelete"`
	model.AuthorHeader
	Ids []int64 `json:"ids" dc:"ids" v:"min-length:1#Id不能为空"`
}

type RealDeleteAttachmentRes struct {
	g.Meta `mime:"application/json"`
}

type RecoveryAttachmentReq struct {
	g.Meta `path:"/attachment/recovery" method:"put" tags:"附件" summary:"单个或批量恢复在回收站的数据." x-permission:"system:attachment:recovery"`
	model.AuthorHeader
	Ids []int64 `json:"ids" dc:"ids" v:"min-length:1#Id不能为空"`
}

type RecoveryAttachmentRes struct {
	g.Meta `mime:"application/json"`
}
