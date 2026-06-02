package websocket

import (
	"encoding/json"
	"testing"
)

// TestGenerateUserAuthSignature 测试用户认证签名生成
func TestGenerateUserAuthSignature(t *testing.T) {
	// 初始化配置
	InitPusherAuth("test-app-key", "test-app-secret")

	socketID := "test-socket.123456"
	userData := map[string]interface{}{
		"id":     "1234",
		"name":   "Test User",
		"email":  "test@example.com",
		"avatar": "https://example.com/avatar.png",
	}

	// 生成签名
	auth, err := GenerateUserAuthSignature(socketID, userData)
	if err != nil {
		t.Fatalf("Failed to generate user auth signature: %v", err)
	}

	// 验证签名格式：app_key:signature
	if len(auth) == 0 {
		t.Error("Auth signature is empty")
	}

	// 验证包含 app_key
	expectedPrefix := "test-app-key:"
	if len(auth) <= len(expectedPrefix) || auth[:len(expectedPrefix)] != expectedPrefix {
		t.Errorf("Auth signature should start with '%s', got: %s", expectedPrefix, auth)
	}

	// 验证签名部分是十六进制（64 字符 for SHA256）
	signature := auth[len(expectedPrefix):]
	if len(signature) != 64 {
		t.Errorf("Signature should be 64 characters (SHA256 hex), got: %d", len(signature))
	}

	t.Logf("✅ User auth signature generated: %s", auth)
}

// TestGenerateUserAuthSignature_Validation 测试输入验证
func TestGenerateUserAuthSignature_Validation(t *testing.T) {
	InitPusherAuth("test-app-key", "test-app-secret")

	socketID := "test-socket.123456"

	testCases := []struct {
		name        string
		userData    map[string]interface{}
		expectError bool
		errorMsg    string
	}{
		{
			name:        "nil userData",
			userData:    nil,
			expectError: true,
			errorMsg:    "userData cannot be nil",
		},
		{
			name:        "missing id field",
			userData:    map[string]interface{}{"name": "Test"},
			expectError: true,
			errorMsg:    "userData must contain 'id' field",
		},
		{
			name: "id not string",
			userData: map[string]interface{}{
				"id":   123, // 数字而非字符串
				"name": "Test",
			},
			expectError: true,
			errorMsg:    "userData['id'] must be a string",
		},
		{
			name: "valid minimal userData",
			userData: map[string]interface{}{
				"id": "1234",
			},
			expectError: false,
		},
		{
			name: "valid full userData",
			userData: map[string]interface{}{
				"id":     "1234",
				"name":   "Test User",
				"email":  "test@example.com",
				"role":   "admin",
				"avatar": "https://example.com/avatar.png",
			},
			expectError: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			auth, err := GenerateUserAuthSignature(socketID, tc.userData)

			if tc.expectError {
				if err == nil {
					t.Errorf("Expected error but got none")
				} else {
					t.Logf("✅ Got expected error: %v", err)
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				} else {
					t.Logf("✅ Generated auth: %s", auth)
				}
			}
		})
	}
}

// TestGenerateUserAuthSignature_Format 测试签名字符串格式
func TestGenerateUserAuthSignature_Format(t *testing.T) {
	InitPusherAuth("test-app-key", "test-app-secret")

	// 测试用例：验证待签名字符串的格式
	// 格式应为：socket_id::user::user_data_json

	socketID := "server1.123456"
	userData := map[string]interface{}{
		"id":   "user123",
		"name": "Alice",
	}

	// 手动构建期望的待签名字符串
	userDataJSON, _ := json.Marshal(userData)
	expectedStringToSign := socketID + "::user::" + string(userDataJSON)

	t.Logf("Expected string to sign: %s", expectedStringToSign)

	// 生成签名
	auth, err := GenerateUserAuthSignature(socketID, userData)
	if err != nil {
		t.Fatalf("Failed to generate signature: %v", err)
	}

	// 验证签名不为空
	if auth == "" {
		t.Error("Generated auth is empty")
	}

	t.Logf("✅ Generated auth signature: %s", auth)
}

// TestUserAuthSignature_ConsistencyCheck 测试签名一致性
func TestUserAuthSignature_ConsistencyCheck(t *testing.T) {
	InitPusherAuth("test-app-key", "test-app-secret")

	socketID := "server1.123456"
	userData := map[string]interface{}{
		"id":   "user123",
		"name": "Alice",
	}

	// 生成两次签名，应该相同（确定性）
	auth1, err1 := GenerateUserAuthSignature(socketID, userData)
	auth2, err2 := GenerateUserAuthSignature(socketID, userData)

	if err1 != nil || err2 != nil {
		t.Fatalf("Failed to generate signatures: %v, %v", err1, err2)
	}

	if auth1 != auth2 {
		t.Errorf("Signatures should be identical:\n  1st: %s\n  2nd: %s", auth1, auth2)
	} else {
		t.Logf("✅ Signatures are consistent: %s", auth1)
	}
}

// TestUserAuthSignature_DifferentUsers 测试不同用户生成不同签名
func TestUserAuthSignature_DifferentUsers(t *testing.T) {
	InitPusherAuth("test-app-key", "test-app-secret")

	socketID := "server1.123456"

	userData1 := map[string]interface{}{
		"id":   "user1",
		"name": "Alice",
	}

	userData2 := map[string]interface{}{
		"id":   "user2",
		"name": "Bob",
	}

	auth1, _ := GenerateUserAuthSignature(socketID, userData1)
	auth2, _ := GenerateUserAuthSignature(socketID, userData2)

	if auth1 == auth2 {
		t.Errorf("Different users should generate different signatures")
	} else {
		t.Logf("✅ User 1: %s", auth1)
		t.Logf("✅ User 2: %s", auth2)
	}
}

// BenchmarkGenerateUserAuthSignature 性能基准测试
func BenchmarkGenerateUserAuthSignature(b *testing.B) {
	InitPusherAuth("test-app-key", "test-app-secret")

	socketID := "server1.123456"
	userData := map[string]interface{}{
		"id":     "user123",
		"name":   "Test User",
		"email":  "test@example.com",
		"avatar": "https://example.com/avatar.png",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := GenerateUserAuthSignature(socketID, userData)
		if err != nil {
			b.Fatalf("Failed to generate signature: %v", err)
		}
	}
}

func TestValidateUserAuthSignature(t *testing.T) {
	InitPusherAuth("test-app-key", "test-app-secret")

	socketID := "server1.123456"
	userData := map[string]interface{}{
		"id":   "user123",
		"name": "Alice",
	}

	auth, err := GenerateUserAuthSignature(socketID, userData)
	if err != nil {
		t.Fatalf("GenerateUserAuthSignature failed: %v", err)
	}

	userDataJSON, _ := json.Marshal(userData)
	if err := ValidateUserAuthSignature(socketID, auth, string(userDataJSON)); err != nil {
		t.Fatalf("ValidateUserAuthSignature should succeed: %v", err)
	}

	if err := ValidateUserAuthSignature(socketID, "bad-app:bad-sign", string(userDataJSON)); err == nil {
		t.Fatalf("ValidateUserAuthSignature should fail for invalid auth")
	}

	if err := ValidateUserAuthSignature(socketID, auth, `{"id":"tampered"}`); err == nil {
		t.Fatalf("ValidateUserAuthSignature should fail for tampered user_data")
	}
}
