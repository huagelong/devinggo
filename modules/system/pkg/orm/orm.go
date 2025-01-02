// Package orm
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package orm

import (
	"context"
	"devinggo/modules/system/pkg/utils/config"
	"devinggo/modules/system/pkg/utils/slice"
	"github.com/gogf/gf/v2/database/gdb"
	"reflect"
	"time"
)

func GetTableName(m *gdb.Model) string {
	v := reflect.ValueOf(m).Elem()
	field := v.FieldByName("tablesInit")
	return field.String()
}

func SetCacheOption(ctx context.Context, duration ...time.Duration) gdb.CacheOption {
	globalCache := config.GetConfigBool(ctx, "settings.enableGlobalDbCache", false)
	var dura time.Duration
	if globalCache {
		if len(duration) > 0 {
			dura = duration[0]
		} else {
			dura = time.Hour * 24
		}
	} else {
		dura = -1
	}
	return gdb.CacheOption{Duration: dura, Force: false}
}

func GetTableFieds(m *gdb.Model) []string {
	return slice.EscapeFieldsToSlice(m.GetFieldsStr())
}
