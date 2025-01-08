// Package hook
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package hook

import (
	"context"
	"devinggo/modules/system/pkg/contexts"
	"devinggo/modules/system/pkg/utils"
	"devinggo/modules/system/service"
	"github.com/gogf/gf/v2/net/ghttp"
)

func (s *sHook) apiAccessLog(r *ghttp.Request) {
	if service.Hook().IsIgnoredRequest(r) {
		return
	}

	var ctx = r.Context()
	if contexts.New().Get(ctx) == nil {
		return
	}

	utils.SafeGo(ctx, func(ctx context.Context) {
		service.SystemApiLog().Push(ctx)
	})
}
