// Package {{.moduleName}}
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package {{.moduleName}}

import (
	"context"
	"devinggo/modules/{{.moduleName}}/service"

	"github.com/gogf/gf/v2/frame/g"
)

type s{{.moduleNameCap}} struct {
}

func New() *s{{.moduleNameCap}} {
	return &s{{.moduleNameCap}}{}
}

func init() {
	service.Register{{.moduleNameCap}}(New())
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
