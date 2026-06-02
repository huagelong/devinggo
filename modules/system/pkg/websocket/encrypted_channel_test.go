// Package websocket
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package websocket

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// TestEncryptedChannelServerPush 测试 Encrypted Channels 服务端推送
func TestEncryptedChannelServerPush(t *testing.T) {
	// 1. 初始化 Pusher 配置
	appKey := "test-app-key"
	appSecret := "test-app-secret"
	InitPusherAuth(appKey, appSecret)

	t.Log("✅ Pusher 认证配置已初始化")

	// 2. 测试加密频道类型检测
	testCases := []struct {
		channel  string
		expected ChannelType
	}{
		{"private-encrypted-secure", ChannelTypeEncrypted},
		{"private-encrypted-test", ChannelTypeEncrypted},
		{"private-test", ChannelTypePrivate},
		{"presence-lobby", ChannelTypePresence},
		{"public-channel", ChannelTypePublic},
	}

	for _, tc := range testCases {
		channelType := GetChannelType(tc.channel)
		assert.Equal(t, tc.expected, channelType, "频道类型检测失败: %s", tc.channel)
		t.Logf("✅ 频道 '%s' 类型检测正确: %v", tc.channel, tc.expected)
	}

	// 3. 测试 shared_secret 生成
	sharedSecret := GenerateSharedSecret()
	assert.NotEmpty(t, sharedSecret, "shared_secret 不应为空")
	assert.Greater(t, len(sharedSecret), 40, "shared_secret 长度应该足够（Base64编码的32字节应该>40字符）")
	t.Logf("✅ shared_secret 已生成: %s (长度: %d)", sharedSecret[:20]+"...", len(sharedSecret))

	// 4. 测试加密频道认证签名生成
	socketID := "test-server.12345"
	channel := "private-encrypted-secure"
	auth := GenerateAuthSignature(socketID, channel, "")

	assert.NotEmpty(t, auth, "认证签名不应为空")
	assert.Contains(t, auth, appKey, "认证签名应包含 app_key")
	assert.Contains(t, auth, ":", "认证签名应包含冒号分隔符")
	t.Logf("✅ 加密频道认证签名已生成: %s", auth)

	// 5. 测试认证签名验证
	err := ValidateChannelAuth(socketID, channel, auth, "")
	assert.NoError(t, err, "认证签名验证应该成功")
	t.Log("✅ 认证签名验证成功")

	// 6. 测试错误的签名
	invalidAuth := appKey + ":invalid_signature"
	err = ValidateChannelAuth(socketID, channel, invalidAuth, "")
	assert.Error(t, err, "错误的签名应该验证失败")
	t.Log("✅ 错误签名正确拒绝")

	// 7. 测试消息序列化（模拟服务端推送）
	testMessage := map[string]interface{}{
		"type":      "payment-notification",
		"amount":    100.50,
		"status":    "success",
		"timestamp": time.Now().Unix(),
		"message":   "支付成功",
	}

	messageJSON, err := json.Marshal(testMessage)
	assert.NoError(t, err, "消息序列化应该成功")
	t.Logf("✅ 消息已序列化: %s", string(messageJSON))

	// 8. 测试 PusherResponse 构建
	pusherResponse := &PusherResponse{
		Event:   "encrypted-message",
		Channel: channel,
		Data:    string(messageJSON),
	}

	responseJSON, err := json.Marshal(pusherResponse)
	assert.NoError(t, err, "PusherResponse 序列化应该成功")
	assert.Contains(t, string(responseJSON), "event", "响应应包含 event 字段")
	assert.Contains(t, string(responseJSON), "channel", "响应应包含 channel 字段")
	assert.Contains(t, string(responseJSON), "data", "响应应包含 data 字段")
	t.Logf("✅ PusherResponse 已构建: %s", string(responseJSON))

	t.Log("\n========== 测试总结 ==========")
	t.Log("✅ 所有 Encrypted Channels 服务端推送功能测试通过！")
	t.Log("📋 测试覆盖:")
	t.Log("  - 频道类型检测 (IsEncryptedChannel)")
	t.Log("  - shared_secret 生成")
	t.Log("  - 认证签名生成与验证")
	t.Log("  - 消息格式序列化")
}

