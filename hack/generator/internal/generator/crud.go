// Package generator
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package generator

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"devinggo/hack/generator/internal/utils"

	"github.com/gogf/gf/v2/text/gstr"
)

// FrontendField 前端字段信息
type FrontendField struct {
	Name         string // 字段名（JSON名）
	Type         string // TypeScript类型
	TSParamType  string // 参数类型（用于可选字段）
	DefaultValue string // 默认值
	IsSearchable bool   // 是否可搜索
	IsList       bool   // 是否在列表中显示
	IsEditable   bool   // 是否可编辑
	IsRequired   bool   // 是否必填
	Optional     bool   // 是否可选（对应TypeScript ?）
	Comment      string // 字段注释
	// 表格列相关
	Align    string // 对齐方式
	ColKey   string // 列key
	MinWidth string // 最小宽度
	Width    string // 宽度
	Title    string // 列标题（i18n键）
	// 搜索表单相关
	Component   string // 组件类型（Input, Select, DateRangePicker）
	LabelKey    string // 标签i18n键
	Placeholder string // 占位符
}

// CRUDGenerator CRUD代码生成器
type CRUDGenerator struct {
	ModuleName  string  // 模块名（例如：system）
	TableName   string  // 表名（例如：system_api）
	EntityName  string  // 实体名（例如：SystemApi）
	VarName     string  // 变量名（例如：api）
	PackageName string  // 包名（例如：system）
	ChineseName string  // 中文名（例如：接口）
	Fields      []Field // 字段列表
	WorkDir     string  // 工作目录
	Force       bool    // 是否覆盖已存在的文件
	DryRun      bool    // 仅预览，不实际写入

	// 前端生成配置
	GenerateFrontend bool   // 是否生成前端代码
	FrontendDir      string // 前端代码目录

	GeneratedFiles []string
	SkippedFiles   []string
}

// Field 字段信息
type Field struct {
	Name         string // 字段名（Go格式，例如：GroupId）
	ColumnName   string // 列名（数据库格式，例如：group_id）
	Type         string // 字段类型（例如：int64, string）
	JSONName     string // JSON标签名（例如：group_id）
	Comment      string // 字段注释
	IsSearchable bool   // 是否可搜索
	IsRequired   bool   // 是否必填
}

// NewCRUDGenerator 创建CRUD生成器实例
func NewCRUDGenerator(moduleName, tableName, chineseName string) (*CRUDGenerator, error) {
	workDir, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("获取工作目录失败：%v", err)
	}

	normalizedPath := filepath.ToSlash(workDir)
	if strings.HasSuffix(normalizedPath, "hack/generator") {
		workDir = filepath.Join(workDir, "..", "..")
	}

	entityName := gstr.CaseCamel(tableName)

	parts := strings.Split(tableName, "_")
	var resourceName string
	if len(parts) > 1 && parts[0] == moduleName {
		resourceName = strings.Join(parts[1:], "_")
	} else {
		resourceName = tableName
	}
	varName := gstr.CaseCamelLower(resourceName)

	return &CRUDGenerator{
		ModuleName:     moduleName,
		TableName:      tableName,
		EntityName:     entityName,
		VarName:        varName,
		PackageName:    moduleName,
		ChineseName:    chineseName,
		WorkDir:        workDir,
		GeneratedFiles: make([]string, 0),
		SkippedFiles:   make([]string, 0),
	}, nil
}

// SetForce 设置是否覆盖已存在文件
func (g *CRUDGenerator) SetForce(force bool) { g.Force = force }

// SetDryRun 设置是否仅预览
func (g *CRUDGenerator) SetDryRun(dryRun bool) { g.DryRun = dryRun }

// SetGenerateFrontend 设置是否生成前端代码
func (g *CRUDGenerator) SetGenerateFrontend(generate bool) { g.GenerateFrontend = generate }

// getFrontendDir 获取前端代码目录
func (g *CRUDGenerator) getFrontendDir() string {
	if g.FrontendDir != "" {
		return g.FrontendDir
	}
	return filepath.Join(g.WorkDir, "admin-ui", "apps", "backend", "src")
}

