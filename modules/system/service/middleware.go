// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"

	"github.com/gogf/gf/v2/net/ghttp"
)

type (
	IMiddleware interface {
		// AdminAuth 后台鉴权中间件
		AdminAuth(r *ghttp.Request)
		Cors(r *ghttp.Request)
		Ctx(r *ghttp.Request)
		Header(r *ghttp.Request)
		I18n(r *ghttp.Request)
		// IsExceptAuth 是否是不需要验证权限的路由地址
		IsExceptAuth(ctx context.Context) bool
		// IsExceptLogin 是否是不需要登录的路由地址
		IsExceptLogin(ctx context.Context) bool
		NeverDoneCtx(r *ghttp.Request)
		// ResponseHandler custom response format.
		ResponseHandler(r *ghttp.Request)
		// TraceID use 'Trace-Id' from client request header by default.
		TraceID(r *ghttp.Request)
		// ws鉴权中间件
		WsAuth(r *ghttp.Request)
	}
)

var (
	localMiddleware IMiddleware
)

func Middleware() IMiddleware {
	if localMiddleware == nil {
		panic("implement not found for interface IMiddleware, forgot register?")
	}
	return localMiddleware
}

func RegisterMiddleware(i IMiddleware) {
	localMiddleware = i
}
