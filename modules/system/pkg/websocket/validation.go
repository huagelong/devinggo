// Package websocket
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package websocket

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

// Pusher 命名约定常量
const (
	// MaxChannelNameLength 频道名称最大长度（Pusher 标准）
	MaxChannelNameLength = 200

	// MaxEventNameLength 事件名称最大长度（Pusher 标准）
	MaxEventNameLength = 200

	// MaxChannelsPerTrigger 每次触发的最大频道数（Pusher 标准：最多 100 个频道）
	MaxChannelsPerTrigger = 100

	// MaxEventDataSize 事件数据最大大小（Pusher 标准：10KB）
	MaxEventDataSize = 10 * 1024 // 10KB
)

// 频道名称验证正则表达式
// 只允许字母、数字、下划线(_)、连字符(-)
var channelNameRegex = regexp.MustCompile(`^[a-zA-Z0-9_\-]+$`)

// ValidateChannelName 验证频道名称
//
// Pusher 频道命名规则：
// - 只能包含：字母、数字、下划线(_)、连字符(-)
// - 最大长度：200 字符
// - 不能为空
// - 不能包含空格和其他特殊字符
//
// 参考：https://pusher.com/docs/channels/using_channels/channels#channel-naming-conventions
func ValidateChannelName(channelName string) error {
	if channelName == "" {
		return errors.New("channel name cannot be empty")
	}

	if len(channelName) > MaxChannelNameLength {
		return errors.New("channel name exceeds maximum length of 200 characters")
	}

	// 移除频道前缀以验证基本名称
	baseName := channelName
	if strings.HasPrefix(channelName, "private-encrypted-") {
		baseName = strings.TrimPrefix(channelName, "private-encrypted-")
	} else if strings.HasPrefix(channelName, "presence-") {
		baseName = strings.TrimPrefix(channelName, "presence-")
	} else if strings.HasPrefix(channelName, "private-") {
		baseName = strings.TrimPrefix(channelName, "private-")
	}

	if baseName == "" {
		return errors.New("channel name cannot be only a prefix")
	}

	if !channelNameRegex.MatchString(channelName) {
		return errors.New("channel name can only contain letters, numbers, underscores and hyphens")
	}

	return nil
}

// ValidateEventName 验证事件名称
//
// Pusher 事件命名规则：
// - 最大长度：200 字符
// - 不能为空
// - 客户端触发的事件必须以 "client-" 开头
//
// 参考：https://pusher.com/docs/channels/server_api/http-api#publishing-events
func ValidateEventName(eventName string) error {
	if eventName == "" {
		return errors.New("event name cannot be empty")
	}

	if len(eventName) > MaxEventNameLength {
		return errors.New("event name exceeds maximum length of 200 characters")
	}

	return nil
}

// ValidateChannels 验证频道列表
//
// 验证规则：
// - 频道数量不能超过 100 个（Pusher 标准）
// - 每个频道名称必须有效
func ValidateChannels(channels []string) error {
	if len(channels) == 0 {
		return errors.New("at least one channel is required")
	}

	if len(channels) > MaxChannelsPerTrigger {
		return errors.New("cannot trigger events on more than 100 channels at once")
	}

	for i, channel := range channels {
		if err := ValidateChannelName(channel); err != nil {
			return errors.New("invalid channel name at index " + strconv.Itoa(i) + ": " + err.Error())
		}
	}

	return nil
}

// ValidateEventData 验证事件数据
//
// 验证规则：
// - 数据不能为空
// - 数据大小不能超过 10KB（Pusher 标准）
//
// 参考：https://pusher.com/docs/channels/server_api/http-api#publishing-events
func ValidateEventData(data string) error {
	if data == "" {
		return errors.New("event data cannot be empty")
	}

	if len(data) > MaxEventDataSize {
		return errors.New("event data exceeds maximum size of 10KB")
	}

	return nil
}

// ValidateChannelsForMultiTrigger 验证多频道触发时的加密频道限制
//
// ⚠️ Pusher Server Library Reference Specification 要求：
// "Triggering an event on multiple channels if any of those channels are encrypted
// MUST be prevented in the library API or MUST result in an error"
//
// 验证规则：
// - 当向多个频道触发事件时，如果其中任何一个是加密频道，则返回错误
// - 单个加密频道可以正常触发事件
//
// 参考：https://pusher.com/docs/channels/library_auth_reference/server-library-reference-specification
func ValidateChannelsForMultiTrigger(channels []string) error {
	if len(channels) <= 1 {
		return nil // 单个频道无需检查
	}

	for _, channel := range channels {
		if IsEncryptedChannel(channel) {
			return errors.New("cannot trigger events on multiple channels when any channel is encrypted")
		}
	}

	return nil
}
