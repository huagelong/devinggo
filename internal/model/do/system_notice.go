// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SystemNotice is the golang structure of table system_notice for DAO operations like Where/Data.
type SystemNotice struct {
	g.Meta       `orm:"table:system_notice, do:true"`
	Id           any         // 主键
	MessageId    any         // 消息ID
	Title        any         // 标题
	Type         any         // 公告类型（1通知 2公告）
	Content      any         // 公告内容
	CreatedBy    any         // 创建者
	UpdatedBy    any         // 更新者
	CreatedAt    *gtime.Time // 创建时间
	UpdatedAt    *gtime.Time // 更新时间
	DeletedAt    *gtime.Time // 删除时间
	Remark       any         // 备注
	ReceiveUsers any         // 接收用户id,隔开
}
