// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SettingCrontab is the golang structure of table setting_crontab for DAO operations like Where/Data.
type SettingCrontab struct {
	g.Meta    `orm:"table:setting_crontab, do:true"`
	Id        any         // 主键
	Name      any         // 任务名称
	Type      any         // 任务类型 (1 command, 2 class, 3 url, 4 eval)
	Target    any         // 调用任务字符串
	Parameter any         // 调用任务参数
	Rule      any         // 任务执行表达式
	Singleton any         // 是否单次执行 (1 是 2 不是)
	Status    any         // 状态 (1正常 2停用)
	CreatedBy any         // 创建者
	UpdatedBy any         // 更新者
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
	Remark    any         // 备注
}
