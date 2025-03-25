// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SystemNotice is the golang structure for table system_notice.
type SystemNotice struct {
	Id           int64       `json:"id"           orm:"id"            description:""` //
	MessageId    int64       `json:"messageId"    orm:"message_id"    description:""` //
	Title        string      `json:"title"        orm:"title"         description:""` //
	Type         int         `json:"type"         orm:"type"          description:""` //
	Content      string      `json:"content"      orm:"content"       description:""` //
	CreatedBy    int64       `json:"createdBy"    orm:"created_by"    description:""` //
	UpdatedBy    int64       `json:"updatedBy"    orm:"updated_by"    description:""` //
	CreatedAt    *gtime.Time `json:"createdAt"    orm:"created_at"    description:""` //
	UpdatedAt    *gtime.Time `json:"updatedAt"    orm:"updated_at"    description:""` //
	DeletedAt    *gtime.Time `json:"deletedAt"    orm:"deleted_at"    description:""` //
	Remark       string      `json:"remark"       orm:"remark"        description:""` //
	ReceiveUsers string      `json:"receiveUsers" orm:"receive_users" description:""` //
}
