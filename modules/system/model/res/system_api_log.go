// Package res
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package res

import "github.com/gogf/gf/v2/os/gtime"

type SystemApiLog struct {
	Id           int64       `json:"id"                       description:"主键"`     // 主键
	ApiId        int64       `json:"api_id"                description:"api ID"`    // api ID
	ApiName      string      `json:"api_name"            description:"接口名称"`        // 接口名称
	AccessName   string      `json:"access_name"      description:"接口访问名称"`         // 接口访问名称
	RequestData  string      `json:"request_data"    description:"请求数据"`            // 请求数据
	ResponseCode string      `json:"response_code"  description:"响应状态码"`            // 响应状态码
	ResponseData string      `json:"response_data"  description:"响应数据"`             // 响应数据
	Ip           string      `json:"ip"                       description:"访问IP地址"` // 访问IP地址
	IpLocation   string      `json:"ip_location"      description:"IP所属地"`          // IP所属地
	AccessTime   *gtime.Time `json:"access_time"      description:"访问时间"`           // 访问时间
	Remark       string      `json:"remark"               description:"备注"`         // 备注
}
