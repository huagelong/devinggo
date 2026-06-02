// Package req
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package req

type SystemApiLogSearch struct {
	ApiName    string   `json:"api_name"`
	Ip         string   `json:"ip"`
	AccessName string   `json:"access_name"`
	AccessTime []string `json:"access_time"`
}
