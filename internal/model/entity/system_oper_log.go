// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SystemOperLog is the golang structure for table system_oper_log.
type SystemOperLog struct {
	Id           int64       `json:"id"           orm:"id"            description:""` //
	Username     string      `json:"username"     orm:"username"      description:""` //
	Method       string      `json:"method"       orm:"method"        description:""` //
	Router       string      `json:"router"       orm:"router"        description:""` //
	ServiceName  string      `json:"serviceName"  orm:"service_name"  description:""` //
	Ip           string      `json:"ip"           orm:"ip"            description:""` //
	IpLocation   string      `json:"ipLocation"   orm:"ip_location"   description:""` //
	RequestData  string      `json:"requestData"  orm:"request_data"  description:""` //
	ResponseCode string      `json:"responseCode" orm:"response_code" description:""` //
	ResponseData string      `json:"responseData" orm:"response_data" description:""` //
	CreatedBy    int64       `json:"createdBy"    orm:"created_by"    description:""` //
	UpdatedBy    int64       `json:"updatedBy"    orm:"updated_by"    description:""` //
	CreatedAt    *gtime.Time `json:"createdAt"    orm:"created_at"    description:""` //
	UpdatedAt    *gtime.Time `json:"updatedAt"    orm:"updated_at"    description:""` //
	DeletedAt    *gtime.Time `json:"deletedAt"    orm:"deleted_at"    description:""` //
	Remark       string      `json:"remark"       orm:"remark"        description:""` //
}
