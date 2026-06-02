// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SystemUser is the golang structure of table system_user for DAO operations like Where/Data.
type SystemUser struct {
	g.Meta         `orm:"table:system_user, do:true"`
	Id             any         //
	Username       any         //
	Password       any         //
	UserType       any         //
	Nickname       any         //
	Phone          any         //
	Email          any         //
	Avatar         any         //
	Signed         any         //
	Dashboard      any         //
	Status         any         //
	LoginIp        any         //
	LoginTime      *gtime.Time //
	BackendSetting *gjson.Json //
	CreatedBy      any         //
	UpdatedBy      any         //
	CreatedAt      *gtime.Time //
	UpdatedAt      *gtime.Time //
	DeletedAt      *gtime.Time //
	Remark         any         //
}
