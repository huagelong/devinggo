// Package req
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package req

type SystemLoginLogSearch struct {
	Ip        string   `json:"ip"`
	Username  string   `json:"username"`
	Status    int      `json:"status"`
	LoginTime []string `json:"login_time"`
}
