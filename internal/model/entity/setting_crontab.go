// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SettingCrontab is the golang structure for table setting_crontab.
type SettingCrontab struct {
	Id        int64       `json:"id"        orm:"id"         description:"主键"`                                       // 主键
	Name      string      `json:"name"      orm:"name"       description:"任务名称"`                                     // 任务名称
	Type      int         `json:"type"      orm:"type"       description:"任务类型 (1 command, 2 class, 3 url, 4 eval)"` // 任务类型 (1 command, 2 class, 3 url, 4 eval)
	Target    string      `json:"target"    orm:"target"     description:"调用任务字符串"`                                  // 调用任务字符串
	Parameter string      `json:"parameter" orm:"parameter"  description:"调用任务参数"`                                   // 调用任务参数
	Rule      string      `json:"rule"      orm:"rule"       description:"任务执行表达式"`                                  // 任务执行表达式
	Singleton int         `json:"singleton" orm:"singleton"  description:"是否单次执行 (1 是 2 不是)"`                        // 是否单次执行 (1 是 2 不是)
	Status    int         `json:"status"    orm:"status"     description:"状态 (1正常 2停用)"`                             // 状态 (1正常 2停用)
	CreatedBy int64       `json:"createdBy" orm:"created_by" description:"创建者"`                                      // 创建者
	UpdatedBy int64       `json:"updatedBy" orm:"updated_by" description:"更新者"`                                      // 更新者
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:""`                                         //
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:""`                                         //
	Remark    string      `json:"remark"    orm:"remark"     description:"备注"`                                       // 备注
}
