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
	"devinggo/modules/system/pkg/handler"
	"devinggo/modules/system/pkg/orm"
	"devinggo/modules/system/pkg/utils"
	"devinggo/modules/system/service"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

type sSettingCrontabLog struct {
	base.BaseService
}

func init() {
	service.RegisterSettingCrontabLog(NewSystemSettingCrontabLog())
}

func NewSystemSettingCrontabLog() *sSettingCrontabLog {
	return &sSettingCrontabLog{}
}

func (s *sSettingCrontabLog) Model(ctx context.Context) *gdb.Model {
	return dao.SettingCrontabLog.Ctx(ctx).OnConflict("id")
}

func (s *sSettingCrontabLog) GetPageList(ctx context.Context, req *model.PageListReq, in *req.SettingCrontabLogSearch) (rs []*res.SettingCrontabLog, total int, err error) {
	m := s.handleSearch(ctx, in).Handler(handler.FilterAuth)
	var entity []*entity.SettingCrontabLog
	err = orm.GetPageList(m, req).ScanAndCount(&entity, &total, false)
	if utils.IsError(err) {
		return nil, 0, err
	}
	rs = make([]*res.SettingCrontabLog, 0)
	if !g.IsEmpty(entity) {
		if err = gconv.Structs(entity, &rs); err != nil {
			return nil, 0, err
		}
	}
	return
}

func (s *sSettingCrontabLog) handleSearch(ctx context.Context, in *req.SettingCrontabLogSearch) (m *gdb.Model) {

	m = s.Model(ctx)

	if !g.IsEmpty(in.CrontabId) {
		m = m.Where("crontab_id", in.CrontabId)
	}
	return
}

func (s *sSettingCrontabLog) Delete(ctx context.Context, ids []int64) (err error) {
	_, err = s.Model(ctx).WhereIn("id", ids).Delete()
	if utils.IsError(err) {
		return err
	}
	return
}

func (s *sSettingCrontabLog) AddLog(ctx context.Context, id int64, status int, exceptionInfo string) (err error) {
	var entity *entity.SettingCrontab
	err = service.SettingCrontab().Model(ctx).Where("id", id).Scan(&entity)
	if utils.IsError(err) {
		return err
	}
	_, err = s.Model(ctx).Insert(do.SettingCrontabLog{
		CrontabId:     entity.Id,
		Name:          entity.Name,
		Target:        entity.Target,
		Parameter:     entity.Parameter,
		Status:        status,
		ExceptionInfo: exceptionInfo,
	})
	if utils.IsError(err) {
		return err
	}
	return
}
