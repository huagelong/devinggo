// Package middleware
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package middleware

import (
	"devinggo/modules/api/service"
	"devinggo/modules/system/pkg/contexts"
	"context"
)

type sMiddleware struct {
}

func init() {
	service.RegisterMiddleware(NewMiddleware())
}

func NewMiddleware() *sMiddleware {
	return &sMiddleware{}
}

// IsExceptAuth 是否是不需要验证权限的路由地址
func (s *sMiddleware) isExceptAuth(ctx context.Context) bool {
	return contexts.New().GetExceptAuth(ctx)
}

// IsExceptLogin 是否是不需要登录的路由地址
func (s *sMiddleware) isExceptLogin(ctx context.Context) bool {
	return contexts.New().GetExceptLogin(ctx)
}
