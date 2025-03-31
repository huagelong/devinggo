// Package glob
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE
package glob

import (
	"context"
	"devinggo/api"
)

var (
	GlobController = globController{}
)

type globController struct {
}

func (c *globController) Health(ctx context.Context, in *api.HealthReq) (out *api.HealthRes, err error) {
	out = &api.HealthRes{}
	out.Data = "Hello, World!"
	return
}
