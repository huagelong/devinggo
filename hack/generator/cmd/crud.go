// Package cmd
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package cmd

import (
	"context"
	"fmt"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"

	"devinggo/hack/generator/internal/config"
	"devinggo/hack/generator/internal/generator"
	"devinggo/hack/generator/internal/utils"
)

var (
	CrudGenerate = &gcmd.Command{
		Name:        "crud:generate",
		Brief:       "生成CRUD代码（API、Model、Controller、Logic）",
		Description: "基于数据库表的Entity模型，生成完整的CRUD代码，包括API定义、请求/响应模型、Controller和Logic层。",
		Arguments: []gcmd.Argument{
			{
				Name:  "module",
				Short: "m",
				Brief: "模块名称（例如：system）",
				IsArg: false,
			},
			{
				Name:  "table",
				Short: "t",
				Brief: "数据库表名（例如：system_user）",
				IsArg: false,
			},
			{
				Name:  "name",
				Short: "n",
				Brief: "资源中文名称（例如：用户）",
				IsArg: false,
			},
			{
				Name:  "output",
				Short: "o",
				Brief: "输出格式(text/json)",
				IsArg: false,
			},
			{
				Name:   "force",
				Short:  "f",
				Brief:  "覆盖已存在的文件",
				IsArg:  false,
				Orphan: true,
			},
			{
				Name:   "dry-run",
				Short:  "d",
				Brief:  "仅预览，不实际生成文件",
				IsArg:  false,
				Orphan: true,
			},
			{
				Name:   "config",
				Short:  "c",
				Brief:  "批量生成配置文件(generator.yaml)",
				IsArg:  false,
			},
			{
				Name:   "frontend",
				Short:  "F",
				Brief:  "同时生成前端代码（admin-ui）",
				IsArg:  false,
				Orphan: true,
			},
		},
		Func: runCrudGenerate,
		Examples: `
生成system模块的system_user表的CRUD代码：
  go run main.go crud:generate -m=system -t=system_user -n=用户

同时生成前端代码：
  go run main.go crud:generate -m=system -t=system_user -n=用户 --frontend

预览生成（不实际写入文件）：
  go run main.go crud:generate -m=system -t=system_user -n=用户 --dry-run

强制覆盖已存在的文件：
  go run main.go crud:generate -m=system -t=system_user -n=用户 --force

使用配置文件批量生成：
  go run main.go crud:generate -c=generator.yaml

JSON格式输出：
  go run main.go crud:generate -m=system -t=system_user -n=用户 -o=json
`,
	}
)

func runCrudGenerate(ctx context.Context, parser *gcmd.Parser) (err error) {
	outputFormat := utils.ParseOutputFormat(parser.GetOpt("output").String())
	
	// orphan 参数在 GetOptAll 中存在但值为空，需要检查键是否存在
	allOpts := parser.GetOptAll()
	force := false
	if _, ok := allOpts["force"]; ok {
		force = true
	}
	if _, ok := allOpts["f"]; ok {
		force = true
	}

	dryRun := false
	if _, ok := allOpts["dry-run"]; ok {
		dryRun = true
	}
	if _, ok := allOpts["d"]; ok {
		dryRun = true
	}

	configPath := parser.GetOpt("config").String()

	frontend := false
	if _, ok := allOpts["frontend"]; ok {
		frontend = true
	}
	if _, ok := allOpts["F"]; ok {
		frontend = true
	}

	if configPath != "" {
		return runBatchGenerate(ctx, configPath, force, dryRun, frontend, outputFormat)
	}

	return runSingleGenerate(ctx, parser, force, dryRun, frontend, outputFormat)
}

