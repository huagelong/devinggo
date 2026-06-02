// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/os/gtime"
)

// SettingCrontab is the golang structure for table setting_crontab.
type SettingCrontab struct {
	Id        int64       `json:"id"        orm:"id"         description:""` //
	Name      string      `json:"name"      orm:"name"       description:""` //
	Type      int         `json:"type"      orm:"type"       description:""` //
	Target    string      `json:"target"    orm:"target"     description:""` //
	Parameter *gjson.Json `json:"parameter" orm:"parameter"  description:""` //
	Rule      string      `json:"rule"      orm:"rule"       description:""` //
	Singleton int         `json:"singleton" orm:"singleton"  description:""` //
	Status    int         `json:"status"    orm:"status"     description:""` //
	CreatedBy int64       `json:"createdBy" orm:"created_by" description:""` //
	UpdatedBy int64       `json:"updatedBy" orm:"updated_by" description:""` //
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:""` //
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:""` //
	Remark    string      `json:"remark"    orm:"remark"     description:""` //
}
