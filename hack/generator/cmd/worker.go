// Package cmd
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package cmd

import (
	"context"
	"fmt"

	"devinggo/hack/generator/internal/generator"
	"devinggo/hack/generator/internal/utils"

	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	// WorkerCreate 创建Worker任务命令
	WorkerCreate = &gcmd.Command{
		Name:  "worker:create",
		Brief: "创建 Task 或 Cron 任务",
		Description: `
创建 Worker 任务（异步任务或定时任务）

快速创建 Task 或 Cron 任务，自动生成规范代码。支持三种模式：
- task: 仅创建异步任务
- cron: 仅创建定时任务
- both: 同时创建任务和定时任务（数据结构共享）

使用示例:
  # 创建发送邮件任务（默认both类型）
  go run hack/generator/main.go worker:create -name send_email -desc "发送邮件"

  # 仅创建异步任务
  go run hack/generator/main.go worker:create -name process_order -type task -desc "处理订单"

  # 仅创建定时任务
  go run hack/generator/main.go worker:create -name clean_logs -type cron -desc "清理日志"

  # 在指定模块中创建
  go run hack/generator/main.go worker:create -name notify_user -module custom -desc "用户通知"

命令选项:
  -name   任务名称（必填），建议使用下划线命名，如: send_email
  -module 模块名称（可选，默认: system）
  -type   任务类型（可选，默认: both）
          task: 仅创建异步任务
          cron: 仅创建定时任务
          both: 同时创建任务和定时任务
  -desc   任务描述（可选），用于生成注释
  -output 输出格式（可选，默认: text）text|json

生成的文件:
  • modules/{module}/worker/cron/{name}_cron.go        (定时任务文件)
  • modules/{module}/worker/server/{name}_worker.go    (异步任务文件)
  • modules/{module}/worker/consts/const.go            (常量定义文件)

注意事项:
  1. 任务名称建议使用小写下划线格式
  2. 目标模块必须已存在
  3. 不能创建重复的任务名称
  4. 生成后需要重启 worker 服务以加载新任务
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

			// 获取参数
			name := parser.GetOpt("name").String()
			moduleName := parser.GetOpt("module", "system").String()
			workerType := parser.GetOpt("type", "both").String()
			description := parser.GetOpt("desc").String()

			if name == "" {
				result := utils.NewCommandResult(false, "缺少必填参数: -name")
				result.Print(outputFormat)
				return nil
			}

			// 转换类型
			var wType generator.WorkerType
			switch workerType {
			case "task":
				wType = generator.WorkerTypeTask
			case "cron":
				wType = generator.WorkerTypeCron
			case "both":
				wType = generator.WorkerTypeBoth
			default:
				result := utils.NewCommandResult(false, fmt.Sprintf("类型必须是 task、cron 或 both，当前值: %s", workerType))
				result.Print(outputFormat)
				return nil
			}

			// 创建生成器
			gen := generator.NewWorkerGenerator(ctx, moduleName, name, description, wType)

			// 执行生成
			if err := gen.Generate(); err != nil {
				result := utils.NewCommandResult(false, fmt.Sprintf("创建Worker任务失败: %v", err))
				result.Print(outputFormat)
				return nil
			}

			result := utils.NewCommandResult(true, fmt.Sprintf("Worker任务 '%s' 创建成功", name))
			result.Print(outputFormat)
			return nil
		},
	}
)
