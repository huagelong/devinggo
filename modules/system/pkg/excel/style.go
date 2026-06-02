// Package excel
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package excel

import (
	"github.com/xuri/excelize/v2"
)

func CreateDefaultHeader() *excelize.Style {
	style := excelize.Style{
		Border: nil,
		Fill:   excelize.Fill{Type: "gradient", Color: []string{"#808080", "#808080"}, Shading: 1},
		Font: &excelize.Font{
			Family: "Arial",
			Size:   10,
			Color:  "ffffff",
			Bold:   true,
		},
		Alignment: &excelize.Alignment{
			Horizontal:      "center",
			Indent:          1,
			JustifyLastLine: true,
			ReadingOrder:    0,
			RelativeIndent:  1,
			ShrinkToFit:     false,
			TextRotation:    0,
			Vertical:        "",
			WrapText:        false,
		},
		Protection:   nil,
		NumFmt:       0,
		CustomNumFmt: nil,
		NegRed:       false,
	}
	return &style
}

func CreateDefaultData() *excelize.Style {
	style := excelize.Style{
		Border: []excelize.Border{
			{Type: "left", Color: "050505", Style: 1},
			{Type: "top", Color: "050505", Style: 1},
			{Type: "bottom", Color: "050505", Style: 1},
			{Type: "right", Color: "050505", Style: 1},
		},
		Alignment: &excelize.Alignment{
			Horizontal:      "center",
			Indent:          1,
			JustifyLastLine: true,
			ReadingOrder:    0,
			RelativeIndent:  1,
			ShrinkToFit:     false,
			TextRotation:    0,
			Vertical:        "",
			WrapText:        false,
		},
		Protection:   nil,
		NumFmt:       0,
		CustomNumFmt: nil,
		NegRed:       false,
	}
	return &style
}
