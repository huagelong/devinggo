// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SystemDeptLeader is the golang structure of table system_dept_leader for DAO operations like Where/Data.
type SystemDeptLeader struct {
	g.Meta    `orm:"table:system_dept_leader, do:true"`
	DeptId    any         // 部门主键
	UserId    any         // 用户主键
	Username  any         // 用户名
	CreatedAt *gtime.Time // 添加时间
}
