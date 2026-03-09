// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SystemMenu is the golang structure of table system_menu for DAO operations like Where/Data.
type SystemMenu struct {
	g.Meta    `orm:"table:system_menu, do:true"`
	Id        any         // 主键
	ParentId  any         // 父ID
	Level     any         // 组级集合
	Name      any         // 菜单名称
	Code      any         // 菜单标识代码
	Icon      any         // 菜单图标
	Route     any         // 路由地址
	Component any         // 组件路径
	Redirect  any         // 跳转地址
	IsHidden  any         // 是否隐藏 (1是 2否)
	Type      any         // 菜单类型, (M菜单 B按钮 L链接 I iframe)
	Status    any         // 状态 (1正常 2停用)
	Sort      any         // 排序
	CreatedBy any         // 创建者
	UpdatedBy any         // 更新者
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
	DeletedAt *gtime.Time // 删除时间
	Remark    any         // 备注
}
