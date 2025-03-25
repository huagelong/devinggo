// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SettingGenerateColumns is the golang structure of table setting_generate_columns for DAO operations like Where/Data.
type SettingGenerateColumns struct {
	g.Meta        `orm:"table:setting_generate_columns, do:true"`
	Id            interface{} //
	TableId       interface{} //
	ColumnName    interface{} //
	ColumnComment interface{} //
	ColumnType    interface{} //
	IsPk          interface{} //
	IsRequired    interface{} //
	IsInsert      interface{} //
	IsEdit        interface{} //
	IsList        interface{} //
	IsQuery       interface{} //
	IsSort        interface{} //
	QueryType     interface{} //
	ViewType      interface{} //
	DictType      interface{} //
	AllowRoles    interface{} //
	Options       interface{} //
	Extra         interface{} //
	Sort          interface{} //
	CreatedBy     interface{} //
	UpdatedBy     interface{} //
	CreatedAt     *gtime.Time //
	UpdatedAt     *gtime.Time //
	Remark        interface{} //
}
