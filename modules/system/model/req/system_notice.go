// Package req
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package req

type SystemNoticeSearch struct {
	Title     string   `json:"title"`
	Type      int      `json:"type"`
	CreatedAt []string `json:"created_at" dc:"created at"`
}

type SystemNoticeSave struct {
	Title   string  `json:"title" v:"required"`
	Type    int     `json:"type" v:"required"`
	Content string  `json:"content" v:"required"`
	Remark  string  `json:"remark"`
	Users   []int64 `json:"users"`
}

type SystemNoticeUpdate struct {
	Id      int64  `json:"id" v:"required"`
	Title   string `json:"title" v:"required"`
	Type    int    `json:"type" v:"required"`
	Content string `json:"content" v:"required"`
	Remark  string `json:"remark"`
}
