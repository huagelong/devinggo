// Package cms
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE
package api

import (
	"context"
	"devinggo/modules/api/controller"
	_ "devinggo/modules/api/logic"
	"devinggo/modules/api/service"
	"devinggo/modules/system/pkg/modules"
	"github.com/gogf/gf/v2/net/ghttp"
)

type apiModule struct {
	Name   string
	Server *ghttp.Server
}

func init() {
	module := &apiModule{}
	module.Name = "api"
	modules.Register(module)
}

func (m *apiModule) Start(ctx context.Context, s *ghttp.Server) error {
	m.Server = s
	s.BindHookHandler("/api/*", ghttp.HookBeforeServe, service.Hook().BeforeServe)
	s.BindHookHandler("/api/*", ghttp.HookAfterOutput, service.Hook().AfterOutput)
	s.Group("/api", func(group *ghttp.RouterGroup) {
		group.Bind(
			controller.TokenController,
			controller.TestController,
		).Middleware(service.Middleware().ApiAuth)
	})
	return nil
}

func (m *apiModule) Stop(ctx context.Context) error {
	return nil
}

func (m *apiModule) GetName() string {
	return m.Name
}
