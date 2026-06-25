// Package system
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package system

import (
	"github.com/gogf/gf/v2/frame/g"
)

// PusherChannelUsersReq 获取 Presence 频道用户列表请求
type PusherChannelUsersReq struct {
	g.Meta `path:"/apps/:app_id/channels/:channel_name/users" method:"get" tags:"WebSocket" summary:"Pusher Channel Users API - 获取频道用户列表" x-exceptAuth:"true"`

	// Path参数
	AppId       string `json:"app_id" in:"path" dc:"应用ID" v:"required"`
	ChannelName string `json:"channel_name" in:"path" dc:"频道名称（仅限 presence-* 频道）" v:"required"`

	// Query参数（用于签名验证）
	AuthKey       string `json:"auth_key" in:"query" dc:"应用Key" v:"required"`
	AuthTimestamp int64  `json:"auth_timestamp" in:"query" dc:"时间戳" v:"required"`
	AuthVersion   string `json:"auth_version" in:"query" dc:"认证版本" v:"required"`
	AuthSignature string `json:"auth_signature" in:"query" dc:"HMAC-SHA256签名" v:"required"`
}

// PusherChannelUsersRes 获取频道用户列表响应
type PusherChannelUsersRes struct {
	g.Meta `mime:"application/json" status:"200"`
	Users  []PusherChannelUser `json:"users" dc:"用户列表"`
}

// PusherChannelUser 频道用户信息
type PusherChannelUser struct {
	ID string `json:"id" dc:"用户ID"`
}

// PusherChannelInfoReq 获取单个频道信息请求
type PusherChannelInfoReq struct {
	g.Meta `path:"/apps/:app_id/channels/:channel_name" method:"get" tags:"WebSocket" summary:"Pusher Channel Info API - 获取频道信息" x-exceptAuth:"true"`

	// Path参数
	AppId       string `json:"app_id" in:"path" dc:"应用ID" v:"required"`
	ChannelName string `json:"channel_name" in:"path" dc:"频道名称" v:"required"`

	// Query参数（用于签名验证）
	AuthKey       string `json:"auth_key" in:"query" dc:"应用Key" v:"required"`
	AuthTimestamp int64  `json:"auth_timestamp" in:"query" dc:"时间戳" v:"required"`
	AuthVersion   string `json:"auth_version" in:"query" dc:"认证版本" v:"required"`
	AuthSignature string `json:"auth_signature" in:"query" dc:"HMAC-SHA256签名" v:"required"`

	// 可选参数
	Info string `json:"info" in:"query" dc:"要返回的信息，逗号分隔：user_count（presence频道）, subscription_count（需开启）"`
}

// PusherChannelInfoRes 获取频道信息响应
type PusherChannelInfoRes struct {
	g.Meta            `mime:"application/json" status:"200"`
	Name              string `json:"name,omitempty" dc:"频道名称"`
	Occupied          bool   `json:"occupied" dc:"频道是否被占用（有订阅者）"`
	UserCount         int    `json:"user_count,omitempty" dc:"用户数（仅 presence 频道）"`
	SubscriptionCount int    `json:"subscription_count,omitempty" dc:"订阅数（需在配置中启用）"`
}

// PusherChannelsListReq 获取频道列表请求
type PusherChannelsListReq struct {
	g.Meta `path:"/apps/:app_id/channels" method:"get" tags:"WebSocket" summary:"Pusher Channels List API - 获取频道列表" x-exceptAuth:"true"`

	// Path参数
	AppId string `json:"app_id" in:"path" dc:"应用ID" v:"required"`

	// Query参数（用于签名验证）
	AuthKey       string `json:"auth_key" in:"query" dc:"应用Key" v:"required"`
	AuthTimestamp int64  `json:"auth_timestamp" in:"query" dc:"时间戳" v:"required"`
	AuthVersion   string `json:"auth_version" in:"query" dc:"认证版本" v:"required"`
	AuthSignature string `json:"auth_signature" in:"query" dc:"HMAC-SHA256签名" v:"required"`

	// 可选参数
	FilterByPrefix string `json:"filter_by_prefix" in:"query" dc:"按前缀过滤频道（如：presence-, private-）"`
	Info           string `json:"info" in:"query" dc:"要返回的信息：user_count（仅 presence 频道）"`
}

// PusherChannelsListRes 获取频道列表响应
type PusherChannelsListRes struct {
	g.Meta   `mime:"application/json" status:"200"`
	Channels map[string]PusherChannelListItem `json:"channels" dc:"频道列表"`
}

// PusherChannelListItem 频道列表项
type PusherChannelListItem struct {
	UserCount int `json:"user_count,omitempty" dc:"用户数（仅 presence 频道，且需请求 info=user_count）"`
}
