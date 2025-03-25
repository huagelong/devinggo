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
	Id           interface{} //
	MessageId    interface{} //
	Title        interface{} //
	Type         interface{} //
	Content      interface{} //
	CreatedBy    interface{} //
	UpdatedBy    interface{} //
	CreatedAt    *gtime.Time //
	UpdatedAt    *gtime.Time //
	DeletedAt    *gtime.Time //
	Remark       interface{} //
	ReceiveUsers interface{} //
}
