// Package middleware
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package middleware

import (
	"devinggo/modules/system/model"
	"devinggo/modules/system/myerror"
	"devinggo/modules/system/pkg/utils"
	websocket2 "devinggo/modules/system/pkg/websocket"
	"devinggo/modules/system/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
)

// ws鉴权中间件
func (s *sMiddleware) WsAuth(r *ghttp.Request) {
	ctx := r.GetCtx()
	sessionId, err := s.parseSessionId(r)
	g.Log().Debug(ctx, "sessionId:", sessionId)

	// ⚠️ Pusher协议支持：允许无token连接（用于Public频道）
	// Private/Presence频道通过HTTP认证端点进行认证
	if utils.IsError(err) && g.IsEmpty(sessionId) {
		// 设置空sessionId，允许连接
		r.SetCtxVar(websocket2.SESSION_ID_KEY, "")
		r.Middleware.Next()
		return
	} else {
		r.SetCtxVar(websocket2.SESSION_ID_KEY, sessionId)
	}

	r.Middleware.Next()
}

func (s *sMiddleware) parseSessionId(r *ghttp.Request) (sessionId string, err error) {
	ctx := r.GetCtx()
	sessionIdTmp := r.GetQuery("sessionId")
	token := r.GetQuery("token")

	//权限检查
	if g.IsEmpty(token) {
		return "", myerror.MissingParameter(ctx, "sessionId或token缺失")
	}

	claims, err := service.Token().ParseToken(ctx, token.String())
	if utils.IsError(err) {
		return "", err
	}
	data := claims.Data
	if g.IsEmpty(data) {
		return "", myerror.ValidationFailed(ctx, "claims为空")
	}

	if g.IsEmpty(sessionIdTmp) {
		var user *model.Identity
		data := claims.Data
		err = gconv.Scan(data, &user)
		if utils.IsError(err) {
			return "", err
		}
		if g.IsEmpty(user) {
			return "", myerror.ValidationFailed(ctx, "sessionId缺失")
		} else {
			return gconv.String(user.Id), nil
		}

	} else {
		return sessionIdTmp.String(), nil
	}
}
