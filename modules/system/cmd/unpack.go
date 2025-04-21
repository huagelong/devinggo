// Package cmd
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package cmd

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gres"
	"github.com/gogf/gf/v2/text/gstr"
)

var (
	Unpack = &gcmd.Command{
		Name:        "unpack",
		Brief:       "unpack the config files and generate the necessary files for the project",
		Description: ``,
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			filePath := ".init.lock"
			if gfile.Exists(filePath) {
				return gerror.New("initialization has been locked, please delete the .init.lock under the project root directory to unlock the initialization.")
			} else {
				exportDirs := g.SliceStr{"resource", "manifest"}
				for _, dir := range exportDirs {
					exportDir(ctx, dir)
				}
				gfile.Create(filePath)
			}
			return
		},
	}
)

func exportDir(ctx context.Context, dir string) {
	//other file
	files := gres.ScanDir(dir, "*.*", true)
	g.Log().Debug(ctx, "files:", files)
	if len(files) > 0 {
		var (
			err  error
			name string
			path string
		)
		for _, file := range files {
			name = file.Name()
			name = gstr.Trim(name, `\/`)
			if name == "" {
				continue
			}
			path = gfile.Join("./", name)
			if gfile.Exists(path) {
				g.Log().Debug(ctx, path+" found")
				continue
			}
			if file.FileInfo().IsDir() {
				err = gfile.Mkdir(path)
			} else {
				err = gfile.PutBytes(path, file.Content())
			}
			if err != nil {
				g.Log().Panic(ctx, "export Error:", err)
			} else {
				g.Log().Debug(ctx, "export success:", path)
			}
		}
	}

}
