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
	Id        interface{} // 主键
	TypeId    interface{} // 字典类型ID
	Label     interface{} // 字典标签
	Value     interface{} // 字典值
	Code      interface{} // 字典标示
	Sort      interface{} // 排序
	Status    interface{} // 状态 (1正常 2停用)
	CreatedBy interface{} // 创建者
	UpdatedBy interface{} // 更新者
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
	DeletedAt *gtime.Time // 删除时间
	Remark    interface{} // 备注
}
