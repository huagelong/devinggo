// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SystemLoginLog is the golang structure for table system_login_log.
type SystemLoginLog struct {
	Id         int64       `json:"id"         orm:"id"          description:""` //
	Username   string      `json:"username"   orm:"username"    description:""` //
	Ip         string      `json:"ip"         orm:"ip"          description:""` //
	IpLocation string      `json:"ipLocation" orm:"ip_location" description:""` //
	Os         string      `json:"os"         orm:"os"          description:""` //
	Browser    string      `json:"browser"    orm:"browser"     description:""` //
	Status     int         `json:"status"     orm:"status"      description:""` //
	Message    string      `json:"message"    orm:"message"     description:""` //
	LoginTime  *gtime.Time `json:"loginTime"  orm:"login_time"  description:""` //
	Remark     string      `json:"remark"     orm:"remark"      description:""` //
}
