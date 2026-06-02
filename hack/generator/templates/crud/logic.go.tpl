// Package {{.PackageName}}
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package {{.PackageName}}

import (
	"context"
	"devinggo/internal/dao"
	"devinggo/internal/model/do"
	"devinggo/internal/model/entity"
	"devinggo/modules/{{.ModuleName}}/logic/base"
	"devinggo/modules/{{.ModuleName}}/model"
	"devinggo/modules/{{.ModuleName}}/model/req"
	"devinggo/modules/{{.ModuleName}}/model/res"
	"devinggo/modules/{{.ModuleName}}/pkg/handler"
	"devinggo/modules/{{.ModuleName}}/pkg/hook"
	"devinggo/modules/{{.ModuleName}}/pkg/orm"
	"devinggo/modules/{{.ModuleName}}/pkg/utils"
	"devinggo/modules/{{.ModuleName}}/service"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

type s{{.EntityName}} struct {
	base.BaseService
}

func init() {
	service.Register{{.EntityName}}(New{{.EntityName}}())
}

func New{{.EntityName}}() *s{{.EntityName}} {
	return &s{{.EntityName}}{}
}

func (s *s{{.EntityName}}) Model(ctx context.Context) *gdb.Model {
	return dao.{{.EntityName}}.Ctx(ctx).Hook(hook.Default()).Cache(orm.SetCacheOption(ctx)).OnConflict("id")
}

func (s *s{{.EntityName}}) GetPageListForSearch(ctx context.Context, req *model.PageListReq, in *req.{{.EntityName}}Search) (rs []*res.{{.EntityName}}, total int, err error) {
	m := s.handleSearch(ctx, in)
	var entity []*entity.{{.EntityName}}
	err = orm.NewQuery(m).WithPageListReq(req).ScanAndCount(&entity, &total)
	if utils.IsError(err) {
		return nil, 0, err
	}
	rs = make([]*res.{{.EntityName}}, 0)
	if !g.IsEmpty(entity) {
		if err = gconv.Structs(entity, &rs); err != nil {
			return nil, 0, err
		}
	}
	return
}

func (s *s{{.EntityName}}) GetList(ctx context.Context, in *req.{{.EntityName}}Search) (out []*res.{{.EntityName}}, err error) {
	inReq := &model.ListReq{
		OrderBy:   dao.{{.EntityName}}.Table() + ".created_at",
		OrderType: "desc",
	}
	m := s.handleSearch(ctx, in).Handler(handler.FilterAuth)
	m = orm.NewQuery(m).WithListReq(inReq).Build()
	err = m.Scan(&out)
	if utils.IsError(err) {
		return
	}
	return
}

func (s *s{{.EntityName}}) handleSearch(ctx context.Context, in *req.{{.EntityName}}Search) (m *gdb.Model) {
	m = s.Model(ctx){{.SearchConditions}}
	return
}

func (s *s{{.EntityName}}) Save(ctx context.Context, in *req.{{.EntityName}}Save) (id int64, err error) {
	saveData := do.{{.EntityName}}{
{{.SaveDoFields}}	}
	id, err = orm.NewQuery(s.Model(ctx)).Data(&saveData).InsertAndGetId()
	if utils.IsError(err) {
		return 0, err
	}
	return
}

func (s *s{{.EntityName}}) GetById(ctx context.Context, id int64) (out *res.{{.EntityName}}, err error) {
	var entity *entity.{{.EntityName}}
	err = s.Model(ctx).Where("id", id).Scan(&entity)
	if utils.IsError(err) {
		return nil, err
	}
	out = &res.{{.EntityName}}{}
	if err = gconv.Struct(entity, out); err != nil {
		return nil, err
	}
	return
}

func (s *s{{.EntityName}}) Update(ctx context.Context, in *req.{{.EntityName}}Update) (err error) {
	updateData := do.{{.EntityName}}{
{{.UpdateDoFields}}	}
	_, err = orm.NewQuery(s.Model(ctx)).Data(&updateData).Where("id", in.Id).Update()
	if utils.IsError(err) {
		return err
	}
	return
}

func (s *s{{.EntityName}}) Delete(ctx context.Context, ids []int64) (err error) {
	_, err = s.Model(ctx).WhereIn("id", ids).Delete()
	if utils.IsError(err) {
		return err
	}
	return
}

func (s *s{{.EntityName}}) RealDelete(ctx context.Context, ids []int64) (err error) {
	_, err = s.Model(ctx).Unscoped().WhereIn("id", ids).Delete()
	if utils.IsError(err) {
		return err
	}
	return
}

func (s *s{{.EntityName}}) Recovery(ctx context.Context, ids []int64) (err error) {
	_, err = s.Model(ctx).Unscoped().WhereIn("id", ids).Update(g.Map{
		"deleted_at": nil,
	})
	if utils.IsError(err) {
		return err
	}
	return
}

func (s *s{{.EntityName}}) ChangeStatus(ctx context.Context, id int64, status int) (err error) {
	_, err = s.Model(ctx).Where("id", id).Update(g.Map{
		"status": status,
	})
	if utils.IsError(err) {
		return err
	}
	return
}
