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
	Id           interface{} // 主键
	Username     interface{} // 用户名
	Method       interface{} // 请求方式
	Router       interface{} // 请求路由
	ServiceName  interface{} // 业务名称
	Ip           interface{} // 请求IP地址
	IpLocation   interface{} // IP所属地
	RequestData  interface{} // 请求数据
	ResponseCode interface{} // 响应状态码
	ResponseData interface{} // 响应数据
	CreatedBy    interface{} // 创建者
	UpdatedBy    interface{} // 更新者
	CreatedAt    *gtime.Time // 创建时间
	UpdatedAt    *gtime.Time // 更新时间
	DeletedAt    *gtime.Time // 删除时间
	Remark       interface{} // 备注
}
