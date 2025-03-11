// Package cmd
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package cmd

import (
	"context"
	"devinggo/modules/system/pkg/cache"
	"devinggo/modules/system/pkg/utils/config"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/util/gmode"
)

func CmdInit(ctx context.Context, parser *gcmd.Parser) {
	once.Do(func() {
		filePath := "manifest/config/config.yaml"
		if !gfile.Exists(filePath) {
			g.Log().Panicf(ctx, "%s：config file not found,please run devinggo unpack", filePath)
		}
		configFile := parser.GetOpt("config").String()
		if configFile != "" {
			g.Cfg().GetAdapter().(*gcfg.AdapterFile).SetFileName(configFile)
		}

		configPath, _ := g.Cfg().GetAdapter().(*gcfg.AdapterFile).GetFilePath()
		g.Log().Debug(ctx, "use config file:", configPath)

		if parser.GetOpt("gf.gmode").IsEmpty() {
			gmodeConfig := config.GetConfigString(ctx, "gf.gmode")
			if !g.IsEmpty(gmodeConfig) {
				gmode.Set(gmodeConfig)
			}
		} else {
			gmode.Set(gmode.PRODUCT)
		}
		g.Log().Debug(ctx, "gmode:", gmode.Mode())

		// json格式日志
		logFormat := config.GetConfigString(ctx, "logger.format", "")
		if logFormat == "json" {
			glog.SetDefaultHandler(glog.HandlerJson)
		}

		// 异步打印日志 & 显示打印错误的文件行号, 对访问日志无效
		g.Log().SetFlags(glog.F_ASYNC | glog.F_TIME_STD | glog.F_FILE_LONG)

		//设置缓存
		cache.SetAdapter(ctx)
	})
}
