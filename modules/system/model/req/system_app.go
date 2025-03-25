// Package req
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package req

type SystemAppSearch struct {
	AppName string `json:"app_name"`
	AppId   string `json:"app_id"`
	GroupId int64  `json:"group_id"`
	Status  int    `json:"status"`
}

type SystemAppSave struct {
	AppName     string `json:"app_name" v:"required"`
	AppId       string `json:"app_id"  v:"required"`
	GroupId     int64  `json:"group_id"  v:"required"`
	AppSecret   string `json:"app_secret"  v:"required"`
	Status      int    `json:"status"`
	Description string `json:"description"`
	Remark      string `json:"remark"`
}

type SystemAppUpdate struct {
	Id          int64  `json:"id" v:"required"`
	AppName     string `json:"app_name" v:"required"`
	AppId       string `json:"app_id"  v:"required"`
	GroupId     int64  `json:"group_id"  v:"required"`
	AppSecret   string `json:"app_secret"  v:"required"`
	Status      int    `json:"status"`
	Description string `json:"description"`
	Remark      string `json:"remark"`
}
