// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SettingGenerateColumns is the golang structure of table setting_generate_columns for DAO operations like Where/Data.
type SettingGenerateColumns struct {
	g.Meta        `orm:"table:setting_generate_columns, do:true"`
	Id            any         // 主键
	TableId       any         // 所属表ID
	ColumnName    any         // 字段名称
	ColumnComment any         // 字段注释
	ColumnType    any         // 字段类型
	IsPk          any         // 1 非主键 2 主键
	IsRequired    any         // 1 非必填 2 必填
	IsInsert      any         // 1 非插入字段 2 插入字段
	IsEdit        any         // 1 非编辑字段 2 编辑字段
	IsList        any         // 1 非列表显示字段 2 列表显示字段
	IsQuery       any         // 1 非查询字段 2 查询字段
	IsSort        any         // 1 不排序 2 排序字段
	QueryType     any         // 查询方式 eq 等于, neq 不等于, gt 大于, lt 小于, like 范围
	ViewType      any         // 页面控件，text, textarea, password, select, checkbox, radio, date, upload, ma-upload（封装的上传控件）
	DictType      any         // 字典类型
	AllowRoles    any         // 允许查看该字段的角色
	Options       any         // 字段其他设置
	Extra         any         // 字段扩展信息
	Sort          any         // 排序
	CreatedBy     any         // 创建者
	UpdatedBy     any         // 更新者
	CreatedAt     *gtime.Time // 创建时间
	UpdatedAt     *gtime.Time // 更新时间
	Remark        any         // 备注
}
