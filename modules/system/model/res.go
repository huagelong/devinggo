// Package res
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package model

type UserRelate struct {
	Id       int64  `json:"id"                 description:"user id"`
	Nickname string `json:"nickname"           description:"nickname"`
	Username string `json:"username"           description:"account"`
	Avatar   string `json:"avatar"             description:"avatar"`
}

type Dict struct {
	Value string `json:"value" description:"value"`
	Label string `json:"label" description:"label"`
}
