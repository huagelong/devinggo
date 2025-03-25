// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SystemApi is the golang structure for table system_api.
type SystemApi struct {
	Id          int64       `json:"id"          orm:"id"           description:""` //
	GroupId     int64       `json:"groupId"     orm:"group_id"     description:""` //
	Name        string      `json:"name"        orm:"name"         description:""` //
	AccessName  string      `json:"accessName"  orm:"access_name"  description:""` //
	AuthMode    int         `json:"authMode"    orm:"auth_mode"    description:""` //
	RequestMode string      `json:"requestMode" orm:"request_mode" description:""` //
	Status      int         `json:"status"      orm:"status"       description:""` //
	CreatedBy   int64       `json:"createdBy"   orm:"created_by"   description:""` //
	UpdatedBy   int64       `json:"updatedBy"   orm:"updated_by"   description:""` //
	CreatedAt   *gtime.Time `json:"createdAt"   orm:"created_at"   description:""` //
	UpdatedAt   *gtime.Time `json:"updatedAt"   orm:"updated_at"   description:""` //
	DeletedAt   *gtime.Time `json:"deletedAt"   orm:"deleted_at"   description:""` //
	Remark      string      `json:"remark"      orm:"remark"       description:""` //
}
