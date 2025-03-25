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

type IndexDictTypeReq struct {
	g.Meta `path:"/dictType/index" method:"get" tags:"数据字典" summary:"获取字典列表." x-permission:"system:dict:index" `
	model.AuthorHeader
	model.PageListReq
	req.SystemDictTypeSearch
}

type IndexDictTypeRes struct {
	g.Meta `mime:"application/json"`
	page.PageRes
	Items []res.SystemDictType `json:"items"  dc:"list" `
}

type DictTypeListReq struct {
	g.Meta `path:"/dictType/list" method:"get" tags:"字典" summary:"字典列表" x-exceptAuth:"true" x-permission:"system:dict:list" `
	model.AuthorHeader
	model.ListReq
	req.SystemDictTypeSearch
}

type DictTypeListRes struct {
	g.Meta `mime:"application/json"`
	Data   []res.SystemDictType `json:"data"`
}

type RecycleDictTypeReq struct {
	g.Meta `path:"/dictType/recycle" method:"get" tags:"字典" summary:"回收站列表." x-permission:"system:dict:recycle" `
	model.AuthorHeader
	model.PageListReq
	req.SystemDictTypeSearch
}

type RecycleDictTypeRes struct {
	g.Meta `mime:"application/json"`
	page.PageRes
	Items []res.SystemDictType `json:"items"  dc:"list" `
}

type SaveDictTypeReq struct {
	g.Meta `path:"/dictType/save" method:"post" tags:"字典" summary:"新增." x-permission:"system:dict:save"`
	model.AuthorHeader
	req.SystemDictTypeSave
}

type SaveDictTypeRes struct {
	g.Meta `mime:"application/json"`
	Id     int64 `json:"id" dc:"id"`
}

type ReadDictTypeReq struct {
	g.Meta `path:"/dictType/read/{Id}" method:"get" tags:"字典" summary:"获取一个字典类型数据." x-permission:"system:dict:read"`
	model.AuthorHeader
	Id int64 `json:"id" dc:"id" v:"required|min:1#Id不能为空"`
}

type ReadDictTypeRes struct {
	g.Meta `mime:"application/json"`
	Data   res.SystemDictType `json:"data" dc:"字典类型信息"`
}

type UpdateDictTypeReq struct {
	g.Meta `path:"/dictType/update/{Id}" method:"put" tags:"字典" summary:"更新." x-permission:"system:dict:update"`
	model.AuthorHeader
	req.SystemDictTypeUpdate
}

type UpdateDictTypeRes struct {
	g.Meta `mime:"application/json"`
}

type DeleteDictTypeReq struct {
	g.Meta `path:"/dictType/delete" method:"delete" tags:"字典" summary:"删除" x-permission:"system:dict:delete"`
	model.AuthorHeader
	Ids []int64 `json:"ids" dc:"ids" v:"min-length:1#Id不能为空"`
}

type DeleteDictTypeRes struct {
	g.Meta `mime:"application/json"`
}

type RealDeleteDictTypeReq struct {
	g.Meta `path:"/dictType/realDelete" method:"delete" tags:"字典" summary:"单个或批量真实删除 （清空回收站）." x-permission:"system:dict:realDelete"`
	model.AuthorHeader
	Ids []int64 `json:"ids" dc:"ids" v:"min-length:1#岗位Id不能为空"`
}

type RealDeleteDictTypeRes struct {
	g.Meta `mime:"application/json"`
}

type RecoveryDictTypeReq struct {
	g.Meta `path:"/dictType/recovery" method:"put" tags:"字典" summary:"单个或批量恢复在回收站的数据." x-permission:"system:dict:recovery"`
	model.AuthorHeader
	Ids []int64 `json:"ids" dc:"ids" v:"min-length:1#Id不能为空"`
}

type RecoveryDictTypeRes struct {
	g.Meta `mime:"application/json"`
}

type ChangeStatusDictTypeReq struct {
	g.Meta `path:"/dictType/changeStatus" method:"put" tags:"字典" summary:"更改状态" x-permission:"system:dict:changeStatus"`
	model.AuthorHeader
	Id     int64 `json:"id" dc:"ids" v:"min:1#Id不能为空"`
	Status int   `json:"status" dc:"status" v:"min:1#状态不能为空"`
}

type ChangeStatusDictTypeRes struct {
	g.Meta `mime:"application/json"`
}

type DataDictListReq struct {
	g.Meta `path:"/dataDict/list" method:"get" tags:"字典" summary:"快捷查询一个字典" x-exceptAuth:"true" x-permission:"system:dataDict:list" `
	model.AuthorHeader
	model.ListReq
	req.SystemDictDataSearch
}

type DataDictListRes struct {
	g.Meta `mime:"application/json"`
	Data   []res.SystemDictData `json:"data"`
}

type IndexDictDataReq struct {
	g.Meta `path:"/dataDict/index" method:"get" tags:"数据字典" summary:"获取字典列表." x-permission:"system:dict:index" `
	model.AuthorHeader
	model.PageListReq
	req.SystemDictDataSearch
}

