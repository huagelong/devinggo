// Package middleware
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package middleware

import (
	"github.com/gogf/gf/v2/net/ghttp"
)

func (s *sMiddleware) NeverDoneCtx(r *ghttp.Request) {
	r.SetCtx(r.GetNeverDoneCtx())
	r.Middleware.Next()
}
