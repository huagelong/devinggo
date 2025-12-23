// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SystemApp is the golang structure of table system_app for DAO operations like Where/Data.
type SystemApp struct {
	g.Meta      `orm:"table:system_app, do:true"`
	Id          any         // 主键
	GroupId     any         // 应用组ID
	AppName     any         // 应用名称
	AppId       any         // 应用ID
	AppSecret   any         // 应用密钥
	Status      any         // 状态 (1正常 2停用)
	Description any         // 应用介绍
	CreatedBy   any         // 创建者
	UpdatedBy   any         // 更新者
	CreatedAt   *gtime.Time // 创建时间
	UpdatedAt   *gtime.Time // 更新时间
	DeletedAt   *gtime.Time // 删除时间
	Remark      any         // 备注
}
