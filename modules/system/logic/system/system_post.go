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

type sSystemPost struct {
	base.GenericService[res.SystemPost]
}

func init() {
	service.RegisterSystemPost(NewSystemPost())
}

func NewSystemPost() *sSystemPost {
	s := &sSystemPost{}
	s.GenericService = base.GenericService[res.SystemPost]{
		ModelFn: func(ctx context.Context) *gdb.Model {
			return dao.SystemPost.Ctx(ctx).Hook(hook.Default()).Cache(orm.SetCacheOption(ctx)).OnConflict("id")
		},
	}
	return s
}

// Model 返回数据库 Model
func (s *sSystemPost) Model(ctx context.Context) *gdb.Model {
	return s.GenericService.Model(ctx)
}

func (s *sSystemPost) handlePostSearch(ctx context.Context, in *req.SystemPostSearch) (m *gdb.Model) {
	m = s.Model(ctx)
	if !g.IsEmpty(in.Code) {
		m = m.Where("code", in.Code)
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

func (s *sSystemPost) GetList(ctx context.Context, in *req.SystemPostSearch) (out []*res.SystemPost, err error) {
	inReq := &model.ListReq{
		OrderBy:   "sort",
		OrderType: "desc",
	}
	m := s.handlePostSearch(ctx, in).Handler(handler.FilterAuth)
	m = orm.NewQuery(m).WithListReq(inReq).Build()
	err = m.Scan(&out)
	if utils.IsError(err) {
		return
	}
	return
}

func (s *sSystemPost) GetPageList(ctx context.Context, req *model.PageListReq, in *req.SystemPostSearch) (rs []*res.SystemPost, total int, err error) {
	m := s.handlePostSearch(ctx, in).Handler(handler.FilterAuth)
	var postEntity []*entity.SystemPost
	err = orm.NewQuery(m).WithPageListReq(req).ScanAndCount(&postEntity, &total)
	if utils.IsError(err) {
		return nil, 0, err
	}
	rs = make([]*res.SystemPost, 0)
	if !g.IsEmpty(postEntity) {
		if err = gconv.Structs(postEntity, &rs); utils.IsError(err) {
			return nil, 0, err
		}
	}
	return
}

func (s *sSystemPost) Save(ctx context.Context, in *req.SystemPostSave) (id int64, err error) {
	saveData := do.SystemPost{
		Name:   in.Name,
		Sort:   in.Sort,
		Status: in.Status,
		Code:   in.Code,
		Remark: in.Remark,
	}
	rs, err := s.Model(ctx).Data(saveData).Insert()
	if utils.IsError(err) {
		return
	}
	tmpId, err := rs.LastInsertId()
	if utils.IsError(err) {
		return
	}
	id = gconv.Int64(tmpId)
	return
}

func (s *sSystemPost) Update(ctx context.Context, in *req.SystemPostSave) (err error) {
	updateData := do.SystemPost{
		Name:   in.Name,
		Sort:   in.Sort,
		Status: in.Status,
		Code:   in.Code,
		Remark: in.Remark,
	}
	_, err = s.Model(ctx).Data(updateData).Where("id", in.Id).Update()
	if utils.IsError(err) {
		return
	}
	return
}

// GetById 由 GenericService 提供，此处声明用于接口生成
func (s *sSystemPost) GetById(ctx context.Context, id int64) (res *res.SystemPost, err error) {
	return s.GenericService.GetById(ctx, id)
}

// ChangeStatus 由 GenericService 提供，此处声明用于接口生成
func (s *sSystemPost) ChangeStatus(ctx context.Context, id int64, status int) (err error) {
	return s.GenericService.ChangeStatus(ctx, id, status)
}

// Recovery 由 GenericService 提供，此处声明用于接口生成
func (s *sSystemPost) Recovery(ctx context.Context, ids []int64) (err error) {
	return s.GenericService.Recovery(ctx, ids)
}

// Delete 由 GenericService 提供，此处声明用于接口生成
func (s *sSystemPost) Delete(ctx context.Context, ids []int64) (err error) {
	return s.GenericService.Delete(ctx, ids)
}

// RealDelete 由 GenericService 提供，此处声明用于接口生成
func (s *sSystemPost) RealDelete(ctx context.Context, ids []int64) (err error) {
	return s.GenericService.RealDelete(ctx, ids)
}

// NumberOperation 由 GenericService 提供，此处声明用于接口生成
func (s *sSystemPost) NumberOperation(ctx context.Context, id int64, numberName string, numberValue int) (err error) {
	return s.GenericService.NumberOperation(ctx, id, numberName, numberValue)
}
