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
	"github.com/gogf/gf/v2/net/ghttp"
)

type GetInfoReq struct {
	g.Meta `path:"/getInfo" method:"get" tags:"管理员信息" summary:"获取登录管理员信息" x-exceptAuth:"true" x-permission:"system:user:getInfo" `
	model.AuthorHeader
}

type GetInfoRes struct {
	g.Meta `mime:"application/json"`
	res.SystemUserInfo
}

type UpdateInfoReq struct {
	g.Meta `path:"/user/updateInfo" method:"post" tags:"管理员信息" summary:"获取登录管理员信息" x-exceptAuth:"true" x-permission:"system:user:updateInfo"`
	model.AuthorHeader
	req.SystemUser
}

type UpdateInfoRes struct {
	g.Meta `mime:"application/json"`
}

type ModifyPasswordReq struct {
	g.Meta `path:"/user/modifyPassword" method:"post" tags:"管理员信息" summary:"修改密码" x-exceptAuth:"true" x-permission:"system:user:modifyPassword"`
	model.AuthorHeader
	NewPassword             string `json:"newPassword" v:"required|length:7,20#请输入新密码|密码长度为7~20位"`
	NewPasswordConfirmation string `json:"newPasswordConfirmation" v:"required|length:7,20#请输入确认密码|确认密码长度为7~20位"`
	OldPassword             string `json:"oldPassword" v:"required|length:7,20#请输入旧密码|密码长度为7~20位"`
}

type ModifyPasswordRes struct {
	g.Meta `mime:"application/json"`
}

type IndexUserReq struct {
	g.Meta `path:"/user/index" method:"get" tags:"管理员信息" summary:"管理员信息列表." x-permission:"system:user:index" `
	model.AuthorHeader
	model.PageListReq
	req.SystemUserSearch
}

type IndexUserRes struct {
	g.Meta `mime:"application/json"`
	page.PageRes
	Items []res.SystemUser `json:"items"  dc:"user list" `
}

type RecycleUserReq struct {
	g.Meta `path:"/user/recycle" method:"get" tags:"管理员信息" summary:"回收站管理员信息列表." x-permission:"system:user:recycle" `
	model.AuthorHeader
	model.PageListReq
	req.SystemUserSearch
}

type RecycleUserRes struct {
	g.Meta `mime:"application/json"`
	page.PageRes
	Items []res.SystemUser `json:"items"  dc:"user list" `
}

type SaveUserReq struct {
	g.Meta `path:"/user/save" method:"post" tags:"管理员信息" summary:"新增管理员信息." x-permission:"system:user:save"`
	model.AuthorHeader
	req.SystemUserSave
}

type SaveUserRes struct {
	g.Meta `mime:"application/json"`
	Id     int64 `json:"id" dc:"管理员信息 id"`
}

type ReadUserReq struct {
	g.Meta `path:"/user/read/{Id}" method:"get" tags:"管理员信息" summary:"更新管理员信息." x-permission:"system:user:read"`
	model.AuthorHeader
	Id int64 `json:"id" dc:"管理员信息 id" v:"required|min:1#管理员信息Id不能为空"`
}

type ReadUserRes struct {
	g.Meta `mime:"application/json"`
	Data   res.SystemUserFullInfo `json:"data" dc:"管理员信息信息"`
}

type ClearCacheReq struct {
	g.Meta `path:"/user/clearCache" method:"post" tags:"管理员信息" summary:"更新管理员信息." x-permission:"system:user:cache"`
	model.AuthorHeader
	Id int64 `json:"id" dc:"管理员信息 id" v:"required|min:1#管理员信息Id不能为空"`
}

type ClearCacheRes struct {
	g.Meta `mime:"application/json"`
}

type ExportReq struct {
	g.Meta `path:"/user/export" method:"post" tags:"管理员信息" summary:"用户导出." x-permission:"system:user:export"`
	model.AuthorHeader
	model.ListReq
	req.SystemUserSearch
}

type ExportRes struct {
	g.Meta `mime:"application/json"`
}

type ImportReq struct {
	g.Meta `path:"/user/import" method:"post" mime:"multipart/form-data" tags:"管理员信息" summary:"用户导入." x-permission:"system:user:import"`
	model.AuthorHeader
	File *ghttp.UploadFile `json:"file" type:"file"  dc:"pls upload file"`
}

type ImportRes struct {
	g.Meta `mime:"application/json"`
}

type DownloadTemplateReq struct {
	g.Meta `path:"/user/downloadTemplate" method:"post,get" tags:"管理员信息" summary:"下载导入模板." x-exceptAuth:"true" x-permission:"system:user:downloadTemplate"`
	model.AuthorHeader
}

