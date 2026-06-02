// Package req
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package req

import "github.com/gogf/gf/v2/encoding/gjson"

type SettingCrontabSave struct {
	Name      string      `json:"name"   v:"required"          description:"任务名称"`
	Type      int         `json:"type"    v:"required"          description:"任务类型 (1 command, 2 class, 3 url, 4 eval)"`
	Rule      string      `json:"rule"    v:"required"          description:"任务执行表达式"`
	Target    string      `json:"target"   v:"required"       description:"调用任务字符串"`
	Status    int         `json:"status"   v:"required"       description:"任务状态 (1 enable, 2 disable)"`
	Singleton int         `json:"singleton"         description:"是否单例"`
	Parameter *gjson.Json `json:"parameter"         description:"调用任务参数"`
	Remark    string      `json:"remark"          description:"任务备注"`
}

type SettingCrontabUpdate struct {
	Id        int64       `json:"id" v:"required" description:"任务ID"`
	Name      string      `json:"name"   v:"required"          description:"任务名称"`
	Type      int         `json:"type"    v:"required"          description:"任务类型 (1 command, 2 class, 3 url, 4 eval)"`
	Rule      string      `json:"rule"    v:"required"          description:"任务执行表达式"`
	Target    string      `json:"target"   v:"required"       description:"调用任务字符串"`
	Status    int         `json:"status"   v:"required"       description:"任务状态 (1 enable, 2 disable)"`
	Singleton int         `json:"singleton"         description:"是否单例"`
	Parameter *gjson.Json `json:"parameter"         description:"调用任务参数"`
	Remark    string      `json:"remark"          description:"任务备注"`
}

type SettingCrontabSearch struct {
	Name      string   `json:"name"        description:"任务名称"`
	Type      int      `json:"type"          description:"任务类型"`
	Status    int      `json:"status"        description:"任务状态"`
	CreatedAt []string `json:"created_at" dc:"created at"`
}
