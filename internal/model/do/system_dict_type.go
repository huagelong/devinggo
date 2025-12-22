// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SystemDictType is the golang structure of table system_dict_type for DAO operations like Where/Data.
type SystemDictType struct {
	g.Meta    `orm:"table:system_dict_type, do:true"`
	Id        any         // 主键
	Name      any         // 字典名称
	Code      any         // 字典标示
	Status    any         // 状态 (1正常 2停用)
	CreatedBy any         // 创建者
	UpdatedBy any         // 更新者
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
	DeletedAt *gtime.Time // 删除时间
	Remark    any         // 备注
}
