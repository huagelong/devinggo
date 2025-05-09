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
	"devinggo/modules/system/pkg/utils/request"
	"devinggo/modules/system/service"
)

var (
	TokenController = tokenController{}
)

type tokenController struct {
	base.BaseController
}

func (c *tokenController) Token(ctx context.Context, in *api.TokenReq) (out *api.TokenRes, err error) {
	out = &api.TokenRes{}
	token, exp, err := service.SystemApp().GetAccessToken(ctx)
	if err != nil {
		return nil, err
	}
	out.Token = token
	out.Expire = exp
	return
}

func (c *tokenController) RefreshToken(ctx context.Context, in *api.RefreshTokenReq) (out *api.RefreshTokenRes, err error) {
	out = &api.RefreshTokenRes{}
	r := request.GetHttpRequest(ctx)
	token, exp, err := service.Token().Refresh(r)
	if err != nil {
		return nil, err
	}
	out.Token = token
	out.Expire = exp
	return
}
