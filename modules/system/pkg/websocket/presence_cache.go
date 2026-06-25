// Package websocket
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package websocket

import (
	"context"
	"sync"
	"time"
)

// PresenceCacheEntry Presence成员列表缓存条目
type PresenceCacheEntry struct {
	Members   map[string]map[string]interface{}
	Count     int
	ExpiredAt time.Time
}

// PresenceCache Presence成员列表缓存（性能优化：减少Redis查询）
type PresenceCache struct {
	cache map[string]*PresenceCacheEntry
	mu    sync.RWMutex
	ttl   time.Duration // 缓存TTL（推荐5秒）
}

var (
	presenceCache     *PresenceCache
	presenceCacheOnce sync.Once
)

// GetPresenceCache 获取全局Presence缓存实例
func GetPresenceCache() *PresenceCache {
	presenceCacheOnce.Do(func() {
		presenceCache = &PresenceCache{
			cache: make(map[string]*PresenceCacheEntry),
			ttl:   5 * time.Second, // 5秒TTL
		}

		// 启动后台清理协程
		go presenceCache.cleanupExpired()
	})
	return presenceCache
}

// GetMembers 获取频道成员列表（带缓存）
func (pc *PresenceCache) GetMembers(ctx context.Context, channel string) (map[string]map[string]interface{}, int, error) {
	// 检查缓存
	pc.mu.RLock()
	entry, exists := pc.cache[channel]
	if exists && time.Now().Before(entry.ExpiredAt) {
		// 缓存命中
		members := entry.Members
		count := entry.Count
		pc.mu.RUnlock()
		return members, count, nil
	}
	pc.mu.RUnlock()

	// 缓存未命中，从Redis查询
	members, err := GetPresenceMembers4Redis(ctx, channel)
	if err != nil {
		return nil, 0, err
	}

	count := len(members)

	// 更新缓存
	pc.mu.Lock()
	pc.cache[channel] = &PresenceCacheEntry{
		Members:   members,
		Count:     count,
		ExpiredAt: time.Now().Add(pc.ttl),
	}
	pc.mu.Unlock()

	return members, count, nil
}

// InvalidateChannel 使指定频道的缓存失效
func (pc *PresenceCache) InvalidateChannel(channel string) {
	pc.mu.Lock()
	delete(pc.cache, channel)
	pc.mu.Unlock()
}

// InvalidateAll 清空所有缓存
func (pc *PresenceCache) InvalidateAll() {
	pc.mu.Lock()
	pc.cache = make(map[string]*PresenceCacheEntry)
	pc.mu.Unlock()
}

// cleanupExpired 后台清理过期缓存
func (pc *PresenceCache) cleanupExpired() {
	ticker := time.NewTicker(10 * time.Second) // 每10秒检查一次
	defer ticker.Stop()

	for range ticker.C {
		now := time.Now()
		pc.mu.Lock()
		for channel, entry := range pc.cache {
			if now.After(entry.ExpiredAt) {
				delete(pc.cache, channel)
			}
		}
		pc.mu.Unlock()
	}
}
