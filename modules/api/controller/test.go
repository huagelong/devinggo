// Package api
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package controller

import (
	"context"
	"devinggo/modules/api/api"
	"devinggo/modules/system/controller/base"
)

var (
	TestController = testController{}
)

type testController struct {
	base.BaseController
}

func (c *testController) Index(ctx context.Context, in *api.IndexReq) (out *api.IndexRes, err error) {
	out = &api.IndexRes{}
	//utils.SafeGo(ctx, func(ctx context.Context) {
	//	var i interface{} = "hello"
	//	num := i.(int) // 运行时 panic: interface conversion: interface {} is string, not int
	//	fmt.Println(num)
	//})
	out.Data = "立即使用 Devinggo，体验全栈开发的便捷与高效！"
	return
}
