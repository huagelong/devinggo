package generator

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"devinggo/hack/generator/internal/scanner"
	"devinggo/hack/generator/internal/utils"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/util/gconv"
)

// getString 从 map 中获取字符串值
func getString(m map[string]interface{}, key string) string {
	if v, ok := m[key].(string); ok {
		return v
	}
	return ""
}

// getBool 从 map 中获取布尔值
func getBool(m map[string]interface{}, key string) bool {
	if v, ok := m[key].(bool); ok {
		return v
	}
	return false
}

// ModuleExporter 模块导出器
type ModuleExporter struct {
	ctx context.Context
}

// NewModuleExporter 创建模块导出器
func NewModuleExporter(ctx context.Context) *ModuleExporter {
	return &ModuleExporter{ctx: ctx}
}

// Export 导出模块为zip包
func (e *ModuleExporter) Export(moduleName string) (string, error) {
	// 扫描模块信息
	moduleInfo, err := scanner.ScanModule(moduleName)
	if err != nil {
		return "", err
	}

	// 创建临时目录
	tmpDir := filepath.Join(utils.GetTmpDir(), "module_export_"+moduleName)
	defer os.RemoveAll(tmpDir)

	// 复制文件到临时目录
	for fileType, paths := range moduleInfo.Files {
		for _, path := range paths {
			// 确保源文件存在
			if _, err := os.Stat(path); os.IsNotExist(err) {
				g.Log().Warningf(e.ctx, "[%s]文件不存在，跳过: %s", fileType, path)
				continue
			}

			// 计算目标路径
			dstPath := filepath.Join(tmpDir, path)

			// 创建目标目录
			if err := utils.EnsureDir(filepath.Dir(dstPath)); err != nil {
				return "", fmt.Errorf("创建目录失败: %w", err)
			}

			// 复制文件或目录
			if gfile.IsDir(path) {
				if err := gfile.CopyDir(path, dstPath); err != nil {
					return "", fmt.Errorf("复制目录失败: %w", err)
				}
			} else {
				if err := utils.CopyFile(path, dstPath); err != nil {
					return "", fmt.Errorf("复制文件失败: %w", err)
				}
			}
		}
	}

	// 清理敏感信息
	fmt.Printf("\n🔍 正在扫描并替换敏感信息...\n")
	sanitizer := utils.NewSanitizer(nil)
	if err := sanitizer.SanitizeDirectoryWithReport(tmpDir); err != nil {
		g.Log().Warningf(e.ctx, "清理敏感信息时出现警告: %v", err)
		// 继续执行，不中断导出流程
	}

	// 创建zip文件
	zipFile := fmt.Sprintf("%s.v%s.zip", moduleName, moduleInfo.Version)
	if err := utils.ZipDirectory(e.ctx, tmpDir, zipFile); err != nil {
		return "", fmt.Errorf("创建zip文件失败: %w", err)
	}

	return zipFile, nil
}

// ModuleImporter 模块导入器
type ModuleImporter struct {
	ctx context.Context
}

// NewModuleImporter 创建模块导入器
func NewModuleImporter(ctx context.Context) *ModuleImporter {
	return &ModuleImporter{ctx: ctx}
}

