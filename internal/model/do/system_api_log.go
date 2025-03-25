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
	Id           interface{} //
	ApiId        interface{} //
	ApiName      interface{} //
	AccessName   interface{} //
	RequestData  interface{} //
	ResponseCode interface{} //
	ResponseData interface{} //
	Ip           interface{} //
	IpLocation   interface{} //
	AccessTime   *gtime.Time //
	Remark       interface{} //
}
