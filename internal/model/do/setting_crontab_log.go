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
	Id            interface{} // 主键
	CrontabId     interface{} // 任务ID
	Name          interface{} // 任务名称
	Target        interface{} // 任务调用目标字符串
	Parameter     interface{} // 任务调用参数
	ExceptionInfo interface{} // 异常信息
	Status        interface{} // 执行状态 (1成功 2失败)
	CreatedAt     *gtime.Time // 创建时间
}
