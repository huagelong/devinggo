// Package api
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE
package api

import "github.com/gogf/gf/v2/frame/g"

type HealthReq struct {
	g.Meta `path:"/health" method:"get" tags:"health" summary:"健康检查." `
}

type HealthRes struct {
	g.Meta `mime:"application/json"`
	Data   string `json:"data"  dc:"data" `
}
