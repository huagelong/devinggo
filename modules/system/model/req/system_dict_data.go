// Package req
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package req

type SystemDictDataSearch struct {
	TypeId    int64    `json:"type_id"`
	Code      string   `json:"code"`
	Codes     string   `json:"codes"`
	Status    int      `json:"status"`
	Value     string   `json:"value"`
	Label     string   `json:"label"`
	CreatedAt []string `json:"created_at" dc:"created at"`
}

type SystemDictDataSave struct {
	TypeId int64  `json:"type_id" v:"required"`
	Code   string `json:"code" v:"required"`
	Value  string `json:"value" v:"required"`
	Label  string `json:"label" v:"required"`
}

type SystemDictDataUpdate struct {
	Id    int64  `json:"id" v:"required"`
	Code  string `json:"code" v:"required"`
	Value string `json:"value" v:"required"`
	Label string `json:"label" v:"required"`
}
