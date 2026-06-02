// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// SettingConfig is the golang structure of table setting_config for DAO operations like Where/Data.
type SettingConfig struct {
	g.Meta           `orm:"table:setting_config, do:true"`
	GroupId          any // 组id
	Key              any // 配置键名
	Value            any // 配置值
	Name             any //
	InputType        any //
	ConfigSelectData any //
	Sort             any //
	Remark           any //
}
