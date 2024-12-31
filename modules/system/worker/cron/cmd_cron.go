// Package cron
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package cron

import (
	"devinggo/modules/system/pkg/worker/cron"
	glob2 "devinggo/modules/system/pkg/worker/glob"
	"devinggo/modules/system/pkg/worker/task"
	"devinggo/modules/system/worker/consts"
	"context"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/hibiken/asynq"
)

var cmdCron = &ccmdCron{
	Type:        consts.CMD_CRON,
	Description: "执行命令",
	Payload:     &glob2.Payload{},
}

type ccmdCron struct {
	Type        string
	Description string
	Payload     *glob2.Payload
}

type CmdCronData struct {
	Cmd string `json:"cmd"`
}

func init() {
	cron.Register(cmdCron)
}

func (s *ccmdCron) GetType() string {
	return s.Type
}

func (s *ccmdCron) GetCronTask() *asynq.Task {
	return task.GetTask(s)
}

func (s *ccmdCron) GetPayload() *glob2.Payload {
	return s.Payload
}

func (s *ccmdCron) GetDescription() string {
	return s.Description
}

func (s *ccmdCron) SetParams(ctx context.Context, params *gjson.Json) {
	if g.IsEmpty(params) {
		return
	}
	data := new(CmdCronData)
	if err := params.Scan(data); err != nil {
		glob2.WithWorkLog().Errorf(ctx, "[%s] cron SetParams failed:%v", s.Type, err)
		return
	}
	s.Payload.Data = data
}
