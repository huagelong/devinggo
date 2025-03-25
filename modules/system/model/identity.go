// Package model
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package model

type Identity struct {
	Id       int64   `json:"id"              description:"用户ID"`
	Username string  `json:"username"        description:"用户名"`
	AppId    string  `json:"app_id"           description:"应用ID"`
	RoleIds  []int64 `json:"roleIds"      description:"角色ID列表"`
	DeptIds  []int64 `json:"deptIds"      description:"部门ID列表"`
}

type NormalIdentity struct {
	Scene     string                 `json:"scene" description:"场景"`
	Data      map[string]interface{} `json:"data" description:"数据"`
	ExpiresAt int64                  `json:"expiresAt" description:"过期时间"`
}
