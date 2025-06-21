// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SystemDictData is the golang structure for table system_dict_data.
type SystemDictData struct {
	Id        int64       `json:"id"        orm:"id"         description:"主键"`           // 主键
	TypeId    int64       `json:"typeId"    orm:"type_id"    description:"字典类型ID"`       // 字典类型ID
	Label     string      `json:"label"     orm:"label"      description:"字典标签"`         // 字典标签
	Value     string      `json:"value"     orm:"value"      description:"字典值"`          // 字典值
	Code      string      `json:"code"      orm:"code"       description:"字典标示"`         // 字典标示
	Sort      int         `json:"sort"      orm:"sort"       description:"排序"`           // 排序
	Status    int         `json:"status"    orm:"status"     description:"状态 (1正常 2停用)"` // 状态 (1正常 2停用)
	CreatedBy int64       `json:"createdBy" orm:"created_by" description:"创建者"`          // 创建者
	UpdatedBy int64       `json:"updatedBy" orm:"updated_by" description:"更新者"`          // 更新者
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:""`             //
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:""`             //
	DeletedAt *gtime.Time `json:"deletedAt" orm:"deleted_at" description:"删除时间"`         // 删除时间
	Remark    string      `json:"remark"    orm:"remark"     description:"备注"`           // 备注
}
