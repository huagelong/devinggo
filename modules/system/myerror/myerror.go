// Package myerror
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package myerror

import (
	"context"
	"devinggo/modules/system/codes"
	"devinggo/modules/system/pkg/i18n"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

func NewCode(ctx context.Context, code gcode.Code) error {
	tfStr := i18n.Newf(ctx, code.Message())
	return gerror.NewCode(code, tfStr)
}

func NewCodef(ctx context.Context, code gcode.Code, message string, params ...string) error {
	tfStr := i18n.Newf(ctx, message, params...)
	return gerror.NewCode(code, tfStr)
}

func NewErrorf(ctx context.Context, message string, params ...string) error {
	tfStr := i18n.Newf(ctx, message, params...)
	return gerror.New(tfStr)
}

func MissingParameter(ctx context.Context, message string, params ...string) error {
	return NewCodef(ctx, gcode.CodeMissingParameter, message, params...)
}

func InvalidParameter(ctx context.Context, message string, params ...string) error {
	return NewCodef(ctx, gcode.CodeInvalidParameter, message, params...)
}

func ValidationFailed(ctx context.Context, message string, params ...string) error {
	return NewCodef(ctx, gcode.CodeValidationFailed, message, params...)
}

func ApiTokenIsExpire(ctx context.Context, message string, params ...string) error {
	return NewCodef(ctx, codes.ApiTokenIsExpire, message, params...)
}

func NotAuthorized(ctx context.Context) error {
	return NewCode(ctx, gcode.CodeNotAuthorized)
}

func NotLogged(ctx context.Context) error {
	return NewCode(ctx, codes.CodeNotLogged)
}
