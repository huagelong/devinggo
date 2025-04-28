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

type IndexRoleReq struct {
	g.Meta `path:"/role/index" method:"get" tags:"角色" summary:"角色列表." x-permission:"system:role:index" `
	model.AuthorHeader
	model.PageListReq
	req.SystemRoleSearch
}

type IndexRoleRes struct {
	g.Meta `mime:"application/json"`
	page.PageRes
	Items []res.SystemRole `json:"items"  dc:"role list" `
}

type RecycleRoleReq struct {
	g.Meta `path:"/role/recycle" method:"get" tags:"角色" summary:"回收站角色列表." x-permission:"system:role:recycle" `
	model.AuthorHeader
	model.PageListReq
	req.SystemRoleSearch
}

type RecycleRoleRes struct {
	g.Meta `mime:"application/json"`
	page.PageRes
	Items []res.SystemRole `json:"items"  dc:"role list" `
}

type ListRoleReq struct {
	g.Meta `path:"/role/list" method:"get" tags:"角色" summary:"前端选择树（不需要权限）." x-permission:"system:role:list" `
	model.AuthorHeader
}

type ListRoleRes struct {
	g.Meta `mime:"application/json"`
	Data   []res.SystemRole `json:"data"  dc:"role tree list" `
}

type SaveRoleReq struct {
	g.Meta `path:"/role/save" method:"post" tags:"角色" summary:"新增角色." x-permission:"system:role:save"`
	model.AuthorHeader
	req.SystemRoleSave
}

type SaveRoleRes struct {
	g.Meta `mime:"application/json"`
	Id     int64 `json:"id" dc:"角色 id"`
}

type UpdateRoleReq struct {
	g.Meta `path:"/role/update/{Id}" method:"put" tags:"角色" summary:"更新角色." x-permission:"system:role:update"`
	model.AuthorHeader
	req.SystemRoleSave
}

type UpdateRoleRes struct {
	g.Meta `mime:"application/json"`
}

type DeleteRoleReq struct {
	g.Meta `path:"/role/delete" method:"delete" tags:"角色" summary:"删除角色" x-permission:"system:role:delete"`
	model.AuthorHeader
	Ids []int64 `json:"ids" dc:"ids" v:"min-length:1#角色Id不能为空"`
}

type DeleteRoleRes struct {
	g.Meta `mime:"application/json"`
}

type RealDeleteRoleReq struct {
	g.Meta `path:"/role/realDelete" method:"delete" tags:"角色" summary:"单个或批量真实删除角色 （清空回收站）." x-permission:"system:role:realDelete"`
	model.AuthorHeader
	Ids []int64 `json:"ids" dc:"ids" v:"min-length:1#角色Id不能为空"`
}

type RealDeleteRoleRes struct {
	g.Meta `mime:"application/json"`
}

type RecoveryRoleReq struct {
	g.Meta `path:"/role/recovery" method:"put" tags:"角色" summary:"单个或批量恢复在回收站的角色." x-permission:"system:role:recovery"`
	model.AuthorHeader
	Ids []int64 `json:"ids" dc:"ids" v:"min-length:1#角色Id不能为空"`
}

type RecoveryRoleRes struct {
	g.Meta `mime:"application/json"`
}

type ChangeStatusRoleReq struct {
	g.Meta `path:"/role/changeStatus" method:"put" tags:"角色" summary:"更改角色状态" x-permission:"system:role:changeStatus"`
	model.AuthorHeader
	Id     int64 `json:"id" dc:"ids" v:"min:1#Id不能为空"`
	Status int   `json:"status" dc:"status" v:"min:1#角色状态不能为空"`
}

type ChangeStatusRoleRes struct {
	g.Meta `mime:"application/json"`
}

type NumberOperationRoleReq struct {
	g.Meta `path:"/role/numberOperation" method:"put" tags:"角色" summary:"数字运算操作."  x-permission:"system:role:update"`
	model.AuthorHeader
	Id          int64  `json:"id" dc:"ids" v:"min:1#Id不能为空"`
	NumberName  string `json:"numberName" dc:"numberName" v:"required#名称不能为空"`
	NumberValue int    `json:"numberValue" dc:"number Value" d:"1" v:"min:1#数字不能为空"`
}

type NumberOperationRoleRes struct {
	g.Meta `mime:"application/json"`
}

type MenuPermissionRoleReq struct {
	g.Meta `path:"/role/menuPermission/{Id}" method:"put" tags:"角色" summary:"更新用户菜单权限." x-permission:"system:role:menuPermission"`
	model.AuthorHeader
	req.SystemRoleSave
}

type MenuPermissionRoleRes struct {
	g.Meta `mime:"application/json"`
}

type DataPermissionRoleReq struct {
	g.Meta `path:"/role/dataPermission/{Id}" method:"put" tags:"角色" summary:"更新用户数据权限." x-permission:"system:role:dataPermission"`
	model.AuthorHeader
	req.SystemRoleSave
}

type DataPermissionRoleRes struct {
	g.Meta `mime:"application/json"`
}

type GetMenuByRoleReq struct {
	g.Meta `path:"/role/getMenuByRole/{Id}" method:"get" tags:"角色" summary:"通过角色获取菜单."  x-exceptAuth:"true"  x-permission:"system:role:getMenuByRole"`
	model.AuthorHeader
	Id int64 `json:"id" dc:"id" v:"required|min:1#Id不能为空"`
}

type GetMenuByRoleRes struct {
	g.Meta `mime:"application/json"`
	Data   []*res.SystemRoleMenus `json:"data" dc:"menu id list" `
}

type GetDeptByRoleReq struct {
	g.Meta `path:"/role/getDeptByRole/{Id}" method:"get" tags:"角色" summary:"通过角色获取部门."  x-exceptAuth:"true"  x-permission:"system:role:getDeptByRole"`
	model.AuthorHeader
	Id int64 `json:"id" dc:"id" v:"required|min:1#Id不能为空"`
}

type GetDeptByRoleRes struct {
	g.Meta `mime:"application/json"`
	Data   []*res.SystemRoleDepts `json:"data" dc:"dept id list" `
}

type RemoteRoleReq struct {
	g.Meta `path:"/role/remote" method:"post" tags:"角色" summary:"远程万能通用列表接口." x-exceptAuth:"true" x-permission:"system:role:remote"`
	model.AuthorHeader
	model.PageListReq
}

type RemoteRoleRes struct {
	g.Meta `mime:"application/json"`
	page.PageRes
	Items []res.SystemRole `json:"items"  dc:"list" `
	Data  []res.SystemRole `json:"data"  dc:"list" `
}
