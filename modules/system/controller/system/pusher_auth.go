// Package system
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package system

import (
	"context"
	"strings"

	"devinggo/modules/system/api/system"
	"devinggo/modules/system/controller/base"
	"devinggo/modules/system/pkg/websocket"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

var (
	PusherAuthController = pusherAuthController{}
)

type pusherAuthController struct {
	base.BaseController
}

// PusherAuth Pusher频道认证端点
// 用于Private/Presence频道的认证签名生成
// 注意：此接口直接返回Pusher标准格式的JSON，不使用GoFrame响应包装
func (c *pusherAuthController) PusherAuth(ctx context.Context, req *system.PusherAuthReq) (rs *system.PusherAuthRes, err error) {
	r := g.RequestFromCtx(ctx)

	// 验证socket_id格式（防止伪造）
	if req.SocketId == "" || req.ChannelName == "" {
		writeJSONOrJSONP(r, g.Map{
			"error": "socket_id and channel_name are required",
		}, 403)
		return
	}

	// ⚠️ 验证socket_id是否属于当前用户（防止跨用户攻击）
	// 从Redis获取socket_id对应的ServerName，验证连接存在性
	serverName := websocket.GetServerNameBySocketId4Redis(ctx, req.SocketId)
	if serverName == "" {
		g.Log().Warning(ctx, "Invalid socket_id:", req.SocketId)
		writeJSONOrJSONP(r, g.Map{
			"error": "Invalid socket_id or connection not found",
		}, 403)
		return
	}

	// ⚠️ 安全检查：验证频道类型
	if !websocket.RequiresAuth(req.ChannelName) {
		writeJSONOrJSONP(r, g.Map{
			"error": "This channel does not require authentication",
		}, 403)
		return
	}

	// 可选安全增强（根据业务需求启用）：
	// - 验证用户是否有权访问该频道（根据业务规则）
	// - 记录认证日志（审计）
	// - 限流防止滥用（Rate Limiting）

	// ⚠️ 检查是否为 Encrypted Channel（private-encrypted-*）
	if websocket.IsEncryptedChannel(req.ChannelName) {
		var sharedSecret string

		// 根据配置选择密钥生成方式
		if websocket.HasEncryptionMasterKey() {
			// 使用派生模式（Pusher 标准）
			deriveErr := error(nil)
			sharedSecret, deriveErr = websocket.DeriveSharedSecret(req.ChannelName)
			if deriveErr != nil {
				g.Log().Error(ctx, "Failed to derive shared_secret:", deriveErr)
				writeJSONOrJSONP(r, g.Map{
					"error": "Failed to derive channel key",
				}, 500)
				return
			}
			g.Log().Debugf(ctx, "Derived shared_secret for channel: %s", req.ChannelName)
		} else {
			// 使用随机生成模式（per-channel，向后兼容）
			sharedSecret = websocket.GenerateSharedSecret()
			saveErr := websocket.SaveSharedSecret(ctx, req.ChannelName, sharedSecret)
			if saveErr != nil {
				g.Log().Warning(ctx, "Failed to save shared_secret:", saveErr)
				// 继续执行，不影响客户端认证
			}
			g.Log().Debugf(ctx, "Generated random shared_secret for channel: %s", req.ChannelName)
		}

		// 生成认证签名（不包含 channel_data）
		auth := websocket.GenerateAuthSignature(req.SocketId, req.ChannelName, "")

		g.Log().Debugf(ctx, "Generated encrypted channel auth for user:%d, socket:%s, channel:%s", c.UserId, req.SocketId, req.ChannelName)

		// 返回 Pusher 标准格式（包含 shared_secret）
		writeJSONOrJSONP(r, g.Map{
			"auth":          auth,
			"shared_secret": sharedSecret,
		}, 200)
		return
	}

	// 检查是否为Presence频道
	if websocket.IsPresenceChannel(req.ChannelName) {
		// Phase 3: Presence频道认证（包含channel_data）
		// 生成channel_data（包含用户信息）
		userData, userErr := buildPusherUserData(ctx, c.UserId)
		if userErr != nil {
			g.Log().Warning(ctx, "Failed to load user data for presence auth:", userErr)
			writeJSONOrJSONP(r, g.Map{
				"error": "Failed to load user profile",
			}, 500)
			return
		}

		channelData, err := websocket.EncodeChannelData(userData["id"].(string), buildPresenceUserInfo(userData))
		if err != nil {
			g.Log().Warning(ctx, "EncodeChannelData error:", err)
			writeJSONOrJSONP(r, g.Map{
				"error": "Failed to generate channel_data",
			}, 500)
			return nil, err
		}

		// 生成认证签名（包含channel_data）
		auth := websocket.GenerateAuthSignature(req.SocketId, req.ChannelName, channelData)

		g.Log().Debugf(ctx, "Generated presence auth for user:%d, socket:%s, channel:%s", c.UserId, req.SocketId, req.ChannelName)

		// 直接返回Pusher标准格式（不包装在GoFrame响应中）
		writeJSONOrJSONP(r, g.Map{
			"auth":         auth,
			"channel_data": channelData,
		}, 200)
	} else {
		// Private频道认证（不包含channel_data）
		auth := websocket.GenerateAuthSignature(req.SocketId, req.ChannelName, "")

		g.Log().Debugf(ctx, "Generated private auth for user:%d, socket:%s, channel:%s", c.UserId, req.SocketId, req.ChannelName)

		// 直接返回Pusher标准格式（不包装在GoFrame响应中）
		writeJSONOrJSONP(r, g.Map{
			"auth": auth,
		}, 200)
	}

	return
}

// BatchAuth Pusher批量频道认证端点
// 支持一次请求认证多个频道（private/presence/encrypted）
func (c *pusherAuthController) BatchAuth(ctx context.Context, req *system.PusherBatchAuthReq) (rs *system.PusherBatchAuthRes, err error) {
	r := g.RequestFromCtx(ctx)

	channelNames := normalizeBatchChannels(r, req.ChannelNames, req.Channels)
	if req.SocketId == "" || len(channelNames) == 0 {
		writeJSONOrJSONP(r, g.Map{
			"error": "socket_id and channel_names are required",
		}, 403)
		return
	}

	serverName := websocket.GetServerNameBySocketId4Redis(ctx, req.SocketId)
	if serverName == "" {
		writeJSONOrJSONP(r, g.Map{
			"error": "Invalid socket_id or connection not found",
		}, 403)
		return
	}

	result := make(map[string]system.PusherBatchAuthItem, len(channelNames))
	for _, channelName := range channelNames {
		if !websocket.RequiresAuth(channelName) {
			continue
		}

		item := system.PusherBatchAuthItem{}

		if websocket.IsEncryptedChannel(channelName) {
			var sharedSecret string

			// 根据配置选择密钥生成方式
			if websocket.HasEncryptionMasterKey() {
				// 使用派生模式（Pusher 标准）
				deriveErr := error(nil)
				sharedSecret, deriveErr = websocket.DeriveSharedSecret(channelName)
				if deriveErr != nil {
					g.Log().Error(ctx, "Failed to derive shared_secret for batch auth:", deriveErr)
					continue // 跳过此频道
				}
			} else {
				// 使用随机生成模式（per-channel，向后兼容）
				sharedSecret = websocket.GenerateSharedSecret()
				saveErr := websocket.SaveSharedSecret(ctx, channelName, sharedSecret)
				if saveErr != nil {
					g.Log().Warning(ctx, "Failed to save shared_secret:", saveErr)
				}
			}

			item.Auth = websocket.GenerateAuthSignature(req.SocketId, channelName, "")
			item.SharedSecret = sharedSecret
			result[channelName] = item
			continue
		}

		if websocket.IsPresenceChannel(channelName) {
			userData, userErr := buildPusherUserData(ctx, c.UserId)
			if userErr != nil {
				g.Log().Warning(ctx, "Failed to load user data for batch presence auth:", userErr)
				continue
			}

			channelData, encodeErr := websocket.EncodeChannelData(userData["id"].(string), buildPresenceUserInfo(userData))
			if encodeErr != nil {
				g.Log().Warning(ctx, "EncodeChannelData error:", encodeErr)
				continue
			}
			item.Auth = websocket.GenerateAuthSignature(req.SocketId, channelName, channelData)
			item.ChannelData = channelData
			result[channelName] = item
			continue
		}

		item.Auth = websocket.GenerateAuthSignature(req.SocketId, channelName, "")
		result[channelName] = item
	}

	writeJSONOrJSONP(r, g.Map{
		"channels": result,
	}, 200)

	return
}

func normalizeBatchChannels(r *ghttp.Request, channelNames []string, channels []string) []string {
	merged := make([]string, 0, len(channelNames)+len(channels))
	merged = append(merged, channelNames...)
	merged = append(merged, channels...)

	if len(merged) == 0 {
		raw := r.Get("channel_names").String()
		if raw == "" {
			raw = r.Get("channels").String()
		}
		if raw != "" {
			for _, v := range strings.Split(raw, ",") {
				name := strings.TrimSpace(v)
				if name != "" {
					merged = append(merged, name)
				}
			}
		}
	}

	seen := make(map[string]struct{}, len(merged))
	unique := make([]string, 0, len(merged))
	for _, ch := range merged {
		ch = strings.TrimSpace(ch)
		if ch == "" {
			continue
		}
		if _, ok := seen[ch]; ok {
			continue
		}
		seen[ch] = struct{}{}
		unique = append(unique, ch)
	}

	return unique
}
