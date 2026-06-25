// Package websocket
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package websocket

import (
	"fmt"
	"strings"
	"testing"
)

// TestValidateChannelName 测试频道名称验证
func TestValidateChannelName(t *testing.T) {
	tests := []struct {
		name        string
		channelName string
		expectError bool
	}{
		// 有效的频道名称
		{
			name:        "简单的公开频道",
			channelName: "my-channel",
			expectError: false,
		},
		{
			name:        "包含数字的频道",
			channelName: "channel-123",
			expectError: false,
		},
		{
			name:        "包含下划线的频道",
			channelName: "my_channel_name",
			expectError: false,
		},
		{
			name:        "私有频道",
			channelName: "private-my-channel",
			expectError: false,
		},
		{
			name:        "在线状态频道",
			channelName: "presence-chatroom",
			expectError: false,
		},
		{
			name:        "加密频道",
			channelName: "private-encrypted-secret-channel",
			expectError: false,
		},
		{
			name:        "最大长度频道名（200字符）",
			channelName: strings.Repeat("a", 200),
			expectError: false,
		},

		// 无效的频道名称
		{
			name:        "空字符串",
			channelName: "",
			expectError: true,
		},
		{
			name:        "只有前缀",
			channelName: "private-",
			expectError: true,
		},
		{
			name:        "只有presence前缀",
			channelName: "presence-",
			expectError: true,
		},
		{
			name:        "只有加密前缀",
			channelName: "private-encrypted-",
			expectError: true,
		},
		{
			name:        "包含空格",
			channelName: "my channel",
			expectError: true,
		},
		{
			name:        "包含特殊字符@",
			channelName: "my@channel",
			expectError: true,
		},
		{
			name:        "包含特殊字符#",
			channelName: "my#channel",
			expectError: true,
		},
		{
			name:        "包含特殊字符!",
			channelName: "my!channel",
			expectError: true,
		},
		{
			name:        "超过最大长度（201字符）",
			channelName: strings.Repeat("a", 201),
			expectError: true,
		},
		{
			name:        "包含中文字符",
			channelName: "频道名称",
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateChannelName(tt.channelName)
			if tt.expectError && err == nil {
				t.Errorf("ValidateChannelName(%q) expected error but got nil", tt.channelName)
			}
			if !tt.expectError && err != nil {
				t.Errorf("ValidateChannelName(%q) unexpected error: %v", tt.channelName, err)
			}
		})
	}
}

// TestValidateEventName 测试事件名称验证
func TestValidateEventName(t *testing.T) {
	tests := []struct {
		name        string
		eventName   string
		expectError bool
	}{
		// 有效的事件名称
		{
			name:        "简单事件名",
			eventName:   "my-event",
			expectError: false,
		},
		{
			name:        "客户端事件",
			eventName:   "client-message",
			expectError: false,
		},
		{
			name:        "包含点号",
			eventName:   "user.login",
			expectError: false,
		},
		{
			name:        "最大长度事件名（200字符）",
			eventName:   strings.Repeat("a", 200),
			expectError: false,
		},

		// 无效的事件名称
		{
			name:        "空字符串",
			eventName:   "",
			expectError: true,
		},
		{
			name:        "超过最大长度（201字符）",
			eventName:   strings.Repeat("a", 201),
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateEventName(tt.eventName)
			if tt.expectError && err == nil {
				t.Errorf("ValidateEventName(%q) expected error but got nil", tt.eventName)
			}
			if !tt.expectError && err != nil {
				t.Errorf("ValidateEventName(%q) unexpected error: %v", tt.eventName, err)
			}
		})
	}
}

