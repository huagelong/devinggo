// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// SystemQueueMessageReceive is the golang structure for table system_queue_message_receive.
type SystemQueueMessageReceive struct {
	MessageId  int64 `json:"messageId"  orm:"message_id"  description:"队列消息主键"`         // 队列消息主键
	UserId     int64 `json:"userId"     orm:"user_id"     description:"接收用户主键"`         // 接收用户主键
	ReadStatus int   `json:"readStatus" orm:"read_status" description:"已读状态 (1未读 2已读)"` // 已读状态 (1未读 2已读)
}
