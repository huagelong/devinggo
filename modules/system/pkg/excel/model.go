// Package excel
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package excel

import (
	"reflect"
	"sort"
	"strconv"
)

var DefaultSheet = "Sheet1"

/*
*
自定义tag
excelName
excelIndex
toExcelFormat
toDataFormat
excelColWidth
*/
type model struct {
	excelName     string
	excelIndex    int
	toExcelFormat string
	toDataFormat  string
	excelColWidth int
	fieldName     string
	fieldIndex    int
}

// 根据类型获取打印相关内容
func getInterfaceExcelModel(field reflect.Type) *[]*model {
	m := make([]*model, 0)

	//获取tag 根据excelName 获取输出内容,根据 excelIndex 序号 excelFormat 格式化函数 excelColWidth 单元格宽度
	for i := 0; i < field.NumField(); i++ {
		tag := field.Field(i).Tag
		excelName := tag.Get("excelName")
		if excelName != "" {
			indexString := tag.Get("excelIndex")
			index := i
			if indexString != "" {
				parseInt, err := strconv.ParseInt(indexString, 10, 64)
				if err == nil {
					index = int(parseInt)
				}
			}
			name := field.Field(i).Name
			toExcelFormat := tag.Get("toExcelFormat")
			toDataFormat := tag.Get("toDataFormat")
			widthString := tag.Get("excelColWidth")
			var width int
			if widthString != "" {
				parseInt, err := strconv.ParseInt(widthString, 10, 64)
				if err == nil {
					width = int(parseInt)
				}
			} else {
				width = len(excelName) * 3
			}
			m = append(m, &model{excelName, index, toExcelFormat, toDataFormat, width, name, -1})
		}
	}
	// 排序
	sort.Slice(m, func(i, j int) bool {
		return m[i].excelIndex < m[j].excelIndex
	})
	return &m
}
