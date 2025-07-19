// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SettingCrontabLog is the golang structure for table setting_crontab_log.
type SettingCrontabLog struct {
	Id            int64       `json:"id"            orm:"id"             description:"主键"`             // 主键
	CrontabId     int64       `json:"crontabId"     orm:"crontab_id"     description:"任务ID"`           // 任务ID
	Name          string      `json:"name"          orm:"name"           description:"任务名称"`           // 任务名称
	Target        string      `json:"target"        orm:"target"         description:"任务调用目标字符串"`      // 任务调用目标字符串
	Parameter     string      `json:"parameter"     orm:"parameter"      description:"任务调用参数"`         // 任务调用参数
	ExceptionInfo string      `json:"exceptionInfo" orm:"exception_info" description:"异常信息"`           // 异常信息
	Status        int         `json:"status"        orm:"status"         description:"执行状态 (1成功 2失败)"` // 执行状态 (1成功 2失败)
	CreatedAt     *gtime.Time `json:"createdAt"     orm:"created_at"     description:"创建时间"`           // 创建时间
}
