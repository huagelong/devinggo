// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SettingConfigGroup is the golang structure for table setting_config_group.
type SettingConfigGroup struct {
	Id        int64       `json:"id"        orm:"id"         description:""` //
	Name      string      `json:"name"      orm:"name"       description:""` //
	Code      string      `json:"code"      orm:"code"       description:""` //
	CreatedBy int64       `json:"createdBy" orm:"created_by" description:""` //
	UpdatedBy int64       `json:"updatedBy" orm:"updated_by" description:""` //
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:""` //
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:""` //
	Remark    string      `json:"remark"    orm:"remark"     description:""` //
}
