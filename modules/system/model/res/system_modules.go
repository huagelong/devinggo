// Package req
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package res

import "github.com/gogf/gf/v2/os/gtime"

type SystemModules struct {
	Id int64 `json:"id"  description:"ID" `

	Name string `json:"name"  description:"模块名称" `

	Label string `json:"label"  description:"模块标记" `

	Installed int `json:"installed"  description:"是否安装" `

	Status int `json:"status"  description:"状态" `

	CreatedAt *gtime.Time `json:"created_at"  description:"创建时间" `

	Description string `json:"description"  description:"描述" `
}

type SystemModulesExcel struct {
	Id int64 `json:"id"  description:"ID"  `

	Name string `json:"name"  v:"required"  description:"模块名称"  `

	Label string `json:"label"  v:"required"  description:"模块标记"  `

	Installed int `json:"installed"  v:"required"  description:"是否安装"  `

	Status int `json:"status"  v:"required"  description:"状态"  `

	Description string `json:"description"  v:"required"  description:"描述"  `
}
