// Package service
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package service

import (
	"context"
)

type I{{.moduleNameCap}} interface {
	GetModuleName(ctx context.Context) string
	GetModuleVersion(ctx context.Context) string
	Install(ctx context.Context) error
	Uninstall(ctx context.Context) error
	Enable(ctx context.Context) error
	Disable(ctx context.Context) error
}

var local{{.moduleNameCap}} I{{.moduleNameCap}}

func {{.moduleNameCap}}() I{{.moduleNameCap}} {
	if local{{.moduleNameCap}} == nil {
		panic("implement not found for interface I{{.moduleNameCap}}, forgot register?")
	}
	return local{{.moduleNameCap}}
}

func Register{{.moduleNameCap}}(i I{{.moduleNameCap}}) {
	local{{.moduleNameCap}} = i
}