// mapGoTypeToTSType 将Go类型映射为TypeScript类型
func mapGoTypeToTSType(goType string) string {
	goType = strings.TrimPrefix(goType, "*")
	switch {
	case strings.Contains(goType, "string"):
		return "string"
	case strings.Contains(goType, "int") || strings.Contains(goType, "uint") || strings.Contains(goType, "float"):
		return "number"
	case strings.Contains(goType, "bool"):
		return "boolean"
	case strings.Contains(goType, "Time"):
		return "string"
	default:
		return "any"
	}
}

// getDefaultValue 获取字段默认值
func getDefaultValue(tsType string, isSearch bool) string {
	if isSearch {
		// 搜索表单字段默认使用 undefined
		switch tsType {
		case "string":
			return "''"
		case "number":
			return "undefined"
		case "boolean":
			return "undefined"
		default:
			return "undefined"
		}
	}
	// 表单字段默认值
	switch tsType {
	case "string":
		return "''"
	case "number":
		return "0"
	case "boolean":
		return "false"
	default:
		return "undefined"
	}
}

// getSearchComponent 根据字段类型获取搜索表单组件
func getSearchComponent(tsType string, fieldName string) string {
	if fieldName == "created_at" || fieldName == "updated_at" {
		return "DateRangePicker"
	}
	if fieldName == "status" {
		return "Select"
	}
	switch tsType {
	case "string":
		return "Input"
	case "number":
		return "InputNumber"
	default:
		return "Input"
	}
}

// getColumnWidth 获取列宽度
func getColumnWidth(fieldName string) string {
	switch fieldName {
	case "id":
		return "80"
	case "status":
		return "120"
	case "sort":
		return "140"
	case "created_at", "updated_at":
		return "180"
	case "email":
		return "150"
	case "phone":
		return "120"
	default:
		return ""
	}
}

// buildFrontendFields 构建前端字段列表
func (g *CRUDGenerator) buildFrontendFields() []FrontendField {
	fields := make([]FrontendField, 0)
	
	// 排除的系统字段
	excludeSystemFields := map[string]bool{
		"created_by": true,
		"updated_by": true,
		"deleted_at": true,
	}
	
	for _, f := range g.Fields {
		if excludeSystemFields[f.JSONName] {
			continue
		}
		
		tsType := mapGoTypeToTSType(f.Type)
		isSearchable := f.IsSearchable && f.Name != "Id"
		isEditable := f.Name != "Id"
		isList := true
		
		// 搜索表单默认值
		searchDefault := getDefaultValue(tsType, true)
		if f.JSONName == "created_at" || f.JSONName == "updated_at" {
			searchDefault = "[]"
		}
		
		// 跳过Id字段的编辑
		if f.Name == "Id" {
			continue
		}
		
		field := FrontendField{
			Name:         f.JSONName,
			Type:         tsType,
			TSParamType:  tsType,
			DefaultValue: searchDefault,
			IsSearchable: isSearchable,
			IsList:       isList,
			IsEditable:   isEditable,
			IsRequired:   f.IsRequired,
			Optional:     !f.IsRequired,
			Comment:      f.Comment,
			Align:        "center",
			ColKey:       f.JSONName,
			Width:        getColumnWidth(f.JSONName),
			Title:        fmt.Sprintf("$t('%s.%s.%s')", g.ModuleName, g.VarName, f.JSONName),
			Component:    getSearchComponent(tsType, f.JSONName),
			LabelKey:     fmt.Sprintf("%s.%s.%s", g.ModuleName, g.VarName, f.JSONName),
			Placeholder:  fmt.Sprintf("$t('ui.placeholder.input')"),
		}
		
		if field.Component == "Select" {
			field.Placeholder = fmt.Sprintf("$t('ui.placeholder.select')")
		}
		
		if field.Width == "" {
			field.MinWidth = "140"
		}
		
		fields = append(fields, field)
	}
	
	return fields
}

