// Package api
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package api

import (
	"devinggo/modules/system/model"
	"github.com/gogf/gf/v2/frame/g"
)

type IndexReq struct {
	g.Meta `path:"/test" method:"get" tags:"test" summary:"测试." x-exceptLogin:"true" x-permission:"api:test" `
	model.ApiSign
}

type IndexRes struct {
	g.Meta `mime:"application/json"`
	Data   string `json:"data"  dc:"test data" `
}
