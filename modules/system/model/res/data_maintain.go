// Package res
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package res

import "github.com/gogf/gf/v2/os/gtime"

type DataMaintain struct {
	Name       string      `json:"name"`
	Collation  string      `json:"collation"`
	Comment    string      `json:"comment"`
	Engine     string      `json:"engine"`
	CreateTime *gtime.Time `json:"create_time"` // 创建时间
	Rows       int64       `json:"rows"`        // 行数
	DataLength int64       `json:"data_length"`
	UpdateTime string      `json:"update_time"`
}

type DataMaintainColumn struct {
	Field        string `json:"field"`
	Type         string `json:"type"`
	Nullable     bool   `json:"nullable"`
	Key          string `json:"key"`
	DefaultValue string `json:"default_value"`
	Comment      string `json:"comment"`
}
