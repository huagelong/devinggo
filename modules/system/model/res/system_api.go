// Package res
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package res

import "github.com/gogf/gf/v2/os/gtime"

type SystemApi struct {
	Id          int64       `json:"id"                     description:"主键"`              // 主键
	GroupId     int64       `json:"group_id"          description:"接口组ID"`                // 接口组ID
	Name        string      `json:"name"                 description:"接口名称"`              // 接口名称
	AccessName  string      `json:"access_name"    description:"接口访问名称"`                  // 接口访问名称
	AuthMode    int         `json:"auth_mode"        description:"认证模式 (1简易 2复杂)"`        // 认证模式 (1简易 2复杂)
	RequestMode string      `json:"request_mode"  description:"请求模式 (A 所有 P POST G GET)"` // 请求模式 (A 所有 P POST G GET)
	Description string      `json:"description"   description:"接口说明介绍"`                   // 接口说明介绍
	Status      int         `json:"status"             description:"状态 (1正常 2停用)"`        // 状态 (1正常 2停用)
	CreatedBy   int64       `json:"created_by"      description:"创建者"`                    // 创建者
	UpdatedBy   int64       `json:"updated_by"      description:"更新者"`                    // 更新者
	CreatedAt   *gtime.Time `json:"created_at"      description:"创建时间"`                   // 创建时间
	UpdatedAt   *gtime.Time `json:"updated_at"      description:"更新时间"`                   // 更新时间
	Remark      string      `json:"remark"             description:"备注"`                  // 备注
}
