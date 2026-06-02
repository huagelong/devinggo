// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SystemDeptLeader is the golang structure for table system_dept_leader.
type SystemDeptLeader struct {
	DeptId    int64       `json:"deptId"    orm:"dept_id"    description:""` //
	UserId    int64       `json:"userId"    orm:"user_id"    description:""` //
	Username  string      `json:"username"  orm:"username"   description:""` //
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:""` //
}
