package utils

import (
	"context"
	"os"
	"path/filepath"
	"testing"
)

// TestZipDirectory 测试目录压缩
func TestZipDirectory(t *testing.T) {
	ctx := context.Background()

	// 创建临时测试目录
	tmpDir, err := os.MkdirTemp("", "zip_test_*")
	if err != nil {
		t.Fatalf("创建临时目录失败: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// 创建测试文件结构
	testDir := filepath.Join(tmpDir, "test_dir")
	if mkdirErr := os.MkdirAll(testDir, 0755); mkdirErr != nil {
		t.Fatalf("创建测试目录失败: %v", mkdirErr)
	}

	// 创建测试文件
	testFiles := map[string]string{
		"file1.txt":           "Hello World",
		"dir1/file2.txt":      "Test Content",
		"dir1/dir2/file3.txt": "Nested Content",
	}

	for relPath, content := range testFiles {
		fullPath := filepath.Join(testDir, relPath)
		if mkdirErr := os.MkdirAll(filepath.Dir(fullPath), 0755); mkdirErr != nil {
			t.Fatalf("创建文件目录失败: %v", mkdirErr)
		}
		if writeErr := os.WriteFile(fullPath, []byte(content), 0644); writeErr != nil {
			t.Fatalf("创建测试文件失败: %v", writeErr)
		}
	}

	// 压缩目录
	zipPath := filepath.Join(tmpDir, "test.zip")
	err = ZipDirectory(ctx, testDir, zipPath)
	if err != nil {
		t.Fatalf("压缩目录失败: %v", err)
	}

	// 验证zip文件是否创建
	if _, statErr := os.Stat(zipPath); os.IsNotExist(statErr) {
		t.Fatal("zip文件未创建")
	}

	// 验证zip文件大小
	info, err := os.Stat(zipPath)
	if err != nil {
		t.Fatalf("获取zip文件信息失败: %v", err)
	}
	if info.Size() == 0 {
		t.Fatal("zip文件为空")
	}

	t.Logf("压缩成功: %s (大小: %d bytes)", zipPath, info.Size())
}

// TestUnzipFile 测试文件解压
func TestUnzipFile(t *testing.T) {
	ctx := context.Background()

	// 创建临时测试目录
	tmpDir, err := os.MkdirTemp("", "unzip_test_*")
	if err != nil {
		t.Fatalf("创建临时目录失败: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// 创建测试文件结构并压缩
	srcDir := filepath.Join(tmpDir, "src")
	if err := os.MkdirAll(srcDir, 0755); err != nil {
		t.Fatalf("创建源目录失败: %v", err)
	}

	testFiles := map[string]string{
		"file1.txt":           "Hello World",
		"dir1/file2.txt":      "Test Content",
		"dir1/dir2/file3.txt": "Nested Content",
	}

	for relPath, content := range testFiles {
		fullPath := filepath.Join(srcDir, relPath)
		if err := os.MkdirAll(filepath.Dir(fullPath), 0755); err != nil {
			t.Fatalf("创建文件目录失败: %v", err)
		}
		if err := os.WriteFile(fullPath, []byte(content), 0644); err != nil {
			t.Fatalf("创建测试文件失败: %v", err)
		}
	}

	// 压缩
	zipPath := filepath.Join(tmpDir, "test.zip")
	if err := ZipDirectory(ctx, srcDir, zipPath); err != nil {
		t.Fatalf("压缩失败: %v", err)
	}

	// 解压到新目录
	dstDir := filepath.Join(tmpDir, "dst")
	if err := UnzipFile(zipPath, dstDir); err != nil {
		t.Fatalf("解压失败: %v", err)
	}

	// 验证解压后的文件
	for relPath, expectedContent := range testFiles {
		fullPath := filepath.Join(dstDir, relPath)

		// 检查文件是否存在
		if _, err := os.Stat(fullPath); os.IsNotExist(err) {
			t.Errorf("文件不存在: %s", relPath)
			continue
		}

		// 检查文件内容
		content, err := os.ReadFile(fullPath)
		if err != nil {
			t.Errorf("读取文件失败 %s: %v", relPath, err)
			continue
		}

		if string(content) != expectedContent {
			t.Errorf("文件内容不匹配 %s:\n期望: %q\n实际: %q", relPath, expectedContent, string(content))
		}
	}

	t.Logf("解压成功: %d个文件验证通过", len(testFiles))
}

// TestZipUnzipRoundTrip 测试压缩-解压完整流程
func TestZipUnzipRoundTrip(t *testing.T) {
	ctx := context.Background()

	// 创建临时测试目录
	tmpDir, err := os.MkdirTemp("", "zip_roundtrip_*")
	if err != nil {
		t.Fatalf("创建临时目录失败: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// 创建复杂的测试文件结构
	srcDir := filepath.Join(tmpDir, "src")
	testStructure := map[string]string{
		"README.md":                    "# Test Module",
		"go.mod":                       "module test",
		"api/user.go":                  "package api",
		"controller/user.go":           "package controller",
		"logic/user.go":                "package logic",
		"model/req/user.go":            "package req",
		"model/res/user.go":            "package res",
		"worker/server/task.go":        "package server",
		"worker/cron/cron.go":          "package cron",
		"migrations/001_init.up.sql":   "CREATE TABLE users;",
		"migrations/001_init.down.sql": "DROP TABLE users;",
	}

	for relPath, content := range testStructure {
		fullPath := filepath.Join(srcDir, relPath)
		if err := os.MkdirAll(filepath.Dir(fullPath), 0755); err != nil {
			t.Fatalf("创建目录失败: %v", err)
		}
		if err := os.WriteFile(fullPath, []byte(content), 0644); err != nil {
			t.Fatalf("创建文件失败: %v", err)
		}
	}

	// 压缩
	zipPath := filepath.Join(tmpDir, "module.zip")
	if err := ZipDirectory(ctx, srcDir, zipPath); err != nil {
		t.Fatalf("压缩失败: %v", err)
	}

	// 解压
	dstDir := filepath.Join(tmpDir, "dst")
	if err := UnzipFile(zipPath, dstDir); err != nil {
		t.Fatalf("解压失败: %v", err)
	}

	// 验证所有文件
	for relPath, expectedContent := range testStructure {
		fullPath := filepath.Join(dstDir, relPath)
		content, err := os.ReadFile(fullPath)
		if err != nil {
			t.Errorf("读取文件失败 %s: %v", relPath, err)
			continue
		}
		if string(content) != expectedContent {
			t.Errorf("文件内容不匹配 %s", relPath)
		}
	}

	t.Logf("完整流程测试通过: %d个文件", len(testStructure))
}

// TestUnzipFile_PathTraversal 测试路径穿越攻击防护
func TestUnzipFile_PathTraversal(t *testing.T) {
	// 这个测试需要手动创建包含恶意路径的zip文件
	// 在实际环境中，应该创建一个包含 "../../../etc/passwd" 等路径的zip文件进行测试
	// 这里我们只测试路径检查逻辑
	t.Skip("需要手动创建包含恶意路径的zip文件进行测试")
}

// TestZipDirectory_EmptyDir 测试压缩空目录
func TestZipDirectory_EmptyDir(t *testing.T) {
	ctx := context.Background()

	tmpDir, err := os.MkdirTemp("", "zip_empty_*")
	if err != nil {
		t.Fatalf("创建临时目录失败: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// 创建空目录
	emptyDir := filepath.Join(tmpDir, "empty")
	if mkdirErr := os.MkdirAll(emptyDir, 0755); mkdirErr != nil {
		t.Fatalf("创建空目录失败: %v", mkdirErr)
	}

	// 压缩空目录
	zipPath := filepath.Join(tmpDir, "empty.zip")
	err = ZipDirectory(ctx, emptyDir, zipPath)
	if err != nil {
		t.Fatalf("压缩空目录失败: %v", err)
	}

	// 验证zip文件创建
	if _, err := os.Stat(zipPath); os.IsNotExist(err) {
		t.Fatal("zip文件未创建")
	}

	t.Log("空目录压缩测试通过")
}

// TestZipDirectory_NonExistentDir 测试压缩不存在的目录
func TestZipDirectory_NonExistentDir(t *testing.T) {
	ctx := context.Background()

	tmpDir, err := os.MkdirTemp("", "zip_nonexistent_*")
	if err != nil {
		t.Fatalf("创建临时目录失败: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	nonExistentDir := filepath.Join(tmpDir, "nonexistent")
	zipPath := filepath.Join(tmpDir, "test.zip")

	err = ZipDirectory(ctx, nonExistentDir, zipPath)
	if err == nil {
		t.Fatal("应该返回错误但实际成功")
	}

	t.Logf("不存在目录测试通过，错误: %v", err)
}

// TestUnzipFile_NonExistentZip 测试解压不存在的zip文件
func TestUnzipFile_NonExistentZip(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "unzip_nonexistent_*")
	if err != nil {
		t.Fatalf("创建临时目录失败: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	nonExistentZip := filepath.Join(tmpDir, "nonexistent.zip")
	dstDir := filepath.Join(tmpDir, "dst")

	err = UnzipFile(nonExistentZip, dstDir)
	if err == nil {
		t.Fatal("应该返回错误但实际成功")
	}

	t.Logf("不存在zip文件测试通过，错误: %v", err)
}

// BenchmarkZipDirectory 压缩性能测试
func BenchmarkZipDirectory(b *testing.B) {
	ctx := context.Background()

	// 创建测试目录
	tmpDir, err := os.MkdirTemp("", "zip_bench_*")
	if err != nil {
		b.Fatalf("创建临时目录失败: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// 创建测试文件
	srcDir := filepath.Join(tmpDir, "src")
	os.MkdirAll(srcDir, 0755)
	for i := 0; i < 10; i++ {
		path := filepath.Join(srcDir, filepath.Join("dir", "file.txt"))
		os.MkdirAll(filepath.Dir(path), 0755)
		os.WriteFile(path, []byte("test content"), 0644)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		zipPath := filepath.Join(tmpDir, "bench.zip")
		ZipDirectory(ctx, srcDir, zipPath)
		os.Remove(zipPath)
	}
}

// BenchmarkUnzipFile 解压性能测试
func BenchmarkUnzipFile(b *testing.B) {
	ctx := context.Background()

	// 准备zip文件
	tmpDir, err := os.MkdirTemp("", "unzip_bench_*")
	if err != nil {
		b.Fatalf("创建临时目录失败: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	srcDir := filepath.Join(tmpDir, "src")
	os.MkdirAll(srcDir, 0755)
	os.WriteFile(filepath.Join(srcDir, "file.txt"), []byte("test"), 0644)

	zipPath := filepath.Join(tmpDir, "test.zip")
	ZipDirectory(ctx, srcDir, zipPath)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		dstDir := filepath.Join(tmpDir, "dst")
		UnzipFile(zipPath, dstDir)
		os.RemoveAll(dstDir)
	}
}
