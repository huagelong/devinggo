// Package system
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE
package system

import (
	"devinggo/internal/model"
	"devinggo/internal/model/page"
	"devinggo/modules/system/model/req"
	"devinggo/modules/system/model/res"
	"github.com/gogf/gf/v2/frame/g"
)

type IndexCodeReq struct {
	g.Meta `path:"/code/index" method:"get" tags:"代码生成" summary:"代码生成列表分页." x-permission:"system:code:index" `
	model.AuthorHeader
	model.PageListReq
	req.SettingGenerateTablesSearch
}

type IndexCodeRes struct {
	g.Meta `mime:"application/json"`
	page.PageRes
	Items []res.SettingGenerateTables `json:"items"  dc:"list" `
}

type GetDataSourceListReq struct {
	g.Meta `path:"/code/getDataSourceList" method:"get" tags:"代码生成" summary:"获取数据源列表." x-permission:"system:code:getDataSourceList" `
	model.AuthorHeader
	model.PageListReq
}

type GetDataSourceListRes struct {
	g.Meta `mime:"application/json"`
	page.PageRes
	Items []model.Dict `json:"items"  dc:"list" `
}

type LoadTableReq struct {
	g.Meta `path:"/code/loadTable" method:"post" tags:"代码生成" summary:"加载数据表." x-permission:"system:code:loadTable" `
	model.AuthorHeader
	req.LoadTable
}

type LoadTableRes struct {
	g.Meta `mime:"application/json"`
}

type ReadTableReq struct {
	g.Meta `path:"/code/readTable" method:"get" tags:"代码生成" summary:"读取表数据." x-permission:"system:code:readTable" `
	Id     uint64 `json:"id" dc:"id"`
}

type ReadTableRes struct {
	g.Meta `mime:"application/json"`
	Data   res.SettingGenerateTables `json:"data"  dc:"detail" `
}

type GetTableColumnsReq struct {
	g.Meta `path:"/code/getTableColumns" method:"get" tags:"代码生成" summary:"获取业务表字段信息." x-permission:"system:code:getTableColumns" `
	model.AuthorHeader
	req.SettingGenerateColumnsSearch
}

type GetTableColumnsRes struct {
	g.Meta `mime:"application/json"`
	Data   []res.SettingGenerateColumns `json:"data"  dc:"list" `
}

type PreviewCodeReq struct {
	g.Meta `path:"/code/preview" method:"get" tags:"代码生成" summary:"预览代码." x-permission:"system:code:preview" `
	model.AuthorHeader
	Id uint64 `json:"id" dc:"id"`
}

type PreviewCodeRes struct {
	g.Meta `mime:"application/json"`
	Data   []res.PreviewTable `json:"data"  dc:"list" `
}

type UpdateTableAndColumnsReq struct {
	g.Meta `path:"/code/update" method:"post" tags:"代码生成" summary:"更新业务表信息." x-permission:"system:code:update"`
	model.AuthorHeader
	req.TableAndColumnsUpdate
}

type UpdateTableAndColumnsRes struct {
	g.Meta `mime:"application/json"`
}

type GenerateCodeReq struct {
	g.Meta `path:"/code/generate" method:"post" tags:"代码生成" summary:"生成代码."  x-permission:"system:code:generate"`
	model.AuthorHeader
	Ids []uint64 `json:"ids" dc:"ids"  v:"min-length:1#Id不能为空"`
}

type GenerateCodeRes struct {
	g.Meta `mime:"application/json"`
}

type DeleteCodeReq struct {
	g.Meta `path:"/code/delete" method:"delete" tags:"代码生成" summary:"删除" x-permission:"system:code:delete"`
	model.AuthorHeader
	Ids []uint64 `json:"ids" dc:"ids" v:"min-length:1#Id不能为空"`
}

type DeleteCodeRes struct {
	g.Meta `mime:"application/json"`
}

type SyncCodeReq struct {
	g.Meta `path:"/code/sync/{Id}" method:"put" tags:"代码生成" summary:"同步数据库中的表信息跟字段." x-permission:"system:code:sync"`
	model.AuthorHeader
	Id uint64 `json:"id" dc:"id" v:"required"`
}

type SyncCodeRes struct {
	g.Meta `mime:"application/json"`
}
