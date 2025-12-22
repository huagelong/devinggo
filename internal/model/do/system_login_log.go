// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SystemLoginLog is the golang structure of table system_login_log for DAO operations like Where/Data.
type SystemLoginLog struct {
	g.Meta     `orm:"table:system_login_log, do:true"`
	Id         any         // 主键
	Username   any         // 用户名
	Ip         any         // 登录IP地址
	IpLocation any         // IP所属地
	Os         any         // 操作系统
	Browser    any         // 浏览器
	Status     any         // 登录状态 (1成功 2失败)
	Message    any         // 提示消息
	LoginTime  *gtime.Time // 登录时间
	Remark     any         // 备注
}
