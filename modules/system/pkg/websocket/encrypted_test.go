package websocket

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"testing"
)

// TestEncryptMessage 测试消息加密功能
func TestEncryptMessage(t *testing.T) {
	ctx := context.Background()

	// 生成测试用的 shared_secret
	sharedSecret := GenerateSharedSecret()
	t.Logf("Generated shared_secret: %s", sharedSecret)

	// 测试数据
	plaintext := `{"message":"Hello World","amount":123.45}`

	// 加密
	encrypted, err := EncryptMessage(ctx, plaintext, sharedSecret)
	if err != nil {
		t.Fatalf("EncryptMessage failed: %v", err)
	}

	t.Logf("Encrypted: %s", encrypted)

	// 验证加密结果格式
	var result EncryptedMessage
	err = json.Unmarshal([]byte(encrypted), &result)
	if err != nil {
		t.Fatalf("Failed to unmarshal encrypted message: %v", err)
	}

	// 验证字段存在
	if result.Ciphertext == "" {
		t.Error("Ciphertext is empty")
	}
	if result.Nonce == "" {
		t.Error("Nonce is empty")
	}

	// 验证 Base64 编码
	_, err = base64.StdEncoding.DecodeString(result.Ciphertext)
	if err != nil {
		t.Errorf("Ciphertext is not valid Base64: %v", err)
	}

	_, err = base64.StdEncoding.DecodeString(result.Nonce)
	if err != nil {
		t.Errorf("Nonce is not valid Base64: %v", err)
	}

	// 验证 nonce 长度（24 bytes = 32 chars in Base64）
	nonceBytes, _ := base64.StdEncoding.DecodeString(result.Nonce)
	if len(nonceBytes) != 24 {
		t.Errorf("Nonce must be 24 bytes, got %d", len(nonceBytes))
	}

	t.Logf("✅ Encryption successful: ciphertext_len=%d, nonce=%s",
		len(result.Ciphertext), result.Nonce[:8]+"...")
}

// TestEncryptChannelData 测试频道数据加密
func TestEncryptChannelData(t *testing.T) {
	ctx := context.Background()

	// 生成测试用的 shared_secret
	sharedSecret := GenerateSharedSecret()

	// 测试数据（Go结构）
	data := map[string]interface{}{
		"type":      "test",
		"message":   "Hello from server",
		"amount":    12345.67,
		"timestamp": "2025-03-24T10:30:00Z",
	}

	// 加密
	encrypted, err := EncryptChannelData(ctx, data, sharedSecret)
	if err != nil {
		t.Fatalf("EncryptChannelData failed: %v", err)
	}

	t.Logf("Encrypted channel data: %s", encrypted)

	// 验证结果
	var result EncryptedMessage
	err = json.Unmarshal([]byte(encrypted), &result)
	if err != nil {
		t.Fatalf("Failed to unmarshal: %v", err)
	}

	if result.Ciphertext == "" || result.Nonce == "" {
		t.Error("Encrypted data missing required fields")
	}

	t.Log("✅ Channel data encryption successful")
}

// TestEncryptDecryptRoundTrip 测试加密解密往返
func TestEncryptDecryptRoundTrip(t *testing.T) {
	// 注意：这个测试需要和浏览器端的 TweetNaCl 配合才能真正验证
	// 这里只验证加密部分的正确性

	ctx := context.Background()
	sharedSecret := GenerateSharedSecret()

	testCases := []string{
		`{"msg":"test"}`,
		`{"amount":123.45}`,
		`{"data":"Special chars: 中文测试 éàü"}`,
		`{}`,
	}

	for i, tc := range testCases {
		encrypted, err := EncryptMessage(ctx, tc, sharedSecret)
		if err != nil {
			t.Errorf("Case %d: encryption failed: %v", i, err)
			continue
		}

		// 验证格式
		var result EncryptedMessage
		err = json.Unmarshal([]byte(encrypted), &result)
		if err != nil {
			t.Errorf("Case %d: invalid encrypted format: %v", i, err)
			continue
		}

		t.Logf("Case %d: ✅ plaintext_len=%d, ciphertext_len=%d",
			i, len(tc), len(result.Ciphertext))
	}
}

// TestInvalidSharedSecret 测试无效的 shared_secret
func TestInvalidSharedSecret(t *testing.T) {
	ctx := context.Background()
	plaintext := `{"test":"data"}`

	testCases := []struct {
		name         string
		sharedSecret string
		expectError  bool
	}{
		{"empty", "", true},
		{"invalid base64", "not-base64!!!", true},
		{"wrong length", base64.StdEncoding.EncodeToString([]byte("short")), true},
		{"valid 32 bytes", base64.StdEncoding.EncodeToString(make([]byte, 32)), false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := EncryptMessage(ctx, plaintext, tc.sharedSecret)

			if tc.expectError && err == nil {
				t.Errorf("Expected error but got none")
			}
			if !tc.expectError && err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
		})
	}
}

// BenchmarkEncryptMessage 性能基准测试
func BenchmarkEncryptMessage(b *testing.B) {
	ctx := context.Background()
	sharedSecret := GenerateSharedSecret()
	plaintext := `{"message":"Benchmark test","amount":12345.67,"timestamp":"2025-03-24T10:30:00Z"}`

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := EncryptMessage(ctx, plaintext, sharedSecret)
		if err != nil {
			b.Fatalf("Encryption failed: %v", err)
		}
	}
}

// BenchmarkEncryptChannelData 性能基准测试
func BenchmarkEncryptChannelData(b *testing.B) {
	ctx := context.Background()
	sharedSecret := GenerateSharedSecret()
	data := map[string]interface{}{
		"message":   "Benchmark test",
		"amount":    12345.67,
		"timestamp": "2025-03-24T10:30:00Z",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := EncryptChannelData(ctx, data, sharedSecret)
		if err != nil {
			b.Fatalf("Encryption failed: %v", err)
		}
	}
}
