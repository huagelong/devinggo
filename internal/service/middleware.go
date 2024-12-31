// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"github.com/gogf/gf/v2/net/ghttp"
)

type (
	IMiddleware interface {
		Cors(r *ghttp.Request)
		Ctx(r *ghttp.Request)
		I18n(r *ghttp.Request)
		NeverDoneCtx(r *ghttp.Request)
		// ResponseHandler custom response format.
		ResponseHandler(r *ghttp.Request)
		// TraceID use 'Trace-Id' from client request header by default.
		TraceID(r *ghttp.Request)
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
