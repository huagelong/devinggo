// Package generator
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package generator

import (
	"context"
	"fmt"
	"path/filepath"
	"strings"
	"time"

	"devinggo/hack/generator/internal/config"
	"devinggo/hack/generator/internal/utils"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/text/gstr"
)

// ModuleCreator 模块创建器
type ModuleCreator struct {
	ctx context.Context
}

// NewModuleCreator 创建模块创建器
func NewModuleCreator(ctx context.Context) *ModuleCreator {
	return &ModuleCreator{ctx: ctx}
}

// Create 创建新模块
func (c *ModuleCreator) Create(moduleName string) error {
	// 检查模块名是否已存在
	modulePath := filepath.Join("modules", moduleName)
	if gfile.Exists(modulePath) {
		return fmt.Errorf("模块 '%s' 已存在", moduleName)
	}

	g.Log().Infof(c.ctx, "开始创建模块: %s", moduleName)

	// 创建主要目录
	dirs := []string{
		filepath.Join("modules", moduleName),
		filepath.Join("modules", moduleName, "api"),
		filepath.Join("modules", moduleName, "controller"),
		filepath.Join("modules", moduleName, "logic"),
		filepath.Join("modules", moduleName, "logic", "hook"),
		filepath.Join("modules", moduleName, "logic", "middleware"),
		filepath.Join("modules", moduleName, "logic", moduleName),
		filepath.Join("modules", moduleName, "service"),
		filepath.Join("modules", moduleName, "worker"),
	}

	for _, dir := range dirs {
		if err := utils.EnsureDir(dir); err != nil {
			return fmt.Errorf("创建目录 '%s' 失败: %w", dir, err)
		}
		g.Log().Debugf(c.ctx, "创建目录: %s", dir)
	}

	// 创建模块文件
	if err := c.createModuleFiles(moduleName); err != nil {
		return err
	}

	return nil
}

// createModuleFiles 创建模块文件
func (c *ModuleCreator) createModuleFiles(moduleName string) error {
	// 首字母大写的模块名（用于类型名称）
	moduleNameCap := gstr.UcFirst(moduleName)

	// 模板数据
	tplData := g.Map{
		"moduleName":    moduleName,
		"moduleNameCap": moduleNameCap,
		"date":          time.Now().Format("2006-01-02 15:04:05"),
	}

	// 定义需要生成的文件
	files := []struct {
		tplPath  string
		filePath string
	}{
		{
			tplPath:  "module.go.tpl",
			filePath: filepath.Join("modules", moduleName, "module.go"),
		},
		{
			tplPath:  "logic.tpl",
			filePath: filepath.Join("modules", moduleName, "logic", "logic.go"),
		},
		{
			tplPath:  "hook_service.go.tpl",
			filePath: filepath.Join("modules", moduleName, "service", "hook.go"),
		},
		{
			tplPath:  "middleware_service.go.tpl",
			filePath: filepath.Join("modules", moduleName, "service", "middleware.go"),
		},
		{
			tplPath:  "mod_service.go.tpl",
			filePath: filepath.Join("modules", moduleName, "service", moduleName+".go"),
		},
		{
			tplPath:  "mod.go.tpl",
			filePath: filepath.Join("modules", moduleName, "logic", moduleName, moduleName+".go"),
		},
		{
			tplPath:  "hook.go.tpl",
			filePath: filepath.Join("modules", moduleName, "logic", "hook", "hook.go"),
		},
		{
			tplPath:  "api_access_log.go.tpl",
			filePath: filepath.Join("modules", moduleName, "logic", "hook", "api_access_log.go"),
		},
		{
			tplPath:  "middleware.go.tpl",
			filePath: filepath.Join("modules", moduleName, "logic", "middleware", "middleware.go"),
		},
		{
			tplPath:  "api_auth.go.tpl",
			filePath: filepath.Join("modules", moduleName, "logic", "middleware", "api_auth.go"),
		},
		{
			tplPath:  "test_api.go.tpl",
			filePath: filepath.Join("modules", moduleName, "api", "test.go"),
		},
		{
			tplPath:  "test_controller.go.tpl",
			filePath: filepath.Join("modules", moduleName, "controller", "test.go"),
		},
		{
			tplPath:  "worker.tpl",
			filePath: filepath.Join("modules", "bootstrap", "worker", moduleName+".go"),
		},
		{
			tplPath:  "modules.tpl",
			filePath: filepath.Join("modules", "bootstrap", "modules", moduleName+".go"),
		},
		{
			tplPath:  "logic_import.go.tpl",
			filePath: filepath.Join("modules", "bootstrap", "logic", moduleName+".go"),
		},
	}

	// 使用g.View渲染模板并生成文件
	view := g.View()
	// 设置模板目录
	templatePath := filepath.Join("hack", "generator", "templates", "module")
	view.SetPath(templatePath)

	// 渲染并生成每个文件
	for _, file := range files {
		content, err := view.Parse(c.ctx, file.tplPath, tplData)
		if err != nil {
			return fmt.Errorf("渲染模板 '%s' 失败: %w", file.tplPath, err)
		}

		if err := gfile.PutContents(file.filePath, content); err != nil {
			return fmt.Errorf("创建文件 '%s' 失败: %w", file.filePath, err)
		}
		g.Log().Debugf(c.ctx, "创建文件: %s", file.filePath)
	}

	// 生成SQL迁移文件
	sqlFiles, err := c.createModuleMigrationFiles(moduleName, tplData)
	if err != nil {
		return err
	}

	// 创建模块配置文件
	err = c.createModuleConfigFile(moduleName, sqlFiles)
	if err != nil {
		return err
	}

	return nil
}

