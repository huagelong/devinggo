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
	"devinggo/modules/system/pkg/worker/cron"
	"devinggo/modules/system/service"
	"github.com/gogf/gf/v2/frame/g"
)

var (
	TestController = testController{}
)

type testController struct {
	base.BaseController
}

func (c *testController) Index(ctx context.Context, in *api.IndexReq) (out *api.IndexRes, err error) {
	out = &api.IndexRes{}
	out.Data = "Hello, World!"
	return
}
