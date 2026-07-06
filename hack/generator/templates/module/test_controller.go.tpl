// Package {{.moduleName}}
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package {{.moduleName}}

import (
	"context"

	{{.moduleName}}api "devinggo/modules/{{.moduleName}}/api/{{.moduleName}}"
	{{.moduleName}}logic "devinggo/modules/{{.moduleName}}/logic/{{.moduleName}}"
)

var (
	TestController = testController{}
)

type testController struct{}

// Test 测试方法
func (c *testController) Test(ctx context.Context, req *{{.moduleName}}api.TestReq) (res *{{.moduleName}}api.TestRes, err error) {
	message := {{.moduleName}}logic.New().Test(ctx)
	return &{{.moduleName}}api.TestRes{Message: message}, nil
}
