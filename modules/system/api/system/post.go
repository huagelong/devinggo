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

type IndexPostReq struct {
	g.Meta `path:"/post/index" method:"get" tags:"岗位" summary:"岗位列表." x-permission:"system:post:index" `
	model.AuthorHeader
	model.PageListReq
	req.SystemPostSearch
}

type IndexPostRes struct {
	g.Meta `mime:"application/json"`
	page.PageRes
	Items []res.SystemPost `json:"items"  dc:"post list" `
}

type RecyclePostReq struct {
	g.Meta `path:"/post/recycle" method:"get" tags:"岗位" summary:"回收站岗位列表." x-permission:"system:post:recycle" `
	model.AuthorHeader
	model.PageListReq
	req.SystemPostSearch
}

type RecyclePostRes struct {
	g.Meta `mime:"application/json"`
	page.PageRes
	Items []res.SystemPost `json:"items"  dc:"post list" `
}

type ListPostReq struct {
	g.Meta `path:"/post/list" method:"get" tags:"岗位" summary:"前端选择树（不需要权限）." x-permission:"system:post:list" `
	model.AuthorHeader
}

type ListPostRes struct {
	g.Meta `mime:"application/json"`
	Data   []res.SystemPost `json:"data"  dc:"post tree list" `
}

type SavePostReq struct {
	g.Meta `path:"/post/save" method:"post" tags:"岗位" summary:"新增岗位." x-permission:"system:post:save"`
	model.AuthorHeader
	req.SystemPostSave
}

type SavePostRes struct {
	g.Meta `mime:"application/json"`
	Id     int64 `json:"id" dc:"岗位 id"`
}

type ReadPostReq struct {
	g.Meta `path:"/post/read/{Id}" method:"get" tags:"岗位" summary:"获取岗位信息." x-permission:"system:post:read"`
	model.AuthorHeader
	Id int64 `json:"id" dc:"岗位 id" v:"required|min:1#岗位Id不能为空"`
}

type ReadPostRes struct {
	g.Meta `mime:"application/json"`
	Data   res.SystemPost `json:"data" dc:"岗位信息"`
}

type UpdatePostReq struct {
	g.Meta `path:"/post/update/{Id}" method:"put" tags:"岗位" summary:"更新岗位." x-permission:"system:post:update"`
	model.AuthorHeader
	req.SystemPostSave
}

type UpdatePostRes struct {
	g.Meta `mime:"application/json"`
}

type DeletePostReq struct {
	g.Meta `path:"/post/delete" method:"delete" tags:"岗位" summary:"删除岗位" x-permission:"system:post:delete"`
	model.AuthorHeader
	Ids []int64 `json:"ids" dc:"ids" v:"min-length:1#岗位Id不能为空"`
}

type DeletePostRes struct {
	g.Meta `mime:"application/json"`
}

type RealDeletePostReq struct {
	g.Meta `path:"/post/realDelete" method:"delete" tags:"岗位" summary:"单个或批量真实删除岗位 （清空回收站）." x-permission:"system:post:realDelete"`
	model.AuthorHeader
	Ids []int64 `json:"ids" dc:"ids" v:"min-length:1#岗位Id不能为空"`
}

type RealDeletePostRes struct {
	g.Meta `mime:"application/json"`
}

type RecoveryPostReq struct {
	g.Meta `path:"/post/recovery" method:"put" tags:"岗位" summary:"单个或批量恢复在回收站的岗位." x-permission:"system:post:recovery"`
	model.AuthorHeader
	Ids []int64 `json:"ids" dc:"ids" v:"min-length:1#岗位Id不能为空"`
}

type RecoveryPostRes struct {
	g.Meta `mime:"application/json"`
}

type ChangeStatusPostReq struct {
	g.Meta `path:"/post/changeStatus" method:"put" tags:"岗位" summary:"更改岗位状态" x-permission:"system:post:changeStatus"`
	model.AuthorHeader
	Id     int64 `json:"id" dc:"ids" v:"min:1#Id不能为空"`
	Status int   `json:"status" dc:"status" v:"min:1#岗位状态不能为空"`
}

type ChangeStatusPostRes struct {
	g.Meta `mime:"application/json"`
}

type NumberOperationPostReq struct {
	g.Meta `path:"/post/numberOperation" method:"put" tags:"岗位" summary:"数字运算操作." x-permission:"system:post:update"`
	model.AuthorHeader
	Id          int64  `json:"id" dc:"ids" v:"min:1#Id不能为空"`
	NumberName  string `json:"numberName" dc:"numberName" v:"required#名称不能为空"`
	NumberValue int    `json:"numberValue" dc:"number Value" d:"1" v:"min:1#数字不能为空"`
}

type NumberOperationPostRes struct {
	g.Meta `mime:"application/json"`
}

type RemotePostReq struct {
	g.Meta `path:"/post/remote" method:"post" tags:"岗位" summary:"远程万能通用列表接口." x-exceptAuth:"true" x-permission:"system:post:remote"`
	model.AuthorHeader
	model.PageListReq
}

type RemotePostRes struct {
	g.Meta `mime:"application/json"`
	page.PageRes
	Items []res.SystemPost `json:"items"  dc:"list" `
	Data  []res.SystemPost `json:"data"  dc:"list" `
}
