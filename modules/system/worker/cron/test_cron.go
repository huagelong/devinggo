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

// TestCronData 测试定时任务数据结构
type TestCronData struct {
	Name string `json:"name"`
}

func init() {
	// 使用新方式注册Cron
	worker.RegisterCronFunc(consts.TEST_CRON, "This is a test cron", handleTestCronParams)
}

// handleTestCronParams 处理测试定时任务参数
func handleTestCronParams(ctx context.Context, payload *glob2.Payload, params *gjson.Json) {
	if g.IsEmpty(params) {
		return
	}
	data := new(TestCronData)
	if err := params.Scan(data); err != nil {
		glob2.WithWorkLog().Errorf(ctx, "[%s] cron SetParams failed:%v data:%s", consts.TEST_CRON, err, data)
		return
	}
	payload.Data = data
}
