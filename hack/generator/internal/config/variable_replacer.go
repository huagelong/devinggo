package config

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/gogf/gf/v2/os/gfile"
)

// VariableReplacer 变量替换引擎
type VariableReplacer struct {
	variables map[string]string
	pattern   *regexp.Regexp
}

// NewVariableReplacer 创建变量替换引擎
func NewVariableReplacer(variables map[string]string) *VariableReplacer {
	// 支持 {{.VarName}} 和 ${VarName} 两种格式
	pattern := regexp.MustCompile(`\{\{\.(\w+)\}\}|\$\{(\w+)\}`)

	return &VariableReplacer{
		variables: variables,
		pattern:   pattern,
	}
}

// ReplaceString 替换字符串中的变量
func (r *VariableReplacer) ReplaceString(content string) string {
	return r.pattern.ReplaceAllStringFunc(content, func(match string) string {
		// 提取变量名
		varName := ""
		if strings.HasPrefix(match, "{{") {
			varName = strings.TrimSuffix(strings.TrimPrefix(match, "{{."), "}}")
		} else if strings.HasPrefix(match, "${") {
			varName = strings.TrimSuffix(strings.TrimPrefix(match, "${"), "}")
		}

		// 查找变量值
		if value, ok := r.variables[varName]; ok {
			return value
		}

		// 如果变量未定义，保持原样
		return match
	})
}

// ReplaceFile 替换文件中的变量
func (r *VariableReplacer) ReplaceFile(filePath string) error {
	if !gfile.Exists(filePath) {
		return fmt.Errorf("文件不存在: %s", filePath)
	}

	content := gfile.GetContents(filePath)
	replaced := r.ReplaceString(content)

	if err := gfile.PutContents(filePath, replaced); err != nil {
		return fmt.Errorf("写入文件失败: %w", err)
	}

	return nil
}

// ReplaceInDirectory 替换目录中所有文件的变量
func (r *VariableReplacer) ReplaceInDirectory(dir string, patterns []string) error {
	if len(patterns) == 0 {
		patterns = []string{"*"}
	}

	for _, pattern := range patterns {
		files, err := gfile.ScanDirFile(dir, pattern, true)
		if err != nil {
			return fmt.Errorf("扫描目录失败: %w", err)
		}

		for _, file := range files {
			if err := r.ReplaceFile(file); err != nil {
				return fmt.Errorf("替换文件 '%s' 失败: %w", file, err)
			}
		}
	}

	return nil
}

// SetVariable 设置或更新变量
func (r *VariableReplacer) SetVariable(key, value string) {
	r.variables[key] = value
}

// GetVariable 获取变量值
func (r *VariableReplacer) GetVariable(key string) (string, bool) {
	value, ok := r.variables[key]
	return value, ok
}

// GetAllVariables 获取所有变量
func (r *VariableReplacer) GetAllVariables() map[string]string {
	// 返回副本以避免外部修改
	result := make(map[string]string)
	for k, v := range r.variables {
		result[k] = v
	}
	return result
}

// MergeVariables 合并新变量
func (r *VariableReplacer) MergeVariables(newVars map[string]string) {
	for k, v := range newVars {
		r.variables[k] = v
	}
}

// ExtractVariables 从模板字符串中提取所有变量名
func ExtractVariables(template string) []string {
	pattern := regexp.MustCompile(`\{\{\.(\w+)\}\}|\$\{(\w+)\}`)
	matches := pattern.FindAllStringSubmatch(template, -1)

	vars := make(map[string]bool)
	for _, match := range matches {
		if match[1] != "" {
			vars[match[1]] = true
		} else if match[2] != "" {
			vars[match[2]] = true
		}
	}

	result := make([]string, 0, len(vars))
	for v := range vars {
		result = append(result, v)
	}

	return result
}

// ValidateVariables 验证所有必需的变量是否已定义
func (r *VariableReplacer) ValidateVariables(template string) []string {
	required := ExtractVariables(template)
	missing := make([]string, 0)

	for _, varName := range required {
		if _, ok := r.variables[varName]; !ok {
			missing = append(missing, varName)
		}
	}

	return missing
}

// BuildDefaultVariables 构建默认变量集
func BuildDefaultVariables(moduleName string) map[string]string {
	// 处理空字符串，避免 panic
	if moduleName == "" {
		moduleName = "module"
	}

	// 使用 gstr.UcFirst 更安全地处理首字母大写
	moduleNameCap := moduleName
	if len(moduleName) > 0 {
		moduleNameCap = strings.ToUpper(moduleName[:1]) + moduleName[1:]
	}

	return map[string]string{
		"moduleName":    moduleName,
		"moduleNameCap": moduleNameCap,
		"ModuleName":    moduleNameCap,
		"MODULE_NAME":   strings.ToUpper(moduleName),
		"module_name":   strings.ToLower(moduleName),
	}
}
