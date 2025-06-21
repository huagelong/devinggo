// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SystemAppGroup is the golang structure for table system_app_group.
type SystemAppGroup struct {
	Id        int64       `json:"id"        orm:"id"         description:"主键"`           // 主键
	Name      string      `json:"name"      orm:"name"       description:"应用组名称"`        // 应用组名称
	Status    int         `json:"status"    orm:"status"     description:"状态 (1正常 2停用)"` // 状态 (1正常 2停用)
	CreatedBy int64       `json:"createdBy" orm:"created_by" description:"创建者"`          // 创建者
	UpdatedBy int64       `json:"updatedBy" orm:"updated_by" description:"更新者"`          // 更新者
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:"创建时间"`         // 创建时间
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:"更新时间"`         // 更新时间
	DeletedAt *gtime.Time `json:"deletedAt" orm:"deleted_at" description:"删除时间"`         // 删除时间
	Remark    string      `json:"remark"    orm:"remark"     description:"备注"`           // 备注
}
