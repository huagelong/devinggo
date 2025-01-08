// Package orm
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package orm

import (
	"devinggo/modules/system/model"
	"devinggo/modules/system/model/page"
	"devinggo/modules/system/pkg/handler"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

func GetPageList(m *gdb.Model, req *model.PageListReq, params ...g.Map) *gdb.Model {

	if !g.IsEmpty(req.Recycle) {
		if req.Recycle {
			m = m.Unscoped().Where("deleted_at is not null")
		}
	}

	if !g.IsEmpty(req.FilterAuth) {
		if req.FilterAuth {
			m = m.Handler(handler.FilterAuth)
		}
	}

	if !g.IsEmpty(params) {
		paramsData := interface{}(nil)
		if len(params) > 0 {
			paramsData = params[0]
		}
		m = m.Where(paramsData)
	}

	if !g.IsEmpty(req.Select) {
		m = m.Fields(req.Select)
	}

	if !g.IsEmpty(req.PageSize) {
		pageNum := page.DefaultPage
		if !g.IsEmpty(req.Page) {
			pageNum = req.Page
		}
		m = m.Page(pageNum, req.PageSize)
	} else {
		pageNum := 1
		if !g.IsEmpty(req.Page) {
			pageNum = req.Page
		}
		m = m.Page(pageNum, page.DefaultPageSize)
	}

	if !g.IsEmpty(req.OrderBy) {
		orderType := "asc"
		if !g.IsEmpty(req.OrderType) {
			orderType = req.OrderType
		}
		m = m.Order(req.OrderBy, orderType)
	}
	return m
}

func GetList(m *gdb.Model, req *model.ListReq, params ...g.Map) *gdb.Model {

	if !g.IsEmpty(req.Recycle) {
		if req.Recycle {
			m = m.Unscoped().Where("deleted_at is not null")
		}
	}

	if !g.IsEmpty(req.FilterAuth) {
		if req.FilterAuth {
			m = m.Handler(handler.FilterAuth)
		}
	}

	if !g.IsEmpty(params) {
		paramsData := interface{}(nil)
		if len(params) > 0 {
			paramsData = params[0]
		}
		m = m.Where(paramsData)
	}

	if !g.IsEmpty(req.Select) {
		m = m.Fields(req.Select)
	}

	if !g.IsEmpty(req.OrderBy) {
		orderType := "asc"
		if !g.IsEmpty(req.OrderType) {
			orderType = req.OrderType
		}
		m = m.Order(req.OrderBy, orderType)
	}
	return m
}