// buildFrontendTemplateData 构建前端模板数据
func (g *CRUDGenerator) buildFrontendTemplateData() map[string]interface{} {
	frontendFields := g.buildFrontendFields()
	
	// 分类字段
	searchFields := make([]FrontendField, 0)
	listFields := make([]FrontendField, 0)
	editableFields := make([]FrontendField, 0)
	tableColumns := make([]FrontendField, 0)
	searchFormItems := make([]FrontendField, 0)
	
	for _, f := range frontendFields {
		if f.IsSearchable {
			searchFields = append(searchFields, f)
			searchFormItems = append(searchFormItems, f)
		}
		if f.IsList {
			listFields = append(listFields, f)
			tableColumns = append(tableColumns, f)
		}
		if f.IsEditable {
			editField := f
			editField.DefaultValue = getDefaultValue(f.Type, false)
			editableFields = append(editableFields, editField)
		}
	}
	
	componentName := g.ModuleName + g.EntityName
	if g.ModuleName == "system" {
		componentName = "System" + g.EntityName
	}
	
	// 生成基于时间戳的菜单ID，避免冲突
	menuId := int(time.Now().Unix())

	return map[string]interface{}{
		"ModuleName":      g.ModuleName,
		"TableName":       g.TableName,
		"EntityName":      g.EntityName,
		"VarName":         g.VarName,
		"PackageName":     g.PackageName,
		"ChineseName":     g.ChineseName,
		"ApiPrefix":       fmt.Sprintf("/%s/%s", g.ModuleName, g.VarName),
		"ComponentName":   componentName,
		"SearchFields":    searchFields,
		"ListFields":      listFields,
		"EditableFields":  editableFields,
		"TableColumns":    tableColumns,
		"SearchFormItems": searchFormItems,
		"FrontendFields":  frontendFields,
		"Date":            time.Now().Format("2006-01-02 15:04:05"),
		"ParentId":        0,
		"MenuId":          menuId,
		"Level":           ",0,",
		"ParentMenuId":    0,
	}
}

func (g *CRUDGenerator) getTemplateDir() string {
	root, err := utils.GetProjectRoot()
	if err != nil {
		return ""
	}
	return filepath.Join(root, "hack", "generator", "templates", "crud")
}

// ParseEntityFields 从Entity文件解析字段信息
func (g *CRUDGenerator) ParseEntityFields() error {
	entityPath := filepath.Join(g.WorkDir, "internal", "model", "entity", g.TableName+".go")
	content, err := os.ReadFile(entityPath)
	if err != nil {
		return fmt.Errorf("读取entity文件失败：%v", err)
	}

	lines := strings.Split(string(content), "\n")
	inStruct := false
	g.Fields = make([]Field, 0)

	for _, line := range lines {
		line = strings.TrimSpace(line)

		if strings.Contains(line, "type "+g.EntityName+" struct") {
			inStruct = true
			continue
		}

		if inStruct {
			if strings.HasPrefix(line, "}") {
				break
			}

			if line == "" || strings.HasPrefix(line, "//") {
				continue
			}

			field := g.parseFieldLine(line)
			if field != nil {
				if g.shouldIncludeField(field.Name) {
					g.Fields = append(g.Fields, *field)
				}
			}
		}
	}

	if len(g.Fields) == 0 {
		return fmt.Errorf("未能解析到任何字段")
	}

	return nil
}

func (g *CRUDGenerator) parseFieldLine(line string) *Field {
	line = strings.TrimSpace(line)

	backquoteIndex := strings.Index(line, "`")
	if backquoteIndex == -1 {
		return nil
	}

	fieldDef := strings.TrimSpace(line[:backquoteIndex])
	tags := line[backquoteIndex:]

	parts := strings.Fields(fieldDef)
	if len(parts) < 2 {
		return nil
	}

	name := parts[0]
	fieldType := parts[1]

	jsonName := g.extractTag(tags, "json")
	columnName := g.extractTag(tags, "orm")
	comment := g.extractTag(tags, "description")

	isSearchable := isSearchableType(fieldType)
	isRequired := !strings.Contains(fieldType, "*") && name != "Id"

	return &Field{
		Name:         name,
		ColumnName:   columnName,
		Type:         fieldType,
		JSONName:     jsonName,
		Comment:      comment,
		IsSearchable: isSearchable,
		IsRequired:   isRequired,
	}
}

