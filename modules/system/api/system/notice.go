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

type IndexNoticeReq struct {
	g.Meta `path:"/notice/index" method:"get" tags:"通知" summary:"通知列表." x-permission:"system:notice:index" `
	model.AuthorHeader
	model.PageListReq
	req.SystemNoticeSearch
}

type IndexNoticeRes struct {
	g.Meta `mime:"application/json"`
	page.PageRes
	Items []res.SystemNotice `json:"items"  dc:"notice list" `
}

type RecycleNoticeReq struct {
	g.Meta `path:"/notice/recycle" method:"get" tags:"通知" summary:"回收站通知列表." x-permission:"system:notice:recycle" `
	model.AuthorHeader
	model.PageListReq
	req.SystemNoticeSearch
}

type RecycleNoticeRes struct {
	g.Meta `mime:"application/json"`
	page.PageRes
	Items []res.SystemNotice `json:"items"  dc:"notice list" `
}

type SaveNoticeReq struct {
	g.Meta `path:"/notice/save" method:"post" tags:"通知" summary:"新增通知." x-permission:"system:notice:save"`
	model.AuthorHeader
	req.SystemNoticeSave
}

type SaveNoticeRes struct {
	g.Meta `mime:"application/json"`
	Id     int64 `json:"id" dc:"通知 id"`
}

type ReadNoticeReq struct {
	g.Meta `path:"/notice/read/{Id}" method:"get" tags:"通知" summary:"更新通知." x-permission:"system:notice:read"`
	model.AuthorHeader
	Id int64 `json:"id" dc:"通知 id" v:"required|min:1#通知Id不能为空"`
}

type ReadNoticeRes struct {
	g.Meta `mime:"application/json"`
	Data   res.SystemNotice `json:"data" dc:"通知信息"`
}

type UpdateNoticeReq struct {
	g.Meta `path:"/notice/update/{Id}" method:"put" tags:"通知" summary:"更新通知." x-permission:"system:notice:update"`
	model.AuthorHeader
	req.SystemNoticeUpdate
}

type UpdateNoticeRes struct {
	g.Meta `mime:"application/json"`
}

type DeleteNoticeReq struct {
	g.Meta `path:"/notice/delete" method:"delete" tags:"通知" summary:"删除通知" x-permission:"system:notice:delete"`
	model.AuthorHeader
	Ids []int64 `json:"ids" dc:"ids" v:"min-length:1#通知Id不能为空"`
}

type DeleteNoticeRes struct {
	g.Meta `mime:"application/json"`
}

type RealDeleteNoticeReq struct {
	g.Meta `path:"/notice/realDelete" method:"delete" tags:"通知" summary:"单个或批量真实删除通知 （清空回收站）." x-permission:"system:notice:realDelete"`
	model.AuthorHeader
	Ids []int64 `json:"ids" dc:"ids" v:"min-length:1#通知Id不能为空"`
}

type RealDeleteNoticeRes struct {
	g.Meta `mime:"application/json"`
}

type RecoveryNoticeReq struct {
	g.Meta `path:"/notice/recovery" method:"put" tags:"通知" summary:"单个或批量恢复在回收站的通知." x-permission:"system:notice:recovery"`
	model.AuthorHeader
	Ids []int64 `json:"ids" dc:"ids" v:"min-length:1#通知Id不能为空"`
}

type RecoveryNoticeRes struct {
	g.Meta `mime:"application/json"`
}
