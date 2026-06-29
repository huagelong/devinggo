// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SystemMenu is the golang structure for table system_menu.
type SystemMenu struct {
	Id        int64       `json:"id"        orm:"id"         description:""` //
	ParentId  int64       `json:"parentId"  orm:"parent_id"  description:""` //
	Level     string      `json:"level"     orm:"level"      description:""` //
	Name      string      `json:"name"      orm:"name"       description:""` //
	Code      string      `json:"code"      orm:"code"       description:""` //
	Icon      string      `json:"icon"      orm:"icon"       description:""` //
	Route     string      `json:"route"     orm:"route"      description:""` //
	Component string      `json:"component" orm:"component"  description:""` //
	Redirect  string      `json:"redirect"  orm:"redirect"   description:""` //
	IsHidden  int         `json:"isHidden"  orm:"is_hidden"  description:""` //
	Type      string      `json:"type"      orm:"type"       description:""` //
	Status    int         `json:"status"    orm:"status"     description:""` //
	Sort      int         `json:"sort"      orm:"sort"       description:""` //
	CreatedBy int64       `json:"createdBy" orm:"created_by" description:""` //
	UpdatedBy int64       `json:"updatedBy" orm:"updated_by" description:""` //
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:""` //
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:""` //
	DeletedAt *gtime.Time `json:"deletedAt" orm:"deleted_at" description:""` //
	Remark    string      `json:"remark"    orm:"remark"     description:""` //
}
