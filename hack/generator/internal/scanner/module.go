// Package scanner
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package scanner

import (
	"fmt"
	"os"
	"path/filepath"

	"devinggo/hack/generator/internal/config"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/os/gfile"
)

// ModuleInfo 模块信息
type ModuleInfo struct {
	Name      string              `json:"name"`
	Author    string              `json:"author"`
	Version   string              `json:"version"`
	License   string              `json:"license"`
	GoVersion string              `json:"goVersion"`
	Mod       map[string]string   `json:"mod"`
	Files     map[string][]string `json:"files"`
	Path      string              `json:"-"` // 模块路径
}

// ScanModule 扫描模块信息
func ScanModule(moduleName string) (*ModuleInfo, error) {
	modulePath := filepath.Join("modules", moduleName)

	// 检查模块目录是否存在
	if _, err := os.Stat(modulePath); os.IsNotExist(err) {
		return nil, fmt.Errorf("模块 '%s' 不存在", moduleName)
	}

	// 优先尝试加载 .module.yaml
	parser := config.NewModuleConfigParser()
	moduleConfig, err := parser.LoadConfig(modulePath)
	if err != nil {
		// 降级到旧的module.json方式
		return scanModuleLegacy(moduleName, modulePath)
	}

	// 转换为ModuleInfo
	moduleInfo := &ModuleInfo{
		Name:      moduleConfig.Name,
		Author:    moduleConfig.Author,
		Version:   moduleConfig.Version,
		License:   moduleConfig.License,
		GoVersion: moduleConfig.GoVersion,
		Path:      modulePath,
		Mod:       moduleConfig.Dependencies,
		Files:     make(map[string][]string),
	}

	// 转换文件列表
	moduleInfo.Files["go"] = moduleConfig.Files.Go
	moduleInfo.Files["sql"] = moduleConfig.Files.SQL
	moduleInfo.Files["static"] = moduleConfig.Files.Static
	moduleInfo.Files["other"] = moduleConfig.Files.Other

	return moduleInfo, nil
}

// scanModuleLegacy 使用旧方式扫描模块（向后兼容）
func scanModuleLegacy(moduleName, modulePath string) (*ModuleInfo, error) {
	// 检查module.json是否存在
	configPath := filepath.Join(modulePath, "module.json")
	if !gfile.Exists(configPath) {
		return nil, fmt.Errorf("模块 '%s' 配置文件不存在", moduleName)
	}

	// 读取模块配置
	config, err := gjson.LoadPath(configPath, gjson.Options{Safe: true})
	if err != nil {
		return nil, fmt.Errorf("读取模块配置失败: %w", err)
	}

	// 解析模块信息
	moduleInfo := &ModuleInfo{
		Name:      config.Get("name").String(),
		Author:    config.Get("author").String(),
		Version:   config.Get("version").String(),
		License:   config.Get("license").String(),
		GoVersion: config.Get("goVersion").String(),
		Path:      modulePath,
	}

	// 解析mod依赖
	modMap := config.Get("mod").Map()
	if modMap != nil {
		moduleInfo.Mod = make(map[string]string)
		for k, v := range modMap {
			moduleInfo.Mod[k] = fmt.Sprintf("%v", v)
		}
	}

	// 解析files
	filesMap := config.Get("files").Map()
	if filesMap != nil {
		moduleInfo.Files = make(map[string][]string)
		for k, v := range filesMap {
			if arr, ok := v.([]interface{}); ok {
				strArr := make([]string, len(arr))
				for i, item := range arr {
					strArr[i] = fmt.Sprintf("%v", item)
				}
				moduleInfo.Files[k] = strArr
			}
		}
	}

	return moduleInfo, nil
}

// ListModules 列出所有模块
func ListModules() ([]*ModuleInfo, error) {
	modulesDir := "modules"

	// 检查modules目录是否存在
	if _, err := os.Stat(modulesDir); os.IsNotExist(err) {
		return nil, fmt.Errorf("modules目录不存在")
	}

	// 读取modules目录
	entries, err := os.ReadDir(modulesDir)
	if err != nil {
		return nil, fmt.Errorf("读取modules目录失败: %w", err)
	}

	modules := make([]*ModuleInfo, 0, len(entries))
	for _, entry := range entries {
		// 跳过非目录和特殊目录
		if !entry.IsDir() || entry.Name() == "_" {
			continue
		}

		// 尝试扫描模块
		moduleInfo, err := ScanModule(entry.Name())
		if err != nil {
			// 跳过没有module.json的目录
			continue
		}

		modules = append(modules, moduleInfo)
	}

	return modules, nil
}

// ValidateModule 验证模块完整性
func ValidateModule(moduleName string) ([]string, []string, error) {
	moduleInfo, err := ScanModule(moduleName)
	if err != nil {
		return nil, nil, err
	}

	var warnings []string
	var errors []string

	// 检查必填字段
	if moduleInfo.Name == "" {
		errors = append(errors, "模块名称为空")
	}
	if moduleInfo.Version == "" {
		warnings = append(warnings, "版本号为空")
	}
	if moduleInfo.Author == "" {
		warnings = append(warnings, "作者信息为空")
	}

	// 检查文件是否存在
	if moduleInfo.Files != nil {
		for fileType, paths := range moduleInfo.Files {
			for _, path := range paths {
				if _, err := os.Stat(path); os.IsNotExist(err) {
					warnings = append(warnings, fmt.Sprintf("%s文件不存在: %s", fileType, path))
				}
			}
		}
	}

	return warnings, errors, nil
}
