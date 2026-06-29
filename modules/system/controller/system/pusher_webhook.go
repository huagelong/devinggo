// Package system
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package system

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"

	"devinggo/modules/system/api/system"

	"github.com/gogf/gf/v2/frame/g"
)

var (
	PusherWebhook = cPusherWebhook{}
)

type cPusherWebhook struct{}

// Webhook Pusher Webhook 验证和处理
// 用于接收和验证来自 Pusher 的 webhook 事件
// 事件类型包括：
// - channel_occupied: 频道被占用（第一个订阅者）
// - channel_vacated: 频道被清空（最后一个订阅者离开）
// - member_added: 成员加入 presence 频道
// - member_removed: 成员离开 presence 频道
// - client_event: 客户端事件（如果启用了事件记录）
//
// 文档：https://pusher.com/docs/channels/server_api/webhooks
func (c *cPusherWebhook) Webhook(ctx context.Context, req *system.PusherWebhookReq) (res *system.PusherWebhookRes, err error) {
	r := g.RequestFromCtx(ctx)

	// 1. 验证应用配置
	config, err := getAppConfig(ctx)
	if err != nil {
		g.Log().Warning(ctx, "Webhook: Failed to get app config:", err)
		return &system.PusherWebhookRes{
			Success: false,
			Message: "Invalid app configuration",
		}, nil
	}

	if req.AppId != config.AppID {
		g.Log().Warning(ctx, "Webhook: Invalid app_id:", req.AppId)
		return &system.PusherWebhookRes{
			Success: false,
			Message: "Invalid app_id",
		}, nil
	}

	// 2. 验证 X-Pusher-Key
	pusherKey := req.XPusherKey
	if pusherKey == "" {
		pusherKey = r.Header.Get("X-Pusher-Key")
	}
	if pusherKey != config.Key {
		g.Log().Warning(ctx, "Webhook: Invalid X-Pusher-Key:", pusherKey)
		return &system.PusherWebhookRes{
			Success: false,
			Message: "Invalid X-Pusher-Key",
		}, nil
	}

	// 3. 验证签名
	pusherSignature := req.XPusherSignature
	if pusherSignature == "" {
		pusherSignature = r.Header.Get("X-Pusher-Signature")
	}

	bodyBytes := r.GetBody()
	if !verifyWebhookSignature(bodyBytes, pusherSignature, config.Secret) {
		g.Log().Warning(ctx, "Webhook: Invalid signature")
		return &system.PusherWebhookRes{
			Success: false,
			Message: "Invalid signature",
		}, nil
	}

	// 4. 处理 webhook 事件
	g.Log().Infof(ctx, "Webhook received: time_ms=%d, events_count=%d", req.TimeMs, len(req.Events))

	for i, event := range req.Events {
		g.Log().Debugf(ctx, "Webhook event[%d]: name=%s, channel=%s, event=%s, socket_id=%s, user_id=%s",
			i, event.Name, event.Channel, event.Event, event.SocketID, event.UserID)

		// 根据事件类型进行处理
		switch event.Name {
		case "channel_occupied":
			// 频道被占用（第一个订阅者加入）
			g.Log().Infof(ctx, "Channel occupied: %s", event.Channel)
			// 扩展点：可在此添加业务逻辑（如记录日志、统计在线用户数等）

		case "channel_vacated":
			// 频道被清空（最后一个订阅者离开）
			g.Log().Infof(ctx, "Channel vacated: %s", event.Channel)
			// 扩展点：可在此添加业务逻辑（如清理频道资源等）

		case "member_added":
			// 成员加入 presence 频道
			g.Log().Infof(ctx, "Member added to %s: user_id=%s", event.Channel, event.UserID)
			// 扩展点：可在此添加业务逻辑（如更新在线状态等）

		case "member_removed":
			// 成员离开 presence 频道
			g.Log().Infof(ctx, "Member removed from %s: user_id=%s", event.Channel, event.UserID)
			// 扩展点：可在此添加业务逻辑（如更新在线状态等）

		case "client_event":
			// 客户端触发的事件
			g.Log().Infof(ctx, "Client event on %s: event=%s, socket_id=%s", event.Channel, event.Event, event.SocketID)
			// 扩展点：可在此添加业务逻辑（如处理自定义事件等）

		default:
			g.Log().Warningf(ctx, "Unknown webhook event type: %s", event.Name)
		}
	}

	// 5. 返回成功响应
	res = &system.PusherWebhookRes{
		Success: true,
		Message: fmt.Sprintf("Processed %d events", len(req.Events)),
	}
	return
}

// verifyWebhookSignature 验证 Webhook 签名
// Pusher 使用 HMAC-SHA256 对请求体进行签名
// 签名算法：hex(HMAC-SHA256(secret, body))
func verifyWebhookSignature(body []byte, providedSignature string, secret string) bool {
	// 计算 HMAC-SHA256
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write(body)
	expectedSignature := hex.EncodeToString(mac.Sum(nil))

	g.Log().Debugf(context.Background(), "Webhook signature verification:")
	g.Log().Debugf(context.Background(), "  Expected: %s", expectedSignature)
	g.Log().Debugf(context.Background(), "  Provided: %s", providedSignature)

	// 比对签名
	return hmac.Equal([]byte(expectedSignature), []byte(providedSignature))
}
