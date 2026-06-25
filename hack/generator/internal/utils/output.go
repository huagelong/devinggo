package utils

import (
	"encoding/json"
	"fmt"
	"os"
)

// OutputFormat 输出格式
type OutputFormat string

const (
	OutputFormatText OutputFormat = "text"
	OutputFormatJSON OutputFormat = "json"
)

// CommandResult 命令执行结果（用于结构化输出）
type CommandResult struct {
	Success  bool     `json:"success"`
	Message  string   `json:"message,omitempty"`
	Files    []string `json:"files,omitempty"`
	Warnings []string `json:"warnings,omitempty"`
	Errors   []string `json:"errors,omitempty"`
}

// NewCommandResult 创建命令结果
func NewCommandResult(success bool, message string) *CommandResult {
	return &CommandResult{
		Success:  success,
		Message:  message,
		Files:    []string{},
		Warnings: []string{},
		Errors:   []string{},
	}
}

// AddFile 添加生成的文件
func (r *CommandResult) AddFile(file string) {
	r.Files = append(r.Files, file)
}

// AddWarning 添加警告
func (r *CommandResult) AddWarning(warning string) {
	r.Warnings = append(r.Warnings, warning)
}

// AddError 添加错误
func (r *CommandResult) AddError(err string) {
	r.Errors = append(r.Errors, err)
}

// Print 输出结果
func (r *CommandResult) Print(format OutputFormat) {
	if format == OutputFormatJSON {
		r.printJSON()
		return
	}
	r.printText()
}

func (r *CommandResult) printJSON() {
	data, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, `{"success":false,"message":"JSON序列化失败: %s"}`, err.Error())
		os.Exit(1)
	}
	fmt.Println(string(data))
	if !r.Success {
		os.Exit(1)
	}
}

func (r *CommandResult) printText() {
	if r.Success {
		fmt.Printf("\n✅ %s\n", r.Message)
	} else {
		fmt.Printf("\n❌ %s\n", r.Message)
	}

	if len(r.Files) > 0 {
		fmt.Println("\n📁 生成的文件:")
		for _, f := range r.Files {
			fmt.Printf("   • %s\n", f)
		}
	}

	if len(r.Warnings) > 0 {
		fmt.Println("\n⚠️  警告:")
		for _, w := range r.Warnings {
			fmt.Printf("   • %s\n", w)
		}
	}

	if len(r.Errors) > 0 {
		fmt.Println("\n❌ 错误:")
		for _, e := range r.Errors {
			fmt.Printf("   • %s\n", e)
		}
	}

	fmt.Println()

	if !r.Success {
		os.Exit(1)
	}
}

// ParseOutputFormat 从字符串解析输出格式
func ParseOutputFormat(output string) OutputFormat {
	if output == "json" {
		return OutputFormatJSON
	}
	return OutputFormatText
}
