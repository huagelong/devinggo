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

type GetLoginLogPageListReq struct {
	g.Meta `path:"/logs/getLoginLogPageList" method:"get" tags:"日志" summary:"获取登录日志列表." x-permission:"system:loginLog" `
	model.AuthorHeader
	model.PageListReq
	req.SystemLoginLogSearch
}

type GetLoginLogPageListRes struct {
	g.Meta `mime:"application/json"`
	page.PageRes
	Items []res.SystemLoginLog `json:"items"  dc:"list" `
}

type GetOperLogPageListReq struct {
	g.Meta `path:"/logs/getOperLogPageList" method:"get" tags:"日志" summary:"获取操作日志列表." x-permission:"system:operLog" `
	model.AuthorHeader
	model.PageListReq
	req.SystemOperLogSearch
}

type GetOperLogPageListRes struct {
	g.Meta `mime:"application/json"`
	page.PageRes
	Items []res.SystemOperLog `json:"items"  dc:"list" `
}

type GetApiLogPageListReq struct {
	g.Meta `path:"/logs/getApiLogPageList" method:"get" tags:"日志" summary:"获取接口日志列表." x-permission:"system:apiLog" `
	model.AuthorHeader
	model.PageListReq
	req.SystemApiLogSearch
}

type GetApiLogPageListRes struct {
	g.Meta `mime:"application/json"`
	page.PageRes
	Items []res.SystemApiLog `json:"items"  dc:"list" `
}

type DeleteOperLogReq struct {
	g.Meta `path:"/logs/deleteOperLog" method:"delete" tags:"日志" summary:"删除操作日志." x-permission:"system:operLog"`
	model.AuthorHeader
	Ids []int64 `json:"ids" dc:"ids" v:"min-length:1#操作日志Id不能为空"`
}

type DeleteOperLogRes struct {
	g.Meta `mime:"application/json"`
}

type DeleteApiLogReq struct {
	g.Meta `path:"/logs/deleteApiLog" method:"delete" tags:"日志" summary:"删除api日志." x-permission:"system:apiLog"`
	model.AuthorHeader
	Ids []int64 `json:"ids" dc:"ids" v:"min-length:1#日志Id不能为空"`
}

type DeleteApiLogRes struct {
	g.Meta `mime:"application/json"`
}

type DeleteLoginLogReq struct {
	g.Meta `path:"/logs/deleteLoginLog" method:"delete" tags:"日志" summary:"删除登录日志." x-permission:"system:loginLog"`
	model.AuthorHeader
	Ids []int64 `json:"ids" dc:"ids" v:"min-length:1#日志Id不能为空"`
}

type DeleteLoginLogRes struct {
	g.Meta `mime:"application/json"`
}
