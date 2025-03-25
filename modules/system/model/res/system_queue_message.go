// Package res
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package res

import (
	"devinggo/modules/system/model"
	"github.com/gogf/gf/v2/os/gtime"
)

type SystemQueueMessage struct {
	Id          int64            `json:"id"                    description:"主键"` // 主键
	ContentType string           `json:"content_type" description:"内容类型"`        // 内容类型
	Title       string           `json:"title"              description:"消息标题"`  // 消息标题
	SendBy      int64            `json:"send_by"           description:"发送人"`    // 发送人
	Content     string           `json:"content"          description:"消息内容"`    // 消息内容
	CreatedBy   int64            `json:"created_by"     description:"创建者"`       // 创建者
	UpdatedBy   int64            `json:"updated_by"     description:"更新者"`       // 更新者
	CreatedAt   *gtime.Time      `json:"created_at"     description:"创建时间"`      // 创建时间
	UpdatedAt   *gtime.Time      `json:"updated_at"     description:"更新时间"`      // 更新时间
	Remark      string           `json:"remark"            description:"备注"`     // 备注
	SendUser    model.UserRelate `json:"send_user"      description:"发送人信息"`     // 发送人信息
}

type MessageReceiveUser struct {
	Username      string `json:"username"  description:"用户名"`
	Nickname      string `json:"nickname"  description:"昵称"`
	ReadStatusInt int    `json:"read_status_int"  description:"阅读状态,1未读，2已读"`
	ReadStatus    string `json:"read_status"   description:"阅读状态,已读，未读"`
}
