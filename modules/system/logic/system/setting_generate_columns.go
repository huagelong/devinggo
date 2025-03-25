// Package system
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE
package system

import (
	"context"
	"devinggo/internal/dao"
	"devinggo/modules/system/logic/base"
	"devinggo/modules/system/model"
	"devinggo/modules/system/model/req"
	"devinggo/modules/system/model/res"
	"devinggo/modules/system/pkg/handler"
	"devinggo/modules/system/pkg/hook"
	"devinggo/modules/system/pkg/orm"
	"devinggo/modules/system/pkg/utils"
	"devinggo/modules/system/service"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

type sSettingGenerateColumns struct {
	base.BaseService
}

func init() {
	service.RegisterSettingGenerateColumns(NewSystemSettingGenerateColumns())
}

func NewSystemSettingGenerateColumns() *sSettingGenerateColumns {
	return &sSettingGenerateColumns{}
}

func (s *sSettingGenerateColumns) Model(ctx context.Context) *gdb.Model {
	return dao.SettingGenerateColumns.Ctx(ctx).Hook(hook.Bind()).Cache(orm.SetCacheOption(ctx)).OnConflict("id")
}

func (s *sSettingGenerateColumns) GetList(ctx context.Context, in *req.SettingGenerateColumnsSearch) (out []*res.SettingGenerateColumns, err error) {
	inReq := &model.ListReq{
		OrderBy:   "sort",
		OrderType: "desc",
	}
	m := s.handleSearch(ctx, in).Handler(handler.FilterAuth)
	m = orm.GetList(m, inReq)
	err = m.Scan(&out)
	if utils.IsError(err) {
		return
	}
	return
}

func (s *sSettingGenerateColumns) handleSearch(ctx context.Context, in *req.SettingGenerateColumnsSearch) (m *gdb.Model) {
	m = s.Model(ctx)
	if !g.IsEmpty(in.TableId) {
		m = m.Where("table_id", in.TableId)
	}
	return
}