// Import 从zip包导入模块
func (i *ModuleImporter) Import(zipPath string) (string, error) {
	// 检查zip文件是否存在
	if !utils.PathExists(zipPath) {
		return "", fmt.Errorf("模块文件 '%s' 不存在", zipPath)
	}

	// 创建临时解压目录
	tmpDir := filepath.Join(utils.GetTmpDir(), "module_import_"+filepath.Base(zipPath))
	defer os.RemoveAll(tmpDir)

	// 解压zip文件
	if err := utils.UnzipFile(zipPath, tmpDir); err != nil {
		return "", fmt.Errorf("解压模块文件失败: %w", err)
	}

	// 读取模块配置文件（优先读取 .module.yaml）
	configPattern := filepath.Join(tmpDir, "modules", "*", ".module.yaml")
	configFiles, err := filepath.Glob(configPattern)
	if err != nil || len(configFiles) == 0 {
		// 尝试读取旧的 module.json
		configPattern = filepath.Join(tmpDir, "modules", "*", "module.json")
		configFiles, err = filepath.Glob(configPattern)
		if err != nil || len(configFiles) == 0 {
			return "", fmt.Errorf("未找到模块配置文件")
		}
	}

	// 读取配置
	config, err := gjson.LoadPath(configFiles[0], gjson.Options{Safe: true})
	if err != nil {
		return "", fmt.Errorf("读取模块配置失败: %w", err)
	}

	moduleName := config.Get("name").String()
	if moduleName == "" {
		return "", fmt.Errorf("模块配置文件中未指定模块名称")
	}

	// 检查模块是否已存在
	modulePath := filepath.Join("modules", moduleName)
	if utils.PathExists(modulePath) {
		return "", fmt.Errorf("模块 '%s' 已存在", moduleName)
	}

	// 解析并执行 preInstall 钩子
	hookExecutor := utils.NewHookExecutor()
	preInstallHooks := i.parseHooks(config.Get("hooks.preInstall"))
	if len(preInstallHooks) > 0 {
		fmt.Println("\n📋 执行安装前钩子...")
		results := hookExecutor.ExecuteHooks(preInstallHooks, "preInstall")
		for _, result := range results {
			if !result.Success {
				return "", fmt.Errorf("preInstall 钩子 '%s' 执行失败: %s", result.Name, result.Error)
			}
		}
	}

	// 收集需要迁移的 SQL 文件和静态资源部署规则
	var sqlFiles []string
	staticDeploy := i.parseStaticDeploy(config.Get("staticDeploy"))

	// 复制文件到目标位置
	filesMap := config.Get("files").Map()
	if filesMap != nil {
		for fileType, v := range filesMap {
			if arr, ok := v.([]interface{}); ok {
				for _, item := range arr {
					path := fmt.Sprintf("%v", item)
					srcPath := filepath.Join(tmpDir, path)
					dstPath := path

					// 确保源文件存在
					if !utils.PathExists(srcPath) {
						g.Log().Warningf(i.ctx, "文件不存在，跳过: %s", srcPath)
						continue
					}

					// 创建目标目录
					if err := utils.EnsureDir(filepath.Dir(dstPath)); err != nil {
						return "", fmt.Errorf("创建目录失败: %w", err)
					}

					// 复制文件或目录
					if gfile.IsDir(srcPath) {
						if err := gfile.CopyDir(srcPath, dstPath); err != nil {
							return "", fmt.Errorf("复制目录失败: %w", err)
						}
					} else {
						if err := utils.CopyFile(srcPath, dstPath); err != nil {
							return "", fmt.Errorf("复制文件失败: %w", err)
						}
					}

					// 收集 SQL 文件
					if fileType == "sql" || strings.HasSuffix(path, ".sql") {
						sqlFiles = append(sqlFiles, path)
					}
				}
			}
		}
	}

	// 执行静态资源部署
	if staticDeploy != nil && staticDeploy.Enabled {
		if err := i.executeStaticDeploy(staticDeploy, tmpDir, hookExecutor); err != nil {
			return "", fmt.Errorf("静态资源部署失败: %w", err)
		}
	}

	// 解析并执行 postInstall 钩子
	postInstallHooks := i.parseHooks(config.Get("hooks.postInstall"))
	if len(postInstallHooks) > 0 {
		fmt.Println("\n📋 执行安装后钩子...")
		hookExecutor.ExecuteHooks(postInstallHooks, "postInstall")
	}

	// 处理数据库迁移提示
	if len(sqlFiles) > 0 {
		fmt.Printf("\n📊 发现 %d 个 SQL 迁移文件:\n", len(sqlFiles))
		for _, sqlFile := range sqlFiles {
			fmt.Printf("  • %s\n", sqlFile)
		}
		fmt.Println("\n💡 提示: 请运行以下命令执行数据库迁移:")
		fmt.Println("   go run main.go migrate:up")
	}

	// 执行安装验证
	fmt.Println("\n🔍 验证安装...")
	if errors := i.validateInstall(moduleName, filesMap); len(errors) > 0 {
		fmt.Println("  ⚠️  安装验证发现问题:")
		for _, err := range errors {
			fmt.Printf("    - %s\n", err)
		}
	} else {
		fmt.Println("  ✅ 安装验证通过，所有文件已正确创建")
	}

	return moduleName, nil
}

// validateInstall 验证模块安装是否完整
func (i *ModuleImporter) validateInstall(moduleName string, filesMap map[string]interface{}) []string {
	var errors []string

	if filesMap == nil {
		return nil
	}

	// 检查声明的文件是否存在
	for fileType, v := range filesMap {
		if arr, ok := v.([]interface{}); ok {
			for _, item := range arr {
				path := fmt.Sprintf("%v", item)

				if !utils.PathExists(path) {
					errors = append(errors, fmt.Sprintf("%s文件缺失: %s", fileType, path))
				}
			}
		}
	}

	// 检查模块目录是否存在
	modulePath := filepath.Join("modules", moduleName)
	if !utils.PathExists(modulePath) {
		errors = append(errors, fmt.Sprintf("模块目录不存在: %s", modulePath))
	}

	// 检查必要的子目录
	requiredDirs := []string{
		filepath.Join(modulePath, "api"),
		filepath.Join(modulePath, "logic"),
		filepath.Join(modulePath, "controller"),
	}

	for _, dir := range requiredDirs {
		if !utils.PathExists(dir) {
			errors = append(errors, fmt.Sprintf("必要目录缺失: %s", dir))
		}
	}

	return errors
}

