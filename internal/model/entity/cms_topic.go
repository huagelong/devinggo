// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// CmsTopic is the golang structure for table cms_topic.
type CmsTopic struct {
	Id        int64       `json:"id"        orm:"id"         description:""`             //
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:"创建时间"`         // 创建时间
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:"更新时间"`         // 更新时间
	DeletedAt *gtime.Time `json:"deletedAt" orm:"deleted_at" description:"删除时间"`         // 删除时间
	Title     string      `json:"title"     orm:"title"      description:"标题"`           // 标题
	Content   string      `json:"content"   orm:"content"    description:"内容"`           // 内容
	Status    int         `json:"status"    orm:"status"     description:"状态 (1正常 2停用)"` // 状态 (1正常 2停用)
	Remark    string      `json:"remark"    orm:"remark"     description:"备注"`           // 备注
}
