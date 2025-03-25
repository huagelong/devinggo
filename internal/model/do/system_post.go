// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SystemPost is the golang structure of table system_post for DAO operations like Where/Data.
type SystemPost struct {
	g.Meta    `orm:"table:system_post, do:true"`
	Id        interface{} //
	Name      interface{} //
	Code      interface{} //
	Sort      interface{} //
	Status    interface{} //
	CreatedBy interface{} //
	UpdatedBy interface{} //
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
	DeletedAt *gtime.Time //
	Remark    interface{} //
}
