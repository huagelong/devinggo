// Package middleware
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package middleware

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

// ApiAuth API认证中间件
func (s *sMiddleware) ApiAuth(r *ghttp.Request) {
	g.Log().Debug(r.GetCtx(), "{{.moduleName}} ApiAuth middleware called")
	r.Middleware.Next()
}
