// Package websocket
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package websocket

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"devinggo/modules/system/pkg/websocket/glob"

	"github.com/gogf/gf/v2/os/gtime"
)

// SigninController Pusher用户认证控制器
// 客户端发送 pusher:signin 后验证签名，成功返回 pusher:signin_success
func SigninController(ctx context.Context, client *Client, req *PusherRequest) {
	dataBytes, err := json.Marshal(req.Data)
	if err != nil {
		glob.WithWsLog().Warning(ctx, "SigninController marshal error:", err)
		client.SendError("Invalid signin data", CodeUnauthorized)
		return
	}

	var payload struct {
		Auth     string          `json:"auth"`
		UserData json.RawMessage `json:"user_data"`
	}
	if err = json.Unmarshal(dataBytes, &payload); err != nil {
		glob.WithWsLog().Warning(ctx, "SigninController unmarshal error:", err)
		client.SendError("Invalid signin payload", CodeUnauthorized)
		return
	}

	if payload.Auth == "" || len(payload.UserData) == 0 {
		client.SendError("Missing auth or user_data", CodeUnauthorized)
		return
	}

	userDataMap, userDataJSON, err := parseSigninUserData(payload.UserData)
	if err != nil {
		glob.WithWsLog().Warning(ctx, "SigninController parse user_data error:", err)
		client.SendError("Invalid user_data", CodeUnauthorized)
		return
	}

	userID, ok := userDataMap["id"].(string)
	if !ok || userID == "" {
		client.SendError("user_data.id is required and must be string", CodeUnauthorized)
		return
	}

	if err = ValidateUserAuthSignature(client.SocketID, payload.Auth, userDataJSON); err != nil {
		glob.WithWsLog().Warning(ctx, "SigninController validate signature failed:", err)
		client.SendError("Invalid user auth signature", CodeUnauthorized)
		return
	}

	// 标记用户级连接状态
	client.UserID = userID
	client.UserInfo = userDataMap

	// 保存 user_id → socket_id 映射到 Redis（用于 Send to User API）
	if err = SaveUserIdSocketIdMapping(ctx, userID, client.SocketID); err != nil {
		glob.WithWsLog().Warning(ctx, "SigninController save user_id mapping failed:", err)
		// 不阻塞登录流程，继续执行
	}

	// 返回用户级登录成功事件
	// ⚠️ pusher-js 期望 user_data 为 JSON 字符串，而不是对象
	if err = client.SendPusherEvent(EventSigninSuccess, "", SigninSuccessData{UserData: userDataJSON}); err != nil {
		glob.WithWsLog().Warning(ctx, "SigninController send signin_success failed:", err)
		return
	}

	glob.WithWsLog().Infof(ctx, "Client signed in: socket_id=%s, user_id=%s", client.SocketID, userID)
}

func parseSigninUserData(raw json.RawMessage) (map[string]interface{}, string, error) {
	trimmed := strings.TrimSpace(string(raw))
	if trimmed == "" {
		return nil, "", fmt.Errorf("empty user_data")
	}

	if strings.HasPrefix(trimmed, "\"") {
		var userDataJSON string
		if err := json.Unmarshal(raw, &userDataJSON); err != nil {
			return nil, "", err
		}
		var userDataMap map[string]interface{}
		if err := json.Unmarshal([]byte(userDataJSON), &userDataMap); err != nil {
			return nil, "", err
		}
		return userDataMap, userDataJSON, nil
	}

	var userDataMap map[string]interface{}
	if err := json.Unmarshal(raw, &userDataMap); err != nil {
		return nil, "", err
	}
	return userDataMap, trimmed, nil
}

