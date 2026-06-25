package cmd

import (
	"context"
	"fmt"

	"devinggo/hack/generator/internal/generator"
	"devinggo/hack/generator/internal/scanner"
	"devinggo/hack/generator/internal/utils"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	// ModuleCreate 创建模块命令
	ModuleCreate = &gcmd.Command{
		Name:  "module:create",
		Brief: "创建新模块",
		Description: `
创建一个新的模块，包含基本的目录结构和文件

用法:
  go run hack/generator/main.go module:create -name 模块名称 [-output text|json]
`,
		Arguments: []gcmd.Argument{
			{
				Name:   "name",
				Short:  "n",
				Brief:  "模块名称",
				IsArg:  false,
				Orphan: false,
			},
			{
				Name:   "output",
				Short:  "o",
				Brief:  "输出格式(text/json)",
				IsArg:  false,
				Orphan: false,
			},
		},
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			outputFormat := utils.ParseOutputFormat(parser.GetOpt("output").String())
			moduleName := parser.GetOpt("name").String()

			if moduleName == "" {
				result := utils.NewCommandResult(false, "缺少必填参数: -name")
				result.Print(outputFormat)
				return nil
			}

			g.Log().Infof(ctx, "开始创建模块: %s", moduleName)

			// 创建模块创建器
			creator := generator.NewModuleCreator(ctx)
			if err := creator.Create(moduleName); err != nil {
				result := utils.NewCommandResult(false, fmt.Sprintf("创建模块失败: %v", err))
				result.Print(outputFormat)
				return nil
			}

			g.Log().Infof(ctx, "模块 '%s' 创建成功！", moduleName)
			result := utils.NewCommandResult(true, fmt.Sprintf("模块 '%s' 创建成功", moduleName))
			result.Print(outputFormat)
			return nil
		},
	}

	// ModuleClone 克隆模块命令
	ModuleClone = &gcmd.Command{
		Name:  "module:clone",
		Brief: "从现有模块克隆新模块",
		Description: `
从现有模块快速克隆一个新模块

用法:
  go run hack/generator/main.go module:clone -source 源模块 -target 目标模块 [-output text|json]
`,
		Arguments: []gcmd.Argument{
			{
				Name:   "source",
				Short:  "s",
				Brief:  "源模块名称",
				IsArg:  false,
				Orphan: false,
			},
			{
				Name:   "target",
				Short:  "t",
				Brief:  "目标模块名称",
				IsArg:  false,
				Orphan: false,
			},
			{
				Name:   "output",
				Short:  "o",
				Brief:  "输出格式(text/json)",
				IsArg:  false,
				Orphan: false,
			},
		},
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			outputFormat := utils.ParseOutputFormat(parser.GetOpt("output").String())
			sourceModule := parser.GetOpt("source").String()
			targetModule := parser.GetOpt("target").String()

			if sourceModule == "" {
				result := utils.NewCommandResult(false, "缺少必填参数: -source")
				result.Print(outputFormat)
				return nil
			}
			if targetModule == "" {
				result := utils.NewCommandResult(false, "缺少必填参数: -target")
				result.Print(outputFormat)
				return nil
			}

			// 禁止克隆系统模块
			if sourceModule == "system" {
				result := utils.NewCommandResult(false, "系统模块不能被克隆")
				result.Print(outputFormat)
				return nil
			}
			if targetModule == "system" {
				result := utils.NewCommandResult(false, "不能创建名为 'system' 的模块")
				result.Print(outputFormat)
				return nil
			}

			g.Log().Infof(ctx, "开始克隆模块: %s -> %s", sourceModule, targetModule)

			// 创建模块克隆器
			cloner := generator.NewModuleCloner(ctx)
			if err := cloner.Clone(sourceModule, targetModule); err != nil {
				result := utils.NewCommandResult(false, fmt.Sprintf("克隆模块失败: %v", err))
				result.Print(outputFormat)
				return nil
			}

			g.Log().Infof(ctx, "模块 '%s' 克隆成功！", targetModule)
			result := utils.NewCommandResult(true, fmt.Sprintf("模块 '%s' 克隆成功（源模块: %s）", targetModule, sourceModule))
			result.Print(outputFormat)
			return nil
		},
	}

	// ModuleExport 导出模块命令
	ModuleExport = &gcmd.Command{
		Name:  "module:export",
		Brief: "导出模块为zip包",
		Description: `
将模块文件导出为zip压缩包
用法: go run hack/generator/main.go module:export -name 模块名称 [-output text|json]
`,
		Arguments: []gcmd.Argument{
			{
				Name:   "name",
				Short:  "n",
				Brief:  "模块名称",
				IsArg:  false,
				Orphan: false,
			},
			{
				Name:   "output",
				Short:  "o",
				Brief:  "输出格式(text/json)",
				IsArg:  false,
				Orphan: false,
			},
		},
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			outputFormat := utils.ParseOutputFormat(parser.GetOpt("output").String())
			moduleName := parser.GetOpt("name").String()
			if moduleName == "" {
				result := utils.NewCommandResult(false, "缺少必填参数: -name")
				result.Print(outputFormat)
				return nil
			}

			// 禁止导出系统模块
			if moduleName == "system" {
				result := utils.NewCommandResult(false, "系统模块不能导出")
				result.Print(outputFormat)
				return nil
			}

			g.Log().Infof(ctx, "开始导出模块: %s", moduleName)

			// 创建导出器
			exporter := generator.NewModuleExporter(ctx)
			zipFile, err := exporter.Export(moduleName)
			if err != nil {
				result := utils.NewCommandResult(false, fmt.Sprintf("导出模块失败: %v", err))
				result.Print(outputFormat)
				return nil
			}

			g.Log().Infof(ctx, "模块 '%s' 导出成功: %s", moduleName, zipFile)
			result := utils.NewCommandResult(true, fmt.Sprintf("模块 '%s' 导出成功", moduleName))
			result.AddFile(zipFile)
			result.Print(outputFormat)
			return nil
		},
	}

	// ModuleImport 导入模块命令
	ModuleImport = &gcmd.Command{
		Name:  "module:import",
		Brief: "从zip包导入模块",
		Description: `
从zip压缩包导入模块

用法:
  go run hack/generator/main.go module:import -file 模块zip文件路径 [-output text|json]
`,
		Arguments: []gcmd.Argument{
			{
				Name:   "file",
				Short:  "f",
				Brief:  "模块zip文件路径",
				IsArg:  false,
				Orphan: false,
			},
			{
				Name:   "output",
				Short:  "o",
				Brief:  "输出格式(text/json)",
				IsArg:  false,
				Orphan: false,
			},
		},
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			outputFormat := utils.ParseOutputFormat(parser.GetOpt("output").String())
			zipPath := parser.GetOpt("file").String()

			if zipPath == "" {
				result := utils.NewCommandResult(false, "缺少必填参数: -file")
				result.Print(outputFormat)
				return nil
			}

			g.Log().Infof(ctx, "开始导入模块: %s", zipPath)

			// 创建导入器
			importer := generator.NewModuleImporter(ctx)
			moduleName, err := importer.Import(zipPath)
			if err != nil {
				result := utils.NewCommandResult(false, fmt.Sprintf("导入模块失败: %v", err))
				result.Print(outputFormat)
				return nil
			}

			g.Log().Infof(ctx, "模块 '%s' 导入成功！", moduleName)
			result := utils.NewCommandResult(true, fmt.Sprintf("模块 '%s' 导入成功", moduleName))
			result.Print(outputFormat)
			return nil
		},
	}

	// ModuleList 列出模块命令
	ModuleList = &gcmd.Command{
		Name:  "module:list",
		Brief: "列出所有已安装模块",
		Description: `
列出所有已安装的模块及其信息
用法: go run hack/generator/main.go module:list [-output text|json]
`,
		Arguments: []gcmd.Argument{
			{
				Name:   "output",
				Short:  "o",
				Brief:  "输出格式(text/json)",
				IsArg:  false,
				Orphan: false,
			},
		},
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			outputFormat := utils.ParseOutputFormat(parser.GetOpt("output").String())

			modules, err := scanner.ListModules()
			if err != nil {
				result := utils.NewCommandResult(false, fmt.Sprintf("列出模块失败: %v", err))
				result.Print(outputFormat)
				return nil
			}

			if len(modules) == 0 {
				result := utils.NewCommandResult(true, "暂无已安装模块")
				result.Print(outputFormat)
				return nil
			}

			result := utils.NewCommandResult(true, fmt.Sprintf("已安装模块 (%d个)", len(modules)))
			for _, module := range modules {
				result.AddFile(fmt.Sprintf("%-10s v%-7s %-12s %-7s",
					truncateString(module.Name, 10),
					truncateString(module.Version, 7),
					truncateString(module.Author, 12),
					truncateString(module.License, 7)))
			}
			result.Print(outputFormat)
			return nil
		},
	}

	// ModuleValidate 验证模块命令
	ModuleValidate = &gcmd.Command{
		Name:  "module:validate",
		Brief: "验证模块完整性",
		Description: `
验证模块的完整性和配置正确性
用法: go run hack/generator/main.go module:validate -name 模块名称 [-output text|json]
`,
		Arguments: []gcmd.Argument{
			{
				Name:   "name",
				Short:  "n",
				Brief:  "模块名称",
				IsArg:  false,
				Orphan: false,
			},
			{
				Name:   "output",
				Short:  "o",
				Brief:  "输出格式(text/json)",
				IsArg:  false,
				Orphan: false,
			},
		},
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			outputFormat := utils.ParseOutputFormat(parser.GetOpt("output").String())
			moduleName := parser.GetOpt("name").String()
			if moduleName == "" {
				result := utils.NewCommandResult(false, "缺少必填参数: -name")
				result.Print(outputFormat)
				return nil
			}

			g.Log().Infof(ctx, "验证模块: %s", moduleName)

			warnings, errors, err := scanner.ValidateModule(moduleName)
			if err != nil {
				result := utils.NewCommandResult(false, fmt.Sprintf("验证模块失败: %v", err))
				result.Print(outputFormat)
				return nil
			}

			result := utils.NewCommandResult(true, fmt.Sprintf("验证模块 '%s'", moduleName))
			for _, w := range warnings {
				result.AddWarning(w)
			}
			for _, e := range errors {
				result.AddError(e)
			}
			if len(errors) == 0 && len(warnings) == 0 {
				result.Message = fmt.Sprintf("模块 '%s' 验证通过，没有发现问题", moduleName)
			} else {
				result.Success = len(errors) == 0
				if len(errors) > 0 {
					result.Message = fmt.Sprintf("模块 '%s' 验证发现 %d 个错误，%d 个警告", moduleName, len(errors), len(warnings))
				} else {
					result.Message = fmt.Sprintf("模块 '%s' 验证通过，但有 %d 个警告", moduleName, len(warnings))
				}
			}
			result.Print(outputFormat)
			return nil
		},
	}
)

// truncateString 截断字符串
func truncateString(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen-3] + "..."
}
