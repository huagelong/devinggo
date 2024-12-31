// Package db
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE
package cache

import (
	"context"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/util/gconv"
	"regexp"
	"time"
)

type Adapter struct{}

func NewAdapter() gcache.Adapter {
	return &Adapter{}
}

func (a Adapter) getTable(ctx context.Context, key interface{}) string {
	keyStr := gconv.String(key)
	pattern := "^" + regexp.QuoteMeta(cachePrefixSelectCache) + "(.*?)@"
	re := regexp.MustCompile(pattern)
	if re.MatchString(keyStr) {
		tableExtracted := re.FindStringSubmatch(keyStr)[1]
		return tableExtracted
	} else {
		return ""
	}
}

func (a Adapter) Set(ctx context.Context, key interface{}, value interface{}, duration time.Duration) error {
	table := a.getTable(ctx, key)
	return Set(ctx, key, value, duration, table)
}

func (a Adapter) SetMap(ctx context.Context, data map[interface{}]interface{}, duration time.Duration) error {
	if len(data) == 0 {
		return nil
	}
	if duration < 0 {
		var (
			index = 0
			keys  = make([]string, len(data))
		)
		for k := range data {
			keys[index] = gconv.String(k)
			index += 1
		}
		for _, key := range keys {
			_, err := a.Remove(ctx, key)
			if err != nil {
				return err
			}
		}

	} else {
		var err error
		for k, v := range data {
			if err = a.Set(ctx, k, v, duration); err != nil {
				return err
			}
		}
	}
	return nil
}

func (a Adapter) SetIfNotExist(ctx context.Context, key interface{}, value interface{}, duration time.Duration) (ok bool, err error) {
	table := a.getTable(ctx, key)
	return SetIfNotExist(ctx, key, value, duration, table)
}

func (a Adapter) SetIfNotExistFunc(ctx context.Context, key interface{}, f gcache.Func, duration time.Duration) (ok bool, err error) {
	table := a.getTable(ctx, key)
	return SetIfNotExistFunc(ctx, key, f, duration, table)
}

func (a Adapter) SetIfNotExistFuncLock(ctx context.Context, key interface{}, f gcache.Func, duration time.Duration) (ok bool, err error) {
	table := a.getTable(ctx, key)
	return SetIfNotExistFuncLock(ctx, key, f, duration, table)
}

func (a Adapter) Get(ctx context.Context, key interface{}) (*gvar.Var, error) {
	return Get(ctx, key)
}

func (a Adapter) GetOrSet(ctx context.Context, key interface{}, value interface{}, duration time.Duration) (result *gvar.Var, err error) {
	table := a.getTable(ctx, key)
	return GetOrSet(ctx, key, value, duration, table)
}

func (a Adapter) GetOrSetFunc(ctx context.Context, key interface{}, f gcache.Func, duration time.Duration) (result *gvar.Var, err error) {
	table := a.getTable(ctx, key)
	return GetOrSetFunc(ctx, key, f, duration, table)
}

func (a Adapter) GetOrSetFuncLock(ctx context.Context, key interface{}, f gcache.Func, duration time.Duration) (result *gvar.Var, err error) {
	return a.GetOrSetFunc(ctx, key, f, duration)
}

func (a Adapter) Contains(ctx context.Context, key interface{}) (bool, error) {
	return Contains(ctx, key)
}

func (a Adapter) Size(ctx context.Context) (size int, err error) {
	return GetAdapterRedis().Size(ctx)
}

func (a Adapter) Data(ctx context.Context) (data map[interface{}]interface{}, err error) {
	return GetAdapterRedis().Data(ctx)
}

func (a Adapter) Keys(ctx context.Context) (keys []interface{}, err error) {
	return GetAdapterRedis().Keys(ctx)
}

func (a Adapter) Values(ctx context.Context) (values []interface{}, err error) {
	return GetAdapterRedis().Values(ctx)
}

func (a Adapter) Update(ctx context.Context, key interface{}, value interface{}) (oldValue *gvar.Var, exist bool, err error) {
	table := a.getTable(ctx, key)
	return Update(ctx, key, value, table)
}

func (a Adapter) UpdateExpire(ctx context.Context, key interface{}, duration time.Duration) (oldDuration time.Duration, err error) {
	table := a.getTable(ctx, key)
	return UpdateExpire(ctx, key, duration, table)
}

func (a Adapter) GetExpire(ctx context.Context, key interface{}) (time.Duration, error) {
	return GetExpire(ctx, key)
}

func (a Adapter) Remove(ctx context.Context, keys ...interface{}) (lastValue *gvar.Var, err error) {
	if !g.IsEmpty(keys) && len(keys) > 0 {
		for _, key := range keys {
			lastValue, err = Remove(ctx, key)
			if err != nil {
				return nil, err
			}
		}
		return lastValue, nil
	}
	return nil, nil
}

func (a Adapter) Clear(ctx context.Context) error {
	return GetAdapterRedis().Clear(ctx)
}

func (a Adapter) Close(ctx context.Context) error {
	return GetAdapterRedis().Close(ctx)
}
