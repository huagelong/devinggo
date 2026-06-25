// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SystemNotice is the golang structure of table system_notice for DAO operations like Where/Data.
type SystemNotice struct {
	g.Meta       `orm:"table:system_notice, do:true"`
	Id           any         //
	MessageId    any         //
	Title        any         //
	Type         any         //
	Content      any         //
	CreatedBy    any         //
	UpdatedBy    any         //
	CreatedAt    *gtime.Time //
	UpdatedAt    *gtime.Time //
	DeletedAt    *gtime.Time //
	Remark       any         //
	ReceiveUsers any         //
}