func (g *CRUDGenerator) extractTag(tags, key string) string {
	keyPattern := key + `:"`
	start := strings.Index(tags, keyPattern)
	if start == -1 {
		return ""
	}
	start += len(keyPattern)

	end := strings.Index(tags[start:], `"`)
	if end == -1 {
		return ""
	}

	return tags[start : start+end]
}

func (g *CRUDGenerator) shouldIncludeField(fieldName string) bool {
	excludeFields := []string{"CreatedBy", "UpdatedBy", "CreatedAt", "UpdatedAt", "DeletedAt"}
	for _, exclude := range excludeFields {
		if fieldName == exclude {
			return false
		}
	}
	return true
}

func isSearchableType(fieldType string) bool {
	searchableTypes := []string{"string", "int", "int64", "int32"}
	for _, t := range searchableTypes {
		if strings.Contains(fieldType, t) {
			return true
		}
	}
	return false
}

// Generate 生成所有CRUD文件
func (g *CRUDGenerator) Generate() error {
	if err := g.ParseEntityFields(); err != nil {
		return err
	}

	generators := []struct {
		name string
		fn   func() error
	}{
		{"API定义", g.GenerateAPI},
		{"请求模型", g.GenerateReq},
		{"响应模型", g.GenerateRes},
		{"控制器", g.GenerateController},
		{"逻辑层", g.GenerateLogic},
	}

	for _, gen := range generators {
		fmt.Printf("正在生成%s...\n", gen.name)
		if err := gen.fn(); err != nil {
			return fmt.Errorf("生成%s失败：%v", gen.name, err)
		}
	}

	// 生成前端代码
	if g.GenerateFrontend {
		fmt.Printf("\n正在生成前端代码...\n")
		if err := g.GenerateFrontendCode(); err != nil {
			return fmt.Errorf("生成前端代码失败：%v", err)
		}
	}

	// 自动注册 service 接口
	fmt.Printf("\n正在注册 service 接口...\n")
	if err := g.RegisterServiceInterface(); err != nil {
		fmt.Printf("  ⚠️ 注册 service 接口失败（可手动添加）：%v\n", err)
	}

	// 自动注册路由
	fmt.Printf("正在注册路由...\n")
	if err := g.RegisterRouter(); err != nil {
		fmt.Printf("  ⚠️ 注册路由失败（可手动添加）：%v\n", err)
	}

	fmt.Printf("\n✓ CRUD代码生成完成！\n")
	return nil
}

// GenerateFrontendCode 生成前端代码
func (g *CRUDGenerator) GenerateFrontendCode() error {
	generators := []struct {
		name string
		fn   func() error
	}{
		{"前端API", g.GenerateFrontendAPI},
		{"前端模型", g.GenerateFrontendModel},
		{"前端列配置", g.GenerateFrontendSchemas},
		{"前端CRUD逻辑", g.GenerateFrontendUseCrud},
		{"前端页面", g.GenerateFrontendIndexVue},
		{"前端菜单SQL", g.GenerateFrontendMenuSQL},
	}

	for _, gen := range generators {
		fmt.Printf("正在生成%s...\n", gen.name)
		if err := gen.fn(); err != nil {
			return fmt.Errorf("生成%s失败：%v", gen.name, err)
		}
	}

	return nil
}

// GenerateFrontendAPI 生成前端API文件
func (g *CRUDGenerator) GenerateFrontendAPI() error {
	data := g.buildFrontendTemplateData()
	frontendDir := g.getFrontendDir()
	outputPath := filepath.Join(frontendDir, "api", g.ModuleName, g.VarName+".ts")
	templatePath := filepath.Join(g.getTemplateDir(), "frontend", "api.ts.tpl")
	return g.renderAndSaveTemplateFrontend(templatePath, outputPath, data)
}

