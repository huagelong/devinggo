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

type sSystemAppGroup struct {
	base.BaseService
}

func init() {
	service.RegisterSystemAppGroup(NewSystemAppGroup())
}

func NewSystemAppGroup() *sSystemAppGroup {
	return &sSystemAppGroup{}
}

func (s *sSystemAppGroup) Model(ctx context.Context) *gdb.Model {
	return dao.SystemAppGroup.Ctx(ctx).Hook(hook.Bind()).Cache(orm.SetCacheOption(ctx)).OnConflict("id")
}

func (s *sSystemAppGroup) GetPageListForSearch(ctx context.Context, req *model.PageListReq, in *req.SystemAppGroupSearch) (rs []*res.SystemAppGroup, total int, err error) {
	m := s.handleSearch(ctx, in)
	var entity []*entity.SystemAppGroup
	err = orm.GetPageList(m, req).ScanAndCount(&entity, &total, false)
	if utils.IsError(err) {
		return nil, 0, err
	}
	rs = make([]*res.SystemAppGroup, 0)
	if !g.IsEmpty(entity) {
		if err = gconv.Structs(entity, &rs); err != nil {
			return nil, 0, err
		}
	}
	return
}

func (s *sSystemAppGroup) GetList(ctx context.Context, in *req.SystemAppGroupSearch) (out []*res.SystemAppGroup, err error) {
	inReq := &model.ListReq{
		OrderBy:   "created_at",
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

func (s *sSystemAppGroup) handleSearch(ctx context.Context, in *req.SystemAppGroupSearch) (m *gdb.Model) {
	m = s.Model(ctx)
	if !g.IsEmpty(in.Name) {
		m = m.Where("name", in.Name)
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

func (s *sSystemAppGroup) Save(ctx context.Context, in *req.SystemAppGroupSave) (id int64, err error) {
	saveData := do.SystemAppGroup{
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

func (s *sSystemAppGroup) GetById(ctx context.Context, id int64) (res *res.SystemAppGroup, err error) {
	err = s.Model(ctx).Where("id", id).Scan(&res)
	if utils.IsError(err) {
		return
	}
	return
}

func (s *sSystemAppGroup) Update(ctx context.Context, in *req.SystemAppGroupUpdate) (err error) {
	updateData := do.SystemAppGroup{
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

func (s *sSystemAppGroup) Delete(ctx context.Context, ids []int64) (err error) {
	_, err = s.Model(ctx).WhereIn("id", ids).Delete()
	if utils.IsError(err) {
		return err
	}
	return
}

func (s *sSystemAppGroup) RealDelete(ctx context.Context, ids []int64) (err error) {
	var res []*res.SystemAppGroup
	err = s.Model(ctx).Unscoped().WhereIn("id", ids).Scan(&res)
	if utils.IsError(err) {
		return
	}
	return
}

func (s *sSystemAppGroup) Recovery(ctx context.Context, ids []int64) (err error) {
	_, err = s.Model(ctx).Unscoped().WhereIn("id", ids).Update(g.Map{"deleted_at": nil})
	if utils.IsError(err) {
		return err
	}
	return
}

func (s *sSystemAppGroup) ChangeStatus(ctx context.Context, id int64, status int) (err error) {
	_, err = s.Model(ctx).Data(g.Map{"status": status}).Where("id", id).Update()
	if utils.IsError(err) {
		return err
	}
	return
}
