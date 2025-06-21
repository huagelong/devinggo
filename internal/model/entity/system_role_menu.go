// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// SystemRoleMenu is the golang structure for table system_role_menu.
type SystemRoleMenu struct {
	RoleId int64 `json:"roleId" orm:"role_id" description:"角色主键"` // 角色主键
	MenuId int64 `json:"menuId" orm:"menu_id" description:"菜单主键"` // 菜单主键
}
