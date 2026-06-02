// Package request
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package request

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

// GetClientIp 获取客户端IP
func GetClientIp(ctx context.Context) string {
	ip := g.RequestFromCtx(ctx).GetClientIp()
	if g.IsEmpty(ip) {
		ip = "127.0.0.1"
	}
	return ip
}

// GetUserAgent 获取user-agent
func GetUserAgent(ctx context.Context) string {
	return ghttp.RequestFromCtx(ctx).Header.Get("User-Agent")
}

func GetHttpRequest(ctx context.Context) *ghttp.Request {
	r := ghttp.RequestFromCtx(ctx)
	if r == nil {
		g.Log().Warningf(ctx, "ctx not http request")
		return nil
	}
	return r
}
