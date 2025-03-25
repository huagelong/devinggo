// Package req
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package req

type SystemMenuSearch struct {
	Status    int      `json:"status"`
	Level     string   `json:"level"`
	Name      string   `json:"name"`
	Recycle   bool     `json:"recycle"`
	NoButtons bool     `json:"no_buttons"`
	CreatedAt []string `json:"created_at" dc:"created at"`
}

type SystemMenuSave struct {
	Id        int64  `json:"id"`
	Level     string `json:"level"`
	ParentId  int64  `json:"parent_id"`
	Type      string `json:"type"`
	Sort      int    `json:"sort"`
	IsHidden  int    `json:"is_hidden"`
	Status    int    `json:"status"`
	Restful   string `json:"restful"`
	Name      string `json:"name"`
	Icon      string `json:"icon"`
	Code      string `json:"code"`
	Route     string `json:"route"`
	Component string `json:"component"`
	Redirect  string `json:"redirect"`
	Remark    string `json:"remark"`
}
