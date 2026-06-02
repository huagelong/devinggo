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

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/text/gstr"
)

// WorkerType 任务类型
type WorkerType string

const (
	WorkerTypeTask WorkerType = "task" // 异步任务
	WorkerTypeCron WorkerType = "cron" // 定时任务
	WorkerTypeBoth WorkerType = "both" // 两者都有
)

// WorkerGenerator Worker任务生成器
type WorkerGenerator struct {
	ctx         context.Context
	moduleName  string
	name        string
	description string
	workerType  WorkerType
	templateDir string
}

// NewWorkerGenerator 创建Worker生成器实例
func NewWorkerGenerator(ctx context.Context, moduleName, name, description string, workerType WorkerType) *WorkerGenerator {
	if description == "" {
		description = name
	}

	return &WorkerGenerator{
		ctx:         ctx,
		moduleName:  moduleName,
		name:        name,
		description: description,
		workerType:  workerType,
		templateDir: "./hack/generator/templates/worker",
	}
}

// Generate 生成Worker任务
func (w *WorkerGenerator) Generate() error {
	// 验证参数
	if err := w.validate(); err != nil {
		return err
	}

	// 打印创建信息
	w.printInfo()

	// 创建必要的目录
	if err := w.createDirectories(); err != nil {
		return err
	}

	// 更新常量文件
	if err := w.updateConstFile(); err != nil {
		return err
	}

	// 根据类型创建文件
	if w.workerType == WorkerTypeCron || w.workerType == WorkerTypeBoth {
		if err := w.createCronFile(); err != nil {
			return err
		}
	}

	if w.workerType == WorkerTypeTask || w.workerType == WorkerTypeBoth {
		if err := w.createTaskFile(); err != nil {
			return err
		}
	}

	// 输出成功信息
	w.printSuccess()

	return nil
}

// validate 验证参数
func (w *WorkerGenerator) validate() error {
	// 验证任务名称
	if w.name == "" {
		return gerror.New("任务名称不能为空")
	}

	// 验证类型
	if w.workerType != WorkerTypeTask && w.workerType != WorkerTypeCron && w.workerType != WorkerTypeBoth {
		return gerror.New("类型必须是 task、cron 或 both")
	}

	// 验证模块是否存在
	modulePath := fmt.Sprintf("./modules/%s", w.moduleName)
	if !gfile.Exists(modulePath) {
		return gerror.Newf("模块 '%s' 不存在，请先创建模块或使用已存在的模块", w.moduleName)
	}

	// 验证模板目录
	if !gfile.Exists(w.templateDir) {
		return gerror.Newf("模板目录 '%s' 不存在", w.templateDir)
	}

	return nil
}

