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
	Id          any         // 主键
	GroupId     any         // 接口组ID
	Name        any         // 接口名称
	AccessName  any         // 接口访问名称
	AuthMode    any         // 认证模式 (1简易 2复杂)
	RequestMode any         // 请求模式 (A 所有 P POST G GET)
	Status      any         // 状态 (1正常 2停用)
	CreatedBy   any         // 创建者
	UpdatedBy   any         // 更新者
	CreatedAt   *gtime.Time // 创建时间
	UpdatedAt   *gtime.Time // 更新时间
	DeletedAt   *gtime.Time // 删除时间
	Remark      any         // 备注
}
