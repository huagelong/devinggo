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
	"devinggo/modules/system/model/req"
	"devinggo/modules/system/model/res"
	"devinggo/modules/system/pkg/cache"
	"devinggo/modules/system/service"
	"github.com/gogf/gf/v2/frame/g"
)

var (
	DictController = dictController{}
)

type dictController struct {
	base.BaseController
}

func (c *dictController) IndexDictType(ctx context.Context, in *system.IndexDictTypeReq) (out *system.IndexDictTypeRes, err error) {
	out = &system.IndexDictTypeRes{}
	items, totalCount, err := service.SystemDictType().GetPageList(ctx, &in.PageListReq, &in.SystemDictTypeSearch)
	if err != nil {
		return
	}

	if !g.IsEmpty(items) {
		for _, item := range items {
			out.Items = append(out.Items, *item)
		}
	} else {
		out.Items = make([]res.SystemDictType, 0)
	}
	out.PageRes.Pack(in, totalCount)
	return
}

func (c *dictController) IndexDictData(ctx context.Context, in *system.IndexDictDataReq) (out *system.IndexDictDataRes, err error) {
	out = &system.IndexDictDataRes{}
	items, totalCount, err := service.SystemDictData().GetPageList(ctx, &in.PageListReq, &in.SystemDictDataSearch)
	if err != nil {
		return
	}

	if !g.IsEmpty(items) {
		for _, item := range items {
			out.Items = append(out.Items, *item)
		}
	} else {
		out.Items = make([]res.SystemDictDataFull, 0)
	}
	out.PageRes.Pack(in, totalCount)
	return
}

func (c *dictController) DataTypeList(ctx context.Context, in *system.DictTypeListReq) (out *system.DictTypeListRes, err error) {
	out = &system.DictTypeListRes{}
	rs, err := service.SystemDictType().GetList(ctx, &in.ListReq, &in.SystemDictTypeSearch)
	if err != nil {
		return
	}
	if !g.IsEmpty(rs) {
		for _, v := range rs {
			out.Data = append(out.Data, *v)
		}
	} else {
		out.Data = make([]res.SystemDictType, 0)
	}

	return
}

func (c *dictController) RecycleDictType(ctx context.Context, in *system.RecycleDictTypeReq) (out *system.RecycleDictTypeRes, err error) {
	out = &system.RecycleDictTypeRes{}
	in.Recycle = true
	items, totalCount, err := service.SystemDictType().GetPageList(ctx, &in.PageListReq, &in.SystemDictTypeSearch)
	if err != nil {
		return
	}

	if !g.IsEmpty(items) {
		for _, item := range items {
			out.Items = append(out.Items, *item)
		}
	} else {
		out.Items = make([]res.SystemDictType, 0)
	}
	out.PageRes.Pack(in, totalCount)
	return
}

func (c *dictController) RecycleDictData(ctx context.Context, in *system.RecycleDictDataReq) (out *system.RecycleDictDataRes, err error) {
	out = &system.RecycleDictDataRes{}
	in.Recycle = true
	items, totalCount, err := service.SystemDictData().GetPageList(ctx, &in.PageListReq, &in.SystemDictDataSearch)
	if err != nil {
		return
	}

	if !g.IsEmpty(items) {
		for _, item := range items {
			out.Items = append(out.Items, *item)
		}
	} else {
		out.Items = make([]res.SystemDictDataFull, 0)
	}
	out.PageRes.Pack(in, totalCount)
	return
}

func (c *dictController) SaveDictType(ctx context.Context, in *system.SaveDictTypeReq) (out *system.SaveDictTypeRes, err error) {
	out = &system.SaveDictTypeRes{}
	id, err := service.SystemDictType().Save(ctx, &in.SystemDictTypeSave)
	if err != nil {
		return
	}
	out.Id = id
	return
}

func (c *dictController) SaveDictData(ctx context.Context, in *system.SaveDictDataReq) (out *system.SaveDictDataRes, err error) {
	out = &system.SaveDictDataRes{}
	id, err := service.SystemDictData().Save(ctx, &in.SystemDictDataSave)
	if err != nil {
		return
	}
	out.Id = id
	return
}

func (c *dictController) ReadDictType(ctx context.Context, in *system.ReadDictTypeReq) (out *system.ReadDictTypeRes, err error) {
	out = &system.ReadDictTypeRes{}
	info, err := service.SystemDictType().GetById(ctx, in.Id)
	if err != nil {
		return
	}
	out.Data = *info
	return
}

func (c *dictController) ReadDictData(ctx context.Context, in *system.ReadDictDataReq) (out *system.ReadDictDataRes, err error) {
	out = &system.ReadDictDataRes{}
	info, err := service.SystemDictData().GetById(ctx, in.Id)
	if err != nil {
		return
	}
	out.Data = *info
	return
}

func (c *dictController) UpdateDictType(ctx context.Context, in *system.UpdateDictTypeReq) (out *system.UpdateDictTypeRes, err error) {
	out = &system.UpdateDictTypeRes{}
	err = service.SystemDictType().Update(ctx, &in.SystemDictTypeUpdate)
	if err != nil {
		return
	}
	return
}