// createModuleMigrationFiles 创建模块的SQL迁移文件
func (c *ModuleCreator) createModuleMigrationFiles(moduleName string, tplData g.Map) ([]string, error) {
	sqlFiles := make([]string, 0)

	// 生成迁移文件名称
	timezone, err := time.LoadLocation("UTC")
	if err != nil {
		return sqlFiles, fmt.Errorf("加载时区失败: %w", err)
	}
	version := time.Now().In(timezone).Format("20060102150405")
	name := fmt.Sprintf("%s_module", gstr.LcFirst(moduleName))

	// 确定迁移文件目录
	directory := "resource/migrations"

	// 创建迁移文件
	g.Log().Infof(c.ctx, "开始创建模块 '%s' 的SQL迁移文件 (PostgreSQL)", moduleName)

	// 使用g.View渲染SQL模板
	view := g.View()
	templatePath := filepath.Join("hack", "generator", "templates", "module")
	view.SetPath(templatePath)

	// 生成up.sql文件
	upFilename := filepath.Join(directory, fmt.Sprintf("%s_%s.up.sql", version, name))
	upContent, err := view.Parse(c.ctx, "module_up_postgres.sql.tpl", tplData)
	if err != nil {
		return sqlFiles, fmt.Errorf("渲染SQL模板失败: %w", err)
	}

	if putErr := gfile.PutContents(upFilename, upContent); putErr != nil {
		return sqlFiles, fmt.Errorf("创建SQL迁移文件 '%s' 失败: %w", upFilename, putErr)
	}
	sqlFiles = append(sqlFiles, upFilename)
	g.Log().Debugf(c.ctx, "创建SQL迁移文件: %s", upFilename)

	// 生成down.sql文件
	downFilename := filepath.Join(directory, fmt.Sprintf("%s_%s.down.sql", version, name))
	downContent, err := view.Parse(c.ctx, "module_down_postgres.sql.tpl", tplData)
	if err != nil {
		return sqlFiles, fmt.Errorf("渲染SQL模板失败: %w", err)
	}

	if err := gfile.PutContents(downFilename, downContent); err != nil {
		return sqlFiles, fmt.Errorf("创建SQL迁移文件 '%s' 失败: %w", downFilename, err)
	}
	sqlFiles = append(sqlFiles, downFilename)
	g.Log().Debugf(c.ctx, "创建SQL迁移文件: %s", downFilename)

	return sqlFiles, nil
}

// createModuleConfigFile 创建模块配置文件
func (c *ModuleCreator) createModuleConfigFile(moduleName string, sqlFiles []string) error {
	// 使用新的配置结构
	moduleConfig := config.DefaultModuleConfig(moduleName)

	// 设置文件列表
	moduleConfig.Files.Go = []string{
		fmt.Sprintf("modules/%s", moduleName),
		fmt.Sprintf("modules/bootstrap/worker/%s.go", moduleName),
		fmt.Sprintf("modules/bootstrap/modules/%s.go", moduleName),
		fmt.Sprintf("modules/bootstrap/logic/%s.go", moduleName),
	}
	moduleConfig.Files.SQL = sqlFiles

	// 设置描述
	moduleConfig.Description = fmt.Sprintf("%s 模块 - 由DevingGo代码生成器创建", gstr.UcFirst(moduleName))

	// 同时生成 .module.yaml (新格式) 和 module.json (向后兼容)
	parser := config.NewModuleConfigParser()

	// 保存YAML格式
	modulePath := filepath.Join("modules", moduleName)
	if err := parser.SaveConfig(moduleConfig, modulePath, "yaml"); err != nil {
		return fmt.Errorf("创建.module.yaml文件失败: %w", err)
	}
	g.Log().Debugf(c.ctx, "创建模块配置文件: %s", filepath.Join(modulePath, ".module.yaml"))

	// 同时保存JSON格式（向后兼容）
	if err := parser.SaveConfig(moduleConfig, modulePath, "json"); err != nil {
		return fmt.Errorf("创建module.json文件失败: %w", err)
	}
	g.Log().Debugf(c.ctx, "创建模块配置文件: %s", filepath.Join(modulePath, "module.json"))

	return nil
}

// ModuleCloner 模块克隆器
type ModuleCloner struct {
	ctx context.Context
}

// NewModuleCloner 创建模块克隆器
func NewModuleCloner(ctx context.Context) *ModuleCloner {
	return &ModuleCloner{ctx: ctx}
}

