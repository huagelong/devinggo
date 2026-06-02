// Package req
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package req

type SystemDictTypeSearch struct {
	Code   string `json:"code"`
	Name   string `json:"name"`
	Status int    `json:"status"`
}

type SystemDictTypeSave struct {
	Code   string `json:"code" v:"required"`
	Name   string `json:"name" v:"required"`
	Status int    `json:"status"`
	Remark string `json:"remark"`
}

type SystemDictTypeUpdate struct {
	Id     int64  `json:"id"`
	Code   string `json:"code" v:"required"`
	Name   string `json:"name" v:"required"`
	Status int    `json:"status"`
	Remark string `json:"remark"`
}