type DownloadTemplateRes struct {
	g.Meta `mime:"application/json"`
}

type SetHomePageReq struct {
	g.Meta `path:"/user/setHomePage" method:"post" tags:"管理员信息" summary:"设置用户首页." x-permission:"system:user:homePage"`
	model.AuthorHeader
	Id        int64  `json:"id" dc:"管理员信息 id" v:"required|min:1#管理员信息Id不能为空"`
	Dashboard string `json:"dashboard" dc:"dashboard" v:"required"`
}

type SetHomePageRes struct {
	g.Meta `mime:"application/json"`
}

type InitUserPasswordReq struct {
	g.Meta `path:"/user/initUserPassword" method:"put" tags:"管理员信息" summary:"初始化用户密码." x-permission:"system:user:initUserPassword"`
	model.AuthorHeader
	Id int64 `json:"id" dc:"管理员信息 id" v:"required|min:1#管理员信息Id不能为空"`
}

type InitUserPasswordRes struct {
	g.Meta `mime:"application/json"`
}

type UpdateUserReq struct {
	g.Meta `path:"/user/update/{Id}" method:"put" tags:"管理员信息" summary:"更新管理员信息." x-permission:"system:user:update"`
	model.AuthorHeader
	req.SystemUserUpdate
}

type UpdateUserRes struct {
	g.Meta `mime:"application/json"`
}

type DeleteUserReq struct {
	g.Meta `path:"/user/delete" method:"delete" tags:"管理员信息" summary:"删除管理员信息" x-permission:"system:user:delete"`
	model.AuthorHeader
	Ids []int64 `json:"ids" dc:"ids" v:"min-length:1#管理员信息Id不能为空"`
}

type DeleteUserRes struct {
	g.Meta `mime:"application/json"`
}

type RealDeleteUserReq struct {
	g.Meta `path:"/user/realDelete" method:"delete" tags:"管理员信息" summary:"单个或批量真实删除管理员信息 （清空回收站）." x-permission:"system:user:realDelete"`
	model.AuthorHeader
	Ids []int64 `json:"ids" dc:"ids" v:"min-length:1#管理员信息Id不能为空"`
}

type RealDeleteUserRes struct {
	g.Meta `mime:"application/json"`
}

type RecoveryUserReq struct {
	g.Meta `path:"/user/recovery" method:"put" tags:"管理员信息" summary:"单个或批量恢复在回收站的管理员信息." x-permission:"system:user:recovery"`
	model.AuthorHeader
	Ids []int64 `json:"ids" dc:"ids" v:"min-length:1#管理员信息Id不能为空"`
}

type RecoveryUserRes struct {
	g.Meta `mime:"application/json"`
}

type ChangeStatusUserReq struct {
	g.Meta `path:"/user/changeStatus" method:"put" tags:"管理员信息" summary:"更改管理员信息状态" x-permission:"system:user:changeStatus"`
	model.AuthorHeader
	Id     int64 `json:"id" dc:"ids" v:"min:1#Id不能为空"`
	Status int   `json:"status" dc:"status" v:"min:1#管理员信息状态不能为空"`
}

type ChangeStatusUserRes struct {
	g.Meta `mime:"application/json"`
}

type IndexOnlineUserReq struct {
	g.Meta `path:"/onlineUser/index" method:"get" tags:"管理员信息" summary:"获取在线用户列表." x-permission:"system:onlineUser:index" `
	model.AuthorHeader
	model.PageListReq
	req.SystemUserSearch
}

type IndexOnlineUserRes struct {
	g.Meta `mime:"application/json"`
	page.PageRes
	Items []res.SystemUser `json:"items"  dc:"user list" `
}

type KickUserReq struct {
	g.Meta `path:"/onlineUser/kick" method:"post" tags:"管理员信息" summary:"强退用户" x-permission:"system:onlineUser:kick"`
	model.AuthorHeader
	Id    int64  `json:"id" dc:"id" v:"min-length:1#管理员信息Id不能为空"`
	AppId string `json:"app_id" dc:"app_id"`
}

type KickUserRes struct {
	g.Meta `mime:"application/json"`
}

type RemoteUserReq struct {
	g.Meta `path:"/user/remote" method:"post" tags:"管理员信息" summary:"远程万能通用列表接口." x-exceptAuth:"true" x-permission:"system:user:remote"`
	model.AuthorHeader
	model.PageListReq
}

type RemoteUserRes struct {
	g.Meta `mime:"application/json"`
	page.PageRes
	Items []res.SystemUser `json:"items"  dc:"list" `
	Data  []res.SystemUser `json:"data"  dc:"list" `
}
