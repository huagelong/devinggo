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

func AutoCreatedUpdatedByInsert(ctx context.Context, in *gdb.HookInsertInput) (err error) {
	//g.Log().Debug(ctx, "GetTableFieds", orm.GetTableFieds(in.Model))
	hasCreatedBy := gstr.InArray(orm.GetTableFieds(in.Model), "created_by") || gstr.InArray(orm.GetTableFieds(in.Model), "\"created_by\"")
	hasUpdatedBy := gstr.InArray(orm.GetTableFieds(in.Model), "updated_by") || gstr.InArray(orm.GetTableFieds(in.Model), "\"updated_by\"")
	if hasCreatedBy || hasUpdatedBy {
		userId := contexts.New().GetUserId(ctx)
		//g.Log().Debug(ctx, "userId", userId)
		if !g.IsEmpty(in.Data) && !g.IsEmpty(userId) {
			for _, data := range in.Data {
				if hasCreatedBy {
					if _, ok := data["created_by"]; !ok {
						data["created_by"] = contexts.New().GetUserId(ctx)
					}
				}
				if hasUpdatedBy {
					if _, ok := data["updated_by"]; !ok {
						data["updated_by"] = contexts.New().GetUserId(ctx)
					}
				}
			}
		}
	}
	//g.Log().Debug(ctx, "GetTableFieds-in", in)
	return
}

func AutoCreatedUpdatedByUpdatefunc(ctx context.Context, in *gdb.HookUpdateInput) (err error) {
	//g.Log().Debug(ctx, "in", in)
	hasUpdatedBy := gstr.InArray(orm.GetTableFieds(in.Model), "updated_by") || gstr.InArray(orm.GetTableFieds(in.Model), "\"updated_by\"")
	if hasUpdatedBy {
		userId := contexts.New().GetUserId(ctx)
		if !g.IsEmpty(in.Data) && !g.IsEmpty(userId) {
			switch in.Data.(type) {
			case map[string]interface{}:
				//g.Log().Info(ctx, "map")
				if _, ok := in.Data.(map[string]interface{})["updated_by"]; !ok {
					in.Data.(map[string]interface{})["updated_by"] = userId
				}
			case string:
				in.Data = in.Data.(string) + ", " + utils.QuoteField("updated_by") + " = " + gconv.String(userId)
			}
		}
	}
	return
}
