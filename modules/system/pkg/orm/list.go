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

// QueryBuilder 查询构建器，提供链式调用接口
type QueryBuilder struct {
	model *gdb.Model
}

// NewQuery 创建新的查询构建器
func NewQuery(m *gdb.Model) *QueryBuilder {
	return &QueryBuilder{model: m}
}

// WithRecycle 设置回收站查询
func (qb *QueryBuilder) WithRecycle(recycle bool) *QueryBuilder {
	if recycle {
		qb.model = qb.model.Unscoped().Where("deleted_at is not null")
	}
	return qb
}

// WithFilterAuth 设置权限过滤
func (qb *QueryBuilder) WithFilterAuth(filterAuth bool) *QueryBuilder {
	if filterAuth {
		qb.model = qb.model.Handler(handler.FilterAuth)
	}
	return qb
}

// WithWhere 设置查询条件
func (qb *QueryBuilder) WithWhere(params ...g.Map) *QueryBuilder {
	if len(params) > 0 && !g.IsEmpty(params[0]) {
		qb.model = qb.model.Where(params[0])
	}
	return qb
}

// WithFields 设置查询字段
func (qb *QueryBuilder) WithFields(fields interface{}) *QueryBuilder {
	if !g.IsEmpty(fields) {
		qb.model = qb.model.Fields(fields)
	}
	return qb
}

// WithOrder 设置排序
func (qb *QueryBuilder) WithOrder(orderBy string, orderType ...string) *QueryBuilder {
	if !g.IsEmpty(orderBy) {
		orderTypeStr := "asc"
		if len(orderType) > 0 && !g.IsEmpty(orderType[0]) {
			orderTypeStr = orderType[0]
		}
		qb.model = qb.model.Order(orderBy + " " + orderTypeStr)
	}
	return qb
}

// WithPage 设置分页
func (qb *QueryBuilder) WithPage(pageNum, pageSize int) *QueryBuilder {
	if pageNum <= 0 {
		pageNum = page.DefaultPage
	}
	if pageSize <= 0 {
		pageSize = page.DefaultPageSize
	}
	qb.model = qb.model.Page(pageNum, pageSize)
	return qb
}

// Build 构建最终的 Model
func (qb *QueryBuilder) Build() *gdb.Model {
	return qb.model
}

// ScanOne 扫描单条数据（便捷方法）
func (qb *QueryBuilder) ScanOne(result interface{}) error {
	return qb.model.Scan(result)
}

// ScanAll 扫描所有数据（便捷方法）
func (qb *QueryBuilder) ScanAll(results interface{}) error {
	return qb.model.Scan(results)
}

// ScanAndCount 扫描数据并统计总数（便捷方法）
func (qb *QueryBuilder) ScanAndCount(results interface{}, total *int) error {
	return qb.model.ScanAndCount(results, total, false)
}

// WithPageListReq 使用 PageListReq 配置查询构建器
func (qb *QueryBuilder) WithPageListReq(req *model.PageListReq, params ...g.Map) *QueryBuilder {
	if !g.IsEmpty(req.Recycle) && req.Recycle {
		qb.WithRecycle(true)
	}

	if !g.IsEmpty(req.FilterAuth) && req.FilterAuth {
		qb.WithFilterAuth(true)
	}

	if len(params) > 0 {
		qb.WithWhere(params...)
	}

	if !g.IsEmpty(req.Select) {
		qb.WithFields(req.Select)
	}

	pageNum := page.DefaultPage
	if !g.IsEmpty(req.Page) {
		pageNum = req.Page
	}
	pageSize := page.DefaultPageSize
	if !g.IsEmpty(req.PageSize) {
		pageSize = req.PageSize
	}
	qb.WithPage(pageNum, pageSize)

	if !g.IsEmpty(req.OrderBy) {
		orderTypeStr := "asc"
		if !g.IsEmpty(req.OrderType) {
			orderTypeStr = req.OrderType
		}
		qb.model = qb.model.Order(req.OrderBy + " " + orderTypeStr)
	} else {
		qb.model = qb.model.OrderDesc("id")
	}

	return qb
}

// WithListReq 使用 ListReq 配置查询构建器
func (qb *QueryBuilder) WithListReq(req *model.ListReq, params ...g.Map) *QueryBuilder {
	if !g.IsEmpty(req.Recycle) && req.Recycle {
		qb.WithRecycle(true)
	}

	if !g.IsEmpty(req.FilterAuth) && req.FilterAuth {
		qb.WithFilterAuth(true)
	}

	if len(params) > 0 {
		qb.WithWhere(params...)
	}

	if !g.IsEmpty(req.Select) {
		qb.WithFields(req.Select)
	}

	if !g.IsEmpty(req.OrderBy) {
		orderTypeStr := "asc"
		if !g.IsEmpty(req.OrderType) {
			orderTypeStr = req.OrderType
		}
		qb.model = qb.model.Order(req.OrderBy + " " + orderTypeStr)
	}

	return qb
}
