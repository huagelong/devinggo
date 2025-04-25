// Package api
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package api

import (
	"devinggo/modules/system/model"
	"github.com/gogf/gf/v2/frame/g"
)

type TokenReq struct {
	g.Meta `path:"/getToken" method:"get" tags:"token" summary:"获取token." x-exceptAuth:"true" x-exceptLogin:"true" x-permission:"api:token:getToken" `
	model.EasyModeVerify
}

type TokenRes struct {
	g.Meta `mime:"application/json"`
	Token  string `json:"token" dc:"token"`
	Expire int64  `json:"expire" dc:"expire time"`
}

type RefreshTokenReq struct {
	g.Meta `path:"/refreshToken" tags:"token" method:"get" summary:"refresh token" x-exceptAuth:"true" x-exceptLogin:"true"  x-permission:"api:token:refreshToken"`
	model.AuthorHeader
}

type RefreshTokenRes struct {
	g.Meta `mime:"application/json"`
	Token  string `json:"token" dc:"new token"`
	Expire int64  `json:"expire" dc:"expire time"`
}
