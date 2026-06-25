// Package system
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package system

import (
	"devinggo/modules/system/controller/system"
	"devinggo/modules/system/service"

	"github.com/gogf/gf/v2/net/ghttp"
)

func BindController(group *ghttp.RouterGroup) {
	// Pusher HTTP Events API (无需认证，通过HMAC签名验证)
	group.Bind(
		system.PusherEvents,
		system.PusherWebhook, // Pusher Webhook 验证
		system.PusherChannel, // Pusher Channel API (查询频道状态)
	)

	group.Group("/system", func(group *ghttp.RouterGroup) {
		group.Bind(
			system.LoginController,
			system.LogoutController,
			system.RefreshController,
			system.UserController,
			system.CommonController,
			system.DictController,
			system.MessageController,
			system.UploadController,
			system.DeptController,
			system.MenuController,
			system.PostController,
			system.RoleController,
			system.LogsController,
			system.ConfigController,
			system.CrontabController,
			system.NoticeController,
			system.AttachmentController,
			system.AppGroupController,
			system.AppController,
			system.ApiController,
			system.ApiGroupController,
			system.CacheController,
			system.DataMaintainController,
			system.SystemModulesController,
			system.DashboardController,
			system.PusherAuthController,
			system.PusherUserAuthController, // Pusher User Authentication
			system.CodeGenController,        // 代码生成
		).Middleware(service.Middleware().AdminAuth)
	})

}
