// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SystemOperLog is the golang structure of table system_oper_log for DAO operations like Where/Data.
type SystemOperLog struct {
	g.Meta       `orm:"table:system_oper_log, do:true"`
	Id           any         //
	Username     any         //
	Method       any         //
	Router       any         //
	ServiceName  any         //
	Ip           any         //
	IpLocation   any         //
	RequestData  any         //
	ResponseCode any         //
	ResponseData any         //
	CreatedBy    any         //
	UpdatedBy    any         //
	CreatedAt    *gtime.Time //
	UpdatedAt    *gtime.Time //
	DeletedAt    *gtime.Time //
	Remark       any         //
}
