// Package websocket
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package websocket

import (
	"sync"
	"time"
)

// RateLimiter Client Events速率限制器（v8.3.0要求：10条/秒）
type RateLimiter struct {
	mu      sync.Mutex
	buckets map[string]*tokenBucket
}

// tokenBucket 令牌桶
type tokenBucket struct {
	tokens     int       // 当前令牌数
	lastRefill time.Time // 上次补充令牌时间
	maxTokens  int       // 最大令牌数
	refillRate int       // 每秒补充令牌数
}

// 全局速率限制器
var globalRateLimiter = NewRateLimiter()

// NewRateLimiter 创建速率限制器
func NewRateLimiter() *RateLimiter {
	rl := &RateLimiter{
		buckets: make(map[string]*tokenBucket),
	}

	// 启动定期清理过期bucket（避免内存泄漏）
	go rl.cleanupExpiredBuckets()

	return rl
}

// AllowClientEvent 检查是否允许发送Client Event（10条/秒）
func (rl *RateLimiter) AllowClientEvent(socketID string) bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	bucket, exists := rl.buckets[socketID]
	if !exists {
		bucket = &tokenBucket{
			tokens:     10,
			lastRefill: time.Now(),
			maxTokens:  10,
			refillRate: 10, // 每秒10条
		}
		rl.buckets[socketID] = bucket
	}

	// 补充令牌（Token Bucket算法）
	now := time.Now()
	elapsed := now.Sub(bucket.lastRefill).Seconds()
	if elapsed >= 1.0 {
		// 每秒补充令牌
		bucket.tokens = bucket.maxTokens
		bucket.lastRefill = now
	}

	// 检查是否有可用令牌
	if bucket.tokens > 0 {
		bucket.tokens--
		return true
	}

	return false
}

// RemoveBucket 移除指定socket_id的bucket（连接断开时调用）
func (rl *RateLimiter) RemoveBucket(socketID string) {
	rl.mu.Lock()
	defer rl.mu.Unlock()
	delete(rl.buckets, socketID)
}

// cleanupExpiredBuckets 定期清理超过5分钟未使用的bucket
func (rl *RateLimiter) cleanupExpiredBuckets() {
	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()

	for range ticker.C {
		rl.mu.Lock()
		now := time.Now()
		for socketID, bucket := range rl.buckets {
			if now.Sub(bucket.lastRefill) > 5*time.Minute {
				delete(rl.buckets, socketID)
			}
		}
		rl.mu.Unlock()
	}
}

// GetRateLimiter 获取全局速率限制器实例
func GetRateLimiter() *RateLimiter {
	return globalRateLimiter
}
