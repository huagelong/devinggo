// Package {{.moduleName}}
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package {{.moduleName}}

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
)

type s{{.moduleNameCap}}Logic struct{}

func New() *s{{.moduleNameCap}}Logic {
	return &s{{.moduleNameCap}}Logic{}
}

// Test 测试方法
func (s *s{{.moduleNameCap}}Logic) Test(ctx context.Context) string {
	g.Log().Info(ctx, "{{.moduleName}} Test method called")
	return "{{.moduleName}} module test success"
}
