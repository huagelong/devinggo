// Package utils
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package utils

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// SensitivePattern 敏感信息模式
type SensitivePattern struct {
	Pattern     string   // 正则表达式模式
	Replacement string   // 替换模板
	Extensions  []string // 匹配的文件扩展名
	Description string   // 描述
}

// DefaultSensitivePatterns 默认的敏感信息模式
var DefaultSensitivePatterns = []SensitivePattern{
	{
		Pattern:     `(?i)(password|passwd|pwd)\s*[:=]\s*["']([^"']+)["']`,
		Replacement: `$1: "{{.DB_PASSWORD}}"`,
		Extensions:  []string{".go", ".yaml", ".yml", ".json", ".env", ".conf"},
		Description: "数据库密码",
	},
	{
		Pattern:     `(?i)(host|hostname|server)\s*[:=]\s*["']([^"']+\.local[^"']*|localhost|127\.0\.0\.1|0\.0\.0\.0)["']`,
		Replacement: `$1: "{{.DB_HOST}}"`,
		Extensions:  []string{".go", ".yaml", ".yml", ".json", ".env", ".conf"},
		Description: "本地数据库地址",
	},
	{
		Pattern:     `(?i)(apikey|api_key|accesskey|access_key|secretkey|secret_key)\s*[:=]\s*["']([^"']+)["']`,
		Replacement: `$1: "${API_KEY}"`,
		Extensions:  []string{".go", ".yaml", ".yml", ".json", ".env"},
		Description: "API密钥",
	},
	{
		Pattern:     `(?i)(jwtsecret|jwt_secret|jwt)\s*[:=]\s*["']([^"']+)["']`,
		Replacement: `$1: "{{.JWT_SECRET}}"`,
		Extensions:  []string{".go", ".yaml", ".yml", ".json", ".env"},
		Description: "JWT密钥",
	},
	{
		Pattern:     `(?i)(redispassword|redis_pwd|redis_pwd)\s*[:=]\s*["']([^"']+)["']`,
		Replacement: `$1: "{{.REDIS_PASSWORD}}"`,
		Extensions:  []string{".go", ".yaml", ".yml", ".json", ".env"},
		Description: "Redis密码",
	},
	{
		Pattern:     `(?i)(user|username|dbuser|db_user)\s*[:=]\s*["']([^"']+)["']`,
		Replacement: `$1: "{{.DB_USER}}"`,
		Extensions:  []string{".go", ".yaml", ".yml", ".json", ".env"},
		Description: "数据库用户名",
	},
	{
		Pattern:     `(?i)(database|dbname|db_name)\s*[:=]\s*["']([^"']+)["']`,
		Replacement: `$1: "{{.DB_NAME}}"`,
		Extensions:  []string{".go", ".yaml", ".yml", ".json", ".env"},
		Description: "数据库名称",
	},
	{
		Pattern:     `(?i)(smtp_pass|smtp_password|mailpassword|mail_password)\s*[:=]\s*["']([^"']+)["']`,
		Replacement: `$1: "{{.SMTP_PASSWORD}}"`,
		Extensions:  []string{".go", ".yaml", ".yml", ".json", ".env"},
		Description: "SMTP密码",
	},
}

// Sanitizer 敏感信息清理器
type Sanitizer struct {
	patterns []SensitivePattern
	dryRun   bool
}

// NewSanitizer 创建敏感信息清理器
func NewSanitizer(patterns []SensitivePattern) *Sanitizer {
	if patterns == nil {
		patterns = DefaultSensitivePatterns
	}
	return &Sanitizer{
		patterns: patterns,
		dryRun:   false,
	}
}

// SetDryRun 设置是否为演练模式（不实际修改文件）
func (s *Sanitizer) SetDryRun(dryRun bool) {
	s.dryRun = dryRun
}

