// Package cmd
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package cmd

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gfile"
)

var (
	Unpack = &gcmd.Command{
		Name:        "unpack",
		Brief:       "unpack the config files and generate the necessary files for the project",
		Description: ``,
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			CmdInit(ctx, parser)
			filePath := ".init.lock"
			if gfile.Exists(filePath) {
				return gerror.New("initialization has been locked, please delete the .init.lock under the project root directory to unlock the initialization.")
			} else {
				gfile.Create(filePath)
			}
			return
		},
	}
)
