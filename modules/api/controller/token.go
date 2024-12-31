// Package api
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package controller

import (
	"devinggo/internal/controller/base"
	"devinggo/modules/api/api"
	"devinggo/modules/system/pkg/utils/request"
	"devinggo/modules/system/service"
	"context"
)

var (
	TokenController = tokenController{}
)

type tokenController struct {
	base.BaseController
}

func (c *tokenController) Token(ctx context.Context, in *api.TokenReq) (out *api.TokenRes, err error) {
	out = &api.TokenRes{}
	params := make(map[string]interface{})
	params["app_id"] = in.AppId
	params["signature"] = in.Signature
	token, exp, err := service.SystemApp().GetAccessToken(ctx, params)
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
