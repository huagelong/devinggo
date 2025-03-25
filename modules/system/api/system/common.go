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

type GetUserListReq struct {
	g.Meta `path:"/common/getUserList" method:"get" tags:"公共方法" summary:"获取用户列表." x-exceptAuth:"true" x-permission:"system:common:getUserList"`
	model.AuthorHeader
	model.PageListReq
}

type GetUserListRes struct {
	g.Meta `mime:"application/json"`
	page.PageRes
	Items []res.SystemUser `json:"items"  dc:"user list" `
}

type GetNoticeListReq struct {
	g.Meta `path:"/common/getNoticeList" method:"get" tags:"公共方法" summary:"获取公告列表." x-exceptAuth:"true" x-permission:"system:common:getNoticeList"`
	model.AuthorHeader
	model.PageListReq
}

type GetNoticeListRes struct {
	g.Meta `mime:"application/json"`
	page.PageRes
	Items []res.SystemNotice `json:"items"  dc:"notie list" `
}

type GetLoginLogListReq struct {
	g.Meta `path:"/common/getLoginLogList" method:"get" tags:"公共方法" summary:"获取登录日志列表." x-exceptAuth:"true" x-permission:"system:common:getLoginLogList"`
	model.AuthorHeader
	model.PageListReq
}

type GetLoginLogListRes struct {
	g.Meta `mime:"application/json"`
	page.PageRes
	Items []res.SystemLoginLog `json:"items"  dc:"loginLog list" `
}

type GetOperationLogListReq struct {
	g.Meta `path:"/common/getOperationLogList" method:"get" tags:"公共方法" summary:"获取操作日志列表." x-exceptAuth:"true" x-permission:"system:common:getOperationLogList"`
	model.AuthorHeader
	model.PageListReq
}

type GetOperationLogListRes struct {
	g.Meta `mime:"application/json"`
	page.PageRes
	Items []res.SystemOperLog `json:"items"  dc:"operationLog list" `
}

type ClearAllCacheReq struct {
	g.Meta `path:"/common/clearAllCache" method:"get" tags:"公共方法" summary:"清除缓存." x-exceptAuth:"true" x-permission:"system:common:clearAllCache"`
	model.AuthorHeader
}

type ClearAllCacheRes struct {
	g.Meta `mime:"application/json"`
}

type GetUserInfoByIdsReq struct {
	g.Meta `path:"/common/getUserInfoByIds" method:"post" tags:"公共方法" summary:"通过 id 列表获取用户基础信息." x-exceptAuth:"true" x-permission:"system:common:getUserInfoByIds"`
	model.AuthorHeader
	Ids []int64 `json:"ids" dc:"id list" v:"required" `
}

type GetUserInfoByIdsRes struct {
	g.Meta `mime:"application/json"`
	Data   []res.SystemUser `json:"data"  dc:"user list" `
}

type GetDeptTreeListReq struct {
	g.Meta `path:"/common/getDeptTreeList" method:"get" tags:"公共方法" summary:"获取部门树列表." x-exceptAuth:"true" x-permission:"system:common:getDeptTreeList"`
	model.AuthorHeader
}

type GetDeptTreeListRes struct {
	g.Meta `mime:"application/json"`
	Data   []res.SystemDeptTree `json:"data"  dc:"dept tree list" `
}

type GetRoleListReq struct {
	g.Meta `path:"/common/getRoleList" method:"get" tags:"公共方法" summary:"获取角色列表." x-exceptAuth:"true" x-permission:"system:common:getRoleList"`
	model.AuthorHeader
}

type GetRoleListRes struct {
	g.Meta `mime:"application/json"`
	Data   []res.SystemRole `json:"data"  dc:"role list" `
}

type GetPostListReq struct {
	g.Meta `path:"/common/getPostList" method:"get" tags:"公共方法" summary:"获取岗位列表." x-exceptAuth:"true" x-permission:"system:common:getPostList"`
	model.AuthorHeader
}

type GetPostListRes struct {
	g.Meta `mime:"application/json"`
	Data   []res.SystemPost `json:"data"  dc:"post list" `
}

type GetResourceListReq struct {
	g.Meta `path:"/common/getResourceList" method:"get" tags:"公共方法" summary:"获取上传资源列表." x-exceptAuth:"true" x-permission:"system:common:getResourceList"`
	model.AuthorHeader
	model.PageListReq
	req.SystemUploadFileSearch
}

type GetResourceListRes struct {
	g.Meta `mime:"application/json"`
	page.PageRes
	Items []res.SystemUploadFile `json:"items"  dc:"resource list" `
}

type GetModuleListReq struct {
	g.Meta `path:"/common/getModuleList" method:"get" tags:"公共方法" summary:"获取Module列表." x-exceptAuth:"true" x-permission:"system:common:getModuleList"`
	model.AuthorHeader
}

type GetModuleListRes struct {
	g.Meta `mime:"application/json"`
	Data   []res.Module `json:"data"  dc:"list" `
}
