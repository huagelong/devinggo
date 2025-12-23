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
	Id             any         // 用户ID，主键
	Username       any         // 用户名
	Password       any         // 密码
	UserType       any         // 用户类型：(100系统用户)
	Nickname       any         // 用户昵称
	Phone          any         // 手机
	Email          any         // 用户邮箱
	Avatar         any         // 用户头像
	Signed         any         // 个人签名
	Dashboard      any         // 后台首页类型
	Status         any         // 状态 (1正常 2停用)
	LoginIp        any         // 最后登陆IP
	LoginTime      *gtime.Time // 最后登陆时间
	BackendSetting any         // 后台设置数据
	CreatedBy      any         // 创建者
	UpdatedBy      any         // 更新者
	CreatedAt      *gtime.Time // 创建时间
	UpdatedAt      *gtime.Time // 更新时间
	DeletedAt      *gtime.Time // 删除时间
	Remark         any         // 备注
}
