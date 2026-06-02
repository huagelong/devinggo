// Package res
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package res

import "github.com/gogf/gf/v2/os/gtime"

type SystemLoginLog struct {
	Id         int64       `json:"id"                   description:"主键"`       // 主键
	Username   string      `json:"username"     description:"用户名"`              // 用户名
	Ip         string      `json:"ip"                 description:"登录IP地址"`     // 登录IP地址
	IpLocation string      `json:"ip_location" description:"IP所属地"`             // IP所属地
	Os         string      `json:"os"                 description:"操作系统"`       // 操作系统
	Browser    string      `json:"browser"       description:"浏览器"`             // 浏览器
	Status     int         `json:"status"         description:"登录状态 (1成功 2失败)"` // 登录状态 (1成功 2失败)
	Message    string      `json:"message"       description:"提示消息"`            // 提示消息
	LoginTime  *gtime.Time `json:"login_time"  description:"登录时间"`              // 登录时间
	Remark     string      `json:"remark"         description:"备注"`             // 备注
}
