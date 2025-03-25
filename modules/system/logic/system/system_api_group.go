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

type sSystemApiGroup struct {
	base.BaseService
}

func init() {
	service.RegisterSystemApiGroup(NewSystemApiGroup())
}

func NewSystemApiGroup() *sSystemApiGroup {
	return &sSystemApiGroup{}
}

func (s *sSystemApiGroup) Model(ctx context.Context) *gdb.Model {
	return dao.SystemApiGroup.Ctx(ctx).Hook(hook.Bind()).Cache(orm.SetCacheOption(ctx)).OnConflict("id")
}

func (s *sSystemApiGroup) GetPageListForSearch(ctx context.Context, req *model.PageListReq, in *req.SystemApiGroupSearch) (rs []*res.SystemApiGroup, total int, err error) {
	m := s.handleSearch(ctx, in)
	var entity []*entity.SystemApiGroup
	err = orm.GetPageList(m, req).ScanAndCount(&entity, &total, false)
	if utils.IsError(err) {
		return nil, 0, err
	}
	rs = make([]*res.SystemApiGroup, 0)
	if !g.IsEmpty(entity) {
		if err = gconv.Structs(entity, &rs); err != nil {
			return nil, 0, err
		}
	}

	if !g.IsEmpty(in.GetApiList) && in.GetApiList {
		if !g.IsEmpty(rs) {
			for _, v := range rs {
				var apis []*res.SystemApi
				err = service.SystemApi().Model(ctx).
					Where(dao.SystemApi.Columns().GroupId, v.Id).
					Where(dao.SystemApi.Columns().Status, 1).
					Scan(&apis)
				if err != nil {
					continue
				}
				v.Apis = apis
			}
		}
	}

	return
}

func (s *sSystemApiGroup) GetList(ctx context.Context, in *req.SystemApiGroupSearch) (out []*res.SystemApiGroup, err error) {
	inReq := &model.ListReq{
		OrderBy:   dao.SystemApiGroup.Table() + ".created_by",
		OrderType: "desc",
	}
	m := s.handleSearch(ctx, in).Handler(handler.FilterAuth)
	m = orm.GetList(m, inReq)
	err = m.Scan(&out)
	if utils.IsError(err) {
		return
	}
	if !g.IsEmpty(in.GetApiList) && in.GetApiList {
		if !g.IsEmpty(out) {
			for _, v := range out {
				var apis []*res.SystemApi
				err = service.SystemApi().Model(ctx).
					Where(dao.SystemApi.Columns().GroupId, v.Id).
					Where(dao.SystemApi.Columns().Status, 1).
					Scan(&apis)
				if err != nil {
					continue
				}
				v.Apis = apis
			}
		}
	}
	return
}

func (s *sSystemApiGroup) handleSearch(ctx context.Context, in *req.SystemApiGroupSearch) (m *gdb.Model) {
	m = s.Model(ctx)
	if !g.IsEmpty(in.Name) {
		m = m.Where("name", in.Name)
	}

	if !g.IsEmpty(in.Status) {
		m = m.Where("status", in.Status)
	}
	return
}

func (s *sSystemApiGroup) Save(ctx context.Context, in *req.SystemApiGroupSave) (id int64, err error) {
	saveData := do.SystemApiGroup{
		Name:   in.Name,
		Status: in.Status,
		Remark: in.Remark,
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

func (s *sSystemApiGroup) GetById(ctx context.Context, id int64) (res *res.SystemApiGroup, err error) {
	err = s.Model(ctx).Where("id", id).Scan(&res)
	if utils.IsError(err) {
		return
	}
	return
}

func (s *sSystemApiGroup) Update(ctx context.Context, in *req.SystemApiGroupUpdate) (err error) {
	updateData := do.SystemApiGroup{
		Name:   in.Name,
		Status: in.Status,
		Remark: in.Remark,
	}
	_, err = s.Model(ctx).Data(updateData).Where("id", in.Id).Update()
	if utils.IsError(err) {
		return
	}
	return
}

func (s *sSystemApiGroup) Delete(ctx context.Context, ids []int64) (err error) {
	_, err = s.Model(ctx).WhereIn("id", ids).Delete()
	if utils.IsError(err) {
		return err
	}
	return
}

func (s *sSystemApiGroup) RealDelete(ctx context.Context, ids []int64) (err error) {
	var res []*res.SystemApiGroup
	err = s.Model(ctx).Unscoped().WhereIn("id", ids).Scan(&res)
	if utils.IsError(err) {
		return
	}
	return
}

func (s *sSystemApiGroup) Recovery(ctx context.Context, ids []int64) (err error) {
	_, err = s.Model(ctx).Unscoped().WhereIn("id", ids).Update(g.Map{"deleted_at": nil})
	if utils.IsError(err) {
		return err
	}
	return
}

func (s *sSystemApiGroup) ChangeStatus(ctx context.Context, id int64, status int) (err error) {
	_, err = s.Model(ctx).Data(g.Map{"status": status}).Where("id", id).Update()
	if utils.IsError(err) {
		return err
	}
	return
}
