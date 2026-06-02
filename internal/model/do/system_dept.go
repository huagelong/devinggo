// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SystemDept is the golang structure of table system_dept for DAO operations like Where/Data.
type SystemDept struct {
	g.Meta    `orm:"table:system_dept, do:true"`
	Id        any         //
	ParentId  any         //
	Level     any         //
	Name      any         //
	Leader    any         //
	Phone     any         //
	Status    any         //
	Sort      any         //
	CreatedBy any         //
	UpdatedBy any         //
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
	DeletedAt *gtime.Time //
	Remark    any         //
}
