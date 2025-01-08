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
	AttachmentController = attachmentController{}
)

type attachmentController struct {
	base.BaseController
}

func (c *attachmentController) Index(ctx context.Context, in *system.IndexAttachmentReq) (out *system.IndexAttachmentRes, err error) {
	out = &system.IndexAttachmentRes{}
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

func (c *attachmentController) Recycle(ctx context.Context, in *system.RecycleAttachmentReq) (out *system.RecycleAttachmentRes, err error) {
	out = &system.RecycleAttachmentRes{}
	in.Recycle = true
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

func (c *attachmentController) Delete(ctx context.Context, in *system.DeleteAttachmentReq) (out *system.DeleteAttachmentRes, err error) {
	out = &system.DeleteAttachmentRes{}
	err = service.SystemUploadfile().Delete(ctx, in.Ids)
	if err != nil {
		return
	}
	return
}

func (c *attachmentController) RealDelete(ctx context.Context, in *system.RealDeleteAttachmentReq) (out *system.RealDeleteAttachmentRes, err error) {
	out = &system.RealDeleteAttachmentRes{}
	err = service.SystemUploadfile().RealDelete(ctx, in.Ids)
	if err != nil {
		return
	}
	return
}

func (c *attachmentController) Recovery(ctx context.Context, in *system.RecoveryAttachmentReq) (out *system.RecoveryAttachmentRes, err error) {
	out = &system.RecoveryAttachmentRes{}
	err = service.SystemUploadfile().Recovery(ctx, in.Ids)
	if err != nil {
		return
	}
	return
}
