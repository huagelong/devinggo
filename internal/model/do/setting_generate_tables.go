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
	Id            interface{} // 主键
	TableName     interface{} // 表名称
	TableComment  interface{} // 表注释
	ModuleName    interface{} // 所属模块
	Namespace     interface{} // 命名空间
	MenuName      interface{} // 生成菜单名
	BelongMenuId  interface{} // 所属菜单
	PackageName   interface{} // controller,api包名
	Type          interface{} // 生成类型，single 单表CRUD，tree 树表CRUD，parent_sub父子表CRUD
	GenerateType  interface{} // 1 压缩包下载 2 生成到模块
	GenerateMenus interface{} // 生成菜单列表
	BuildMenu     interface{} // 是否构建菜单
	ComponentType interface{} // 组件显示方式
	Options       interface{} // 其他业务选项
	CreatedBy     interface{} // 创建者
	UpdatedBy     interface{} // 更新者
	CreatedAt     *gtime.Time // 创建时间
	UpdatedAt     *gtime.Time // 更新时间
	Remark        interface{} // 备注
	Source        interface{} // db连接群组
}
