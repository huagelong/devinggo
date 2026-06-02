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

// UserRelate 为查询结果关联用户信息
// 根据指定的字段名查询对应的用户信息，并添加到结果集中
func UserRelate(ctx context.Context, result gdb.Result, fieldNames []string) (gdb.Result, error) {
	// 收集所有需要查询的用户ID
	memberIds := collectMemberIds(result, fieldNames)

	g.Log().Debug(ctx, "UserRelate memberIds:", memberIds)

	if len(memberIds) == 0 {
		return fillEmptyUserRelate(result, fieldNames), nil
	}

	// 去重
	memberIds = slice.Unique(memberIds)

	// 批量查询用户信息
	var members []*model.UserRelate
	if err := dao.SystemUser.Ctx(ctx).Unscoped().WhereIn(dao.SystemUser.Columns().Id, memberIds).Scan(&members); err != nil {
		return result, err
	}

	if g.IsEmpty(members) {
		return fillEmptyUserRelate(result, fieldNames), nil
	}

	// 构建用户ID到用户信息的映射
	memberMap := buildMemberMap(members)

	// 填充用户关联信息
	fillUserRelate(result, fieldNames, memberMap)

	return result, nil
}

// collectMemberIds 从结果集中收集所有需要查询的用户ID
func collectMemberIds(result gdb.Result, fieldNames []string) []int64 {
	var memberIds []int64
	for _, record := range result {
		for _, fieldName := range fieldNames {
			if val, ok := record[fieldName]; ok && val.Int64() > 0 {
				memberIds = append(memberIds, val.Int64())
			}
		}
	}
	return memberIds
}

// buildMemberMap 构建用户ID到用户信息的映射
func buildMemberMap(members []*model.UserRelate) map[int64]*model.UserRelate {
	memberMap := make(map[int64]*model.UserRelate, len(members))
	for _, member := range members {
		memberMap[member.Id] = member
	}
	return memberMap
}

// fillUserRelate 填充用户关联信息到结果集
func fillUserRelate(result gdb.Result, fieldNames []string, memberMap map[int64]*model.UserRelate) {
	for _, record := range result {
		for _, fieldName := range fieldNames {
			relateName := fieldName + "_relate"
			if val, ok := record[fieldName]; ok {
				if member, exists := memberMap[val.Int64()]; exists {
					record[relateName] = gvar.New(member)
				} else {
					record[relateName] = gvar.New(g.Map{})
				}
			}
		}
	}
}

// fillEmptyUserRelate 为结果集填充空的用户关联信息
func fillEmptyUserRelate(result gdb.Result, fieldNames []string) gdb.Result {
	for _, record := range result {
		for _, fieldName := range fieldNames {
			if val, ok := record[fieldName]; ok && val.Int64() > 0 {
				relateName := fieldName + "_relate"
				record[relateName] = gvar.New(g.Map{})
			}
		}
	}
	return result
}
