// Package req
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package req

// CodeGenSearch 代码生成搜索条件
type CodeGenSearch struct {
	TableName string `json:"table_name"`
	Type      string `json:"type"`
}

// CodeGenUpdate 代码生成更新请求
type CodeGenUpdate struct {
	Id            int64          `json:"id" v:"required|min:1#Id不能为空"`
	TableName     string         `json:"table_name"`
	TableComment  string         `json:"table_comment"`
	Remark        string         `json:"remark"`
	ModuleName    string         `json:"module_name"`
	BelongMenuId  int64          `json:"belong_menu_id"`
	Type          string         `json:"type"`
	MenuName      string         `json:"menu_name"`
	ComponentType int            `json:"component_type"`
	TplType       string         `json:"tpl_type"`
	TreeId        string         `json:"tree_id"`
	TreeParentId  string         `json:"tree_parent_id"`
	TreeName      string         `json:"tree_name"`
	TagId         string         `json:"tag_id"`
	TagName       string         `json:"tag_name"`
	TagViewName   string         `json:"tag_view_name"`
	Fields        []CodeGenField `json:"fields"`
	MenuButtons   []string       `json:"menu_buttons"`
}

// CodeGenField 字段配置
type CodeGenField struct {
	Id            int64  `json:"id"`
	ColumnName    string `json:"column_name" v:"required#字段名称不能为空"`
	ColumnComment string `json:"column_comment"`
	ColumnType    string `json:"column_type"`
	Sort          int    `json:"sort"`
	IsRequired    int    `json:"is_required"`
	IsInsert      int    `json:"is_insert"`
	IsEdit        int    `json:"is_edit"`
	IsList        int    `json:"is_list"`
	IsQuery       int    `json:"is_query"`
	IsSort        int    `json:"is_sort"`
	QueryType     string `json:"query_type"`
	ViewType      string `json:"view_type"`
	DictType      string `json:"dict_type"`
	AllowRoles    string `json:"allow_roles"`
}

// CodeGenLoadTable 装载数据表请求
type CodeGenLoadTable struct {
	Source string                 `json:"source" v:"required#数据源不能为空"`
	Names  []CodeGenLoadTableItem `json:"names" v:"required#表名不能为空"`
}

// CodeGenLoadTableItem 装载数据表项
type CodeGenLoadTableItem struct {
	Name       string `json:"name" v:"required#表名不能为空"`
	Comment    string `json:"comment"`
	SourceName string `json:"sourceName"`
}

// CodeGenGenerate 生成代码请求
type CodeGenGenerate struct {
	Ids string `json:"ids" v:"required#生成ID不能为空"`
}
