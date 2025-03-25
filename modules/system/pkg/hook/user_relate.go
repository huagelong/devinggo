// Package hook
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package hook

import (
	"context"
	"devinggo/internal/dao"
	"devinggo/modules/system/model"
	"devinggo/modules/system/pkg/utils/slice"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

func UserRelate(ctx context.Context, result gdb.Result, fieldNames []string) (gdb.Result, error) {
	var (
		memberIds []int64
	)
	for _, record := range result {
		for _, fieldName := range fieldNames {
			if _, ok := record[fieldName]; ok && (record[fieldName].Int64() > 0) {
				memberIds = append(memberIds, record[fieldName].Int64())
			}
		}
	}

	if len(memberIds) == 0 {
		return result, nil
	}
	memberIds = slice.Unique(memberIds)

	var members []*model.UserRelate
	if err := dao.SystemUser.Ctx(ctx).Unscoped().WhereIn(dao.SystemUser.Columns().Id, memberIds).Scan(&members); err != nil {
		return result, err
	}

	if g.IsEmpty(members) {
		return emptyUserRelate(result, fieldNames), nil
	}

	findMember := func(id *gvar.Var) *model.UserRelate {
		for _, v := range members {
			if v.Id == id.Int64() {
				return v
			}
		}
		return nil
	}

	for _, record := range result {
		for _, fieldName := range fieldNames {
			cacheName := fieldName + "_relate"
			record[cacheName] = gvar.New(findMember(record[fieldName]))
		}
	}
	return result, nil
}

func emptyUserRelate(result gdb.Result, fieldNames []string) gdb.Result {
	for _, record := range result {
		for _, fieldName := range fieldNames {
			if _, ok := record[fieldName]; ok && (record[fieldName].Int64() > 0) {
				cacheName := fieldName + "_relate"
				record[cacheName] = gvar.New(g.Map{})
			}
		}
	}
	return result
}
