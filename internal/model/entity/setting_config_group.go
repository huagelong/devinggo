// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SettingConfigGroup is the golang structure for table setting_config_group.
type SettingConfigGroup struct {
	Id        int64       `json:"id"        orm:"id"         description:"主键"`    // 主键
	Name      string      `json:"name"      orm:"name"       description:"配置组名称"` // 配置组名称
	Code      string      `json:"code"      orm:"code"       description:"配置组标识"` // 配置组标识
	CreatedBy int64       `json:"createdBy" orm:"created_by" description:"创建者"`   // 创建者
	UpdatedBy int64       `json:"updatedBy" orm:"updated_by" description:"更新者"`   // 更新者
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:""`      //
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:""`      //
	Remark    string      `json:"remark"    orm:"remark"     description:"备注"`    // 备注
}
