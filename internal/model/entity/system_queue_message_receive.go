// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// SystemQueueMessageReceive is the golang structure for table system_queue_message_receive.
type SystemQueueMessageReceive struct {
	MessageId  int64 `json:"messageId"  orm:"message_id"  description:""` //
	UserId     int64 `json:"userId"     orm:"user_id"     description:""` //
	ReadStatus int   `json:"readStatus" orm:"read_status" description:""` //
}