// Clone 从现有模块克隆
func (c *ModuleCloner) Clone(sourceModule, targetModule string) error {
	// 检查源模块是否存在
	sourceModulePath := filepath.Join("modules", sourceModule)
	if !gfile.Exists(sourceModulePath) {
		return fmt.Errorf("源模块 '%s' 不存在", sourceModule)
	}

	// 检查目标模块是否已存在
	targetModulePath := filepath.Join("modules", targetModule)
	if gfile.Exists(targetModulePath) {
		return fmt.Errorf("目标模块 '%s' 已存在", targetModule)
	}

	g.Log().Infof(c.ctx, "开始克隆模块: %s -> %s", sourceModule, targetModule)

	// 复制整个模块目录
	if err := gfile.CopyDir(sourceModulePath, targetModulePath); err != nil {
		return fmt.Errorf("复制模块目录失败: %w", err)
	}

	// 替换所有文件中的模块名称
	if err := c.replaceModuleName(targetModulePath, sourceModule, targetModule); err != nil {
		return fmt.Errorf("替换模块名称失败: %w", err)
	}

	// 复制并更新 modules/bootstrap/ 下的文件
	if err := c.cloneBootstrapFiles(sourceModule, targetModule); err != nil {
		return fmt.Errorf("克隆引导文件失败: %w", err)
	}

	g.Log().Infof(c.ctx, "模块 '%s' 克隆成功！", targetModule)
	return nil
}

// replaceModuleName 替换文件中的模块名称
func (c *ModuleCloner) replaceModuleName(dir, oldName, newName string) error {
	// 首字母大写的模块名
	oldNameCap := gstr.UcFirst(oldName)
	newNameCap := gstr.UcFirst(newName)

	// 遍历所有.go文件
	files, err := gfile.ScanDirFile(dir, "*.go", true)
	if err != nil {
		return err
	}

	for _, file := range files {
		content := gfile.GetContents(file)

		// 替换包名和类型名
		content = strings.ReplaceAll(content, oldName, newName)
		content = strings.ReplaceAll(content, oldNameCap, newNameCap)

		// 写回文件
		if err := gfile.PutContents(file, content); err != nil {
			return fmt.Errorf("更新文件 '%s' 失败: %w", file, err)
		}
		g.Log().Debugf(c.ctx, "更新文件: %s", file)
	}

	// 更新module.json文件
	configPath := filepath.Join(dir, "module.json")
	if gfile.Exists(configPath) {
		configObj, err := gjson.LoadPath(configPath, gjson.Options{Safe: true})
		if err != nil {
			return fmt.Errorf("读取模块配置失败: %w", err)
		}

		// 更新配置中的模块名
		configObj.Set("name", newName)

		// 更新文件路径
		if files := configObj.Get("files"); files != nil {
			for fileType, paths := range files.Map() {
				newPaths := make([]string, 0)
				// 确保paths是切片类型
				if pathList, ok := paths.([]interface{}); ok {
					for _, path := range pathList {
						pathStr := fmt.Sprintf("%v", path)
						newPath := strings.ReplaceAll(pathStr, oldName, newName)
						newPaths = append(newPaths, newPath)
					}
				}
				configObj.Set(fmt.Sprintf("files.%s", fileType), newPaths)
			}
		}

		// 保存配置
		configContent, err := gjson.MarshalIndent(configObj, "", "    ")
		if err != nil {
			return fmt.Errorf("编码模块配置失败: %w", err)
		}

		if err := gfile.PutContents(configPath, string(configContent)); err != nil {
			return fmt.Errorf("保存模块配置失败: %w", err)
		}
	}

	return nil
}

// cloneBootstrapFiles 克隆引导文件
func (c *ModuleCloner) cloneBootstrapFiles(sourceModule, targetModule string) error {
	bootstrapFiles := []string{
		filepath.Join("modules", "bootstrap", "worker", sourceModule+".go"),
		filepath.Join("modules", "bootstrap", "modules", sourceModule+".go"),
		filepath.Join("modules", "bootstrap", "logic", sourceModule+".go"),
	}

	oldNameCap := gstr.UcFirst(sourceModule)
	newNameCap := gstr.UcFirst(targetModule)

	for _, sourceFile := range bootstrapFiles {
		if !gfile.Exists(sourceFile) {
			g.Log().Warningf(c.ctx, "引导文件不存在，跳过: %s", sourceFile)
			continue
		}

		// 目标文件路径
		targetFile := strings.ReplaceAll(sourceFile, sourceModule, targetModule)

		// 读取并替换内容
		content := gfile.GetContents(sourceFile)
		content = strings.ReplaceAll(content, sourceModule, targetModule)
		content = strings.ReplaceAll(content, oldNameCap, newNameCap)

		// 写入目标文件
		if err := gfile.PutContents(targetFile, content); err != nil {
			return fmt.Errorf("创建引导文件 '%s' 失败: %w", targetFile, err)
		}
		g.Log().Debugf(c.ctx, "创建引导文件: %s", targetFile)
	}

	return nil
}
