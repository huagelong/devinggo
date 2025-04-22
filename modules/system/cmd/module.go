// Package cmd
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package cmd

import (
	"context"
	"devinggo/modules/system/pkg/utils"
	"fmt"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/util/gconv"
	"strings"
	"time"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/text/gstr"
)

var (
	CreateModule = &gcmd.Command{
		Name:  "module:create",
		Brief: "创建一个新模块",
		Description: `
		创建一个新的模块，包含基本的目录结构和文件
		用法: go run main.go module:create -name 模块名称
		`,
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			CmdInit(ctx, parser)
			opts := gcmd.GetOpt("name")
			if opts == nil {
				return gerror.New("模块名称必须输入，使用 -name 参数指定")
			}

			moduleName := opts.String()
			if moduleName == "" {
				return gerror.New("模块名称不能为空")
			}

			// 检查模块名是否已存在
			modulePath := fmt.Sprintf("./modules/%s", moduleName)
			if gfile.Exists(modulePath) {
				return gerror.Newf("模块 '%s' 已存在", moduleName)
			}

			// 创建模块目录结构
			g.Log().Infof(ctx, "开始创建模块: %s", moduleName)

			// 创建主要目录
			dirs := []string{
				fmt.Sprintf("./modules/%s", moduleName),
				fmt.Sprintf("./modules/%s/api", moduleName),
				fmt.Sprintf("./modules/%s/controller", moduleName),
				fmt.Sprintf("./modules/%s/logic", moduleName),
				fmt.Sprintf("./modules/%s/logic/hook", moduleName),
				fmt.Sprintf("./modules/%s/logic/middleware", moduleName),
				fmt.Sprintf("./modules/%s/service", moduleName),
				fmt.Sprintf("./modules/%s/worker", moduleName),
			}

			for _, dir := range dirs {
				if err := gfile.Mkdir(dir); err != nil {
					return gerror.Wrapf(err, "创建目录 '%s' 失败", dir)
				}
				g.Log().Debugf(ctx, "创建目录: %s", dir)
			}

			// 创建模块文件
			if err := createModuleFiles(ctx, moduleName); err != nil {
				return err
			}

			// 使用更明显的方式输出成功信息和提示
			successMsg := fmt.Sprintf("模块 '%s' 创建成功!", moduleName)
			tipMsg := "提示: 请运行 'go run main.go migrate:up' 命令开启应用"

			// 记录到日志
			g.Log().Info(ctx, successMsg)
			g.Log().Info(ctx, tipMsg)

			// 同时直接输出到控制台，确保用户能看到
			fmt.Printf("\n%s\n%s\n\n", successMsg, tipMsg)
			return nil
		},
	}
	ExportModule = &gcmd.Command{
		Name:  "module:export",
		Brief: "导出模块文件",
		Description: `
		将模块文件导出为zip压缩包
		用法: go run main.go module:export -name 模块名称
		`,
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			CmdInit(ctx, parser)
			opts := gcmd.GetOpt("name")
			if opts == nil {
				return gerror.New("模块名称必须输入，使用 -name 参数指定")
			}

			moduleName := opts.String()
			if moduleName == "" {
				return gerror.New("模块名称不能为空")
			}

			if moduleName == "system" {
				return gerror.New("系统模块不能导出")
			}

			// 检查模块是否存在
			modulePath := fmt.Sprintf("./modules/%s", moduleName)
			if !gfile.Exists(modulePath) {
				return gerror.Newf("模块 '%s' 不存在", moduleName)
			}

			// 读取模块配置文件
			configPath := fmt.Sprintf("./modules/%s/module.json", moduleName)
			if !gfile.Exists(configPath) {
				return gerror.Newf("模块 '%s' 配置文件module.json不存在", moduleName)
			}
			config, err := gjson.LoadPath(configPath, gjson.Options{
				Safe: true,
			})
			if err != nil {
				return gerror.Wrapf(err, "读取模块配置文件失败")
			}

			// 创建临时目录
			tmpDir := utils.GetTmpDir() + "/" + "module_export_" + moduleName
			defer gfile.RemoveAll(tmpDir)

			// 复制文件到临时目录
			files := config.Get("files")
			for _, paths := range files.Map() {
				for _, path := range gconv.Strings(paths) {
					srcPath := path
					// 确保源文件存在
					if !gfile.Exists(srcPath) {
						g.Log().Warningf(ctx, "文件不存在，跳过: %s", srcPath)
						continue
					}

					// 计算目标路径
					relPath := strings.TrimPrefix(srcPath, "./")
					dstPath := gfile.Join(tmpDir, relPath)

					// 创建目标目录
					if err := gfile.Mkdir(gfile.Dir(dstPath)); err != nil {
						return gerror.Wrapf(err, "创建目录失败: %s", dstPath)
					}

					// 复制文件或目录
					if gfile.IsDir(srcPath) {
						if err := gfile.CopyDir(srcPath, dstPath); err != nil {
							return gerror.Wrapf(err, "复制目录失败: %s", srcPath)
						}
					} else {
						if err := gfile.Copy(srcPath, dstPath); err != nil {
							return gerror.Wrapf(err, "复制文件失败: %s", srcPath)
						}
					}
				}
			}
			version := config.Get("version")
			// 创建zip文件
			zipFile := fmt.Sprintf("%s.v%s.zip", moduleName, version)
			if err := utils.ZipDirectory(ctx, tmpDir, zipFile); err != nil {
				return gerror.Wrapf(err, "创建zip文件失败")
			}
			g.Log().Infof(ctx, "模块 '%s' 导出成功: %s", moduleName, zipFile)
			return nil
		},
	}
	ImportModule = &gcmd.Command{
		Name:  "module:import",
		Brief: "导入模块文件",
		Description: `
		从zip压缩包导入模块
		用法: go run main.go module:import -file 模块zip文件路径
		`,
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			CmdInit(ctx, parser)
			opts := gcmd.GetOpt("file")
			if opts == nil {
				return gerror.New("模块文件路径必须输入，使用 -file 参数指定")
			}

			zipPath := opts.String()
			if zipPath == "" {
				return gerror.New("模块文件路径不能为空")
			}

			// 检查zip文件是否存在
			if !gfile.Exists(zipPath) {
				return gerror.Newf("模块文件 '%s' 不存在", zipPath)
			}

			// 创建临时解压目录
			tmpDir := utils.GetTmpDir() + "/" + "module_import_" + gfile.Name(zipPath)
			defer gfile.RemoveAll(tmpDir)

			// 解压zip文件
			if err := utils.UnzipFile(zipPath, tmpDir); err != nil {
				return gerror.Wrapf(err, "解压模块文件失败")
			}

			// 读取模块配置文件
			configPath := gfile.Join(tmpDir, "modules/*/module.json")
			configFiles, err := gfile.Glob(configPath)
			if err != nil || len(configFiles) == 0 {
				return gerror.New("未找到模块配置文件")
			}

			config, err := gjson.LoadPath(configFiles[0], gjson.Options{
				Safe: true,
			})
			if err != nil {
				return gerror.Wrapf(err, "读取模块配置文件失败")
			}

			moduleName := config.Get("name").String()
			if moduleName == "" {
				return gerror.New("模块配置文件中未指定模块名称")
			}

			if moduleName == "system" {
				return gerror.New("系统模块不能导入")
			}

			// 目标模块路径
			modulePath := fmt.Sprintf("./modules/%s", moduleName)

			// 如果模块已存在
			if gfile.Exists(modulePath) {
				return gerror.Wrapf(err, "模块 '%s' 已存在", moduleName)
			}

			// 复制文件到目标位置
			files := config.Get("files")
			for _, paths := range files.Map() {
				for _, path := range gconv.Strings(paths) {
					srcPath := gfile.Join(tmpDir, path)
					dstPath := path

					// 确保源文件存在
					if !gfile.Exists(srcPath) {
						g.Log().Warningf(ctx, "文件不存在，跳过: %s", srcPath)
						continue
					}

					// 创建目标目录
					if err := gfile.Mkdir(gfile.Dir(dstPath)); err != nil {
						return gerror.Wrapf(err, "创建目录失败: %s", dstPath)
					}

					// 复制文件或目录
					if gfile.IsDir(srcPath) {
						if err := gfile.CopyDir(srcPath, dstPath); err != nil {
							return gerror.Wrapf(err, "复制目录失败: %s", srcPath)
						}
					} else {
						if err := gfile.Copy(srcPath, dstPath); err != nil {
							return gerror.Wrapf(err, "复制文件失败: %s", srcPath)
						}
					}
				}
			}

			g.Log().Infof(ctx, "模块 '%s' 导入成功!", moduleName)
			tipMsg := "提示: 请运行 'go run main.go migrate:up' 命令开启应用"
			g.Log().Infof(ctx, tipMsg)
			return nil
		},
	}
)

