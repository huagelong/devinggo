// Package system
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package system

import (
	"devinggo/internal/dao"
	"devinggo/internal/logic/base"
	"devinggo/internal/model"
	"devinggo/internal/model/do"
	"devinggo/modules/system/model/req"
	"devinggo/modules/system/model/res"
	"devinggo/modules/system/pkg/hook"
	"devinggo/modules/system/pkg/orm"
	"devinggo/modules/system/pkg/utils"
	"devinggo/modules/system/service"
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/util/gconv"
)

type sSettingConfigGroup struct {
	base.BaseService
}

func init() {
	service.RegisterSettingConfigGroup(NewSystemSettingConfigGroup())
}

func NewSystemSettingConfigGroup() *sSettingConfigGroup {
	return &sSettingConfigGroup{}
}

func (s *sSettingConfigGroup) Model(ctx context.Context) *gdb.Model {
	return dao.SettingConfigGroup.Ctx(ctx).Hook(hook.Bind()).Cache(orm.SetCacheOption(ctx))
}

func (s *sSettingConfigGroup) GetList(ctx context.Context) (out []*res.SettingConfigGroup, err error) {
	inReq := &model.ListReq{}
	m := s.Model(ctx)
	m = orm.GetList(m, inReq)
	err = m.Scan(&out)
	if utils.IsError(err) {
		return
	}
	return
}

func (s *sSettingConfigGroup) SaveConfigGroup(ctx context.Context, data *req.SettingConfigGroupSave) (id uint64, err error) {
	saveData := do.SettingConfigGroup{
		Name:   data.Name,
		Code:   data.Code,
		Remark: data.Remark,
	}
	rs, err := s.Model(ctx).Data(saveData).Save()
	if utils.IsError(err) {
		return
	}
	tmpId, err := rs.LastInsertId()
	if err != nil {
		return
	}
	id = gconv.Uint64(tmpId)
	return
}

func (s *sSettingConfigGroup) UpdateConfigGroup(ctx context.Context, data *req.SettingConfigGroupUpdate) (err error) {
	saveData := do.SettingConfigGroup{
		Name:   data.Name,
		Code:   data.Code,
		Remark: data.Remark,
	}
	_, err = s.Model(ctx).Where(dao.SettingConfigGroup.Columns().Id, data.Id).Data(saveData).Update()
	if utils.IsError(err) {
		return err
	}
	return
}

func (s *sSettingConfigGroup) DeleteConfigGroup(ctx context.Context, id uint64) (err error) {
	_, err = s.Model(ctx).Where("id", id).Delete()
	if utils.IsError(err) {
		return err
	}
	_, err = service.SettingConfig().Model(ctx).Where("group_id", id).Delete()
	if utils.IsError(err) {
		return err
	}
	return
}