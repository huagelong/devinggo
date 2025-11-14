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

type sSystemUserRole struct {
	base.BaseService
}

func init() {
	service.RegisterSystemUserRole(NewSystemUserRole())
}

func NewSystemUserRole() *sSystemUserRole {
	return &sSystemUserRole{}
}

func (s *sSystemUserRole) Model(ctx context.Context) *gdb.Model {
	return dao.SystemUserRole.Ctx(ctx).Hook(hook.Bind()).Cache(orm.SetCacheOption(ctx)).OnConflict("user_id", "role_id")
}

func (s *sSystemUserRole) GetRoleIdsByUserId(ctx context.Context, userId int64) (roleIds []int64, err error) {
	var list []*entity.SystemUserRole
	err = s.Model(ctx).Fields(dao.SystemUserRole.Columns().RoleId).Where(dao.SystemUserRole.Columns().UserId, userId).Scan(&list)
	if utils.IsError(err) {
		return
	}
	if g.IsEmpty(list) {
		return
	}
	for _, item := range list {
		roleIds = append(roleIds, item.RoleId)
	}
	return
}
