// Package websocket
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package websocket

import (
	"devinggo/modules/system/pkg/websocket/glob"
	"context"
	"github.com/gogf/gf/v2/util/gconv"
)

const (
	Connected        = "connected"
	Close            = "close"
	Subscribe        = "subscribe"
	Unsubscribe      = "unsubscribe"
	Ping             = "ping"
	PingAll          = "pingAll"
	Pong             = "pong"
	BroadcastMessage = "broadcastMsg"
	IdMessage        = "idMsg"
	Publish          = "publish"
)

// ProcessData
func ProcessData(ctx context.Context, client *Client, message []byte) {
	defer func() {
		if r := recover(); r != nil {
			glob.WithWsLog().Warning(ctx, "ProcessData error:", r)
		}
	}()

	request := &Request{}
	err := gconv.Struct(message, request)
	if err != nil {
		glob.WithWsLog().Warning(ctx, "ProcessData gconv error:", err)
		return
	}
	glob.WithWsLog().Debug(ctx, "ws request：", request)

	switch request.Event {
	case Subscribe:
		SubscribeController(ctx, client, request)
		break
	case Unsubscribe:
		UnsubscribeController(ctx, client, request)
		break
	case Ping:
		PingController(ctx, client, request)
		break
	case Pong:
		PongController(ctx, client, request)
		break
	case IdMessage:
		IdMessageController(ctx, client, request)
		break
	case Publish:
		PublishController(ctx, client, request)
		break
	case BroadcastMessage:
		BroadcastMessageController(ctx, client, request)
		break
	case Close:
		CloseController(ctx, client, request)
		break
	}
}
