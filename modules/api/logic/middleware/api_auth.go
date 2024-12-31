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

func (s *sMiddleware) ApiAuth(r *ghttp.Request) {
	ctx := r.Context()
	// 不需要验证登录的路由地址
	if !s.isExceptLogin(ctx) {
		// 检查登录
		if g.IsEmpty(contexts.New().GetUser(ctx)) {
			response.JsonError(r, codes.CodeNotLogged)
			return
		}
	}

	// 不需要验证权限的路由地址
	if !s.isExceptAuth(ctx) {
		// 验证路由访问权限
		check, err := service.SystemApp().Verify(r)
		if err != nil {
			response.JsonError(r, codes.ApiNotAuthorized, err.Error())
			return
		}
		if !check {
			response.JsonError(r, codes.ApiNotAuthorized, "接口鉴权失败")
			return
		}
	}

	r.Middleware.Next()
}
