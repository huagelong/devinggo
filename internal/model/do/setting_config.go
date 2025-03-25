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
	GroupId          interface{} // 组id
	Key              interface{} // 配置键名
	Value            interface{} // 配置值
	Name             interface{} //
	InputType        interface{} //
	ConfigSelectData interface{} //
	Sort             interface{} //
	Remark           interface{} //
}
