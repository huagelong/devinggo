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
)

type IndexMenuReq struct {
	g.Meta `path:"/menu/index" method:"get" tags:"菜单" summary:"菜单树列表." x-permission:"system:menu:index" `
	model.AuthorHeader
	req.SystemMenuSearch
}

type IndexMenuRes struct {
	g.Meta `mime:"application/json"`
	Data   []res.SystemMenuTree `json:"data"  dc:"menu tree list" `
}

type RecycleMenuReq struct {
	g.Meta `path:"/menu/recycle" method:"get" tags:"菜单" summary:"回收站部门树列表." x-permission:"system:menu:recycle" `
	model.AuthorHeader
	req.SystemMenuSearch
}

type RecycleMenuRes struct {
	g.Meta `mime:"application/json"`
	Data   []res.SystemMenuTree `json:"data"  dc:"menu tree list" `
}

type TreeMenuReq struct {
	g.Meta `path:"/menu/tree" method:"get" tags:"菜单" summary:"前端选择树（不需要权限）." x-exceptAuth:"true" x-permission:"system:menu:tree" `
	model.AuthorHeader
	OnlyMenu bool `json:"onlyMenu" dc:"是否只返回菜单"`
	Scope    bool `json:"scope" dc:"是否只返回指定范围"`
}

type TreeMenuRes struct {
	g.Meta `mime:"application/json"`
	Data   []res.SystemDeptSelectTree `json:"data"  dc:"menu tree list" `
}

type SaveMenuReq struct {
	g.Meta `path:"/menu/save" method:"post" tags:"菜单" summary:"新增菜单." x-permission:"system:menu:save"`
	model.AuthorHeader
	req.SystemMenuSave
}

type SaveMenuRes struct {
	g.Meta `mime:"application/json"`
	Id     int64 `json:"id" dc:"菜单 id"`
}

type UpdateMenuReq struct {
	g.Meta `path:"/menu/update/{Id}" method:"put" tags:"菜单" summary:"更新菜单." x-permission:"system:menu:update"`
	model.AuthorHeader
	req.SystemMenuSave
}

type UpdateMenuRes struct {
	g.Meta `mime:"application/json"`
}

type DeleteMenuReq struct {
	g.Meta `path:"/menu/delete" method:"delete" tags:"菜单" summary:"删除菜单" x-permission:"system:menu:delete"`
	model.AuthorHeader
	Ids []int64 `json:"ids" dc:"ids" v:"min-length:1#菜单Id不能为空"`
}

type DeleteMenuRes struct {
	g.Meta `mime:"application/json"`
}

type RealDeleteMenuReq struct {
	g.Meta `path:"/menu/realDelete" method:"delete" tags:"菜单" summary:"单个或批量真实删除菜单 （清空回收站）." x-permission:"system:menu:realDelete"`
	model.AuthorHeader
	Ids []int64 `json:"ids" dc:"ids" v:"min-length:1#菜单Id不能为空"`
}

type RealDeleteMenuRes struct {
	g.Meta `mime:"application/json"`
}

type RecoveryMenuReq struct {
	g.Meta `path:"/menu/recovery" method:"put" tags:"菜单" summary:"单个或批量恢复在回收站的菜单." x-permission:"system:menu:recovery"`
	model.AuthorHeader
	Ids []int64 `json:"ids" dc:"ids" v:"min-length:1#菜单Id不能为空"`
}

type RecoveryMenuRes struct {
	g.Meta `mime:"application/json"`
}

type ChangeStatusMenuReq struct {
	g.Meta `path:"/menu/changeStatus" method:"put" tags:"菜单" summary:"更改菜单状态" x-permission:"system:menu:changeStatus"`
	model.AuthorHeader
	Id     int64 `json:"id" dc:"ids" v:"min:1#Id不能为空"`
	Status int   `json:"status" dc:"status" v:"min:1#菜单状态不能为空"`
}

type ChangeStatusMenuRes struct {
	g.Meta `mime:"application/json"`
}

type NumberOperationMenuReq struct {
	g.Meta `path:"/menu/numberOperation" method:"put" tags:"菜单" summary:"数字运算操作." x-permission:"system:menu:update"`
	model.AuthorHeader
	Id          int64  `json:"id" dc:"ids" v:"min:1#Id不能为空"`
	NumberName  string `json:"numberName" dc:"numberName" v:"required#名称不能为空"`
	NumberValue int    `json:"numberValue" dc:"number Value" d:"1" v:"min:1#数字不能为空"`
}

type NumberOperationMenuRes struct {
	g.Meta `mime:"application/json"`
}
