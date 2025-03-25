// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SystemApi is the golang structure of table system_api for DAO operations like Where/Data.
type SystemApi struct {
	g.Meta      `orm:"table:system_api, do:true"`
	Id          interface{} //
	GroupId     interface{} //
	Name        interface{} //
	AccessName  interface{} //
	AuthMode    interface{} //
	RequestMode interface{} //
	Status      interface{} //
	CreatedBy   interface{} //
	UpdatedBy   interface{} //
	CreatedAt   *gtime.Time //
	UpdatedAt   *gtime.Time //
	DeletedAt   *gtime.Time //
	Remark      interface{} //
}
