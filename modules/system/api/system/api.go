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

type IndexApiReq struct {
	g.Meta `path:"/api/index" method:"get" tags:"接口" summary:"接口列表." x-permission:"system:api:index" `
	model.AuthorHeader
	model.PageListReq
	req.SystemApiSearch
}

type IndexApiRes struct {
	g.Meta `mime:"application/json"`
	page.PageRes
	Items []res.SystemApi `json:"items"  dc:"api list" `
}

type ListApiReq struct {
	g.Meta `path:"/api/list" method:"get" tags:"接口" summary:"列表，无分页.." x-permission:"system:api:list" `
	model.AuthorHeader
}

type ListApiRes struct {
	g.Meta `mime:"application/json"`
	Data   []res.SystemApi `json:"data"  dc:"list" `
}

type RecycleApiReq struct {
	g.Meta `path:"/api/recycle" method:"get" tags:"接口" summary:"回收站接口列表." x-permission:"system:api:recycle" `
	model.AuthorHeader
	model.PageListReq
	req.SystemApiSearch
}

type RecycleApiRes struct {
	g.Meta `mime:"application/json"`
	page.PageRes
	Items []res.SystemApi `json:"items"  dc:"api list" `
}

type SaveApiReq struct {
	g.Meta `path:"/api/save" method:"post" tags:"接口" summary:"新增接口." x-permission:"system:api:save"`
	model.AuthorHeader
	req.SystemApiSave
}

type SaveApiRes struct {
	g.Meta `mime:"application/json"`
	Id     int64 `json:"id" dc:"接口 id"`
}

type ReadApiReq struct {
	g.Meta `path:"/api/read/{Id}" method:"get" tags:"接口" summary:"更新接口." x-permission:"system:api:read"`
	model.AuthorHeader
	Id int64 `json:"id" dc:"接口 id" v:"required|min:1#接口Id不能为空"`
}

type ReadApiRes struct {
	g.Meta `mime:"application/json"`
	Data   res.SystemApi `json:"data" dc:"接口信息"`
}

type UpdateApiReq struct {
	g.Meta `path:"/api/update/{Id}" method:"put" tags:"接口" summary:"更新接口." x-permission:"system:api:update"`
	model.AuthorHeader
	req.SystemApiUpdate
}

type UpdateApiRes struct {
	g.Meta `mime:"application/json"`
}

type DeleteApiReq struct {
	g.Meta `path:"/api/delete" method:"delete" tags:"接口" summary:"删除接口" x-permission:"system:api:delete"`
	model.AuthorHeader
	Ids []int64 `json:"ids" dc:"ids" v:"min-length:1#接口Id不能为空"`
}

type DeleteApiRes struct {
	g.Meta `mime:"application/json"`
}

type RealDeleteApiReq struct {
	g.Meta `path:"/api/realDelete" method:"delete" tags:"接口" summary:"单个或批量真实删除接口 （清空回收站）." x-permission:"system:api:realDelete"`
	model.AuthorHeader
	Ids []int64 `json:"ids" dc:"ids" v:"min-length:1#接口Id不能为空"`
}

type RealDeleteApiRes struct {
	g.Meta `mime:"application/json"`
}

type RecoveryApiReq struct {
	g.Meta `path:"/api/recovery" method:"put" tags:"接口" summary:"单个或批量恢复在回收站的接口." x-permission:"system:api:recovery"`
	model.AuthorHeader
	Ids []int64 `json:"ids" dc:"ids" v:"min-length:1#接口Id不能为空"`
}

type RecoveryApiRes struct {
	g.Meta `mime:"application/json"`
}

type ChangeStatusApiReq struct {
	g.Meta `path:"/api/changeStatus" method:"put" tags:"接口" summary:"更改状态" x-permission:"system:api:update"`
	model.AuthorHeader
	Id     int64 `json:"id" dc:"ids" v:"min:1#Id不能为空"`
	Status int   `json:"status" dc:"status" v:"min:1#状态不能为空"`
}

type ChangeStatusApiRes struct {
	g.Meta `mime:"application/json"`
}
