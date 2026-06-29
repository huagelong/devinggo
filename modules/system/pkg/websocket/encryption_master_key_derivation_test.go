// Package websocket
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package websocket

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestInitPusherEncryption tests encryption master key initialization
func TestInitPusherEncryption(t *testing.T) {
	t.Log("========== Test Init Pusher Encryption ==========")

	// Generate test master key
	masterKey := make([]byte, 32)
	for i := range masterKey {
		masterKey[i] = byte(i)
	}
	masterKeyBase64 := base64.StdEncoding.EncodeToString(masterKey)

	testCases := []struct {
		name        string
		keyBase64   string
		expectError bool
		expectHas   bool
	}{
		{"empty key", "", false, false},
		{"valid 32 bytes", masterKeyBase64, false, true},
		{"invalid base64", "not-valid-base64!!!", true, false},
		{"wrong length (16 bytes)", base64.StdEncoding.EncodeToString(make([]byte, 16)), true, false},
		{"wrong length (64 bytes)", base64.StdEncoding.EncodeToString(make([]byte, 64)), true, false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := InitPusherEncryption(tc.keyBase64)
			if tc.expectError {
				assert.Error(t, err, "Expected error for: "+tc.name)
			} else {
				assert.NoError(t, err, "Expected no error for: "+tc.name)
			}
			assert.Equal(t, tc.expectHas, HasEncryptionMasterKey(), "HasEncryptionMasterKey mismatch for: "+tc.name)
		})
	}

	// Cleanup
	InitPusherEncryption("")

	t.Log("\n✅ Init Pusher Encryption test passed")
}

// TestDeriveSharedSecret tests key derivation
func TestDeriveSharedSecret(t *testing.T) {
	t.Log("========== Test Derive Shared Secret ==========")

	// Generate test master key
	masterKey := make([]byte, 32)
	rand.Read(masterKey)
	masterKeyBase64 := base64.StdEncoding.EncodeToString(masterKey)

	// Initialize
	err := InitPusherEncryption(masterKeyBase64)
	assert.NoError(t, err)
	assert.True(t, HasEncryptionMasterKey())

	channel := "private-encrypted-test-channel"

	// Test same channel derives same secret
	secret1, err := DeriveSharedSecret(channel)
	assert.NoError(t, err, "First derivation should succeed")
	assert.NotEmpty(t, secret1, "Derived secret should not be empty")

	secret2, err := DeriveSharedSecret(channel)
	assert.NoError(t, err, "Second derivation should succeed")
	assert.Equal(t, secret1, secret2, "Same channel should derive same secret")

	t.Logf("  Channel: %s", channel)
	t.Logf("  Derived key: %s... (length: %d)", secret1[:20], len(secret1))

	// Cleanup
	InitPusherEncryption("")

	t.Log("\n✅ Derive Shared Secret test passed")
}

// TestDeriveSharedSecret_Uniqueness tests different channels derive different keys
func TestDeriveSharedSecret_Uniqueness(t *testing.T) {
	t.Log("========== Test Derive Shared Secret Uniqueness ==========")

	// Generate test master key
	masterKey := make([]byte, 32)
	rand.Read(masterKey)
	InitPusherEncryption(base64.StdEncoding.EncodeToString(masterKey))

	channels := []string{
		"private-encrypted-channel-1",
		"private-encrypted-channel-2",
		"private-encrypted-user-123",
		"private-encrypted-user-124",
		"private-encrypted-payment",
	}

	secrets := make(map[string]string)

	for _, ch := range channels {
		secret, err := DeriveSharedSecret(ch)
		assert.NoError(t, err, "Derivation should succeed for: "+ch)
		assert.NotEmpty(t, secret, "Secret should not be empty for: "+ch)

		// Verify uniqueness
		if existingChannel, exists := secrets[secret]; exists {
			t.Errorf("Different channels derived same secret: %s and %s", existingChannel, ch)
		}
		secrets[secret] = ch
		t.Logf("  Channel %s -> key %s...", ch, secret[:20])
	}

	assert.Equal(t, len(channels), len(secrets), "All channels should derive unique secrets")

	// Cleanup
	InitPusherEncryption("")

	t.Log("\n✅ Derive Shared Secret Uniqueness test passed")
}

// TestEncryptionWithDerivedKey tests encryption with derived key
func TestEncryptionWithDerivedKey(t *testing.T) {
	ctx := context.Background()

	t.Log("========== Test Encryption With Derived Key ==========")

	// Initialize master key
	masterKey := make([]byte, 32)
	rand.Read(masterKey)
	InitPusherEncryption(base64.StdEncoding.EncodeToString(masterKey))

	channel := "private-encrypted-test"
	plaintext := `{"message":"Hello World","amount":123.45}`

	// Derive key and encrypt
	secret, err := DeriveSharedSecret(channel)
	assert.NoError(t, err)

	encrypted, err := EncryptMessage(ctx, plaintext, secret)
	assert.NoError(t, err)
	assert.NotEmpty(t, encrypted)

	// Verify encrypted message format
	var result EncryptedMessage
	err = json.Unmarshal([]byte(encrypted), &result)
	assert.NoError(t, err, "Encrypted message should be valid JSON")
	assert.NotEmpty(t, result.Ciphertext, "Should have ciphertext")
	assert.NotEmpty(t, result.Nonce, "Should have nonce")

	t.Logf("  Channel: %s", channel)
	t.Logf("  Plaintext: %s", plaintext)
	t.Logf("  Encrypted: %s...", encrypted[:50])

	// Cleanup
	InitPusherEncryption("")

	t.Log("\n✅ Encryption With Derived Key test passed")
}

// TestBackwardCompatibility tests backward compatibility without master key
func TestBackwardCompatibility(t *testing.T) {
	ctx := context.Background()

	t.Log("========== Test Backward Compatibility ==========")

	// Ensure no master key configured
	InitPusherEncryption("")
	assert.False(t, HasEncryptionMasterKey(), "Should not have master key configured")

	// Test GenerateSharedSecret still works
	secret := GenerateSharedSecret()
	assert.NotEmpty(t, secret, "GenerateSharedSecret should work without master key")

	t.Logf("  Random key: %s... (length: %d)", secret[:20], len(secret))

	// Test encryption still works
	plaintext := `{"test":"data"}`
	encrypted, err := EncryptMessage(ctx, plaintext, secret)
	assert.NoError(t, err)
	assert.NotEmpty(t, encrypted)

	t.Logf("  Encrypted: %s...", encrypted[:50])

	t.Log("\n✅ Backward Compatibility test passed")
}

// TestDeriveSharedSecret_WithoutMasterKey tests derivation without master key
func TestDeriveSharedSecret_WithoutMasterKey(t *testing.T) {
	t.Log("========== Test Derive Shared Secret Without Master Key ==========")

	// Ensure no master key configured
	InitPusherEncryption("")
	assert.False(t, HasEncryptionMasterKey())

	// Derive should fail
	_, err := DeriveSharedSecret("private-encrypted-test")
	assert.Error(t, err, "DeriveSharedSecret should fail without master key")
	assert.Contains(t, err.Error(), "not configured", "Error should mention master key not configured")

	t.Log("\n✅ Derive Without Master Key test passed")
}
