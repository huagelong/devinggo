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

type ReceiveListReq struct {
	g.Meta `path:"/queueMessage/receiveList" method:"get" tags:"消息中心" summary:"接收消息列表." x-exceptAuth:"true" x-permission:"system:queueMessage:receiveList"`
	model.AuthorHeader
	model.PageListReq
	req.SystemQueueMessageSearch
}

type ReceiveListRes struct {
	g.Meta `mime:"application/json"`
	page.PageRes
	Items []res.SystemQueueMessage `json:"items"  dc:"message list" `
}

type UpdateReadStatusReq struct {
	g.Meta `path:"/queueMessage/updateReadStatus" method:"put" tags:"消息中心" summary:"更新状态到已读" x-exceptAuth:"true" x-permission:"system:queueMessage:updateReadStatus"`
	model.AuthorHeader
	Ids []int64 `json:"ids" v:"required" dc:"message ids"`
}

type UpdateReadStatusRes struct {
	g.Meta `mime:"application/json"`
}

type DeletesReq struct {
	g.Meta `path:"/queueMessage/deletes" method:"delete" tags:"消息中心" summary:"单个或批量删除数据到回收站" x-exceptAuth:"true" x-permission:"system:queueMessage:deletes"`
	model.AuthorHeader
	Ids []int64 `json:"ids" v:"required" dc:"message ids"`
}

type DeletesRes struct {
	g.Meta `mime:"application/json"`
}
