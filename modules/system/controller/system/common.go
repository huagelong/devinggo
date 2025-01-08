// Package system
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package system

import (
	"context"
	"devinggo/internal/model/entity"
	"devinggo/modules/system/api/system"
	"devinggo/modules/system/consts"
	"devinggo/modules/system/controller/base"
	"devinggo/modules/system/model/req"
	"devinggo/modules/system/model/res"
	"devinggo/modules/system/pkg/cache"
	"devinggo/modules/system/pkg/utils"
	"devinggo/modules/system/service"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

var (
	CommonController = commonController{}
)

type commonController struct {
	base.BaseController
}

func (c *commonController) GetNoticeList(ctx context.Context, req *system.GetNoticeListReq) (rs *system.GetNoticeListRes, err error) {
	rs = &system.GetNoticeListRes{}
	items, totalCount, err := service.SystemNotice().GetPageList(ctx, &req.PageListReq)
	if err != nil {
		return
	}

	if !g.IsEmpty(items) {
		for _, item := range items {
			rs.Items = append(rs.Items, *item)
		}
	} else {
		rs.Items = make([]res.SystemNotice, 0)
	}
	rs.PageRes.Pack(req, totalCount)
	return
}

func (c *commonController) GetLoginLogList(ctx context.Context, req *system.GetLoginLogListReq) (rs *system.GetLoginLogListRes, err error) {
	rs = &system.GetLoginLogListRes{}
	systemUser, err := service.SystemUser().GetInfoById(ctx, c.UserId)
	if err != nil {
		return
	}
	items, totalCount, err := service.SystemLoginLog().GetPageList(ctx, &req.PageListReq, systemUser.Username)
	if err != nil {
		return
	}
	if !g.IsEmpty(items) {
		for _, item := range items {
			rs.Items = append(rs.Items, *item)
		}
	} else {
		rs.Items = make([]res.SystemLoginLog, 0)
	}
	rs.PageRes.Pack(req, totalCount)
	return
}

func (c *commonController) GetOperationLogList(ctx context.Context, req *system.GetOperationLogListReq) (rs *system.GetOperationLogListRes, err error) {
	rs = &system.GetOperationLogListRes{}
	systemUser, err := service.SystemUser().GetInfoById(ctx, c.UserId)
	if err != nil {
		return
	}
	items, totalCount, err := service.SystemOperLog().GetPageList(ctx, &req.PageListReq, systemUser.Username)
	if err != nil {
		return
	}

	if !g.IsEmpty(items) {
		for _, item := range items {
			rs.Items = append(rs.Items, *item)
		}
	} else {
		rs.Items = make([]res.SystemOperLog, 0)
	}
	rs.PageRes.Pack(req, totalCount)
	return
}

func (c *commonController) ClearAllCache(ctx context.Context, in *system.ClearAllCacheReq) (out *system.ClearAllCacheRes, err error) {
	out = &system.ClearAllCacheRes{}
	utils.SafeGo(ctx, func(ctx context.Context) {
		cache.RemoveByTag(ctx, consts.USER_CACHE_TAG+gconv.String(c.UserId))
	})
	return
}

func (c *commonController) GetUserList(ctx context.Context, in *system.GetUserListReq) (out *system.GetUserListRes, err error) {
	out = &system.GetUserListRes{}
	items, totalCount, err := service.SystemUser().GetPageList(ctx, &in.PageListReq)
	if err != nil {
		return
	}

	if !g.IsEmpty(items) {
		for _, item := range items {
			out.Items = append(out.Items, *item)
		}
	} else {
		out.Items = make([]res.SystemUser, 0)
	}
	out.PageRes.Pack(in, totalCount)
	return
}

func (c *commonController) GetUserInfoByIds(ctx context.Context, in *system.GetUserInfoByIdsReq) (out *system.GetUserInfoByIdsRes, err error) {
	out = &system.GetUserInfoByIdsRes{}
	items, err := service.SystemUser().GetInfoByIds(ctx, in.Ids)
	if err != nil {
		return
	}
	if !g.IsEmpty(items) {
		for _, item := range items {
			out.Data = append(out.Data, *item)
		}
	}
	return
}

func (c *commonController) GetDeptTreeList(ctx context.Context, in *system.GetDeptTreeListReq) (out *system.GetDeptTreeListRes, err error) {
	out = &system.GetDeptTreeListRes{}
	items, err := service.SystemDept().GetSelectTree(ctx, c.UserId)
	if err != nil {
		return
	}
	if !g.IsEmpty(items) {
		for _, item := range items {
			out.Data = append(out.Data, *item)
		}
	} else {
		out.Data = make([]res.SystemDeptTree, 0)
	}
	return
}

func (c *commonController) GetRoleList(ctx context.Context, in *system.GetRoleListReq) (out *system.GetRoleListRes, err error) {
	out = &system.GetRoleListRes{}
	rs, err := service.SystemRole().GetList(ctx, &req.SystemRoleSearch{}, true)
	if err != nil {
		return
	}

	if !g.IsEmpty(rs) {
		for _, v := range rs {
			out.Data = append(out.Data, *v)
		}
	} else {
		out.Data = make([]res.SystemRole, 0)
	}

	return
}

func (c *commonController) GetPostList(ctx context.Context, in *system.GetPostListReq) (out *system.GetPostListRes, err error) {
	out = &system.GetPostListRes{}
	rs, err := service.SystemPost().GetList(ctx, &req.SystemPostSearch{})
	if err != nil {
		return
	}

	if !g.IsEmpty(rs) {
		for _, v := range rs {
			out.Data = append(out.Data, *v)
		}
	} else {
		out.Data = make([]res.SystemPost, 0)
	}

	return
}

func (c *commonController) GetResourceList(ctx context.Context, in *system.GetResourceListReq) (out *system.GetResourceListRes, err error) {
	out = &system.GetResourceListRes{}
	items, totalCount, err := service.SystemUploadfile().GetPageList(ctx, &in.PageListReq, &in.SystemUploadFileSearch)
	if err != nil {
		return
	}

	if !g.IsEmpty(items) {
		for _, item := range items {
			out.Items = append(out.Items, *item)
		}
	} else {
		out.Items = make([]res.SystemUploadFile, 0)
	}
	out.PageRes.Pack(in, totalCount)

	return
}

func (c *commonController) GetModuleList(ctx context.Context, in *system.GetModuleListReq) (out *system.GetModuleListRes, err error) {
	out = &system.GetModuleListRes{}
	var dbModules []*entity.SystemModules
	err = service.SystemModules().Model(ctx).Where("status", 1).Scan(&dbModules)
	if utils.IsError(err) {
		return
	}
	data := make([]res.Module, 0)
	if !g.IsEmpty(dbModules) {
		for _, item := range dbModules {
			isEnable := false
			if item.Status == 1 {
				isEnable = true
			}

			installed := false
			if item.Installed == 1 {
				installed = true
			}
			data = append(data, res.Module{
				Name:        item.Name,
				Label:       item.Label,
				Description: item.Description,
				Installed:   installed,
				Enabled:     isEnable,
			})
		}
	}
	out.Data = data
	return
}