// SanitizeFile 清理单个文件中的敏感信息
func (s *Sanitizer) SanitizeFile(filePath string) ([]string, error) {
	// 检查文件扩展名
	ext := strings.ToLower(filepath.Ext(filePath))
	shouldProcess := false
	for _, pattern := range s.patterns {
		for _, allowedExt := range pattern.Extensions {
			if ext == allowedExt {
				shouldProcess = true
				break
			}
		}
		if shouldProcess {
			break
		}
	}

	if !shouldProcess {
		return nil, nil
	}

	// 读取文件
	content, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("读取文件失败: %w", err)
	}

	originalContent := string(content)
	newContent := originalContent
	changes := []string{}

	// 应用所有模式
	for _, pattern := range s.patterns {
		// 检查是否应该处理此扩展名
		shouldApply := false
		for _, allowedExt := range pattern.Extensions {
			if ext == allowedExt {
				shouldApply = true
				break
			}
		}
		if !shouldApply {
			continue
		}

		// 编译正则表达式
		re, err := regexp.Compile(pattern.Pattern)
		if err != nil {
			continue
		}

		// 查找匹配项
		matches := re.FindAllStringSubmatch(newContent, -1)
		if len(matches) > 0 {
			for _, match := range matches {
				if len(match) > 0 {
					changes = append(changes, fmt.Sprintf("  [%s] %s -> %s", pattern.Description, match[0], pattern.Replacement))
				}
			}

			// 替换内容
			newContent = re.ReplaceAllString(newContent, pattern.Replacement)
		}
	}

	// 如果有变化，写回文件
	if newContent != originalContent && !s.dryRun {
		if err := os.WriteFile(filePath, []byte(newContent), 0644); err != nil {
			return nil, fmt.Errorf("写入文件失败: %w", err)
		}
	}

	return changes, nil
}

// SanitizeDirectory 清理目录中的所有文件
func (s *Sanitizer) SanitizeDirectory(dirPath string) (map[string][]string, error) {
	results := make(map[string][]string)

	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 跳过目录和隐藏文件
		if info.IsDir() || strings.HasPrefix(filepath.Base(path), ".") {
			return nil
		}

		// 跳过 vendor、node_modules 等
		if strings.Contains(path, "vendor") || strings.Contains(path, "node_modules") {
			return nil
		}

		// 获取相对路径
		relPath, err := filepath.Rel(dirPath, path)
		if err != nil {
			return nil
		}

		// 清理文件
		changes, err := s.SanitizeFile(path)
		if err != nil {
			return fmt.Errorf("清理文件 %s 失败: %w", relPath, err)
		}

		if len(changes) > 0 {
			results[relPath] = changes
		}

		return nil
	})

	return results, err
}

// SanitizeDirectoryWithReport 清理目录并生成报告
func (s *Sanitizer) SanitizeDirectoryWithReport(dirPath string) error {
	fmt.Printf("\n🔍 扫描敏感信息: %s\n", dirPath)

	results, err := s.SanitizeDirectory(dirPath)
	if err != nil {
		return err
	}

	if len(results) == 0 {
		fmt.Println("  ✅ 未发现敏感信息")
		return nil
	}

	totalChanges := 0
	for file, changes := range results {
		fmt.Printf("\n  📝 %s:\n", file)
		for _, change := range changes {
			fmt.Println(change)
		}
		totalChanges += len(changes)
	}

	fmt.Printf("\n  📊 共处理 %d 个文件，替换 %d 处敏感信息\n", len(results), totalChanges)

	if s.dryRun {
		fmt.Println("  ⚠️  演练模式：文件未被实际修改")
	}

	return nil
}

// ShouldSanitizeFile 检查文件是否应该被清理
func ShouldSanitizeFile(filePath string) bool {
	ext := strings.ToLower(filepath.Ext(filePath))
	sanitizeExts := map[string]bool{
		".go":     true,
		".yaml":   true,
		".yml":    true,
		".json":   true,
		".env":    true,
		".conf":   true,
		".config": true,
		".ini":    true,
	}
	return sanitizeExts[ext]
}

// FindSensitiveInFile 在文件中查找敏感信息（仅查找，不替换）
func FindSensitiveInFile(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var findings []string
	scanner := bufio.NewScanner(file)
	lineNum := 0

	for scanner.Scan() {
		lineNum++
		line := scanner.Text()

		// 检查常见敏感信息模式
		for _, pattern := range DefaultSensitivePatterns {
			re, err := regexp.Compile(pattern.Pattern)
			if err != nil {
				continue
			}

			if re.MatchString(line) {
				findings = append(findings, fmt.Sprintf("%s:%d - [%s] %s", filepath.Base(filePath), lineNum, pattern.Description, strings.TrimSpace(line)))
			}
		}
	}

	return findings, scanner.Err()
}
