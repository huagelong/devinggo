// Package glob
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package glob

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/hibiken/asynq"
)

type Logger struct {
	base *glog.Logger
	ctx  context.Context
}

func NewLogger(ctx context.Context) asynq.Logger {
	return &Logger{
		base: WithWorkLog(),
		ctx:  ctx,
	}
}

func WithWorkLog() *glog.Logger {
	return g.Log(LOG_NAME)
}

func (l *Logger) Debug(args ...interface{}) {
	l.base.Debug(l.ctx, args...)
}

func (l *Logger) Info(args ...interface{}) {
	l.base.Info(l.ctx, args...)
}

func (l *Logger) Warn(args ...interface{}) {
	l.base.Warning(l.ctx, args...)
}

func (l *Logger) Error(args ...interface{}) {
	l.base.Error(l.ctx, args...)
}

func (l *Logger) Fatal(args ...interface{}) {
	l.base.Fatal(l.ctx, args...)
}
