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
	"devinggo/modules/system/service"
	"github.com/gogf/gf/v2/database/gdb"
)

type sSystemRoleDept struct {
	base.BaseService
}

func init() {
	service.RegisterSystemRoleDept(NewSystemRoleDept())
}

func NewSystemRoleDept() *sSystemRoleDept {
	return &sSystemRoleDept{}
}

func (s *sSystemRoleDept) Model(ctx context.Context) *gdb.Model {
	return dao.SystemRoleDept.Ctx(ctx).Hook(hook.Bind()).Cache(orm.SetCacheOption(ctx)).OnConflict("role_id", "dept_id")
}
