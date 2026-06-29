// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// CodeGenFields is the golang structure of table code_gen_fields for DAO operations like Where/Data.
type CodeGenFields struct {
	g.Meta        `orm:"table:code_gen_fields, do:true"`
	Id            any         //
	TableId       any         // 所属表ID
	ColumnName    any         // 字段名称
	ColumnComment any         // 字段描述
	ColumnType    any         // 物理类型
	DataType      any         // 数据类型
	IsNullable    any         // 是否可空
	Sort          any         // 排序
	IsRequired    any         // 必填: 1=是, 2=否
	IsInsert      any         // 新增: 1=是, 2=否
	IsEdit        any         // 编辑: 1=是, 2=否
	IsList        any         // 列表显示: 1=是, 2=否
	IsQuery       any         // 查询: 1=是, 2=否
	IsSort        any         // 排序: 1=是, 2=否
	QueryType     any         // 查询方式
	ViewType      any         // 页面控件
	DictType      any         // 数据字典
	AllowRoles    any         // 允许角色
	Options       *gjson.Json // 组件扩展配置
	CreatedAt     *gtime.Time // 创建时间
	UpdatedAt     *gtime.Time // 更新时间
}
