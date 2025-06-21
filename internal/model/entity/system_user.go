// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SystemUser is the golang structure for table system_user.
type SystemUser struct {
	Id             int64       `json:"id"             orm:"id"              description:"用户ID，主键"`        // 用户ID，主键
	Username       string      `json:"username"       orm:"username"        description:"用户名"`            // 用户名
	Password       string      `json:"password"       orm:"password"        description:"密码"`             // 密码
	UserType       string      `json:"userType"       orm:"user_type"       description:"用户类型：(100系统用户)"` // 用户类型：(100系统用户)
	Nickname       string      `json:"nickname"       orm:"nickname"        description:"用户昵称"`           // 用户昵称
	Phone          string      `json:"phone"          orm:"phone"           description:"手机"`             // 手机
	Email          string      `json:"email"          orm:"email"           description:"用户邮箱"`           // 用户邮箱
	Avatar         string      `json:"avatar"         orm:"avatar"          description:"用户头像"`           // 用户头像
	Signed         string      `json:"signed"         orm:"signed"          description:"个人签名"`           // 个人签名
	Dashboard      string      `json:"dashboard"      orm:"dashboard"       description:"后台首页类型"`         // 后台首页类型
	Status         int         `json:"status"         orm:"status"          description:"状态 (1正常 2停用)"`   // 状态 (1正常 2停用)
	LoginIp        string      `json:"loginIp"        orm:"login_ip"        description:"最后登陆IP"`         // 最后登陆IP
	LoginTime      *gtime.Time `json:"loginTime"      orm:"login_time"      description:"最后登陆时间"`         // 最后登陆时间
	BackendSetting string      `json:"backendSetting" orm:"backend_setting" description:"后台设置数据"`         // 后台设置数据
	CreatedBy      int64       `json:"createdBy"      orm:"created_by"      description:"创建者"`            // 创建者
	UpdatedBy      int64       `json:"updatedBy"      orm:"updated_by"      description:"更新者"`            // 更新者
	CreatedAt      *gtime.Time `json:"createdAt"      orm:"created_at"      description:"创建时间"`           // 创建时间
	UpdatedAt      *gtime.Time `json:"updatedAt"      orm:"updated_at"      description:"更新时间"`           // 更新时间
	DeletedAt      *gtime.Time `json:"deletedAt"      orm:"deleted_at"      description:"删除时间"`           // 删除时间
	Remark         string      `json:"remark"         orm:"remark"          description:"备注"`             // 备注
}
