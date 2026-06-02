// Package system
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package system

import (
	"devinggo/modules/system/model"

	"github.com/gogf/gf/v2/frame/g"
)

// PusherAuthReq Pusher频道认证请求
type PusherAuthReq struct {
	g.Meta `path:"/pusher/auth" method:"post,get" tags:"WebSocket" summary:"Pusher频道认证" x-exceptAuth:"true"`
	model.AuthorHeader
	SocketId    string `json:"socket_id" dc:"WebSocket连接的socket_id" v:"required"`
	ChannelName string `json:"channel_name" dc:"要订阅的频道名称" v:"required"`
}

// PusherAuthRes Pusher频道认证响应
type PusherAuthRes struct {
	g.Meta      `mime:"application/json"`
	Auth        string `json:"auth" dc:"认证签名，格式：{app_key}:{signature}"`
	ChannelData string `json:"channel_data,omitempty" dc:"Presence频道的用户数据（可选），格式：{user_id, user_info}"`
}

// PusherBatchAuthReq Pusher批量频道认证请求
// 用于一次请求认证多个频道，减少前端认证请求次数
type PusherBatchAuthReq struct {
	g.Meta `path:"/pusher/auth/batch" method:"post,get" tags:"WebSocket" summary:"Pusher批量频道认证" x-exceptAuth:"true"`
	model.AuthorHeader
	SocketId     string   `json:"socket_id" dc:"WebSocket连接的socket_id" v:"required"`
	ChannelNames []string `json:"channel_names" dc:"频道名称列表（推荐）"`
	Channels     []string `json:"channels" dc:"频道名称列表（兼容字段）"`
}

// PusherBatchAuthItem 单个频道认证结果
type PusherBatchAuthItem struct {
	Auth         string `json:"auth"`
	ChannelData  string `json:"channel_data,omitempty"`
	SharedSecret string `json:"shared_secret,omitempty"`
}

// PusherBatchAuthRes Pusher批量频道认证响应
type PusherBatchAuthRes struct {
	g.Meta   `mime:"application/json"`
	Channels map[string]PusherBatchAuthItem `json:"channels"`
}
