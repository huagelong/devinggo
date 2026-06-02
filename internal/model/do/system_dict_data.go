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
	Id        any         //
	TypeId    any         //
	Label     any         //
	Value     any         //
	Code      any         //
	Sort      any         //
	Status    any         //
	CreatedBy any         //
	UpdatedBy any         //
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
	DeletedAt *gtime.Time //
	Remark    any         //
}
