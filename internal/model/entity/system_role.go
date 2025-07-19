// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SystemRole is the golang structure for table system_role.
type SystemRole struct {
	Id        int64       `json:"id"        orm:"id"         description:"主键"`                                                       // 主键
	Name      string      `json:"name"      orm:"name"       description:"角色名称"`                                                     // 角色名称
	Code      string      `json:"code"      orm:"code"       description:"角色代码"`                                                     // 角色代码
	DataScope int         `json:"dataScope" orm:"data_scope" description:"数据范围（1：全部数据权限 2：自定义数据权限 3：本部门数据权限 4：本部门及以下数据权限 5：本人数据权限）"` // 数据范围（1：全部数据权限 2：自定义数据权限 3：本部门数据权限 4：本部门及以下数据权限 5：本人数据权限）
	Status    int         `json:"status"    orm:"status"     description:"状态 (1正常 2停用)"`                                             // 状态 (1正常 2停用)
	Sort      int         `json:"sort"      orm:"sort"       description:"排序"`                                                       // 排序
	CreatedBy int64       `json:"createdBy" orm:"created_by" description:"创建者"`                                                      // 创建者
	UpdatedBy int64       `json:"updatedBy" orm:"updated_by" description:"更新者"`                                                      // 更新者
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:"创建时间"`                                                     // 创建时间
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:"更新时间"`                                                     // 更新时间
	DeletedAt *gtime.Time `json:"deletedAt" orm:"deleted_at" description:"删除时间"`                                                     // 删除时间
	Remark    string      `json:"remark"    orm:"remark"     description:"备注"`                                                       // 备注
}