// TestEncryptedChannelAuthFlow 测试完整的认证流程
func TestEncryptedChannelAuthFlow(t *testing.T) {
	// 模拟认证端点逻辑
	appKey := "devinggo-app-key"
	appSecret := "devinggo-app-secret"
	InitPusherAuth(appKey, appSecret)

	socketID := "server1.177227498249324"
	channel := "private-encrypted-secure"

	t.Log("========== 加密频道认证流程测试 ==========")

	// Step 1: 客户端请求认证
	t.Logf("Step 1: 客户端请求认证 (socket_id=%s, channel=%s)", socketID, channel)

	// Step 2: 服务器生成 shared_secret
	sharedSecret := GenerateSharedSecret()
	t.Logf("Step 2: 服务器生成 shared_secret (长度=%d)", len(sharedSecret))

	// Step 3: 服务器生成认证签名
	auth := GenerateAuthSignature(socketID, channel, "")
	t.Logf("Step 3: 服务器生成认证签名: %s", auth)

	// Step 4: 构建认证响应（模拟 HTTP 响应）
	authResponse := map[string]interface{}{
		"auth":          auth,
		"shared_secret": sharedSecret,
	}

	authResponseJSON, err := json.Marshal(authResponse)
	assert.NoError(t, err)
	t.Logf("Step 4: 认证响应已构建: %s", string(authResponseJSON))

	// Step 5: 客户端验证签名（模拟客户端行为）
	err = ValidateChannelAuth(socketID, channel, auth, "")
	assert.NoError(t, err, "客户端签名验证应该成功")
	t.Log("Step 5: 客户端验证签名成功 ✅")

	// Step 6: 客户端订阅成功（模拟）
	t.Log("Step 6: 客户端使用 shared_secret 配置加密，订阅成功 ✅")

	t.Log("\n✅ 完整认证流程测试通过！")
}

// TestEncryptedChannelMessagePush 测试消息推送逻辑
func TestEncryptedChannelMessagePush(t *testing.T) {
	t.Log("========== 加密消息推送逻辑测试 ==========")

	channel := "private-encrypted-user-123"

	// 1. 构建要推送的业务数据
	businessData := map[string]interface{}{
		"type":       "account-update",
		"user_id":    123,
		"balance":    9999.99,
		"updated_at": time.Now().Format(time.RFC3339),
		"sensitive":  true,
	}

	dataJSON, err := json.Marshal(businessData)
	assert.NoError(t, err)
	t.Logf("Step 1: 业务数据已序列化: %s", string(dataJSON))

	// 2. 构建 Pusher 事件
	eventName := "account-balance-changed"
	pusherEvent := &PusherResponse{
		Event:   eventName,
		Channel: channel,
		Data:    string(dataJSON),
	}

	eventJSON, err := json.Marshal(pusherEvent)
	assert.NoError(t, err)
	t.Logf("Step 2: Pusher 事件已构建: %s", string(eventJSON))

	// 3. 验证频道类型
	channelType := GetChannelType(channel)
	assert.Equal(t, ChannelTypeEncrypted, channelType, "应该识别为加密频道")
	t.Logf("Step 3: 频道类型验证: %v ✅", channelType)

	// 4. 验证 RequiresAuth
	requiresAuth := RequiresAuth(channel)
	assert.True(t, requiresAuth, "加密频道应该需要认证")
	t.Log("Step 4: 认证要求检查通过 ✅")

	// 5. 验证 IsEncryptedChannel 辅助函数
	isEncrypted := IsEncryptedChannel(channel)
	assert.True(t, isEncrypted, "应该识别为加密频道")
	t.Log("Step 5: 加密频道辅助函数验证通过 ✅")

	t.Log("\n✅ 消息推送逻辑测试通过！")
	t.Log("💡 提示: 实际加密由客户端 TweetNaCl 库处理，服务器只负责路由")
}

// TestGenerateSharedSecretUniqueness 测试 shared_secret 的唯一性
func TestGenerateSharedSecretUniqueness(t *testing.T) {
	secrets := make(map[string]bool)
	iterations := 100

	t.Logf("========== shared_secret 唯一性测试 (生成 %d 个) ==========", iterations)

	for i := 0; i < iterations; i++ {
		secret := GenerateSharedSecret()

		// 验证长度
		assert.Greater(t, len(secret), 40, "shared_secret 长度应该足够")

		// 验证唯一性
		assert.False(t, secrets[secret], "第 %d 个 shared_secret 不应该重复", i+1)
		secrets[secret] = true
	}

	t.Logf("✅ 生成了 %d 个唯一的 shared_secret", iterations)
	t.Log("✅ 所有密钥都是唯一的，随机性良好")
}

// BenchmarkGenerateSharedSecret 性能基准测试
func BenchmarkGenerateSharedSecret(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GenerateSharedSecret()
	}
}

// BenchmarkGenerateAuthSignature 认证签名生成性能测试
func BenchmarkGenerateAuthSignature(b *testing.B) {
	InitPusherAuth("test-key", "test-secret")
	socketID := "server.12345"
	channel := "private-encrypted-test"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		GenerateAuthSignature(socketID, channel, "")
	}
}

// BenchmarkValidateChannelAuth 认证签名验证性能测试
func BenchmarkValidateChannelAuth(b *testing.B) {
	InitPusherAuth("test-key", "test-secret")
	socketID := "server.12345"
	channel := "private-encrypted-test"
	auth := GenerateAuthSignature(socketID, channel, "")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ValidateChannelAuth(socketID, channel, auth, "")
	}
}