// GenerateFrontendModel 生成前端模型文件
func (g *CRUDGenerator) GenerateFrontendModel() error {
	data := g.buildFrontendTemplateData()
	frontendDir := g.getFrontendDir()
	outputPath := filepath.Join(frontendDir, "views", g.ModuleName, g.VarName, "model.ts")
	templatePath := filepath.Join(g.getTemplateDir(), "frontend", "model.ts.tpl")
	return g.renderAndSaveTemplateFrontend(templatePath, outputPath, data)
}

// GenerateFrontendSchemas 生成前端列配置文件
func (g *CRUDGenerator) GenerateFrontendSchemas() error {
	data := g.buildFrontendTemplateData()
	frontendDir := g.getFrontendDir()
	outputPath := filepath.Join(frontendDir, "views", g.ModuleName, g.VarName, "schemas.ts")
	templatePath := filepath.Join(g.getTemplateDir(), "frontend", "schemas.ts.tpl")
	return g.renderAndSaveTemplateFrontend(templatePath, outputPath, data)
}

// GenerateFrontendUseCrud 生成前端CRUD逻辑文件
func (g *CRUDGenerator) GenerateFrontendUseCrud() error {
	data := g.buildFrontendTemplateData()
	frontendDir := g.getFrontendDir()
	outputPath := filepath.Join(frontendDir, "views", g.ModuleName, g.VarName, "use-"+g.VarName+"-crud.ts")
	templatePath := filepath.Join(g.getTemplateDir(), "frontend", "use-crud.ts.tpl")
	return g.renderAndSaveTemplateFrontend(templatePath, outputPath, data)
}

// GenerateFrontendIndexVue 生成前端页面文件
func (g *CRUDGenerator) GenerateFrontendIndexVue() error {
	data := g.buildFrontendTemplateData()
	frontendDir := g.getFrontendDir()
	outputPath := filepath.Join(frontendDir, "views", g.ModuleName, g.VarName, "index.vue")
	templatePath := filepath.Join(g.getTemplateDir(), "frontend", "index.vue.tpl")
	return g.renderAndSaveTemplateFrontend(templatePath, outputPath, data)
}

// GenerateFrontendMenuSQL 生成前端菜单SQL文件
func (g *CRUDGenerator) GenerateFrontendMenuSQL() error {
	data := g.buildFrontendTemplateData()
	outputPath := filepath.Join(g.WorkDir, "resource", "migrations", fmt.Sprintf("menu_%s_%s.sql", g.ModuleName, g.VarName))
	templatePath := filepath.Join(g.getTemplateDir(), "frontend", "menu.sql.tpl")
	return g.renderAndSaveTemplateFrontend(templatePath, outputPath, data)
}

// renderAndSaveTemplateFrontend 渲染并保存前端模板
func (g *CRUDGenerator) renderAndSaveTemplateFrontend(templatePath string, outputPath string, data interface{}) error {
	if utils.PathExists(outputPath) && !g.Force {
		fmt.Printf("  ⚠️  跳过已存在的前端文件: %s (使用 --force 覆盖)\n", outputPath)
		g.SkippedFiles = append(g.SkippedFiles, outputPath)
		return nil
	}

	if g.DryRun {
		fmt.Printf("  [dry-run] 将生成前端文件: %s\n", outputPath)
		g.GeneratedFiles = append(g.GeneratedFiles, outputPath)
		return nil
	}

	result, err := RenderFrontendTemplate(templatePath, data)
	if err != nil {
		return fmt.Errorf("渲染前端模板失败: %w", err)
	}

	dir := filepath.Dir(outputPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("创建前端目录失败：%v", err)
	}

	if err := os.WriteFile(outputPath, []byte(result), 0644); err != nil {
		return fmt.Errorf("写入前端文件失败：%v", err)
	}

	fmt.Printf("  ✓ 已生成前端文件：%s\n", outputPath)
	g.GeneratedFiles = append(g.GeneratedFiles, outputPath)
	return nil
}

