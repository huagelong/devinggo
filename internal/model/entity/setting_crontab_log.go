// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SettingCrontabLog is the golang structure for table setting_crontab_log.
type SettingCrontabLog struct {
	Id            int64       `json:"id"            orm:"id"             description:""` //
	CrontabId     int64       `json:"crontabId"     orm:"crontab_id"     description:""` //
	Name          string      `json:"name"          orm:"name"           description:""` //
	Target        string      `json:"target"        orm:"target"         description:""` //
	Parameter     string      `json:"parameter"     orm:"parameter"      description:""` //
	ExceptionInfo string      `json:"exceptionInfo" orm:"exception_info" description:""` //
	Status        int         `json:"status"        orm:"status"         description:""` //
	CreatedAt     *gtime.Time `json:"createdAt"     orm:"created_at"     description:""` //
}
