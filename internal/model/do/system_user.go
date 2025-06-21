// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SystemUser is the golang structure of table system_user for DAO operations like Where/Data.
type SystemUser struct {
	g.Meta         `orm:"table:system_user, do:true"`
	Id             interface{} // 用户ID，主键
	Username       interface{} // 用户名
	Password       interface{} // 密码
	UserType       interface{} // 用户类型：(100系统用户)
	Nickname       interface{} // 用户昵称
	Phone          interface{} // 手机
	Email          interface{} // 用户邮箱
	Avatar         interface{} // 用户头像
	Signed         interface{} // 个人签名
	Dashboard      interface{} // 后台首页类型
	Status         interface{} // 状态 (1正常 2停用)
	LoginIp        interface{} // 最后登陆IP
	LoginTime      *gtime.Time // 最后登陆时间
	BackendSetting interface{} // 后台设置数据
	CreatedBy      interface{} // 创建者
	UpdatedBy      interface{} // 更新者
	CreatedAt      *gtime.Time // 创建时间
	UpdatedAt      *gtime.Time // 更新时间
	DeletedAt      *gtime.Time // 删除时间
	Remark         interface{} // 备注
}
