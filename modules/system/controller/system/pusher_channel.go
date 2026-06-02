// Package system
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package system

import (
	"context"
	"fmt"
	"strings"
	"time"

	"devinggo/modules/system/api/system"
	"devinggo/modules/system/pkg/websocket"

	"github.com/gogf/gf/v2/frame/g"
)

var (
	PusherChannel = cPusherChannel{}
)

type cPusherChannel struct{}

// GetChannelUsers 获取 Presence 频道的用户列表
// 此 API 仅适用于 presence-* 频道
// GET /apps/{app_id}/channels/{channel_name}/users
func (c *cPusherChannel) GetChannelUsers(ctx context.Context, req *system.PusherChannelUsersReq) (res *system.PusherChannelUsersRes, err error) {
	r := g.RequestFromCtx(ctx)

	// 1. 验证应用配置
	config, err := getAppConfig(ctx)
	if err != nil {
		return nil, err
	}

	if req.AppId != config.AppID {
		return nil, invalidAppIdError(r)
	}

	if req.AuthKey != config.Key {
		return nil, invalidAppKeyError(r)
	}

	// 2. 验证时间戳
	now := time.Now().Unix()
	if abs(now-req.AuthTimestamp) > 600 {
		return nil, timestampExpiredError(r)
	}

	// 3. 验证签名（GET 请求，body 为空）
	if !verifySignature(req.AuthKey, req.AuthTimestamp, req.AuthVersion, "", req.AuthSignature, config.Secret, "GET", fmt.Sprintf("/apps/%s/channels/%s/users", req.AppId, req.ChannelName), []byte{}) {
		return nil, signatureInvalidError(r)
	}

	// 4. 验证频道类型（必须是 presence 频道）
	if !websocket.IsPresenceChannel(req.ChannelName) {
		r.Response.Status = 400
		r.Response.WriteJson(g.Map{
			"error": "This API only works for presence channels",
		})
		r.ExitAll()
		return nil, nil
	}

	// 5. 从 Redis 获取频道成员列表
	members, err := websocket.GetPresenceMembers4Redis(ctx, req.ChannelName)
	if err != nil {
		g.Log().Warning(ctx, "GetChannelUsers: Failed to get presence members:", err)
		r.Response.Status = 500
		r.Response.WriteJson(g.Map{
			"error": "Failed to retrieve channel members",
		})
		r.ExitAll()
		return nil, nil
	}

	// 6. 构建用户列表（只返回 user_id，不包含 user_info）
	users := make([]system.PusherChannelUser, 0, len(members))
	for userID := range members {
		users = append(users, system.PusherChannelUser{
			ID: userID,
		})
	}

	g.Log().Debugf(ctx, "GetChannelUsers: channel=%s, users_count=%d", req.ChannelName, len(users))

	// 7. 返回成功响应
	res = &system.PusherChannelUsersRes{
		Users: users,
	}
	return
}

// GetChannelInfo 获取频道信息
// 返回频道的状态信息，包括是否占用、用户数（presence 频道）、订阅数等
// GET /apps/{app_id}/channels/{channel_name}
func (c *cPusherChannel) GetChannelInfo(ctx context.Context, req *system.PusherChannelInfoReq) (res *system.PusherChannelInfoRes, err error) {
	r := g.RequestFromCtx(ctx)

	// 1. 验证应用配置
	config, err := getAppConfig(ctx)
	if err != nil {
		return nil, err
	}

	if req.AppId != config.AppID {
		return nil, invalidAppIdError(r)
	}

	if req.AuthKey != config.Key {
		return nil, invalidAppKeyError(r)
	}

	// 2. 验证时间戳
	now := time.Now().Unix()
	if abs(now-req.AuthTimestamp) > 600 {
		return nil, timestampExpiredError(r)
	}

	// 3. 构建查询字符串（包含 info 参数）
	path := fmt.Sprintf("/apps/%s/channels/%s", req.AppId, req.ChannelName)

	// 4. 验证签名（GET 请求，body 为空）
	if !verifySignature(req.AuthKey, req.AuthTimestamp, req.AuthVersion, "", req.AuthSignature, config.Secret, "GET", path, []byte{}) {
		return nil, signatureInvalidError(r)
	}

	// 5. 查询频道订阅者数量
	socketIds := websocket.GetAllSocketIdByChannel4Redis(ctx, req.ChannelName)
	subscriptionCount := len(socketIds)
	occupied := subscriptionCount > 0

	g.Log().Debugf(ctx, "GetChannelInfo: channel=%s, occupied=%v, subscription_count=%d", req.ChannelName, occupied, subscriptionCount)

	// 6. 构建响应
	res = &system.PusherChannelInfoRes{
		Occupied: occupied,
	}

	// 7. 根据 info 参数返回额外信息
	if req.Info != "" {
		infoParts := splitInfo(req.Info)

		// 返回频道名称（如果请求了任何信息）
		if len(infoParts) > 0 {
			res.Name = req.ChannelName
		}

		for _, info := range infoParts {
			switch info {
			case "user_count":
				// 只有 presence 频道才返回 user_count
				if websocket.IsPresenceChannel(req.ChannelName) {
					members, err := websocket.GetPresenceMembers4Redis(ctx, req.ChannelName)
					if err != nil {
						g.Log().Warning(ctx, "GetChannelInfo: Failed to get presence members:", err)
					} else {
						res.UserCount = len(members)
					}
				}
			case "subscription_count":
				// 返回订阅数
				res.SubscriptionCount = subscriptionCount
			}
		}
	}

	return
}

