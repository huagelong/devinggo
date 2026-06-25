// Package cache
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package cache

import (
	"bufio"
	"context"
	"strings"
	"time"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/util/gconv"
)

// ──────────────────────────────────────────────
// 包级变量
// ──────────────────────────────────────────────

var (
	_cache                 *gcache.Cache
	_tagCache              *tagCache
	groupKey               = "cache"
	cachePrefixSelectCache = "SelectCache:"
	scanCount              = 100
)

// ──────────────────────────────────────────────
// 包内辅助
// ──────────────────────────────────────────────

func getTagCacheInstance() (*tagCache, error) {
	if _tagCache == nil {
		return nil, gerror.New("tag cache uninitialized, call SetAdapter first")
	}
	return _tagCache, nil
}

func GetRedisClient() *gredis.Redis {
	return g.Redis(groupKey)
}

func getAdapterRedis() gcache.Adapter {
	return gcache.NewAdapterRedis(g.Redis(groupKey))
}

// ──────────────────────────────────────────────
// 初始化
// ──────────────────────────────────────────────

// SetAdapter 初始化缓存适配器和 TagCache 单例。
func SetAdapter(ctx context.Context) error {
	adapter := newCacheAdapter()
	g.DB().GetCache().SetAdapter(adapter)
	_cache = gcache.New()
	_cache.SetAdapter(adapter)
	tc, err := newTagCache(ctx, g.Redis(groupKey))
	if err != nil {
		return gerror.Wrap(err, "failed to initialize tag cache")
	}
	_tagCache = tc
	return nil
}

// GetCache 返回通用缓存实例。
func GetCache() (*gcache.Cache, error) {
	if _cache == nil {
		return nil, gerror.New("cache uninitialized")
	}
	return _cache, nil
}

// ──────────────────────────────────────────────
// 包内实现：写
// ──────────────────────────────────────────────

func set(ctx context.Context, key interface{}, value interface{}, duration time.Duration, tag ...interface{}) error {
	if value == nil || duration < 0 {
		_, err := Remove(ctx, key)
		return err
	}
	var tags []string
	for _, t := range gconv.Strings(tag) {
		if t != "" {
			tags = append(tags, t)
		}
	}
	tci, tciErr := getTagCacheInstance()
	if tciErr != nil {
		return tciErr
	}
	return tci.set(ctx, gconv.String(key), value, duration, tags)
}

func setIfNotExist(ctx context.Context, key interface{}, value interface{}, duration time.Duration, tag ...interface{}) (ok bool, err error) {
	f, isFn := value.(gcache.Func)
	if !isFn {
		f, isFn = value.(func(ctx context.Context) (value interface{}, err error))
	}
	if isFn {
		if value, err = f(ctx); err != nil {
			return false, err
		}
	}
	if duration < 0 || value == nil {
		tci, tciErr := getTagCacheInstance()
		if tciErr != nil {
			return false, tciErr
		}
		return false, tci.delete(ctx, gconv.String(key))
	}
	defaultKey := gconv.String(key)
	ok, err = GetRedisClient().SetNX(ctx, defaultKey, value)
	if err != nil || !ok {
		return ok, err
	}
	var tags []string
	for _, t := range gconv.Strings(tag) {
		if t != "" {
			tags = append(tags, t)
		}
	}
	tci2, tci2Err := getTagCacheInstance()
	if tci2Err != nil {
		return false, tci2Err
	}
	return true, tci2.set(ctx, defaultKey, value, duration, tags)
}

func setIfNotExistFunc(ctx context.Context, key interface{}, f gcache.Func, duration time.Duration, tag ...interface{}) (ok bool, err error) {
	value, err := f(ctx)
	if err != nil {
		return false, err
	}
	return setIfNotExist(ctx, key, value, duration, tag...)
}

func setIfNotExistFuncLock(ctx context.Context, key interface{}, f gcache.Func, duration time.Duration, tag ...interface{}) (ok bool, err error) {
	value, err := f(ctx)
	if err != nil {
		return false, err
	}
	return setIfNotExist(ctx, key, value, duration, tag...)
}

// ──────────────────────────────────────────────
// 包内实现：读 / 查
// ──────────────────────────────────────────────

func getOrSet(ctx context.Context, key interface{}, value interface{}, duration time.Duration, tag ...interface{}) (result *gvar.Var, err error) {
	if result, err = Get(ctx, key); err != nil || !result.IsNil() {
		return
	}
	return gvar.New(value), set(ctx, key, value, duration, tag...)
}

func getOrSetFunc(ctx context.Context, key interface{}, f gcache.Func, duration time.Duration, tag ...interface{}) (result *gvar.Var, err error) {
	if result, err = Get(ctx, key); err != nil || !result.IsNil() {
		return
	}
	value, err := f(ctx)
	if err != nil {
		return nil, err
	}
	if value == nil {
		return nil, nil
	}
	return gvar.New(value), set(ctx, key, value, duration, tag...)
}

