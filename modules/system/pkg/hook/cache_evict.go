// Package hook
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package hook

import (
	"context"
	"fmt"

	"devinggo/modules/system/pkg/cache"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
)

// CleanCache 清理指定表的缓存
// 支持 Insert、Update、Delete 操作的钩子输入
func CleanCache[T gdb.HookInsertInput | gdb.HookUpdateInput | gdb.HookDeleteInput](ctx context.Context, in *T) error {
	var table string

	// 根据输入参数的不同类型，获取表名
	switch input := any(in).(type) {
	case *gdb.HookInsertInput:
		table = input.Table
	case *gdb.HookUpdateInput:
		table = input.Table
	case *gdb.HookDeleteInput:
		table = input.Table
	default:
		return fmt.Errorf("unsupported hook input type")
	}

	// 如果表名为空，直接返回
	if g.IsEmpty(table) {
		return nil
	}

	// 格式化表名：移除空格、逗号、引号等
	table = normalizeTableName(table)

	// 根据表名清理缓存
	_ = cache.ClearByTable(ctx, table)
	g.Log().Debug(ctx, "cache cleared for table:", table)

	return nil
}

// normalizeTableName 规范化表名，移除多余的字符
func normalizeTableName(table string) string {
	// 移除空格后的部分
	table = gstr.SplitAndTrim(table, " ")[0]
	// 移除逗号后的部分
	table = gstr.SplitAndTrim(table, ",")[0]
	// 移除引号
	table = gstr.Replace(table, "\"", "")
	table = gstr.Replace(table, "`", "")
	return table
}
