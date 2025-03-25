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
	Id          interface{} //
	ContentType interface{} //
	Title       interface{} //
	SendBy      interface{} //
	Content     interface{} //
	CreatedBy   interface{} //
	UpdatedBy   interface{} //
	CreatedAt   *gtime.Time //
	UpdatedAt   *gtime.Time //
	Remark      interface{} //
}
