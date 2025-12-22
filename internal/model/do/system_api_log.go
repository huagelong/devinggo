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
	Id           any         // 主键
	ApiId        any         // api ID
	ApiName      any         // 接口名称
	AccessName   any         // 接口访问名称
	RequestData  any         // 请求数据
	ResponseCode any         // 响应状态码
	ResponseData any         // 响应数据
	Ip           any         // 访问IP地址
	IpLocation   any         // IP所属地
	AccessTime   *gtime.Time // 访问时间
	Remark       any         // 备注
}
