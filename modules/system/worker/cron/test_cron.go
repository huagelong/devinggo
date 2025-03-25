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

var testCron = &ctestCron{
	Type:        consts.TEST_CRON,
	Description: "This is a test cron",
	Payload:     &glob2.Payload{},
}

type ctestCron struct {
	Type        string
	Description string
	Payload     *glob2.Payload
}

type TestCronData struct {
	Name string `json:"name"`
}

func init() {
	cron.Register(testCron)
}

func (s *ctestCron) GetType() string {
	return s.Type
}

func (s *ctestCron) GetCronTask() *asynq.Task {
	return task.GetTask(s)
}

func (s *ctestCron) GetPayload() *glob2.Payload {
	return s.Payload
}

func (s *ctestCron) GetDescription() string {
	return s.Description
}

func (s *ctestCron) SetParams(ctx context.Context, params *gjson.Json) {
	if g.IsEmpty(params) {
		return
	}
	data := new(TestCronData)
	if err := params.Scan(data); err != nil {
		glob2.WithWorkLog().Errorf(ctx, "[%s] cron SetParams failed:%v data:%s", s.Type, err, data)
		return
	}
	s.Payload.Data = data
}
