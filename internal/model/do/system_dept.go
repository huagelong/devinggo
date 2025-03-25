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
	Id        interface{} //
	ParentId  interface{} //
	Level     interface{} //
	Name      interface{} //
	Leader    interface{} //
	Phone     interface{} //
	Status    interface{} //
	Sort      interface{} //
	CreatedBy interface{} //
	UpdatedBy interface{} //
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
	DeletedAt *gtime.Time //
	Remark    interface{} //
}
