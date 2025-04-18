// Package system
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE
package system

import (
	"context"
	"devinggo/internal/dao"
	"devinggo/modules/system/api/system"
	"devinggo/modules/system/controller/base"
	"devinggo/modules/system/model"
	"devinggo/modules/system/model/res"
	"devinggo/modules/system/pkg/utils/config"
	"devinggo/modules/system/pkg/utils/request"
	"devinggo/modules/system/service"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

var (
	CodeController = codeController{}
)

type codeController struct {
	base.BaseController
}

func (c *codeController) Index(ctx context.Context, in *system.IndexCodeReq) (out *system.IndexCodeRes, err error) {
	out = &system.IndexCodeRes{}
	items, totalCount, err := service.SettingGenerateTables().GetPageListForSearch(ctx, &in.PageListReq, &in.SettingGenerateTablesSearch)
	if err != nil {
		return
	}

	if !g.IsEmpty(items) {
		for _, item := range items {
			out.Items = append(out.Items, *item)
		}
	} else {
		out.Items = make([]res.SettingGenerateTables, 0)
	}
	out.PageRes.Pack(in, totalCount)
	return
}

func (c *codeController) GetDataSourceList(ctx context.Context, in *system.GetDataSourceListReq) (out *system.GetDataSourceListRes, err error) {
	out = &system.GetDataSourceListRes{}
	databaseConfig := config.GetConfigMap(ctx, "database")
	dicts := make([]model.Dict, 0)
	var dict model.Dict
	totalCount := 0
	if !g.IsEmpty(databaseConfig) {
		for key, _ := range databaseConfig {
			if key == "logger" {
				continue
			}
			dict = model.Dict{}
			dict.Label = key
			dict.Value = key
			dicts = append(dicts, dict)
		}
		out.Items = dicts
	} else {
		out.Items = make([]model.Dict, 0)
	}
	totalCount = len(out.Items)
	out.PageRes.Pack(in, totalCount)
	return
}

func (c *codeController) LoadTable(ctx context.Context, in *system.LoadTableReq) (out *system.LoadTableRes, err error) {
	out = &system.LoadTableRes{}
	err = service.SettingGenerateTables().LoadTable(ctx, &in.LoadTable)
	return
}

func (c *codeController) ReadTable(ctx context.Context, in *system.ReadTableReq) (out *system.ReadTableRes, err error) {
	out = &system.ReadTableRes{}
	data, err := service.SettingGenerateTables().GetById(ctx, in.Id)
	if err != nil {
		return
	}
	out.Data = *data
	return
}

func (c *codeController) GetTableColumns(ctx context.Context, in *system.GetTableColumnsReq) (out *system.GetTableColumnsRes, err error) {
	out = &system.GetTableColumnsRes{}
	rs, err := service.SettingGenerateColumns().GetList(ctx, &in.SettingGenerateColumnsSearch)
	if err != nil {
		return
	}

	if !g.IsEmpty(rs) {
		for _, item := range rs {
			out.Data = append(out.Data, *item)
		}
	} else {
		out.Data = make([]res.SettingGenerateColumns, 0)
	}
	return
}

func (c *codeController) PreviewCode(ctx context.Context, in *system.PreviewCodeReq) (out *system.PreviewCodeRes, err error) {
	out = &system.PreviewCodeRes{}
	rs, err := service.SettingGenerateTables().Preview(ctx, in.Id)
	if err != nil {
		return
	}
	out.Data = rs
	return
}

func (c *codeController) UpdateTableAndColumns(ctx context.Context, in *system.UpdateTableAndColumnsReq) (out *system.UpdateTableAndColumnsRes, err error) {
	out = &system.UpdateTableAndColumnsRes{}
	err = dao.SettingGenerateTables.Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		err = service.SettingGenerateTables().UpdateTableAndColumns(ctx, &in.TableAndColumnsUpdate)
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

func (c *codeController) GenerateCode(ctx context.Context, in *system.GenerateCodeReq) (out *system.GenerateCodeRes, err error) {
	out = &system.GenerateCodeRes{}
	filePath, err := service.SettingGenerateTables().GenerateCode(ctx, in.Ids)
	if err != nil {
		return
	}
	r := request.GetHttpRequest(ctx)
	r.Response.ServeFileDownload(filePath)
	return
}

func (c *codeController) DeleteCode(ctx context.Context, in *system.DeleteCodeReq) (out *system.DeleteCodeRes, err error) {
	out = &system.DeleteCodeRes{}
	err = dao.SettingGenerateTables.Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		err = service.SettingGenerateTables().Delete(ctx, in.Ids)
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

func (c *codeController) SyncCode(ctx context.Context, in *system.SyncCodeReq) (out *system.SyncCodeRes, err error) {
	out = &system.SyncCodeRes{}
	err = dao.SettingGenerateTables.Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		err = service.SettingGenerateTables().SyncCode(ctx, in.Id)
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
