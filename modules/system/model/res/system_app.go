// Package res
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package res

import "github.com/gogf/gf/v2/os/gtime"

type SystemApp struct {
	Id          int64       `json:"id"                    description:"主键"`       // 主键
	GroupId     int64       `json:"group_id"         description:"应用组ID"`         // 应用组ID
	AppName     string      `json:"app_name"         description:"应用名称"`          // 应用名称
	AppId       string      `json:"app_id"             description:"应用ID"`        // 应用ID
	AppSecret   string      `json:"app_secret"     description:"应用密钥"`            // 应用密钥
	Status      int         `json:"status"            description:"状态 (1正常 2停用)"` // 状态 (1正常 2停用)
	Description string      `json:"description"  description:"应用介绍"`              // 应用介绍
	CreatedBy   int64       `json:"created_by"     description:"创建者"`             // 创建者
	UpdatedBy   int64       `json:"updated_by"     description:"更新者"`             // 更新者
	CreatedAt   *gtime.Time `json:"created_at"     description:"创建时间"`            // 创建时间
	UpdatedAt   *gtime.Time `json:"updated_at"     description:"更新时间"`            // 更新时间
	Remark      string      `json:"remark"            description:"备注"`           // 备注
}
