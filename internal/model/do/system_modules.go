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
	Id          interface{} // 主键
	CreatedAt   *gtime.Time // 创建时间
	UpdatedAt   *gtime.Time // 更新时间
	CreatedBy   interface{} // 创建者
	UpdatedBy   interface{} // 更新者
	Name        interface{} // 模块名称
	Label       interface{} // 模块标记
	Description interface{} // 描述
	Installed   interface{} // 是否安装1-否，2-是
	Status      interface{} // 状态 (1正常 2停用)
	DeletedAt   *gtime.Time // 删除时间
}