// parseHooks 解析钩子配置
func (i *ModuleImporter) parseHooks(hookData interface{}) []utils.HookCommand {
	if hookData == nil {
		return nil
	}

	// 使用 gconv 转换为 []interface{}
	hooksArray := gconv.SliceAny(hookData)

	if len(hooksArray) == 0 {
		return nil
	}

	hooks := make([]utils.HookCommand, 0, len(hooksArray))
	for _, h := range hooksArray {
		hookMap, ok := h.(map[string]interface{})
		if !ok {
			continue
		}

		hook := utils.HookCommand{
			Name:        getString(hookMap, "name"),
			Command:     getString(hookMap, "command"),
			WorkDir:     getString(hookMap, "workDir"),
			IgnoreError: getBool(hookMap, "ignoreError"),
		}

		// 解析环境变量
		if envData, ok := hookMap["env"].(map[string]interface{}); ok {
			hook.Env = make(map[string]string)
			for k, v := range envData {
				if strVal, ok := v.(string); ok {
					hook.Env[k] = strVal
				}
			}
		}

		if hook.Command != "" {
			hooks = append(hooks, hook)
		}
	}

	return hooks
}

// parseStaticDeploy 解析静态资源部署配置
func (i *ModuleImporter) parseStaticDeploy(deployData interface{}) *StaticDeployConfig {
	if deployData == nil {
		return nil
	}

	// 使用 gconv 转换为 map[string]interface{}
	deployMap := gconv.Map(deployData)

	if len(deployMap) == 0 {
		return nil
	}

	config := &StaticDeployConfig{
		Enabled: getBool(deployMap, "enabled"),
		Rules:   []StaticDeployRule{},
	}

	// 解析 rules
	if rulesData, ok := deployMap["rules"].([]interface{}); ok {
		for _, r := range rulesData {
			ruleMap, ok := r.(map[string]interface{})
			if !ok {
				continue
			}

			rule := StaticDeployRule{
				Source:    getString(ruleMap, "source"),
				Target:    getString(ruleMap, "target"),
				Method:    getString(ruleMap, "method"),
				Overwrite: getBool(ruleMap, "overwrite"),
			}

			if rule.Source != "" && rule.Target != "" {
				config.Rules = append(config.Rules, rule)
			}
		}
	}

	return config
}

// executeStaticDeploy 执行静态资源部署
func (i *ModuleImporter) executeStaticDeploy(config *StaticDeployConfig, tmpDir string, hookExecutor *utils.HookExecutor) error {
	if !config.Enabled || len(config.Rules) == 0 {
		return nil
	}

	fmt.Println("\n📦 执行静态资源部署...")

	for _, rule := range config.Rules {
		fmt.Printf("  • %s -> %s (%s)\n", rule.Source, rule.Target, rule.Method)

		srcPath := filepath.Join(tmpDir, rule.Source)

		// 确保源路径存在
		if !utils.PathExists(srcPath) {
			g.Log().Warningf(i.ctx, "静态资源源路径不存在，跳过: %s", srcPath)
			continue
		}

		// 创建目标目录
		if err := utils.EnsureDir(rule.Target); err != nil {
			return fmt.Errorf("创建目标目录失败: %w", err)
		}

		switch rule.Method {
		case "copy":
			if gfile.IsDir(srcPath) {
				if err := gfile.CopyDir(srcPath, rule.Target); err != nil {
					return fmt.Errorf("复制目录失败: %w", err)
				}
			} else {
				if err := utils.CopyFile(srcPath, rule.Target); err != nil {
					return fmt.Errorf("复制文件失败: %w", err)
				}
			}

		case "symlink":
			// 删除已存在的链接
			if utils.PathExists(rule.Target) {
				os.Remove(rule.Target)
			}
			if err := os.Symlink(srcPath, rule.Target); err != nil {
				g.Log().Warningf(i.ctx, "创建符号链接失败（可能需要管理员权限）: %v", err)
				// 回退到复制
				if gfile.IsDir(srcPath) {
					gfile.CopyDir(srcPath, rule.Target)
				} else {
					utils.CopyFile(srcPath, rule.Target)
				}
			}

		case "merge":
			if gfile.IsDir(srcPath) {
				// 合并目录：复制不存在的文件
				if err := i.mergeDirectories(srcPath, rule.Target, rule.Overwrite); err != nil {
					return fmt.Errorf("合并目录失败: %w", err)
				}
			}
		}
	}

	return nil
}

// mergeDirectories 合并两个目录
func (i *ModuleImporter) mergeDirectories(src, dst string, overwrite bool) error {
	return filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 计算相对路径
		relPath, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}

		dstPath := filepath.Join(dst, relPath)

		if info.IsDir() {
			// 创建目标目录
			return utils.EnsureDir(dstPath)
		}

		// 检查目标文件是否存在
		if utils.PathExists(dstPath) && !overwrite {
			// 文件已存在，跳过
			return nil
		}

		// 复制文件
		return utils.CopyFile(path, dstPath)
	})
}

// StaticDeployConfig 静态资源部署配置
type StaticDeployConfig struct {
	Enabled bool
	Rules   []StaticDeployRule
}

// StaticDeployRule 静态资源部署规则
type StaticDeployRule struct {
	Source    string
	Target    string
	Method    string // copy, symlink, merge
	Overwrite bool
}
