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

type sSystemDictData struct {
	base.BaseService
}

func init() {
	service.RegisterSystemDictData(NewSystemDictData())
}

func NewSystemDictData() *sSystemDictData {
	return &sSystemDictData{}
}

func (s *sSystemDictData) Model(ctx context.Context) *gdb.Model {
	return dao.SystemDictData.Ctx(ctx).Hook(hook.Bind()).Cache(orm.SetCacheOption(ctx)).OnConflict("id")
}

func (s *sSystemDictData) GetList(ctx context.Context, listReq *model.ListReq, in *req.SystemDictDataSearch) (out []*res.SystemDictData, err error) {
	inReq := &model.ListReq{
		OrderBy:   "sort",
		OrderType: "desc",
	}

	mergo.Merge(&listReq, inReq)
	dbType := utils.GetDbType()
	if dbType == "mysql" {
		listReq.Select = "id, `label` as `title`, `value` as `key`,code"
	} else {
		listReq.Select = "id, label as title, value as key,code"
	}
	m := s.handleSearch(ctx, in)
	err = orm.GetList(m, listReq).Scan(&out)
	if utils.IsError(err) {
		return nil, err
	}
	return
}

func (s *sSystemDictData) GetPageList(ctx context.Context, req *model.PageListReq, in *req.SystemDictDataSearch) (rs []*res.SystemDictDataFull, total int, err error) {
	m := s.handleSearch(ctx, in).Handler(handler.FilterAuth)
	var entity []*entity.SystemDictData
	err = orm.GetPageList(m, req).ScanAndCount(&entity, &total, false)
	if utils.IsError(err) {
		return nil, 0, err
	}
	rs = make([]*res.SystemDictDataFull, 0)
	if !g.IsEmpty(entity) {
		if err = gconv.Structs(entity, &rs); err != nil {
			return nil, 0, err
		}
	}
	return
}

func (s *sSystemDictData) Save(ctx context.Context, in *req.SystemDictDataSave) (id int64, err error) {
	saveData := do.SystemDictData{
		TypeId: in.TypeId,
		Value:  in.Value,
		Code:   in.Code,
		Label:  in.Label,
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

func (s *sSystemDictData) GetById(ctx context.Context, id int64) (res *res.SystemDictDataFull, err error) {
	err = s.Model(ctx).Where("id", id).Scan(&res)
	if utils.IsError(err) {
		return
	}
	return
}

func (s *sSystemDictData) Update(ctx context.Context, in *req.SystemDictDataUpdate) (err error) {
	updateData := do.SystemDictData{
		Value: in.Value,
		Code:  in.Code,
		Label: in.Label,
	}
	_, err = s.Model(ctx).Data(updateData).Where("id", in.Id).Update()
	if utils.IsError(err) {
		return
	}
	return
}

func (s *sSystemDictData) Delete(ctx context.Context, ids []int64) (err error) {
	_, err = s.Model(ctx).WhereIn("id", ids).Delete()
	if utils.IsError(err) {
		return err
	}
	return
}

func (s *sSystemDictData) RealDelete(ctx context.Context, ids []int64) (err error) {
	_, err = s.Model(ctx).Unscoped().WhereIn("id", ids).Delete()
	if utils.IsError(err) {
		return
	}
	return
}

func (s *sSystemDictData) Recovery(ctx context.Context, ids []int64) (err error) {
	_, err = s.Model(ctx).Unscoped().WhereIn("id", ids).Update(g.Map{"deleted_at": nil})
	if utils.IsError(err) {
		return err
	}
	return
}

func (s *sSystemDictData) ChangeStatus(ctx context.Context, id int64, status int) (err error) {
	_, err = s.Model(ctx).Data(g.Map{"status": status}).Where("id", id).Update()
	if utils.IsError(err) {
		return err
	}
	return
}

func (s *sSystemDictData) NumberOperation(ctx context.Context, id int64, numberName string, numberValue int) (err error) {
	_, err = s.Model(ctx).Data(g.Map{numberName: numberValue}).Where("id", id).Update()
	if utils.IsError(err) {
		return err
	}
	return
}

func (s *sSystemDictData) handleSearch(ctx context.Context, in *req.SystemDictDataSearch) (m *gdb.Model) {

	m = s.Model(ctx)

	if !g.IsEmpty(in.Code) {
		m = m.Where("code", in.Code)
	}

	if !g.IsEmpty(in.TypeId) {
		m = m.Where("type_id", in.TypeId)
	}

	if !g.IsEmpty(in.Status) {
		m = m.Where("status", in.Status)
	}

	if !g.IsEmpty(in.Value) {
		m = m.Where("value like ? ", "%"+in.Value+"%")
	}

	if !g.IsEmpty(in.Label) {
		m = m.Where("label like ? ", "%"+in.Label+"%")
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
