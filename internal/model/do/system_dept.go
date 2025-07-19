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
	Id        interface{} // 主键
	ParentId  interface{} // 父ID
	Level     interface{} // 组级集合
	Name      interface{} // 部门名称
	Leader    interface{} // 负责人
	Phone     interface{} // 联系电话
	Status    interface{} // 状态 (1正常 2停用)
	Sort      interface{} // 排序
	CreatedBy interface{} // 创建者
	UpdatedBy interface{} // 更新者
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
	DeletedAt *gtime.Time // 删除时间
	Remark    interface{} // 备注
}
