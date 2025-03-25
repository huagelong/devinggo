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
	"devinggo/modules/system/pkg/worker/glob"
	"github.com/gogf/gf/v2/frame/g"
)

type IndexCrontabReq struct {
	g.Meta `path:"/setting/crontab/index" method:"get" tags:"定时任务" summary:"获取列表分页数据." x-permission:"system:crontab:index" `
	model.AuthorHeader
	model.PageListReq
	req.SettingCrontabSearch
}

type IndexCrontabRes struct {
	g.Meta `mime:"application/json"`
	page.PageRes
	Items []res.SettingCrontab `json:"items"  dc:"list" `
}

type LogPageListReq struct {
	g.Meta `path:"/setting/crontab/logPageList" method:"get" tags:"定时任务" summary:"获取日志列表分页数据." x-permission:"system:crontab:logPageList" `
	model.AuthorHeader
	model.PageListReq
	req.SettingCrontabLogSearch
}

type LogPageListRes struct {
	g.Meta `mime:"application/json"`
	page.PageRes
	Items []res.SettingCrontabLog `json:"items"  dc:"list" `
}

type SaveCrontabReq struct {
	g.Meta `path:"setting/crontab/save" method:"post" tags:"定时任务" summary:"保存数据." x-permission:"system:crontab:save"`
	model.AuthorHeader
	req.SettingCrontabSave
}

type SaveCrontabRes struct {
	g.Meta `mime:"application/json"`
	Id     int64 `json:"id" dc:"id"`
}

type RunCrontabReq struct {
	g.Meta `path:"setting/crontab/run" method:"post" tags:"定时任务" summary:"立即执行定时任务." x-permission:"system:crontab:run"`
	model.AuthorHeader
	Id int64 `json:"id" dc:"id"`
}

type RunCrontabRes struct {
	g.Meta `mime:"application/json"`
}

type ReadCrontabReq struct {
	g.Meta `path:"setting/crontab/read/{Id}" method:"get" tags:"定时任务" summary:"立即执行定时任务." x-permission:"system:crontab:read"`
	model.AuthorHeader
	Id int64 `json:"id" dc:"id" v:"required|min:1" `
}

type ReadCrontabRes struct {
	g.Meta `mime:"application/json"`
	Data   res.SettingCrontab `json:"data"  dc:"list" `
}

type UpdateCrontabReq struct {
	g.Meta `path:"setting/crontab/update/{Id}" method:"put" tags:"定时任务" summary:"更新." x-permission:"system:crontab:update"`
	model.AuthorHeader
	req.SettingCrontabUpdate
}

type UpdateCrontabRes struct {
	g.Meta `mime:"application/json"`
}

type DeleteCrontabReq struct {
	g.Meta `path:"setting/crontab/delete" method:"delete" tags:"定时任务" summary:"删除" x-permission:"system:crontab:delete"`
	model.AuthorHeader
	Ids []int64 `json:"ids" dc:"ids" v:"min-length:1#Id不能为空"`
}

type DeleteCrontabRes struct {
	g.Meta `mime:"application/json"`
}

type DeleteCrontabLogReq struct {
	g.Meta `path:"setting/crontab/deleteCrontabLog" method:"delete" tags:"定时任务" summary:"删除定时任务日志" x-permission:"system:crontab:deleteCrontabLog"`
	model.AuthorHeader
	Ids []int64 `json:"ids" dc:"ids" v:"min-length:1#Id不能为空"`
}

type DeleteCrontabLogRes struct {
	g.Meta `mime:"application/json"`
}

type ChangeStatusCrontabReq struct {
	g.Meta `path:"setting/crontab/changeStatus" method:"put" tags:"定时任务" summary:"更改状态" x-permission:"system:crontab:update"`
	model.AuthorHeader
	Id     int64 `json:"id" dc:"ids" v:"min:1#Id不能为空"`
	Status int   `json:"status" dc:"status" v:"min:1#状态不能为空"`
}

type ChangeStatusCrontabRes struct {
	g.Meta `mime:"application/json"`
}

type GetCrontabTargetReq struct {
	g.Meta `path:"setting/crontab/getTarget" method:"get" tags:"定时任务" summary:"获取target." x-permission:"system:crontab:getTarget"`
	model.AuthorHeader
	Type int `json:"type" dc:"type"`
}

type GetCrontabTargetRes struct {
	g.Meta `mime:"application/json"`
	Data   []glob.TaskStruct `json:"data"  dc:"list" `
}
