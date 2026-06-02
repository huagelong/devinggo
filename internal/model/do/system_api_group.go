// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SystemApiGroup is the golang structure of table system_api_group for DAO operations like Where/Data.
type SystemApiGroup struct {
	g.Meta    `orm:"table:system_api_group, do:true"`
	Id        any         //
	Name      any         //
	Status    any         //
	CreatedBy any         //
	UpdatedBy any         //
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
	DeletedAt *gtime.Time //
	Remark    any         //
}
