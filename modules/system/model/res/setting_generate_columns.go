// Package req
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE
package res

import "github.com/gogf/gf/v2/os/gtime"

type SettingGenerateColumns struct {
	Id            int64       `json:"id"                         description:"主键"`                                                                                 // 主键
	TableId       int64       `json:"table_id"              description:"所属表ID"`                                                                                   // 所属表ID
	ColumnName    string      `json:"column_name"        description:"字段名称"`                                                                                       // 字段名称
	ColumnComment string      `json:"column_comment"  description:"字段注释"`                                                                                          // 字段注释
	ColumnType    string      `json:"column_type"        description:"字段类型"`                                                                                       // 字段类型
	IsPk          int         `json:"is_pk"                    description:"1 非主键 2 主键"`                                                                           // 1 非主键 2 主键
	IsRequired    int         `json:"is_required"        description:"1 非必填 2 必填"`                                                                                 // 1 非必填 2 必填
	IsInsert      int         `json:"is_insert"            description:"1 非插入字段 2 插入字段"`                                                                           // 1 非插入字段 2 插入字段
	IsEdit        int         `json:"is_edit"                description:"1 非编辑字段 2 编辑字段"`                                                                         // 1 非编辑字段 2 编辑字段
	IsList        int         `json:"is_list"                description:"1 非列表显示字段 2 列表显示字段"`                                                                     // 1 非列表显示字段 2 列表显示字段
	IsQuery       int         `json:"is_query"              description:"1 非查询字段 2 查询字段"`                                                                          // 1 非查询字段 2 查询字段
	IsSort        int         `json:"is_sort"                description:"1 不排序 2 排序字段"`                                                                           // 1 不排序 2 排序字段
	QueryType     string      `json:"query_type"          description:"查询方式 eq 等于, neq 不等于, gt 大于, lt 小于, like 范围"`                                                // 查询方式 eq 等于, neq 不等于, gt 大于, lt 小于, like 范围
	ViewType      string      `json:"view_type"            description:"页面控件，text, textarea, password, select, checkbox, radio, date, upload, ma-upload（封装的上传控件）"` // 页面控件，text, textarea, password, select, checkbox, radio, date, upload, ma-upload（封装的上传控件）
	DictType      string      `json:"dict_type"            description:"字典类型"`                                                                                     // 字典类型
	AllowRoles    string      `json:"allow_roles"        description:"允许查看该字段的角色"`                                                                                 // 允许查看该字段的角色
	Options       string      `json:"options"               description:"字段其他设置"`                                                                                  // 字段其他设置
	Extra         string      `json:"extra"                   description:"字段扩展信息"`                                                                                // 字段扩展信息
	Sort          int         `json:"sort"                     description:"排序"`                                                                                   // 排序
	CreatedBy     int64       `json:"created_by"          description:"创建者"`                                                                                       // 创建者
	UpdatedBy     int64       `json:"updated_by"          description:"更新者"`                                                                                       // 更新者
	CreatedAt     *gtime.Time `json:"created_at"          description:"创建时间"`                                                                                      // 创建时间
	UpdatedAt     *gtime.Time `json:"updated_at"          description:"更新时间"`                                                                                      // 更新时间
	Remark        string      `json:"remark"                 description:"备注"`                                                                                     // 备注
}
