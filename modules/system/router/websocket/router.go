// Package websocket
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package websocket

import (
	"devinggo/modules/system/pkg/websocket"
	"devinggo/modules/system/service"
	"github.com/gogf/gf/v2/net/ghttp"
)

func BindController(group *ghttp.RouterGroup) {
	group.Group("/ws", func(group *ghttp.RouterGroup) {
		group.Bind(
			websocket.WsPage,
		).Middleware(service.Middleware().WsAuth)
	})

}
