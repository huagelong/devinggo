// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SystemApp is the golang structure of table system_app for DAO operations like Where/Data.
type SystemApp struct {
	g.Meta      `orm:"table:system_app, do:true"`
	Id          any         //
	GroupId     any         //
	AppName     any         //
	AppId       any         //
	AppSecret   any         //
	Status      any         //
	Description any         //
	CreatedBy   any         //
	UpdatedBy   any         //
	CreatedAt   *gtime.Time //
	UpdatedAt   *gtime.Time //
	DeletedAt   *gtime.Time //
	Remark      any         //
}
