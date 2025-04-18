// Package cache
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package cache

import (
	"context"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/util/gconv"
	"time"
)

func getTagCache(ctx context.Context) (*TagCache, error) {
	return NewTagCache(ctx, g.Redis(groupKey))
}

func Set(ctx context.Context, key interface{}, value interface{}, duration time.Duration, tag ...interface{}) error {
	if !g.IsEmpty(tag) && len(tag) > 0 {
		tagCache, err := getTagCache(ctx)
		if err != nil {
			return err
		}
		if value == nil || duration < 0 {
			_, err := Remove(ctx, key)
			if err != nil {
				return err
			}
		} else {
			err = tagCache.Set(ctx, gconv.String(key), value, duration, gconv.Strings(tag))
			if err != nil {
				return err
			}
		}
	} else {
		return GetAdapterRedis().Set(ctx, key, value, duration)
	}
	return nil
}

func SetIfNotExist(ctx context.Context, key interface{}, value interface{}, duration time.Duration, tag ...interface{}) (ok bool, err error) {
	if !g.IsEmpty(tag) && len(tag) > 0 {
		// Execute the function and retrieve the result.
		f, ok := value.(gcache.Func)
		if !ok {
			// Compatible with raw function value.
			f, ok = value.(func(ctx context.Context) (value interface{}, err error))
		}
		if ok {
			if value, err = f(ctx); err != nil {
				return false, err
			}
		}

		tagCache, err := getTagCache(ctx)
		if err != nil {
			return false, err
		}

		// DEL.
		if duration < 0 || value == nil {
			err := tagCache.Delete(ctx, gconv.String(key))
			if err != nil {
				return false, err
			}
			return false, err
		}
		defaultKey := gconv.String(key)
		ok, err = g.Redis(groupKey).SetNX(ctx, defaultKey, value)
		if err != nil {
			return ok, err
		}
		if ok && duration > 0 {
			_, err = g.Redis(groupKey).PExpire(ctx, defaultKey, duration.Milliseconds())
			if err != nil {
				return ok, err
			}
			return ok, err
		}
		return ok, err
	}
	return GetAdapterRedis().SetIfNotExist(ctx, key, value, duration)
}

func SetIfNotExistFunc(ctx context.Context, key interface{}, f gcache.Func, duration time.Duration, tag ...interface{}) (ok bool, err error) {
	if !g.IsEmpty(tag) && len(tag) > 0 {
		value, err := f(ctx)
		if err != nil {
			return false, err
		}
		return SetIfNotExist(ctx, key, value, duration, tag...)
	}
	return GetAdapterRedis().SetIfNotExistFunc(ctx, key, f, duration)
}

func Get(ctx context.Context, key interface{}) (*gvar.Var, error) {
	tagCache, err := getTagCache(ctx)
	if err != nil {
		return nil, err
	}
	return tagCache.Get(ctx, gconv.String(key))
}

func GetOrSet(ctx context.Context, key interface{}, value interface{}, duration time.Duration, tag ...interface{}) (result *gvar.Var, err error) {
	if !g.IsEmpty(tag) && len(tag) > 0 {
		result, err = Get(ctx, key)
		if err != nil {
			return
		}
		if result.IsNil() {
			return gvar.New(value), Set(ctx, key, value, duration, tag...)
		}
		return
	}
	return GetAdapterRedis().GetOrSet(ctx, key, value, duration)
}

func GetOrSetFunc(ctx context.Context, key interface{}, f gcache.Func, duration time.Duration, tag ...interface{}) (result *gvar.Var, err error) {
	if !g.IsEmpty(tag) && len(tag) > 0 {
		result, err = Get(ctx, key)
		if err != nil {
			return
		}
		if result.IsNil() {
			value, err := f(ctx)
			if err != nil {
				return nil, err
			}
			if value == nil {
				return nil, nil
			}
			return gvar.New(value), Set(ctx, key, value, duration, tag...)
		}
		return
	}
	return GetAdapterRedis().GetOrSetFunc(ctx, key, f, duration)
}

func SetIfNotExistFuncLock(ctx context.Context, key interface{}, f gcache.Func, duration time.Duration, tag ...interface{}) (ok bool, err error) {
	if !g.IsEmpty(tag) && len(tag) > 0 {
		value, err := f(ctx)
		if err != nil {
			return false, err
		}
		return SetIfNotExist(ctx, key, value, duration, tag...)
	}
	return GetAdapterRedis().SetIfNotExistFuncLock(ctx, key, f, duration)
}

