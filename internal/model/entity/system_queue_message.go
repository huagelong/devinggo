// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SystemQueueMessage is the golang structure for table system_queue_message.
type SystemQueueMessage struct {
	Id          int64       `json:"id"          orm:"id"           description:""` //
	ContentType string      `json:"contentType" orm:"content_type" description:""` //
	Title       string      `json:"title"       orm:"title"        description:""` //
	SendBy      int64       `json:"sendBy"      orm:"send_by"      description:""` //
	Content     string      `json:"content"     orm:"content"      description:""` //
	CreatedBy   int64       `json:"createdBy"   orm:"created_by"   description:""` //
	UpdatedBy   int64       `json:"updatedBy"   orm:"updated_by"   description:""` //
	CreatedAt   *gtime.Time `json:"createdAt"   orm:"created_at"   description:""` //
	UpdatedAt   *gtime.Time `json:"updatedAt"   orm:"updated_at"   description:""` //
	Remark      string      `json:"remark"      orm:"remark"       description:""` //
}
