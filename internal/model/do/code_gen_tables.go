// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// CodeGenTables is the golang structure of table code_gen_tables for DAO operations like Where/Data.
type CodeGenTables struct {
	g.Meta        `orm:"table:code_gen_tables, do:true"`
	Id            any         //
	TableName     any         // 表名称
	TableComment  any         // 表描述
	Remark        any         // 备注信息
	ModuleName    any         // 所属模块
	BelongMenuId  any         // 所属菜单ID
	Type          any         // 生成类型: single=单表, tree=树表
	MenuName      any         // 菜单名称
	ComponentType any         // 组件类型: 1=模态框, 2=抽屉, 3=Tag页
	TplType       any         // 模板类型: default
	TreeId        any         // 树表主ID字段
	TreeParentId  any         // 树表父ID字段
	TreeName      any         // 树表显示名称字段
	TagId         any         // Tag页ID
	TagName       any         // Tag页名称
	TagViewName   any         // Tag页显示字段
	GenerateMenus any         // 生成的菜单按钮
	Options       *gjson.Json // 扩展配置
	CreatedAt     *gtime.Time // 创建时间
	UpdatedAt     *gtime.Time // 更新时间
	DeletedAt     *gtime.Time // 删除时间
	CreatedBy     any         // 创建者ID
	UpdatedBy     any         // 更新者ID
	Status        any         // 状态: 1=正常, 0=停用
	Sort          any         // 排序
}
