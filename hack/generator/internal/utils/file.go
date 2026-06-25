package utils

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// GetTmpDir 获取临时目录
func GetTmpDir() string {
	return os.TempDir()
}

// PathExists 检查路径是否存在
func PathExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

// IsDir 检查路径是否为目录
func IsDir(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return info.IsDir()
}

// IsFile 检查路径是否为文件
func IsFile(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return !info.IsDir()
}

// EnsureDir 确保目录存在，如果不存在则创建
func EnsureDir(dir string) error {
	if !PathExists(dir) {
		return os.MkdirAll(dir, 0755)
	}
	return nil
}

// CopyFile 复制文件
func CopyFile(src, dst string) error {
	// 读取源文件
	data, err := os.ReadFile(src)
	if err != nil {
		return fmt.Errorf("读取源文件失败: %w", err)
	}

	// 确保目标目录存在
	dstDir := filepath.Dir(dst)
	if err := EnsureDir(dstDir); err != nil {
		return fmt.Errorf("创建目标目录失败: %w", err)
	}

	// 写入目标文件
	if err := os.WriteFile(dst, data, 0644); err != nil {
		return fmt.Errorf("写入目标文件失败: %w", err)
	}

	return nil
}

// FormatGoCode 格式化Go代码文件
// 自动调用 gofmt 和 goimports（如果可用）
func FormatGoCode(filePath string) error {
	if !strings.HasSuffix(filePath, ".go") {
		return nil
	}

	// 先尝试 goimports
	if err := runCommand("goimports", "-w", filePath); err == nil {
		return nil
	}

	// 如果 goimports 不可用，使用 gofmt
	if err := runCommand("gofmt", "-w", filePath); err != nil {
		return fmt.Errorf("格式化代码失败: %w", err)
	}

	return nil
}

// FormatGoCodeInDir 格式化目录下的所有Go代码
func FormatGoCodeInDir(dir string) error {
	return filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && strings.HasSuffix(path, ".go") {
			if err := FormatGoCode(path); err != nil {
				// 继续处理其他文件，不中断
				fmt.Printf("警告: 格式化 %s 失败: %v\n", path, err)
			}
		}

		return nil
	})
}

// runCommand 执行命令
func runCommand(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// GetProjectRoot 获取项目根目录
// 通过查找go.mod文件来确定
func GetProjectRoot() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	// 向上查找go.mod文件
	for {
		if PathExists(filepath.Join(dir, "go.mod")) {
			return dir, nil
		}

		parent := filepath.Dir(dir)
		if parent == dir {
			// 已到根目录
			break
		}
		dir = parent
	}

	return "", fmt.Errorf("未找到项目根目录(go.mod)")
}

// GetModuleName 获取Go模块名称
func GetModuleName() (string, error) {
	root, err := GetProjectRoot()
	if err != nil {
		return "", err
	}

	// 读取go.mod文件
	data, err := os.ReadFile(filepath.Join(root, "go.mod"))
	if err != nil {
		return "", err
	}

	// 解析module名称
	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "module ") {
			return strings.TrimSpace(strings.TrimPrefix(line, "module")), nil
		}
	}

	return "", fmt.Errorf("未找到module声明")
}

// WriteFile 写入文件，如果已存在则提示
func WriteFile(path string, content []byte, overwrite bool) error {
	if PathExists(path) && !overwrite {
		return fmt.Errorf("文件已存在: %s", path)
	}

	// 确保目录存在
	dir := filepath.Dir(path)
	if err := EnsureDir(dir); err != nil {
		return err
	}

	return os.WriteFile(path, content, 0644)
}
