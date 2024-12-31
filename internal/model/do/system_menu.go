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
	Id        interface{} // 主键
	ParentId  interface{} // 父ID
	Level     interface{} // 组级集合
	Name      interface{} // 菜单名称
	Code      interface{} // 菜单标识代码
	Icon      interface{} // 菜单图标
	Route     interface{} // 路由地址
	Component interface{} // 组件路径
	Redirect  interface{} // 跳转地址
	IsHidden  interface{} // 是否隐藏 (1是 2否)
	Type      interface{} // 菜单类型, (M菜单 B按钮 L链接 I iframe)
	Status    interface{} // 状态 (1正常 2停用)
	Sort      interface{} // 排序
	CreatedBy interface{} // 创建者
	UpdatedBy interface{} // 更新者
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
	DeletedAt *gtime.Time // 删除时间
	Remark    interface{} // 备注
}
