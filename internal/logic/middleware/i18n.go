// Package middleware
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package middleware

import (
	"devinggo/modules/system/pkg/i18n"
	"github.com/gogf/gf/v2/net/ghttp"
)

func (s *sMiddleware) I18n(r *ghttp.Request) {
	i18n.InitI18n(r.GetCtx())
	r.Middleware.Next()
}
