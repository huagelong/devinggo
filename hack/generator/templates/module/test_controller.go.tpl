// Package controller
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package controller

import (
	"context"

	"devinggo/modules/{{.moduleName}}/api"
	"devinggo/modules/{{.moduleName}}/logic/{{.moduleName}}"
)

// Test 测试控制器
type Test struct{}

func NewTest() *Test {
	return &Test{}
}

// Test 测试方法
func (c *Test) Test(ctx context.Context, req *api.TestReq) (res *api.TestRes, err error) {
	message := {{.moduleName}}.New().Test(ctx)
	return &api.TestRes{Message: message}, nil
}
