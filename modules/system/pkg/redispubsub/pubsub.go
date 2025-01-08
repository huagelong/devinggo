// Package redispubsub
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package redispubsub

import (
	"context"
	"devinggo/modules/system/pkg/utils"
	"sync"

	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"
)

type PubSub struct {
	client    *gredis.Redis
	subscribe gredis.Conn
	channels  []string             // 频道列表
	messages  chan *gredis.Message // 订阅的消息放到该通道
	unsub     chan struct{}        // 订阅取消信号通道
	lock      sync.Mutex
	logger    *glog.Logger
}

type Options struct {
	Group      string // redis group name
	LoggerName string // logger name
}

type Option func(*Options)

func New(opts ...Option) *PubSub {
	options := Options{
		Group:      "default", // redis 默认分组
		LoggerName: "",        // 主程序默认日志
	}

	for _, opt := range opts {
		opt(&options)
	}

	return &PubSub{
		client:    g.Redis(options.Group),
		subscribe: nil,
		channels:  nil,
		messages:  make(chan *gredis.Message),
		unsub:     make(chan struct{}),
		logger:    g.Log(options.LoggerName),
	}
}

// Subscribe 订阅频道, 支持订阅1个或多个频道.
func (ps *PubSub) Subscribe(ctx context.Context, channel string, channels ...string) error {
	ps.lock.Lock()
	defer ps.lock.Unlock()

	channels = append(channels, channel)

	// 已经订阅过一些频道了, 继续订阅其他频道的场景.
	if ps.subscribe != nil {
		var (
			newChannels    []interface{}
			newChannelsStr []string
		)
		for _, c := range channels {
			if !contains(ps.channels, c) {
				newChannels = append(newChannels, c)
				newChannelsStr = append(newChannelsStr, c)
			}
		}

		if _, err := ps.subscribe.Do(ctx, "SUBSCRIBE", newChannels...); err != nil {
			return err
		}

		ps.channels = append(ps.channels, newChannelsStr...)

		return nil
	}

	// 以下是首次订阅的场景.
	ps.channels = channels

	conn, err := ps.client.Conn(ctx)
	if err != nil {
		return err
	}

	if _, err := conn.Subscribe(ctx, ps.channels[0], ps.channels[1:]...); err != nil {
		return err
	}

	ps.subscribe = conn

	// 启动goroutine接收订阅的消息
	utils.SafeGo(ctx, func(ctx context.Context) {
		func() {
			defer conn.Close(ctx)
			for {
				res, err := conn.Receive(ctx)
				if err != nil {
					ps.logger.Errorf(ctx, "[PubSub] receive message error: %v", err)
					continue
				}

				switch msg := res.Val().(type) {
				case *gredis.Message:
					ps.messages <- msg
				case *gredis.Subscription: // 处理订阅和取消订阅事件
					if msg.Kind == "subscribe" {
						ps.logger.Infof(ctx, "[PubSub] subscribed to channel: %s, now channels: %v\n", msg.Channel, ps.channels)
					} else if msg.Kind == "unsubscribe" {
						ps.logger.Infof(ctx, "[PubSub] unsubscribed from channel: %s, now channels: %v\n", msg.Channel, ps.channels)
					}
				}
			}
		}()
	})

	return nil
}

// Publish 发布消息到频道.
func (ps *PubSub) Publish(ctx context.Context, channel, message string) error {
	_, err := ps.client.Do(ctx, "PUBLISH", channel, message)
	if err != nil {
		return err
	}

	return nil
}

// UnsubscribeChannel 取消订阅指定频道.
func (ps *PubSub) UnsubscribeChannel(ctx context.Context, channel string) error {
	_, err := ps.subscribe.Do(ctx, "UNSUBSCRIBE", channel)
	if err != nil {
		return err
	}

	ps.delSubChannels(channel)

	return nil
}

// UnsubscribeAll 停止订阅所有频道.
func (ps *PubSub) UnsubscribeAll(ctx context.Context) error {
	for _, c := range ps.channels {
		err := ps.UnsubscribeChannel(ctx, c)
		if err != nil {
			ps.logger.Errorf(ctx, "[PubSub] unsubscribe channel error: %v", err)
		}
	}

	return nil
}

// Messages 返回订阅的消息通道.
func (ps *PubSub) Messages() <-chan *gredis.Message {
	return ps.messages
}

// Unsubscribe 返回取消订阅的信号通道.
func (ps *PubSub) Unsubscribe() <-chan struct{} {
	return ps.unsub
}

// Close 关闭PubSub长连接实例.
func (ps *PubSub) Close() {
	close(ps.unsub)
	close(ps.messages)
}

// 从订阅的频道列表中删除指定的频道.
func (ps *PubSub) delSubChannels(channel string) {
	var newChannels []string
	for _, c := range ps.channels {
		if c != channel {
			newChannels = append(newChannels, c)
		}
	}
	ps.channels = newChannels
}

// WithRedisGroup 设置使用配置的哪组redis服务器(或集群).
func WithRedisGroup(group string) Option {
	return func(o *Options) {
		o.Group = group
	}
}

// WithLoggerName 使用指定的logger, 如果存在子项目logger, 则需要使用此函数提供指定的logger名称.
// 默认是gf框架配置文件中配置的default logger(apiserver).
func WithLoggerName(loggerName string) Option {
	return func(o *Options) {
		o.LoggerName = loggerName
	}
}

func contains(items []string, item string) bool {
	for _, v := range items {
		if v == item {
			return true
		}
	}
	return false
}
