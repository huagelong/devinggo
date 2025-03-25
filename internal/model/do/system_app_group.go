// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SystemAppGroup is the golang structure of table system_app_group for DAO operations like Where/Data.
type SystemAppGroup struct {
	g.Meta    `orm:"table:system_app_group, do:true"`
	Id        interface{} //
	Name      interface{} //
	Status    interface{} //
	CreatedBy interface{} //
	UpdatedBy interface{} //
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
	DeletedAt *gtime.Time //
	Remark    interface{} //
}
