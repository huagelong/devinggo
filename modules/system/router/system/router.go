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
			system.CodeController,
			system.DataMaintainController,
			system.SystemModulesController,
		).Middleware(service.Middleware().AdminAuth)
	})

}
