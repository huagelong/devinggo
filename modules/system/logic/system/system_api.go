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

type sSystemApi struct {
	base.BaseService
}

func init() {
	service.RegisterSystemApi(NewSystemApi())
}

func NewSystemApi() *sSystemApi {
	return &sSystemApi{}
}

func (s *sSystemApi) Model(ctx context.Context) *gdb.Model {
	return dao.SystemApi.Ctx(ctx).Hook(hook.Bind()).Cache(orm.SetCacheOption(ctx)).OnConflict("id")
}

func (s *sSystemApi) GetPageListForSearch(ctx context.Context, req *model.PageListReq, in *req.SystemApiSearch) (rs []*res.SystemApi, total int, err error) {
	m := s.handleSearch(ctx, in)
	var entity []*entity.SystemApi
	err = orm.GetPageList(m, req).ScanAndCount(&entity, &total, false)
	if utils.IsError(err) {
		return nil, 0, err
	}
	rs = make([]*res.SystemApi, 0)
	if !g.IsEmpty(entity) {
		if err = gconv.Structs(entity, &rs); err != nil {
			return nil, 0, err
		}
	}
	return
}

func (s *sSystemApi) GetList(ctx context.Context, in *req.SystemApiSearch) (out []*res.SystemApi, err error) {
	inReq := &model.ListReq{
		OrderBy:   dao.SystemApi.Table() + ".created_by",
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

func (s *sSystemApi) handleSearch(ctx context.Context, in *req.SystemApiSearch) (m *gdb.Model) {
	m = s.Model(ctx)
	if !g.IsEmpty(in.GroupId) {
		m = m.Where("group_id", in.GroupId)
	}

	if !g.IsEmpty(in.Name) {
		m = m.Where("name", in.Name)
	}

	if !g.IsEmpty(in.AccessName) {
		m = m.Where("access_name", in.AccessName)
	}

	if !g.IsEmpty(in.Status) {
		m = m.Where("status", in.Status)
	}
	return
}

func (s *sSystemApi) Save(ctx context.Context, in *req.SystemApiSave) (id int64, err error) {
	saveData := do.SystemApi{
		GroupId:     in.GroupId,
		Name:        in.Name,
		AccessName:  in.AccessName,
		AuthMode:    in.AuthMode,
		RequestMode: in.RequestMode,
		Status:      in.Status,
		Remark:      in.Remark,
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

func (s *sSystemApi) GetById(ctx context.Context, id int64) (res *res.SystemApi, err error) {
	err = s.Model(ctx).Where("id", id).Scan(&res)
	if utils.IsError(err) {
		return
	}
	return
}

func (s *sSystemApi) Update(ctx context.Context, in *req.SystemApiUpdate) (err error) {
	updateData := do.SystemApi{
		GroupId:     in.GroupId,
		Name:        in.Name,
		AccessName:  in.AccessName,
		AuthMode:    in.AuthMode,
		RequestMode: in.RequestMode,
		Status:      in.Status,
		Remark:      in.Remark,
	}
	_, err = s.Model(ctx).Data(updateData).Where("id", in.Id).Update()
	if utils.IsError(err) {
		return
	}
	return
}

func (s *sSystemApi) Delete(ctx context.Context, ids []int64) (err error) {
	_, err = s.Model(ctx).WhereIn("id", ids).Delete()
	if utils.IsError(err) {
		return err
	}
	return
}

func (s *sSystemApi) RealDelete(ctx context.Context, ids []int64) (err error) {
	var res []*res.SystemApi
	err = s.Model(ctx).Unscoped().WhereIn("id", ids).Scan(&res)
	if utils.IsError(err) {
		return
	}
	return
}

func (s *sSystemApi) Recovery(ctx context.Context, ids []int64) (err error) {
	_, err = s.Model(ctx).Unscoped().WhereIn("id", ids).Update(g.Map{"deleted_at": nil})
	if utils.IsError(err) {
		return err
	}
	return
}

func (s *sSystemApi) ChangeStatus(ctx context.Context, id int64, status int) (err error) {
	_, err = s.Model(ctx).Data(g.Map{"status": status}).Where("id", id).Update()
	if utils.IsError(err) {
		return err
	}
	return
}
