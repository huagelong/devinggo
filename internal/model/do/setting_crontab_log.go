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
	Id            any         //
	CrontabId     any         //
	Name          any         //
	Target        any         //
	Parameter     any         //
	ExceptionInfo any         //
	Status        any         //
	CreatedAt     *gtime.Time //
	StartTime     *gtime.Time // 任务开始执行时间
	EndTime       *gtime.Time // 任务执行结束时间
	Output        any         // 任务执行输出
}
