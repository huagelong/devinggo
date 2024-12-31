// Package hook
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package hook

import (
	baseService "devinggo/internal/service"
	"devinggo/modules/system/pkg/contexts"
	"devinggo/modules/system/pkg/utils"
	"devinggo/modules/system/service"
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
)

// system 访问日志
func (s *sHook) AccessLog(r *ghttp.Request) {
	if baseService.Hook().IsIgnoredRequest(r) {
		return
	}

	var ctx = r.Context()
	if contexts.New().Get(ctx) == nil {
		return
	}

	if contexts.New().GetExceptAccessLog(ctx) {
		return
	}

	utils.SafeGo(ctx, func(ctx context.Context) {
		service.SystemOperLog().Push(ctx)
	})
}
