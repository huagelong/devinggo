// Package system
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package system

import (
	"github.com/gogf/gf/v2/frame/g"
)

// PusherWebhookReq Pusher Webhook 验证请求
// Pusher 会在特定事件发生时（如频道被占用/清空、成员加入/离开等）向配置的 webhook URL 发送 POST 请求
// 文档：https://pusher.com/docs/channels/server_api/webhooks
type PusherWebhookReq struct {
	g.Meta `path:"/apps/:app_id/webhooks" method:"post" tags:"WebSocket" summary:"Pusher Webhook 验证" x-exceptAuth:"true"`

	// Path参数
	AppId string `json:"app_id" in:"path" dc:"应用ID" v:"required"`

	// Header参数（用于签名验证）
	XPusherKey       string `json:"X-Pusher-Key" in:"header" dc:"应用Key"`
	XPusherSignature string `json:"X-Pusher-Signature" in:"header" dc:"HMAC-SHA256签名"`

	// Body参数会自动从请求体解析
	TimeMs int64                `json:"time_ms" dc:"时间戳（毫秒）" v:"required"`
	Events []PusherWebhookEvent `json:"events" dc:"事件列表" v:"required"`
}

// PusherWebhookEvent Webhook 事件项
type PusherWebhookEvent struct {
	Name     string `json:"name" dc:"事件名称（如：channel_occupied, channel_vacated 等）"`
	Channel  string `json:"channel" dc:"频道名称"`
	Event    string `json:"event,omitempty" dc:"客户端事件名称（仅用于 client_event 类型）"`
	Data     string `json:"data,omitempty" dc:"事件数据"`
	SocketID string `json:"socket_id,omitempty" dc:"Socket ID"`
	UserID   string `json:"user_id,omitempty" dc:"用户ID（仅用于 member_added/removed 事件）"`
}

// PusherWebhookRes Pusher Webhook 响应
type PusherWebhookRes struct {
	g.Meta  `mime:"application/json" status:"200"`
	Success bool   `json:"success" dc:"是否成功"`
	Message string `json:"message,omitempty" dc:"消息"`
}
