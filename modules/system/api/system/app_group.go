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

type IndexAppGroupReq struct {
	g.Meta `path:"/appGroup/index" method:"get" tags:"应用分组" summary:"应用分组列表." x-permission:"system:appGroup:index" `
	model.AuthorHeader
	model.PageListReq
	req.SystemAppGroupSearch
}

type IndexAppGroupRes struct {
	g.Meta `mime:"application/json"`
	page.PageRes
	Items []res.SystemAppGroup `json:"items"  dc:"appGroup list" `
}

type ListAppGroupReq struct {
	g.Meta `path:"/appGroup/list" method:"get" tags:"应用分组" summary:"列表，无分页.." x-exceptAuth:"true" x-permission:"system:appGroup:list" `
	model.AuthorHeader
}

type ListAppGroupRes struct {
	g.Meta `mime:"application/json"`
	Data   []res.SystemAppGroup `json:"data"  dc:"list" `
}

type RecycleAppGroupReq struct {
	g.Meta `path:"/appGroup/recycle" method:"get" tags:"应用分组" summary:"回收站应用分组列表." x-permission:"system:appGroup:recycle" `
	model.AuthorHeader
	model.PageListReq
	req.SystemAppGroupSearch
}

type RecycleAppGroupRes struct {
	g.Meta `mime:"application/json"`
	page.PageRes
	Items []res.SystemAppGroup `json:"items"  dc:"appGroup list" `
}

type SaveAppGroupReq struct {
	g.Meta `path:"/appGroup/save" method:"post" tags:"应用分组" summary:"新增应用分组." x-permission:"system:appGroup:save"`
	model.AuthorHeader
	req.SystemAppGroupSave
}

type SaveAppGroupRes struct {
	g.Meta `mime:"application/json"`
	Id     int64 `json:"id" dc:"应用分组 id"`
}

type ReadAppGroupReq struct {
	g.Meta `path:"/appGroup/read/{Id}" method:"get" tags:"应用分组" summary:"更新应用分组." x-permission:"system:appGroup:read"`
	model.AuthorHeader
	Id int64 `json:"id" dc:"应用分组 id" v:"required|min:1#应用分组Id不能为空"`
}

type ReadAppGroupRes struct {
	g.Meta `mime:"application/json"`
	Data   res.SystemAppGroup `json:"data" dc:"应用分组信息"`
}

type UpdateAppGroupReq struct {
	g.Meta `path:"/appGroup/update/{Id}" method:"put" tags:"应用分组" summary:"更新应用分组." x-permission:"system:appGroup:update"`
	model.AuthorHeader
	req.SystemAppGroupUpdate
}

type UpdateAppGroupRes struct {
	g.Meta `mime:"application/json"`
}

type DeleteAppGroupReq struct {
	g.Meta `path:"/appGroup/delete" method:"delete" tags:"应用分组" summary:"删除应用分组" x-permission:"system:appGroup:delete"`
	model.AuthorHeader
	Ids []int64 `json:"ids" dc:"ids" v:"min-length:1#应用分组Id不能为空"`
}

type DeleteAppGroupRes struct {
	g.Meta `mime:"application/json"`
}

type RealDeleteAppGroupReq struct {
	g.Meta `path:"/appGroup/realDelete" method:"delete" tags:"应用分组" summary:"单个或批量真实删除应用分组 （清空回收站）." x-permission:"system:appGroup:realDelete"`
	model.AuthorHeader
	Ids []int64 `json:"ids" dc:"ids" v:"min-length:1#应用分组Id不能为空"`
}

type RealDeleteAppGroupRes struct {
	g.Meta `mime:"application/json"`
}

type RecoveryAppGroupReq struct {
	g.Meta `path:"/appGroup/recovery" method:"put" tags:"应用分组" summary:"单个或批量恢复在回收站的应用分组." x-permission:"system:appGroup:recovery"`
	model.AuthorHeader
	Ids []int64 `json:"ids" dc:"ids" v:"min-length:1#应用分组Id不能为空"`
}

type RecoveryAppGroupRes struct {
	g.Meta `mime:"application/json"`
}

type ChangeStatusAppGroupReq struct {
	g.Meta `path:"/appGroup/changeStatus" method:"put" tags:"应用分组" summary:"更改状态" x-permission:"system:appGroup:update"`
	model.AuthorHeader
	Id     int64 `json:"id" dc:"ids" v:"min:1#Id不能为空"`
	Status int   `json:"status" dc:"status" v:"min:1#状态不能为空"`
}

type ChangeStatusAppGroupRes struct {
	g.Meta `mime:"application/json"`
}
