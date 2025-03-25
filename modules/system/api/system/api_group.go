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

type IndexApiGroupReq struct {
	g.Meta `path:"/apiGroup/index" method:"get" tags:"接口分组" summary:"接口分组列表." x-permission:"system:apiGroup:index" `
	model.AuthorHeader
	model.PageListReq
	req.SystemApiGroupSearch
}

type IndexApiGroupRes struct {
	g.Meta `mime:"application/json"`
	page.PageRes
	Items []res.SystemApiGroup `json:"items"  dc:"apiGroup list" `
}

type ListApiGroupReq struct {
	g.Meta `path:"/apiGroup/list" method:"get" tags:"接口分组" summary:"列表，无分页.." x-exceptAuth:"true" x-permission:"system:apiGroup:list" `
	model.AuthorHeader
	req.SystemApiGroupSearch
}

type ListApiGroupRes struct {
	g.Meta `mime:"application/json"`
	Data   []res.SystemApiGroup `json:"data"  dc:"list" `
}

type RecycleApiGroupReq struct {
	g.Meta `path:"/apiGroup/recycle" method:"get" tags:"接口分组" summary:"回收站接口分组列表." x-permission:"system:apiGroup:recycle" `
	model.AuthorHeader
	model.PageListReq
	req.SystemApiGroupSearch
}

type RecycleApiGroupRes struct {
	g.Meta `mime:"application/json"`
	page.PageRes
	Items []res.SystemApiGroup `json:"items"  dc:"apiGroup list" `
}

type SaveApiGroupReq struct {
	g.Meta `path:"/apiGroup/save" method:"post" tags:"接口分组" summary:"新增接口分组." x-permission:"system:apiGroup:save"`
	model.AuthorHeader
	req.SystemApiGroupSave
}

type SaveApiGroupRes struct {
	g.Meta `mime:"application/json"`
	Id     int64 `json:"id" dc:"接口分组 id"`
}

type ReadApiGroupReq struct {
	g.Meta `path:"/apiGroup/read/{Id}" method:"get" tags:"接口分组" summary:"更新接口分组." x-permission:"system:apiGroup:read"`
	model.AuthorHeader
	Id int64 `json:"id" dc:"接口分组 id" v:"required|min:1#接口分组Id不能为空"`
}

type ReadApiGroupRes struct {
	g.Meta `mime:"application/json"`
	Data   res.SystemApiGroup `json:"data" dc:"接口分组信息"`
}

type UpdateApiGroupReq struct {
	g.Meta `path:"/apiGroup/update/{Id}" method:"put" tags:"接口分组" summary:"更新接口分组." x-permission:"system:apiGroup:update"`
	model.AuthorHeader
	req.SystemApiGroupUpdate
}

type UpdateApiGroupRes struct {
	g.Meta `mime:"application/json"`
}

type DeleteApiGroupReq struct {
	g.Meta `path:"/apiGroup/delete" method:"delete" tags:"接口分组" summary:"删除接口分组" x-permission:"system:apiGroup:delete"`
	model.AuthorHeader
	Ids []int64 `json:"ids" dc:"ids" v:"min-length:1#接口分组Id不能为空"`
}

type DeleteApiGroupRes struct {
	g.Meta `mime:"application/json"`
}

type RealDeleteApiGroupReq struct {
	g.Meta `path:"/apiGroup/realDelete" method:"delete" tags:"接口分组" summary:"单个或批量真实删除接口分组 （清空回收站）." x-permission:"system:apiGroup:realDelete"`
	model.AuthorHeader
	Ids []int64 `json:"ids" dc:"ids" v:"min-length:1#接口分组Id不能为空"`
}

type RealDeleteApiGroupRes struct {
	g.Meta `mime:"application/json"`
}

type RecoveryApiGroupReq struct {
	g.Meta `path:"/apiGroup/recovery" method:"put" tags:"接口分组" summary:"单个或批量恢复在回收站的接口分组." x-permission:"system:apiGroup:recovery"`
	model.AuthorHeader
	Ids []int64 `json:"ids" dc:"ids" v:"min-length:1#接口分组Id不能为空"`
}

type RecoveryApiGroupRes struct {
	g.Meta `mime:"application/json"`
}

type ChangeStatusApiGroupReq struct {
	g.Meta `path:"/apiGroup/changeStatus" method:"put" tags:"接口分组" summary:"更改状态" x-permission:"system:apiGroup:update"`
	model.AuthorHeader
	Id     int64 `json:"id" dc:"ids" v:"min:1#Id不能为空"`
	Status int   `json:"status" dc:"status" v:"min:1#状态不能为空"`
}

type ChangeStatusApiGroupRes struct {
	g.Meta `mime:"application/json"`
}
