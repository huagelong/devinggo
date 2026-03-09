// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SystemModules is the golang structure of table system_modules for DAO operations like Where/Data.
type SystemModules struct {
	g.Meta      `orm:"table:system_modules, do:true"`
	Id          any         // 主键
	CreatedAt   *gtime.Time // 创建时间
	UpdatedAt   *gtime.Time // 更新时间
	CreatedBy   any         // 创建者
	UpdatedBy   any         // 更新者
	Name        any         // 模块名称
	Label       any         // 模块标记
	Description any         // 描述
	Installed   any         // 是否安装1-否，2-是
	Status      any         // 状态 (1正常 2停用)
	DeletedAt   *gtime.Time // 删除时间
}
