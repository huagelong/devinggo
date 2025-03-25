// Package req
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package req

type SystemDeptSearch struct {
	Status    int      `json:"status"`
	Level     string   `json:"level"`
	Name      string   `json:"name"`
	Leader    string   `json:"leader"`
	Phone     string   `json:"phone"`
	Recycle   bool     `json:"recycle"`
	CreatedAt []string `json:"created_at" dc:"created at"`
}

type SystemDeptSave struct {
	Id       int64  `json:"id"`
	Sort     int    `json:"sort"`
	Status   string `json:"status"`
	ParentId int64  `json:"parent_id"`
	Name     string `json:"name" v:"required|length:1,30#请输入部门名称|部门名称长度必须在1到30之间"`
	Leader   string `json:"leader"`
	Phone    string `json:"phone"`
	Level    string `json:"level"`
	Remark   string `json:"remark"`
}

type SystemDeptAddLeader struct {
	Id    int64        `json:"id" v:"required"`
	Users []SystemUser `json:"users" v:"required"`
}
