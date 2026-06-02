package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/encoding/gyaml"
	"github.com/gogf/gf/v2/os/gfile"
)

// ModuleConfigParser 模块配置解析器
type ModuleConfigParser struct{}

// NewModuleConfigParser 创建配置解析器
func NewModuleConfigParser() *ModuleConfigParser {
	return &ModuleConfigParser{}
}

// LoadConfig 加载模块配置
// 优先加载 .module.yaml，不存在则加载 module.json
func (p *ModuleConfigParser) LoadConfig(modulePath string) (*ModuleConfig, error) {
	// 尝试加载 .module.yaml
	yamlPath := filepath.Join(modulePath, ".module.yaml")
	if gfile.Exists(yamlPath) {
		return p.loadYAMLConfig(yamlPath)
	}

	// 尝试加载 module.json（向后兼容）
	jsonPath := filepath.Join(modulePath, "module.json")
	if gfile.Exists(jsonPath) {
		return p.loadJSONConfig(jsonPath)
	}

	return nil, fmt.Errorf("未找到模块配置文件: .module.yaml 或 module.json")
}

// loadYAMLConfig 加载YAML格式配置
func (p *ModuleConfigParser) loadYAMLConfig(path string) (*ModuleConfig, error) {
	content := gfile.GetContents(path)
	if content == "" {
		return nil, fmt.Errorf("配置文件为空: %s", path)
	}

	config := &ModuleConfig{}
	if err := gyaml.DecodeTo([]byte(content), config); err != nil {
		return nil, fmt.Errorf("解析YAML配置失败: %w", err)
	}

	return config, nil
}

// loadJSONConfig 加载JSON格式配置（向后兼容）
func (p *ModuleConfigParser) loadJSONConfig(path string) (*ModuleConfig, error) {
	jsonObj, err := gjson.LoadPath(path, gjson.Options{Safe: true})
	if err != nil {
		return nil, fmt.Errorf("加载JSON配置失败: %w", err)
	}

	// 转换为ModuleConfig
	config := &ModuleConfig{
		Name:         jsonObj.Get("name").String(),
		Version:      jsonObj.Get("version").String(),
		Author:       jsonObj.Get("author").String(),
		License:      jsonObj.Get("license").String(),
		GoVersion:    jsonObj.Get("goVersion").String(),
		Dependencies: make(map[string]string),
		Variables:    make(map[string]string),
	}

	// 转换mod字段为dependencies
	if mod := jsonObj.Get("mod"); mod != nil {
		for k, v := range mod.Map() {
			config.Dependencies[k] = fmt.Sprintf("%v", v)
		}
	}

	// 转换files字段
	if files := jsonObj.Get("files"); files != nil {
		filesMap := files.Map()
		if go_files, ok := filesMap["go"]; ok {
			for _, f := range go_files.([]interface{}) {
				config.Files.Go = append(config.Files.Go, fmt.Sprintf("%v", f))
			}
		}
		if sql_files, ok := filesMap["sql"]; ok {
			for _, f := range sql_files.([]interface{}) {
				config.Files.SQL = append(config.Files.SQL, fmt.Sprintf("%v", f))
			}
		}
		if static_files, ok := filesMap["static"]; ok {
			for _, f := range static_files.([]interface{}) {
				config.Files.Static = append(config.Files.Static, fmt.Sprintf("%v", f))
			}
		}
		if other_files, ok := filesMap["other"]; ok {
			for _, f := range other_files.([]interface{}) {
				config.Files.Other = append(config.Files.Other, fmt.Sprintf("%v", f))
			}
		}
	}

	// 设置默认值
	if config.Security.Signature.Algorithm == "" {
		config.Security.Signature.Algorithm = "RSA"
	}
	config.Security.Permissions.FileSystem = true
	config.Security.Permissions.Database = true

	return config, nil
}

