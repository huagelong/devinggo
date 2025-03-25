// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SystemDictType is the golang structure of table system_dict_type for DAO operations like Where/Data.
type SystemDictType struct {
	g.Meta    `orm:"table:system_dict_type, do:true"`
	Id        interface{} //
	Name      interface{} //
	Code      interface{} //
	Status    interface{} //
	CreatedBy interface{} //
	UpdatedBy interface{} //
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
	DeletedAt *gtime.Time //
	Remark    interface{} //
}
