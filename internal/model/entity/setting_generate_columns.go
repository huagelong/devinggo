// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SettingGenerateColumns is the golang structure for table setting_generate_columns.
type SettingGenerateColumns struct {
	Id            int64       `json:"id"            orm:"id"             description:""` //
	TableId       int64       `json:"tableId"       orm:"table_id"       description:""` //
	ColumnName    string      `json:"columnName"    orm:"column_name"    description:""` //
	ColumnComment string      `json:"columnComment" orm:"column_comment" description:""` //
	ColumnType    string      `json:"columnType"    orm:"column_type"    description:""` //
	IsPk          int         `json:"isPk"          orm:"is_pk"          description:""` //
	IsRequired    int         `json:"isRequired"    orm:"is_required"    description:""` //
	IsInsert      int         `json:"isInsert"      orm:"is_insert"      description:""` //
	IsEdit        int         `json:"isEdit"        orm:"is_edit"        description:""` //
	IsList        int         `json:"isList"        orm:"is_list"        description:""` //
	IsQuery       int         `json:"isQuery"       orm:"is_query"       description:""` //
	IsSort        int         `json:"isSort"        orm:"is_sort"        description:""` //
	QueryType     string      `json:"queryType"     orm:"query_type"     description:""` //
	ViewType      string      `json:"viewType"      orm:"view_type"      description:""` //
	DictType      string      `json:"dictType"      orm:"dict_type"      description:""` //
	AllowRoles    string      `json:"allowRoles"    orm:"allow_roles"    description:""` //
	Options       string      `json:"options"       orm:"options"        description:""` //
	Extra         string      `json:"extra"         orm:"extra"          description:""` //
	Sort          int         `json:"sort"          orm:"sort"           description:""` //
	CreatedBy     int64       `json:"createdBy"     orm:"created_by"     description:""` //
	UpdatedBy     int64       `json:"updatedBy"     orm:"updated_by"     description:""` //
	CreatedAt     *gtime.Time `json:"createdAt"     orm:"created_at"     description:""` //
	UpdatedAt     *gtime.Time `json:"updatedAt"     orm:"updated_at"     description:""` //
	Remark        string      `json:"remark"        orm:"remark"         description:""` //
}
