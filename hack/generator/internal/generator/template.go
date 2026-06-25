// Package generator
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package generator

import (
	"bytes"
	"strings"
	"text/template"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gfile"
)

// RenderTemplate 渲染模板文件
// templatePath: 模板文件路径
// data: 模板数据
// 返回值: 渲染后的内容
func RenderTemplate(templatePath string, data interface{}) (string, error) {
	// 检查模板文件是否存在
	if !gfile.Exists(templatePath) {
		return "", gerror.Newf("模板文件不存在: %s", templatePath)
	}

	// 读取模板内容
	templateContent := gfile.GetContents(templatePath)
	if templateContent == "" {
		return "", gerror.Newf("模板文件为空: %s", templatePath)
	}

	// 创建模板
	tmpl, err := template.New(gfile.Basename(templatePath)).Parse(templateContent)
	if err != nil {
		return "", gerror.Wrapf(err, "解析模板失败: %s", templatePath)
	}

	// 渲染模板
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", gerror.Wrapf(err, "渲染模板失败: %s", templatePath)
	}

	return buf.String(), nil
}

// RenderTemplateString 渲染模板字符串
// templateStr: 模板字符串
// data: 模板数据
// 返回值: 渲染后的内容
func RenderTemplateString(templateStr string, data interface{}) (string, error) {
	// 创建模板
	tmpl, err := template.New("template").Parse(templateStr)
	if err != nil {
		return "", gerror.Wrap(err, "解析模板失败")
	}

	// 渲染模板
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", gerror.Wrap(err, "渲染模板失败")
	}

	return buf.String(), nil
}

// RenderFrontendTemplate 渲染前端模板（处理Vue的{{ }}与Go模板冲突）
// templatePath: 模板文件路径
// data: 模板数据
// 返回值: 渲染后的内容
func RenderFrontendTemplate(templatePath string, data interface{}) (string, error) {
	// 检查模板文件是否存在
	if !gfile.Exists(templatePath) {
		return "", gerror.Newf("模板文件不存在: %s", templatePath)
	}

	// 读取模板内容
	templateContent := gfile.GetContents(templatePath)
	if templateContent == "" {
		return "", gerror.Newf("模板文件为空: %s", templatePath)
	}

	// 将Vue的{{ }}替换为临时占位符，避免与Go模板冲突
	vueLBrace := "__VUE_LBRACE__"
	vueRBrace := "__VUE_RBRACE__"
	templateContent = strings.ReplaceAll(templateContent, "{{", vueLBrace)
	templateContent = strings.ReplaceAll(templateContent, "}}", vueRBrace)

	// 创建模板，使用自定义分隔符 <%, %>
	tmpl, err := template.New(gfile.Basename(templatePath)).Delims("<%", "%>").Parse(templateContent)
	if err != nil {
		return "", gerror.Wrapf(err, "解析前端模板失败: %s", templatePath)
	}

	// 渲染模板
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", gerror.Wrapf(err, "渲染前端模板失败: %s", templatePath)
	}

	// 将占位符恢复为Vue的{{ }}
	result := buf.String()
	result = strings.ReplaceAll(result, vueLBrace, "{{")
	result = strings.ReplaceAll(result, vueRBrace, "}}")

	return result, nil
}
