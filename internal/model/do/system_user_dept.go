// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// SystemUserDept is the golang structure of table system_user_dept for DAO operations like Where/Data.
type SystemUserDept struct {
	g.Meta `orm:"table:system_user_dept, do:true"`
	UserId interface{} // 用户主键
	DeptId interface{} // 部门主键
}