// TestValidateChannels 测试频道列表验证
func TestValidateChannels(t *testing.T) {
	tests := []struct {
		name        string
		channels    []string
		expectError bool
		errorMsg    string
	}{
		// 有效的频道列表
		{
			name:        "单个频道",
			channels:    []string{"my-channel"},
			expectError: false,
		},
		{
			name:        "多个频道（10个）",
			channels:    []string{"ch1", "ch2", "ch3", "ch4", "ch5", "ch6", "ch7", "ch8", "ch9", "ch10"},
			expectError: false,
		},
		{
			name:        "混合类型的频道",
			channels:    []string{"public-ch", "private-ch", "presence-room", "private-encrypted-secret"},
			expectError: false,
		},

		// 无效的频道列表
		{
			name:        "空列表",
			channels:    []string{},
			expectError: true,
			errorMsg:    "at least one channel is required",
		},
		{
			name:        "超过100个频道",
			channels:    generateChannels(101),
			expectError: true,
			errorMsg:    "cannot trigger events on more than 100 channels at once",
		},
		{
			name:        "最多100个频道",
			channels:    generateChannels(100),
			expectError: false,
		},
		{
			name:        "包含无效频道名",
			channels:    []string{"valid-channel", "invalid channel"},
			expectError: true,
		},
		{
			name:        "包含空频道名",
			channels:    []string{"valid-channel", ""},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateChannels(tt.channels)
			if tt.expectError && err == nil {
				t.Errorf("ValidateChannels(%v) expected error but got nil", tt.channels)
			}
			if !tt.expectError && err != nil {
				t.Errorf("ValidateChannels(%v) unexpected error: %v", tt.channels, err)
			}
			if tt.expectError && err != nil && tt.errorMsg != "" {
				if !strings.Contains(err.Error(), tt.errorMsg) {
					t.Errorf("ValidateChannels(%v) error = %v, want error containing %q", tt.channels, err, tt.errorMsg)
				}
			}
		})
	}
}

// TestValidateChannelsForMultiTrigger 测试多频道触发时的加密频道限制
func TestValidateChannelsForMultiTrigger(t *testing.T) {
	tests := []struct {
		name        string
		channels    []string
		expectError bool
		errorMsg    string
	}{
		// 有效的场景
		{
			name:        "单个加密频道（允许）",
			channels:    []string{"private-encrypted-secret"},
			expectError: false,
		},
		{
			name:        "单个普通频道（允许）",
			channels:    []string{"my-channel"},
			expectError: false,
		},
		{
			name:        "多个普通频道（允许）",
			channels:    []string{"ch1", "ch2", "ch3"},
			expectError: false,
		},
		{
			name:        "多个私有频道（允许）",
			channels:    []string{"private-ch1", "private-ch2"},
			expectError: false,
		},
		{
			name:        "多个presence频道（允许）",
			channels:    []string{"presence-room1", "presence-room2"},
			expectError: false,
		},

		// 无效的场景
		{
			name:        "多个频道中包含加密频道（禁止）",
			channels:    []string{"my-channel", "private-encrypted-secret"},
			expectError: true,
			errorMsg:    "cannot trigger events on multiple channels when any channel is encrypted",
		},
		{
			name:        "加密频道+普通频道（禁止）",
			channels:    []string{"private-encrypted-ch1", "public-ch2"},
			expectError: true,
			errorMsg:    "cannot trigger events on multiple channels when any channel is encrypted",
		},
		{
			name:        "多个加密频道（禁止）",
			channels:    []string{"private-encrypted-ch1", "private-encrypted-ch2"},
			expectError: true,
			errorMsg:    "cannot trigger events on multiple channels when any channel is encrypted",
		},
		{
			name:        "加密频道+私有频道（禁止）",
			channels:    []string{"private-encrypted-ch1", "private-ch2"},
			expectError: true,
			errorMsg:    "cannot trigger events on multiple channels when any channel is encrypted",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateChannelsForMultiTrigger(tt.channels)
			if tt.expectError && err == nil {
				t.Errorf("ValidateChannelsForMultiTrigger(%v) expected error but got nil", tt.channels)
			}
			if !tt.expectError && err != nil {
				t.Errorf("ValidateChannelsForMultiTrigger(%v) unexpected error: %v", tt.channels, err)
			}
			if tt.expectError && err != nil && tt.errorMsg != "" {
				if !strings.Contains(err.Error(), tt.errorMsg) {
					t.Errorf("ValidateChannelsForMultiTrigger(%v) error = %v, want error containing %q", tt.channels, err, tt.errorMsg)
				}
			}
		})
	}
}

// BenchmarkValidateChannelName 频道名称验证性能测试
func BenchmarkValidateChannelName(b *testing.B) {
	channelName := "private-encrypted-my-channel-123"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ValidateChannelName(channelName)
	}
}

// BenchmarkValidateEventName 事件名称验证性能测试
func BenchmarkValidateEventName(b *testing.B) {
	eventName := "client-message-sent"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ValidateEventName(eventName)
	}
}

// generateChannels 生成指定数量的频道名称（用于测试）
func generateChannels(count int) []string {
	channels := make([]string, count)
	for i := 0; i < count; i++ {
		channels[i] = fmt.Sprintf("channel-%d", i)
	}
	return channels
}

// BenchmarkValidateChannels 频道列表验证性能测试
func BenchmarkValidateChannels(b *testing.B) {
	channels := []string{"ch1", "ch2", "ch3", "ch4", "ch5"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ValidateChannels(channels)
	}
}
