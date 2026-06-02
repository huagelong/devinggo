//go:build ignore
// +build ignore

// Package redispubsub
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package redispubsub

import (
	"log"
	"testing"
	"time"

	_ "github.com/gogf/gf/contrib/nosql/redis/v2"

	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/os/gctx"
)

// 使用流程示例: go test -v
func TestRedispubsub(t *testing.T) {
	ctx := gctx.New()

	// *** 此处配置仅测试使用, 正式使用请在manifest/config/config.yml中配置
	redisConfig := gredis.Config{
		Address: "127.0.0.1:6379",
		Db:      1,
		Pass:    "",
	}

	gredis.SetConfig(&redisConfig)
	// ****************************************

	pubsub := New()

	// 订阅频道1和2
	err := pubsub.Subscribe(ctx, "channel1", "channel2")
	if err != nil {
		log.Fatal(err)
	}
	t.Log("[Sub] subscribed to channel1 and channel2")

	// 启动goroutine接收订阅的消息
	go func() {
		for {
			select {
			case msg := <-pubsub.Messages():
				t.Logf("[Sub] received message: %#v, from channel: %s\n", msg.Payload, msg.Channel)
			case <-pubsub.Unsubscribe():
				t.Log("[Sub] goroutine for fetching subscribed messages stopped")
				return
			}
		}
	}()

	// 发布消息到频道1
	err = pubsub.Publish(ctx, "channel1", "hello, channel1")
	if err != nil {
		log.Fatal(err)
	}
	t.Log("[Pub] published message: hello, channel1")

	// 取消订阅频道1
	err = pubsub.UnsubscribeChannel(ctx, "channel1")
	if err != nil {
		log.Fatal(err)
	}
	t.Log("[Sub] unsubscribed from channel1")

	// 发布消息到频道1
	err = pubsub.Publish(ctx, "channel1", "this message won't be received")
	if err != nil {
		log.Fatal(err)
	}
	t.Log("[Pub] published message to channel1: this message won't be received")

	// 发布消息到频道2
	err = pubsub.Publish(ctx, "channel2", "hello, channel2")
	if err != nil {
		log.Fatal(err)
	}
	t.Log("[Pub] published message: hello, channel2")

	// 在原有订阅基础上新增订阅频道3和4
	err = pubsub.Subscribe(ctx, "channel3", "channel4")
	if err != nil {
		log.Fatal(err)
	}
	t.Log("[Sub] subscribed to channel3 and channel4")

	// 发布消息到频道3
	err = pubsub.Publish(ctx, "channel3", "hello, channel3")
	if err != nil {
		log.Fatal(err)
	}
	t.Log("[Pub] published message: hello, channel3")

	// 发布消息到频道2
	err = pubsub.Publish(ctx, "channel2", "hello, channel2")
	if err != nil {
		log.Fatal(err)
	}
	t.Log("[Pub] published message: hello, channel2")

	time.Sleep(100 * time.Microsecond)

	// 停止订阅所有频道
	pubsub.UnsubscribeAll(ctx)

	// 等待一段时间以确保接收到所有消息
	time.Sleep(2 * time.Second)

	// 关闭PubSub实例
	pubsub.Close()

	time.Sleep(100 * time.Millisecond)
}
