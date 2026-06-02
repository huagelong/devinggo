// Package res
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package res

import "github.com/gogf/gf/v2/os/gtime"

type SystemDictData struct {
	Id    int64  `json:"id"`
	Key   string `json:"key"`
	Title string `json:"title"`
	Code  string `json:"code"`
}

type SystemDictDataFull struct {
	Id        int64       `json:"id"                 description:"主键"`       // 主键
	TypeId    int64       `json:"typeId"        description:"字典类型ID"`        // 字典类型ID
	Label     string      `json:"label"           description:"字典标签"`        // 字典标签
	Value     string      `json:"value"           description:"字典值"`         // 字典值
	Code      string      `json:"code"             description:"字典标示"`       // 字典标示
	Sort      int         `json:"sort"             description:"排序"`         // 排序
	Status    int         `json:"status"         description:"状态 (1正常 2停用)"` // 状态 (1正常 2停用)
	CreatedBy int64       `json:"created_by"  description:"创建者"`             // 创建者
	UpdatedBy int64       `json:"updated_by"  description:"更新者"`             // 更新者
	CreatedAt *gtime.Time `json:"created_at"  description:""`                //
	UpdatedAt *gtime.Time `json:"updated_at"  description:""`                //
	Remark    string      `json:"remark"         description:"备注"`           // 备注
}
