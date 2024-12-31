// Package slice
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package slice

import (
	"fmt"
	"github.com/gogf/gf/v2/text/gstr"
	"strings"
)

// 切片去重
func Unique[K comparable](languages []K) []K {
	result := make([]K, 0, len(languages))
	temp := map[K]struct{}{}
	for _, item := range languages {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}

// DifferenceSlice 比较两个切片，返回他们的差集
// slice1 := []int{1, 2, 3, 4, 5}
// slice2 := []int{4, 5, 6, 7, 8}
// fmt.Println(Difference(slice1, slice2)) // Output: [1 2 3]
func Difference[T comparable](s1, s2 []T) []T {
	m := make(map[T]bool)
	for _, item := range s1 {
		m[item] = true
	}

	var diff []T
	for _, item := range s2 {
		if _, ok := m[item]; !ok {
			diff = append(diff, item)
		}
	}
	return diff
}

// 是否包含
func Contains[T comparable](slice []T, str T) bool {
	for _, s := range slice {
		if s == str {
			return true
		}
	}
	return false
}

// 从切片中删除第一个匹配的元素
func Remove[T comparable](slice []T, element T) []T {
	for i, v := range slice {
		if v == element {
			// 创建一个新的切片，包含除了要删除的元素之外的所有元素
			return append(slice[:i], slice[i+1:]...)
		}
	}
	// 如果没有找到元素，返回原始切片的副本
	return slice
}

// 字符串转slice，逗号分隔
func EscapeFieldsToSlice(s string) []string {
	return gstr.Explode(",", gstr.Replace(gstr.Replace(s, "`,`", ","), "`", ""))
}

// 字符串转数组 逗号分隔
func Explode(s string, sep string) []string {
	return gstr.Explode(sep, s)
}

// 数组转字符串
func Join(ints []string, separator string) string {
	// 使用 strings.Builder 来构建字符串
	var builder strings.Builder
	for i, v := range ints {
		builder.WriteString(fmt.Sprintf("%s", v))
		if i < len(ints)-1 {
			builder.WriteString(separator)
		}
	}
	return builder.String()
}

// Paginate 函数将数组分成多个页面，每个页面包含 pageSize 个元素，并支持当前页
func Paginate[T any](arr []T, pageSize int, currentPage int) ([]T, error) {
	var result []T

	// 检查当前页是否有效
	if currentPage < 1 {
		return result, nil
	}

	// 如果数组为空，直接返回空结果
	if len(arr) == 0 {
		return result, nil
	}

	// 计算总页数
	totalPages := (len(arr) + pageSize - 1) / pageSize

	// 如果当前页超出总页数，返回错误
	if currentPage > totalPages {
		return nil, nil
	}

	// 计算当前页的起始和结束索引
	start := (currentPage - 1) * pageSize
	end := start + pageSize

	// 如果 end 超出数组长度，则将其设置为数组长度
	if end > len(arr) {
		end = len(arr)
	}

	// 将当前页的元素添加到结果中
	result = append(result, arr[start:end]...)

	return result, nil
}
