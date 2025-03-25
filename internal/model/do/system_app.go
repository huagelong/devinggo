// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SystemApp is the golang structure of table system_app for DAO operations like Where/Data.
type SystemApp struct {
	g.Meta      `orm:"table:system_app, do:true"`
	Id          interface{} //
	GroupId     interface{} //
	AppName     interface{} //
	AppId       interface{} //
	AppSecret   interface{} //
	Status      interface{} //
	Description interface{} //
	CreatedBy   interface{} //
	UpdatedBy   interface{} //
	CreatedAt   *gtime.Time //
	UpdatedAt   *gtime.Time //
	DeletedAt   *gtime.Time //
	Remark      interface{} //
}
