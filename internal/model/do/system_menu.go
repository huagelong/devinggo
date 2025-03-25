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
	Id        interface{} //
	ParentId  interface{} //
	Level     interface{} //
	Name      interface{} //
	Code      interface{} //
	Icon      interface{} //
	Route     interface{} //
	Component interface{} //
	Redirect  interface{} //
	IsHidden  interface{} //
	Type      interface{} //
	Status    interface{} //
	Sort      interface{} //
	CreatedBy interface{} //
	UpdatedBy interface{} //
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
	DeletedAt *gtime.Time //
	Remark    interface{} //
}
