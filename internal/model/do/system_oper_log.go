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
	Id           any         // 主键
	Username     any         // 用户名
	Method       any         // 请求方式
	Router       any         // 请求路由
	ServiceName  any         // 业务名称
	Ip           any         // 请求IP地址
	IpLocation   any         // IP所属地
	RequestData  any         // 请求数据
	ResponseCode any         // 响应状态码
	ResponseData any         // 响应数据
	CreatedBy    any         // 创建者
	UpdatedBy    any         // 更新者
	CreatedAt    *gtime.Time // 创建时间
	UpdatedAt    *gtime.Time // 更新时间
	DeletedAt    *gtime.Time // 删除时间
	Remark       any         // 备注
}
