// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SystemApiLog is the golang structure for table system_api_log.
type SystemApiLog struct {
	Id           int64       `json:"id"           orm:"id"            description:""` //
	ApiId        int64       `json:"apiId"        orm:"api_id"        description:""` //
	ApiName      string      `json:"apiName"      orm:"api_name"      description:""` //
	AccessName   string      `json:"accessName"   orm:"access_name"   description:""` //
	RequestData  string      `json:"requestData"  orm:"request_data"  description:""` //
	ResponseCode string      `json:"responseCode" orm:"response_code" description:""` //
	ResponseData string      `json:"responseData" orm:"response_data" description:""` //
	Ip           string      `json:"ip"           orm:"ip"            description:""` //
	IpLocation   string      `json:"ipLocation"   orm:"ip_location"   description:""` //
	AccessTime   *gtime.Time `json:"accessTime"   orm:"access_time"   description:""` //
	Remark       string      `json:"remark"       orm:"remark"        description:""` //
}
