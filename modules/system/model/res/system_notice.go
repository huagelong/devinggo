// Package res
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package res

import "github.com/gogf/gf/v2/os/gtime"

type SystemNotice struct {
	Id        int64       `json:"id"                description:"主键"`        // 主键
	MessageId int64       `json:"message_id"  description:"消息ID"`            // 消息ID
	Title     string      `json:"title"           description:"标题"`          // 标题
	Type      int         `json:"type"          description:"公告类型（1通知 2公告）"` // 公告类型（1通知 2公告）
	Content   string      `json:"content"     description:"公告内容"`            // 公告内容
	CreatedBy int64       `json:"created_by"  description:"创建者"`             // 创建者
	UpdatedBy int64       `json:"updated_by"  description:"更新者"`             // 更新者
	CreatedAt *gtime.Time `json:"created_at"  description:"创建时间"`            // 创建时间
	UpdatedAt *gtime.Time `json:"updated_at"  description:"更新时间"`            // 更新时间
	Users     []int64     `json:"users"        description:"接收用户"`           // 接收用户
	Remark    string      `json:"remark"       description:"备注"`             // 备注
}
