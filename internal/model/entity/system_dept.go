// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SystemDept is the golang structure for table system_dept.
type SystemDept struct {
	Id        int64       `json:"id"        orm:"id"         description:"主键"`           // 主键
	ParentId  int64       `json:"parentId"  orm:"parent_id"  description:"父ID"`          // 父ID
	Level     string      `json:"level"     orm:"level"      description:"组级集合"`         // 组级集合
	Name      string      `json:"name"      orm:"name"       description:"部门名称"`         // 部门名称
	Leader    string      `json:"leader"    orm:"leader"     description:"负责人"`          // 负责人
	Phone     string      `json:"phone"     orm:"phone"      description:"联系电话"`         // 联系电话
	Status    int         `json:"status"    orm:"status"     description:"状态 (1正常 2停用)"` // 状态 (1正常 2停用)
	Sort      int         `json:"sort"      orm:"sort"       description:"排序"`           // 排序
	CreatedBy int64       `json:"createdBy" orm:"created_by" description:"创建者"`          // 创建者
	UpdatedBy int64       `json:"updatedBy" orm:"updated_by" description:"更新者"`          // 更新者
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:"创建时间"`         // 创建时间
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:"更新时间"`         // 更新时间
	DeletedAt *gtime.Time `json:"deletedAt" orm:"deleted_at" description:"删除时间"`         // 删除时间
	Remark    string      `json:"remark"    orm:"remark"     description:"备注"`           // 备注
}
