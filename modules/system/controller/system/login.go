// Package system
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package system

import (
	"context"
	"devinggo/modules/system/api/system"
	"devinggo/modules/system/controller/base"
	"devinggo/modules/system/service"
)

var (
	LoginController = loginController{}
)

type loginController struct {
	base.BaseController
}

func (c *loginController) Login(ctx context.Context, req *system.LoginReq) (res *system.LoginRes, err error) {
	res = &system.LoginRes{}
	token, expire, err := service.Login().Login(ctx, req.Username, req.Password)
	if err != nil {
		return
	}
	res.Token = token
	res.Expire = expire
	return
}
