// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SystemMenu is the golang structure for table system_menu.
type SystemMenu struct {
	Id        int64       `json:"id"        orm:"id"         description:"主键"`                           // 主键
	ParentId  int64       `json:"parentId"  orm:"parent_id"  description:"父ID"`                          // 父ID
	Level     string      `json:"level"     orm:"level"      description:"组级集合"`                         // 组级集合
	Name      string      `json:"name"      orm:"name"       description:"菜单名称"`                         // 菜单名称
	Code      string      `json:"code"      orm:"code"       description:"菜单标识代码"`                       // 菜单标识代码
	Icon      string      `json:"icon"      orm:"icon"       description:"菜单图标"`                         // 菜单图标
	Route     string      `json:"route"     orm:"route"      description:"路由地址"`                         // 路由地址
	Component string      `json:"component" orm:"component"  description:"组件路径"`                         // 组件路径
	Redirect  string      `json:"redirect"  orm:"redirect"   description:"跳转地址"`                         // 跳转地址
	IsHidden  int         `json:"isHidden"  orm:"is_hidden"  description:"是否隐藏 (1是 2否)"`                 // 是否隐藏 (1是 2否)
	Type      string      `json:"type"      orm:"type"       description:"菜单类型, (M菜单 B按钮 L链接 I iframe)"` // 菜单类型, (M菜单 B按钮 L链接 I iframe)
	Status    int         `json:"status"    orm:"status"     description:"状态 (1正常 2停用)"`                 // 状态 (1正常 2停用)
	Sort      int         `json:"sort"      orm:"sort"       description:"排序"`                           // 排序
	CreatedBy int64       `json:"createdBy" orm:"created_by" description:"创建者"`                          // 创建者
	UpdatedBy int64       `json:"updatedBy" orm:"updated_by" description:"更新者"`                          // 更新者
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:""`                             //
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:""`                             //
	DeletedAt *gtime.Time `json:"deletedAt" orm:"deleted_at" description:"删除时间"`                         // 删除时间
	Remark    string      `json:"remark"    orm:"remark"     description:"备注"`                           // 备注
}
