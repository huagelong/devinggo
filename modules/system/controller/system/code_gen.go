// Package system
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package system

import (
	"context"

	"devinggo/modules/system/api/system"
	"devinggo/modules/system/controller/base"
	"devinggo/modules/system/model/res"
	"devinggo/modules/system/service"

	"github.com/gogf/gf/v2/frame/g"
)

var (
	CodeGenController = codeGenController{}
)

type codeGenController struct {
	base.BaseController
}

// Index 获取代码生成列表
func (c *codeGenController) Index(ctx context.Context, in *system.CodeGenIndexReq) (out *system.CodeGenIndexRes, err error) {
	out = &system.CodeGenIndexRes{}
	items, totalCount, err := service.CodeGen().GetPageList(ctx, &in.PageListReq, &in.CodeGenSearch)
	if err != nil {
		return
	}
	if !g.IsEmpty(items) {
		out.Items = make([]res.CodeGenTable, 0, len(items))
		for _, item := range items {
			out.Items = append(out.Items, *item)
		}
	} else {
		out.Items = make([]res.CodeGenTable, 0)
	}
	out.PageRes.Pack(in, totalCount)
	return
}

// Delete 删除代码生成记录
func (c *codeGenController) Delete(ctx context.Context, in *system.CodeGenDeleteReq) (out *system.CodeGenDeleteRes, err error) {
	out = &system.CodeGenDeleteRes{}
	err = service.CodeGen().Delete(ctx, in.Ids)
	if err != nil {
		return
	}
	return
}

// Update 更新代码生成配置
func (c *codeGenController) Update(ctx context.Context, in *system.CodeGenUpdateReq) (out *system.CodeGenUpdateRes, err error) {
	out = &system.CodeGenUpdateRes{}
	err = service.CodeGen().Update(ctx, &in.CodeGenUpdate, c.UserId)
	if err != nil {
		return
	}
	return
}

// LoadTable 装载数据表
func (c *codeGenController) LoadTable(ctx context.Context, in *system.CodeGenLoadTableReq) (out *system.CodeGenLoadTableRes, err error) {
	out = &system.CodeGenLoadTableRes{}
	err = service.CodeGen().LoadTable(ctx, &in.CodeGenLoadTable, c.UserId)
	if err != nil {
		return
	}
	return
}

// Sync 同步数据表结构
func (c *codeGenController) Sync(ctx context.Context, in *system.CodeGenSyncReq) (out *system.CodeGenSyncRes, err error) {
	out = &system.CodeGenSyncRes{}
	err = service.CodeGen().SyncTable(ctx, in.Id, c.UserId)
	if err != nil {
		return
	}
	return
}

// Generate 生成代码
func (c *codeGenController) Generate(ctx context.Context, in *system.CodeGenGenerateReq) (out *system.CodeGenGenerateRes, err error) {
	out = &system.CodeGenGenerateRes{}
	fileBytes, err := service.CodeGen().GenerateCode(ctx, in.Ids)
	if err != nil {
		return
	}
	out.FileBytes = fileBytes
	return
}

// Preview 预览代码
func (c *codeGenController) Preview(ctx context.Context, in *system.CodeGenPreviewReq) (out *system.CodeGenPreviewRes, err error) {
	out = &system.CodeGenPreviewRes{}
	preview, err := service.CodeGen().PreviewCode(ctx, in.Id)
	if err != nil {
		return
	}
	out.Data = preview
	return
}

// ReadTable 读取表信息
func (c *codeGenController) ReadTable(ctx context.Context, in *system.CodeGenReadTableReq) (out *system.CodeGenReadTableRes, err error) {
	out = &system.CodeGenReadTableRes{}
	info, err := service.CodeGen().ReadTable(ctx, in.Id)
	if err != nil {
		return
	}
	out.Data = info
	return
}

// ListTable 获取数据源表列表
func (c *codeGenController) ListTable(ctx context.Context, in *system.CodeGenListTableReq) (out *system.CodeGenListTableRes, err error) {
	out = &system.CodeGenListTableRes{}
	tables, err := service.CodeGen().ListSourceTables(ctx, in.Source)
	if err != nil {
		return
	}
	out.Data = tables
	return
}
