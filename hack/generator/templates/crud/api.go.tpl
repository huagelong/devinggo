// Package {{.PackageName}}
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package {{.PackageName}}

import (
	"devinggo/modules/{{.ModuleName}}/model"
	"devinggo/modules/{{.ModuleName}}/model/page"
	"devinggo/modules/{{.ModuleName}}/model/req"
	"devinggo/modules/{{.ModuleName}}/model/res"

	"github.com/gogf/gf/v2/frame/g"
)

type Index{{.EntityName}}Req struct {
	g.Meta `path:"/{{.VarName}}/index" method:"get" tags:"{{.ChineseName}}" summary:"{{.ChineseName}}列表." x-permission:"{{.ModuleName}}:{{.VarName}}:index"`
	model.AuthorHeader
	model.PageListReq
	req.{{.EntityName}}Search
}

type Index{{.EntityName}}Res struct {
	g.Meta `mime:"application/json"`
	page.PageRes
	Items []res.{{.EntityName}} `json:"items"  dc:"{{.ChineseName}} list"`
}

type List{{.EntityName}}Req struct {
	g.Meta `path:"/{{.VarName}}/list" method:"get" tags:"{{.ChineseName}}" summary:"列表，无分页.." x-permission:"{{.ModuleName}}:{{.VarName}}:list"`
	model.AuthorHeader
}

type List{{.EntityName}}Res struct {
	g.Meta `mime:"application/json"`
	Data   []res.{{.EntityName}} `json:"data"  dc:"list"`
}

type Recycle{{.EntityName}}Req struct {
	g.Meta `path:"/{{.VarName}}/recycle" method:"get" tags:"{{.ChineseName}}" summary:"回收站{{.ChineseName}}列表." x-permission:"{{.ModuleName}}:{{.VarName}}:recycle"`
	model.AuthorHeader
	model.PageListReq
	req.{{.EntityName}}Search
}

type Recycle{{.EntityName}}Res struct {
	g.Meta `mime:"application/json"`
	page.PageRes
	Items []res.{{.EntityName}} `json:"items"  dc:"{{.ChineseName}} list"`
}

type Save{{.EntityName}}Req struct {
	g.Meta `path:"/{{.VarName}}/save" method:"post" tags:"{{.ChineseName}}" summary:"新增{{.ChineseName}}." x-permission:"{{.ModuleName}}:{{.VarName}}:save"`
	model.AuthorHeader
	req.{{.EntityName}}Save
}

type Save{{.EntityName}}Res struct {
	g.Meta `mime:"application/json"`
	Id     int64 `json:"id" dc:"{{.ChineseName}} id"`
}

type Read{{.EntityName}}Req struct {
	g.Meta `path:"/{{.VarName}}/read/{Id}" method:"get" tags:"{{.ChineseName}}" summary:"读取{{.ChineseName}}." x-permission:"{{.ModuleName}}:{{.VarName}}:read"`
	model.AuthorHeader
	Id int64 `json:"id" dc:"{{.ChineseName}} id" v:"required|min:1#{{.ChineseName}}Id不能为空"`
}

type Read{{.EntityName}}Res struct {
	g.Meta `mime:"application/json"`
	Data   res.{{.EntityName}} `json:"data" dc:"{{.ChineseName}}信息"`
}

type Update{{.EntityName}}Req struct {
	g.Meta `path:"/{{.VarName}}/update/{Id}" method:"put" tags:"{{.ChineseName}}" summary:"更新{{.ChineseName}}." x-permission:"{{.ModuleName}}:{{.VarName}}:update"`
	model.AuthorHeader
	req.{{.EntityName}}Update
}

type Update{{.EntityName}}Res struct {
	g.Meta `mime:"application/json"`
}

type Delete{{.EntityName}}Req struct {
	g.Meta `path:"/{{.VarName}}/delete" method:"delete" tags:"{{.ChineseName}}" summary:"删除{{.ChineseName}}" x-permission:"{{.ModuleName}}:{{.VarName}}:delete"`
	model.AuthorHeader
	Ids []int64 `json:"ids" dc:"ids" v:"min-length:1#{{.ChineseName}}Id不能为空"`
}

type Delete{{.EntityName}}Res struct {
	g.Meta `mime:"application/json"`
}

type RealDelete{{.EntityName}}Req struct {
	g.Meta `path:"/{{.VarName}}/realDelete" method:"delete" tags:"{{.ChineseName}}" summary:"单个或批量真实删除{{.ChineseName}} （清空回收站）." x-permission:"{{.ModuleName}}:{{.VarName}}:realDelete"`
	model.AuthorHeader
	Ids []int64 `json:"ids" dc:"ids" v:"min-length:1#{{.ChineseName}}Id不能为空"`
}

type RealDelete{{.EntityName}}Res struct {
	g.Meta `mime:"application/json"`
}

type Recovery{{.EntityName}}Req struct {
	g.Meta `path:"/{{.VarName}}/recovery" method:"put" tags:"{{.ChineseName}}" summary:"单个或批量恢复在回收站的{{.ChineseName}}." x-permission:"{{.ModuleName}}:{{.VarName}}:recovery"`
	model.AuthorHeader
	Ids []int64 `json:"ids" dc:"ids" v:"min-length:1#{{.ChineseName}}Id不能为空"`
}

type Recovery{{.EntityName}}Res struct {
	g.Meta `mime:"application/json"`
}

type ChangeStatus{{.EntityName}}Req struct {
	g.Meta `path:"/{{.VarName}}/changeStatus" method:"put" tags:"{{.ChineseName}}" summary:"更改状态" x-permission:"{{.ModuleName}}:{{.VarName}}:update"`
	model.AuthorHeader
	Id     int64 `json:"id" dc:"ids" v:"min:1#Id不能为空"`
	Status int   `json:"status" dc:"status" v:"min:1#状态不能为空"`
}

type ChangeStatus{{.EntityName}}Res struct {
	g.Meta `mime:"application/json"`
}
