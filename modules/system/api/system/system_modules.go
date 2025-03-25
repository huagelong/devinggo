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

type IndexSystemModulesReq struct {
	g.Meta `path:"/systemModules/index" method:"get" tags:"modules" summary:"分页列表" x-permission:"system:systemModules:index" `
	model.AuthorHeader
	model.PageListReq
	req.SystemModulesSearch
}

type IndexSystemModulesRes struct {
	g.Meta `mime:"application/json"`
	page.PageRes
	Items []res.SystemModules `json:"items"  dc:"list" `
}

type ListSystemModulesReq struct {
	g.Meta `path:"/systemModules/list" method:"get" tags:"modules" summary:"列表" x-permission:"system:systemModules:list" `
	model.AuthorHeader
	model.ListReq
	req.SystemModulesSearch
}

type ListSystemModulesRes struct {
	g.Meta `mime:"application/json"`
	Data   []res.SystemModules `json:"data"  dc:"list" `
}

type SaveSystemModulesReq struct {
	g.Meta `path:"/systemModules/save" method:"post" tags:"modules" summary:"新增" x-permission:"system:systemModules:save"`
	model.AuthorHeader
	req.SystemModulesSave
}

type SaveSystemModulesRes struct {
	g.Meta `mime:"application/json"`
	Id     int64 `json:"id" dc:"id"`
}

type ReadSystemModulesReq struct {
	g.Meta `path:"/systemModules/read/{Id}" method:"get" tags:"modules" summary:"获取单个信息" x-permission:"system:systemModules:read"`
	model.AuthorHeader
	Id int64 `json:"id" dc:"modules id" v:"required|min:1#Id不能为空"`
}

type ReadSystemModulesRes struct {
	g.Meta `mime:"application/json"`
	Data   res.SystemModules `json:"data" dc:"信息数据"`
}

type UpdateSystemModulesReq struct {
	g.Meta `path:"/systemModules/update/{Id}" method:"put" tags:"modules" summary:"更新" x-permission:"system:systemModules:update"`
	model.AuthorHeader
	req.SystemModulesUpdate
}

type UpdateSystemModulesRes struct {
	g.Meta `mime:"application/json"`
}

type DeleteSystemModulesReq struct {
	g.Meta `path:"/systemModules/delete" method:"delete" tags:"modules" summary:"删除" x-permission:"system:systemModules:delete"`
	model.AuthorHeader
	Ids []int64 `json:"ids" dc:"ids" v:"min-length:1#Id不能为空"`
}

type DeleteSystemModulesRes struct {
	g.Meta `mime:"application/json"`
}

type RecycleSystemModulesReq struct {
	g.Meta `path:"/systemModules/recycle" method:"get" tags:"modules" summary:"回收站列表" x-permission:"system:systemModules:recycle" `
	model.AuthorHeader
	model.PageListReq
	req.SystemModulesSearch
}

type RecycleSystemModulesRes struct {
	g.Meta `mime:"application/json"`
	page.PageRes
	Items []res.SystemModules `json:"items"  dc:"list" `
}

type RealDeleteSystemModulesReq struct {
	g.Meta `path:"/systemModules/realDelete" method:"delete" tags:"modules" summary:"单个或批量真实删除 （清空回收站）" x-permission:"system:systemModules:realDelete"`
	model.AuthorHeader
	Ids []int64 `json:"ids" dc:"ids" v:"min-length:1#Id不能为空"`
}

type RealDeleteSystemModulesRes struct {
	g.Meta `mime:"application/json"`
}

type RecoverySystemModulesReq struct {
	g.Meta `path:"/systemModules/recovery" method:"put" tags:"modules" summary:"单个或批量恢复在回收站的" x-permission:"system:systemModules:recovery"`
	model.AuthorHeader
	Ids []int64 `json:"ids" dc:"ids" v:"min-length:1#Id不能为空"`
}

type RecoverySystemModulesRes struct {
	g.Meta `mime:"application/json"`
}

type ChangeStatusSystemModulesReq struct {
	g.Meta `path:"/systemModules/changeStatus" method:"put" tags:"modules" summary:"更改状态" x-permission:"system:systemModules:changeStatus"`
	model.AuthorHeader
	Id     int64 `json:"id" dc:"ids" v:"min:1#Id不能为空"`
	Status int   `json:"status" dc:"status" v:"min:1#状态不能为空"`
}

type ChangeStatusSystemModulesRes struct {
	g.Meta `mime:"application/json"`
}

type RemoteSystemModulesReq struct {
	g.Meta `path:"/systemModules/remote" method:"post" tags:"modules" summary:"远程万能通用列表接口" x-exceptAuth:"true" x-permission:"system:systemModules:remote"`
	model.AuthorHeader
	model.PageListReq
}

type RemoteSystemModulesRes struct {
	g.Meta `mime:"application/json"`
	page.PageRes
	Items []res.SystemModules `json:"items"  dc:"list" `
	Data  []res.SystemModules `json:"data"  dc:"list" `
}
