// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SystemDept is the golang structure of table system_dept for DAO operations like Where/Data.
type SystemDept struct {
	g.Meta    `orm:"table:system_dept, do:true"`
	Id        any         // 主键
	ParentId  any         // 父ID
	Level     any         // 组级集合
	Name      any         // 部门名称
	Leader    any         // 负责人
	Phone     any         // 联系电话
	Status    any         // 状态 (1正常 2停用)
	Sort      any         // 排序
	CreatedBy any         // 创建者
	UpdatedBy any         // 更新者
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
	DeletedAt *gtime.Time // 删除时间
	Remark    any         // 备注
}
