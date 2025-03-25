// Package res
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package res

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/os/gtime"
)

type SystemUserInfo struct {
	User    SystemUser `json:"user" description:"用户信息"`    // 用户信息
	Roles   []string   `json:"roles" description:"用户角色"`   // 用户角色
	Codes   []string   `json:"codes" description:"导航菜单编码"` // 导航菜单编码
	Routers []*Router  `json:"routers" description:"菜单路由"` // 路由
}

type Router struct {
	SystemMenu
	Children []*Router `json:"children" dc:"Children node"`
}

type SystemMenu struct {
	Id        int64  `json:"id"                description:"主键"`   // 主键
	ParentId  int64  `json:"parent_id"    description:"父ID"`       // 父ID
	Name      string `json:"name"            description:"菜单标识代码"` // 菜单标识代码
	Path      string `json:"path"        description:"路由地址"`       // 路由地址
	Component string `json:"component"   description:"组件路径"`       // 组件路径
	Redirect  string `json:"redirect"   description:"跳转地址"`        // 跳转地址
	Meta      Meta   `json:"meta"        description:"meta"`       // meta
}

type Meta struct {
	Title            string `json:"title" dc:"name"`
	Icon             string `json:"icon" dc:"icon"`
	Type             string `json:"type" dc:"菜单类型, (M菜单 B按钮 L链接 I iframe)"`
	Hidden           bool   `json:"hidden" dc:"hidden"`
	HiddenBreadcrumb bool   `json:"hiddenBreadcrumb" dc:"hiddenBreadcrumb"`
}

type SystemUserSimple struct {
	Id       int64  `json:"id"                   description:"用户ID，主键"` // 用户ID，主键
	Username string `json:"username"             description:"用户名"`     // 用户名
}

type SystemUser struct {
	Id             int64       `json:"id"                   description:"用户ID，主键"`         // 用户ID，主键
	Username       string      `json:"username"             description:"用户名"`             // 用户名
	UserType       string      `json:"user_type"           description:"用户类型：(100系统用户)"`   // 用户类型：(100系统用户)
	Nickname       string      `json:"nickname"            description:"用户昵称"`             // 用户昵称
	Phone          string      `json:"phone"                 description:"手机"`             // 手机
	Email          string      `json:"email"                  description:"用户邮箱"`          // 用户邮箱
	Avatar         string      `json:"avatar"                description:"用户头像"`           // 用户头像
	Signed         string      `json:"signed"               description:"个人签名"`            // 个人签名
	Dashboard      string      `json:"dashboard"            description:"后台首页类型"`          // 后台首页类型
	Status         int         `json:"status"                  description:"状态 (1正常 2停用)"` // 状态 (1正常 2停用)
	LoginIp        string      `json:"login_ip"              description:"最后登陆IP"`         // 最后登陆IP
	LoginTime      *gtime.Time `json:"login_time"         description:"最后登陆时间"`            // 最后登陆时间
	BackendSetting *gjson.Json `json:"backend_setting"     description:"后台设置数据"`           // 后台设置数据
	CreatedBy      int64       `json:"created_by"           description:"创建者"`             // 创建者
	UpdatedBy      int64       `json:"updated_by"           description:"更新者"`             // 更新者
	CreatedAt      *gtime.Time `json:"created_at"          description:"创建时间"`             // 创建时间
	UpdatedAt      *gtime.Time `json:"updated_at"            description:"更新时间"`           // 更新时间
	Remark         string      `json:"remark"                description:"备注"`             // 备注
	AppId          string      `json:"app_id"               description:"应用ID"`            // 应用ID
	//CreatedBySumma *bind_hook.UserSumma `json:"createdBySumma" dc:"创建者摘要信息"`
}

type SystemUserFullInfo struct {
	SystemUser
	RoleList []*SystemRole `json:"roleList" dc:"角色列表"` // 角色列表
	PostList []*SystemPost `json:"postList" dc:"岗位列表"` // 岗位列表
	DeptList []*SystemDept `json:"deptList" dc:"部门列表"` // 部门列表
}

type SystemUserExport struct {
	Username string `json:"username"       excelName:"用户名" excelIndex:"0"   `                                                              // 用户名
	Nickname string `json:"nickname"       excelName:"用户昵称" excelIndex:"1"  `                                                              // 用户昵称 	// 密码
	Phone    string `json:"phone"          excelName:"手机" excelIndex:"2" `                                                                 // 手机
	Status   int    `json:"status"   excelName:"状态" excelIndex:"3" toDataFormat:"ToDataStatusFormat"  toExcelFormat:"ToExcelStatusFormat"` // 状态 (1正常 2停用)
}

func (n SystemUserExport) ToExcelStatusFormat() string {
	if n.Status == 1 {
		return "正常"
	}
	if n.Status == 2 {
		return "停用"
	}
	return "待定"
}

func (n SystemUserExport) ToDataStatusFormat(status string) int {
	if status == "正常" {
		return 1
	}
	if status == "停用" {
		return 2
	}
	return 0
}

type SystemUserApp struct {
	UserId int64  `json:"user_id" dc:"用户ID"` // 用户ID
	AppId  string `json:"app_id" dc:"应用ID"`  // 应用ID
}
