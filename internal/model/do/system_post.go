// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SystemPost is the golang structure of table system_post for DAO operations like Where/Data.
type SystemPost struct {
	g.Meta    `orm:"table:system_post, do:true"`
	Id        interface{} // 主键
	Name      interface{} // 岗位名称
	Code      interface{} // 岗位代码
	Sort      interface{} // 排序
	Status    interface{} // 状态 (1正常 2停用)
	CreatedBy interface{} // 创建者
	UpdatedBy interface{} // 更新者
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
	DeletedAt *gtime.Time // 删除时间
	Remark    interface{} // 备注
}
