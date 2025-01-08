// Package system
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package system

import (
	"context"
	"devinggo/modules/system/api/system"
	"devinggo/modules/system/controller/base"
	"devinggo/modules/system/model/res"
	"devinggo/modules/system/pkg/cache"
	"github.com/gogf/gf/v2/util/gconv"
)

var (
	CacheController = cacheController{}
)

type cacheController struct {
	base.BaseController
}

func (c *cacheController) GetCacheInfo(ctx context.Context, in *system.GetCacheInfoReq) (out *system.GetCacheInfoRes, err error) {
	out = &system.GetCacheInfoRes{}
	keys, err := cache.GetKeys(ctx)
	if err != nil {
		return
	}
	infoArr, err := cache.GetInfo(ctx)
	if err != nil {
		return
	}
	var infoServer = infoArr["Server"]
	var infoPersistence = infoArr["Persistence"]
	var infoMemory = infoArr["Memory"]
	var infoClients = infoArr["Clients"]
	var infoStats = infoArr["Stats"]
	redisMode := "单机"
	if gconv.String(infoServer["redis_mode"]) != "standalone" {
		redisMode = "集群"
	}
	aofEnabled := "关闭"
	if gconv.Int(infoPersistence["aof_enabled"]) != 0 {
		aofEnabled = "开启"
	}

	rs := res.CacheInfo{
		Keys: keys,
		Server: res.ServerInfo{
			Version:      gconv.String(infoServer["redis_version"]),
			RedisMode:    redisMode,
			RunDays:      gconv.String(infoServer["uptime_in_days"]),
			AofEnabled:   aofEnabled,
			UseMemory:    gconv.String(infoMemory["used_memory_human"]),
			Port:         gconv.String(infoServer["tcp_port"]),
			Clients:      gconv.String(infoClients["connected_clients"]),
			ExpiredKeys:  gconv.String(infoStats["expired_keys"]),
			SysTotalKeys: len(keys),
		},
	}
	out.Data = rs
	return
}

func (c *cacheController) ViewCache(ctx context.Context, in *system.ViewCacheReq) (out *system.ViewCacheRes, err error) {
	out = &system.ViewCacheRes{}
	rs, err := cache.Get(ctx, in.Key)
	if err != nil {
		return
	}
	content := rs.String()
	out.Content = content
	return
}

func (c *cacheController) DeleteCache(ctx context.Context, in *system.DeleteCacheReq) (out *system.DeleteCacheRes, err error) {
	out = &system.DeleteCacheRes{}
	_, err = cache.Remove(ctx, in.Key)
	if err != nil {
		return
	}
	return
}

func (c *cacheController) DeleteAllCache(ctx context.Context, in *system.DeleteAllCacheReq) (out *system.DeleteAllCacheRes, err error) {
	err = cache.ClearCacheAll(ctx)
	if err != nil {
		return
	}
	return
}
