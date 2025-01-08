// Package redis
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package redis

import (
	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcache"
)

var group = "default"

func GetCacheClient() *gcache.Cache {
	c := gcache.New()
	adapter := gcache.NewAdapterRedis(g.Redis(group))
	c.SetAdapter(adapter)
	return c
}

func GetRedis() *gredis.Redis {
	return g.Redis(group)
}
