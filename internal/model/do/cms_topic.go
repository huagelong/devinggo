// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// CmsTopic is the golang structure of table cms_topic for DAO operations like Where/Data.
type CmsTopic struct {
	g.Meta    `orm:"table:cms_topic, do:true"`
	Id        interface{} //
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
	DeletedAt *gtime.Time // 删除时间
	Title     interface{} // 标题
	Content   interface{} // 内容
	Status    interface{} // 状态 (1正常 2停用)
	Remark    interface{} // 备注
}
