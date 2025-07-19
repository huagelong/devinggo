// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// SystemUserRole is the golang structure for table system_user_role.
type SystemUserRole struct {
	UserId int64 `json:"userId" orm:"user_id" description:"用户主键"` // 用户主键
	RoleId int64 `json:"roleId" orm:"role_id" description:"角色主键"` // 角色主键
}
