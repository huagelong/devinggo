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
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

type sContexts struct {
}

func New() *sContexts {
	return &sContexts{}
}

const ContextHTTPKey = "contextHTTPKey"

func (s *sContexts) Init(r *ghttp.Request, customCtx *model.Context) {
	r.SetCtxVar(ContextHTTPKey, customCtx)
}

// Get 获得上下文变量，如果没有设置，那么返回nil
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

func (s *sContexts) GetModule(ctx context.Context) string {
	c := s.Get(ctx)
	if c == nil {
		return ""
	}
	return c.Module
}

func (s *sContexts) SetUser(ctx context.Context, user *model.Identity) {
	c := s.Get(ctx)
	if c == nil {
		g.Log().Warning(ctx, "contexts.SetUser, c == nil ")
		return
	}
	c.User = user
}

func (s *sContexts) DelUser(ctx context.Context) {
	c := s.Get(ctx)
	c.User = &model.Identity{
		Id: 0,
	}
}

func (s *sContexts) GetUser(ctx context.Context) *model.Identity {
	c := s.Get(ctx)
	if c == nil {
		return nil
	}
	return c.User
}

// GetUserId 获取用户ID
func (s *sContexts) GetUserId(ctx context.Context) int64 {
	user := s.GetUser(ctx)
	if user == nil {
		return 0
	}
	return user.Id
}

func (s *sContexts) SetAppId(ctx context.Context, appId string) {
	c := s.Get(ctx)
	if c == nil {
		g.Log().Warning(ctx, "contexts.SetAppId, c == nil ")
		return
	}
	c.AppId = appId
}

func (s *sContexts) GetAppId(ctx context.Context) string {
	c := s.Get(ctx)
	if c == nil {
		g.Log().Warning(ctx, "contexts.GetAppId, c == nil ")
		return ""
	}
	if g.IsEmpty(c.AppId) {
		c.AppId = ""
	}
	return c.AppId
}

// SetData 设置额外数据
func (s *sContexts) SetData(ctx context.Context, k string, v interface{}) {
	c := s.Get(ctx)
	if c == nil {
		g.Log().Warning(ctx, "contexts.SetData, c == nil ")
		return
	}
	s.Get(ctx).Data[k] = v
}

// SetDataMap 设置额外数据
func (s *sContexts) SetDataMap(ctx context.Context, vs g.Map) {
	c := s.Get(ctx)
	if c == nil {
		g.Log().Warning(ctx, "contexts.SetData, c == nil ")
		return
	}
	for k, v := range vs {
		s.Get(ctx).Data[k] = v
	}
}

// GetData 获取额外数据
func (s *sContexts) GetData(ctx context.Context) g.Map {
	c := s.Get(ctx)
	if c == nil {
		return nil
	}
	return c.Data
}

func (s *sContexts) GetPermission(ctx context.Context) string {
	c := s.Get(ctx)
	if c == nil {
		return ""
	}
	if g.IsEmpty(c.Permission) {
		return ""
	}
	return c.Permission
}

func (s *sContexts) SetPermission(ctx context.Context, permission string) {
	c := s.Get(ctx)
	if c == nil {
		g.Log().Warning(ctx, "contexts.SetPermission, c == nil ")
		return
	}
	c.Permission = permission
}

func (s *sContexts) GetExceptAuth(ctx context.Context) bool {
	c := s.Get(ctx)
	if c == nil {
		return false
	}
	if g.IsEmpty(c.ExceptAuth) {
		return false
	}
	return c.ExceptAuth
}

func (s *sContexts) SetExceptAuth(ctx context.Context, exceptAuth bool) {
	c := s.Get(ctx)
	if c == nil {
		g.Log().Warning(ctx, "contexts.SetExceptAuth, c == nil ")
		return
	}
	c.ExceptAuth = exceptAuth
}

func (s *sContexts) GetExceptLogin(ctx context.Context) bool {
	c := s.Get(ctx)
	if c == nil {
		return false
	}
	if g.IsEmpty(c.ExceptLogin) {
		return false
	}
	return c.ExceptLogin
}

func (s *sContexts) SetExceptLogin(ctx context.Context, exceptLogin bool) {
	c := s.Get(ctx)
	if c == nil {
		g.Log().Warning(ctx, "contexts.SetExceptLogin, c == nil ")
		return
	}
	c.ExceptLogin = exceptLogin
}

func (s *sContexts) GetExceptAccessLog(ctx context.Context) bool {
	c := s.Get(ctx)
	if c == nil {
		return false
	}
	if g.IsEmpty(c.ExceptAccessLog) {
		return false
	}
	return c.ExceptAccessLog
}

func (s *sContexts) SetExceptAccessLog(ctx context.Context, exceptAccessLog bool) {
	c := s.Get(ctx)
	if c == nil {
		g.Log().Warning(ctx, "contexts.SetExceptAccessLog, c == nil ")
		return
	}
	c.ExceptAccessLog = exceptAccessLog
}

func (s *sContexts) SetTenantId(ctx context.Context, tenantId int64) {
	c := s.Get(ctx)
	if c == nil {
		g.Log().Warning(ctx, "contexts.SetTenantId, c == nil ")
		return
	}
	c.TenantId = tenantId
}

func (s *sContexts) GetTenantId(ctx context.Context) int64 {
	c := s.Get(ctx)
	if c == nil {
		panic("TenantId is empty")
	}

	enable := config.GetConfigBool(ctx, "tenant.enable", false)
	if enable {
		if g.IsEmpty(c.TenantId) {
			panic("TenantId is empty")
		}
	} else {
		return gconv.Int64(consts.DefaultTenantId)
	}
	return c.TenantId
}

func (s *sContexts) GetTakeUpTime(ctx context.Context) int64 {
	r := request.GetHttpRequest(ctx)
	return gtime.Now().Sub(gtime.New(r.EnterTime)).Milliseconds()
}

func (s *sContexts) GetRequestBody(ctx context.Context) string {
	c := s.Get(ctx)
	if c == nil {
		return "{}"
	}
	if g.IsEmpty(c.Permission) {
		return "{}"
	}
	return c.RequestBody
}

func (s *sContexts) SetRequestBody(ctx context.Context, requestBody string) {
	c := s.Get(ctx)
	if c == nil {
		g.Log().Warning(ctx, "contexts.SetRequestBody, c == nil ")
		return
	}
	c.RequestBody = requestBody
}
