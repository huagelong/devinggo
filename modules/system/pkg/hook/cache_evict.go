// Package hook
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package hook

import (
	"context"
	"database/sql"
	"devinggo/modules/system/pkg/cache"
	"fmt"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
)

type iHookInput interface {
	IsTransaction() bool

	Next(ctx context.Context) (result sql.Result, err error)
}

func CleanCache[T gdb.HookInsertInput | gdb.HookUpdateInput | gdb.HookDeleteInput](ctx context.Context, in *T) (err error) {
	// 将输入参数 in 转换为 iHookInput 接口类型，以便调用 Next 方法。
	_, ok := interface{}(in).(iHookInput)
	if !ok {
		// 如果转换失败，返回错误。
		return fmt.Errorf("input does not implement iHookInput")
	}

	var table string
	// 根据输入参数的不同类型，清理相应的缓存。
	if input, ok := interface{}(in).(*gdb.HookInsertInput); ok == true {
		table = input.Table
	} else if input, ok := interface{}(in).(*gdb.HookUpdateInput); ok == true {
		table = input.Table
	} else if input, ok := interface{}(in).(*gdb.HookDeleteInput); ok == true {
		table = input.Table
	}
	// 清理完缓存后，如果有表名，则进行表名的格式化处理,根据表名从缓存中移除对应的缓存项。
	if !g.IsEmpty(table) {
		table = gstr.SplitAndTrim(table, " ")[0]
		table = gstr.SplitAndTrim(table, ",")[0]
		table = gstr.Replace(table, "\"", "")
		table = gstr.Replace(table, "`", "")
		cache.ClearByTable(ctx, table)
		g.Log().Debug(ctx, "clean cache by table:", table)
	}
	return
}
