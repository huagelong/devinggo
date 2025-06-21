// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// SystemQueueMessageReceive is the golang structure of table system_queue_message_receive for DAO operations like Where/Data.
type SystemQueueMessageReceive struct {
	g.Meta     `orm:"table:system_queue_message_receive, do:true"`
	MessageId  interface{} // 队列消息主键
	UserId     interface{} // 接收用户主键
	ReadStatus interface{} // 已读状态 (1未读 2已读)
}
