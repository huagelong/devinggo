// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/os/gtime"
)

// CodeGenFields is the golang structure for table code_gen_fields.
type CodeGenFields struct {
	Id            int64       `json:"id"            orm:"id"             description:""`               //
	TableId       int64       `json:"tableId"       orm:"table_id"       description:"所属表ID"`          // 所属表ID
	ColumnName    string      `json:"columnName"    orm:"column_name"    description:"字段名称"`           // 字段名称
	ColumnComment string      `json:"columnComment" orm:"column_comment" description:"字段描述"`           // 字段描述
	ColumnType    string      `json:"columnType"    orm:"column_type"    description:"物理类型"`           // 物理类型
	DataType      string      `json:"dataType"      orm:"data_type"      description:"数据类型"`           // 数据类型
	IsNullable    string      `json:"isNullable"    orm:"is_nullable"    description:"是否可空"`           // 是否可空
	Sort          int         `json:"sort"          orm:"sort"           description:"排序"`             // 排序
	IsRequired    int         `json:"isRequired"    orm:"is_required"    description:"必填: 1=是, 2=否"`   // 必填: 1=是, 2=否
	IsInsert      int         `json:"isInsert"      orm:"is_insert"      description:"新增: 1=是, 2=否"`   // 新增: 1=是, 2=否
	IsEdit        int         `json:"isEdit"        orm:"is_edit"        description:"编辑: 1=是, 2=否"`   // 编辑: 1=是, 2=否
	IsList        int         `json:"isList"        orm:"is_list"        description:"列表显示: 1=是, 2=否"` // 列表显示: 1=是, 2=否
	IsQuery       int         `json:"isQuery"       orm:"is_query"       description:"查询: 1=是, 2=否"`   // 查询: 1=是, 2=否
	IsSort        int         `json:"isSort"        orm:"is_sort"        description:"排序: 1=是, 2=否"`   // 排序: 1=是, 2=否
	QueryType     string      `json:"queryType"     orm:"query_type"     description:"查询方式"`           // 查询方式
	ViewType      string      `json:"viewType"      orm:"view_type"      description:"页面控件"`           // 页面控件
	DictType      string      `json:"dictType"      orm:"dict_type"      description:"数据字典"`           // 数据字典
	AllowRoles    string      `json:"allowRoles"    orm:"allow_roles"    description:"允许角色"`           // 允许角色
	Options       *gjson.Json `json:"options"       orm:"options"        description:"组件扩展配置"`         // 组件扩展配置
	CreatedAt     *gtime.Time `json:"createdAt"     orm:"created_at"     description:"创建时间"`           // 创建时间
	UpdatedAt     *gtime.Time `json:"updatedAt"     orm:"updated_at"     description:"更新时间"`           // 更新时间
}
