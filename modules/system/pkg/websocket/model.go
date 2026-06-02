// Package websocket
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package websocket

import "sync"

// PusherRequest Pusher协议请求结构（客户端→服务器）
type PusherRequest struct {
	Event   string      `json:"event"`             // 事件名（如 "pusher:subscribe"）
	Data    interface{} `json:"data,omitempty"`    // JSON对象或字符串
	Channel string      `json:"channel,omitempty"` // 频道名（可选）
}

// PusherResponse Pusher协议响应结构（服务器→客户端）
// ⚠️ v8.3.0关键要求：data字段必须是JSON字符串
type PusherResponse struct {
	Event   string `json:"event"`             // 事件名
	Channel string `json:"channel,omitempty"` // 频道名（可选）
	Data    string `json:"data"`              // 必须是JSON字符串，不能是对象
}

// ConnectionEstablishedData 连接建立时的数据结构
type ConnectionEstablishedData struct {
	SocketID        string `json:"socket_id"`        // 连接唯一标识
	ActivityTimeout int    `json:"activity_timeout"` // 活动超时（秒，v8.3.0推荐120）
}

// ErrorData 错误消息数据结构
type ErrorData struct {
	Message string `json:"message"` // 错误消息
	Code    int    `json:"code"`    // 错误码
}

// SubscriptionErrorData 订阅错误数据结构（⚠️ v8.3.0要求）
type SubscriptionErrorData struct {
	Type   string `json:"type"`   // 错误类型（"AuthError" | "ChannelLimitReached"）
	Error  string `json:"error"`  // 错误描述
	Status int    `json:"status"` // HTTP状态码 (401/403/500)
}

// SubscribeRequestData 订阅请求的data字段结构
type SubscribeRequestData struct {
	Channel     string `json:"channel"`                // 频道名
	Auth        string `json:"auth,omitempty"`         // 认证签名（Private/Presence）
	ChannelData string `json:"channel_data,omitempty"` // 频道数据（Presence）
}

// SigninRequestData 用户登录请求数据
type SigninRequestData struct {
	Auth     string      `json:"auth"`      // 用户认证签名
	UserData interface{} `json:"user_data"` // 用户数据（对象或JSON字符串）
}

// SigninSuccessData 用户登录成功数据
type SigninSuccessData struct {
	UserData string `json:"user_data"`
}

// PresenceData Presence频道订阅成功时的数据结构
type PresenceData struct {
	Presence PresenceMemberList `json:"presence"`
}

// PresenceMemberList Presence成员列表结构
type PresenceMemberList struct {
	Count int                    `json:"count"`        // 成员数量
	Ids   []string               `json:"ids"`          // 成员ID列表
	Hash  map[string]interface{} `json:"hash"`         // 成员信息哈希表（⚠️ v8.3.0：只含user_info）
	Me    *PresenceMember        `json:"me,omitempty"` // 当前用户信息
}

// PresenceMember Presence成员结构
type PresenceMember struct {
	UserID   string                 `json:"user_id"`   // 用户ID
	UserInfo map[string]interface{} `json:"user_info"` // 用户信息
}

// MemberAddedData 成员加入事件数据
type MemberAddedData struct {
	UserID   string                 `json:"user_id"`   // 用户ID
	UserInfo map[string]interface{} `json:"user_info"` // 用户信息
}

// MemberRemovedData 成员离开事件数据
type MemberRemovedData struct {
	UserID string `json:"user_id"` // 用户ID
}

// Redis消息传递结构
type TopicWResponse struct {
	Topic           string          `json:"topic"`
	ExcludeSocketID string          `json:"exclude_socket_id,omitempty"`
	PusherResponse  *PusherResponse `json:"pusherResponse"`
}

type ClientIdWResponse struct {
	SocketID       string          `json:"socketId"`
	PusherResponse *PusherResponse `json:"pusherResponse"`
}

type BroadcastWResponse struct {
	Broadcast      string          `json:"broadcast"`
	PusherResponse *PusherResponse `json:"pusherResponse"`
}

// TopicTerminateConnection 终止连接消息结构（用于跨服务器）
type TopicTerminateConnection struct {
	SocketID string `json:"socketId"`
}

// Pusher错误码常量（v8.3.0完整版）
const (
	CodeNormalClosure        = 4000 // 正常关闭
	CodeGoingAway            = 4001 // 服务器下线
	CodeMaxConnections       = 4004 // 连接数超限
	CodePathNotFound         = 4005 // 路径错误
	CodeOverCapacity         = 4008 // 服务过载
	CodeUnauthorized         = 4009 // 未授权
	CodeAppDisabled          = 4100 // 应用已禁用
	CodePongTimeout          = 4200 // 心跳超时
	CodeClosedByClient       = 4201 // 客户端主动关闭
	CodeClientEventForbidden = 4301 // 客户端事件错误
)

// Pusher系统事件常量
const (
	EventConnectionEstablished = "pusher:connection_established"
	EventError                 = "pusher:error"
	EventPing                  = "pusher:ping"
	EventPong                  = "pusher:pong"
	EventSignin                = "pusher:signin"
	EventSigninSuccess         = "pusher:signin_success"
	EventSubscribe             = "pusher:subscribe"
	EventSubscriptionSucceeded = "pusher:subscription_succeeded"
	EventSubscriptionError     = "pusher:subscription_error"
	EventUnsubscribe           = "pusher:unsubscribe"
	EventMemberAdded           = "pusher:member_added"
	EventMemberRemoved         = "pusher:member_removed"
)

// 客户端事件前缀
const (
	ClientEventPrefix = "client-"
)

// pusherResponsePool 用于复用PusherResponse对象（性能优化）
var pusherResponsePool = sync.Pool{
	New: func() interface{} {
		return &PusherResponse{}
	},
}

// AcquirePusherResponse 从池中获取PusherResponse对象
func AcquirePusherResponse() *PusherResponse {
	return pusherResponsePool.Get().(*PusherResponse)
}

// ReleasePusherResponse 将PusherResponse对象归还池中
func ReleasePusherResponse(resp *PusherResponse) {
	// 重置字段避免脏数据
	resp.Event = ""
	resp.Channel = ""
	resp.Data = ""
	pusherResponsePool.Put(resp)
}
