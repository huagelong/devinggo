// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SystemQueueMessage is the golang structure of table system_queue_message for DAO operations like Where/Data.
type SystemQueueMessage struct {
	g.Meta      `orm:"table:system_queue_message, do:true"`
	Id          interface{} // 主键
	ContentType interface{} // 内容类型
	Title       interface{} // 消息标题
	SendBy      interface{} // 发送人
	Content     interface{} // 消息内容
	CreatedBy   interface{} // 创建者
	UpdatedBy   interface{} // 更新者
	CreatedAt   *gtime.Time // 创建时间
	UpdatedAt   *gtime.Time // 更新时间
	Remark      interface{} // 备注
}
