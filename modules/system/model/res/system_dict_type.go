// Package res
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package res

import "github.com/gogf/gf/v2/os/gtime"

type SystemDictType struct {
	Id        int64       `json:"id"                 description:"主键"`       // 主键
	Name      string      `json:"name"             description:"字典名称"`       // 字典名称
	Code      string      `json:"code"             description:"字典标示"`       // 字典标示
	Status    int         `json:"status"         description:"状态 (1正常 2停用)"` // 状态 (1正常 2停用)
	CreatedBy int64       `json:"created_by"  description:"创建者"`             // 创建者
	UpdatedBy int64       `json:"updated_by"  description:"更新者"`             // 更新者
	CreatedAt *gtime.Time `json:"created_at"  description:""`                //
	UpdatedAt *gtime.Time `json:"updated_at"  description:""`                //
	Remark    string      `json:"remark"         description:"备注"`           // 备注
}
