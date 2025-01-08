// Package orm
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package orm

import (
	"context"
	"devinggo/modules/system/model/page"
	"devinggo/modules/system/pkg/handler"
	"devinggo/modules/system/pkg/utils"
	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gstructs"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

type remote[T any] struct {
	Model *gdb.Model
	req   T
}

func NewRemote[T any](m *gdb.Model, req T) *remote[T] {
	return &remote[T]{
		Model: m,
		req:   req,
	}
}

func (s *remote[T]) GetRemote(ctx context.Context, params *gmap.StrAnyMap) (res []T, total int, err error) {
	m := s.remoteOption(params)
	openPage := params.GetVar("openPage")
	if !g.IsEmpty(openPage) && openPage.Bool() {
		pageNameParam := params.GetVar("pageName")
		pageName := "page"
		if !g.IsEmpty(pageNameParam) {
			pageName = pageNameParam.String()
		}

		pageIntParam := params.GetVar(pageName)
		pageInt := page.DefaultPage
		if !g.IsEmpty(pageIntParam) {
			pageInt = pageIntParam.Int()
		}
		pageSizeParam := params.GetVar("pageSize")
		pageSize := page.DefaultPageSize
		if !g.IsEmpty(pageSizeParam) {
			pageSize = pageSizeParam.Int()
		}
		return s.Paginate(m, page.PageReq{
			Page:     pageInt,
			PageSize: pageSize,
		})
	} else {
		res, err = s.GetList(m)
		return
	}
}

func (s *remote[T]) remoteOption(params *gmap.StrAnyMap) *gdb.Model {
	m := s.Model
	if !g.IsEmpty(s.req) {
		tagNames, _ := gstructs.TagMapName(s.req, []string{"orm"})
		if !g.IsEmpty(tagNames) {
			for _, tagName := range tagNames {
				if gstr.Contains(tagName, "with") {
					m = m.WithAll()
					break
				}
			}
		}
	}

	if g.IsEmpty(params) {
		return m
	}

	options := params.GetVar("remoteOption")
	if g.IsEmpty(options) {
		return m
	}
	optionsMap := gmap.NewStrAnyMapFrom(options.Map())
	if optionsMap.Contains("relations") {
		optionsMap.Remove("relations")
	}

	selectParams := optionsMap.GetVar("select")
	if !g.IsEmpty(selectParams) {
		selectFields := gconv.SliceStr(selectParams)
		m = m.Fields(selectFields)
	}
	sort := optionsMap.GetVar("sort")
	if !g.IsEmpty(sort) {
		sortMap := sort.Map()
		for k, v := range sortMap {
			m = m.Order(gconv.String(k) + " " + gconv.String(v))
		}
	}
	group := optionsMap.GetVar("group")
	if !g.IsEmpty(group) {
		groups := gconv.SliceStr(group)
		m = m.Group(groups...)
	}

	dataScope := optionsMap.GetVar("dataScope")
	if !g.IsEmpty(dataScope) {
		m = m.Handler(handler.FilterAuth)
	}

	filter := optionsMap.GetVar("filter")
	if !g.IsEmpty(filter) {
		filterMaps := filter.Map()
		for k, v := range filterMaps {
			tmpv := gconv.SliceAny(v)
			if gstr.ToLower(gconv.String(tmpv[0])) == "like" {
				m = m.Where(k+" LIKE ?", "%"+gconv.String(tmpv[1])+"%")
			} else if gstr.ToLower(gconv.String(tmpv[0])) == "like%" {
				m = m.Where(k+" LIKE ?", gconv.String(tmpv[1])+"%")
			} else {
				m = m.Where(k, tmpv...)
			}
		}
	}
	return m
}

func (s *remote[T]) Paginate(m *gdb.Model, page page.PageReq) (res []T, total int, err error) {
	err = m.Page(page.Page, page.PageSize).ScanAndCount(&res, &total, false)
	if utils.IsError(err) {
		return nil, 0, err
	}
	return
}

func (s *remote[T]) GetList(m *gdb.Model) (res []T, err error) {
	err = m.Scan(&res)
	if utils.IsError(err) {
		return nil, err
	}
	return
}
