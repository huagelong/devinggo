// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SystemDictData is the golang structure of table system_dict_data for DAO operations like Where/Data.
type SystemDictData struct {
	g.Meta    `orm:"table:system_dict_data, do:true"`
	Id        interface{} //
	TypeId    interface{} //
	Label     interface{} //
	Value     interface{} //
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
