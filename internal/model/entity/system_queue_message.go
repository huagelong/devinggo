// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SystemQueueMessage is the golang structure for table system_queue_message.
type SystemQueueMessage struct {
	Id          int64       `json:"id"          orm:"id"           description:"主键"`   // 主键
	ContentType string      `json:"contentType" orm:"content_type" description:"内容类型"` // 内容类型
	Title       string      `json:"title"       orm:"title"        description:"消息标题"` // 消息标题
	SendBy      int64       `json:"sendBy"      orm:"send_by"      description:"发送人"`  // 发送人
	Content     string      `json:"content"     orm:"content"      description:"消息内容"` // 消息内容
	CreatedBy   int64       `json:"createdBy"   orm:"created_by"   description:"创建者"`  // 创建者
	UpdatedBy   int64       `json:"updatedBy"   orm:"updated_by"   description:"更新者"`  // 更新者
	CreatedAt   *gtime.Time `json:"createdAt"   orm:"created_at"   description:"创建时间"` // 创建时间
	UpdatedAt   *gtime.Time `json:"updatedAt"   orm:"updated_at"   description:"更新时间"` // 更新时间
	Remark      string      `json:"remark"      orm:"remark"       description:"备注"`   // 备注
}
