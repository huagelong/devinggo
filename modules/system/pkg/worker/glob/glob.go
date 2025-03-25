// Package glob
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package glob

import (
	"context"
	"devinggo/modules/system/pkg/utils/config"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/hibiken/asynq"
	"runtime"
)

type Payload struct {
	Data      interface{}  `json:"data"`
	CrontabId int64        `json:"crontab_id"`
	Time      asynq.Option `json:"time"`
	QueueName string       `json:"queue_name"` // 队列名称, 默认为default
	TaskID    string       `json:"task_id"`    // 任务ID, 用于唯一标识一个任务
}

type TaskStruct struct {
	Name        string `json:"value"`
	Description string `json:"label"`
}

func GetConnConfig(ctx context.Context) asynq.RedisClientOpt {
	return asynq.RedisClientOpt{
		Addr:         config.GetConfigString(ctx, "worker.redis.address", "localhost:6379"),
		Password:     config.GetConfigString(ctx, "worker.redis.pass", ""),
		DB:           config.GetConfigInt(ctx, "worker.redis.db", 3),
		DialTimeout:  config.GetConfigDuration(ctx, "worker.redis.dialTimeout", "30s"),
		ReadTimeout:  config.GetConfigDuration(ctx, "worker.redis.readTimeout", "30s"),
		WriteTimeout: config.GetConfigDuration(ctx, "worker.redis.writeTimeout", "30s"),
		PoolSize:     runtime.NumCPU() * 2,
	}
}

func GetPayload(ctx context.Context, t *asynq.Task) (data *Payload, err error) {
	payload := t.Payload()
	if j, err := gjson.DecodeToJson(payload); err != nil {
		WithWorkLog().Errorf(ctx, "LoggingMiddleware Error decoding json: %s", err)
		return nil, err
	} else {
		//WithWorkLog().Debugf(ctx, "Payload: %s", j.String())
		data := new(Payload)
		//if err := j.Scan(data); err != nil {
		//	WithWorkLog().Errorf(ctx, "LoggingMiddleware Error scanning json: %s", err)
		//	return nil, err
		//}
		data.Data = j.Get("data")
		data.CrontabId = j.Get("crontab_id").Int64()
		data.QueueName = j.Get("queue_name").String()
		data.TaskID = j.Get("task_id").String()
		return data, nil
	}
}

func GetParamters[T any](ctx context.Context, t *asynq.Task) (T, error) {
	var zero T
	payload, err := GetPayload(ctx, t)
	if err != nil {
		return zero, err
	}
	payloadData := payload.Data
	if j, err := gjson.DecodeToJson(payloadData); err != nil {
		WithWorkLog().Warningf(ctx, "Error decoding json: %s", err)
		return zero, err
	} else {
		data := new(T)
		if err := j.Scan(data); err != nil {
			WithWorkLog().Warningf(ctx, "Error scanning json: %s", err)
			return zero, err
		}
		return *data, nil
	}
}
