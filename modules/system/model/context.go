// Package model
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package model

import (
	"github.com/gogf/gf/v2/frame/g"
)

type Context struct {
	Module          string    // 应用模块 system｜api｜home｜websocket
	User            *Identity // 上下文用户信息
	Data            g.Map     // 自定kv变量 业务模块根据需要设置，不固定
	AppId           string    // 应用ID
	Permission      string
	ExceptAuth      bool  // 是否排除权限验证
	ExceptLogin     bool  // 是否排除登录验证
	ExceptAccessLog bool  // 是否排除记录访问日志
	TenantId        int64 // 租户ID
	RequestBody     string
}
