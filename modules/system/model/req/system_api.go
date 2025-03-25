// Package req
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package req

type SystemApiSearch struct {
	GroupId    int64  `json:"group_id"`
	Name       string `json:"name"`
	AccessName string `json:"access_name"`
	Status     int    `json:"status"`
}

type SystemApiSave struct {
	GroupId     int64  `json:"group_id"     v:"required"     description:"接口组ID"`                   // 接口组ID
	Name        string `json:"name"      v:"required"             description:"接口名称"`               // 接口名称
	AccessName  string `json:"access_name"   v:"required"   description:"接口访问名称"`                   // 接口访问名称
	AuthMode    int    `json:"auth_mode"     v:"required"     description:"认证模式 (1简易 2复杂)"`         // 认证模式 (1简易 2复杂)
	RequestMode string `json:"request_mode"  v:"required"   description:"请求模式 (A 所有 P POST G GET)"` // 请求模式 (A 所有 P POST G GET)
	Status      int    `json:"status"             description:"状态 (1正常 2停用)"`                       // 状态 (1正常 2停用)
	Remark      string `json:"remark"             description:"备注"`                                 // 备注
}

type SystemApiUpdate struct {
	Id          int64  `json:"id"          v:"required"     description:"接口ID"`                     // 接口ID
	GroupId     int64  `json:"group_id"      description:"接口组ID"`                                   // 接口组ID
	Name        string `json:"name"      v:"required"             description:"接口名称"`               // 接口名称
	AccessName  string `json:"access_name"   v:"required"   description:"接口访问名称"`                   // 接口访问名称
	AuthMode    int    `json:"auth_mode"     v:"required"     description:"认证模式 (1简易 2复杂)"`         // 认证模式 (1简易 2复杂)
	RequestMode string `json:"request_mode"  v:"required"   description:"请求模式 (A 所有 P POST G GET)"` // 请求模式 (A 所有 P POST G GET)
	Status      int    `json:"status"             description:"状态 (1正常 2停用)"`                       // 状态 (1正常 2停用)
	Remark      string `json:"remark"             description:"备注"`                                 // 备注
}
