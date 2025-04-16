// Package middleware
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package middleware

import (
	"devinggo/modules/system/pkg/response"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/text/gstr"
	"reflect"
	"strings"
)

// ResponseHandler custom response format.
func (s *sMiddleware) ResponseHandler(r *ghttp.Request) {
	ctx := r.Context()
	//g.Log().Debug(ctx, "IsExited:", r.IsExited())
	r.Middleware.Next()
	// For /swagger
	if r.Request.URL.Path == "/api.json" {
		return
	}

	// For PProf
	if g.Cfg().MustGet(ctx, "server.pprofEnabled").Bool() {
		rPath := g.Cfg().MustGet(ctx, "server.pprofPattern").String()
		if rPath == "" {
			rPath = "/debug/pprof"
		}
		if strings.HasPrefix(r.Request.URL.Path, rPath) {
			return
		}
	}

	var (
		err = r.GetError()
	)
	if err == nil && r.Response.BufferLength() > 0 {
		return
	}

	if err == nil && gstr.Contains(r.Response.Header().Get("Content-Type"), "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet") {
		return
	}

	res, bizCode := response.ResponseHandler(r)
	data, has := getField(res, "Data")
	if has {
		res = data
	}
	response.JsonExit(r, bizCode, res)
}

func getField(s interface{}, fieldName string) (interface{}, bool) {
	if g.IsNil(s) {
		return nil, false
	}
	// 获取s的反射值对象
	v := reflect.ValueOf(s)
	// 获取结构体的类型
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	} else {
		return nil, false
	}

	t := v.Type()

	hasItems := false
	hasPageInfo := false
	hasData := false
	// 遍历结构体的所有字段
	var dataValue interface{}
	var itemsValue interface{}
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		//g.Log().Debug(context.Background(), "field.Name:", field.Name)
		if field.Name == fieldName {
			hasData = true
			dataValue = v.Field(i).Interface()
		}
		if field.Name == "Items" {
			hasItems = true
			itemsValue = v.Field(i).Interface()
		}
		if field.Name == "PageRes" {
			hasPageInfo = true
		}
	}
	//g.Log().Debug(context.Background(), "t.NumField():", t.NumField(), "hasItems:", hasItems, "hasPageInfo:", hasPageInfo, "hasData:", hasData, "itemsValue:", itemsValue, "dataValue:", dataValue)
	if hasItems && hasPageInfo && hasData && g.IsEmpty(itemsValue) {
		//判断是否remote接口返回字段
		return dataValue, true
	} else {
		if t.NumField() > 2 {
			return nil, false
		}
		return dataValue, hasData
	}
}
