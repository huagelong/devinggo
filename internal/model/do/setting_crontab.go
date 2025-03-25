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
	Id        interface{} //
	Name      interface{} //
	Type      interface{} //
	Target    interface{} //
	Parameter *gjson.Json //
	Rule      interface{} //
	Singleton interface{} //
	Status    interface{} //
	CreatedBy interface{} //
	UpdatedBy interface{} //
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
	Remark    interface{} //
}
