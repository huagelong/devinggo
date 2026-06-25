// Package res
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package res

import "github.com/gogf/gf/v2/os/gtime"

// CodeGenTable 代码生成表信息
type CodeGenTable struct {
	Id           int64       `json:"id" description:"主键"`
	TableName    string      `json:"table_name" description:"表名称"`
	TableComment string      `json:"table_comment" description:"表描述"`
	Type         string      `json:"type" description:"生成类型: single=单表, tree=树表"`
	ModuleName   string      `json:"module_name" description:"所属模块"`
	MenuName     string      `json:"menu_name" description:"菜单名称"`
	Status       int         `json:"status" description:"状态"`
	Sort         int         `json:"sort" description:"排序"`
	CreatedBy    int64       `json:"created_by" description:"创建者"`
	UpdatedBy    int64       `json:"updated_by" description:"更新者"`
	CreatedAt    *gtime.Time `json:"created_at" description:"创建时间"`
	UpdatedAt    *gtime.Time `json:"updated_at" description:"更新时间"`
}

// CodeGenPreview 代码预览项
type CodeGenPreview struct {
	Name    string `json:"name" description:"文件名称"`
	TabName string `json:"tab_name" description:"Tab标签名称"`
	Code    string `json:"code" description:"代码内容"`
	Lang    string `json:"lang" description:"语言"`
}

// CodeGenReadTable 读取表信息响应
type CodeGenReadTable struct {
	TableName    string          `json:"table_name" description:"表名称"`
	TableComment string          `json:"table_comment" description:"表描述"`
	Columns      []CodeGenColumn `json:"columns" description:"字段列表"`
}

// CodeGenColumn 字段信息
type CodeGenColumn struct {
	ColumnName    string `json:"column_name" description:"字段名称"`
	ColumnComment string `json:"column_comment" description:"字段描述"`
	ColumnType    string `json:"column_type" description:"物理类型"`
	IsNullable    string `json:"is_nullable" description:"是否可空"`
	DataType      string `json:"data_type" description:"数据类型"`
}

// CodeGenSourceTable 数据源表信息
type CodeGenSourceTable struct {
	Name       string `json:"name" description:"表名称"`
	Comment    string `json:"comment" description:"表描述"`
	SourceName string `json:"sourceName" description:"数据源名称"`
}
