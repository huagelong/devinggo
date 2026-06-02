// Package req
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package req

type SystemPostSearch struct {
	Name      string   `json:"name"`
	Code      string   `json:"code"`
	Status    int      `json:"status"`
	CreatedAt []string `json:"created_at" dc:"created at"`
}

type SystemPostSave struct {
	Id     int64  `json:"id"`
	Sort   int    `json:"sort"`
	Status int    `json:"status"`
	Name   string `json:"name"`
	Code   string `json:"code"`
	Remark string `json:"remark"`
}
