// Package req
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE
package req

type SettingGenerateColumnsSearch struct {
	TableId int `json:"table_id"`
}

type SettingGenerateColumnsUpdate struct {
	Id            int64  `json:"id"`
	IsInsert      bool   `json:"is_insert" d:"false" `
	IsEdit        bool   `json:"is_edit"  d:"false" `
	IsList        bool   `json:"is_list"  d:"false" `
	IsQuery       bool   `json:"is_query"  d:"false" `
	IsSort        bool   `json:"is_sort"  d:"false" `
	IsRequired    bool   `json:"is_required"  d:"false" `
	AllowRoles    string `json:"allow_roles"`
	ColumnComment string `json:"column_comment"`
	ColumnName    string `json:"column_name"`
	ColumnType    string `json:"column_type"`
	DictType      string `json:"dict_type"`
	Extra         string `json:"extra"`
	IsPk          int    `json:"is_pk"`
	Options       string `json:"options"`
	QueryType     string `json:"query_type"`
	Remark        string `json:"remark"`
	Sort          int    `json:"sort"`
	ViewType      string `json:"view_type"`
}
