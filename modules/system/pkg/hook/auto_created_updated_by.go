// Package hook
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package hook

import (
	"context"

	"devinggo/modules/system/pkg/contexts"
	"devinggo/modules/system/pkg/orm"
	"devinggo/modules/system/pkg/utils"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

// AutoCreatedUpdatedByInsert 在插入数据时自动填充 created_by 和 updated_by 字段
func AutoCreatedUpdatedByInsert(ctx context.Context, in *gdb.HookInsertInput) (err error) {
	fields := orm.GetTableFieds(in.Model)
	hasCreatedBy := gstr.InArray(fields, "created_by") || gstr.InArray(fields, "\"created_by\"")
	hasUpdatedBy := gstr.InArray(fields, "updated_by") || gstr.InArray(fields, "\"updated_by\"")

	if !hasCreatedBy && !hasUpdatedBy {
		return nil
	}

	userId := contexts.GetUserId(ctx)
	if g.IsEmpty(in.Data) || g.IsEmpty(userId) {
		return nil
	}

	for _, data := range in.Data {
		if hasCreatedBy {
			if _, ok := data["created_by"]; !ok {
				data["created_by"] = userId
			}
		}
		if hasUpdatedBy {
			if _, ok := data["updated_by"]; !ok {
				data["updated_by"] = userId
			}
		}
	}
	return nil
}

// AutoCreatedUpdatedByUpdate 在更新数据时自动填充 updated_by 字段
func AutoCreatedUpdatedByUpdate(ctx context.Context, in *gdb.HookUpdateInput) (err error) {
	fields := orm.GetTableFieds(in.Model)
	hasUpdatedBy := gstr.InArray(fields, "updated_by") || gstr.InArray(fields, "\"updated_by\"")

	if !hasUpdatedBy {
		return nil
	}

	userId := contexts.GetUserId(ctx)
	if g.IsEmpty(in.Data) || g.IsEmpty(userId) {
		return nil
	}

	switch data := in.Data.(type) {
	case map[string]interface{}:
		if _, ok := data["updated_by"]; !ok {
			data["updated_by"] = userId
		}
	case string:
		in.Data = data + ", " + utils.QuoteField("updated_by") + " = " + gconv.String(userId)
	}
	return nil
}
