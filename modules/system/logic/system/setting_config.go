// Package system
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package system

import (
	"context"
	"devinggo/internal/dao"
	"devinggo/internal/model/do"
	"devinggo/internal/model/entity"
	"devinggo/modules/system/logic/base"
	"devinggo/modules/system/model"
	"devinggo/modules/system/model/req"
	"devinggo/modules/system/model/res"
	"devinggo/modules/system/pkg/hook"
	"devinggo/modules/system/pkg/orm"
	"devinggo/modules/system/pkg/utils"
	"devinggo/modules/system/service"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

type sSettingConfig struct {
	base.BaseService
}

func init() {
	service.RegisterSettingConfig(NewSystemSettingConfig())
}

func NewSystemSettingConfig() *sSettingConfig {
	return &sSettingConfig{}
}

func (s *sSettingConfig) Model(ctx context.Context) *gdb.Model {
	return dao.SettingConfig.Ctx(ctx).Hook(hook.Bind()).Cache(orm.SetCacheOption(ctx)).OnConflict("key")
}

func (s *sSettingConfig) GetConfigByKey(ctx context.Context, key string, groupKey ...string) (rs string, err error) {
	gkey := ""
	var groupId int64
	if !g.IsEmpty(groupKey) && len(groupKey) > 0 {
		gkey = groupKey[0]
		var settingConfigGroup *entity.SettingConfigGroup
		err = service.SettingConfigGroup().Model(ctx).Where(dao.SettingConfigGroup.Columns().Code, gkey).Scan(&settingConfigGroup)
		if utils.IsError(err) {
			return "", err
		}
		groupId = settingConfigGroup.Id
	}
	var settingConfig *entity.SettingConfig
	m := s.Model(ctx).Where(dao.SettingConfig.Columns().Key, key)
	if !g.IsEmpty(groupId) {
		m = m.Where("group_id", groupId)
	}
	err = m.Scan(&settingConfig)
	if utils.IsError(err) {
		return "", err
	}

	if !g.IsEmpty(settingConfig) {
		rs = settingConfig.Value
	} else {
		rs = ""
	}
	return
}

func (s *sSettingConfig) GetList(ctx context.Context, in *req.SettingConfigSearch) (out []*res.SettingConfig, err error) {
	inReq := &model.ListReq{
		OrderBy:   "sort",
		OrderType: "desc",
	}
	m := s.handleSearch(ctx, in)
	m = orm.GetList(m, inReq)
	err = m.Scan(&out)
	if utils.IsError(err) {
		return
	}
	return
}

func (s *sSettingConfig) SaveConfig(ctx context.Context, data *req.SettingConfigSave) (id int64, err error) {
	saveData := do.SettingConfig{
		Name:             data.Name,
		GroupId:          data.GroupId,
		Key:              data.Key,
		Value:            data.Value,
		InputType:        data.InputType,
		ConfigSelectData: data.ConfigSelectData,
		Sort:             data.Sort,
		Remark:           data.Remark,
	}
	rs, err := s.Model(ctx).Data(saveData).Insert()
	if utils.IsError(err) {
		return
	}
	tmpId, err := rs.LastInsertId()
	if err != nil {
		return
	}
	id = gconv.Int64(tmpId)
	return
}

func (s *sSettingConfig) UpdateConfig(ctx context.Context, data *req.SettingConfigUpdate) (err error) {
	saveData := do.SettingConfig{
		Name:             data.Name,
		Key:              data.Key,
		Value:            data.Value,
		InputType:        data.InputType,
		ConfigSelectData: data.ConfigSelectData,
		Sort:             data.Sort,
		Remark:           data.Remark,
	}
	_, err = s.Model(ctx).Where(dao.SettingConfig.Columns().Key, data.Key).Data(saveData).Update()
	if utils.IsError(err) {
		return err
	}
	return
}

func (s *sSettingConfig) DeleteConfig(ctx context.Context, ids []string) (err error) {
	_, err = s.Model(ctx).WhereIn("key", ids).Delete()
	if utils.IsError(err) {
		return err
	}
	return
}

func (s *sSettingConfig) handleSearch(ctx context.Context, in *req.SettingConfigSearch) (m *gdb.Model) {

	m = s.Model(ctx)

	if !g.IsEmpty(in.GroupId) {
		m = m.Where("group_id", in.GroupId)
	}

	if !g.IsEmpty(in.Name) {
		m = m.Where("name", in.Name)
	}

	if !g.IsEmpty(in.Key) {
		m = m.Where("key like ? ", "%"+in.Key+"%")
	}
	return
}
