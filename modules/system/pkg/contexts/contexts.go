// Package contexts
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package contexts

import (
	"context"

	"devinggo/modules/system/consts"
	"devinggo/modules/system/model"
	"devinggo/modules/system/pkg/utils/config"
	"devinggo/modules/system/pkg/utils/request"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

type sContexts struct{}

// std 包级单例，避免每次调用 New() 创建临时对象
var std = &sContexts{}

// New 返回包级单例（保留以向后兼容）
func New() *sContexts { return std }

// ---- 包级函数，外部可直接 contexts.GetUserId(ctx) 调用 ----

func Init(r *ghttp.Request, customCtx *model.Context)        { std.Init(r, customCtx) }
func Get(ctx context.Context) *model.Context                 { return std.Get(ctx) }
func GetModule(ctx context.Context) string                   { return std.GetModule(ctx) }
func SetUser(ctx context.Context, user *model.Identity)      { std.SetUser(ctx, user) }
func DelUser(ctx context.Context)                            { std.DelUser(ctx) }
func GetUser(ctx context.Context) *model.Identity            { return std.GetUser(ctx) }
func GetUserId(ctx context.Context) int64                    { return std.GetUserId(ctx) }
func SetAppId(ctx context.Context, appId string)             { std.SetAppId(ctx, appId) }
func GetAppId(ctx context.Context) string                    { return std.GetAppId(ctx) }
func SetData(ctx context.Context, k string, v interface{})   { std.SetData(ctx, k, v) }
func SetDataMap(ctx context.Context, vs g.Map)               { std.SetDataMap(ctx, vs) }
func GetData(ctx context.Context) g.Map                      { return std.GetData(ctx) }
func GetPermission(ctx context.Context) string               { return std.GetPermission(ctx) }
func SetPermission(ctx context.Context, permission string)   { std.SetPermission(ctx, permission) }
func GetExceptAuth(ctx context.Context) bool                 { return std.GetExceptAuth(ctx) }
func SetExceptAuth(ctx context.Context, v bool)              { std.SetExceptAuth(ctx, v) }
func GetExceptLogin(ctx context.Context) bool                { return std.GetExceptLogin(ctx) }
func SetExceptLogin(ctx context.Context, v bool)             { std.SetExceptLogin(ctx, v) }
func GetExceptAccessLog(ctx context.Context) bool            { return std.GetExceptAccessLog(ctx) }
func SetExceptAccessLog(ctx context.Context, v bool)         { std.SetExceptAccessLog(ctx, v) }
func SetTenantId(ctx context.Context, tenantId int64)        { std.SetTenantId(ctx, tenantId) }
func GetTenantId(ctx context.Context) (int64, error)         { return std.GetTenantId(ctx) }
func GetTakeUpTime(ctx context.Context) int64                { return std.GetTakeUpTime(ctx) }
func GetRequestBody(ctx context.Context) string              { return std.GetRequestBody(ctx) }
func SetRequestBody(ctx context.Context, requestBody string) { std.SetRequestBody(ctx, requestBody) }
func SetLanguage(ctx context.Context, lang string)           { std.SetLanguage(ctx, lang) }
func GetLanguage(ctx context.Context) string                 { return std.GetLanguage(ctx) }

const (
	ContextHTTPKey = "contextHTTPKey"
)

func (s *sContexts) Init(r *ghttp.Request, customCtx *model.Context) {
	r.SetCtxVar(ContextHTTPKey, customCtx)
}

// Get 获得上下文变量，如果没有设置则返回 nil
func (s *sContexts) Get(ctx context.Context) *model.Context {
	value := ctx.Value(ContextHTTPKey)
	if value == nil {
		return nil
	}
	if localCtx, ok := value.(*model.Context); ok {
		return localCtx
	}
	return nil
}

// getCtx 获取上下文，若为 nil 则打印警告日志
func (s *sContexts) getCtx(ctx context.Context, caller string) *model.Context {
	c := s.Get(ctx)
	if c == nil {
		g.Log().Warningf(ctx, "%s, c == nil", caller)
	}
	return c
}

// getField 通用字段读取：c 为 nil 时返回零值
func getField[T any](s *sContexts, ctx context.Context, getter func(*model.Context) T) T {
	c := s.Get(ctx)
	if c == nil {
		var zero T
		return zero
	}
	return getter(c)
}

// setField 通用字段写入：c 为 nil 时记录警告后跳过
func (s *sContexts) setField(ctx context.Context, caller string, setter func(*model.Context)) {
	if c := s.getCtx(ctx, caller); c != nil {
		setter(c)
	}
}

func (s *sContexts) GetModule(ctx context.Context) string {
	return getField(s, ctx, func(c *model.Context) string { return c.Module })
}

func (s *sContexts) SetUser(ctx context.Context, user *model.Identity) {
	s.setField(ctx, "contexts.user", func(c *model.Context) { c.User = user })
}

// DelUser 清除当前用户信息
func (s *sContexts) DelUser(ctx context.Context) {
	s.setField(ctx, "contexts.user", func(c *model.Context) { c.User = &model.Identity{Id: 0} })
}

func (s *sContexts) GetUser(ctx context.Context) *model.Identity {
	return getField(s, ctx, func(c *model.Context) *model.Identity { return c.User })
}

// GetUserId 获取用户ID
func (s *sContexts) GetUserId(ctx context.Context) int64 {
	if user := s.GetUser(ctx); user != nil {
		return user.Id
	}
	return 0
}

func (s *sContexts) SetAppId(ctx context.Context, appId string) {
	s.setField(ctx, "contexts.appId", func(c *model.Context) { c.AppId = appId })
}

func (s *sContexts) GetAppId(ctx context.Context) string {
	return getField(s, ctx, func(c *model.Context) string { return c.AppId })
}

// SetData 设置额外数据
func (s *sContexts) SetData(ctx context.Context, k string, v interface{}) {
	s.setField(ctx, "contexts.data", func(c *model.Context) { c.Data[k] = v })
}

// SetDataMap 批量设置额外数据
func (s *sContexts) SetDataMap(ctx context.Context, vs g.Map) {
	s.setField(ctx, "contexts.data", func(c *model.Context) {
		for k, v := range vs {
			c.Data[k] = v
		}
	})
}

// GetData 获取额外数据
func (s *sContexts) GetData(ctx context.Context) g.Map {
	return getField(s, ctx, func(c *model.Context) g.Map { return c.Data })
}

func (s *sContexts) GetPermission(ctx context.Context) string {
	return getField(s, ctx, func(c *model.Context) string { return c.Permission })
}

func (s *sContexts) SetPermission(ctx context.Context, permission string) {
	s.setField(ctx, "contexts.permission", func(c *model.Context) { c.Permission = permission })
}

func (s *sContexts) GetExceptAuth(ctx context.Context) bool {
	return getField(s, ctx, func(c *model.Context) bool { return c.ExceptAuth })
}

func (s *sContexts) SetExceptAuth(ctx context.Context, exceptAuth bool) {
	s.setField(ctx, "contexts.exceptAuth", func(c *model.Context) { c.ExceptAuth = exceptAuth })
}

func (s *sContexts) GetExceptLogin(ctx context.Context) bool {
	return getField(s, ctx, func(c *model.Context) bool { return c.ExceptLogin })
}

func (s *sContexts) SetExceptLogin(ctx context.Context, exceptLogin bool) {
	s.setField(ctx, "contexts.exceptLogin", func(c *model.Context) { c.ExceptLogin = exceptLogin })
}

func (s *sContexts) GetExceptAccessLog(ctx context.Context) bool {
	return getField(s, ctx, func(c *model.Context) bool { return c.ExceptAccessLog })
}

func (s *sContexts) SetExceptAccessLog(ctx context.Context, exceptAccessLog bool) {
	s.setField(ctx, "contexts.exceptAccessLog", func(c *model.Context) { c.ExceptAccessLog = exceptAccessLog })
}

func (s *sContexts) SetTenantId(ctx context.Context, tenantId int64) {
	s.setField(ctx, "contexts.tenantId", func(c *model.Context) { c.TenantId = tenantId })
}

// GetTenantId 获取租户ID，未开启多租户时返回默认租户ID
func (s *sContexts) GetTenantId(ctx context.Context) (int64, error) {
	if !config.GetConfigBool(ctx, "tenant.enable", false) {
		return gconv.Int64(consts.DefaultTenantId), nil
	}
	c := s.Get(ctx)
	if c == nil || c.TenantId == 0 {
		return 0, gerror.New("TenantId is empty")
	}
	return c.TenantId, nil
}

func (s *sContexts) GetTakeUpTime(ctx context.Context) int64 {
	r := request.GetHttpRequest(ctx)
	return gtime.Now().Sub(gtime.New(r.EnterTime)).Milliseconds()
}

// GetRequestBody 获取请求体内容，若为空则返回 "{}"
func (s *sContexts) GetRequestBody(ctx context.Context) string {
	if body := getField(s, ctx, func(c *model.Context) string { return c.RequestBody }); body != "" {
		return body
	}
	return "{}"
}

func (s *sContexts) SetRequestBody(ctx context.Context, requestBody string) {
	s.setField(ctx, "contexts.requestBody", func(c *model.Context) { c.RequestBody = requestBody })
}

// SetLanguage 设置当前请求语言
func (s *sContexts) SetLanguage(ctx context.Context, lang string) {
	s.setField(ctx, "contexts.language", func(c *model.Context) { c.Language = lang })
}

// GetLanguage 获取当前请求语言
func (s *sContexts) GetLanguage(ctx context.Context) string {
	return getField(s, ctx, func(c *model.Context) string { return c.Language })
}
