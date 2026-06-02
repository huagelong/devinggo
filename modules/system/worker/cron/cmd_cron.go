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

// CmdCronData 命令执行定时任务数据结构
type CmdCronData struct {
	Cmd string `json:"cmd"`
}

func init() {
	// 使用新方式注册Cron
	worker.RegisterCronFunc(consts.CMD_CRON, "执行命令", handleCmdCronParams)
}

// handleCmdCronParams 处理命令定时任务参数
func handleCmdCronParams(ctx context.Context, payload *glob2.Payload, params *gjson.Json) {
	if g.IsEmpty(params) {
		return
	}
	data := new(CmdCronData)
	if err := params.Scan(data); err != nil {
		glob2.WithWorkLog().Errorf(ctx, "[%s] cron SetParams failed:%v", consts.CMD_CRON, err)
		return
	}
	payload.Data = data
}
