// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SystemPost is the golang structure of table system_post for DAO operations like Where/Data.
type SystemPost struct {
	g.Meta    `orm:"table:system_post, do:true"`
	Id        any         // 主键
	Name      any         // 岗位名称
	Code      any         // 岗位代码
	Sort      any         // 排序
	Status    any         // 状态 (1正常 2停用)
	CreatedBy any         // 创建者
	UpdatedBy any         // 更新者
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
	DeletedAt *gtime.Time // 删除时间
	Remark    any         // 备注
}
