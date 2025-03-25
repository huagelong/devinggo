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

type GetAppIdReq struct {
	g.Meta `path:"/app/getAppId" method:"get" tags:"应用管理" summary:"获取应用Id." x-exceptAuth:"true" x-permission:"system:app:getAppId" `
	model.AuthorHeader
}

type GetAppIdRes struct {
	g.Meta `mime:"application/json"`
	AppId  string `json:"app_id" dc:"应用Id"`
}

type GetAppSecretReq struct {
	g.Meta `path:"/app/getAppSecret" method:"get" tags:"应用管理" summary:"获取应用秘钥." x-exceptAuth:"true" x-permission:"system:app:getAppSecret" `
	model.AuthorHeader
}

type GetAppSecretRes struct {
	g.Meta    `mime:"application/json"`
	AppSecret string `json:"app_secret" dc:"应用秘钥"`
}

type GetApiListReq struct {
	g.Meta `path:"/app/getApiList" method:"get" tags:"应用管理" summary:"获取绑定接口列表." x-exceptAuth:"true" x-permission:"system:app:getApiList" `
	model.AuthorHeader
	Id int64 `json:"id" dc:"应用Id" v:"required"`
}

type GetApiListRes struct {
	g.Meta `mime:"application/json"`
	Data   []int64 `json:"data" dc:"获取绑定接口id"`
}

type IndexAppReq struct {
	g.Meta `path:"/app/index" method:"get" tags:"应用管理" summary:"应用管理列表." x-permission:"system:app:index" `
	model.AuthorHeader
	model.PageListReq
	req.SystemAppSearch
}

type IndexAppRes struct {
	g.Meta `mime:"application/json"`
	page.PageRes
	Items []res.SystemApp `json:"items"  dc:"app list" `
}

type RecycleAppReq struct {
	g.Meta `path:"/app/recycle" method:"get" tags:"应用管理" summary:"回收站应用管理列表." x-permission:"system:app:recycle" `
	model.AuthorHeader
	model.PageListReq
	req.SystemAppSearch
}

type RecycleAppRes struct {
	g.Meta `mime:"application/json"`
	page.PageRes
	Items []res.SystemApp `json:"items"  dc:"app list" `
}

type SaveAppReq struct {
	g.Meta `path:"/app/save" method:"post" tags:"应用管理" summary:"新增应用管理." x-permission:"system:app:save"`
	model.AuthorHeader
	req.SystemAppSave
}

type SaveAppRes struct {
	g.Meta `mime:"application/json"`
	Id     int64 `json:"id" dc:"应用管理 id"`
}

type ReadAppReq struct {
	g.Meta `path:"/app/read/{Id}" method:"get" tags:"应用管理" summary:"更新应用管理." x-permission:"system:app:read"`
	model.AuthorHeader
	Id int64 `json:"id" dc:"应用管理 id" v:"required|min:1#应用管理Id不能为空"`
}

type ReadAppRes struct {
	g.Meta `mime:"application/json"`
	Data   res.SystemApp `json:"data" dc:"应用管理信息"`
}
type UpdateAppReq struct {
	g.Meta `path:"/app/update/{Id}" method:"put" tags:"应用管理" summary:"更新应用管理." x-permission:"system:app:update"`
	model.AuthorHeader
	req.SystemAppUpdate
}

type UpdateAppRes struct {
	g.Meta `mime:"application/json"`
}

type BindAppReq struct {
	g.Meta `path:"/app/bind/{Id}" method:"put" tags:"应用管理" summary:"更新应用管理." x-permission:"system:app:bind"`
	model.AuthorHeader
	Id     int64   `json:"id" dc:"应用管理 id" v:"required"`
	ApiIds []int64 `json:"apiIds" dc:"apiIds" v:"required"`
}

type BindAppRes struct {
	g.Meta `mime:"application/json"`
}

type DeleteAppReq struct {
	g.Meta `path:"/app/delete" method:"delete" tags:"应用管理" summary:"删除应用管理" x-permission:"system:app:delete"`
	model.AuthorHeader
	Ids []int64 `json:"ids" dc:"ids" v:"min-length:1#应用管理Id不能为空"`
}

type DeleteAppRes struct {
	g.Meta `mime:"application/json"`
}

type RealDeleteAppReq struct {
	g.Meta `path:"/app/realDelete" method:"delete" tags:"应用管理" summary:"单个或批量真实删除应用管理 （清空回收站）." x-permission:"system:app:realDelete"`
	model.AuthorHeader
	Ids []int64 `json:"ids" dc:"ids" v:"min-length:1#应用管理Id不能为空"`
}

type RealDeleteAppRes struct {
	g.Meta `mime:"application/json"`
}

type RecoveryAppReq struct {
	g.Meta `path:"/app/recovery" method:"put" tags:"应用管理" summary:"单个或批量恢复在回收站的应用管理." x-permission:"system:app:recovery"`
	model.AuthorHeader
	Ids []int64 `json:"ids" dc:"ids" v:"min-length:1#应用管理Id不能为空"`
}

type RecoveryAppRes struct {
	g.Meta `mime:"application/json"`
}
type ChangeStatusAppReq struct {
	g.Meta `path:"/app/changeStatus" method:"put" tags:"应用管理" summary:"更改状态" x-permission:"system:app:update"`
	model.AuthorHeader
	Id     int64 `json:"id" dc:"ids" v:"min:1#Id不能为空"`
	Status int   `json:"status" dc:"status" v:"min:1#状态不能为空"`
}

type ChangeStatusAppRes struct {
	g.Meta `mime:"application/json"`
}

type RemoteAppReq struct {
	g.Meta `path:"/app/remote" method:"post" tags:"应用管理" summary:"远程万能通用列表接口." x-exceptLogin:"true" x-permission:"system:app:remote"`
	model.AuthorHeader
	model.PageListReq
}

type RemoteAppRes struct {
	g.Meta `mime:"application/json"`
	page.PageRes
	Items []res.SystemApp `json:"items"  dc:"list" `
	Data  []res.SystemApp `json:"data"  dc:"list" `
}
