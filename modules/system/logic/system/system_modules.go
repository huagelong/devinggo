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
	"devinggo/modules/system/pkg/hook"
	"devinggo/modules/system/pkg/orm"
	"devinggo/modules/system/pkg/utils"
	"devinggo/modules/system/service"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

type sSystemModules struct {
	base.BaseService
}

func init() {
	service.RegisterSystemModules(NewSystemModules())
}

func NewSystemModules() *sSystemModules {
	return &sSystemModules{}
}

func (s *sSystemModules) Model(ctx context.Context) *gdb.Model {
	return dao.SystemModules.Ctx(ctx).Hook(hook.Bind()).Cache(orm.SetCacheOption(ctx)).OnConflict("id")
}

func (s *sSystemModules) handleSearch(ctx context.Context, in *req.SystemModulesSearch) (m *gdb.Model) {
	m = s.Model(ctx)

	if !g.IsEmpty(in.Id) {
		m = m.Where("id", in.Id)
	}

	if !g.IsEmpty(in.Name) {
		m = m.Where("name like ? ", "%"+in.Name+"%")
	}

	if !g.IsEmpty(in.Label) {
		m = m.Where("label", in.Label)
	}

	if !g.IsEmpty(in.Installed) {
		m = m.Where("installed", in.Installed)
	}

	if !g.IsEmpty(in.Status) {
		m = m.Where("status", in.Status)
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

func (s *sSystemModules) GetList(ctx context.Context, inReq *model.ListReq, in *req.SystemModulesSearch) (out []*res.SystemModules, err error) {
	m := s.handleSearch(ctx, in).Handler(handler.FilterAuth)
	m = orm.GetList(m, inReq)
	err = m.Scan(&out)
	if utils.IsError(err) {
		return
	}
	return
}

func (s *sSystemModules) GetPageList(ctx context.Context, req *model.PageListReq, in *req.SystemModulesSearch) (rs []*res.SystemModules, total int, err error) {
	m := s.handleSearch(ctx, in).Handler(handler.FilterAuth)
	var entity []*entity.SystemModules
	err = orm.GetPageList(m, req).ScanAndCount(&entity, &total, false)
	if utils.IsError(err) {
		return nil, 0, err
	}
	rs = make([]*res.SystemModules, 0)
	if !g.IsEmpty(entity) {
		if err = gconv.Structs(entity, &rs); err != nil {
			return nil, 0, err
		}
	}
	return
}

func (s *sSystemModules) Save(ctx context.Context, in *req.SystemModulesSave) (id int64, err error) {
	var saveData *do.SystemModules
	if err = gconv.Struct(in, &saveData); err != nil {
		return
	}
	rs, err := s.Model(ctx).OmitEmptyData().Data(saveData).Insert()
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

func (s *sSystemModules) GetById(ctx context.Context, id int64) (res *res.SystemModules, err error) {
	err = s.Model(ctx).Where("id", id).Scan(&res)
	if utils.IsError(err) {
		return
	}
	return
}

func (s *sSystemModules) Update(ctx context.Context, in *req.SystemModulesUpdate) (err error) {
	var updateData *do.SystemModules
	if err = gconv.Struct(in, &updateData); err != nil {
		return
	}
	_, err = s.Model(ctx).OmitEmptyData().Data(updateData).Where("id", in.Id).Update()
	if utils.IsError(err) {
		return
	}
	return
}

func (s *sSystemModules) Delete(ctx context.Context, ids []int64) (err error) {
	_, err = s.Model(ctx).WhereIn("id", ids).Delete()
	if utils.IsError(err) {
		return err
	}
	return
}

func (s *sSystemModules) RealDelete(ctx context.Context, ids []int64) (err error) {
	_, err = s.Model(ctx).Unscoped().WhereIn("id", ids).Delete()
	if utils.IsError(err) {
		return
	}
	return
}

func (s *sSystemModules) Recovery(ctx context.Context, ids []int64) (err error) {
	_, err = s.Model(ctx).Unscoped().WhereIn("id", ids).Update(g.Map{"deleted_at": nil})
	if utils.IsError(err) {
		return err
	}
	return
}

func (s *sSystemModules) ChangeStatus(ctx context.Context, id int64, status int) (err error) {
	_, err = s.Model(ctx).Data(g.Map{"status": status}).Where("id", id).Update()
	if utils.IsError(err) {
		return err
	}
	return
}