// 创建模块所需的基本文件
func createModuleFiles(ctx context.Context, moduleName string) error {
	// 首字母大写的模块名（用于类型名称）
	moduleNameCap := strings.ToUpper(moduleName[:1]) + moduleName[1:]

	// 模板数据
	tplData := g.Map{
		"moduleName":    moduleName,
		"moduleNameCap": moduleNameCap,
	}

	// 定义需要生成的文件
	files := []struct {
		tplPath  string
		filePath string
	}{
		{
			tplPath:  "modules/module.go.html",
			filePath: fmt.Sprintf("./modules/%s/module.go", moduleName),
		},
		{
			tplPath:  "modules/logic.html",
			filePath: fmt.Sprintf("./modules/%s/logic/logic.go", moduleName),
		},
		{
			tplPath:  "modules/hook_service.go.html",
			filePath: fmt.Sprintf("./modules/%s/service/hook.go", moduleName),
		},
		{
			tplPath:  "modules/middleware_service.go.html",
			filePath: fmt.Sprintf("./modules/%s/service/middleware.go", moduleName),
		},
		{
			tplPath:  "modules/hook.go.html",
			filePath: fmt.Sprintf("./modules/%s/logic/hook/hook.go", moduleName),
		},
		{
			tplPath:  "modules/api_access_log.go.html",
			filePath: fmt.Sprintf("./modules/%s/logic/hook/api_access_log.go", moduleName),
		},
		{
			tplPath:  "modules/middleware.go.html",
			filePath: fmt.Sprintf("./modules/%s/logic/middleware/middleware.go", moduleName),
		},
		{
			tplPath:  "modules/api_auth.go.html",
			filePath: fmt.Sprintf("./modules/%s/logic/middleware/api_auth.go", moduleName),
		},
		{
			tplPath:  "modules/test_api.go.html",
			filePath: fmt.Sprintf("./modules/%s/api/test.go", moduleName),
		},
		{
			tplPath:  "modules/test_controller.go.html",
			filePath: fmt.Sprintf("./modules/%s/controller/test.go", moduleName),
		},
		{
			tplPath:  "modules/worker.html",
			filePath: fmt.Sprintf("./modules/_/worker/%s.go", moduleName),
		},
		{
			tplPath:  "modules/modules.html",
			filePath: fmt.Sprintf("./modules/_/modules/%s.go", moduleName),
		},
		{
			tplPath:  "modules/logic_import.go.html",
			filePath: fmt.Sprintf("./modules/_/logic/%s.go", moduleName),
		},
	}

	// 使用g.View渲染模板并生成文件
	view := g.View()
	// 设置模板目录
	view.SetPath("resource/generate")

	// 渲染并生成每个文件
	for _, file := range files {
		content, err := view.Parse(ctx, file.tplPath, tplData)
		if err != nil {
			return gerror.Wrapf(err, "渲染模板 '%s' 失败", file.tplPath)
		}

		// 替换特殊格式的模板变量
		content = gstr.Replace(content, "{% .table.ModuleName %}", moduleName)
		content = gstr.Replace(content, "{% .table.PackageName %}", moduleName)
		content = gstr.Replace(content, "{% .tableCaseCamelName %}", moduleNameCap)
		content = gstr.Replace(content, "{% .tableCaseCamelLowerName %}", moduleName)

		// 处理条件语句，默认保留条件内容
		content = gstr.Replace(content, "{% if eq .table.Type \"single\"  %}", "")
		content = gstr.Replace(content, "{% if eq .table.Type \"tree\"  %}", "")
		content = gstr.Replace(content, "{% endif %}", "")

		if err := gfile.PutContents(file.filePath, content); err != nil {
			return gerror.Wrapf(err, "创建文件 '%s' 失败", file.filePath)
		}
		g.Log().Debugf(ctx, "创建文件: %s", file.filePath)
	}

	// 生成SQL迁移文件
	sqlfiles, err := createModuleMigrationFiles(ctx, moduleName, tplData)
	if err != nil {
		return err
	}

	// 创建模块所需的基本文件
	err = createModuleConfigFile(ctx, moduleName, sqlfiles)
	if err != nil {
		return err
	}

	return nil
}

