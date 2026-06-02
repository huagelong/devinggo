// Package cache
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package cache

import (
	"context"
	"regexp"
	"time"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/util/gconv"
)

// ──────────────────────────────────────────────
// gcache.Adapter 实现
// 用于 GoFrame ORM 查询缓存和 GetCache() 实例
// ──────────────────────────────────────────────

// Adapter 实现 gcache.Adapter 接口，为 GoFrame 框架提供缓存支持。
// 该适配器会自动从 ORM 缓存键中提取表名作为标签。
type Adapter struct{}

func newCacheAdapter() gcache.Adapter {
	return &Adapter{}
}

// getTable 从缓存键中提取表名作为标签。
// 用于 ORM 查询缓存的自动标签关联。
// 例如：SelectCache:system_user@hash... -> "system_user"
func (a Adapter) getTable(ctx context.Context, key interface{}) string {
	keyStr := gconv.String(key)
	pattern := "^" + regexp.QuoteMeta(cachePrefixSelectCache) + "(.*?)@"
	re := regexp.MustCompile(pattern)
	if !re.MatchString(keyStr) {
		return ""
	}
	return re.FindStringSubmatch(keyStr)[1]
}

func (a Adapter) Set(ctx context.Context, key interface{}, value interface{}, duration time.Duration) error {
	return set(ctx, key, value, duration, a.getTable(ctx, key))
}

func (a Adapter) SetMap(ctx context.Context, data map[interface{}]interface{}, duration time.Duration) error {
	if len(data) == 0 {
		return nil
	}
	if duration < 0 {
		keys := make([]string, 0, len(data))
		for k := range data {
			keys = append(keys, gconv.String(k))
		}
		for _, key := range keys {
			if _, err := a.Remove(ctx, key); err != nil {
				return err
			}
		}
	} else {
		for k, v := range data {
			if err := a.Set(ctx, k, v, duration); err != nil {
				return err
			}
		}
	}
	return nil
}

func (a Adapter) SetIfNotExist(ctx context.Context, key interface{}, value interface{}, duration time.Duration) (bool, error) {
	return setIfNotExist(ctx, key, value, duration, a.getTable(ctx, key))
}

func (a Adapter) SetIfNotExistFunc(ctx context.Context, key interface{}, f gcache.Func, duration time.Duration) (bool, error) {
	return setIfNotExistFunc(ctx, key, f, duration, a.getTable(ctx, key))
}

func (a Adapter) SetIfNotExistFuncLock(ctx context.Context, key interface{}, f gcache.Func, duration time.Duration) (bool, error) {
	return setIfNotExistFuncLock(ctx, key, f, duration, a.getTable(ctx, key))
}

func (a Adapter) Get(ctx context.Context, key interface{}) (*gvar.Var, error) {
	return Get(ctx, key)
}

func (a Adapter) GetOrSet(ctx context.Context, key interface{}, value interface{}, duration time.Duration) (*gvar.Var, error) {
	return getOrSet(ctx, key, value, duration, a.getTable(ctx, key))
}

func (a Adapter) GetOrSetFunc(ctx context.Context, key interface{}, f gcache.Func, duration time.Duration) (*gvar.Var, error) {
	return getOrSetFunc(ctx, key, f, duration, a.getTable(ctx, key))
}

func (a Adapter) GetOrSetFuncLock(ctx context.Context, key interface{}, f gcache.Func, duration time.Duration) (*gvar.Var, error) {
	return a.GetOrSetFunc(ctx, key, f, duration)
}

func (a Adapter) Contains(ctx context.Context, key interface{}) (bool, error) {
	return contains(ctx, key)
}

func (a Adapter) Size(ctx context.Context) (int, error) {
	return getAdapterRedis().Size(ctx)
}

func (a Adapter) Data(ctx context.Context) (map[interface{}]interface{}, error) {
	return getAdapterRedis().Data(ctx)
}

func (a Adapter) Keys(ctx context.Context) ([]interface{}, error) {
	return getAdapterRedis().Keys(ctx)
}

func (a Adapter) Values(ctx context.Context) ([]interface{}, error) {
	return getAdapterRedis().Values(ctx)
}

func (a Adapter) Update(ctx context.Context, key interface{}, value interface{}) (*gvar.Var, bool, error) {
	return update(ctx, key, value, a.getTable(ctx, key))
}

func (a Adapter) UpdateExpire(ctx context.Context, key interface{}, duration time.Duration) (time.Duration, error) {
	return updateExpire(ctx, key, duration, a.getTable(ctx, key))
}

func (a Adapter) GetExpire(ctx context.Context, key interface{}) (time.Duration, error) {
	return getExpire(ctx, key)
}

func (a Adapter) Remove(ctx context.Context, keys ...interface{}) (*gvar.Var, error) {
	var lastValue *gvar.Var
	for _, key := range keys {
		v, err := Remove(ctx, key)
		if err != nil {
			return nil, err
		}
		lastValue = v
	}
	return lastValue, nil
}

func (a Adapter) Clear(ctx context.Context) error {
	return getAdapterRedis().Clear(ctx)
}

func (a Adapter) Close(ctx context.Context) error {
	return getAdapterRedis().Close(ctx)
}
