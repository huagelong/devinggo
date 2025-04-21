// Package cmd
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package cmd

import (
	"context"
	"devinggo/internal/router"
	_ "devinggo/modules/_/modules"
	"devinggo/modules/system/pkg/modules"
	"devinggo/modules/system/pkg/redis"
	"devinggo/modules/system/pkg/response"
	"devinggo/modules/system/pkg/upload"
	"devinggo/modules/system/pkg/utils/config"
	systemService "devinggo/modules/system/service"
	"github.com/gogf/gf/contrib/trace/jaeger/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/net/goai"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/genv"
	"github.com/gogf/gf/v2/os/gsession"
	"github.com/gogf/gf/v2/util/gmode"
)

var (
	Http = &gcmd.Command{
		Name:  "http",
		Usage: "http",
		Brief: "HTTP服务，也可以称为主服务，包含http、websocket",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			CmdInit(ctx, parser)
			s := g.Server()

			serverPort := genv.GetWithCmd("server.port")
			if !g.IsEmpty(serverPort) {
				s.SetPort(serverPort.Int())
			}
			s.SetSessionStorage(gsession.NewStorageRedisHashTable(redis.GetRedis()))
			s.Use(
				systemService.Middleware().I18n,
				systemService.Middleware().Ctx,
				systemService.Middleware().Cors,
				systemService.Middleware().TraceID,
				systemService.Middleware().NeverDoneCtx,
				systemService.Middleware().ResponseHandler,
			)
			// static dir setting
			uploadPath := upload.GetUploadPath(ctx)
			s.AddStaticPath("/upload", uploadPath)
			// doc
			if gmode.IsDevelop() {
				enhanceOpenAPIDoc(s)
			}

			modules.StartModules(ctx, s)
			s.Group("/", func(group *ghttp.RouterGroup) {
				router.BindController(group)
			})

			jaegerEnable := config.GetConfigBool(ctx, "jaeger.enable", false)
			if jaegerEnable {
				serviceName := config.GetConfigString(ctx, "jaeger.serviceName")
				udpEndpoint := config.GetConfigString(ctx, "jaeger.udpEndpoint")
				tp, err := jaeger.Init(serviceName, udpEndpoint)
				if err != nil {
					g.Log().Panic(ctx, err)
				}
				defer tp.Shutdown(ctx)
			}

			ServerWg.Add(1)
			// 信号监听
			SignalListen(ctx, SignalHandlerForOverall)
			s.Run()
			<-ServerCloseSignal
			modules.StopModules(ctx)
			g.Log().Info(ctx, "http server successfully closed ..")
			ServerWg.Done()
			return
		},
	}
)

func enhanceOpenAPIDoc(s *ghttp.Server) {
	openapi := s.GetOpenApi()
	// OpenApi自定义信息
	openapi.Info.Title = `devinggo Project`
	openapi.Config.CommonResponse = response.Response{}
	openapi.Config.CommonResponseDataField = `Data`
	openapi.Info.Description = ``
	openapi.Tags = &goai.Tags{}
}
