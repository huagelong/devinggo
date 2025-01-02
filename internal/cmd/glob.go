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
	"github.com/gogf/gf/v2/os/gres"
	"github.com/gogf/gf/v2/util/gmode"
)

func CmdInit(ctx context.Context, parser *gcmd.Parser) {
	once.Do(func() {
		configFile := parser.GetOpt("config").String()
		if configFile != "" {
			g.Cfg().GetAdapter().(*gcfg.AdapterFile).SetFileName(configFile)
		} else {
			if gfile.Exists("./config.yaml") {
				g.Cfg().GetAdapter().(*gcfg.AdapterFile).SetFileName("./config.yaml")
			}
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

		exportDirs := g.SliceStr{"resource", "config.yaml"}
		for _, dir := range exportDirs {
			exportDir(ctx, dir)
		}
		//设置缓存
		cache.SetAdapter(ctx)
	})
}

func exportDir(ctx context.Context, dir string) {
	if !gfile.Exists(dir) {
		if dir == "config.yaml" {
			configFullPath := "manifest/config"
			if gfile.Exists(configFullPath) {
				//g.Log().Debug(ctx, "manifest/config found")
				return
			}
			options := gres.ExportOption{RemovePrefix: "manifest/config/"}
			err := gres.Export(configFullPath, "./", options)
			if err != nil {
				g.Log().Panic(ctx, "export Error:", err)
			} else {
				g.Log().Debug(ctx, "export success:", dir)
			}
			return
		}
		//other file
		files := gres.ScanDir(dir, "*.*", true)
		if len(files) > 0 {
			err := gres.Export(dir, "./")
			if err != nil {
				g.Log().Panic(ctx, "export Error:", err)
			} else {
				g.Log().Debug(ctx, "export success:", dir)
			}
		}
	}
}
