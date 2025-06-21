// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SystemDeptLeader is the golang structure for table system_dept_leader.
type SystemDeptLeader struct {
	DeptId    int64       `json:"deptId"    orm:"dept_id"    description:"部门主键"` // 部门主键
	UserId    int64       `json:"userId"    orm:"user_id"    description:"用户主键"` // 用户主键
	Username  string      `json:"username"  orm:"username"   description:"用户名"`  // 用户名
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:"添加时间"` // 添加时间
}
