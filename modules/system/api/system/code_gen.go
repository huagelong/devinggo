// Package system
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package system

import (
	"devinggo/modules/system/model"
	"devinggo/modules/system/model/page"
	"devinggo/modules/system/model/req"
	"devinggo/modules/system/model/res"

	"github.com/gogf/gf/v2/frame/g"
)

// CodeGenIndexReq 代码生成列表请求
type CodeGenIndexReq struct {
	g.Meta `path:"/code/index" method:"get" tags:"代码生成" summary:"获取代码生成列表." x-permission:"system:code:index"`
	model.AuthorHeader
	model.PageListReq
	req.CodeGenSearch
}

// CodeGenIndexRes 代码生成列表响应
type CodeGenIndexRes struct {
	g.Meta `mime:"application/json"`
	page.PageRes
	Items []res.CodeGenTable `json:"items" dc:"list"`
}

// CodeGenDeleteReq 删除代码生成请求
type CodeGenDeleteReq struct {
	g.Meta `path:"/code/delete" method:"delete" tags:"代码生成" summary:"删除代码生成记录." x-permission:"system:code:delete"`
	model.AuthorHeader
	Ids []int64 `json:"ids" dc:"ids" v:"min-length:1#Id不能为空"`
}

// CodeGenDeleteRes 删除代码生成响应
type CodeGenDeleteRes struct {
	g.Meta `mime:"application/json"`
}

// CodeGenUpdateReq 更新代码生成配置请求
type CodeGenUpdateReq struct {
	g.Meta `path:"/code/update" method:"post" tags:"代码生成" summary:"更新代码生成配置." x-permission:"system:code:update"`
	model.AuthorHeader
	req.CodeGenUpdate
}

// CodeGenUpdateRes 更新代码生成配置响应
type CodeGenUpdateRes struct {
	g.Meta `mime:"application/json"`
}

// CodeGenLoadTableReq 装载数据表请求
type CodeGenLoadTableReq struct {
	g.Meta `path:"/code/loadTable" method:"post" tags:"代码生成" summary:"装载数据表." x-permission:"system:code:loadTable"`
	model.AuthorHeader
	req.CodeGenLoadTable
}

// CodeGenLoadTableRes 装载数据表响应
type CodeGenLoadTableRes struct {
	g.Meta `mime:"application/json"`
}

// CodeGenSyncReq 同步数据表请求
type CodeGenSyncReq struct {
	g.Meta `path:"/code/sync/{Id}" method:"put" tags:"代码生成" summary:"同步数据表结构." x-permission:"system:code:sync"`
	model.AuthorHeader
	Id int64 `json:"id" dc:"id" v:"required|min:1#Id不能为空"`
}

// CodeGenSyncRes 同步数据表响应
type CodeGenSyncRes struct {
	g.Meta `mime:"application/json"`
}

// CodeGenGenerateReq 生成代码请求
type CodeGenGenerateReq struct {
	g.Meta `path:"/code/generate" method:"post" tags:"代码生成" summary:"生成代码." x-permission:"system:code:generate"`
	model.AuthorHeader
	req.CodeGenGenerate
}

// CodeGenGenerateRes 生成代码响应
type CodeGenGenerateRes struct {
	g.Meta    `mime:"application/json" types:"blob"`
	FileBytes []byte `json:"fileBytes" dc:"生成的代码压缩文件"`
}

// CodeGenPreviewReq 预览代码请求
type CodeGenPreviewReq struct {
	g.Meta `path:"/code/preview" method:"get" tags:"代码生成" summary:"预览代码." x-permission:"system:code:preview"`
	model.AuthorHeader
	Id int64 `json:"id" dc:"id" v:"required|min:1#Id不能为空"`
}

// CodeGenPreviewRes 预览代码响应
type CodeGenPreviewRes struct {
	g.Meta `mime:"application/json"`
	Data   []res.CodeGenPreview `json:"data"`
}

// CodeGenReadTableReq 读取表信息请求
type CodeGenReadTableReq struct {
	g.Meta `path:"/code/readTable/{Id}" method:"get" tags:"代码生成" summary:"读取表信息." x-permission:"system:code:readTable"`
	model.AuthorHeader
	Id int64 `json:"id" dc:"id" v:"required|min:1#Id不能为空"`
}

// CodeGenReadTableRes 读取表信息响应
type CodeGenReadTableRes struct {
	g.Meta `mime:"application/json"`
	Data   res.CodeGenReadTable `json:"data"`
}

// CodeGenListTableReq 获取数据源表列表请求
type CodeGenListTableReq struct {
	g.Meta `path:"/code/listTable" method:"get" tags:"代码生成" summary:"获取数据源表列表." x-permission:"system:code:listTable"`
	model.AuthorHeader
	Source string `json:"source" dc:"数据源"`
}

// CodeGenListTableRes 获取数据源表列表响应
type CodeGenListTableRes struct {
	g.Meta `mime:"application/json"`
	Data   []res.CodeGenSourceTable `json:"data"`
}
