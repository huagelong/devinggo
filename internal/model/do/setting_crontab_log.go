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
	Id            any         // 主键
	CrontabId     any         // 任务ID
	Name          any         // 任务名称
	Target        any         // 任务调用目标字符串
	Parameter     any         // 任务调用参数
	ExceptionInfo any         // 异常信息
	Status        any         // 执行状态 (1成功 2失败)
	CreatedAt     *gtime.Time // 创建时间
}
