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

	"devinggo/hack/generator/internal/utils"

	"github.com/gogf/gf/v2/text/gstr"
)

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

	fmt.Printf("\n✓ CRUD代码生成完成！\n")
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
