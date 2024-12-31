// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// Migrations is the golang structure of table migrations for DAO operations like Where/Data.
type Migrations struct {
	g.Meta    `orm:"table:migrations, do:true"`
	Id        interface{} //
	Migration interface{} //
	Batch     interface{} //
}