// 创建模块配置文件
func createModuleConfigFile(ctx context.Context, moduleName string, sqlfiles []string) error {
	config := g.Map{
		"name":      moduleName,
		"author":    "devinggo",
		"version":   "1.0.0",
		"license":   "MIT",
		"goVersion": "1.23+",
		"mod":       g.Map{},
		"files": g.Map{
			"go": []string{
				fmt.Sprintf("modules/%s", moduleName),
				fmt.Sprintf("modules/_/worker/%s.go", moduleName),
				fmt.Sprintf("modules/_/modules/%s.go", moduleName),
				fmt.Sprintf("modules/_/logic/%s.go", moduleName),
			},
			"sql":    sqlfiles,
			"static": []string{},
			"other":  []string{},
		},
	}

	configPath := fmt.Sprintf("./modules/%s/module.json", moduleName)
	configContent, err := gjson.MarshalIndent(config, "", "    ")
	if err != nil {
		return gerror.Wrapf(err, "编码模块配置失败")
	}
	if err := gfile.PutContents(configPath, gconv.String(configContent)); err != nil {
		return gerror.Wrapf(err, "创建模块配置文件 '%s' 失败", configPath)
	}
	g.Log().Debugf(ctx, "创建模块配置文件: %s", configPath)
	return nil
}

