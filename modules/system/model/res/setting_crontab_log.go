// Package res
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package res

import "github.com/gogf/gf/v2/os/gtime"

type SettingCrontabLog struct {
	Id            int64       `json:"id"                         description:"主键"`         // 主键
	CrontabId     int64       `json:"crontab_id"          description:"任务ID"`              // 任务ID
	Name          string      `json:"name"                     description:"任务名称"`         // 任务名称
	Target        string      `json:"target"                 description:"任务调用目标字符串"`      // 任务调用目标字符串
	Parameter     string      `json:"parameter"           description:"任务调用参数"`            // 任务调用参数
	ExceptionInfo string      `json:"exception_info"  description:"异常信息"`                  // 异常信息
	Status        int         `json:"status"                 description:"执行状态 (1成功 2失败)"` // 执行状态 (1成功 2失败)
	CreatedAt     *gtime.Time `json:"created_at"          description:"创建时间"`              // 创建时间
}
