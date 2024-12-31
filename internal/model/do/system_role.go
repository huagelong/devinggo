// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SystemRole is the golang structure of table system_role for DAO operations like Where/Data.
type SystemRole struct {
	g.Meta    `orm:"table:system_role, do:true"`
	Id        interface{} // 主键
	Name      interface{} // 角色名称
	Code      interface{} // 角色代码
	DataScope interface{} // 数据范围（1：全部数据权限 2：自定义数据权限 3：本部门数据权限 4：本部门及以下数据权限 5：本人数据权限）
	Status    interface{} // 状态 (1正常 2停用)
	Sort      interface{} // 排序
	CreatedBy interface{} // 创建者
	UpdatedBy interface{} // 更新者
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
	DeletedAt *gtime.Time // 删除时间
	Remark    interface{} // 备注
}
