// Package res
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package res

import "github.com/gogf/gf/v2/os/gtime"

type SystemMenuTree struct {
	SystemMenuItem
	Children []*SystemMenuTree `json:"children" description:"子菜单"` // 子菜单
}

type SystemMenuItem struct {
	Id        int64       `json:"id"                 description:"主键"`                         // 主键
	ParentId  int64       `json:"parent_id"    description:"父ID"`                              // 父ID
	Level     string      `json:"level"           description:"组级集合"`                          // 组级集合
	Name      string      `json:"name"             description:"菜单名称"`                         // 菜单名称
	Code      string      `json:"code"             description:"菜单标识代码"`                       // 菜单标识代码
	Icon      string      `json:"icon"             description:"菜单图标"`                         // 菜单图标
	Route     string      `json:"route"           description:"路由地址"`                          // 路由地址
	Component string      `json:"component"   description:"组件路径"`                              // 组件路径
	Redirect  string      `json:"redirect"     description:"跳转地址"`                             // 跳转地址
	IsHidden  int         `json:"is_hidden"    description:"是否隐藏 (1是 2否)"`                     // 是否隐藏 (1是 2否)
	Type      string      `json:"type"             description:"菜单类型, (M菜单 B按钮 L链接 I iframe)"` // 菜单类型, (M菜单 B按钮 L链接 I iframe)
	Status    int         `json:"status"         description:"状态 (1正常 2停用)"`                   // 状态 (1正常 2停用)
	Sort      int         `json:"sort"             description:"排序"`                           // 排序
	CreatedBy int64       `json:"created_by"  description:"创建者"`                               // 创建者
	UpdatedBy int64       `json:"updated_by"  description:"更新者"`                               // 更新者
	CreatedAt *gtime.Time `json:"created_at"  description:""`                                  //
	UpdatedAt *gtime.Time `json:"updated_at"  description:""`                                  //
	DeletedAt *gtime.Time `json:"deleted_at"  description:"删除时间"`                              // 删除时间
	Remark    string      `json:"remark"         description:"备注"`                             // 备注
}

type SystemDeptSelectTree struct {
	SystemDeptSelectItem
	Children []*SystemDeptSelectTree `json:"children"`
}

type SystemDeptSelectItem struct {
	Id       int64  `json:"id"`
	ParentId int64  `json:"parent_id"`
	Value    int64  `json:"value"`
	Label    string `json:"label"`
}
