// Package res
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE
package res

import "github.com/gogf/gf/v2/os/gtime"

type SettingGenerateTables struct {
	Id            int64       `json:"id"                         description:"主键"`                                             // 主键
	TableName     string      `json:"table_name"          description:"表名称"`                                                   // 表名称
	TableComment  string      `json:"table_comment"    description:"表注释"`                                                      // 表注释
	ModuleName    string      `json:"module_name"        description:"所属模块"`                                                   // 所属模块
	Namespace     string      `json:"namespace"           description:"命名空间"`                                                  // 命名空间
	MenuName      string      `json:"menu_name"            description:"生成菜单名"`                                                // 生成菜单名
	BelongMenuId  int64       `json:"belong_menu_id"   description:"所属菜单"`                                                     // 所属菜单
	PackageName   string      `json:"package_name"      description:"控制器包名"`                                                   // 控制器包名
	Type          string      `json:"type"                     description:"生成类型，single 单表CRUD，tree 树表CRUD，parent_sub父子表CRUD"` // 生成类型，single 单表CRUD，tree 树表CRUD，parent_sub父子表CRUD
	GenerateType  int         `json:"generate_type"    description:"1 压缩包下载 2 生成到模块"`                                          // 1 压缩包下载 2 生成到模块
	GenerateMenus string      `json:"generate_menus"  description:"生成菜单列表"`                                                    // 生成菜单列表
	BuildMenu     int         `json:"build_menu"          description:"是否构建菜单"`                                                // 是否构建菜单
	ComponentType int         `json:"component_type"  description:"组件显示方式"`                                                    // 组件显示方式
	Options       string      `json:"options"               description:"其他业务选项"`                                              // 其他业务选项
	CreatedBy     int64       `json:"created_by"          description:"创建者"`                                                   // 创建者
	UpdatedBy     int64       `json:"updated_by"          description:"更新者"`                                                   // 更新者
	CreatedAt     *gtime.Time `json:"created_at"          description:"创建时间"`                                                  // 创建时间
	UpdatedAt     *gtime.Time `json:"updated_at"          description:"更新时间"`                                                  // 更新时间
	Remark        string      `json:"remark"                 description:"备注"`                                                 // 备注
}

type PreviewTable struct {
	TabName string `json:"tab_name"`
	Name    string `json:"name"`
	Code    string `json:"code"`
	Lang    string `json:"lang"`
}
