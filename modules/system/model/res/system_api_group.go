// Package res
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package res

import "github.com/gogf/gf/v2/os/gtime"

type SystemApiGroup struct {
	Id         int64        `json:"id"                 description:"主键"`       // 主键
	Name       string       `json:"name"             description:"接口组名称"`      // 接口组名称
	GroupId    int64        `json:"group_id"         description:"接口组ID"`      // 接口组ID
	AccessName string       `json:"access_name"  description:"接口访问名称"`         // 接口访问名称
	Status     int          `json:"status"         description:"状态 (1正常 2停用)"` // 状态 (1正常 2停用)
	CreatedBy  int64        `json:"created_by"  description:"创建者"`             // 创建者
	UpdatedBy  int64        `json:"updated_by"  description:"更新者"`             // 更新者
	CreatedAt  *gtime.Time  `json:"created_at"  description:"创建时间"`            // 创建时间
	UpdatedAt  *gtime.Time  `json:"updated_at"  description:"更新时间"`            // 更新时间
	Remark     string       `json:"remark"         description:"备注"`           // 备注
	Apis       []*SystemApi `json:"apis"            description:"接口列表"`        // 接口列表
}