// GenerateAPI 生成API定义文件
func (g *CRUDGenerator) GenerateAPI() error {
	data := map[string]string{
		"ModuleName":  g.ModuleName,
		"EntityName":  g.EntityName,
		"VarName":     g.VarName,
		"PackageName": g.PackageName,
		"ChineseName": g.ChineseName,
	}

	outputPath := filepath.Join(g.WorkDir, "modules", g.ModuleName, "api", g.PackageName, g.VarName+".go")
	templatePath := filepath.Join(g.getTemplateDir(), "api.go.tpl")
	return g.renderAndSaveTemplate(templatePath, outputPath, data)
}

// GenerateReq 生成请求模型文件
func (g *CRUDGenerator) GenerateReq() error {
	var searchFields strings.Builder
	for _, field := range g.Fields {
		if field.IsSearchable && field.Name != "Id" {
			searchFields.WriteString(fmt.Sprintf("\t%s %s `json:\"%s\"`\n",
				field.Name, field.Type, field.JSONName))
		}
	}

	var saveFields strings.Builder
	for _, field := range g.Fields {
		if field.Name == "Id" {
			continue
		}
		required := ""
		if field.IsRequired {
			required = ` v:"required"`
		}
		saveFields.WriteString(fmt.Sprintf("\t%s %s `json:\"%s\"%s description:\"%s\"`\n",
			field.Name, field.Type, field.JSONName, required, field.Comment))
	}

	var updateFields strings.Builder
	for _, field := range g.Fields {
		required := ""
		if field.Name == "Id" {
			required = ` v:"required"`
		} else if field.IsRequired {
			required = ` v:"required"`
		}
		updateFields.WriteString(fmt.Sprintf("\t%s %s `json:\"%s\"%s description:\"%s\"`\n",
			field.Name, field.Type, field.JSONName, required, field.Comment))
	}

	data := map[string]string{
		"EntityName":   g.EntityName,
		"SearchFields": searchFields.String(),
		"SaveFields":   saveFields.String(),
		"UpdateFields": updateFields.String(),
	}

	outputPath := filepath.Join(g.WorkDir, "modules", g.ModuleName, "model", "req", g.TableName+".go")
	templatePath := filepath.Join(g.getTemplateDir(), "req.go.tpl")
	return g.renderAndSaveTemplate(templatePath, outputPath, data)
}

// GenerateRes 生成响应模型文件
func (g *CRUDGenerator) GenerateRes() error {
	var fields strings.Builder
	fields.WriteString(fmt.Sprintf("\tId %s `json:\"%s\" description:\"%s\"`\n",
		"int64", "id", "主键"))

	for _, field := range g.Fields {
		if field.Name == "Id" {
			continue
		}
		fields.WriteString(fmt.Sprintf("\t%s %s `json:\"%s\" description:\"%s\"`\n",
			field.Name, field.Type, field.JSONName, field.Comment))
	}

	timestampFields := []string{
		"CreatedBy   int64       `json:\"created_by\" description:\"创建者\"`",
		"UpdatedBy   int64       `json:\"updated_by\" description:\"更新者\"`",
		"CreatedAt   *gtime.Time `json:\"created_at\" description:\"创建时间\"`",
		"UpdatedAt   *gtime.Time `json:\"updated_at\" description:\"更新时间\"`",
	}
	for _, ts := range timestampFields {
		fields.WriteString("\t" + ts + "\n")
	}

	data := map[string]string{
		"EntityName": g.EntityName,
		"Fields":     fields.String(),
	}

	outputPath := filepath.Join(g.WorkDir, "modules", g.ModuleName, "model", "res", g.TableName+".go")
	templatePath := filepath.Join(g.getTemplateDir(), "res.go.tpl")
	return g.renderAndSaveTemplate(templatePath, outputPath, data)
}

// GenerateController 生成控制器文件
func (g *CRUDGenerator) GenerateController() error {
	data := map[string]string{
		"ModuleName":  g.ModuleName,
		"EntityName":  g.EntityName,
		"VarName":     g.VarName,
		"PackageName": g.PackageName,
		"ChineseName": g.ChineseName,
	}

	outputPath := filepath.Join(g.WorkDir, "modules", g.ModuleName, "controller", g.PackageName, g.VarName+".go")
	templatePath := filepath.Join(g.getTemplateDir(), "controller.go.tpl")
	return g.renderAndSaveTemplate(templatePath, outputPath, data)
}

