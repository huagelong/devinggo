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
	Id             interface{} //
	Username       interface{} //
	Password       interface{} //
	UserType       interface{} //
	Nickname       interface{} //
	Phone          interface{} //
	Email          interface{} //
	Avatar         interface{} //
	Signed         interface{} //
	Dashboard      interface{} //
	Status         interface{} //
	LoginIp        interface{} //
	LoginTime      *gtime.Time //
	BackendSetting *gjson.Json //
	CreatedBy      interface{} //
	UpdatedBy      interface{} //
	CreatedAt      *gtime.Time //
	UpdatedAt      *gtime.Time //
	DeletedAt      *gtime.Time //
	Remark         interface{} //
}
