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
		// system 访问日志
		AccessLog(r *ghttp.Request)
		BeforeServe(r *ghttp.Request)
		AfterOutput(r *ghttp.Request)
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