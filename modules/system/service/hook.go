// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"github.com/gogf/gf/v2/net/ghttp"
)

type (
	IHook interface {
		BeforeServe(r *ghttp.Request)
		AfterOutput(r *ghttp.Request)
		// 是否忽略请求
		IsIgnoredRequest(r *ghttp.Request) bool
	}
)

var (
	localHook IHook
)

func Hook() IHook {
	if localHook == nil {
		panic("implement not found for interface IHook, forgot register?")
	}
	return localHook
}

func RegisterHook(i IHook) {
	localHook = i
}
