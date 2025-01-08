// Package base
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package base

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
)

type Service interface {
	Model(ctx context.Context) *gdb.Model
}

type BaseService struct {
}
