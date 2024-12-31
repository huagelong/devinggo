// Package codes
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package codes

import (
	"devinggo/modules/system/pkg/i18n"
	"context"
	"github.com/gogf/gf/v2/errors/gcode"
)

var (
	CodeNotLogged    = gcode.New(1000, "NotLogged", nil)        //401
	CodeForbidden    = gcode.New(1001, "Forbidden", nil)        //403
	ApiNotAuthorized = gcode.New(3004, "ApiNotAuthorized", nil) // api未认证
)

func NewCode(ctx context.Context, code gcode.Code) gcode.Code {
	tfStr := i18n.Newf(ctx, code.Message())
	return gcode.New(code.Code(), tfStr, nil)
}
