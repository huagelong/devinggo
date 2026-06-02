// Package req
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package req

type SystemOperLogSearch struct {
	Ip          string   `json:"ip"`
	ServiceName string   `json:"service_name"`
	Username    string   `json:"username"`
	Status      int      `json:"status"`
	CreatedAt   []string `json:"created_at"`
}
