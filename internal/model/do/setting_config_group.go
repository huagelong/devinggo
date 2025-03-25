// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SettingConfigGroup is the golang structure of table setting_config_group for DAO operations like Where/Data.
type SettingConfigGroup struct {
	g.Meta    `orm:"table:setting_config_group, do:true"`
	Id        interface{} //
	Name      interface{} //
	Code      interface{} //
	CreatedBy interface{} //
	UpdatedBy interface{} //
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
	Remark    interface{} //
}
