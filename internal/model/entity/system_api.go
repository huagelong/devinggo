// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SystemApi is the golang structure for table system_api.
type SystemApi struct {
	Id          int64       `json:"id"          orm:"id"           description:"主键"`                       // 主键
	GroupId     int64       `json:"groupId"     orm:"group_id"     description:"接口组ID"`                    // 接口组ID
	Name        string      `json:"name"        orm:"name"         description:"接口名称"`                     // 接口名称
	AccessName  string      `json:"accessName"  orm:"access_name"  description:"接口访问名称"`                   // 接口访问名称
	AuthMode    int         `json:"authMode"    orm:"auth_mode"    description:"认证模式 (1简易 2复杂)"`           // 认证模式 (1简易 2复杂)
	RequestMode string      `json:"requestMode" orm:"request_mode" description:"请求模式 (A 所有 P POST G GET)"` // 请求模式 (A 所有 P POST G GET)
	Status      int         `json:"status"      orm:"status"       description:"状态 (1正常 2停用)"`             // 状态 (1正常 2停用)
	CreatedBy   int64       `json:"createdBy"   orm:"created_by"   description:"创建者"`                      // 创建者
	UpdatedBy   int64       `json:"updatedBy"   orm:"updated_by"   description:"更新者"`                      // 更新者
	CreatedAt   *gtime.Time `json:"createdAt"   orm:"created_at"   description:"创建时间"`                     // 创建时间
	UpdatedAt   *gtime.Time `json:"updatedAt"   orm:"updated_at"   description:"更新时间"`                     // 更新时间
	DeletedAt   *gtime.Time `json:"deletedAt"   orm:"deleted_at"   description:"删除时间"`                     // 删除时间
	Remark      string      `json:"remark"      orm:"remark"       description:"备注"`                       // 备注
}
