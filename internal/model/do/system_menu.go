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
	Id        any         //
	ParentId  any         //
	Level     any         //
	Name      any         //
	Code      any         //
	Icon      any         //
	Route     any         //
	Component any         //
	Redirect  any         //
	IsHidden  any         //
	Type      any         //
	Status    any         //
	Sort      any         //
	CreatedBy any         //
	UpdatedBy any         //
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
	DeletedAt *gtime.Time //
	Remark    any         //
}
