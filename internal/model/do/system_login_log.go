// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SystemLoginLog is the golang structure of table system_login_log for DAO operations like Where/Data.
type SystemLoginLog struct {
	g.Meta     `orm:"table:system_login_log, do:true"`
	Id         any         //
	Username   any         //
	Ip         any         //
	IpLocation any         //
	Os         any         //
	Browser    any         //
	Status     any         //
	Message    any         //
	LoginTime  *gtime.Time //
	Remark     any         //
}
