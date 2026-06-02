// Package res
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE
package res

import "github.com/gogf/gf/v2/os/gtime"

type DataMaintain struct {
	Name       string      `json:"name"`
	Collation  string      `json:"collation"`
	Comment    string      `json:"comment"`
	Engine     string      `json:"engine"`
	CreateTime *gtime.Time `json:"create_time" ` // 创建时间
	Rows       int64       `json:"rows"`         // 行数
}
