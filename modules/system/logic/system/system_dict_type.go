// Package system
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package system

import (
	"context"
	"dario.cat/mergo"
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

type sSystemDictType struct {
	base.BaseService
}

func init() {
	service.RegisterSystemDictType(NewSystemDictType())
}

func NewSystemDictType() *sSystemDictType {
	return &sSystemDictType{}
}

func (s *sSystemDictType) Model(ctx context.Context) *gdb.Model {
	return dao.SystemDictType.Ctx(ctx).Hook(hook.Bind()).Cache(orm.SetCacheOption(ctx)).OnConflict("id")
}

func (s *sSystemDictType) GetPageList(ctx context.Context, req *model.PageListReq, in *req.SystemDictTypeSearch) (rs []*res.SystemDictType, total int, err error) {
	m := s.handleSearch(ctx, in).Handler(handler.FilterAuth)
	var entity []*entity.SystemDictType
	err = orm.GetPageList(m, req).ScanAndCount(&entity, &total, false)
	if utils.IsError(err) {
		return nil, 0, err
	}
	rs = make([]*res.SystemDictType, 0)
	if !g.IsEmpty(entity) {
		if err = gconv.Structs(entity, &rs); err != nil {
			return nil, 0, err
		}
	}
	return
}

func (s *sSystemDictType) GetList(ctx context.Context, listReq *model.ListReq, in *req.SystemDictTypeSearch) (out []*res.SystemDictType, err error) {
	inReq := &model.ListReq{
		OrderBy:   "sort",
		OrderType: "desc",
	}
	mergo.Merge(&listReq, inReq)
	m := s.handleSearch(ctx, in)
	err = orm.GetList(m, listReq).Scan(&out)
	if utils.IsError(err) {
		return nil, err
	}
	return
}

func (s *sSystemDictType) Save(ctx context.Context, in *req.SystemDictTypeSave) (id int64, err error) {
	saveData := do.SystemDictType{
		Name:   in.Name,
		Status: in.Status,
		Code:   in.Code,
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

func (s *sSystemDictType) GetById(ctx context.Context, id int64) (res *res.SystemDictType, err error) {
	err = s.Model(ctx).Where("id", id).Scan(&res)
	if utils.IsError(err) {
		return
	}
	return
}

func (s *sSystemDictType) Update(ctx context.Context, in *req.SystemDictTypeUpdate) (err error) {
	updateData := do.SystemPost{
		Name:   in.Name,
		Status: in.Status,
		Code:   in.Code,
		Remark: in.Remark,
	}
	_, err = s.Model(ctx).Data(updateData).Where("id", in.Id).Update()
	if utils.IsError(err) {
		return
	}

	_, err = service.SystemDictData().Model(ctx).Where("type_id", in.Id).Update(g.Map{"code": in.Code})
	if utils.IsError(err) {
		return
	}
	return
}

func (s *sSystemDictType) Delete(ctx context.Context, ids []int64) (err error) {
	_, err = s.Model(ctx).WhereIn("id", ids).Delete()
	if utils.IsError(err) {
		return err
	}
	_, err = service.SystemDictData().Model(ctx).Where("type_id in (?)", ids).Delete()
	if utils.IsError(err) {
		return
	}
	return
}

func (s *sSystemDictType) RealDelete(ctx context.Context, ids []int64) (err error) {
	_, err = s.Model(ctx).Unscoped().WhereIn("id", ids).Delete()
	if utils.IsError(err) {
		return
	}
	_, err = service.SystemDictData().Model(ctx).Unscoped().Where("type_id in (?)", ids).Delete()
	if utils.IsError(err) {
		return
	}
	return
}

func (s *sSystemDictType) Recovery(ctx context.Context, ids []int64) (err error) {
	_, err = s.Model(ctx).Unscoped().WhereIn("id", ids).Update(g.Map{"deleted_at": nil})
	if utils.IsError(err) {
		return err
	}
	return
}

func (s *sSystemDictType) ChangeStatus(ctx context.Context, id int64, status int) (err error) {
	_, err = s.Model(ctx).Data(g.Map{"status": status}).Where("id", id).Update()
	if utils.IsError(err) {
		return err
	}
	return
}

func (s *sSystemDictType) handleSearch(ctx context.Context, in *req.SystemDictTypeSearch) (m *gdb.Model) {
	m = s.Model(ctx)

	if !g.IsEmpty(in.Status) {
		m = m.Where("status", in.Status)
	}
	if !g.IsEmpty(in.Code) {
		m = m.Where("code like ? ", "%"+in.Code+"%")
	}
	if !g.IsEmpty(in.Name) {
		m = m.Where("name like ? ", "%"+in.Name+"%")
	}
	return
}
