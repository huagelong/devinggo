// Package config
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package config

import (
	"context"
	"time"

	"github.com/gogf/gf/v2/frame/g"
)

func GetConfigString(ctx context.Context, key string, defaultValue ...string) string {
	return g.Cfg().MustGet(ctx, key, defaultValue).String()
}

func GetConfigStrings(ctx context.Context, key string, defaultValue ...[]string) []string {
	return g.Cfg().MustGet(ctx, key, defaultValue).Strings()
}

func GetConfigArray(ctx context.Context, key string, defaultValue ...[]interface{}) []interface{} {
	return g.Cfg().MustGet(ctx, key, defaultValue).Array()
}

func GetConfigMap(ctx context.Context, key string, defaultValue ...map[string]interface{}) map[string]interface{} {
	return g.Cfg().MustGet(ctx, key, defaultValue).Map()
}

func GetConfigMaps(ctx context.Context, key string, defaultValue ...[]map[string]interface{}) []map[string]interface{} {
	return g.Cfg().MustGet(ctx, key, defaultValue).Maps()
}

func GetConfigInt64(ctx context.Context, key string, defaultValue ...int64) int64 {
	return g.Cfg().MustGet(ctx, key, defaultValue).Int64()
}

func GetConfigint64(ctx context.Context, key string, defaultValue ...int64) int64 {
	return g.Cfg().MustGet(ctx, key, defaultValue).Int64()
}

func GetConfigInt(ctx context.Context, key string, defaultValue ...int) int {
	return g.Cfg().MustGet(ctx, key, defaultValue).Int()
}

func GetConfigBool(ctx context.Context, key string, defaultValue ...bool) bool {
	return g.Cfg().MustGet(ctx, key, defaultValue).Bool()
}

func GetConfigDuration(ctx context.Context, key string, defaultValue ...string) time.Duration {
	return g.Cfg().MustGet(ctx, key, defaultValue).Duration()
}
