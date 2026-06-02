// Package utils
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package utils

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

// HookExecutor 钩子执行器
type HookExecutor struct {
	verbose bool
	dryRun  bool
	timeout time.Duration
}

// NewHookExecutor 创建钩子执行器
func NewHookExecutor() *HookExecutor {
	return &HookExecutor{
		verbose: true,
		dryRun:  false,
		timeout: 30 * time.Second,
	}
}

// SetVerbose 设置是否输出详细日志
func (h *HookExecutor) SetVerbose(verbose bool) {
	h.verbose = verbose
}

// SetDryRun 设置是否为演练模式
func (h *HookExecutor) SetDryRun(dryRun bool) {
	h.dryRun = dryRun
}

// SetTimeout 设置命令超时时间
func (h *HookExecutor) SetTimeout(timeout time.Duration) {
	h.timeout = timeout
}

// HookCommand 钩子命令定义
type HookCommand struct {
	Name        string            `json:"name"`
	Command     string            `json:"command"`
	WorkDir     string            `json:"workDir"`
	Env         map[string]string `json:"env"`
	IgnoreError bool              `json:"ignoreError"`
}

// HookResult 钩子执行结果
type HookResult struct {
	Name     string
	Command  string
	Success  bool
	Output   string
	Error    string
	Duration time.Duration
}

// ExecuteHook 执行单个钩子命令
func (h *HookExecutor) ExecuteHook(hook HookCommand) *HookResult {
	result := &HookResult{
		Name:    hook.Name,
		Command: hook.Command,
	}

	startTime := time.Now()
	defer func() {
		result.Duration = time.Since(startTime)
	}()

	if h.verbose {
		fmt.Printf("  🔧 执行钩子: %s\n", hook.Name)
		if hook.Command != "" {
			fmt.Printf("     命令: %s\n", hook.Command)
		}
	}

	// 演练模式
	if h.dryRun {
		result.Success = true
		result.Output = "[演练模式] 命令未执行"
		if h.verbose {
			fmt.Println("     ⚠️  演练模式：命令未实际执行")
		}
		return result
	}

	// 空命令
	if hook.Command == "" {
		result.Success = true
		return result
	}

	// 解析命令
	parts := strings.Fields(hook.Command)
	if len(parts) == 0 {
		result.Success = false
		result.Error = "命令为空"
		return result
	}

	// 确定工作目录
	workDir := hook.WorkDir
	if workDir == "" {
		workDir, _ = os.Getwd()
	}

	// 创建命令
	cmd := exec.Command(parts[0], parts[1:]...)
	cmd.Dir = workDir

	// 设置环境变量
	if len(hook.Env) > 0 {
		env := os.Environ()
		for k, v := range hook.Env {
			env = append(env, fmt.Sprintf("%s=%s", k, v))
		}
		cmd.Env = env
	}

	// 执行命令
	output, err := cmd.CombinedOutput()
	result.Output = string(output)

	if err != nil {
		result.Success = false
		result.Error = err.Error()

		if !hook.IgnoreError {
			if h.verbose {
				fmt.Printf("     ❌ 失败: %s\n", err.Error())
				if len(output) > 0 {
					fmt.Printf("     输出: %s\n", strings.TrimSpace(string(output)))
				}
			}
		} else {
			result.Success = true
			if h.verbose {
				fmt.Printf("     ⚠️  失败但已忽略: %s\n", err.Error())
			}
		}
		return result
	}

	result.Success = true
	if h.verbose && len(output) > 0 {
		fmt.Printf("     输出: %s\n", strings.TrimSpace(string(output)))
	}

	return result
}

// ExecuteHooks 执行多个钩子命令
func (h *HookExecutor) ExecuteHooks(hooks []HookCommand, stage string) []*HookResult {
	if len(hooks) == 0 {
		return nil
	}

	if h.verbose {
		fmt.Printf("\n📋 执行 %s 钩子 (%d 个)...\n", stage, len(hooks))
	}

	results := make([]*HookResult, 0, len(hooks))
	for _, hook := range hooks {
		result := h.ExecuteHook(hook)
		results = append(results, result)

		// 如果钩子执行失败且不忽略错误，停止执行
		if !result.Success && !h.dryRun {
			// 检查是否应该忽略错误
			shouldIgnore := false
			for _, h := range hooks {
				if h.Name == hook.Name && h.IgnoreError {
					shouldIgnore = true
					break
				}
			}

			if !shouldIgnore {
				if h.verbose {
					fmt.Printf("  ❌ 钩子执行失败，中止后续操作\n")
				}
				return results
			}
		}
	}

	return results
}

// ExecuteHooksInDir 在指定目录中执行钩子
func (h *HookExecutor) ExecuteHooksInDir(hooks []HookCommand, stage, workDir string) []*HookResult {
	// 保存当前目录
	originalDir, _ := os.Getwd()
	defer func() {
		_ = os.Chdir(originalDir)
	}()

	// 切换到工作目录
	if workDir != "" {
		if err := os.Chdir(workDir); err != nil {
			if h.verbose {
				fmt.Printf("  ❌ 无法切换到工作目录 %s: %v\n", workDir, err)
			}
			return nil
		}
	}

	return h.ExecuteHooks(hooks, stage)
}

// ExecuteScriptFile 执行脚本文件
func (h *HookExecutor) ExecuteScriptFile(scriptPath string) error {
	if h.verbose {
		fmt.Printf("  🔧 执行脚本: %s\n", scriptPath)
	}

	if h.dryRun {
		if h.verbose {
			fmt.Println("     ⚠️  演练模式：脚本未实际执行")
		}
		return nil
	}

	// 检查文件是否存在
	if _, err := os.Stat(scriptPath); os.IsNotExist(err) {
		return fmt.Errorf("脚本文件不存在: %s", scriptPath)
	}

	// 根据扩展名确定执行方式
	ext := strings.ToLower(filepath.Ext(scriptPath))
	var cmd *exec.Cmd

	switch ext {
	case ".sh":
		cmd = exec.Command("bash", scriptPath)
	case ".bash":
		cmd = exec.Command("bash", scriptPath)
	case ".bat", ".cmd":
		cmd = exec.Command("cmd", "/C", scriptPath)
	case ".ps1":
		cmd = exec.Command("powershell", "-File", scriptPath)
	case ".py":
		cmd = exec.Command("python", scriptPath)
	default:
		// 尝试直接执行
		cmd = exec.Command(scriptPath)
	}

	// 设置输出
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// 执行
	return cmd.Run()
}

// PrintHookSummary 打印钩子执行摘要
func (h *HookExecutor) PrintHookSummary(results []*HookResult) {
	if len(results) == 0 {
		return
	}

	fmt.Println("\n📊 钩子执行摘要:")

	successCount := 0
	for _, result := range results {
		if result.Success {
			successCount++
		}
	}

	fmt.Printf("  总计: %d, 成功: %d, 失败: %d\n", len(results), successCount, len(results)-successCount)

	if h.verbose {
		for _, result := range results {
			status := "✅"
			if !result.Success {
				status = "❌"
			}
			fmt.Printf("  %s %s (%.2fs)\n", status, result.Name, float64(result.Duration.Milliseconds())/1000)
		}
	}
}
