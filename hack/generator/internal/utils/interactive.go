// Package utils
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// PromptBool 提示用户输入布尔值
func PromptBool(prompt string, defaultValue bool) bool {
	reader := bufio.NewReader(os.Stdin)

	defaultStr := "y/N"
	if defaultValue {
		defaultStr = "Y/n"
	}

	for {
		fmt.Printf("%s [%s]: ", prompt, defaultStr)

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("读取输入失败，请重试")
			continue
		}

		input = strings.TrimSpace(input)

		// 空输入返回默认值
		if input == "" {
			return defaultValue
		}

		// 解析输入
		switch strings.ToLower(input) {
		case "y", "yes", "是", "1":
			return true
		case "n", "no", "否", "0":
			return false
		default:
			fmt.Println("请输入 y/yes 或 n/no")
		}
	}
}

// PromptString 提示用户输入字符串
func PromptString(prompt string, defaultValue string) string {
	reader := bufio.NewReader(os.Stdin)

	if defaultValue != "" {
		fmt.Printf("%s [%s]: ", prompt, defaultValue)
	} else {
		fmt.Printf("%s: ", prompt)
	}

	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("读取输入失败")
		return defaultValue
	}

	input = strings.TrimSpace(input)

	if input == "" {
		return defaultValue
	}

	return input
}

// PromptRequiredString 提示用户输入必填字符串
func PromptRequiredString(prompt string) string {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("%s: ", prompt)

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("读取输入失败，请重试")
			continue
		}

		input = strings.TrimSpace(input)

		if input == "" {
			fmt.Println("此项不能为空，请输入")
			continue
		}

		return input
	}
}

// PromptSelect 提示用户从选项中选择
func PromptSelect(prompt string, options []string, defaultIndex int) int {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("\n%s\n", prompt)
		for i, option := range options {
			prefix := " "
			if i == defaultIndex {
				prefix = "*"
			}
			fmt.Printf("  %s %d. %s\n", prefix, i+1, option)
		}

		fmt.Printf("请选择 [1-%d]: ", len(options))

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("读取输入失败，请重试")
			continue
		}

		input = strings.TrimSpace(input)

		// 空输入返回默认值
		if input == "" && defaultIndex >= 0 && defaultIndex < len(options) {
			return defaultIndex
		}

		// 解析数字
		var choice int
		if _, err := fmt.Sscanf(input, "%d", &choice); err != nil {
			fmt.Printf("请输入 1-%d 之间的数字\n", len(options))
			continue
		}

		if choice < 1 || choice > len(options) {
			fmt.Printf("请输入 1-%d 之间的数字\n", len(options))
			continue
		}

		return choice - 1
	}
}

// PromptConfirm 提示用户确认
func PromptConfirm(prompt string) bool {
	return PromptBool(prompt, false)
}
