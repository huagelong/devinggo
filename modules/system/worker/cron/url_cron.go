// Package cron
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package cron

import (
	"context"

	"devinggo/modules/system/pkg/worker"
	glob2 "devinggo/modules/system/pkg/worker/glob"
	"devinggo/modules/system/worker/consts"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
)

// UrlCronData URL请求定时任务数据结构
type UrlCronData struct {
	Url         string                 `json:"url"`
	Method      string                 `json:"method"`
	Headers     map[string]string      `json:"headers"`
	Params      map[string]interface{} `json:"params"`
	Timeout     int64                  `json:"timeout"`
	Retry       int                    `json:"retry"`
	Cookies     map[string]string      `json:"cookie"`
	ContentType string                 `json:"content_type"`
	Proxy       string                 `json:"proxy"`
}

func init() {
	// 使用新方式注册Cron
	worker.RegisterCronFunc(consts.URL_CRON, "执行http请求", handleUrlCronParams)
}

// handleUrlCronParams 处理URL请求定时任务参数
func handleUrlCronParams(ctx context.Context, payload *glob2.Payload, params *gjson.Json) {
	if g.IsEmpty(params) {
		return
	}
	data := new(UrlCronData)
	if err := params.Scan(data); err != nil {
		glob2.WithWorkLog().Errorf(ctx, "[%s] cron SetParams failed:%v", consts.URL_CRON, err)
		return
	}
	payload.Data = data
}
