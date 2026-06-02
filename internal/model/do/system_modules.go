// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SystemModules is the golang structure of table system_modules for DAO operations like Where/Data.
type SystemModules struct {
	g.Meta      `orm:"table:system_modules, do:true"`
	Id          any         //
	CreatedAt   *gtime.Time //
	UpdatedAt   *gtime.Time //
	CreatedBy   any         //
	UpdatedBy   any         //
	Name        any         //
	Label       any         //
	Description any         //
	Installed   any         //
	Status      any         //
	DeletedAt   *gtime.Time //
}