func (c *dictController) UpdateDictData(ctx context.Context, in *system.UpdateDictDataReq) (out *system.UpdateDictDataRes, err error) {
	out = &system.UpdateDictDataRes{}
	err = service.SystemDictData().Update(ctx, &in.SystemDictDataUpdate)
	if err != nil {
		return
	}
	return
}

func (c *dictController) DeleteDictType(ctx context.Context, in *system.DeleteDictTypeReq) (out *system.DeleteDictTypeRes, err error) {
	out = &system.DeleteDictTypeRes{}
	err = service.SystemDictType().Delete(ctx, in.Ids)
	if err != nil {
		return
	}
	return
}

func (c *dictController) DeleteDictData(ctx context.Context, in *system.DeleteDictDataReq) (out *system.DeleteDictDataRes, err error) {
	out = &system.DeleteDictDataRes{}
	err = service.SystemDictData().Delete(ctx, in.Ids)
	if err != nil {
		return
	}
	return
}

func (c *dictController) RealDeleteDictType(ctx context.Context, in *system.RealDeleteDictTypeReq) (out *system.RealDeleteDictTypeRes, err error) {
	out = &system.RealDeleteDictTypeRes{}
	err = service.SystemDictType().RealDelete(ctx, in.Ids)
	if err != nil {
		return
	}
	return
}

func (c *dictController) RealDeleteDictData(ctx context.Context, in *system.RealDeleteDictDataReq) (out *system.RealDeleteDictDataRes, err error) {
	out = &system.RealDeleteDictDataRes{}
	err = service.SystemDictData().RealDelete(ctx, in.Ids)
	if err != nil {
		return
	}
	return
}

func (c *dictController) RecoveryDictType(ctx context.Context, in *system.RecoveryDictTypeReq) (out *system.RecoveryDictTypeRes, err error) {
	out = &system.RecoveryDictTypeRes{}
	err = service.SystemDictType().Recovery(ctx, in.Ids)
	if err != nil {
		return
	}
	return
}

func (c *dictController) RecoveryDictData(ctx context.Context, in *system.RecoveryDictDataReq) (out *system.RecoveryDictDataRes, err error) {
	out = &system.RecoveryDictDataRes{}
	err = service.SystemDictData().Recovery(ctx, in.Ids)
	if err != nil {
		return
	}
	return
}

func (c *dictController) ChangeStatusDictType(ctx context.Context, in *system.ChangeStatusDictTypeReq) (out *system.ChangeStatusDictTypeRes, err error) {
	out = &system.ChangeStatusDictTypeRes{}
	err = service.SystemDictType().ChangeStatus(ctx, in.Id, in.Status)
	if err != nil {
		return
	}
	return
}

func (c *dictController) ChangeStatusDictData(ctx context.Context, in *system.ChangeStatusDictDataReq) (out *system.ChangeStatusDictDataRes, err error) {
	out = &system.ChangeStatusDictDataRes{}
	err = service.SystemDictData().ChangeStatus(ctx, in.Id, in.Status)
	if err != nil {
		return
	}
	return
}

func (c *dictController) DataDictList(ctx context.Context, in *system.DataDictListReq) (out *system.DataDictListRes, err error) {
	out = &system.DataDictListRes{}
	rs, err := service.SystemDictData().GetList(ctx, &in.ListReq, &in.SystemDictDataSearch)
	if err != nil {
		return
	}
	if !g.IsEmpty(rs) {
		for _, v := range rs {
			out.Data = append(out.Data, *v)
		}
	} else {
		out.Data = make([]res.SystemDictData, 0)
	}

	return
}

func (c *dictController) DataDictLists(ctx context.Context, in *system.DataDictListsReq) (out *system.DataDictListsRes, err error) {
	out = &system.DataDictListsRes{}
	search := &req.SystemDictDataSearch{
		Codes: in.Codes,
	}
	rs, err := service.SystemDictData().GetList(ctx, &in.ListReq, search)
	if err != nil {
		return
	}
	if !g.IsEmpty(rs) {
		mapTmp := make(map[string]res.SystemDictData)
		for _, v := range rs {
			mapTmp[v.Code] = *v
		}
		out.Data = mapTmp
	} else {
		out.Data = make(map[string]res.SystemDictData, 0)
	}

	return
}

func (c *dictController) ClearCacheDictData(ctx context.Context, in *system.ClearCacheDictDataReq) (out *system.ClearCacheDictDataRes, err error) {
	err = cache.ClearByTable(ctx, "system_dict_type")
	if err != nil {
		return
	}
	err = cache.ClearByTable(ctx, "system_dict_data")
	if err != nil {
		return
	}
	return
}

func (c *dictController) NumberOperation(ctx context.Context, in *system.NumberOperationDictDataReq) (out *system.NumberOperationDictDataRes, err error) {
	out = &system.NumberOperationDictDataRes{}
	err = service.SystemDictData().NumberOperation(ctx, in.Id, in.NumberName, in.NumberValue)
	if err != nil {
		return
	}
	return
}
