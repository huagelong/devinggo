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
	Id           interface{} //
	Username     interface{} //
	Method       interface{} //
	Router       interface{} //
	ServiceName  interface{} //
	Ip           interface{} //
	IpLocation   interface{} //
	RequestData  interface{} //
	ResponseCode interface{} //
	ResponseData interface{} //
	CreatedBy    interface{} //
	UpdatedBy    interface{} //
	CreatedAt    *gtime.Time //
	UpdatedAt    *gtime.Time //
	DeletedAt    *gtime.Time //
	Remark       interface{} //
}
