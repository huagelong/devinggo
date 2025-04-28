// Package secure
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package secure

import (
	"bytes"
	"crypto/aes"
	"encoding/base64"
	"fmt"
	"github.com/gogf/gf/v2/util/gconv"
	"golang.org/x/crypto/bcrypt"
)

func PasswordHash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func PasswordVerify(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// PKCS7 填充函数
func pkcs7Pad(data []byte, blockSize int) []byte {
	padding := blockSize - (len(data) % blockSize)
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padText...)
}

// PKCS7 去填充函数
func pkcs7Unpad(data []byte) []byte {
	padding := int(data[len(data)-1])
	return data[:len(data)-padding]
}

// AES加密 (ECB模式)
func AESEncrypt(str string, salt string) (string, error) {
	plaintext := []byte(str)
	key := []byte(salt)
	// 检查密钥长度
	if len(key) != 16 && len(key) != 24 && len(key) != 32 {
		return "", fmt.Errorf("密钥必须为 16/24/32 字节 (AES-128/192/256)")
	}

	// PKCS7 填充
	blockSize := aes.BlockSize // 16字节
	paddedData := pkcs7Pad(plaintext, blockSize)

	// 创建 AES 密码块
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// ECB 模式加密（手动分块）
	encrypted := make([]byte, len(paddedData))
	for i := 0; i < len(paddedData); i += blockSize {
		block.Encrypt(encrypted[i:], paddedData[i:i+blockSize])
	}

	// 返回 Base64
	return base64.StdEncoding.EncodeToString(encrypted), nil
}

// AES-ECB 解密（手动分块）
func AESDecrypt(str string, salt string) (string, error) {
	ciphertext := str
	key := []byte(salt)
	// Base64 解码
	encrypted, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}

	// 创建 AES 密码块
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// ECB 模式解密（手动分块）
	decrypted := make([]byte, len(encrypted))
	for i := 0; i < len(encrypted); i += aes.BlockSize {
		block.Decrypt(decrypted[i:], encrypted[i:i+aes.BlockSize])
	}

	// PKCS7 去填充
	return gconv.String(pkcs7Unpad(decrypted)), nil
}
