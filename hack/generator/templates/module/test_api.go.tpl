// Package test
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package api

import (
	"github.com/gogf/gf/v2/frame/g"
)

// TestReq 测试请求
type TestReq struct {
	g.Meta `path:"/{{.moduleName}}/test" method:"get" tags:"{{.moduleName}}" summary:"测试接口"`
}

// TestRes 测试响应
type TestRes struct {
	Message string `json:"message" dc:"返回消息"`
}
