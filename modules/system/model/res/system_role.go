// Package res
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package res

import "github.com/gogf/gf/v2/os/gtime"

type SystemRole struct {
	Id        int64       `json:"id"                description:"主键"`                                                // 主键
	Name      string      `json:"name"            description:"角色名称"`                                                // 角色名称
	Code      string      `json:"code"            description:"角色代码"`                                                // 角色代码
	DataScope int         `json:"data_scope" description:"数据范围（1：全部数据权限 2：自定义数据权限 3：本部门数据权限 4：本部门及以下数据权限 5：本人数据权限）"` // 数据范围（1：全部数据权限 2：自定义数据权限 3：本部门数据权限 4：本部门及以下数据权限 5：本人数据权限）
	Status    int         `json:"status"        description:"状态 (1正常 2停用)"`                                          // 状态 (1正常 2停用)
	Sort      int         `json:"sort"            description:"排序"`                                                  // 排序
	CreatedBy int64       `json:"created_by" description:"创建者"`                                                      // 创建者
	UpdatedBy int64       `json:"updated_by" description:"更新者"`                                                      // 更新者
	CreatedAt *gtime.Time `json:"created_at" description:"创建时间"`                                                     // 创建时间
	UpdatedAt *gtime.Time `json:"updated_at" description:"更新时间"`                                                     // 更新时间
	Remark    string      `json:"remark"        description:"备注"`                                                    // 备注
}

type SystemRoleMenus struct {
	Id    int64        `json:"id"         description:"主键"`   // 主键
	Menus []MenuIdsArr `json:"menus"      description:"菜单ID"` // 菜单ID
}

type MenuIdsArr struct {
	Id    int64 `json:"id"         description:"主键"`
	Pivot struct {
		RoleId int64 `json:"role_id" description:"角色ID"`
		MenuId int64 `json:"menu_id" description:"菜单ID"`
	} `json:"pivot" description:"关联"` // 关联
}

type SystemRoleDepts struct {
	Id    int64        `json:"id"         description:"主键"`   // 主键
	Depts []DeptIdsArr `json:"depts"      description:"菜单ID"` // 菜单ID
}

type DeptIdsArr struct {
	Id    int64 `json:"id"         description:"主键"`
	Pivot struct {
		RoleId int64 `json:"role_id" description:"角色ID"`
		DeptId int64 `json:"dept_id" description:"部门ID"`
	} `json:"pivot" description:"关联"` // 关联
}
