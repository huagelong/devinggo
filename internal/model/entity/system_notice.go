// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SystemNotice is the golang structure for table system_notice.
type SystemNotice struct {
	Id           int64       `json:"id"           orm:"id"            description:"主键"`            // 主键
	MessageId    int64       `json:"messageId"    orm:"message_id"    description:"消息ID"`          // 消息ID
	Title        string      `json:"title"        orm:"title"         description:"标题"`            // 标题
	Type         int         `json:"type"         orm:"type"          description:"公告类型（1通知 2公告）"` // 公告类型（1通知 2公告）
	Content      string      `json:"content"      orm:"content"       description:"公告内容"`          // 公告内容
	CreatedBy    int64       `json:"createdBy"    orm:"created_by"    description:"创建者"`           // 创建者
	UpdatedBy    int64       `json:"updatedBy"    orm:"updated_by"    description:"更新者"`           // 更新者
	CreatedAt    *gtime.Time `json:"createdAt"    orm:"created_at"    description:"创建时间"`          // 创建时间
	UpdatedAt    *gtime.Time `json:"updatedAt"    orm:"updated_at"    description:"更新时间"`          // 更新时间
	DeletedAt    *gtime.Time `json:"deletedAt"    orm:"deleted_at"    description:"删除时间"`          // 删除时间
	Remark       string      `json:"remark"       orm:"remark"        description:"备注"`            // 备注
	ReceiveUsers string      `json:"receiveUsers" orm:"receive_users" description:"接收用户id,隔开"`     // 接收用户id,隔开
}
