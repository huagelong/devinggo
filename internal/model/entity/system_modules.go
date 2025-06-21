// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SystemModules is the golang structure for table system_modules.
type SystemModules struct {
	Id          int64       `json:"id"          orm:"id"          description:"主键"`           // 主键
	CreatedAt   *gtime.Time `json:"createdAt"   orm:"created_at"  description:"创建时间"`         // 创建时间
	UpdatedAt   *gtime.Time `json:"updatedAt"   orm:"updated_at"  description:"更新时间"`         // 更新时间
	CreatedBy   int64       `json:"createdBy"   orm:"created_by"  description:"创建者"`          // 创建者
	UpdatedBy   int64       `json:"updatedBy"   orm:"updated_by"  description:"更新者"`          // 更新者
	Name        string      `json:"name"        orm:"name"        description:"模块名称"`         // 模块名称
	Label       string      `json:"label"       orm:"label"       description:"模块标记"`         // 模块标记
	Description string      `json:"description" orm:"description" description:"描述"`           // 描述
	Installed   int         `json:"installed"   orm:"installed"   description:"是否安装1-否，2-是"`  // 是否安装1-否，2-是
	Status      int         `json:"status"      orm:"status"      description:"状态 (1正常 2停用)"` // 状态 (1正常 2停用)
	DeletedAt   *gtime.Time `json:"deletedAt"   orm:"deleted_at"  description:"删除时间"`         // 删除时间
}
