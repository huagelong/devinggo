// Package cms
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE
package cms

import (
	"context"
	"devinggo/modules/cms/controller"
	_ "devinggo/modules/cms/logic"
	"devinggo/modules/cms/service"
	"devinggo/modules/system/pkg/modules"
	systemService "devinggo/modules/system/service"
	"github.com/gogf/gf/v2/net/ghttp"
)

type CmsModule struct {
	Name   string
	Server *ghttp.Server
}

func init() {
	module := &CmsModule{}
	module.Name = "cms"
	modules.Register(module)
}

func (m *CmsModule) Start(ctx context.Context, s *ghttp.Server) error {
	m.Server = s
	s.BindHookHandler("/cms/*", ghttp.HookBeforeServe, service.Hook().BeforeServe)
	s.BindHookHandler("/cms/*", ghttp.HookAfterOutput, service.Hook().AfterOutput)
	s.Group("/cms", func(group *ghttp.RouterGroup) {
		group.Group("/api", func(group *ghttp.RouterGroup) {
			group.Bind(
			controller.TestController,
			).Middleware(service.Middleware().ApiAuth)
		})
		group.Group("/", func(group *ghttp.RouterGroup) {
			group.Bind(
				controller.TestController,
				controller.CmsTopicController,
			).Middleware(systemService.Middleware().AdminAuth)
		})
	})
	return nil
}

func (m *CmsModule) Stop(ctx context.Context) error {
	return nil
}

func (m *CmsModule) GetName() string {
	return m.Name
}