// splitInfo 分割 info 参数（按逗号分割）
func splitInfo(info string) []string {
	if info == "" {
		return []string{}
	}
	result := make([]string, 0)
	for _, part := range strings.Split(info, ",") {
		trimmed := strings.TrimSpace(part)
		if trimmed != "" {
			result = append(result, trimmed)
		}
	}
	return result
}

// GetChannelsList 获取频道列表
// 返回应用中所有活跃的频道列表，支持按前缀过滤
// GET /apps/{app_id}/channels
func (c *cPusherChannel) GetChannelsList(ctx context.Context, req *system.PusherChannelsListReq) (res *system.PusherChannelsListRes, err error) {
	r := g.RequestFromCtx(ctx)

	// 1. 验证应用配置
	config, err := getAppConfig(ctx)
	if err != nil {
		return nil, err
	}

	if req.AppId != config.AppID {
		return nil, invalidAppIdError(r)
	}

	if req.AuthKey != config.Key {
		return nil, invalidAppKeyError(r)
	}

	// 2. 验证时间戳
	now := time.Now().Unix()
	if abs(now-req.AuthTimestamp) > 600 {
		return nil, timestampExpiredError(r)
	}

	// 3. 验证签名（GET 请求，body 为空）
	path := fmt.Sprintf("/apps/%s/channels", req.AppId)
	if !verifySignature(req.AuthKey, req.AuthTimestamp, req.AuthVersion, "", req.AuthSignature, config.Secret, "GET", path, []byte{}) {
		return nil, signatureInvalidError(r)
	}

	// 4. 获取所有活跃频道
	allChannels := websocket.GetAllChannels(ctx)

	// 5. 按前缀过滤
	filteredChannels := make([]string, 0)
	if req.FilterByPrefix != "" {
		for _, channel := range allChannels {
			if len(channel) >= len(req.FilterByPrefix) && channel[:len(req.FilterByPrefix)] == req.FilterByPrefix {
				filteredChannels = append(filteredChannels, channel)
			}
		}
	} else {
		filteredChannels = allChannels
	}

	g.Log().Debugf(ctx, "GetChannelsList: total=%d, filtered=%d, prefix=%s", len(allChannels), len(filteredChannels), req.FilterByPrefix)

	// 6. 构建响应
	channels := make(map[string]system.PusherChannelListItem)

	// 7. 根据 info 参数返回额外信息
	includeUserCount := false
	if req.Info != "" {
		infoParts := splitInfo(req.Info)
		for _, info := range infoParts {
			if info == "user_count" {
				includeUserCount = true
				break
			}
		}
	}

	for _, channel := range filteredChannels {
		item := system.PusherChannelListItem{}

		// 如果请求了 user_count 且是 presence 频道，则返回用户数
		if includeUserCount && websocket.IsPresenceChannel(channel) {
			members, err := websocket.GetPresenceMembers4Redis(ctx, channel)
			if err != nil {
				g.Log().Warning(ctx, "GetChannelsList: Failed to get presence members for", channel, err)
			} else {
				item.UserCount = len(members)
			}
		}

		channels[channel] = item
	}

	res = &system.PusherChannelsListRes{
		Channels: channels,
	}

	return
}