// SubscribeController Pusher订阅控制器
func SubscribeController(ctx context.Context, client *Client, req *PusherRequest) {
	// 解析data字段
	var subData SubscribeRequestData
	dataBytes, err := json.Marshal(req.Data)
	if err != nil {
		glob.WithWsLog().Warning(ctx, "SubscribeController marshal error:", err)
		client.SendError("Invalid data", CodeNormalClosure)
		return
	}

	err = json.Unmarshal(dataBytes, &subData)
	if err != nil {
		glob.WithWsLog().Warning(ctx, "SubscribeController unmarshal error:", err)
		client.SendError("Invalid subscribe data", CodeNormalClosure)
		return
	}

	channel := subData.Channel
	if channel == "" {
		client.SendError("Missing channel", CodeNormalClosure)
		return
	}

	// 获取频道类型
	channelType := GetChannelType(channel)

	// Phase 2: Private/Encrypted Channel认证
	// ⚠️ Encrypted channels 使用与 Private channels 相同的认证流程
	if channelType == ChannelTypePrivate || channelType == ChannelTypeEncrypted {
		auth := subData.Auth
		if auth == "" {
			glob.WithWsLog().Warning(ctx, "Private/Encrypted channel missing auth:", channel)
			client.SendSubscriptionError(channel, "AuthError", "Auth signature required for private/encrypted channel", CodeUnauthorized)
			return
		}

		// 验证认证签名
		err = ValidateChannelAuth(client.SocketID, channel, auth, "")
		if err != nil {
			glob.WithWsLog().Warning(ctx, "Private/Encrypted channel auth failed:", err)
			client.SendSubscriptionError(channel, "AuthError", "Invalid auth signature", CodeUnauthorized)
			return
		}

		// 验证成功：加入频道
		if !client.HasChannel(channel) {
			client.AddChannel(channel)
		}

		err = JoinChannel4Redis(ctx, client.SocketID, channel)
		if err != nil {
			glob.WithWsLog().Warning(ctx, "JoinChannel4Redis error:", err)
			client.SendSubscriptionError(channel, "ServerError", err.Error(), 500)
			return
		}

		// 发送订阅成功事件
		client.SendPusherEvent(EventSubscriptionSucceeded, channel, map[string]interface{}{})

		if channelType == ChannelTypeEncrypted {
			glob.WithWsLog().Debugf(ctx, "Client %s subscribed to encrypted channel %s", client.SocketID, channel)
		} else {
			glob.WithWsLog().Debugf(ctx, "Client %s subscribed to private channel %s", client.SocketID, channel)
		}
		return
	}

	// Phase 3: Presence Channel认证
	if channelType == ChannelTypePresence {
		auth := subData.Auth
		channelData := subData.ChannelData

		if auth == "" || channelData == "" {
			glob.WithWsLog().Warning(ctx, "Presence channel missing auth or channel_data:", channel)
			client.SendSubscriptionError(channel, "AuthError", "Auth signature and channel_data required for presence channel", CodeUnauthorized)
			return
		}

		// 验证认证签名（包含channel_data）
		err = ValidateChannelAuth(client.SocketID, channel, auth, channelData)
		if err != nil {
			glob.WithWsLog().Warning(ctx, "Presence channel auth failed:", err)
			client.SendSubscriptionError(channel, "AuthError", "Invalid auth signature", CodeUnauthorized)
			return
		}

		// 解析channel_data获取user_id和user_info
		member, parseErr := ParseChannelData(channelData)
		if parseErr != nil || member.UserID == "" {
			glob.WithWsLog().Warning(ctx, "Invalid channel_data:", parseErr)
			client.SendSubscriptionError(channel, "AuthError", "Invalid channel_data format", CodeUnauthorized)
			return
		}

		// 保存用户信息到Client
		client.UserID = member.UserID
		client.UserInfo = member.UserInfo

		// 添加成员到Redis
		err = AddPresenceMember4Redis(ctx, channel, member.UserID, member.UserInfo)
		if err != nil {
			glob.WithWsLog().Warning(ctx, "AddPresenceMember4Redis error:", err)
			client.SendSubscriptionError(channel, "ServerError", err.Error(), 500)
			return
		}

		// 加入频道
		if !client.HasChannel(channel) {
			client.AddChannel(channel)
		}

		err = JoinChannel4Redis(ctx, client.SocketID, channel)
		if err != nil {
			glob.WithWsLog().Warning(ctx, "JoinChannel4Redis error:", err)
			client.SendSubscriptionError(channel, "ServerError", err.Error(), 500)
			return
		}

		// 获取所有成员列表（使用缓存优化性能）
		members, _, memErr := GetPresenceCache().GetMembers(ctx, channel)
		if memErr != nil {
			members = make(map[string]map[string]interface{})
		}

		// 🔍 调试：打印成员列表
		glob.WithWsLog().Debugf(ctx, "🔍 [Presence订阅] channel=%s, members数量=%d", channel, len(members))
		for userID, userInfo := range members {
			glob.WithWsLog().Debugf(ctx, "🔍 [Presence成员] userID=%s, userInfo=%+v", userID, userInfo)
		}

		// 发送订阅成功事件（包含完整成员列表）
		// ⚠️ 传入当前用户ID以构造 me 字段
		// ⚠️ Pusher.js 8.3.0 需要扁平化结构，直接发送 PresenceMemberList，不包装在 presence 字段中
		presenceData := FormatPresenceData(members, member.UserID)

		// 🔍 调试：验证 presenceData.Presence 可以被正确序列化
		testJSON, testErr := json.Marshal(presenceData.Presence) // 使用 .Presence 字段（扁平化）
		if testErr != nil {
			glob.WithWsLog().Errorf(ctx, "❌ [Presence数据] JSON序列化测试失败: %v", testErr)
		} else {
			glob.WithWsLog().Debugf(ctx, "🔍 [Presence数据] JSON序列化成功（扁平化），长度=%d, 内容=%s", len(testJSON), string(testJSON))
		}

		// 直接发送 PresenceMemberList（扁平化），而不是包装在 PresenceData 中
		client.SendPusherEvent(EventSubscriptionSucceeded, channel, presenceData.Presence)

		// 向频道内其他成员广播member_added事件
		memberAddedData := MemberAddedData{
			UserID:   member.UserID,
			UserInfo: member.UserInfo,
		}
		BroadcastToChannel(ctx, channel, EventMemberAdded, memberAddedData, client.SocketID)

		// 使缓存失效（因为新增了成员）
		GetPresenceCache().InvalidateChannel(channel)

		glob.WithWsLog().Debugf(ctx, "Client %s subscribed to presence channel %s as user %s", client.SocketID, channel, member.UserID)
		return
	}

	// Public Channel订阅
	if !client.HasChannel(channel) {
		client.AddChannel(channel)
	}

	err = JoinChannel4Redis(ctx, client.SocketID, channel)
	if err != nil {
		glob.WithWsLog().Warning(ctx, "JoinChannel4Redis error:", err)
		client.SendSubscriptionError(channel, "ServerError", err.Error(), 500)
		return
	}

	// 发送订阅成功事件（⚠️ Public Channel的data为空对象字符串 "{}"）
	client.SendPusherEvent(EventSubscriptionSucceeded, channel, map[string]interface{}{})
	glob.WithWsLog().Debugf(ctx, "Client %s subscribed to channel %s", client.SocketID, channel)
}

