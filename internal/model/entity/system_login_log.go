// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SystemLoginLog is the golang structure for table system_login_log.
type SystemLoginLog struct {
	Id         int64       `json:"id"         orm:"id"          description:"主键"`             // 主键
	Username   string      `json:"username"   orm:"username"    description:"用户名"`            // 用户名
	Ip         string      `json:"ip"         orm:"ip"          description:"登录IP地址"`         // 登录IP地址
	IpLocation string      `json:"ipLocation" orm:"ip_location" description:"IP所属地"`          // IP所属地
	Os         string      `json:"os"         orm:"os"          description:"操作系统"`           // 操作系统
	Browser    string      `json:"browser"    orm:"browser"     description:"浏览器"`            // 浏览器
	Status     int         `json:"status"     orm:"status"      description:"登录状态 (1成功 2失败)"` // 登录状态 (1成功 2失败)
	Message    string      `json:"message"    orm:"message"     description:"提示消息"`           // 提示消息
	LoginTime  *gtime.Time `json:"loginTime"  orm:"login_time"  description:"登录时间"`           // 登录时间
	Remark     string      `json:"remark"     orm:"remark"      description:"备注"`             // 备注
}