// contains 用 EXISTS 命令判断，避免将 redis.Nil 误当错误传播。
func contains(ctx context.Context, key interface{}) (bool, error) {
	n, err := GetRedisClient().Exists(ctx, gconv.String(key))
	if err != nil {
		return false, err
	}
	return n > 0, nil
}

func update(ctx context.Context, key interface{}, value interface{}, tag ...interface{}) (oldValue *gvar.Var, exist bool, err error) {
	defaultKey := gconv.String(key)
	oldPTTL, err := GetRedisClient().PTTL(ctx, defaultKey)
	if err != nil || oldPTTL == -2 || oldPTTL == 0 {
		return
	}
	if oldValue, err = Get(ctx, key); err != nil {
		return
	}
	if value == nil {
		_, err = Remove(ctx, key)
		return
	}
	if oldPTTL == -1 {
		err = set(ctx, key, value, 0, tag...)
	} else {
		err = set(ctx, key, value, time.Duration(oldPTTL/1000)*time.Second, tag...)
	}
	return oldValue, true, err
}

func updateExpire(ctx context.Context, key interface{}, duration time.Duration, tag ...interface{}) (oldDuration time.Duration, err error) {
	oldPTTL, err := GetRedisClient().PTTL(ctx, gconv.String(key))
	if err != nil || oldPTTL == -2 || oldPTTL == 0 {
		return
	}
	oldDuration = time.Duration(oldPTTL) * time.Millisecond
	if duration < 0 {
		_, err = Remove(ctx, key)
		return
	}
	if duration > 0 {
		_, err = GetRedisClient().PExpire(ctx, gconv.String(key), duration.Milliseconds())
		return
	}
	// duration == 0：持久化，重新写入不带 EX 的 tagCache 条目
	v, err := Get(ctx, key)
	if err != nil {
		return
	}
	var tags []string
	for _, t := range gconv.Strings(tag) {
		if t != "" {
			tags = append(tags, t)
		}
	}
	tc, err := getTagCacheInstance()
	if err != nil {
		return oldDuration, err
	}
	err = tc.set(ctx, gconv.String(key), v.Val(), 0, tags)
	return
}

func getExpire(ctx context.Context, key interface{}) (time.Duration, error) {
	return getAdapterRedis().GetExpire(ctx, key)
}

// ──────────────────────────────────────────────
// 公开 API
// ──────────────────────────────────────────────

// Get 获取缓存值。
func Get(ctx context.Context, key interface{}) (*gvar.Var, error) {
	tc, err := getTagCacheInstance()
	if err != nil {
		return nil, err
	}
	return tc.get(ctx, gconv.String(key))
}

// Set 设置缓存。可通过可选参数 tag 指定标签，用于批量清除缓存。
//
// 参数：
//
//	ctx: 上下文
//	key: 缓存键
//	value: 缓存值
//	duration: 过期时间，0 表示永不过期
//	tag: 可选的标签，支持多个
//
// 示例：
//
//	cache.Set(ctx, "user:1", userData, 5*time.Minute)                    // 不带标签
//	cache.Set(ctx, "user:1", userData, 5*time.Minute, "user")            // 单个标签
//	cache.Set(ctx, "user:1", userData, 5*time.Minute, "user", "profile") // 多个标签
func Set(ctx context.Context, key interface{}, value interface{}, duration time.Duration, tag ...interface{}) error {
	return set(ctx, key, value, duration, tag...)
}

// SetIfNotExist 仅在 key 不存在时设置缓存。可通过可选参数 tag 指定标签。
//
// 返回：
//
//	ok: true 表示设置成功，false 表示键已存在
//	err: 错误信息
func SetIfNotExist(ctx context.Context, key interface{}, value interface{}, duration time.Duration, tag ...interface{}) (bool, error) {
	return setIfNotExist(ctx, key, value, duration, tag...)
}

// SetMap 批量设置缓存。可通过可选参数 tag 为所有键指定相同的标签。
func SetMap(ctx context.Context, data map[interface{}]interface{}, duration time.Duration, tag ...interface{}) error {
	if len(data) == 0 {
		return nil
	}
	for k, v := range data {
		if err := set(ctx, k, v, duration, tag...); err != nil {
			return err
		}
	}
	return nil
}

// GetOrSet 获取缓存，如果不存在则设置。可通过可选参数 tag 指定标签。
func GetOrSet(ctx context.Context, key interface{}, value interface{}, duration time.Duration, tag ...interface{}) (*gvar.Var, error) {
	return getOrSet(ctx, key, value, duration, tag...)
}

// GetOrSetFunc 使用函数获取或设置缓存。可通过可选参数 tag 指定标签。
//
// 仅在缓存不存在时才会执行函数 f，避免不必要的计算。
func GetOrSetFunc(ctx context.Context, key interface{}, f gcache.Func, duration time.Duration, tag ...interface{}) (*gvar.Var, error) {
	return getOrSetFunc(ctx, key, f, duration, tag...)
}

