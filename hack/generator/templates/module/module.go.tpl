// Package {{.moduleName}}
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package {{.moduleName}}

import (
	"context"

	_ "devinggo/modules/{{.moduleName}}/logic"
	"devinggo/modules/{{.moduleName}}/router/{{.moduleName}}"
	"devinggo/modules/{{.moduleName}}/service"
	"devinggo/modules/system/pkg/modules"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type s{{.moduleNameCap}} struct{}

type {{.moduleName}}Module struct {
	Name   string
	Server *ghttp.Server
}

func New() *s{{.moduleNameCap}} {
	return &s{{.moduleNameCap}}{}
}

func init() {
	service.Register{{.moduleNameCap}}(New())

	module := &{{.moduleName}}Module{}
	module.Name = "{{.moduleName}}"
	if err := modules.Register(module); err != nil {
		panic(err)
	}
}

func (m *{{.moduleName}}Module) Start(ctx context.Context, s *ghttp.Server) error {
	m.Server = s
	s.BindHookHandler("/{{.moduleName}}/*", ghttp.HookBeforeServe, service.Hook().BeforeServe)
	s.BindHookHandler("/{{.moduleName}}/*", ghttp.HookAfterOutput, service.Hook().AfterOutput)
	s.Group("/{{.moduleName}}", func(group *ghttp.RouterGroup) {
		group.Middleware(service.Middleware().ApiAuth)
		{{.moduleName}}.BindController(group)
	})
	return nil
}

func (m *{{.moduleName}}Module) Stop(ctx context.Context) error {
	return nil
}

func (m *{{.moduleName}}Module) GetName() string {
	return m.Name
}

// GetModuleName 获取模块名称
func (s *s{{.moduleNameCap}}) GetModuleName(ctx context.Context) string {
	return "{{.moduleName}}"
}

// GetModuleVersion 获取模块版本
func (s *s{{.moduleNameCap}}) GetModuleVersion(ctx context.Context) string {
	return "1.0.0"
}

// Install 安装模块
func (s *s{{.moduleNameCap}}) Install(ctx context.Context) error {
	g.Log().Info(ctx, "{{.moduleName}} 模块安装")
	return nil
}

// Uninstall 卸载模块
func (s *s{{.moduleNameCap}}) Uninstall(ctx context.Context) error {
	g.Log().Info(ctx, "{{.moduleName}} 模块卸载")
	return nil
}

// Enable 启用模块
func (s *s{{.moduleNameCap}}) Enable(ctx context.Context) error {
	g.Log().Info(ctx, "{{.moduleName}} 模块启用")
	return nil
}

// Disable 禁用模块
func (s *s{{.moduleNameCap}}) Disable(ctx context.Context) error {
	g.Log().Info(ctx, "{{.moduleName}} 模块禁用")
	return nil
}
