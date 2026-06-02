// Package server
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package server

import (
	"context"
	"strings"
	"time"

	"devinggo/modules/system/myerror"
	"devinggo/modules/system/pkg/worker"
	glob2 "devinggo/modules/system/pkg/worker/glob"
	"devinggo/modules/system/worker/consts"
	"devinggo/modules/system/worker/cron"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/hibiken/asynq"
)

func init() {
	// 使用新方式注册Worker
	worker.RegisterWorkerFunc(consts.URL_CRON, executeUrlCronWorker)
}

// executeUrlCronWorker 执行URL请求Worker
func executeUrlCronWorker(ctx context.Context, t *asynq.Task) error {
	data, err := glob2.GetParamters[cron.UrlCronData](ctx, t)
	if err != nil {
		return err
	}
	glob2.WithWorkLog().Infof(ctx, `type:%s, jsonData:%+v`, t.Type(), data)

	url := data.Url
	if g.IsEmpty(url) {
		return myerror.MissingParameter(ctx, `url为空`)
	}
	method := data.Method
	if g.IsEmpty(method) {
		method = "GET"
	}
	method = strings.ToUpper(method)
	params := data.Params
	timeout := data.Timeout
	if g.IsEmpty(timeout) {
		timeout = 30
	}
	client := g.Client()
	var dura time.Duration
	dura = time.Second * gconv.Duration(timeout)
	client.SetTimeout(dura)
	if !g.IsEmpty(data.Retry) {
		client.Retry(data.Retry, time.Second*5)
	}
	if !g.IsEmpty(data.Cookies) {
		client.SetCookieMap(data.Cookies)
	}

	if !g.IsEmpty(data.Headers) {
		client.SetHeaderMap(data.Headers)
	}

	if !g.IsEmpty(data.ContentType) {
		client.ContentType(data.ContentType)
	}

	if !g.IsEmpty(data.Proxy) {
		client.Proxy(data.Proxy)
	}

	resContent := ""
	if method == "GET" {
		r := client.GetContent(ctx, url, params)
		resContent = r
	}

	if method == "POST" {
		r := client.PostContent(ctx, url, params)
		resContent = r
	}

	if method == "PUT" {
		r := client.PutContent(ctx, url, params)
		resContent = r
	}

	if method == "DELETE" {
		r := client.DeleteContent(ctx, url, params)
		resContent = r
	}

	glob2.WithWorkLog().Infof(ctx, `type:%s, response:%+v`, t.Type(), resContent)
	return nil
}
