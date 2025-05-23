// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SystemApp is the golang structure for table system_app.
type SystemApp struct {
	Id          int64       `json:"id"          orm:"id"          description:""` //
	GroupId     int64       `json:"groupId"     orm:"group_id"    description:""` //
	AppName     string      `json:"appName"     orm:"app_name"    description:""` //
	AppId       string      `json:"appId"       orm:"app_id"      description:""` //
	AppSecret   string      `json:"appSecret"   orm:"app_secret"  description:""` //
	Status      int         `json:"status"      orm:"status"      description:""` //
	Description string      `json:"description" orm:"description" description:""` //
	CreatedBy   int64       `json:"createdBy"   orm:"created_by"  description:""` //
	UpdatedBy   int64       `json:"updatedBy"   orm:"updated_by"  description:""` //
	CreatedAt   *gtime.Time `json:"createdAt"   orm:"created_at"  description:""` //
	UpdatedAt   *gtime.Time `json:"updatedAt"   orm:"updated_at"  description:""` //
	DeletedAt   *gtime.Time `json:"deletedAt"   orm:"deleted_at"  description:""` //
	Remark      string      `json:"remark"      orm:"remark"      description:""` //
}
