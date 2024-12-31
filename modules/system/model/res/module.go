// Package res
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE
package res

type Modules struct {
	System Module `json:"System"`
}

type Module struct {
	Name        string `json:"name"`
	Label       string `json:"label"`
	Description string `json:"description"`
	Installed   bool   `json:"installed"`
	Enabled     bool   `json:"enabled"`
}
