// Package cron
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package cron

import (
	"context"
	"devinggo/modules/system/pkg/worker/cron"
	glob2 "devinggo/modules/system/pkg/worker/glob"
	"devinggo/modules/system/pkg/worker/task"
	"devinggo/modules/system/worker/consts"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/hibiken/asynq"
)

var urlCron = &curlCron{
	Type:        consts.URL_CRON,
	Description: "执行http请求",
	Payload:     &glob2.Payload{},
}

type curlCron struct {
	Type        string
	Description string
	Payload     *glob2.Payload
}

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
	cron.Register(urlCron)
}

func (s *curlCron) GetType() string {
	return s.Type
}

func (s *curlCron) GetCronTask() *asynq.Task {
	return task.GetTask(s)
}

func (s *curlCron) GetPayload() *glob2.Payload {
	return s.Payload
}

func (s *curlCron) GetDescription() string {
	return s.Description
}

func (s *curlCron) SetParams(ctx context.Context, params *gjson.Json) {
	if g.IsEmpty(params) {
		return
	}
	data := new(UrlCronData)
	if err := params.Scan(data); err != nil {
		glob2.WithWorkLog().Errorf(ctx, "[%s] cron SetParams failed:%v", s.Type, err)
		return
	}
	s.Payload.Data = data
}
