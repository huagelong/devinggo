// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SystemDictData is the golang structure of table system_dict_data for DAO operations like Where/Data.
type SystemDictData struct {
	g.Meta    `orm:"table:system_dict_data, do:true"`
	Id        any         // 主键
	TypeId    any         // 字典类型ID
	Label     any         // 字典标签
	Value     any         // 字典值
	Code      any         // 字典标示
	Sort      any         // 排序
	Status    any         // 状态 (1正常 2停用)
	CreatedBy any         // 创建者
	UpdatedBy any         // 更新者
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
	DeletedAt *gtime.Time // 删除时间
	Remark    any         // 备注
}
