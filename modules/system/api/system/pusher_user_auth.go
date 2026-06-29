// Package system
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package system

import (
	"github.com/gogf/gf/v2/frame/g"
)

// PusherUserAuthReq Pusher User Authentication 请求
// 用于识别和认证用户身份（不同于频道授权）
// 端点：/system/pusher/user-auth
// 文档：https://pusher.com/docs/channels/server_api/authenticating-users/
type PusherUserAuthReq struct {
	g.Meta   `path:"/pusher/user-auth" method:"post,get" tags:"Pusher" summary:"Pusher User Authentication"`
	SocketId string `json:"socket_id" v:"required" dc:"客户端 socket_id"`
}

// PusherUserAuthRes Pusher User Authentication 响应
type PusherUserAuthRes struct {
	Auth     string                 `json:"auth" dc:"认证签名：app_key:hmac_signature"`
	UserData map[string]interface{} `json:"user_data" dc:"用户数据（必须包含 id 字段）"`
}
