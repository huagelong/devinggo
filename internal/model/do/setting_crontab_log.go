// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SettingCrontabLog is the golang structure of table setting_crontab_log for DAO operations like Where/Data.
type SettingCrontabLog struct {
	g.Meta        `orm:"table:setting_crontab_log, do:true"`
	Id            interface{} //
	CrontabId     interface{} //
	Name          interface{} //
	Target        interface{} //
	Parameter     interface{} //
	ExceptionInfo interface{} //
	Status        interface{} //
	CreatedAt     *gtime.Time //
}
