// Package middleware
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package middleware

import (
	"devinggo/modules/system/pkg/utils"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/net/gtrace"
)

const (
	clientTraceIDHeader = "Trace-Id"
)

// TraceID use 'Trace-Id' from client request header by default.
func (s *sMiddleware) TraceID(r *ghttp.Request) {
	traceID := r.GetHeader(clientTraceIDHeader)
	if traceID != "" {
		newCtx, err := gtrace.WithUUID(r.Context(), traceID)
		if utils.IsError(err) {
			g.Log().Errorf(r.Context(), "设置TraceID失败: %v", err)
		} else {
			r.SetCtx(newCtx)
		}
	}

	r.Middleware.Next()
}