// 创建模块的SQL迁移文件
func createModuleMigrationFiles(ctx context.Context, moduleName string, tplData g.Map) ([]string, error) {
	// 生成迁移文件名称
	sqlFiles := make([]string, 0)
	timezone, err := time.LoadLocation("UTC")
	if err != nil {
		return sqlFiles, gerror.Wrap(err, "加载时区失败")
	}
	version := time.Now().In(timezone).Format("20060102150405")
	name := fmt.Sprintf("%s_module", gstr.LcFirst(moduleName))

	// 确定迁移文件目录和SQL模板
	dbType := utils.GetDbType()
	directory := "resource/migrations"
	upTemplate := "sql/module_up_mysql.html"
	downTemplate := "sql/module_down_mysql.html"

	if dbType == "postgres" {
		directory = "resource/migrations_pgsql"
		upTemplate = "sql/module_up_postgres.html"
		downTemplate = "sql/module_down_postgres.html"
	}

	// 创建迁移文件
	g.Log().Infof(ctx, "开始创建模块 '%s' 的SQL迁移文件 (数据库类型: %s)", moduleName, dbType)

	// 使用g.View渲染SQL模板
	view := g.View()
	view.SetPath("resource/generate")

	// 生成up.sql文件
	upFilename := fmt.Sprintf("%s/%s_%s.up.sql", directory, version, name)
	upContent, err := view.Parse(ctx, upTemplate, tplData)
	if err != nil {
		return sqlFiles, gerror.Wrapf(err, "渲染SQL模板 '%s' 失败", upTemplate)
	}

	// 替换SQL模板中的特殊变量格式
	upContent = gstr.Replace(upContent, "{%.moduleName%}", moduleName)
	upContent = gstr.Replace(upContent, "{%.moduleNameCap%}", tplData["moduleNameCap"].(string))

	if err := gfile.PutContents(upFilename, upContent); err != nil {
		return sqlFiles, gerror.Wrapf(err, "创建SQL迁移文件 '%s' 失败", upFilename)
	}
	sqlFiles = append(sqlFiles, upFilename)
	g.Log().Debugf(ctx, "创建SQL迁移文件: %s", upFilename)

	// 生成down.sql文件
	downFilename := fmt.Sprintf("%s/%s_%s.down.sql", directory, version, name)
	downContent, err := view.Parse(ctx, downTemplate, tplData)
	if err != nil {
		return sqlFiles, gerror.Wrapf(err, "渲染SQL模板 '%s' 失败", downTemplate)
	}
	sqlFiles = append(sqlFiles, downFilename)
	// 替换SQL模板中的特殊变量格式
	downContent = gstr.Replace(downContent, "{%.moduleName%}", moduleName)
	downContent = gstr.Replace(downContent, "{%.moduleNameCap%}", tplData["moduleNameCap"].(string))
	if err := gfile.PutContents(downFilename, downContent); err != nil {
		return sqlFiles, gerror.Wrapf(err, "创建SQL迁移文件 '%s' 失败", downFilename)
	}
	g.Log().Debugf(ctx, "创建SQL迁移文件: %s", downFilename)
	// 使用更明显的方式输出SQL迁移文件创建成功信息和提示
	successMsg := fmt.Sprintf("模块 '%s' 的SQL迁移文件创建成功!", moduleName)
	// 记录到日志
	g.Log().Info(ctx, successMsg)

	return sqlFiles, nil
}