// GenerateLogic 生成逻辑层文件
func (g *CRUDGenerator) GenerateLogic() error {
	var searchConditions strings.Builder
	for _, field := range g.Fields {
		if field.IsSearchable && field.Name != "Id" {
			searchConditions.WriteString(fmt.Sprintf(`
	if !g.IsEmpty(in.%s) {
		m = m.Where("%s", in.%s)
	}
`, field.Name, field.ColumnName, field.Name))
		}
	}

	var saveDoFields strings.Builder
	for _, field := range g.Fields {
		if field.Name == "Id" {
			continue
		}
		saveDoFields.WriteString(fmt.Sprintf("\t\t%s: in.%s,\n", field.Name, field.Name))
	}

	var updateDoFields strings.Builder
	for _, field := range g.Fields {
		if field.Name == "Id" {
			continue
		}
		updateDoFields.WriteString(fmt.Sprintf("\t\t%s: in.%s,\n", field.Name, field.Name))
	}

	data := map[string]string{
		"PackageName":      g.PackageName,
		"ModuleName":       g.ModuleName,
		"EntityName":       g.EntityName,
		"SearchConditions": searchConditions.String(),
		"SaveDoFields":     saveDoFields.String(),
		"UpdateDoFields":   updateDoFields.String(),
	}

	outputPath := filepath.Join(g.WorkDir, "modules", g.ModuleName, "logic", g.PackageName, g.TableName+".go")
	templatePath := filepath.Join(g.getTemplateDir(), "logic.go.tpl")
	return g.renderAndSaveTemplate(templatePath, outputPath, data)
}

func (g *CRUDGenerator) renderAndSaveTemplate(templatePath string, outputPath string, data map[string]string) error {
	if utils.PathExists(outputPath) && !g.Force {
		fmt.Printf("  ⚠️  跳过已存在的文件: %s (使用 --force 覆盖)\n", outputPath)
		g.SkippedFiles = append(g.SkippedFiles, outputPath)
		return nil
	}

	if g.DryRun {
		fmt.Printf("  [dry-run] 将生成: %s\n", outputPath)
		g.GeneratedFiles = append(g.GeneratedFiles, outputPath)
		return nil
	}

	result, err := RenderTemplate(templatePath, data)
	if err != nil {
		return fmt.Errorf("渲染模板失败: %w", err)
	}

	dir := filepath.Dir(outputPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("创建目录失败：%v", err)
	}

	if err := os.WriteFile(outputPath, []byte(result), 0644); err != nil {
		return fmt.Errorf("写入文件失败：%v", err)
	}

	fmt.Printf("  ✓ 已生成：%s\n", outputPath)
	g.GeneratedFiles = append(g.GeneratedFiles, outputPath)
	return nil
}

