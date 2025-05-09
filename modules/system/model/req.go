// Package req
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package model

import (
	"devinggo/modules/system/model/page"
)

type AuthorHeader struct {
	Authorization string `json:"Authorization" in:"header" v:"required" default:""   dc:"token"`
	Lang          string `json:"Accept-Language" in:"header" default:"zh_CN"  dc:"i18n lang"`
	Xappid        string `json:"X-App-Id" in:"header" default:"1000"   dc:"app id"`
}

type EasyModeVerify struct {
	Signature string `json:"signature"  v:"required"  default:""    dc:"sign"`
	AppId     string `json:"app_id"  v:"required" default:""   dc:"app id"`
	Timestamp string `json:"timestamp"  v:"required" default:""    dc:"timestamp"`
	Nonce     string `json:"nonce"  v:"required" default:""    dc:"nonce"`
	Lang      string `json:"language" default:"zh_CN"  dc:"i18n lang"`
}

type NormalModeVerify struct {
	Authorization string `json:"Authorization" in:"header" v:"required" default:""   dc:"token"`
	Lang          string `json:"Accept-Language" in:"header" v:"required" default:"zh_CN"  dc:"i18n lang"`
	Xappid        string `json:"X-App-Id" in:"header" v:"required" default:"1000"   dc:"app id"`
}

type ApiSign struct {
	Authorization string `json:"Authorization" in:"header"  default:""   dc:"token"`
	Lang          string `json:"Accept-Language" in:"header" default:"zh_CN"  dc:"i18n lang"`
	Xappid        string `json:"X-App-Id" in:"header"  default:"1000"   dc:"app id"`
	Signature     string `json:"signature"   default:""    dc:"sign"`
	AppId         string `json:"app_id"  default:""   dc:"app id"`
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
