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
	Id        any         //
	Name      any         //
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
