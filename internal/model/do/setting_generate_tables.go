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
	Id            interface{} //
	TableName     interface{} //
	TableComment  interface{} //
	ModuleName    interface{} //
	Namespace     interface{} //
	MenuName      interface{} //
	BelongMenuId  interface{} //
	PackageName   interface{} //
	Type          interface{} //
	GenerateType  interface{} //
	GenerateMenus interface{} //
	BuildMenu     interface{} //
	ComponentType interface{} //
	Options       interface{} //
	CreatedBy     interface{} //
	UpdatedBy     interface{} //
	CreatedAt     *gtime.Time //
	UpdatedAt     *gtime.Time //
	Remark        interface{} //
	Source        interface{} //
}
