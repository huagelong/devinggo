// Package req
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package req

type {{.EntityName}}Search struct {
{{.SearchFields}}}

type {{.EntityName}}Save struct {
{{.SaveFields}}}

type {{.EntityName}}Update struct {
{{.UpdateFields}}}
