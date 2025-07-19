// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SystemApi is the golang structure of table system_api for DAO operations like Where/Data.
type SystemApi struct {
	g.Meta      `orm:"table:system_api, do:true"`
	Id          interface{} // 主键
	GroupId     interface{} // 接口组ID
	Name        interface{} // 接口名称
	AccessName  interface{} // 接口访问名称
	AuthMode    interface{} // 认证模式 (1简易 2复杂)
	RequestMode interface{} // 请求模式 (A 所有 P POST G GET)
	Status      interface{} // 状态 (1正常 2停用)
	CreatedBy   interface{} // 创建者
	UpdatedBy   interface{} // 更新者
	CreatedAt   *gtime.Time // 创建时间
	UpdatedAt   *gtime.Time // 更新时间
	DeletedAt   *gtime.Time // 删除时间
	Remark      interface{} // 备注
}
