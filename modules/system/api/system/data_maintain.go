// Package system
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE
package system

import (
	"devinggo/modules/system/model"
	"devinggo/modules/system/model/page"
	"devinggo/modules/system/model/req"
	"devinggo/modules/system/model/res"
	"github.com/gogf/gf/v2/frame/g"
)

type IndexDataMaintainReq struct {
	g.Meta `path:"/dataMaintain/index" method:"get" tags:"数据维护" summary:"列表." x-permission:"system:dataMaintain:index" `
	model.AuthorHeader
	model.PageListReq
	req.DataMaintainSearch
}

type IndexDataMaintainRes struct {
	g.Meta `mime:"application/json"`
	page.PageRes
	Items []res.DataMaintain `json:"items"  dc:"list" `
}
