// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/os/gtime"
)

// SystemUser is the golang structure for table system_user.
type SystemUser struct {
	Id             int64       `json:"id"             orm:"id"              description:""` //
	Username       string      `json:"username"       orm:"username"        description:""` //
	Password       string      `json:"password"       orm:"password"        description:""` //
	UserType       string      `json:"userType"       orm:"user_type"       description:""` //
	Nickname       string      `json:"nickname"       orm:"nickname"        description:""` //
	Phone          string      `json:"phone"          orm:"phone"           description:""` //
	Email          string      `json:"email"          orm:"email"           description:""` //
	Avatar         string      `json:"avatar"         orm:"avatar"          description:""` //
	Signed         string      `json:"signed"         orm:"signed"          description:""` //
	Dashboard      string      `json:"dashboard"      orm:"dashboard"       description:""` //
	Status         int         `json:"status"         orm:"status"          description:""` //
	LoginIp        string      `json:"loginIp"        orm:"login_ip"        description:""` //
	LoginTime      *gtime.Time `json:"loginTime"      orm:"login_time"      description:""` //
	BackendSetting *gjson.Json `json:"backendSetting" orm:"backend_setting" description:""` //
	CreatedBy      int64       `json:"createdBy"      orm:"created_by"      description:""` //
	UpdatedBy      int64       `json:"updatedBy"      orm:"updated_by"      description:""` //
	CreatedAt      *gtime.Time `json:"createdAt"      orm:"created_at"      description:""` //
	UpdatedAt      *gtime.Time `json:"updatedAt"      orm:"updated_at"      description:""` //
	DeletedAt      *gtime.Time `json:"deletedAt"      orm:"deleted_at"      description:""` //
	Remark         string      `json:"remark"         orm:"remark"          description:""` //
}
