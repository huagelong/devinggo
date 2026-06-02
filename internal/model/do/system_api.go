// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SystemApi is the golang structure of table system_api for DAO operations like Where/Data.
type SystemApi struct {
	g.Meta      `orm:"table:system_api, do:true"`
	Id          any         //
	GroupId     any         //
	Name        any         //
	AccessName  any         //
	AuthMode    any         //
	RequestMode any         //
	Status      any         //
	CreatedBy   any         //
	UpdatedBy   any         //
	CreatedAt   *gtime.Time //
	UpdatedAt   *gtime.Time //
	DeletedAt   *gtime.Time //
	Remark      any         //
}
