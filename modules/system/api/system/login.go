// Package system
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package system

import (
	"devinggo/modules/system/model"
	"github.com/gogf/gf/v2/frame/g"
)

type LoginReq struct {
	g.Meta   `path:"/login" method:"post" tags:"登录" summary:"登录" x-exceptAuth:"true" x-exceptLogin:"true" x-permission:"system:user:login"`
	Username string `json:"username" v:"required|length:4,12#请输入用户名|用户名长度为4~12位" dc:"用户名"`
	Password string `json:"password" v:"required#请输入密码" dc:"密码" `
}

type LoginRes struct {
	g.Meta `mime:"application/json"`
	Token  string `json:"token"` // 登录成功后返回的token
	Expire int64  `json:"expire" dc:"expire time"`
}

type LogoutReq struct {
	g.Meta `path:"/logout" method:"post" tags:"登录" summary:"退出" x-exceptAuth:"true" x-exceptLogin:"true" x-permission:"system:user:logout"`
	model.AuthorHeader
}

type LogoutRes struct {
	g.Meta `mime:"application/json"`
}

type RefreshReq struct {
	g.Meta `path:"/refresh" method:"post" tags:"登录" summary:"刷新" x-permission:"system:user:refresh"`
	model.AuthorHeader
}

type RefreshRes struct {
	g.Meta `mime:"application/json"`
	Token  string `json:"token"` // 返回的token
	Expire int64  `json:"expire" dc:"expire time"`
}
