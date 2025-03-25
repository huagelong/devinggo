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
	"devinggo/modules/system/myerror"
	"devinggo/modules/system/pkg/handler"
	"devinggo/modules/system/pkg/hook"
	"devinggo/modules/system/pkg/orm"
	"devinggo/modules/system/pkg/utils"
	"devinggo/modules/system/pkg/worker/cron"
	"devinggo/modules/system/pkg/worker/task"
	"devinggo/modules/system/service"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

type sSettingCrontab struct {
	base.BaseService
}

func init() {
	service.RegisterSettingCrontab(NewSystemSettingCrontab())
}

func NewSystemSettingCrontab() *sSettingCrontab {
	return &sSettingCrontab{}
}

func (s *sSettingCrontab) Model(ctx context.Context) *gdb.Model {
	return dao.SettingCrontab.Ctx(ctx).Hook(hook.Bind()).Cache(orm.SetCacheOption(ctx)).OnConflict("id")
}

func (s *sSettingCrontab) GetValidateCron(ctx context.Context) (rs []*res.SettingCrontabOne, err error) {
	err = s.Model(ctx).Where("status", 1).Scan(&rs)
	if utils.IsError(err) {
		return
	}
	return
}

func (s *sSettingCrontab) GetPageList(ctx context.Context, req *model.PageListReq, in *req.SettingCrontabSearch) (rs []*res.SettingCrontab, total int, err error) {
	m := s.handleSearch(ctx, in).Handler(handler.FilterAuth)
	var entity []*entity.SettingCrontab
	err = orm.GetPageList(m, req).ScanAndCount(&entity, &total, false)
	if utils.IsError(err) {
		return nil, 0, err
	}
	rs = make([]*res.SettingCrontab, 0)
	if !g.IsEmpty(entity) {
		if err = gconv.Structs(entity, &rs); err != nil {
			return nil, 0, err
		}
	}
	return
}

func (s *sSettingCrontab) handleSearch(ctx context.Context, in *req.SettingCrontabSearch) (m *gdb.Model) {

	m = s.Model(ctx)

	if !g.IsEmpty(in.Type) {
		m = m.Where("type", in.Type)
	}

	if !g.IsEmpty(in.Status) {
		m = m.Where("status", in.Status)
	}

	if !g.IsEmpty(in.Name) {
		m = m.Where("name like ? ", "%"+in.Name+"%")
	}
	if !g.IsEmpty(in.CreatedAt) {
		if len(in.CreatedAt) > 0 {
			m = m.WhereGTE("created_at", in.CreatedAt[0]+" 00:00:00")
		}
		if len(in.CreatedAt) > 1 {
			m = m.WhereLTE("created_at", in.CreatedAt[1]+"23:59:59")
		}
	}
	return
}

func (s *sSettingCrontab) Save(ctx context.Context, in *req.SettingCrontabSave) (id int64, err error) {
	saveData := do.SettingCrontab{
		Name:      in.Name,
		Type:      in.Type,
		Rule:      in.Rule,
		Target:    in.Target,
		Parameter: in.Parameter,
		Remark:    in.Remark,
		Status:    in.Status,
		Singleton: in.Singleton,
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

func (s *sSettingCrontab) GetById(ctx context.Context, id int64) (res *res.SettingCrontab, err error) {
	err = s.Model(ctx).Where("id", id).Scan(&res)
	if utils.IsError(err) {
		return
	}
	return
}

func (s *sSettingCrontab) Run(ctx context.Context, id int64) (err error) {
	var dbCrons *res.SettingCrontabOne
	err = s.Model(ctx).Where("id", id).Scan(&dbCrons)
	if utils.IsError(err) {
		return
	}
	if g.IsEmpty(dbCrons) {
		err = myerror.ValidationFailed(ctx, "定时任务不存在")
		return
	}
	bindCron := cron.GetWorkerList(dbCrons.Target)
	if g.IsEmpty(bindCron) {
		err = myerror.ValidationFailed(ctx, "任务标的不存在")
		return
	}
	bindCron.SetParams(ctx, dbCrons.Parameter)
	bindCron.GetPayload().CrontabId = dbCrons.Id
	if dbCrons.Singleton == 1 {
		bindCron.GetPayload().TaskID = dbCrons.Target + "_" + gconv.String(dbCrons.Id)
	}
	err = task.NewSimpleTask(ctx, bindCron)
	return
}

func (s *sSettingCrontab) Update(ctx context.Context, in *req.SettingCrontabUpdate) (err error) {
	updateData := do.SettingCrontab{
		Name:      in.Name,
		Type:      in.Type,
		Rule:      in.Rule,
		Target:    in.Target,
		Parameter: in.Parameter,
		Remark:    in.Remark,
		Status:    in.Status,
		Singleton: in.Singleton,
	}
	_, err = s.Model(ctx).Data(updateData).Where("id", in.Id).Update()
	if utils.IsError(err) {
		return
	}
	return
}

func (s *sSettingCrontab) Delete(ctx context.Context, ids []int64) (err error) {
	_, err = s.Model(ctx).WhereIn("id", ids).Delete()
	if utils.IsError(err) {
		return err
	}
	return
}

func (s *sSettingCrontab) ChangeStatus(ctx context.Context, id int64, status int) (err error) {
	_, err = s.Model(ctx).Data(g.Map{"status": status}).Where("id", id).Update()
	if utils.IsError(err) {
		return err
	}
	return
}
