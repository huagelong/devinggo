// Package middleware
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package middleware

import (
	"reflect"

	"devinggo/modules/system/model"
	"devinggo/modules/system/pkg/contexts"
	"devinggo/modules/system/pkg/utils"
	"devinggo/modules/system/service"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/gmeta"
)

func (s *sMiddleware) Ctx(r *ghttp.Request) {

	if r.Method == "OPTIONS" {
		r.Middleware.Next()
		return
	}

	context := &model.Context{
		Module: utils.GetModule(r.URL.Path),
		Data:   make(g.Map),
	}
	contexts.Init(r, context)
	ctx := r.GetCtx()
	appId := r.GetHeader("X-App-Id")
	if g.IsEmpty(appId) {
		appIdRs := r.Get("app_id")
		if !g.IsEmpty(appIdRs) {
			appId = appIdRs.String()
		}
	}
	if !g.IsEmpty(appId) {
		contexts.SetAppId(ctx, appId)
	}
	if !gstr.Contains(gstr.ToLower(r.GetHeader("content-type")), "multipart/form-data") {
		gjson := gjson.New(r.GetBodyString())
		if gjson.Contains("password") {
			gjson.Remove("password")
			contexts.SetRequestBody(ctx, gjson.String())
		} else {
			contexts.SetRequestBody(ctx, r.GetBodyString())
		}
	}

	tenantIdStr := r.GetHeader("X-Tenant-Id")
	if !g.IsEmpty(tenantIdStr) {
		contexts.SetTenantId(ctx, gconv.Int64(tenantIdStr))
	}
	s.meta(r)
	err := s.userCtx(r)
	if err != nil {
		g.Log().Debug(ctx, "userCtx error:", err)
	}
	r.Middleware.Next()
}

func (s *sMiddleware) meta(r *ghttp.Request) {
	ctx := r.GetCtx()
	permission := ""
	exceptAuth := false
	exceptLogin := false
	exceptAccessLog := false
	if g.IsEmpty(r.GetServeHandler()) {
		return
	}
	handler := r.GetServeHandler().Handler
	if handler.Info.Type != nil && handler.Info.Type.NumIn() == 2 {
		var objectReq = reflect.New(handler.Info.Type.In(1))
		if v := gmeta.Get(objectReq, "x-permission"); !v.IsEmpty() {
			permission = v.String()
		}

		if v := gmeta.Get(objectReq, "x-exceptAuth"); !v.IsEmpty() {
			exceptAuth = v.Bool()
		}

		if v := gmeta.Get(objectReq, "x-exceptLogin"); !v.IsEmpty() {
			exceptLogin = v.Bool()
		}

		if v := gmeta.Get(objectReq, "x-exceptAccessLog"); !v.IsEmpty() {
			exceptAccessLog = v.Bool()
		}
	}
	contexts.SetPermission(ctx, permission)
	contexts.SetExceptAuth(ctx, exceptAuth)
	contexts.SetExceptLogin(ctx, exceptLogin)
	contexts.SetExceptAccessLog(ctx, exceptAccessLog)

}

func (s *sMiddleware) userCtx(r *ghttp.Request, appId ...string) (err error) {
	ctx := r.GetCtx()
	newAppId := ""
	if len(appId) > 0 {
		newAppId = appId[0]
		contexts.SetAppId(ctx, newAppId)
	} else {
		newAppId = contexts.GetAppId(ctx)
	}
	user, err := service.Token().ParseLoginUser(r, newAppId)
	if utils.IsError(err) {
		contexts.DelUser(ctx)
		return
	}
	if !g.IsEmpty(user) {
		contexts.SetUser(ctx, user)
		contexts.SetAppId(ctx, user.AppId)
	}
	g.Log().Debug(ctx, "ctx-appId:", contexts.GetAppId(ctx))
	return
}
