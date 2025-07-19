// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SettingGenerateTables is the golang structure for table setting_generate_tables.
type SettingGenerateTables struct {
	Id            int64       `json:"id"            orm:"id"             description:"主键"`                                               // 主键
	TableName     string      `json:"tableName"     orm:"table_name"     description:"表名称"`                                              // 表名称
	TableComment  string      `json:"tableComment"  orm:"table_comment"  description:"表注释"`                                              // 表注释
	ModuleName    string      `json:"moduleName"    orm:"module_name"    description:"所属模块"`                                             // 所属模块
	Namespace     string      `json:"namespace"     orm:"namespace"      description:"命名空间"`                                             // 命名空间
	MenuName      string      `json:"menuName"      orm:"menu_name"      description:"生成菜单名"`                                            // 生成菜单名
	BelongMenuId  int64       `json:"belongMenuId"  orm:"belong_menu_id" description:"所属菜单"`                                             // 所属菜单
	PackageName   string      `json:"packageName"   orm:"package_name"   description:"controller,api包名"`                                 // controller,api包名
	Type          string      `json:"type"          orm:"type"           description:"生成类型，single 单表CRUD，tree 树表CRUD，parent_sub父子表CRUD"` // 生成类型，single 单表CRUD，tree 树表CRUD，parent_sub父子表CRUD
	GenerateType  int         `json:"generateType"  orm:"generate_type"  description:"1 压缩包下载 2 生成到模块"`                                  // 1 压缩包下载 2 生成到模块
	GenerateMenus string      `json:"generateMenus" orm:"generate_menus" description:"生成菜单列表"`                                           // 生成菜单列表
	BuildMenu     int         `json:"buildMenu"     orm:"build_menu"     description:"是否构建菜单"`                                           // 是否构建菜单
	ComponentType int         `json:"componentType" orm:"component_type" description:"组件显示方式"`                                           // 组件显示方式
	Options       string      `json:"options"       orm:"options"        description:"其他业务选项"`                                           // 其他业务选项
	CreatedBy     int64       `json:"createdBy"     orm:"created_by"     description:"创建者"`                                              // 创建者
	UpdatedBy     int64       `json:"updatedBy"     orm:"updated_by"     description:"更新者"`                                              // 更新者
	CreatedAt     *gtime.Time `json:"createdAt"     orm:"created_at"     description:"创建时间"`                                             // 创建时间
	UpdatedAt     *gtime.Time `json:"updatedAt"     orm:"updated_at"     description:"更新时间"`                                             // 更新时间
	Remark        string      `json:"remark"        orm:"remark"         description:"备注"`                                               // 备注
	Source        string      `json:"source"        orm:"source"         description:"db连接群组"`                                           // db连接群组
}
