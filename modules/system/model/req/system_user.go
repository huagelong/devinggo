// Package req
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package req

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/os/gtime"
)

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
}

type SystemUserSearch struct {
	DeptId           int64    `json:"dept_id" description:"部门ID"`                         // 部门ID
	Username         string   `json:"username"             description:"用户名"`             // 用户名
	Nickname         string   `json:"nickname"            description:"用户昵称"`             // 用户昵称
	Phone            string   `json:"phone"                 description:"手机"`             // 手机
	Email            string   `json:"email"                  description:"用户邮箱"`          // 用户邮箱
	Status           int      `json:"status"                  description:"状态 (1正常 2停用)"` // 状态 (1正常 2停用)
	CreatedAt        []string `json:"created_at" dc:"created at"`
	UserIds          []int64  `json:"user_ids" description:"用户ID列表"`
	ShowDept         bool     `json:"show_dept" description:"是否显示部门"`
	FilterSuperAdmin bool     `json:"filter_super_admin" description:"是否过滤超级管理员"`
	RoleId           int64    `json:"role_id" description:"角色ID"`
	PostId           int64    `json:"post_id" description:"岗位ID"`
}

type SystemUserSave struct {
	Username string  `json:"username"     v:"required|max-length:20"        description:"用户名"`             // 用户名
	UserType string  `json:"user_type"    v:"required"                description:"用户类型：(100系统用户)"`        // 用户类型：(100系统用户)
	Nickname string  `json:"nickname"            description:"用户昵称"`                                       // 用户昵称
	Password string  `json:"password"     v:"required|min-length:6|max-length:20"        description:"密码"` // 密码
	Phone    string  `json:"phone"                 description:"手机"`                                       // 手机
	Email    string  `json:"email"                  description:"邮箱"`                                      // 邮箱
	Status   int     `json:"status"                  description:"状态 (1正常 2停用)"`                           // 状态 (1正常 2停用)
	DeptIds  []int64 `json:"dept_ids" description:"部门ID列表"`
	RoleIds  []int64 `json:"role_ids" description:"角色ID列表"`
	PostIds  []int64 `json:"post_ids" description:"岗位ID列表"`
	Remark   string  `json:"remark"   v:"max-length:255"             description:"备注"` // 备注
}

type SystemUserUpdate struct {
	Id       int64   `json:"id" v:"required" description:"用户ID"`
	Username string  `json:"username"             description:"用户名"`                   // 用户名
	UserType string  `json:"user_type"                   description:"用户类型：(100系统用户)"` // 用户类型：(100系统用户)
	Nickname string  `json:"nickname"            description:"用户昵称"`                   // 用户昵称
	Phone    string  `json:"phone"                 description:"手机"`                   // 手机
	Email    string  `json:"email"                  description:"邮箱"`                  // 邮箱
	Status   int     `json:"status"               description:"状态 (1正常 2停用)"`          // 状态 (1正常 2停用)
	DeptIds  []int64 `json:"dept_ids" v:"required" description:"部门ID列表"`
	RoleIds  []int64 `json:"role_ids" v:"required" description:"角色ID列表"`
	PostIds  []int64 `json:"post_ids" v:"required" description:"岗位ID列表"`
	Remark   string  `json:"remark"   v:"max-length:255"             description:"备注"` // 备注
}
