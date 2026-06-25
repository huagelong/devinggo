// Package system
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package system

import (
	"context"

	"devinggo/modules/system/api/system"
	"devinggo/modules/system/controller/base"
	"devinggo/modules/system/pkg/websocket"
	"devinggo/modules/system/service"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

var (
	PusherUserAuthController = cPusherUserAuthController{}
)

type cPusherUserAuthController struct {
	base.BaseController
}

// UserAuth Pusher User Authentication - 用户身份认证
//
// 用途：
// - 识别和认证用户身份（区别于频道授权）
// - 用户级别的事件订阅和触发
// - 跨频道的用户状态追踪
//
// 使用场景：
// - 用户在线状态管理
// - 用户级别的通知
// - 多设备登录检测
// - 用户行为追踪
//
// 客户端用法：
//
//	const pusher = new Pusher('app-key', {
//	    userAuthentication: {
//	        endpoint: '/system/pusher/user-auth'
//	    }
//	});
//	pusher.signin();
//
// 文档：https://pusher.com/docs/channels/server_api/authenticating-users/
func (c *cPusherUserAuthController) UserAuth(ctx context.Context, req *system.PusherUserAuthReq) (res *system.PusherUserAuthRes, err error) {
	r := g.RequestFromCtx(ctx)

	// 验证用户是否已登录
	if c.UserId == 0 {
		writeJSONOrJSONP(r, g.Map{
			"error": "Unauthorized - User not logged in",
		}, 401)
		return
	}

	g.Log().Debugf(ctx, "Pusher User Authentication: user_id=%d, socket_id=%s", c.UserId, req.SocketId)

	userData, err := buildPusherUserData(ctx, c.UserId)
	if err != nil {
		g.Log().Error(ctx, "Failed to load user data for pusher user auth:", err)
		writeJSONOrJSONP(r, g.Map{
			"error": "Failed to load user profile",
		}, 500)
		return
	}

	// 生成用户认证签名
	auth, err := websocket.GenerateUserAuthSignature(req.SocketId, userData)
	if err != nil {
		g.Log().Error(ctx, "Failed to generate user auth signature:", err)
		writeJSONOrJSONP(r, g.Map{
			"error": "Failed to generate authentication signature",
		}, 500)
		return
	}

	g.Log().Infof(ctx, "User authenticated successfully: user_id=%d, socket_id=%s", c.UserId, req.SocketId)

	// 返回认证响应
	res = &system.PusherUserAuthRes{
		Auth:     auth,
		UserData: userData,
	}

	writeJSONOrJSONP(r, res, 200)

	return
}

func buildPusherUserData(ctx context.Context, userId int64) (map[string]interface{}, error) {
	user, err := service.SystemUser().GetInfoById(ctx, userId)
	if err != nil {
		return nil, err
	}
	if g.IsEmpty(user) {
		return nil, gerror.New("user not found")
	}

	name := user.Nickname
	if name == "" {
		name = user.Username
	}

	userData := map[string]interface{}{
		"id":        gconv.String(user.Id),
		"name":      name,
		"username":  user.Username,
		"nickname":  user.Nickname,
		"email":     user.Email,
		"avatar":    user.Avatar,
		"user_type": user.UserType,
		"status":    user.Status,
	}

	return userData, nil
}

func buildPresenceUserInfo(userData map[string]interface{}) map[string]interface{} {
	userInfo := make(map[string]interface{}, len(userData))
	for k, v := range userData {
		if k == "id" {
			continue
		}
		userInfo[k] = v
	}
	return userInfo
}
