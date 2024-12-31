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
	Id           interface{} // 主键
	ApiId        interface{} // api ID
	ApiName      interface{} // 接口名称
	AccessName   interface{} // 接口访问名称
	RequestData  interface{} // 请求数据
	ResponseCode interface{} // 响应状态码
	ResponseData interface{} // 响应数据
	Ip           interface{} // 访问IP地址
	IpLocation   interface{} // IP所属地
	AccessTime   *gtime.Time // 访问时间
	Remark       interface{} // 备注
}
