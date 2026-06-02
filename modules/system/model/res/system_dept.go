// Package res
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package res

import "github.com/gogf/gf/v2/os/gtime"

type SystemDeptTree struct {
	Id       int64             `json:"id"`
	ParentId int64             `json:"parent_id"`
	Value    int64             `json:"value"`
	Label    string            `json:"label"`
	Leader   string            `json:"leader"   `
	Children []*SystemDeptTree `json:"children"`
}

type SystemListDeptTree struct {
	Id        int64                 `json:"id"`
	ParentId  int64                 `json:"parent_id"`
	Name      string                `json:"name"`
	Phone     string                `json:"phone"`
	Leader    string                `json:"leader"`
	Sort      int                   `json:"sort"`
	Status    int                   `json:"status"`
	CreatedAt *gtime.Time           `json:"created_at"`
	Children  []*SystemListDeptTree `json:"children"`
}

type SystemDeptItem struct {
	Id       int64  `json:"id"`
	ParentId int64  `json:"parent_id"`
	Value    int64  `json:"value"`
	Label    string `json:"label"`
}

type SystemDept struct {
	Id        int64       `json:"id"                description:"主键"`       // 主键
	ParentId  int64       `json:"parent_id"   description:"父ID"`            // 父ID
	Level     string      `json:"level"          description:"组级集合"`        // 组级集合
	Name      string      `json:"name"            description:"部门名称"`       // 部门名称
	Leader    string      `json:"leader"        description:"负责人"`          // 负责人
	Phone     string      `json:"phone"          description:"联系电话"`        // 联系电话
	Status    int         `json:"status"        description:"状态 (1正常 2停用)"` // 状态 (1正常 2停用)
	Sort      int         `json:"sort"            description:"排序"`         // 排序
	CreatedBy int64       `json:"created_by" description:"创建者"`             // 创建者
	UpdatedBy int64       `json:"updated_by" description:"更新者"`             // 更新者
	CreatedAt *gtime.Time `json:"created_at" description:"创建时间"`            // 创建时间
	UpdatedAt *gtime.Time `json:"updated_at" description:"更新时间"`            // 更新时间
	Remark    string      `json:"remark"        description:"备注"`           // 备注
}
