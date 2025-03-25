// Package orm
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package handler

import (
	"context"
	"devinggo/internal/dao"
	"devinggo/internal/model/entity"
	"devinggo/modules/system/consts"
	"devinggo/modules/system/pkg/contexts"
	"devinggo/modules/system/pkg/utils"
	"devinggo/modules/system/pkg/utils/slice"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"reflect"
)

// FilterAuth 过滤数据权限
// 通过上下文中的用户角色权限和表中是否含有需要过滤的字段附加查询条件
func FilterAuth(m *gdb.Model) *gdb.Model {
	var (
		needAuth    bool
		filterField string
		fields      = slice.EscapeFieldsToSlice(m.GetFieldsStr())
	)

	if gstr.InArray(fields, "created_by") {
		needAuth = true
		filterField = "created_by"
	}

	if !needAuth {
		return m
	}
	return m.Handler(FilterAuthWithField(filterField))
}

// FilterAuthWithField 过滤数据权限，设置指定字段
func FilterAuthWithField(filterField string) func(m *gdb.Model) *gdb.Model {
	return func(m *gdb.Model) *gdb.Model {
		var (
			roles []*entity.SystemRole
			ctx   = m.GetCtx()
			user  = contexts.New().GetUser(ctx)
		)

		if user == nil {
			return m
		}

		getUserRoleIds := func(ctx context.Context, userId int64) (roles []int64) {
			result, err := dao.SystemUserRole.Ctx(ctx).Fields(dao.SystemUserRole.Columns().RoleId).Where(dao.SystemUserRole.Columns().UserId, userId).Array()
			if utils.IsError(err) {
				g.Log().Panicf(ctx, "get user roleIds err:%+v", err)
				return
			}

			if g.IsEmpty(result) {
				return
			}

			roles = gconv.SliceInt64(result)
			return
		}

		getUserDeptIds := func(ctx context.Context, userId int64) (depts []int64) {
			result, err := dao.SystemUserDept.Ctx(ctx).Fields(dao.SystemUserDept.Columns().DeptId).Where(dao.SystemUserDept.Columns().UserId, userId).Array()
			if utils.IsError(err) {
				g.Log().Panicf(ctx, "get user deptIds err:%+v", err)
				return
			}

			if g.IsEmpty(result) {
				return
			}

			depts = gconv.SliceInt64(result)
			return
		}

		user.RoleIds = getUserRoleIds(ctx, user.Id)
		user.DeptIds = getUserDeptIds(ctx, user.Id)

		err := dao.SystemRole.Ctx(ctx).WhereIn(dao.SystemRole.Columns().Id, user.RoleIds).Scan(&roles)
		if err != nil {
			g.Log().Panicf(ctx, "failed to role information err:%+v", err)
		}

		if roles == nil {
			g.Log().Panic(ctx, "failed to role information roleModel == nil")
		}

		// 超管拥有全部权限
		for _, role := range roles {
			if role.Code == consts.SuperRoleKey {
				return m
			}
		}

		getFromDeptIds := func(ctx context.Context, in []int64) []int64 {
			result, err := dao.SystemUserDept.Ctx(ctx).Fields(dao.SystemUserDept.Columns().UserId).WhereIn(dao.SystemUserDept.Columns().DeptId, in).Array()
			if utils.IsError(err) {
				g.Log().Panic(ctx, "failed to get member dept data", err)
			}

			if g.IsEmpty(result) {
				g.Log().Debug(ctx, "getFromDeptIds userIds is null")
				return nil
			}

			return gconv.SliceInt64(result)
		}

		getFromRoles := func(ctx context.Context, in []int64) []int64 {
			deptIds, err := dao.SystemRoleDept.Ctx(ctx).Fields(dao.SystemRoleDept.Columns().DeptId).WhereIn(dao.SystemRoleDept.Columns().RoleId, in).Array()
			if utils.IsError(err) {
				g.Log().Panic(ctx, "failed to get role_dept dept data")
			}

			result, err := dao.SystemUserDept.Ctx(ctx).Fields(dao.SystemUserDept.Columns().UserId).WhereIn(dao.SystemUserDept.Columns().DeptId, deptIds).Array()
			if utils.IsError(err) {
				g.Log().Panic(ctx, "failed to get member dept data", err)
			}
			if g.IsEmpty(result) {
				g.Log().Debug(ctx, "getFromRoles userIds is null")
				return nil
			}

			return gconv.SliceInt64(result)
		}

		getDeptIdsAndSub := func(ctx context.Context, in []int64) []int64 {
			deptIds := make([]int64, 0)
			if len(in) > 0 {
				for _, deptId := range in {
					newDeptIds := make([]int64, 0)
					result, err := dao.SystemDept.Ctx(ctx).Fields(dao.SystemDept.Columns().Id).Where(dao.SystemDept.Columns().Id, deptId).WhereOr("level like  ? ", "%,"+gconv.String(deptId)+",%").Array()
					if utils.IsError(err) {
						g.Log().Panic(ctx, "failed to get system_dept dept data")
					}
					if g.IsEmpty(result) {
						continue
					}
					newDeptIds = gconv.SliceInt64(result)
					deptIds = append(deptIds, newDeptIds...)
				}
			}
			return deptIds
		}

		/**
		数据范围（1：全部数据权限 2：自定义数据权限 3：本部门数据权限 4：本部门及以下数据权限 5：本人数据权限）
		*/
		userIds := make([]int64, 0)
		for _, role := range roles {
			switch role.DataScope {
			case 1:
			case 2:
				userIds = append(userIds, getFromRoles(ctx, user.RoleIds)...)
			case 3:
				userIds = append(userIds, getFromDeptIds(ctx, user.DeptIds)...)
			case 4:
				userIds = append(userIds, getFromDeptIds(ctx, getDeptIdsAndSub(ctx, user.DeptIds))...)
			case 5:
				userIds = append(userIds, user.Id)
			case 6:
				deptIds := getDeptIdsAndSub(ctx, user.DeptIds)
				if !g.IsEmpty(deptIds) {
					tableName := getTableName(m)
					if tableName == "system_dept" {
						m = m.WhereIn("id", deptIds)
						break
					}
					if !gstr.InArray(getTableFieds(m), "dept_id") {
						break
					}
					m = m.WhereIn("dept_id", deptIds)
				}

			}
		}

		if len(userIds) == 0 {
			return m
		} else {
			m = m.WhereIn(filterField, userIds)
			return m
		}
	}
}

func getTableName(m *gdb.Model) string {
	v := reflect.ValueOf(m).Elem()
	field := v.FieldByName("tablesInit")
	return field.String()
}

func getTableFieds(m *gdb.Model) []string {
	return slice.EscapeFieldsToSlice(m.GetFieldsStr())
}
