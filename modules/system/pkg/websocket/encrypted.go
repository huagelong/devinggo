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
	"fmt"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"golang.org/x/crypto/nacl/secretbox"
)

const (
	// Redis key prefix for storing shared secrets
	sharedSecretKeyPrefix = "pusher:encrypted_secret:"
	// TTL for shared secrets (24 hours)
	sharedSecretTTL = 24 * time.Hour
)

// SaveSharedSecret 保存 shared_secret 到 Redis
// 用于服务器端推送加密消息
// 注意：如果配置了 encryptionMasterKey，派生模式下无需保存
func SaveSharedSecret(ctx context.Context, channelName string, sharedSecret string) error {
	// 如果配置了主密钥，派生模式下无需保存（密钥可重复派生）
	if HasEncryptionMasterKey() {
		g.Log().Debugf(ctx, "Encryption master key mode: skipping save for channel: %s", channelName)
		return nil
	}

	// 否则保存到 Redis（per-channel 模式，向后兼容）
	redisKey := sharedSecretKeyPrefix + channelName

	_, err := g.Redis().Set(ctx, redisKey, sharedSecret)
	if err != nil {
		return fmt.Errorf("failed to save shared_secret: %w", err)
	}

	// 设置过期时间
	_, err = g.Redis().Expire(ctx, redisKey, int64(sharedSecretTTL.Seconds()))
	if err != nil {
		g.Log().Warning(ctx, "Failed to set TTL for shared_secret:", err)
	}

	g.Log().Debugf(ctx, "Saved shared_secret for channel: %s", channelName)
	return nil
}

// GetSharedSecret 从 Redis 获取 shared_secret 或从主密钥派生
// 优先使用派生模式（如果配置了 encryptionMasterKey）
func GetSharedSecret(ctx context.Context, channelName string) (string, error) {
	// 如果配置了主密钥，使用派生模式
	if HasEncryptionMasterKey() {
		return DeriveSharedSecret(channelName)
	}

	// 否则从 Redis 获取（per-channel 模式，向后兼容）
	redisKey := sharedSecretKeyPrefix + channelName

	result, err := g.Redis().Get(ctx, redisKey)
	if err != nil {
		return "", fmt.Errorf("failed to get shared_secret: %w", err)
	}

	sharedSecret := result.String()
	if sharedSecret == "" {
		return "", fmt.Errorf("shared_secret not found for channel: %s", channelName)
	}

	return sharedSecret, nil
}

// EncryptedMessage Pusher Encrypted Channels 消息格式
type EncryptedMessage struct {
	Nonce      string `json:"nonce"`      // Base64 encoded nonce (24 bytes)
	Ciphertext string `json:"ciphertext"` // Base64 encoded encrypted data
}

// EncryptMessage 使用 NaCl secretbox 加密消息
// 兼容 Pusher.js/TweetNaCl 格式
func EncryptMessage(ctx context.Context, plaintext string, sharedSecretBase64 string) (string, error) {
	// 1. 解码 shared_secret (Base64 -> bytes)
	sharedSecretBytes, err := base64.StdEncoding.DecodeString(sharedSecretBase64)
	if err != nil {
		return "", fmt.Errorf("invalid shared_secret base64: %w", err)
	}

	if len(sharedSecretBytes) != 32 {
		return "", fmt.Errorf("shared_secret must be 32 bytes, got %d", len(sharedSecretBytes))
	}

	// 转换为 secretbox 密钥格式 [32]byte
	var key [32]byte
	copy(key[:], sharedSecretBytes)

	// 2. 生成随机 nonce (24 bytes for NaCl secretbox)
	var nonce [24]byte
	if _, randErr := rand.Read(nonce[:]); randErr != nil {
		return "", fmt.Errorf("failed to generate nonce: %w", randErr)
	}

	// 3. 加密数据
	plaintextBytes := []byte(plaintext)
	encrypted := secretbox.Seal(nil, plaintextBytes, &nonce, &key)

	// 4. 构建 Pusher 格式的加密消息
	encryptedMsg := EncryptedMessage{
		Nonce:      base64.StdEncoding.EncodeToString(nonce[:]),
		Ciphertext: base64.StdEncoding.EncodeToString(encrypted),
	}

	// 5. 序列化为 JSON
	result, err := json.Marshal(encryptedMsg)
	if err != nil {
		return "", fmt.Errorf("failed to marshal encrypted message: %w", err)
	}

	g.Log().Debugf(ctx, "Encrypted message: nonce=%s, ciphertext_len=%d",
		encryptedMsg.Nonce[:8]+"...", len(encrypted))

	return string(result), nil
}

// EncryptChannelData 加密频道数据（用于服务器端推送）
// 自动处理 JSON 编码和加密
func EncryptChannelData(ctx context.Context, data interface{}, sharedSecret string) (string, error) {
	// 1. 将数据序列化为 JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", fmt.Errorf("failed to marshal data: %w", err)
	}

	// 2. 加密 JSON 数据
	encrypted, err := EncryptMessage(ctx, string(jsonData), sharedSecret)
	if err != nil {
		return "", fmt.Errorf("failed to encrypt data: %w", err)
	}

	return encrypted, nil
}
