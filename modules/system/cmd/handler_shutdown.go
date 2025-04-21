// Package cmd
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package cmd

import (
	"context"
	"devinggo/modules/system/consts"
	"devinggo/modules/system/pkg/utils"
	"devinggo/modules/system/pkg/utils/event"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gproc"
	"os"
	"sync"
)

var (
	ServerCloseSignal = make(chan struct{}, 1)
	ServerWg          = sync.WaitGroup{}
	once              sync.Once
)

// signalHandlerForOverall 关闭信号处理
func SignalHandlerForOverall(sig os.Signal) {
	serverCloseEvent(gctx.GetInitCtx())
	ServerCloseSignal <- struct{}{}
}

// signalListen 信号监听
func SignalListen(ctx context.Context, handler ...gproc.SigHandler) {
	utils.SafeGo(ctx, func(ctx context.Context) {
		gproc.AddSigHandlerShutdown(handler...)
		gproc.Listen()
	})
}

// serverCloseEvent 关闭事件
// 区别于服务收到退出信号后的处理，只会执行一次
func serverCloseEvent(ctx context.Context) {
	once.Do(func() {
		event.Event().Call(consts.EventServerClose, ctx)
	})
}
