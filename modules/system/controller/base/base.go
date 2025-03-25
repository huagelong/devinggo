// Package base
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package base

import (
	"devinggo/modules/system/pkg/contexts"
	"github.com/gogf/gf/v2/net/ghttp"
)

type BaseController struct {
	UserId int64
}

// Init 自动执行的初始化方法
func (c *BaseController) Init(r *ghttp.Request) {
	c.UserId = contexts.New().GetUserId(r.Context())
}