// SaveConfig 保存模块配置
func (p *ModuleConfigParser) SaveConfig(config *ModuleConfig, modulePath string, format string) error {
	var content []byte
	var err error
	var filePath string

	switch format {
	case "yaml":
		content, err = gyaml.Encode(config)
		if err != nil {
			return fmt.Errorf("编码YAML配置失败: %w", err)
		}
		filePath = filepath.Join(modulePath, ".module.yaml")

	case "json":
		// 转换为ModuleMetadata格式（向后兼容）
		metadata := p.configToMetadata(config)
		content, err = gjson.MarshalIndent(metadata, "", "    ")
		if err != nil {
			return fmt.Errorf("编码JSON配置失败: %w", err)
		}
		filePath = filepath.Join(modulePath, "module.json")

	default:
		return fmt.Errorf("不支持的配置格式: %s", format)
	}

	if err := gfile.PutContents(filePath, string(content)); err != nil {
		return fmt.Errorf("保存配置文件失败: %w", err)
	}

	return nil
}

// configToMetadata 将ModuleConfig转换为ModuleMetadata（向后兼容）
func (p *ModuleConfigParser) configToMetadata(config *ModuleConfig) *ModuleMetadata {
	return &ModuleMetadata{
		Name:      config.Name,
		Author:    config.Author,
		Version:   config.Version,
		License:   config.License,
		GoVersion: config.GoVersion,
		Mod:       config.Dependencies,
		Files: map[string][]string{
			"go":     config.Files.Go,
			"sql":    config.Files.SQL,
			"static": config.Files.Static,
			"other":  config.Files.Other,
		},
	}
}

// ValidateConfig 验证模块配置
func (p *ModuleConfigParser) ValidateConfig(config *ModuleConfig) []error {
	var errors []error

	// 验证必填字段
	if config.Name == "" {
		errors = append(errors, fmt.Errorf("模块名称不能为空"))
	}
	if config.Version == "" {
		errors = append(errors, fmt.Errorf("模块版本不能为空"))
	}
	if config.Author == "" {
		errors = append(errors, fmt.Errorf("模块作者不能为空"))
	}

	// 验证文件路径
	for _, path := range config.Files.Go {
		if !p.isValidPath(path) {
			errors = append(errors, fmt.Errorf("无效的Go文件路径: %s", path))
		}
	}
	for _, path := range config.Files.SQL {
		if !p.isValidPath(path) {
			errors = append(errors, fmt.Errorf("无效的SQL文件路径: %s", path))
		}
	}

	// 验证配置合并设置
	if config.ConfigMerge.Enabled {
		for _, file := range config.ConfigMerge.Files {
			if file.Strategy != "merge" && file.Strategy != "replace" && file.Strategy != "skip" {
				errors = append(errors, fmt.Errorf("无效的合并策略: %s", file.Strategy))
			}
		}
	}

	// 验证静态资源部署设置
	if config.StaticDeploy.Enabled {
		for _, rule := range config.StaticDeploy.Rules {
			if rule.Method != "copy" && rule.Method != "symlink" && rule.Method != "merge" {
				errors = append(errors, fmt.Errorf("无效的部署方式: %s", rule.Method))
			}
		}
	}

	return errors
}

// isValidPath 验证路径是否有效
func (p *ModuleConfigParser) isValidPath(path string) bool {
	// 基本路径验证
	if path == "" {
		return false
	}
	// 检查是否包含危险字符（简化版本）
	if len(path) > 0 && path[0] == '/' {
		return false // 不允许根路径
	}
	// 检查是否是绝对路径
	if filepath.IsAbs(path) {
		return false
	}
	return true
}

// MigrateJSONToYAML 将module.json迁移到.module.yaml
func (p *ModuleConfigParser) MigrateJSONToYAML(modulePath string) error {
	jsonPath := filepath.Join(modulePath, "module.json")
	if !gfile.Exists(jsonPath) {
		return fmt.Errorf("module.json 不存在")
	}

	// 加载JSON配置
	config, err := p.loadJSONConfig(jsonPath)
	if err != nil {
		return err
	}

	// 保存为YAML格式
	if err := p.SaveConfig(config, modulePath, "yaml"); err != nil {
		return err
	}

	// 成功后可以选择删除或重命名旧的JSON文件
	backupPath := jsonPath + ".backup"
	if err := os.Rename(jsonPath, backupPath); err != nil {
		return fmt.Errorf("备份JSON文件失败: %w", err)
	}

	return nil
}
