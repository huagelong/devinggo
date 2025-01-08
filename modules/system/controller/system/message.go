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
	"devinggo/modules/system/service"
	"github.com/gogf/gf/v2/frame/g"
)

var (
	MessageController = messageController{}
)

type messageController struct {
	base.BaseController
}

func (c *messageController) ReceiveList(ctx context.Context, in *system.ReceiveListReq) (out *system.ReceiveListRes, err error) {
	out = &system.ReceiveListRes{}
	params := &req.SystemQueueMessageSearch{
		ReadStatus:  in.ReadStatus,
		ContentType: in.ContentType,
		Title:       in.Title,
		CreatedAt:   in.CreatedAt,
	}
	items, totalCount, err := service.SystemQueueMessage().GetPageList(ctx, &in.PageListReq, c.UserId, params)
	if err != nil {
		return
	}

	if !g.IsEmpty(items) {
		for _, item := range items {
			out.Items = append(out.Items, *item)
		}
	} else {
		out.Items = make([]res.SystemQueueMessage, 0)
	}
	out.PageRes.Pack(in, totalCount)
	return
}

func (c *messageController) UpdateReadStatus(ctx context.Context, in *system.UpdateReadStatusReq) (out *system.UpdateReadStatusRes, err error) {
	out = &system.UpdateReadStatusRes{}
	err = service.SystemQueueMessageReceive().UpdateReadStatus(ctx, in.Ids, c.UserId, 2)
	return
}

func (c *messageController) Deletes(ctx context.Context, in *system.DeletesReq) (out *system.DeletesRes, err error) {
	out = &system.DeletesRes{}
	err = service.SystemQueueMessage().DeletesRelated(ctx, in.Ids, c.UserId)
	return
}
