// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SystemApp is the golang structure for table system_app.
type SystemApp struct {
	Id          int64       `json:"id"          orm:"id"          description:"主键"`           // 主键
	GroupId     int64       `json:"groupId"     orm:"group_id"    description:"应用组ID"`        // 应用组ID
	AppName     string      `json:"appName"     orm:"app_name"    description:"应用名称"`         // 应用名称
	AppId       string      `json:"appId"       orm:"app_id"      description:"应用ID"`         // 应用ID
	AppSecret   string      `json:"appSecret"   orm:"app_secret"  description:"应用密钥"`         // 应用密钥
	Status      int         `json:"status"      orm:"status"      description:"状态 (1正常 2停用)"` // 状态 (1正常 2停用)
	Description string      `json:"description" orm:"description" description:"应用介绍"`         // 应用介绍
	CreatedBy   int64       `json:"createdBy"   orm:"created_by"  description:"创建者"`          // 创建者
	UpdatedBy   int64       `json:"updatedBy"   orm:"updated_by"  description:"更新者"`          // 更新者
	CreatedAt   *gtime.Time `json:"createdAt"   orm:"created_at"  description:"创建时间"`         // 创建时间
	UpdatedAt   *gtime.Time `json:"updatedAt"   orm:"updated_at"  description:"更新时间"`         // 更新时间
	DeletedAt   *gtime.Time `json:"deletedAt"   orm:"deleted_at"  description:"删除时间"`         // 删除时间
	Remark      string      `json:"remark"      orm:"remark"      description:"备注"`           // 备注
}
