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

type IndexConfigGroupReq struct {
	g.Meta `path:"/setting/configGroup/index" method:"get" tags:"配置" summary:"获取系统组配置." x-permission:"system:config:index" `
	model.AuthorHeader
}

type IndexConfigGroupRes struct {
	g.Meta `mime:"application/json"`
	Data   []res.SettingConfigGroup `json:"data"  dc:"list" `
}

type SaveConfigGroupReq struct {
	g.Meta `path:"/setting/configGroup/save" method:"post" tags:"配置" summary:"保存配置组." x-permission:"system:config:save"`
	model.AuthorHeader
	req.SettingConfigGroupSave
}

type SaveConfigGroupRes struct {
	g.Meta `mime:"application/json"`
	Id     int64 `json:"id" dc:"id"`
}

type UpdateConfigGroupReq struct {
	g.Meta `path:"/setting/configGroup/update" method:"post" tags:"配置" summary:"更新配置组." x-permission:"system:config:update"`
	model.AuthorHeader
	req.SettingConfigGroupUpdate
}

type UpdateConfigGroupRes struct {
	g.Meta `mime:"application/json"`
}

type DeleteConfigGroupReq struct {
	g.Meta `path:"/setting/configGroup/delete" method:"delete" tags:"配置" summary:"删除" x-permission:"system:config:delete"`
	model.AuthorHeader
	Id int64 `json:"id" dc:"id" v:"min:1#Id不能为空"`
}

type DeleteConfigGroupRes struct {
	g.Meta `mime:"application/json"`
}

type IndexConfigReq struct {
	g.Meta `path:"/setting/config/index" method:"get" tags:"配置" summary:"获取配置列表." x-permission:"system:config:index" `
	model.AuthorHeader
	req.SettingConfigSearch
}

type IndexConfigRes struct {
	g.Meta `mime:"application/json"`
	Data   []res.SettingConfig `json:"data"  dc:"list" `
}

type SaveConfigReq struct {
	g.Meta `path:"/setting/config/save" method:"post" tags:"配置" summary:"保存配置." x-permission:"system:config:save"`
	model.AuthorHeader
	req.SettingConfigSave
}

type SaveConfigRes struct {
	g.Meta `mime:"application/json"`
	Id     int64 `json:"id" dc:"id"`
}

type UpdateConfigReq struct {
	g.Meta `path:"/setting/config/update" method:"post" tags:"配置" summary:"更新配置." x-permission:"system:config:update"`
	model.AuthorHeader
	req.SettingConfigUpdate
}

type UpdateConfigRes struct {
	g.Meta `mime:"application/json"`
}

type UpdateByKeysConfigReq struct {
	g.Meta `path:"/setting/config/updateByKeys" method:"post" tags:"配置" summary:"按 keys 更新配置." x-permission:"system:config:update"`
	model.AuthorHeader
}

type UpdateByKeysConfigRes struct {
	g.Meta `mime:"application/json"`
}

type DeleteConfigReq struct {
	g.Meta `path:"/setting/config/delete" method:"delete" tags:"配置" summary:"删除" x-permission:"system:config:delete"`
	model.AuthorHeader
	Ids []string `json:"ids" dc:"ids" v:"required#请选择要删除的配置"`
}

type DeleteConfigRes struct {
	g.Meta `mime:"application/json"`
}
