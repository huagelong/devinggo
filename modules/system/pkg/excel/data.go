// Package excel
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package excel

import (
	"github.com/xuri/excelize/v2"
	"strconv"
)

func (e *Export[T]) createHead() {
	for r, m := range *e.mod {
		name, err := excelize.ColumnNumberToName(r + 1)
		if err != nil {
			e.err = err
		}
		s := name + strconv.Itoa(1)
		e.f.SetCellValue(e.sheetName, s, m.excelName)
		e.f.SetColWidth(e.sheetName, name, name, float64(m.excelColWidth))
	}
}
