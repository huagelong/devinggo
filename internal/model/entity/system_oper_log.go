// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SystemOperLog is the golang structure for table system_oper_log.
type SystemOperLog struct {
	Id           int64       `json:"id"           orm:"id"            description:"主键"`     // 主键
	Username     string      `json:"username"     orm:"username"      description:"用户名"`    // 用户名
	Method       string      `json:"method"       orm:"method"        description:"请求方式"`   // 请求方式
	Router       string      `json:"router"       orm:"router"        description:"请求路由"`   // 请求路由
	ServiceName  string      `json:"serviceName"  orm:"service_name"  description:"业务名称"`   // 业务名称
	Ip           string      `json:"ip"           orm:"ip"            description:"请求IP地址"` // 请求IP地址
	IpLocation   string      `json:"ipLocation"   orm:"ip_location"   description:"IP所属地"`  // IP所属地
	RequestData  string      `json:"requestData"  orm:"request_data"  description:"请求数据"`   // 请求数据
	ResponseCode string      `json:"responseCode" orm:"response_code" description:"响应状态码"`  // 响应状态码
	ResponseData string      `json:"responseData" orm:"response_data" description:"响应数据"`   // 响应数据
	CreatedBy    int64       `json:"createdBy"    orm:"created_by"    description:"创建者"`    // 创建者
	UpdatedBy    int64       `json:"updatedBy"    orm:"updated_by"    description:"更新者"`    // 更新者
	CreatedAt    *gtime.Time `json:"createdAt"    orm:"created_at"    description:"创建时间"`   // 创建时间
	UpdatedAt    *gtime.Time `json:"updatedAt"    orm:"updated_at"    description:"更新时间"`   // 更新时间
	DeletedAt    *gtime.Time `json:"deletedAt"    orm:"deleted_at"    description:"删除时间"`   // 删除时间
	Remark       string      `json:"remark"       orm:"remark"        description:"备注"`     // 备注
}
