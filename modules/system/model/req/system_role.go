// Package req
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package req

type SystemRoleSearch struct {
	Name      string   `json:"name"`
	Code      string   `json:"code"`
	Status    int      `json:"status"`
	CreatedAt []string `json:"created_at" dc:"created at"`
}

type SystemRoleSave struct {
	Id        int64   `json:"id"`
	Name      string  `json:"name"`
	Code      string  `json:"code"`
	DataScope int     `json:"data_scope"`
	Status    int     `json:"status"`
	Sort      int     `json:"sort"`
	Remark    string  `json:"remark"`
	MenuIds   []int64 `json:"menu_ids"`
	DeptIds   []int64 `json:"dept_ids"`
}
