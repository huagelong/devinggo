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

type IndexReq struct {
	g.Meta `path:"/dept/index" method:"get" tags:"部门" summary:"部门树列表." x-permission:"system:dept:index" `
	model.AuthorHeader
	req.SystemDeptSearch
}

type IndexRes struct {
	g.Meta `mime:"application/json"`
	Data   []res.SystemListDeptTree `json:"data"  dc:"dept tree list" `
}

type RecycleReq struct {
	g.Meta `path:"/dept/recycle" method:"get" tags:"部门" summary:"回收站部门树列表." x-permission:"system:dept:recycle" `
	model.AuthorHeader
	req.SystemDeptSearch
}

type RecycleRes struct {
	g.Meta `mime:"application/json"`
	Data   []res.SystemListDeptTree `json:"data"  dc:"dept tree list" `
}

type TreeReq struct {
	g.Meta `path:"/dept/tree" method:"get" tags:"部门" summary:"前端选择树（不需要权限）." x-exceptAuth:"true" x-permission:"system:dept:tree" `
	model.AuthorHeader
}

type TreeRes struct {
	g.Meta `mime:"application/json"`
	Data   []res.SystemDeptTree `json:"data"  dc:"dept tree list" `
}

type GetLeaderListReq struct {
	g.Meta `path:"/dept/getLeaderList" method:"get" tags:"部门" summary:"获取上传资源列表." x-permission:"system:dept:getLeaderList"`
	model.AuthorHeader
	model.PageListReq
	req.SystemDeptLeaderSearch
}

type GetLeaderListRes struct {
	g.Meta `mime:"application/json"`
	page.PageRes
	Items []res.SystemDeptLeaderInfo `json:"items"  dc:"leader list" `
}

type SaveReq struct {
	g.Meta `path:"/dept/save" method:"post" tags:"部门" summary:"新增部门." x-permission:"system:dept:save"`
	model.AuthorHeader
	req.SystemDeptSave
}

type SaveRes struct {
	g.Meta `mime:"application/json"`
	Id     int64 `json:"id" dc:"dept id"`
}

type AddLeaderReq struct {
	g.Meta `path:"/dept/addLeader" method:"post" tags:"部门" summary:"新增部门领导" x-permission:"system:dept:update"`
	model.AuthorHeader
	req.SystemDeptAddLeader
}

type AddLeaderRes struct {
	g.Meta `mime:"application/json"`
}

type DelLeaderReq struct {
	g.Meta `path:"/dept/delLeader" method:"delete" tags:"部门" summary:"删除部门领导" x-permission:"system:dept:delete"`
	model.AuthorHeader
	Id  int64   `json:"id" dc:"dept id" v:"min:1#部门Id不能为空"`
	Ids []int64 `json:"ids" dc:"ids" v:"min-length:1#用户Id不能为空"`
}

type DelLeaderRes struct {
	g.Meta `mime:"application/json"`
}

type UpdateReq struct {
	g.Meta `path:"/dept/update/{Id}" method:"put" tags:"部门" summary:"更新部门." x-permission:"system:dept:update"`
	model.AuthorHeader
	req.SystemDeptSave
}

type UpdateRes struct {
	g.Meta `mime:"application/json"`
}

type DeleteReq struct {
	g.Meta `path:"/dept/delete" method:"delete" tags:"部门" summary:"删除部门" x-permission:"system:dept:delete"`
	model.AuthorHeader
	Ids []int64 `json:"ids" dc:"ids" v:"min-length:1#部门Id不能为空"`
}

type DeleteRes struct {
	g.Meta `mime:"application/json"`
}

type RealDeleteReq struct {
	g.Meta `path:"/dept/realDelete" method:"delete" tags:"部门" summary:"单个或批量真实删除部门 （清空回收站）." x-permission:"system:dept:realDelete"`
	model.AuthorHeader
	Ids []int64 `json:"ids" dc:"ids" v:"min-length:1#部门Id不能为空"`
}

type RealDeleteRes struct {
	g.Meta `mime:"application/json"`
}

type RecoveryReq struct {
	g.Meta `path:"/dept/recovery" method:"put" tags:"部门" summary:"单个或批量恢复在回收站的部门." x-permission:"system:dept:recovery"`
	model.AuthorHeader
	Ids []int64 `json:"ids" dc:"ids" v:"min-length:1#部门Id不能为空"`
}

type RecoveryRes struct {
	g.Meta `mime:"application/json"`
}

type ChangeStatusReq struct {
	g.Meta `path:"/dept/changeStatus" method:"put" tags:"部门" summary:"更改部门状态" x-permission:"system:dept:changeStatus"`
	model.AuthorHeader
	Id     int64 `json:"id" dc:"ids" v:"min:1#部门Id不能为空"`
	Status int   `json:"status" dc:"dept status" v:"min:1#部门状态不能为空"`
}

type ChangeStatusRes struct {
	g.Meta `mime:"application/json"`
}

type NumberOperationReq struct {
	g.Meta `path:"/dept/numberOperation" method:"put" tags:"部门" summary:"数字运算操作." x-permission:"system:dept:update"`
	model.AuthorHeader
	Id          int64  `json:"id" dc:"ids" v:"min:1#部门Id不能为空"`
	NumberName  string `json:"numberName" dc:"numberName" v:"required#名称不能为空"`
	NumberValue int    `json:"numberValue" dc:"number Value" d:"1" v:"min:1#数字不能为空"`
}

type NumberOperationRes struct {
	g.Meta `mime:"application/json"`
}

type RemoteDeptReq struct {
	g.Meta `path:"/dept/remote" method:"post" tags:"部门" summary:"远程万能通用列表接口." x-exceptAuth:"true" x-permission:"system:dept:remote"`
	model.AuthorHeader
	model.PageListReq
}

type RemoteDeptRes struct {
	g.Meta `mime:"application/json"`
	page.PageRes
	Items []res.SystemDept `json:"items"  dc:"list" `
	Data  []res.SystemDept `json:"data"  dc:"list" `
}
