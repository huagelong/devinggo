// Package req
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package req

type SystemAppGroupSearch struct {
	Name      string   `json:"name"`
	Status    int      `json:"status"`
	CreatedAt []string `json:"created_at" dc:"created at"`
}

type SystemAppGroupSave struct {
	Name   string `json:"name" v:"required|length:1,32#请输入分组名称|分组名称长度必须在1到32之间"`
	Status int    `json:"status"`
	Remark string `json:"remark"`
}

type SystemAppGroupUpdate struct {
	Id     int64  `json:"id" v:"required"`
	Name   string `json:"name" v:"required|length:1,32#请输入分组名称|分组名称长度必须在1到32之间"`
	Status int    `json:"status"`
	Remark string `json:"remark"`
}
