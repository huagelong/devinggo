// Package system
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package system

import (
	"context"
	"devinggo/internal/dao"
	"devinggo/internal/model/entity"
	"devinggo/modules/system/logic/base"
	"devinggo/modules/system/pkg/hook"
	"devinggo/modules/system/pkg/orm"
	"devinggo/modules/system/pkg/utils"
	"devinggo/modules/system/service"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

type sSystemUserDept struct {
	base.BaseService
}

func init() {
	service.RegisterSystemUserDept(NewSystemUserDept())
}

func NewSystemUserDept() *sSystemUserDept {
	return &sSystemUserDept{}
}

func (s *sSystemUserDept) Model(ctx context.Context) *gdb.Model {
	return dao.SystemUserDept.Ctx(ctx).Hook(hook.Bind()).Cache(orm.SetCacheOption(ctx)).OnConflict("user_id", "dept_id")
}

func (s *sSystemUserDept) GetDeptIdsByUserId(ctx context.Context, userId int64) (deptIds []int64, err error) {
	var list []*entity.SystemUserDept
	err = s.Model(ctx).Fields(dao.SystemUserDept.Columns().DeptId).Where(dao.SystemUserDept.Columns().UserId, userId).Scan(&list)
	if utils.IsError(err) {
		return
	}
	if g.IsEmpty(list) {
		return
	}
	for _, item := range list {
		deptIds = append(deptIds, item.DeptId)
	}
	return
}
