// Package myerror
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package myerror

import (
	"context"
	"testing"

	"devinggo/modules/system/codes"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/stretchr/testify/assert"
)

func TestNewCode(t *testing.T) {
	ctx := context.Background()
	code := gcode.New(100, "TestError", nil)

	err := NewCode(ctx, code)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "TestError")
}

func TestNewCodef(t *testing.T) {
	ctx := context.Background()
	code := gcode.New(100, "TestError", nil)

	err := NewCodef(ctx, code, "test %s %d", "message", 123)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "test message 123")
}

func TestNewErrorf(t *testing.T) {
	ctx := context.Background()

	err := NewErrorf(ctx, "error: %s", "details")

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "error: details")
}

func TestMissingParameter(t *testing.T) {
	ctx := context.Background()

	err := MissingParameter(ctx, "field %s", "username")

	assert.Error(t, err)
	gErr := gerror.Code(err)
	assert.Equal(t, gcode.CodeMissingParameter.Code(), gErr.Code())
}

func TestInvalidParameter(t *testing.T) {
	ctx := context.Background()

	err := InvalidParameter(ctx, "field %s", "email")

	assert.Error(t, err)
	gErr := gerror.Code(err)
	assert.Equal(t, gcode.CodeInvalidParameter.Code(), gErr.Code())
}

func TestValidationFailed(t *testing.T) {
	ctx := context.Background()

	err := ValidationFailed(ctx, "field %s", "password")

	assert.Error(t, err)
	gErr := gerror.Code(err)
	assert.Equal(t, gcode.CodeValidationFailed.Code(), gErr.Code())
}

func TestApiTokenIsExpire(t *testing.T) {
	ctx := context.Background()

	err := ApiTokenIsExpire(ctx, "token %s", "expired")

	assert.Error(t, err)
	gErr := gerror.Code(err)
	assert.Equal(t, codes.ApiTokenIsExpire.Code(), gErr.Code())
}

func TestNotAuthorized(t *testing.T) {
	ctx := context.Background()

	err := NotAuthorized(ctx)

	assert.Error(t, err)
	gErr := gerror.Code(err)
	assert.Equal(t, gcode.CodeNotAuthorized.Code(), gErr.Code())
}

func TestNotLogged(t *testing.T) {
	ctx := context.Background()

	err := NotLogged(ctx)

	assert.Error(t, err)
	gErr := gerror.Code(err)
	assert.Equal(t, codes.CodeNotLogged.Code(), gErr.Code())
}
