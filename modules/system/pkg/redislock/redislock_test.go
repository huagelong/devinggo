//go:build ignore
// +build ignore

// Package redislock
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package redislock

import (
	"testing"
	"time"

	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/os/gctx"
)

// TestRedisLock go test -v 测试redis分布式锁流程
func TestRedisLock(t *testing.T) {
	ctx := gctx.New()

	// *** 此处配置仅测试使用, 正式使用请在manifest/config/config.yml中配置
	redisConfig := gredis.Config{
		Address: "127.0.0.1:6379",
		Db:      1,
		Pass:    "",
	}
	gredis.SetConfig(&redisConfig)
	// ****************************************

	expiration := WithExpiration(5 * time.Second)

	lock := New("test_lock", expiration)

	if ok := lock.Acquire(ctx); !ok {
		t.Logf("lock acquire failed, lock name: %s", lock.key)
		return
	}

	t.Logf("lock acquire success, lock name: %s", lock.key)

	go func() {
		time.Sleep(3 * time.Second)
		lock.Release(ctx)
		t.Logf("lock release success, lock name: %s", lock.key)
	}()

	for {
		time.Sleep(1 * time.Second)

		if ok := lock.Acquire(ctx); !ok {
			t.Logf("lock acquire failed, lock name: %s", lock.key)
			continue
		}

		t.Logf("lock acquire success, lock name: %s", lock.key)
		break
	}
}
