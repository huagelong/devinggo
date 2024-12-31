// Package hook
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package hook

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gmode"
)

func (s *sHook) header(r *ghttp.Request) {
	if gmode.IsDevelop() {
		r.Response.Header().Set("Cache-Control", "no-store")
	} else {
		r.Response.Header().Set("Cache-Control", "max-age=86400")
	}
}