func Contains(ctx context.Context, key interface{}) (bool, error) {
	return GetAdapterRedis().Contains(ctx, key)
}

func Update(ctx context.Context, key interface{}, value interface{}, tag ...interface{}) (oldValue *gvar.Var, exist bool, err error) {
	if !g.IsEmpty(tag) && len(tag) > 0 {
		var (
			v       *gvar.Var
			oldPTTL int64
		)
		defaultKey := gconv.String(key)
		// TTL.
		oldPTTL, err = g.Redis(groupKey).PTTL(ctx, defaultKey) // update ttl -> pttl(millisecond)
		if err != nil {
			return
		}
		if oldPTTL == -2 || oldPTTL == 0 {
			// It does not exist or expired.
			return
		}
		// Check existence.
		v, err = Get(ctx, key)
		if err != nil {
			return
		}
		oldValue = v
		// DEL.
		if value == nil {
			_, err = Remove(ctx, key)
			if err != nil {
				return
			}
			return
		}
		// Update the value.
		if oldPTTL == -1 {
			err = Set(ctx, key, value, 0, tag...)
		} else {
			err = Set(ctx, key, value, time.Duration(oldPTTL/1000)*time.Second, tag...)
		}
		return oldValue, true, err
	}
	return GetAdapterRedis().Update(ctx, key, value)
}

func UpdateExpire(ctx context.Context, key interface{}, duration time.Duration, tag ...interface{}) (oldDuration time.Duration, err error) {
	if !g.IsEmpty(tag) && len(tag) > 0 {
		var (
			v       *gvar.Var
			oldPTTL int64
		)
		// TTL.
		oldPTTL, err = g.Redis(groupKey).PTTL(ctx, gconv.String(key))
		if err != nil {
			return
		}
		if oldPTTL == -2 || oldPTTL == 0 {
			return
		}
		oldDuration = time.Duration(oldPTTL) * time.Millisecond
		if duration < 0 {
			_, err = Remove(ctx, key)
			return
		}
		// Update the expiration.
		if duration > 0 {
			_, err = g.Redis(groupKey).PExpire(ctx, gconv.String(key), duration.Milliseconds())
			if err != nil {
				return
			}
		}
		// No expire.
		if duration == 0 {
			v, err = Get(ctx, key)
			if err != nil {
				return
			}
			err = Set(ctx, key, v.Val(), 0, tag...)
		}
	}
	return GetAdapterRedis().UpdateExpire(ctx, key, duration)
}

func GetExpire(ctx context.Context, key interface{}) (time.Duration, error) {
	return GetAdapterRedis().GetExpire(ctx, key)
}

func Remove(ctx context.Context, key interface{}) (lastValue *gvar.Var, err error) {
	tagCache, err := getTagCache(ctx)
	if err != nil {
		return nil, err
	}

	// 新增数组类型处理逻辑
	if keys, ok := key.([]interface{}); ok {
		// 处理多个 key 的情况
		for _, k := range keys {
			strKey := gconv.String(k)
			if lastValue, err = tagCache.Get(ctx, strKey); err != nil {
				continue
			}
			if err = tagCache.Delete(ctx, strKey); err != nil {
				return lastValue, err
			}
		}
		return lastValue, nil
	}

	// 原有单个 key 处理逻辑
	strKey := gconv.String(key)
	if lastValue, err = tagCache.Get(ctx, strKey); err != nil {
		return nil, err
	}
	err = tagCache.Delete(ctx, strKey)
	return
}

func RemoveByTag(ctx context.Context, tags ...interface{}) (err error) {
	g.Log().Debug(ctx, "RemoveByTag:", tags)
	if !g.IsEmpty(tags) && len(tags) > 0 {
		tagCache, err := getTagCache(ctx)
		if err != nil {
			g.Log().Debug(ctx, "getTagCacheerr:", err)
			return err
		}
		err = tagCache.InvalidateTags(ctx, gconv.Strings(tags))
		if err != nil {
			g.Log().Debug(ctx, "InvalidateTagserr:", err)
			return err
		}
	}
	return
}

func ClearCacheAll(ctx context.Context) error {
	return GetAdapterRedis().Clear(ctx)
}
