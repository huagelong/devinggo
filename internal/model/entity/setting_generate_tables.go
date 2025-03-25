// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SettingGenerateTables is the golang structure for table setting_generate_tables.
type SettingGenerateTables struct {
	Id            int64       `json:"id"            orm:"id"             description:""` //
	TableName     string      `json:"tableName"     orm:"table_name"     description:""` //
	TableComment  string      `json:"tableComment"  orm:"table_comment"  description:""` //
	ModuleName    string      `json:"moduleName"    orm:"module_name"    description:""` //
	Namespace     string      `json:"namespace"     orm:"namespace"      description:""` //
	MenuName      string      `json:"menuName"      orm:"menu_name"      description:""` //
	BelongMenuId  int64       `json:"belongMenuId"  orm:"belong_menu_id" description:""` //
	PackageName   string      `json:"packageName"   orm:"package_name"   description:""` //
	Type          string      `json:"type"          orm:"type"           description:""` //
	GenerateType  int         `json:"generateType"  orm:"generate_type"  description:""` //
	GenerateMenus string      `json:"generateMenus" orm:"generate_menus" description:""` //
	BuildMenu     int         `json:"buildMenu"     orm:"build_menu"     description:""` //
	ComponentType int         `json:"componentType" orm:"component_type" description:""` //
	Options       string      `json:"options"       orm:"options"        description:""` //
	CreatedBy     int64       `json:"createdBy"     orm:"created_by"     description:""` //
	UpdatedBy     int64       `json:"updatedBy"     orm:"updated_by"     description:""` //
	CreatedAt     *gtime.Time `json:"createdAt"     orm:"created_at"     description:""` //
	UpdatedAt     *gtime.Time `json:"updatedAt"     orm:"updated_at"     description:""` //
	Remark        string      `json:"remark"        orm:"remark"         description:""` //
	Source        string      `json:"source"        orm:"source"         description:""` //
}
