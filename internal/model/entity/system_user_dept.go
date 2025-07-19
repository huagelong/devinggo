// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// SystemUserDept is the golang structure for table system_user_dept.
type SystemUserDept struct {
	UserId int64 `json:"userId" orm:"user_id" description:"用户主键"` // 用户主键
	DeptId int64 `json:"deptId" orm:"dept_id" description:"部门主键"` // 部门主键
}