// RegisterServiceInterface 自动在 service/system.go 中注册接口
func (g *CRUDGenerator) RegisterServiceInterface() error {
	serviceFile := filepath.Join(g.WorkDir, "modules", g.ModuleName, "service", "system.go")
	if !utils.PathExists(serviceFile) {
		return fmt.Errorf("service 文件不存在: %s", serviceFile)
	}

	content, err := os.ReadFile(serviceFile)
	if err != nil {
		return fmt.Errorf("读取 service 文件失败: %v", err)
	}

	contentStr := string(content)

	// 检查是否已注册
	if strings.Contains(contentStr, fmt.Sprintf("I%s interface", g.EntityName)) {
		fmt.Printf("  ℹ️ service 接口已存在，跳过\n")
		return nil
	}

	// 1. 插入接口定义（在 ISystemUser 之前）
	interfaceDef := fmt.Sprintf(`	// I%s defines the interface for %s operations.
	I%s interface {
		// Model returns the database Model for %s operations.
		Model(ctx context.Context) *gdb.Model
		// GetPageListForSearch retrieves a paginated list of %s with search criteria.
		GetPageListForSearch(ctx context.Context, req *model.PageListReq, in *req.%sSearch) (rs []*res.%s, total int, err error)
		// GetList retrieves a list of %s with search criteria.
		GetList(ctx context.Context, in *req.%sSearch) (out []*res.%s, err error)
		// Save creates a new %s.
		Save(ctx context.Context, in *req.%sSave) (id int64, err error)
		// GetById retrieves a %s by ID.
		GetById(ctx context.Context, id int64) (res *res.%s, err error)
		// Update modifies an existing %s.
		Update(ctx context.Context, in *req.%sUpdate) (err error)
		// Delete soft deletes %s by IDs.
		Delete(ctx context.Context, ids []int64) (err error)
		// RealDelete permanently removes %s by IDs.
		RealDelete(ctx context.Context, ids []int64) (err error)
		// Recovery restores soft-deleted %s.
		Recovery(ctx context.Context, ids []int64) (err error)
		// ChangeStatus changes the status of a %s.
		ChangeStatus(ctx context.Context, id int64, status int) (err error)
	}
`, g.EntityName, g.VarName, g.EntityName, g.VarName, g.VarName, g.EntityName, g.EntityName,
		g.VarName, g.EntityName, g.EntityName, g.VarName, g.EntityName, g.VarName, g.EntityName,
		g.VarName, g.EntityName, g.VarName, g.EntityName, g.VarName, g.VarName, g.VarName)

	contentStr = strings.Replace(contentStr, "\t// ISystemUser defines the interface for user management operations.",
		interfaceDef+"\t// ISystemUser defines the interface for user management operations.", 1)

	// 2. 插入局部变量（在 localSystemUser 之前）
	localVar := fmt.Sprintf("\tlocal%s\t\t\t\t   I%s\n", g.EntityName, g.EntityName)
	contentStr = strings.Replace(contentStr, "\tlocalSystemUser                ISystemUser",
		localVar+"\tlocalSystemUser                ISystemUser", 1)

	// 3. 插入 getter 和 register 函数（在 SystemUser() 之前）
	getterFunc := fmt.Sprintf(`func %s() I%s {
	if local%s == nil {
		panic("implement not found for interface I%s, forgot register?")
	}
	return local%s
}

func Register%s(i I%s) {
	local%s = i
}

`, g.EntityName, g.EntityName, g.EntityName, g.EntityName, g.EntityName, g.EntityName, g.EntityName, g.EntityName)

	contentStr = strings.Replace(contentStr, "func SystemUser() ISystemUser {",
		getterFunc+"func SystemUser() ISystemUser {", 1)

	if err := os.WriteFile(serviceFile, []byte(contentStr), 0644); err != nil {
		return fmt.Errorf("写入 service 文件失败: %v", err)
	}

	fmt.Printf("  ✓ 已注册 service 接口：%s\n", g.EntityName)
	return nil
}

// RegisterRouter 自动在 router/system/router.go 中注册路由
func (g *CRUDGenerator) RegisterRouter() error {
	routerFile := filepath.Join(g.WorkDir, "modules", g.ModuleName, "router", "system", "router.go")
	if !utils.PathExists(routerFile) {
		return fmt.Errorf("router 文件不存在: %s", routerFile)
	}

	content, err := os.ReadFile(routerFile)
	if err != nil {
		return fmt.Errorf("读取 router 文件失败: %v", err)
	}

	contentStr := string(content)

	// 检查是否已注册
	controllerName := fmt.Sprintf("system.%sController", g.EntityName)
	if strings.Contains(contentStr, controllerName) {
		fmt.Printf("  ℹ️ 路由已存在，跳过\n")
		return nil
	}

	// 在 DashboardController 之前插入新的 controller
	newController := fmt.Sprintf("\t\t\tsystem.%sController,\n", g.EntityName)
	contentStr = strings.Replace(contentStr, "\t\t\tsystem.DashboardController,",
		newController+"\t\t\tsystem.DashboardController,", 1)

	if err := os.WriteFile(routerFile, []byte(contentStr), 0644); err != nil {
		return fmt.Errorf("写入 router 文件失败: %v", err)
	}

	fmt.Printf("  ✓ 已注册路由：%s\n", g.EntityName)
	return nil
}
