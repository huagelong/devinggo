// Package service
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package service

import (
	"github.com/gogf/gf/v2/net/ghttp"
)

type IHook interface {
	BeforeServe(r *ghttp.Request)
	AfterOutput(r *ghttp.Request)
}

var localHook IHook

func Hook() IHook {
	if localHook == nil {
		panic("implement not found for interface IHook, forgot register?")
	}
	return localHook
}

func RegisterHook(i IHook) {
	localHook = i
}
