// Package system
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package system

import (
	"devinggo/modules/system/model"
	"devinggo/modules/system/model/res"
	"github.com/gogf/gf/v2/frame/g"
)

type GetCacheInfoReq struct {
	g.Meta `path:"/cache/monitor" method:"get" tags:"缓存" summary:"缓存信息." x-exceptAccessLog:"true" x-permission:"system:cache:monitor" `
	model.AuthorHeader
}

type GetCacheInfoRes struct {
	g.Meta `mime:"application/json"`
	Data   res.CacheInfo `json:"data"  dc:"cache info" `
}

type ViewCacheReq struct {
	g.Meta `path:"/cache/view" method:"post" tags:"缓存" summary:"查看key内容." x-permission:"system:cache:view" `
	model.AuthorHeader
	Key string `json:"key" v:"required"`
}

type ViewCacheRes struct {
	g.Meta  `mime:"application/json"`
	Content string `json:"content"  dc:"cache data" `
}

type DeleteCacheReq struct {
	g.Meta `path:"/cache/delete" method:"delete" tags:"缓存" summary:"根据key删除缓存." x-permission:"system:cache:delete" `
	model.AuthorHeader
	Key []string `json:"key" v:"required"`
}

type DeleteCacheRes struct {
	g.Meta `mime:"application/json"`
}

type DeleteAllCacheReq struct {
	g.Meta `path:"/cache/clear" method:"delete" tags:"缓存" summary:"晴空所有缓存." x-permission:"system:cache:clear" `
	model.AuthorHeader
}

type DeleteAllCacheRes struct {
	g.Meta `mime:"application/json"`
}
