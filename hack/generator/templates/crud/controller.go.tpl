// Package {{.PackageName}}
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package {{.PackageName}}

import (
	"context"
	"devinggo/internal/dao"
	"devinggo/modules/{{.ModuleName}}/api/{{.PackageName}}"
	"devinggo/modules/{{.ModuleName}}/controller/base"
	"devinggo/modules/{{.ModuleName}}/model/res"
	"devinggo/modules/{{.ModuleName}}/service"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

var (
	{{.EntityName}}Controller = {{.VarName}}Controller{}
)

type {{.VarName}}Controller struct {
	base.BaseController
}

func (c *{{.VarName}}Controller) Index(ctx context.Context, in *{{.PackageName}}.Index{{.EntityName}}Req) (out *{{.PackageName}}.Index{{.EntityName}}Res, err error) {
	out = &{{.PackageName}}.Index{{.EntityName}}Res{}
	items, totalCount, err := service.{{.EntityName}}().GetPageListForSearch(ctx, &in.PageListReq, &in.{{.EntityName}}Search)
	if err != nil {
		return
	}

	if !g.IsEmpty(items) {
		for _, item := range items {
			out.Items = append(out.Items, *item)
		}
	} else {
		out.Items = make([]res.{{.EntityName}}, 0)
	}
	out.PageRes.Pack(in, totalCount)
	return
}

func (c *{{.VarName}}Controller) List(ctx context.Context, in *{{.PackageName}}.List{{.EntityName}}Req) (out *{{.PackageName}}.List{{.EntityName}}Res, err error) {
	out = &{{.PackageName}}.List{{.EntityName}}Res{}
	items, err := service.{{.EntityName}}().GetList(ctx, &{{.PackageName}}.{{.EntityName}}Search{})
	if err != nil {
		return
	}

	if !g.IsEmpty(items) {
		for _, item := range items {
			out.Data = append(out.Data, *item)
		}
	} else {
		out.Data = make([]res.{{.EntityName}}, 0)
	}
	return
}

func (c *{{.VarName}}Controller) Recycle(ctx context.Context, in *{{.PackageName}}.Recycle{{.EntityName}}Req) (out *{{.PackageName}}.Recycle{{.EntityName}}Res, err error) {
	out = &{{.PackageName}}.Recycle{{.EntityName}}Res{}
	in.Recycle = true
	items, totalCount, err := service.{{.EntityName}}().GetPageListForSearch(ctx, &in.PageListReq, &in.{{.EntityName}}Search)
	if err != nil {
		return
	}

	if !g.IsEmpty(items) {
		for _, item := range items {
			out.Items = append(out.Items, *item)
		}
	} else {
		out.Items = make([]res.{{.EntityName}}, 0)
	}
	out.PageRes.Pack(in, totalCount)
	return
}

func (c *{{.VarName}}Controller) Save(ctx context.Context, in *{{.PackageName}}.Save{{.EntityName}}Req) (out *{{.PackageName}}.Save{{.EntityName}}Res, err error) {
	out = &{{.PackageName}}.Save{{.EntityName}}Res{}
	id, err := service.{{.EntityName}}().Save(ctx, &in.{{.EntityName}}Save)
	if err != nil {
		return
	}
	out.Id = id
	return
}

func (c *{{.VarName}}Controller) Read(ctx context.Context, in *{{.PackageName}}.Read{{.EntityName}}Req) (out *{{.PackageName}}.Read{{.EntityName}}Res, err error) {
	out = &{{.PackageName}}.Read{{.EntityName}}Res{}
	info, err := service.{{.EntityName}}().GetById(ctx, in.Id)
	if err != nil {
		return
	}
	out.Data = *info
	return
}

func (c *{{.VarName}}Controller) Update(ctx context.Context, in *{{.PackageName}}.Update{{.EntityName}}Req) (out *{{.PackageName}}.Update{{.EntityName}}Res, err error) {
	err = dao.{{.EntityName}}.Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		out = &{{.PackageName}}.Update{{.EntityName}}Res{}
		err = service.{{.EntityName}}().Update(ctx, &in.{{.EntityName}}Update)
		if err != nil {
			return
		}
		return
	})
	if err != nil {
		return
	}
	return
}

func (c *{{.VarName}}Controller) Delete(ctx context.Context, in *{{.PackageName}}.Delete{{.EntityName}}Req) (out *{{.PackageName}}.Delete{{.EntityName}}Res, err error) {
	out = &{{.PackageName}}.Delete{{.EntityName}}Res{}
	err = service.{{.EntityName}}().Delete(ctx, in.Ids)
	if err != nil {
		return
	}
	return
}

func (c *{{.VarName}}Controller) RealDelete(ctx context.Context, in *{{.PackageName}}.RealDelete{{.EntityName}}Req) (out *{{.PackageName}}.RealDelete{{.EntityName}}Res, err error) {
	err = dao.{{.EntityName}}.Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		out = &{{.PackageName}}.RealDelete{{.EntityName}}Res{}
		err = service.{{.EntityName}}().RealDelete(ctx, in.Ids)
		if err != nil {
			return
		}
		return
	})
	if err != nil {
		return
	}
	return
}

func (c *{{.VarName}}Controller) Recovery(ctx context.Context, in *{{.PackageName}}.Recovery{{.EntityName}}Req) (out *{{.PackageName}}.Recovery{{.EntityName}}Res, err error) {
	err = dao.{{.EntityName}}.Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		out = &{{.PackageName}}.Recovery{{.EntityName}}Res{}
		err = service.{{.EntityName}}().Recovery(ctx, in.Ids)
		if err != nil {
			return
		}
		return
	})
	if err != nil {
		return
	}
	return
}

func (c *{{.VarName}}Controller) ChangeStatus(ctx context.Context, in *{{.PackageName}}.ChangeStatus{{.EntityName}}Req) (out *{{.PackageName}}.ChangeStatus{{.EntityName}}Res, err error) {
	out = &{{.PackageName}}.ChangeStatus{{.EntityName}}Res{}
	err = service.{{.EntityName}}().ChangeStatus(ctx, in.Id, in.Status)
	if err != nil {
		return
	}
	return
}
