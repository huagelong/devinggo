// Package system
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE
package system

import (
	"context"
	"devinggo/modules/system/pkg/modules"
	swebsocket "devinggo/modules/system/pkg/websocket"
	"devinggo/modules/system/router/system"
	"devinggo/modules/system/router/websocket"
	"devinggo/modules/system/service"
	_ "devinggo/modules/system/worker"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/genv"
	"github.com/gogf/gf/v2/util/gconv"
)

type systemModule struct {
	Name   string
	Server *ghttp.Server
}

func init() {
	module := &systemModule{}
	module.Name = "system"
	modules.Register(module)
}

func (m *systemModule) Start(ctx context.Context, s *ghttp.Server) error {
	m.Server = s
	serverName := getServerName(ctx)
	g.Log().Debug(ctx, "serverName:", serverName)
	swebsocket.StartWebSocket(ctx, serverName)
	s.BindHookHandler("/system/*", ghttp.HookBeforeServe, service.Hook().BeforeServe)
	s.BindHookHandler("/system/*", ghttp.HookAfterOutput, service.Hook().AfterOutput)
	s.Group("/", func(group *ghttp.RouterGroup) {
		system.BindController(group)
		websocket.BindController(group)
	})
	return nil
}

func (m *systemModule) Stop(ctx context.Context) error {
	return nil
}

func (m *systemModule) GetName() string {
	return m.Name
}

func getServerName(ctx context.Context) string {
	serverName := genv.GetWithCmd("server.name")
	if g.IsEmpty(serverName) {
		if v, err := g.Cfg().GetWithEnv(ctx, "server.name"); err == nil {
			return gconv.String(v)
		}
	}
	return gconv.String(serverName)
}