type IndexDictDataRes struct {
	g.Meta `mime:"application/json"`
	page.PageRes
	Items []res.SystemDictDataFull `json:"items"  dc:"list" `
}

type DataDictListsReq struct {
	g.Meta `path:"/dataDict/lists" method:"get" tags:"字典" summary:"快捷查询一个字典" x-exceptAuth:"true" x-permission:"system:dataDict:lists" `
	model.AuthorHeader
	model.ListReq
	Codes string `json:"codes" dc:"codes" v:"required#字典编码不能为空"`
}

type DataDictListsRes struct {
	g.Meta `mime:"application/json"`
	Data   map[string]res.SystemDictData `json:"data"`
}

type ClearCacheDictDataReq struct {
	g.Meta `path:"/dataDict/clearCache" method:"post" tags:"字典" summary:"清除字典缓存." x-permission:"system:dict:clearCache"`
	model.AuthorHeader
}

type ClearCacheDictDataRes struct {
	g.Meta `mime:"application/json"`
}

type RecycleDictDataReq struct {
	g.Meta `path:"/dataDict/recycle" method:"get" tags:"字典" summary:"回收站列表." x-permission:"system:dict:recycle" `
	model.AuthorHeader
	model.PageListReq
	req.SystemDictDataSearch
}

type RecycleDictDataRes struct {
	g.Meta `mime:"application/json"`
	page.PageRes
	Items []res.SystemDictDataFull `json:"items"  dc:"list" `
}

type SaveDictDataReq struct {
	g.Meta `path:"/dataDict/save" method:"post" tags:"字典" summary:"新增." x-permission:"system:dict:save"`
	model.AuthorHeader
	req.SystemDictDataSave
}

type SaveDictDataRes struct {
	g.Meta `mime:"application/json"`
	Id     int64 `json:"id" dc:"id"`
}

type ReadDictDataReq struct {
	g.Meta `path:"/dataDict/read/{Id}" method:"get" tags:"字典" summary:"获取一个字典数据." x-permission:"system:dict:read"`
	model.AuthorHeader
	Id int64 `json:"id" dc:"id" v:"required|min:1#Id不能为空"`
}

type ReadDictDataRes struct {
	g.Meta `mime:"application/json"`
	Data   res.SystemDictDataFull `json:"data" dc:"字典类型信息"`
}

type UpdateDictDataReq struct {
	g.Meta `path:"/dataDict/update/{Id}" method:"put" tags:"字典" summary:"更新." x-permission:"system:dict:update"`
	model.AuthorHeader
	req.SystemDictDataUpdate
}

type UpdateDictDataRes struct {
	g.Meta `mime:"application/json"`
}

type DeleteDictDataReq struct {
	g.Meta `path:"/dataDict/delete" method:"delete" tags:"字典" summary:"删除" x-permission:"system:dict:delete"`
	model.AuthorHeader
	Ids []int64 `json:"ids" dc:"ids" v:"min-length:1#Id不能为空"`
}

type DeleteDictDataRes struct {
	g.Meta `mime:"application/json"`
}

type RealDeleteDictDataReq struct {
	g.Meta `path:"/dataDict/realDelete" method:"delete" tags:"字典" summary:"单个或批量真实删除 （清空回收站）." x-permission:"system:dict:realDelete"`
	model.AuthorHeader
	Ids []int64 `json:"ids" dc:"ids" v:"min-length:1#岗位Id不能为空"`
}

type RealDeleteDictDataRes struct {
	g.Meta `mime:"application/json"`
}

type RecoveryDictDataReq struct {
	g.Meta `path:"/dataDict/recovery" method:"put" tags:"字典" summary:"单个或批量恢复在回收站的数据." x-permission:"system:dict:recovery"`
	model.AuthorHeader
	Ids []int64 `json:"ids" dc:"ids" v:"min-length:1#Id不能为空"`
}

type RecoveryDictDataRes struct {
	g.Meta `mime:"application/json"`
}

type ChangeStatusDictDataReq struct {
	g.Meta `path:"/dataDict/changeStatus" method:"put" tags:"字典" summary:"更改状态" x-permission:"system:dict:changeStatus"`
	model.AuthorHeader
	Id     int64 `json:"id" dc:"ids" v:"min:1#Id不能为空"`
	Status int   `json:"status" dc:"status" v:"min:1#状态不能为空"`
}

type ChangeStatusDictDataRes struct {
	g.Meta `mime:"application/json"`
}

type NumberOperationDictDataReq struct {
	g.Meta `path:"/dataDict/numberOperation" method:"put" tags:"字典" summary:"数字运算操作." x-permission:"system:dict:update"`
	model.AuthorHeader
	Id          int64  `json:"id" dc:"ids" v:"min:1#Id不能为空"`
	NumberName  string `json:"numberName" dc:"numberName" v:"required#名称不能为空"`
	NumberValue int    `json:"numberValue" dc:"number Value" d:"1" v:"min:1#数字不能为空"`
}

type NumberOperationDictDataRes struct {
	g.Meta `mime:"application/json"`
}