// UnsubscribeController Pusher退订控制器
func UnsubscribeController(ctx context.Context, client *Client, req *PusherRequest) {
	// 解析data字段
	var subData SubscribeRequestData
	dataBytes, err := json.Marshal(req.Data)
	if err != nil {
		glob.WithWsLog().Warning(ctx, "UnsubscribeController marshal error:", err)
		return
	}

	err = json.Unmarshal(dataBytes, &subData)
	if err != nil {
		glob.WithWsLog().Warning(ctx, "UnsubscribeController unmarshal error:", err)
		return
	}

	channel := subData.Channel
	if channel == "" {
		return
	}

	if client.HasChannel(channel) {
		client.RemoveChannel(channel)
	}

	LeaveChannel4Redis(ctx, channel, client.SocketID)

	// Phase 3: Presence Channel退订（触发Grace Period）
	channelType := GetChannelType(channel)
	if channelType == ChannelTypePresence && client.UserID != "" {
		// ⚠️ 标记断线，启动Grace Period（30秒）
		MarkPresenceDisconnect4Redis(ctx, client.SocketID)

		// 30秒后移除成员（如果未重连）
		// 注意：实际延迟清理在client.close()中实现
		glob.WithWsLog().Debugf(ctx, "Client %s marked for presence grace period on channel %s", client.SocketID, channel)
	}

	// ⚠️ Pusher规范：unsubscribe不发送响应事件
	glob.WithWsLog().Debugf(ctx, "Client %s unsubscribed from channel %s", client.SocketID, channel)
}

// PingController Pusher心跳控制器
func PingController(ctx context.Context, client *Client, req *PusherRequest) {
	currentTime := int64(gtime.Now().Unix())
	client.Heartbeat(currentTime)
	UpdateSocketIdHeartbeatTime4Redis(ctx, client.SocketID, currentTime)

	// 发送pong响应（⚠️ data为空对象字符串 "{}"）
	client.SendPusherEvent(EventPong, "", map[string]interface{}{})
}

