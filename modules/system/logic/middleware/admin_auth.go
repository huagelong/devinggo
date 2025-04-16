// Package middleware
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package middleware

import (
	"devinggo/modules/system/codes"
	"devinggo/modules/system/pkg/contexts"
	"devinggo/modules/system/pkg/response"
	"devinggo/modules/system/service"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

// AdminAuth 后台鉴权中间件
func (s *sMiddleware) AdminAuth(r *ghttp.Request) {

	ctx := r.Context()
	// 不需要验证登录的路由地址
	if !s.IsExceptLogin(ctx) {
		// 检查登录
		if g.IsEmpty(contexts.New().GetUser(ctx)) || contexts.New().GetUser(ctx).Id == 0 {
			response.JsonError(r, codes.CodeNotLogged, "未登录或登录状态已过期，需要重新登录")
			return
		}
	}
	// 不需要验证权限的路由地址
	if !s.IsExceptAuth(ctx) {
		// 验证路由访问权限
		if !service.SystemRole().Verify(r) {
			response.JsonError(r, codes.CodeForbidden, "没有权限访问该资源")
			return
		}
	}

	r.Middleware.Next()
}
