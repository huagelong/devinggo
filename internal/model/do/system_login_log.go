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
	Id         interface{} // 主键
	Username   interface{} // 用户名
	Ip         interface{} // 登录IP地址
	IpLocation interface{} // IP所属地
	Os         interface{} // 操作系统
	Browser    interface{} // 浏览器
	Status     interface{} // 登录状态 (1成功 2失败)
	Message    interface{} // 提示消息
	LoginTime  *gtime.Time // 登录时间
	Remark     interface{} // 备注
}
