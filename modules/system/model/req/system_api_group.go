// Package req
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package req

type SystemApiGroupSearch struct {
	Name       string `json:"name"`
	Status     int    `json:"status"`
	GetApiList bool   `json:"getApiList"`
}

type SystemApiGroupSave struct {
	Name   string `json:"name" v:"required"`
	Status int    `json:"status"`
	Remark string `json:"remark"`
}

type SystemApiGroupUpdate struct {
	Id     int    `json:"id" v:"required"`
	Name   string `json:"name" v:"required"`
	Status int    `json:"status"`
	Remark string `json:"remark"`
}
