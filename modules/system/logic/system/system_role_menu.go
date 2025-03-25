// Package system
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package system

import (
	"context"
	"devinggo/internal/dao"
	"devinggo/modules/system/logic/base"
	"devinggo/modules/system/pkg/hook"
	"devinggo/modules/system/pkg/orm"
	"devinggo/modules/system/pkg/utils"
	"devinggo/modules/system/service"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

type sSystemRoleMenu struct {
	base.BaseService
}

func init() {
	service.RegisterSystemRoleMenu(NewSystemRoleMenu())
}

func NewSystemRoleMenu() *sSystemRoleMenu {
	return &sSystemRoleMenu{}
}

func (s *sSystemRoleMenu) Model(ctx context.Context) *gdb.Model {
	return dao.SystemRoleMenu.Ctx(ctx).Hook(hook.Bind()).Cache(orm.SetCacheOption(ctx)).OnConflict("role_id", "menu_id")
}

func (s *sSystemRoleMenu) GetMenuIdsByRoleIds(ctx context.Context, roleIds []int64) (rmenuIds []int64, err error) {
	menuIdsResult, err := s.Model(ctx).Fields("menu_id").WhereIn(dao.SystemRoleMenu.Columns().RoleId, roleIds).Array()
	if utils.IsError(err) {
		return
	}
	if g.IsEmpty(menuIdsResult) {
		return
	}

	menuIds := gconv.SliceInt64(menuIdsResult)
	rmenuIdsResult, err := service.SystemMenu().Model(ctx).Fields("id").WhereIn(dao.SystemMenu.Columns().Id, menuIds).Where(dao.SystemMenu.Columns().Status, 1).OrderDesc(dao.SystemMenu.Columns().Sort).Array()
	if utils.IsError(err) {
		return
	}

	if g.IsEmpty(rmenuIdsResult) {
		return
	}

	rmenuIds = gconv.SliceInt64(rmenuIdsResult)
	return
}
