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
	Id        any         // 主键
	Name      any         // 配置组名称
	Code      any         // 配置组标识
	CreatedBy any         // 创建者
	UpdatedBy any         // 更新者
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
	Remark    any         // 备注
}
