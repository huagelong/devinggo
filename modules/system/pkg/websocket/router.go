// Package websocket
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package websocket

import (
	"context"
	"encoding/json"
	"strings"

	"devinggo/modules/system/pkg/websocket/glob"
)

// ProcessData 处理客户端消息（Pusher协议）
func ProcessData(ctx context.Context, client *Client, message []byte) {
	defer func() {
		if r := recover(); r != nil {
			glob.WithWsLog().Warning(ctx, "ProcessData error:", r)
		}
	}()

	request := &PusherRequest{}
	err := json.Unmarshal(message, request)
	if err != nil {
		glob.WithWsLog().Warning(ctx, "ProcessData json.Unmarshal error:", err)
		client.SendError("Invalid JSON", CodeNormalClosure)
		return
	}

	// glob.WithWsLog().Debug(ctx, "ws request：", request)

	// ⚠️ 优先检查 client- 事件前缀（Client Events）
	if strings.HasPrefix(request.Event, ClientEventPrefix) {
		ClientEventController(ctx, client, request)
		return
	}

	// 路由Pusher系统事件
	switch request.Event {
	case EventSignin:
		SigninController(ctx, client, request)
	case EventSubscribe:
		SubscribeController(ctx, client, request)
	case EventUnsubscribe:
		UnsubscribeController(ctx, client, request)
	case EventPing:
		PingController(ctx, client, request)
	default:
		glob.WithWsLog().Warning(ctx, "Unknown event:", request.Event)
		client.SendError("Unknown event: "+request.Event, CodeNormalClosure)
	}
}
