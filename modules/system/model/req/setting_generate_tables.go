// Package req
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE
package req

type SettingGenerateTablesSearch struct {
	TableName string `json:"table_name"          description:"表名称"` // 表名称
}

type LoadTable struct {
	Source string           `json:"source"`
	Names  []LoadTableNames `json:"names"`
}

type LoadTableNames struct {
	Name       string `json:"name"`
	Comment    string `json:"comment"`
	SourceName string `json:"sourceName"`
}

type TableAndColumnsUpdate struct {
	Id            int64                          `json:"id"  v:"required" description:"主键"`
	GenerateType  int                            `json:"generate_type"   v:"required#生成类型不能为空" description:"生成类型"`
	BuildMenu     int                            `json:"build_menu"   v:"required#是否生成菜单不能为空" description:"是否生成菜单"`
	GenerateMenus []string                       `json:"generate_menus"   v:"required#生成菜单不能为空" description:"生成菜单"`
	MenuName      string                         `json:"menu_name"   v:"required#菜单名称不能为空" description:"菜单名称"`
	ModuleName    string                         `json:"module_name"   v:"required#模块名称不能为空" description:"模块名称"`
	TableComment  string                         `json:"table_comment"   v:"required#表注释不能为空" description:"表注释"`
	TableName     string                         `json:"table_name"   v:"required#表名称不能为空" description:"表名称"`
	Type          string                         `json:"type"   v:"required#类型不能为空" description:"类型"`
	ComponentType int                            `json:"component_type"   v:"required#组件类型不能为空" description:"组件类型"`
	Columns       []SettingGenerateColumnsUpdate `json:"columns"   v:"required#列不能为空" description:"列"`
	PackageName   string                         `json:"package_name"    description:"包名"`
	BelongMenuId  int64                          `json:"belong_menu_id"  description:"所属菜单ID"`
	Options       string                         `json:"options"    description:"其他配置"`
	Remark        string                         `json:"remark"    description:"备注"`
}
