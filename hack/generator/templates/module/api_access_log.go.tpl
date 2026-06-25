// Package hook
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package hook

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
)

// ApiAccessLog API访问日志钩子
func (s *sHook) ApiAccessLog(ctx context.Context) error {
	g.Log().Debug(ctx, "{{.moduleName}} ApiAccessLog hook called")
	return nil
}
