// Package hook
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package hook

import (
	"devinggo/modules/system/service"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/text/gstr"
	"strings"
)

type sHook struct {
}

func init() {
	service.RegisterHook(NewHook())
}

func NewHook() *sHook {
	return &sHook{}
}

func (s *sHook) BeforeServe(r *ghttp.Request) {
}

func (s *sHook) AfterOutput(r *ghttp.Request) {
	s.accessLog(r)
}

// 忽略的请求方式
var ignoredRequestMethods = []string{"HEAD", "PRI", "OPTIONS"}

// 是否忽略请求
func (s *sHook) IsIgnoredRequest(r *ghttp.Request) bool {
	if r.IsFileRequest() {
		return true
	}

	if gstr.InArray(ignoredRequestMethods, strings.ToUpper(r.Method)) {
		return true
	}

	return false
}
