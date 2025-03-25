// Package res
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package res

import "github.com/gogf/gf/v2/os/gtime"

type SystemOperLog struct {
	Id           int64       `json:"id"                      description:"主键"`     // 主键
	Username     string      `json:"username"          description:"用户名"`          // 用户名
	Method       string      `json:"method"              description:"请求方式"`       // 请求方式
	Router       string      `json:"router"              description:"请求路由"`       // 请求路由
	ServiceName  string      `json:"service_name"   description:"业务名称"`            // 业务名称
	Ip           string      `json:"ip"                      description:"请求IP地址"` // 请求IP地址
	IpLocation   string      `json:"ip_location"     description:"IP所属地"`          // IP所属地
	RequestData  string      `json:"request_data"   description:"请求数据"`            // 请求数据
	ResponseCode string      `json:"response_code" description:"响应状态码"`            // 响应状态码
	ResponseData string      `json:"response_data" description:"响应数据"`             // 响应数据
	CreatedBy    int64       `json:"created_by"       description:"创建者"`           // 创建者
	UpdatedBy    int64       `json:"updated_by"       description:"更新者"`           // 更新者
	CreatedAt    *gtime.Time `json:"created_at"       description:"创建时间"`          // 创建时间
	UpdatedAt    *gtime.Time `json:"updated_at"       description:"更新时间"`          // 更新时间
	Remark       string      `json:"remark"              description:"备注"`         // 备注
}
