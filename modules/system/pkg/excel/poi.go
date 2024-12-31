// Package excel
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package excel

import (
	"bufio"
	"devinggo/modules/system/myerror"
	"devinggo/modules/system/pkg/utils"
	"devinggo/modules/system/pkg/utils/request"
	"context"
	"fmt"
	"github.com/xuri/excelize/v2"
	"io"
	"net/url"
	"os"
)

type Export[T any] struct {
	excelModel[T]
}

type Import[T any] struct {
	excelModel[T]
}

func NewExcelImportFile[T any](fileName string, t T) *Import[T] {
	e := Import[T]{}
	return e.newExcelImportFile(fileName, "", t)
}

func NewExcelImportWriter[T any](reader io.Reader, t T) *Import[T] {
	e := Import[T]{}
	return e.newExcelImportWriter(reader, "", t)
}

func NewExcelImportSheetFile[T any](fileName string, sheetName string, t T) *Import[T] {
	e := Import[T]{}
	return e.newExcelImportFile(fileName, sheetName, t)
}

func NewExcelImportSheetWriter[T any](reader io.Reader, sheetName string, t T) *Import[T] {
	e := Import[T]{}
	return e.newExcelImportWriter(reader, sheetName, t)
}

func (e *Import[T]) ImportRead(fu func(row T)) *Import[T] {
	return e.importRead(fu)
}

func (e *Import[T]) ImportDataToStruct(t *[]T) *Import[T] {
	return e.importDataToStruct(t)
}

// NewExcelExport 导出初始化
func NewExcelExport[T any](sheetName string, t T) *Export[T] {
	e := Export[T]{}
	return e.newExcelExport(sheetName, t)
}

func (e *Export[T]) SetHeadStyle(style *excelize.Style) *Export[T] {
	return e.setHeadStyle(style)
}

func (e *Export[T]) SetDataStyle(style *excelize.Style) *Export[T] {
	return e.setDataStyle(style)
}

func (e *Export[T]) ExportSmallExcelByStruct(object []T) *Export[T] {
	return e.exportData(object, 1)
}

// ExportData 指定位置导出 start 默认从1开始 1 数据开始的位置
func (e *Export[T]) ExportData(object []T, start int) *Export[T] {
	return e.exportData(object, start)
}

func (e *excelModel[T]) WriteInWriter(writer io.Writer) *excelModel[T] {
	e.paddingDataStyle()
	err := e.f.Write(writer)
	if err != nil {
		e.err = err
	}
	return e
}

func (e *excelModel[T]) Download(ctx context.Context, fileName string) *excelModel[T] {
	r := request.GetHttpRequest(ctx)
	if r == nil {
		err := myerror.ValidationFailed(ctx, "ctx not http request")
		e.err = err
		return e
	}
	fileName = fileName + ".xlsx"
	r.Response.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	r.Response.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s;filename*=UTF-8''%s", url.QueryEscape(fileName), url.QueryEscape(fileName)))
	r.Response.Header().Set("Content-Transfer-Encoding", "binary")
	r.Response.Header().Set("content-description", "File Transfer")
	r.Response.Header().Set("pragma", "public")
	r.Response.Header().Set("Access-Control-Expose-Headers", "Content-Disposition")

	path := utils.GetTmpDir() + "/" + fileName

	e.WriteInFileName(path)
	if e.err != nil {
		return e
	}
	//g.Log().Info(ctx, "path:", path)
	r.Response.ServeFile(path)
	return e
}

func (e *excelModel[T]) WriteInFileName(resultFile string) *excelModel[T] {
	file, err := os.OpenFile(resultFile, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, os.ModeDevice|os.ModePerm)
	defer file.Close()

	if err == nil {
		writer := bufio.NewWriter(file)
		e.WriteInWriter(writer)
		writer.Flush()
	}
	e.err = err
	return e
}

func (e *excelModel[T]) Close() {
	err := e.f.Close()
	if err != nil {
		return
	}
}

func (e *excelModel[T]) Error() error {
	return e.err
}