func runSingleGenerate(ctx context.Context, parser *gcmd.Parser, force bool, dryRun bool, frontend bool, outputFormat utils.OutputFormat) error {
	moduleName := parser.GetOpt("module", "").String()
	tableName := parser.GetOpt("table", "").String()
	chineseName := parser.GetOpt("name", "").String()

	result := utils.NewCommandResult(true, "CRUD代码生成")

	if moduleName == "" || tableName == "" || chineseName == "" {
		result.Success = false
		result.Message = "缺少必要参数：-m (模块名), -t (表名), -n (中文名) 为必填项"
		result.Print(outputFormat)
		return nil
	}

	g.Log().Infof(ctx, "开始生成CRUD代码...")
	g.Log().Infof(ctx, "  模块名: %s", moduleName)
	g.Log().Infof(ctx, "  表名: %s", tableName)
	g.Log().Infof(ctx, "  中文名: %s", chineseName)
	if frontend {
		g.Log().Infof(ctx, "  模式: 同时生成前端代码")
	}
	if dryRun {
		g.Log().Infof(ctx, "  模式: dry-run (仅预览)")
	}
	if force {
		g.Log().Infof(ctx, "  模式: force (覆盖已有文件)")
	}
	fmt.Println()

	gen, err := generator.NewCRUDGenerator(moduleName, tableName, chineseName)
	if err != nil {
		result.Success = false
		result.Message = fmt.Sprintf("创建CRUD生成器失败：%v", err)
		result.Print(outputFormat)
		return nil
	}

	gen.SetForce(force)
	gen.SetDryRun(dryRun)
	gen.SetGenerateFrontend(frontend)

	if err := gen.Generate(); err != nil {
		result.Success = false
		result.Message = fmt.Sprintf("生成CRUD代码失败：%v", err)
		result.Print(outputFormat)
		return nil
	}

	for _, f := range gen.GeneratedFiles {
		result.AddFile(f)
	}
	for _, f := range gen.SkippedFiles {
		result.AddWarning(fmt.Sprintf("跳过已存在的文件: %s", f))
	}

	result.Print(outputFormat)
	return nil
}

func runBatchGenerate(ctx context.Context, configPath string, force bool, dryRun bool, frontend bool, outputFormat utils.OutputFormat) error {
	result := utils.NewCommandResult(true, "批量CRUD代码生成")

	cfg, err := config.LoadGeneratorConfig(configPath)
	if err != nil {
		result.Success = false
		result.Message = fmt.Sprintf("加载配置文件失败：%v", err)
		result.Print(outputFormat)
		return nil
	}

	g.Log().Infof(ctx, "批量生成CRUD代码，模块: %s，共 %d 张表", cfg.Module, len(cfg.Tables))
	if frontend {
		g.Log().Infof(ctx, "  模式: 同时生成前端代码")
	}
	if dryRun {
		g.Log().Infof(ctx, "  模式: dry-run (仅预览)")
	}
	if force {
		g.Log().Infof(ctx, "  模式: force (覆盖已有文件)")
	}
	fmt.Println()

	for _, table := range cfg.Tables {
		business := table.Business
		if business == "" {
			business = table.Table
		}

		description := table.Description
		if description == "" {
			description = business
		}

		g.Log().Infof(ctx, "正在处理表: %s (%s)", table.Table, description)
		fmt.Println()

		gen, err := generator.NewCRUDGenerator(cfg.Module, table.Table, description)
		if err != nil {
			errMsg := fmt.Sprintf("表 %s 创建生成器失败：%v", table.Table, err)
			result.AddError(errMsg)
			g.Log().Errorf(ctx, errMsg)
			continue
		}

		gen.SetForce(force)
		gen.SetDryRun(dryRun)
		gen.SetGenerateFrontend(frontend)

		if err := gen.Generate(); err != nil {
			errMsg := fmt.Sprintf("表 %s 生成失败：%v", table.Table, err)
			result.AddError(errMsg)
			g.Log().Errorf(ctx, errMsg)
			continue
		}

		for _, f := range gen.GeneratedFiles {
			result.AddFile(f)
		}
		for _, f := range gen.SkippedFiles {
			result.AddWarning(fmt.Sprintf("[%s] 跳过: %s", table.Table, f))
		}

		fmt.Println()
	}

	if len(result.Errors) > 0 {
		result.Success = false
		result.Message = fmt.Sprintf("批量生成完成，%d 个失败", len(result.Errors))
	} else {
		result.Message = fmt.Sprintf("批量生成完成，共处理 %d 张表", len(cfg.Tables))
	}

	result.Print(outputFormat)
	return nil
}
