// Package system
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package system

import (
	"devinggo/internal/dao"
	"devinggo/internal/logic/base"
	"devinggo/modules/system/pkg/hook"
	"devinggo/modules/system/pkg/orm"
	"devinggo/modules/system/service"
	"context"
	"github.com/gogf/gf/v2/database/gdb"
)

type sSystemAppApi struct {
	base.BaseService
}

func init() {
	service.RegisterSystemAppApi(NewSystemAppApi())
}

func NewSystemAppApi() *sSystemAppApi {
	return &sSystemAppApi{}
}

func (s *sSystemAppApi) Model(ctx context.Context) *gdb.Model {
	return dao.SystemAppApi.Ctx(ctx).Hook(hook.Bind()).Cache(orm.SetCacheOption(ctx))
}
