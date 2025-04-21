// Package cmd
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package cmd

import (
	"context"
	"devinggo/modules/system/pkg/utils"
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Version = &gcmd.Command{
		Name:        "version",
		Brief:       "print version info",
		Description: ``,
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			CmdInit(ctx, parser)
			utils.PrintVersionInfo()
			return
		},
	}
)
