// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SystemModules is the golang structure for table system_modules.
type SystemModules struct {
	Id          int64       `json:"id"          orm:"id"          description:""` //
	CreatedAt   *gtime.Time `json:"createdAt"   orm:"created_at"  description:""` //
	UpdatedAt   *gtime.Time `json:"updatedAt"   orm:"updated_at"  description:""` //
	CreatedBy   int64       `json:"createdBy"   orm:"created_by"  description:""` //
	UpdatedBy   int64       `json:"updatedBy"   orm:"updated_by"  description:""` //
	Name        string      `json:"name"        orm:"name"        description:""` //
	Label       string      `json:"label"       orm:"label"       description:""` //
	Description string      `json:"description" orm:"description" description:""` //
	Installed   int         `json:"installed"   orm:"installed"   description:""` //
	Status      int         `json:"status"      orm:"status"      description:""` //
	DeletedAt   *gtime.Time `json:"deletedAt"   orm:"deleted_at"  description:""` //
}
