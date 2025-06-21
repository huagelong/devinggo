// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SystemApiLog is the golang structure for table system_api_log.
type SystemApiLog struct {
	Id           int64       `json:"id"           orm:"id"            description:"主键"`     // 主键
	ApiId        int64       `json:"apiId"        orm:"api_id"        description:"api ID"` // api ID
	ApiName      string      `json:"apiName"      orm:"api_name"      description:"接口名称"`   // 接口名称
	AccessName   string      `json:"accessName"   orm:"access_name"   description:"接口访问名称"` // 接口访问名称
	RequestData  string      `json:"requestData"  orm:"request_data"  description:"请求数据"`   // 请求数据
	ResponseCode string      `json:"responseCode" orm:"response_code" description:"响应状态码"`  // 响应状态码
	ResponseData string      `json:"responseData" orm:"response_data" description:"响应数据"`   // 响应数据
	Ip           string      `json:"ip"           orm:"ip"            description:"访问IP地址"` // 访问IP地址
	IpLocation   string      `json:"ipLocation"   orm:"ip_location"   description:"IP所属地"`  // IP所属地
	AccessTime   *gtime.Time `json:"accessTime"   orm:"access_time"   description:"访问时间"`   // 访问时间
	Remark       string      `json:"remark"       orm:"remark"        description:"备注"`     // 备注
}
