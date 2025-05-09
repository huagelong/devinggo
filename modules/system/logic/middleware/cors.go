// Package middleware
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package middleware

import (
	"devinggo/modules/system/pkg/utils/config"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/text/gstr"
)

func (s *sMiddleware) Cors(r *ghttp.Request) {
	ctx := r.Context()
	corsOptions := r.Response.DefaultCORSOptions()
	allowDomainConfig := config.GetConfigString(ctx, "settings.cors.allowDomain", "*")
	allowDomain := gstr.Split(allowDomainConfig, ",")
	allowHeaders := config.GetConfigString(ctx, "settings.cors.allowHeaders", "")
	if !g.IsEmpty(allowHeaders) {
		allowHeaders = "," + allowHeaders
	}
	g.Log().Debug(ctx, "cors allow domain:", allowDomain, "allow headers:", allowHeaders)
	if !g.IsEmpty(allowDomain) {
		corsOptions.AllowDomain = allowDomain
	}
	corsOptions.AllowHeaders = corsOptions.AllowHeaders + allowHeaders
	corsOptions.AllowOrigin = "*"
	r.Response.CORS(corsOptions)
	r.Middleware.Next()
}
