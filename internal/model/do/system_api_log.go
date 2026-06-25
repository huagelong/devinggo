// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SystemApiLog is the golang structure of table system_api_log for DAO operations like Where/Data.
type SystemApiLog struct {
	g.Meta       `orm:"table:system_api_log, do:true"`
	Id           any         //
	ApiId        any         //
	ApiName      any         //
	AccessName   any         //
	RequestData  any         //
	ResponseCode any         //
	ResponseData any         //
	Ip           any         //
	IpLocation   any         //
	AccessTime   *gtime.Time //
	Remark       any         //
}