// ClientEventController 客户端事件控制器（Phase 4: Client Events）
func ClientEventController(ctx context.Context, client *Client, req *PusherRequest) {
	// ⚠️ v8.3.0要求：验证事件名必须以client-开头
	if !strings.HasPrefix(req.Event, "client-") {
		glob.WithWsLog().Warning(ctx, "Invalid client event name:", req.Event)
		client.SendError("Client event must start with 'client-'", CodeClientEventForbidden)
		return
	}

	// ⚠️ v8.3.0要求：验证事件名长度（最大200字节）
	if len(req.Event) > 200 {
		glob.WithWsLog().Warning(ctx, "Client event name too long:", req.Event)
		client.SendError("Event name exceeds 200 bytes", CodeClientEventForbidden)
		return
	}

	// 验证channel字段
	if req.Channel == "" {
		glob.WithWsLog().Warning(ctx, "Client event missing channel")
		client.SendError("Channel required for client events", CodeClientEventForbidden)
		return
	}

	// ⚠️ v8.3.0要求：仅允许Private/Presence频道使用Client Events
	// ⚠️ Encrypted Channels 不支持 Client Events（Pusher.js 官方限制）
	channelType := GetChannelType(req.Channel)
	if channelType == ChannelTypePublic {
		glob.WithWsLog().Warning(ctx, "Client events not allowed on public channels:", req.Channel)
		client.SendError("Client events only allowed on private and presence channels", CodeClientEventForbidden)
		return
	}

	// ⚠️ 加密频道不支持 Client Events（Pusher Protocol 规范）
	if channelType == ChannelTypeEncrypted {
		glob.WithWsLog().Warning(ctx, "Client events not supported on encrypted channels:", req.Channel)
		client.SendError("Client events are not supported for encrypted channels", CodeClientEventForbidden)
		return
	}

	// 验证客户端是否已订阅该频道
	if !client.HasChannel(req.Channel) {
		glob.WithWsLog().Warning(ctx, "Client not subscribed to channel:", req.Channel)
		client.SendError("Must subscribe to channel before sending client events", CodeClientEventForbidden)
		return
	}

	// ⚠️ v8.3.0要求：速率限制（10条/秒）
	rateLimiter := GetRateLimiter()
	if !rateLimiter.AllowClientEvent(client.SocketID) {
		glob.WithWsLog().Warning(ctx, "Rate limit exceeded for client:", client.SocketID)
		client.SendError("Rate limit exceeded (max 10 events/sec)", CodeClientEventForbidden)
		return
	}

	// 转发给频道内其他客户端（不包括发送者）
	// ⚠️ Presence Channel 需要包含发送者的 user_id
	if channelType == ChannelTypePresence && client.UserID != "" {
		// 为 Presence Channel 的 Client Events 添加发送者信息
		BroadcastToChannelWithSender(ctx, req.Channel, req.Event, req.Data, client.SocketID, client.UserID)
	} else {
		BroadcastToChannel(ctx, req.Channel, req.Event, req.Data, client.SocketID)
	}

	glob.WithWsLog().Debugf(ctx, "Client event forwarded: socket=%s, channel=%s, event=%s",
		client.SocketID, req.Channel, req.Event)
}

// BroadcastToChannel 向频道内除了指定客户端以外的所有成员广播消息
func BroadcastToChannel(ctx context.Context, channel, event string, data interface{}, excludeSocketID string) {
	BroadcastToChannelWithSender(ctx, channel, event, data, excludeSocketID, "")
}

// BroadcastToChannelWithSender 向频道内除了指定客户端以外的所有成员广播消息（可选包含发送者信息）
// ⚠️ 用于 Presence Channel 的 Client Events，需要包含 user_id
func BroadcastToChannelWithSender(ctx context.Context, channel, event string, data interface{}, excludeSocketID, senderUserID string) {
	// 获取频道内所有socket_id
	socketIds := GetAllSocketIdByChannel4Redis(ctx, channel)

	for _, socketId := range socketIds {
		// 跳过指定的客户端
		if socketId == excludeSocketID {
			continue
		}

		// 获取客户端并发送消息
		targetClient := clientManager.GetClientBySocketID(socketId)
		if targetClient != nil {
			// ⚠️ Presence Channel Client Events: 包装数据以包含发送者 user_id
			if senderUserID != "" && strings.HasPrefix(event, "client-") {
				// 创建包含 metadata 的包装对象
				// Pusher.js 会自动解析这个格式并将 user_id 作为第二个参数传递给回调
				wrappedData := map[string]interface{}{
					"data":    data,
					"user_id": senderUserID, // Pusher.js 需要的 metadata
				}
				targetClient.SendPusherEvent(event, channel, wrappedData)
			} else {
				targetClient.SendPusherEvent(event, channel, data)
			}
		}
	}
}

// BroadcastMemberRemoved 广播member_removed事件
func BroadcastMemberRemoved(ctx context.Context, channel, userID string) {
	memberRemovedData := MemberRemovedData{
		UserID: userID,
	}

	// 向频道内所有成员广播
	socketIds := GetAllSocketIdByChannel4Redis(ctx, channel)
	for _, socketId := range socketIds {
		targetClient := clientManager.GetClientBySocketID(socketId)
		if targetClient != nil {
			targetClient.SendPusherEvent(EventMemberRemoved, channel, memberRemovedData)
		}
	}

	// 使Presence缓存失效（成员已移除）
	GetPresenceCache().InvalidateChannel(channel)
}
