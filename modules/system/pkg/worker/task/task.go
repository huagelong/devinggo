// Package task
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package task

import (
	glob2 "devinggo/modules/system/pkg/worker/glob"
	"context"
	"dario.cat/mergo"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/hibiken/asynq"
	"sync"
)

type TaskInterface interface {
	GetType() string                                  //获取任务类型
	GetPayload() *glob2.Payload                       // 获取数据
	Send(ctx context.Context, data interface{}) error // 发送任务
}

type TaskSimpleInterface interface {
	GetType() string            //获取任务类型
	GetPayload() *glob2.Payload // 获取数据
}

var clientInstance *asynq.Client
var once sync.Once

func GetClient(ctx context.Context) *asynq.Client {
	once.Do(func() {
		clientInstance = asynq.NewClient(glob2.GetConnConfig(ctx))
	})
	return clientInstance
}

func GetTask(taskItem TaskSimpleInterface) *asynq.Task {
	payload := taskItem.GetPayload()
	var payloadDefault = &glob2.Payload{
		QueueName: "default",
		TaskID:    gconv.String(gtime.TimestampNano()),
		Time:      asynq.ProcessIn(0),
	}
	mergo.Merge(payload, payloadDefault)
	var task *asynq.Task
	var bodyBytes []byte
	if !g.IsNil(payload) {
		j := gjson.New(payload)
		bodyBytes = gconv.Bytes(j.String())
	}
	task = asynq.NewTask(taskItem.GetType(), bodyBytes, payload.Time, asynq.Queue(payload.QueueName), asynq.TaskID(payload.TaskID))

	return task
}

func NewTask(ctx context.Context, taskItem TaskInterface) error {
	task := GetTask(taskItem)
	client := GetClient(ctx)
	var info *asynq.TaskInfo
	var err error
	info, err = client.Enqueue(task)
	if err != nil {
		glob2.WithWorkLog().Errorf(ctx, "[%s] task enqueue failed:%v", task.Type(), err)
		return err
	}
	glob2.WithWorkLog().Debugf(ctx, "[%s] task enqueued with ID: %s", task.Type(), info.ID)
	return nil
}

func NewSimpleTask(ctx context.Context, taskItem TaskSimpleInterface) error {
	task := GetTask(taskItem)
	client := GetClient(ctx)
	var info *asynq.TaskInfo
	var err error
	info, err = client.Enqueue(task)
	if err != nil {
		glob2.WithWorkLog().Errorf(ctx, "[%s] task enqueue failed:%v", task.Type(), err)
		return err
	}
	glob2.WithWorkLog().Debugf(ctx, "[%s] task enqueued with ID: %s", task.Type(), info.ID)
	return nil
}
