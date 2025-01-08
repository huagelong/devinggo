// Package hook
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package hook

import (
	"devinggo/modules/api/service"
	"github.com/gogf/gf/v2/net/ghttp"
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
	s.apiAccessLog(r)
}
