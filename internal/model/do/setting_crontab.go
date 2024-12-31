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
	Id        interface{} // 主键
	Name      interface{} // 任务名称
	Type      interface{} // 任务类型 (1 command, 2 class, 3 url, 4 eval)
	Target    interface{} // 调用任务字符串
	Parameter interface{} // 调用任务参数
	Rule      interface{} // 任务执行表达式
	Singleton interface{} // 是否单次执行 (1 是 2 不是)
	Status    interface{} // 状态 (1正常 2停用)
	CreatedBy interface{} // 创建者
	UpdatedBy interface{} // 更新者
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
	Remark    interface{} // 备注
}
