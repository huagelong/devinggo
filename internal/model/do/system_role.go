// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SystemRole is the golang structure of table system_role for DAO operations like Where/Data.
type SystemRole struct {
	g.Meta    `orm:"table:system_role, do:true"`
	Id        interface{} //
	Name      interface{} //
	Code      interface{} //
	DataScope interface{} //
	Status    interface{} //
	Sort      interface{} //
	CreatedBy interface{} //
	UpdatedBy interface{} //
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
	DeletedAt *gtime.Time //
	Remark    interface{} //
}