// printInfo 打印创建信息
func (w *WorkerGenerator) printInfo() {
	fmt.Println("\n━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Printf("📦 任务名称: %s\n", w.name)
	fmt.Printf("📂 所属模块: %s\n", w.moduleName)
	fmt.Printf("🔖 任务类型: %s\n", w.workerType)
	fmt.Printf("📝 任务描述: %s\n", w.description)
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
}

// printSuccess 打印成功信息
func (w *WorkerGenerator) printSuccess() {
	fmt.Println("\n✅ 创建成功！")
	fmt.Println("\n📁 生成的文件:")
	if w.workerType == WorkerTypeCron || w.workerType == WorkerTypeBoth {
		fmt.Printf("   • modules/%s/worker/cron/%s_cron.go\n", w.moduleName, w.name)
	}
	if w.workerType == WorkerTypeTask || w.workerType == WorkerTypeBoth {
		fmt.Printf("   • modules/%s/worker/server/%s_worker.go\n", w.moduleName, w.name)
	}
	fmt.Printf("   • modules/%s/worker/consts/const.go (已更新)\n", w.moduleName)

	fmt.Println("\n💡 下一步:")
	fmt.Println("   1. 编辑生成的文件，添加业务逻辑")
	fmt.Println("   2. 如果 worker 服务正在运行，需要重启以加载新任务")
	if w.workerType == WorkerTypeCron || w.workerType == WorkerTypeBoth {
		fmt.Println("   3. 在后台管理系统中配置定时任务的执行时间")
	}
	fmt.Println()
}

// createDirectories 创建必要的目录
func (w *WorkerGenerator) createDirectories() error {
	dirs := []string{
		fmt.Sprintf("./modules/%s/worker", w.moduleName),
		fmt.Sprintf("./modules/%s/worker/consts", w.moduleName),
	}

	if w.workerType == WorkerTypeCron || w.workerType == WorkerTypeBoth {
		dirs = append(dirs, fmt.Sprintf("./modules/%s/worker/cron", w.moduleName))
	}

	if w.workerType == WorkerTypeTask || w.workerType == WorkerTypeBoth {
		dirs = append(dirs, fmt.Sprintf("./modules/%s/worker/server", w.moduleName))
	}

	for _, dir := range dirs {
		if !gfile.Exists(dir) {
			if err := gfile.Mkdir(dir); err != nil {
				return gerror.Wrapf(err, "创建目录 '%s' 失败", dir)
			}
			g.Log().Debugf(w.ctx, "创建目录: %s", dir)
		}
	}
	return nil
}

// updateConstFile 使用AST方式更新常量文件
func (w *WorkerGenerator) updateConstFile() error {
	constPath := fmt.Sprintf("./modules/%s/worker/consts/const.go", w.moduleName)
	constName := strings.ToUpper(gstr.CaseSnake(w.name))

	// 检查文件是否存在
	if !gfile.Exists(constPath) {
		// 创建新的常量文件
		return w.createNewConstFile(constPath, constName)
	}

	// 使用AST更新现有文件
	return w.updateExistingConstFile(constPath, constName)
}

// createNewConstFile 创建新的常量文件
func (w *WorkerGenerator) createNewConstFile(constPath, constName string) error {
	var constants []string

	if w.workerType == WorkerTypeCron || w.workerType == WorkerTypeBoth {
		constants = append(constants, fmt.Sprintf("\t%s_CRON = \"%s_cron\" // %s", constName, w.name, w.description))
	}
	if w.workerType == WorkerTypeTask || w.workerType == WorkerTypeBoth {
		constants = append(constants, fmt.Sprintf("\t%s_TASK = \"%s_task\" // %s", constName, w.name, w.description))
	}

	content := fmt.Sprintf(`// Package consts
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package consts

var (
%s
)
`, strings.Join(constants, "\n"))

	if err := gfile.PutContents(constPath, content); err != nil {
		return gerror.Wrapf(err, "创建常量文件失败")
	}
	g.Log().Debugf(w.ctx, "创建常量文件: %s", constPath)
	return nil
}

// updateExistingConstFile 使用AST更新现有常量文件
func (w *WorkerGenerator) updateExistingConstFile(constPath, constName string) error {
	// 读取原始文件内容
	content := gfile.GetContents(constPath)
	if content == "" {
		return gerror.New("常量文件为空")
	}

	// 准备要添加的常量行
	var newConstants []string

	if w.workerType == WorkerTypeCron || w.workerType == WorkerTypeBoth {
		cronConstName := fmt.Sprintf("%s_CRON", constName)
		if strings.Contains(content, cronConstName) {
			return gerror.Newf("常量 %s 已存在，请检查是否重复创建", cronConstName)
		}
		newConstants = append(newConstants, fmt.Sprintf("\t%s = \"%s_cron\" // %s", cronConstName, w.name, w.description))
	}

	if w.workerType == WorkerTypeTask || w.workerType == WorkerTypeBoth {
		taskConstName := fmt.Sprintf("%s_TASK", constName)
		if strings.Contains(content, taskConstName) {
			return gerror.Newf("常量 %s 已存在，请检查是否重复创建", taskConstName)
		}
		newConstants = append(newConstants, fmt.Sprintf("\t%s = \"%s_task\" // %s", taskConstName, w.name, w.description))
	}

	// 找到最后一个 ) 的位置
	lastParen := strings.LastIndex(content, ")")
	if lastParen == -1 {
		return gerror.New("常量文件格式错误：未找到右括号")
	}

	// 在右括号前插入新常量
	newContent := content[:lastParen]
	for _, constLine := range newConstants {
		newContent += constLine + "\n"
	}
	newContent += content[lastParen:]

	// 写回文件
	if err := gfile.PutContents(constPath, newContent); err != nil {
		return gerror.Wrapf(err, "写入文件失败")
	}

	g.Log().Debugf(w.ctx, "更新常量文件: %s", constPath)
	return nil
}

// createCronFile 创建Cron文件
func (w *WorkerGenerator) createCronFile() error {
	cronPath := fmt.Sprintf("./modules/%s/worker/cron/%s_cron.go", w.moduleName, w.name)

	// 检查文件是否已存在
	if gfile.Exists(cronPath) {
		return gerror.Newf("Cron 文件 '%s' 已存在，请检查是否重复创建", cronPath)
	}

	// 准备模板变量
	data := map[string]interface{}{
		"ModuleName":  w.moduleName,
		"Name":        w.name,
		"Description": w.description,
		"StructName":  gstr.CaseCamel(w.name) + "Data",
		"ConstName":   strings.ToUpper(gstr.CaseSnake(w.name)),
		"HandlerName": "handle" + gstr.CaseCamel(w.name),
	}

	// 渲染模板
	content, err := RenderTemplate(filepath.Join(w.templateDir, "cron.go.tpl"), data)
	if err != nil {
		return gerror.Wrap(err, "渲染Cron模板失败")
	}

	// 写入文件
	if err := gfile.PutContents(cronPath, content); err != nil {
		return gerror.Wrapf(err, "创建 Cron 文件失败")
	}

	g.Log().Debugf(w.ctx, "创建 Cron 文件: %s", cronPath)
	return nil
}

// createTaskFile 创建Task文件
func (w *WorkerGenerator) createTaskFile() error {
	taskPath := fmt.Sprintf("./modules/%s/worker/server/%s_worker.go", w.moduleName, w.name)

	// 检查文件是否已存在
	if gfile.Exists(taskPath) {
		return gerror.Newf("Task 文件 '%s' 已存在，请检查是否重复创建", taskPath)
	}

	// 准备导入和数据类型别名
	hasCron := w.workerType == WorkerTypeBoth
	var importCron, dataTypeAlias string

	structName := gstr.CaseCamel(w.name) + "Data"

	if hasCron {
		importCron = fmt.Sprintf("\t\"devinggo/modules/%s/worker/cron\"\n", w.moduleName)
		dataTypeAlias = fmt.Sprintf("// 复用 Cron 的数据结构\ntype %s = cron.%s\n", structName, structName)
	} else {
		dataTypeAlias = fmt.Sprintf(`// %s %s的数据结构
type %s struct {
	// TODO: 在这里定义你的参数字段
	// 例如：
	// Name  string `+"`json:\"name\" v:\"required#名称不能为空\"`"+`
	// Email string `+"`json:\"email\" v:\"required|email#邮箱不能为空|邮箱格式错误\"`"+`
}
`, structName, w.description, structName)
	}

	// 准备模板变量
	data := map[string]interface{}{
		"ModuleName":    w.moduleName,
		"Name":          w.name,
		"Description":   w.description,
		"StructName":    structName,
		"ConstName":     strings.ToUpper(gstr.CaseSnake(w.name)),
		"FuncName":      gstr.CaseCamel(w.name),
		"ImportCron":    importCron,
		"DataTypeAlias": dataTypeAlias,
	}

	// 渲染模板
	content, err := RenderTemplate(filepath.Join(w.templateDir, "task.go.tpl"), data)
	if err != nil {
		return gerror.Wrap(err, "渲染Task模板失败")
	}

	// 写入文件
	if err := gfile.PutContents(taskPath, content); err != nil {
		return gerror.Wrapf(err, "创建 Task 文件失败")
	}

	g.Log().Debugf(w.ctx, "创建 Task 文件: %s", taskPath)
	return nil
}
