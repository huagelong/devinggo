// Package modules
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE
package modules

import (
	"context"
	"devinggo/internal/dao"
	"devinggo/internal/model/entity"
	"devinggo/modules/system/pkg/utils"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"sync"
)

var (
	modules = make(map[string]Module)
	mLock   sync.Mutex
)

type Module interface {
	Start(ctx context.Context, server *ghttp.Server) error
	Stop(ctx context.Context) error
	GetName() string
}

func StartModules(ctx context.Context, server *ghttp.Server) (err error) {
	realModules, err := GetAllFilterModules(ctx)
	if err != nil {
		return
	}
	g.Log().Debug(ctx, "start modules", realModules)
	for name, module := range realModules {
		g.Log().Debug(ctx, name, "module start")
		if err = module.Start(ctx, server); err != nil {
			g.Log().Error(ctx, err)
			return
		}
	}
	return
}

func StopModules(ctx context.Context) (err error) {
	realModules, err := GetAllFilterModules(ctx)
	if err != nil {
		return
	}
	g.Log().Debug(ctx, "stop modules", realModules)
	for name, module := range realModules {
		g.Log().Debug(ctx, name, "module stop")
		if err = module.Stop(ctx); err != nil {
			g.Log().Error(ctx, err)
			return
		}
	}
	return
}

func Register(m Module) Module {
	mLock.Lock()
	defer mLock.Unlock()
	name := m.GetName()
	if _, ok := modules[name]; ok {
		panic("module already registered: " + name)
	}
	modules[name] = m
	return m
}

func GetAllFilterModules(ctx context.Context) (list map[string]Module, err error) {
	var dbModules []*entity.SystemModules
	list = make(map[string]Module, 0)
	err = dao.SystemModules.Ctx(ctx).Where("status", 1).Scan(&dbModules)
	if utils.IsError(err) {
		//保留system模块
		list["system"] = getSystemModule()
		return list, nil
	}
	if !g.IsEmpty(dbModules) {
		for key, m := range modules {
			for _, dbModule := range dbModules {
				if dbModule.Name == key {
					list[key] = m
					break
				}
			}
		}
	}
	if _, ok := list["key"]; !ok {
		list["system"] = getSystemModule()
	}

	return
}

func getSystemModule() Module {
	for key, m := range modules {
		if "system" == key {
			return m
		}
	}
	return nil
}
