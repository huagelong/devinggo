// Package req
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package model

import (
	"devinggo/internal/model/page"
)

type AuthorHeader struct {
	Authorization string `json:"Authorization" in:"header" v:"required" default:""   dc:"token"`
	Lang          string `json:"Accept-Language" in:"header" v:"required" default:"zh_CN"  dc:"i18n lang"`
}

type SignHeader struct {
	Xtimestamp string `json:"X-Timestamp" in:"header" v:"required"  default:"1693529263225"    dc:"time stamp"`
	Xnonce     string `json:"X-Nonce" in:"header" v:"required" default:"617266194"   dc:"rand number"`
	Xsign      string `json:"X-Sign" in:"header" v:"required" default:"e01844a1413a236a60e0167fbd62283f"    dc:"sign"`
	Xappid     string `json:"X-Appid" in:"header" v:"required" default:"1000"   dc:"app id"`
}

type PageListReq struct {
	page.PageReq
	OrderBy    string `json:"orderBy" default:"" dc:"order by"`
	OrderType  string `json:"orderType" default:"" dc:"order by type"`
	Select     string `json:"select" default:"" dc:"select"`
	Recycle    bool   `json:"recycle" default:"false" dc:"show deleted data"`
	FilterAuth bool   `json:"filterAuth" default:"false" dc:"filter auth data"`
}

type ListReq struct {
	OrderBy    string `json:"orderBy" default:"" dc:"order by"`
	OrderType  string `json:"orderType" default:"" dc:"order by type"`
	Select     string `json:"select" default:"*" dc:"select"`
	Recycle    bool   `json:"recycle" default:"false" dc:"show deleted data"`
	FilterAuth bool   `json:"filterAuth" default:"false" dc:"filter auth data"`
}
