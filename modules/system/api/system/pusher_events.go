// Package system
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package system

import (
	"github.com/gogf/gf/v2/frame/g"
)

// PusherEventsReq Pusher HTTP Events API请求（服务端推送事件）
type PusherEventsReq struct {
	g.Meta `path:"/apps/:app_id/events" method:"post" tags:"WebSocket" summary:"Pusher Events API - 推送事件到频道" x-exceptAuth:"true"`

	// Path参数
	AppId string `json:"app_id" in:"path" dc:"应用ID" v:"required"`

	// Query参数（用于签名验证）
	AuthKey       string `json:"auth_key" in:"query" dc:"应用Key" v:"required"`
	AuthTimestamp int64  `json:"auth_timestamp" in:"query" dc:"时间戳" v:"required"`
	AuthVersion   string `json:"auth_version" in:"query" dc:"认证版本" v:"required"`
	BodyMd5       string `json:"body_md5" in:"query" dc:"请求体MD5" v:"required"`
	AuthSignature string `json:"auth_signature" in:"query" dc:"HMAC-SHA256签名" v:"required"`

	// Body参数
	Name     string   `json:"name" dc:"事件名称" v:"required|max-length:200"`
	Channels []string `json:"channels" dc:"频道列表" v:"required|length:1,100"`
	Data     string   `json:"data" dc:"事件数据（JSON字符串，最大10KB）" v:"required|max-length:10240"`
	SocketId string   `json:"socket_id" dc:"排除的socket_id（可选）"`
	Info     string   `json:"info" dc:"要返回的频道信息（可选），如：user_count"`
}

// PusherEventsRes Pusher HTTP Events API响应
type PusherEventsRes struct {
	g.Meta   `mime:"application/json" status:"200"`
	Channels map[string]PusherTriggerChannelAttributes `json:"channels,omitempty" dc:"频道信息（当请求 info 参数时返回）"`
}

// PusherTriggerChannelAttributes 触发事件后返回的频道属性
type PusherTriggerChannelAttributes struct {
	UserCount int `json:"user_count,omitempty" dc:"用户数（仅 presence 频道）"`
}

// PusherBatchEventsReq Pusher HTTP Batch Events API请求（批量推送）
type PusherBatchEventsReq struct {
	g.Meta `path:"/apps/:app_id/batch_events" method:"post" tags:"WebSocket" summary:"Pusher Batch Events API - 批量推送事件" x-exceptAuth:"true"`

	// Path参数
	AppId string `json:"app_id" in:"path" dc:"应用ID" v:"required"`

	// Query参数（用于签名验证）
	AuthKey       string `json:"auth_key" in:"query" dc:"应用Key" v:"required"`
	AuthTimestamp int64  `json:"auth_timestamp" in:"query" dc:"时间戳" v:"required"`
	AuthVersion   string `json:"auth_version" in:"query" dc:"认证版本" v:"required"`
	BodyMd5       string `json:"body_md5" in:"query" dc:"请求体MD5" v:"required"`
	AuthSignature string `json:"auth_signature" in:"query" dc:"HMAC-SHA256签名" v:"required"`

	// Body参数
	Batch []PusherBatchEvent `json:"batch" dc:"事件列表" v:"required"`
}

// PusherBatchEvent 批量事件项
type PusherBatchEvent struct {
	Name     string `json:"name" dc:"事件名称" v:"required|max-length:200"`
	Channel  string `json:"channel" dc:"频道名称" v:"required|max-length:200"`
	Data     string `json:"data" dc:"事件数据（JSON字符串，最大10KB）" v:"required|max-length:10240"`
	SocketId string `json:"socket_id" dc:"排除的socket_id（可选）"`
	Info     string `json:"info" dc:"要返回的频道信息（可选），如：user_count"`
}

// PusherBatchEventsRes Pusher HTTP Batch Events API响应
type PusherBatchEventsRes struct {
	g.Meta `mime:"application/json" status:"200"`
	Batch  []PusherBatchEventResult `json:"batch,omitempty" dc:"批量事件结果（当请求 info 参数时返回）"`
}

// PusherBatchEventResult 批量事件结果项
type PusherBatchEventResult struct {
	UserCount *int `json:"user_count,omitempty" dc:"用户数（仅 presence 频道且请求了 info）"`
}

// PusherSendToUserReq Pusher HTTP Send to User API请求（向特定用户发送事件）
type PusherSendToUserReq struct {
	g.Meta `path:"/apps/:app_id/users/:user_id/events" method:"post" tags:"WebSocket" summary:"Pusher Send to User API - 向特定用户发送事件" x-exceptAuth:"true"`

	// Path参数
	AppId  string `json:"app_id" in:"path" dc:"应用ID" v:"required"`
	UserId string `json:"user_id" in:"path" dc:"目标用户ID" v:"required"`

	// Query参数（用于签名验证）
	AuthKey       string `json:"auth_key" in:"query" dc:"应用Key" v:"required"`
	AuthTimestamp int64  `json:"auth_timestamp" in:"query" dc:"时间戳" v:"required"`
	AuthVersion   string `json:"auth_version" in:"query" dc:"认证版本" v:"required"`
	BodyMd5       string `json:"body_md5" in:"query" dc:"请求体MD5" v:"required"`
	AuthSignature string `json:"auth_signature" in:"query" dc:"HMAC-SHA256签名" v:"required"`

	// Body参数
	Name string `json:"name" dc:"事件名称" v:"required|max-length:200"`
	Data string `json:"data" dc:"事件数据（JSON字符串，最大10KB）" v:"required|max-length:10240"`
}

// PusherSendToUserRes Pusher HTTP Send to User API响应
type PusherSendToUserRes struct {
	g.Meta `mime:"application/json" status:"200"`
}

// PusherTerminateConnectionsReq Pusher HTTP Terminate Connections API请求（终止用户的所有连接）
type PusherTerminateConnectionsReq struct {
	g.Meta `path:"/apps/:app_id/users/:user_id/terminate_connections" method:"post" tags:"WebSocket" summary:"Pusher Terminate Connections API - 终止用户的所有连接" x-exceptAuth:"true"`

	// Path参数
	AppId  string `json:"app_id" in:"path" dc:"应用ID" v:"required"`
	UserId string `json:"user_id" in:"path" dc:"目标用户ID" v:"required"`

	// Query参数（用于签名验证）
	AuthKey       string `json:"auth_key" in:"query" dc:"应用Key" v:"required"`
	AuthTimestamp int64  `json:"auth_timestamp" in:"query" dc:"时间戳" v:"required"`
	AuthVersion   string `json:"auth_version" in:"query" dc:"认证版本" v:"required"`
	BodyMd5       string `json:"body_md5" in:"query" dc:"请求体MD5" v:"required"`
	AuthSignature string `json:"auth_signature" in:"query" dc:"HMAC-SHA256签名" v:"required"`

	// Body参数（可选）
	SocketId string `json:"socket_id" dc:"要排除的socket_id（可选）"`
}

// PusherTerminateConnectionsRes Pusher HTTP Terminate Connections API响应
type PusherTerminateConnectionsRes struct {
	g.Meta `mime:"application/json" status:"200"`
}
