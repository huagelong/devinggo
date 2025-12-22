// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SettingGenerateTables is the golang structure of table setting_generate_tables for DAO operations like Where/Data.
type SettingGenerateTables struct {
	g.Meta        `orm:"table:setting_generate_tables, do:true"`
	Id            any         // 主键
	TableName     any         // 表名称
	TableComment  any         // 表注释
	ModuleName    any         // 所属模块
	Namespace     any         // 命名空间
	MenuName      any         // 生成菜单名
	BelongMenuId  any         // 所属菜单
	PackageName   any         // controller,api包名
	Type          any         // 生成类型，single 单表CRUD，tree 树表CRUD，parent_sub父子表CRUD
	GenerateType  any         // 1 压缩包下载 2 生成到模块
	GenerateMenus any         // 生成菜单列表
	BuildMenu     any         // 是否构建菜单
	ComponentType any         // 组件显示方式
	Options       any         // 其他业务选项
	CreatedBy     any         // 创建者
	UpdatedBy     any         // 更新者
	CreatedAt     *gtime.Time // 创建时间
	UpdatedAt     *gtime.Time // 更新时间
	Remark        any         // 备注
	Source        any         // db连接群组
	TplType       any         // Vue模板类型: default(Arco Design) / ruoyi(RuoYi)
}