// Update 更新缓存值，保持原有过期时间。可通过可选参数 tag 更新标签。
//
// 返回：
//
//	oldValue: 更新前的旧值
//	exist: true 表示键存在并已更新，false 表示键不存在
//	err: 错误信息
func Update(ctx context.Context, key interface{}, value interface{}, tag ...interface{}) (*gvar.Var, bool, error) {
	return update(ctx, key, value, tag...)
}

// UpdateExpire 更新缓存过期时间。可通过可选参数 tag 更新标签。
//
// 返回：
//
//	oldDuration: 更新前的剩余过期时间
//	err: 错误信息
func UpdateExpire(ctx context.Context, key interface{}, duration time.Duration, tag ...interface{}) (time.Duration, error) {
	return updateExpire(ctx, key, duration, tag...)
}

// GetExpire 获取缓存的剩余过期时间。
func GetExpire(ctx context.Context, key interface{}) (time.Duration, error) {
	return getExpire(ctx, key)
}

// Contains 判断缓存键是否存在。
func Contains(ctx context.Context, key interface{}) (bool, error) {
	return contains(ctx, key)
}

// Remove 删除一个或多个缓存 key，返回最后一个被删除 key 的旧值。
func Remove(ctx context.Context, key interface{}) (lastValue *gvar.Var, err error) {
	tc, err := getTagCacheInstance()
	if err != nil {
		return nil, err
	}
	if keys, ok := key.([]interface{}); ok {
		for _, k := range keys {
			strKey := gconv.String(k)
			if lastValue, err = tc.get(ctx, strKey); err != nil {
				continue
			}
			if err = tc.delete(ctx, strKey); err != nil {
				return lastValue, err
			}
		}
		return lastValue, nil
	}
	strKey := gconv.String(key)
	if lastValue, err = tc.get(ctx, strKey); err != nil {
		return nil, err
	}
	err = tc.delete(ctx, strKey)
	return
}

// RemoveByTag 按 tag 批量清除缓存。
func RemoveByTag(ctx context.Context, tags ...interface{}) (err error) {
	g.Log().Debug(ctx, "RemoveByTag:", tags)
	if len(tags) > 0 {
		tc, err := getTagCacheInstance()
		if err != nil {
			return err
		}
		if err = tc.invalidateTags(ctx, gconv.Strings(tags)); err != nil {
			g.Log().Debug(ctx, "InvalidateTags err:", err)
		}
	}
	return
}

// ClearByTable 清除指定表关联的所有缓存。
func ClearByTable(ctx context.Context, table string) error {
	return RemoveByTag(ctx, table)
}

// ClearCacheAll 清空所有缓存。
func ClearCacheAll(ctx context.Context) error {
	return getAdapterRedis().Clear(ctx)
}

// GetKeys 返回所有缓存 key 列表。
func GetKeys(ctx context.Context) (keys []string, err error) {
	keys = make([]string, 0)
	iterator := uint64(0)
	for {
		var listKeys []string
		iterator, listKeys, err = GetRedisClient().Scan(ctx, iterator, gredis.ScanOption{
			Match: "*",
			Count: scanCount,
		})
		if err != nil {
			g.Log().Error(ctx, "Scan error:", err)
			break
		}
		for _, k := range listKeys {
			if !strings.HasPrefix(k, "tags:") && !strings.HasPrefix(k, "item_tags:") {
				keys = append(keys, k)
			}
		}
		if iterator == 0 {
			break
		}
	}
	return
}

// GetInfo 返回 Redis INFO 信息，按节分组。
func GetInfo(ctx context.Context) (map[string]map[string]interface{}, error) {
	info, err := GetRedisClient().Do(ctx, "INFO")
	if err != nil {
		return nil, err
	}
	sections := make(map[string][]map[string]interface{})
	scanner := bufio.NewScanner(strings.NewReader(info.String()))
	var section string
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "#") {
			section = strings.TrimSpace(strings.SplitN(line, "#", 2)[1])
			sections[section] = nil
			continue
		}
		if line == "" {
			continue
		}
		kv := strings.SplitN(line, ":", 2)
		if len(kv) != 2 {
			continue
		}
		m := make(map[string]interface{})
		if strings.Contains(kv[1], ",") {
			sub := make(map[string]interface{})
			for _, s := range strings.Split(kv[1], ",") {
				if skv := strings.SplitN(s, "=", 2); len(skv) == 2 {
					sub[skv[0]] = skv[1]
				}
			}
			m[kv[0]] = sub
		} else {
			m[kv[0]] = kv[1]
		}
		sections[section] = append(sections[section], m)
	}
	res := make(map[string]map[string]interface{}, len(sections))
	for k, vList := range sections {
		flat := make(map[string]interface{})
		for _, v := range vList {
			for k1, v1 := range v {
				flat[k1] = v1
			}
		}
		res[k] = flat
	}
	return res, nil
}
