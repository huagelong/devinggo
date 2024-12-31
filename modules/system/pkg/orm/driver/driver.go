// Package driver
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package driver

import (
	"devinggo/modules/system/consts"
	"devinggo/modules/system/pkg/contexts"
	"context"
)

func GetTenant(ctx context.Context) uint64 {
	return contexts.New().GetTenantId(ctx)
}

func GetColumnName() string {
	return consts.TenantId
}
