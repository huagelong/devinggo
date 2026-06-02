// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SettingCrontab is the golang structure of table setting_crontab for DAO operations like Where/Data.
type SettingCrontab struct {
	g.Meta    `orm:"table:setting_crontab, do:true"`
	Id        any         //
	Name      any         //
	Type      any         //
	Target    any         //
	Parameter *gjson.Json //
	Rule      any         //
	Singleton any         //
	Status    any         //
	CreatedBy any         //
	UpdatedBy any         //
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
	Remark    any         //
}
