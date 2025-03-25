// Package res
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package res

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/os/gtime"
)

type SettingCrontab struct {
	Id        int64       `json:"id"                 description:"主键"`                                     // 主键
	Name      string      `json:"name"             description:"任务名称"`                                     // 任务名称
	Type      int         `json:"type"             description:"任务类型 (1 command, 2 class, 3 url, 4 eval)"` // 任务类型 (1 command, 2 class, 3 url, 4 eval)
	Target    string      `json:"target"         description:"调用任务字符串"`                                    // 调用任务字符串
	Parameter string      `json:"parameter"   description:"调用任务参数"`                                        // 调用任务参数
	Rule      string      `json:"rule"             description:"任务执行表达式"`                                  // 任务执行表达式
	Singleton int         `json:"singleton"   description:"是否单次执行 (1 是 2 不是)"`                             // 是否单次执行 (1 是 2 不是)
	Status    int         `json:"status"         description:"状态 (1正常 2停用)"`                               // 状态 (1正常 2停用)
	CreatedBy int64       `json:"created_by"  description:"创建者"`                                           // 创建者
	UpdatedBy int64       `json:"updated_by"  description:"更新者"`                                           // 更新者
	CreatedAt *gtime.Time `json:"created_at"  description:""`                                              //
	UpdatedAt *gtime.Time `json:"updated_at"  description:""`                                              //
	Remark    string      `json:"remark"         description:"备注"`                                         // 备注
}

type SettingCrontabOne struct {
	Id        int64       `json:"id"                 description:"主键"`                                     // 主键
	Name      string      `json:"name"             description:"任务名称"`                                     // 任务名称
	Type      int         `json:"type"             description:"任务类型 (1 command, 2 class, 3 url, 4 eval)"` // 任务类型 (1 command, 2 class, 3 url, 4 eval)
	Target    string      `json:"target"         description:"调用任务字符串"`                                    // 调用任务字符串
	Parameter *gjson.Json `json:"parameter"   description:"调用任务参数"`                                        // 调用任务参数
	Rule      string      `json:"rule"             description:"任务执行表达式"`                                  // 任务执行表达式
	Singleton int         `json:"singleton"   description:"是否单次执行 (1 是 2 不是)"`                             // 是否单次执行 (1 是 2 不是)
	Status    int         `json:"status"         description:"状态 (1正常 2停用)"`                               // 状态 (1正常 2停用)
	CreatedBy int64       `json:"created_by"  description:"创建者"`                                           // 创建者
	UpdatedBy int64       `json:"updated_by"  description:"更新者"`                                           // 更新者
	CreatedAt *gtime.Time `json:"created_at"  description:""`                                              //
	UpdatedAt *gtime.Time `json:"updated_at"  description:""`                                              //
	Remark    string      `json:"remark"         description:"备注"`                                         // 备注
}
