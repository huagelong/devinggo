// Package hook
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package hook

import (
	"context"
	"database/sql"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// Hook 封装了数据库钩子的配置
type Hook struct {
	autoCreatedUpdatedBy bool
	cacheEvict           bool
	userRelate           bool
	userRelateFields     []string
}

// Option 定义钩子配置选项函数类型
type Option func(*Hook)

// WithAutoCreatedUpdatedBy 启用自动填充 created_by 和 updated_by 字段
func WithAutoCreatedUpdatedBy() Option {
	return func(h *Hook) {
		h.autoCreatedUpdatedBy = true
	}
}

// WithoutAutoCreatedUpdatedBy 禁用自动填充 created_by 和 updated_by 字段
func WithoutAutoCreatedUpdatedBy() Option {
	return func(h *Hook) {
		h.autoCreatedUpdatedBy = false
	}
}

// WithCacheEvict 启用缓存清理
func WithCacheEvict() Option {
	return func(h *Hook) {
		h.cacheEvict = true
	}
}

// WithoutCacheEvict 禁用缓存清理
func WithoutCacheEvict() Option {
	return func(h *Hook) {
		h.cacheEvict = false
	}
}

// WithUserRelate 启用用户关联查询，指定需要关联的字段名
func WithUserRelate(fieldNames ...string) Option {
	return func(h *Hook) {
		h.userRelate = true
		h.userRelateFields = fieldNames
	}
}

// WithoutUserRelate 禁用用户关联查询
func WithoutUserRelate() Option {
	return func(h *Hook) {
		h.userRelate = false
		h.userRelateFields = nil
	}
}

// New 创建一个新的 Hook 实例，默认启用所有功能
func New(opts ...Option) gdb.HookHandler {
	h := &Hook{
		autoCreatedUpdatedBy: true,
		cacheEvict:           true,
		userRelate:           false,
		userRelateFields:     nil,
	}

	for _, opt := range opts {
		opt(h)
	}

	return h.build()
}

// Default 返回默认配置的 Hook（启用自动填充和缓存清理）
func Default() gdb.HookHandler {
	return New()
}

// ReadOnly 返回只读模式的 Hook（只启用用户关联查询）
func ReadOnly(fieldNames ...string) gdb.HookHandler {
	return New(
		WithoutAutoCreatedUpdatedBy(),
		WithoutCacheEvict(),
		WithUserRelate(fieldNames...),
	)
}

// Minimal 返回最小配置的 Hook（禁用所有功能）
func Minimal() gdb.HookHandler {
	return New(
		WithoutAutoCreatedUpdatedBy(),
		WithoutCacheEvict(),
		WithoutUserRelate(),
	)
}

// build 构建并返回 gdb.HookHandler
func (h *Hook) build() gdb.HookHandler {
	return gdb.HookHandler{
		Select: func(ctx context.Context, in *gdb.HookSelectInput) (result gdb.Result, err error) {
			result, err = in.Next(ctx)
			if err != nil {
				return result, err
			}
			if h.userRelate && len(h.userRelateFields) > 0 {
				return UserRelate(ctx, result, h.userRelateFields)
			}
			return
		},

		Insert: func(ctx context.Context, in *gdb.HookInsertInput) (result sql.Result, err error) {
			if h.autoCreatedUpdatedBy {
				err = AutoCreatedUpdatedByInsert(ctx, in)
				if err != nil {
					return nil, err
				}
			}

			if h.cacheEvict {
				err = CleanCache[gdb.HookInsertInput](ctx, in)
				if err != nil {
					return nil, err
				}
			}

			result, err = in.Next(ctx)
			if err != nil {
				g.Log().Debug(ctx, "in:", in)
				g.Log().Debug(ctx, "Insert error:", err)
			}
			return
		},

		Update: func(ctx context.Context, in *gdb.HookUpdateInput) (result sql.Result, err error) {
			if h.autoCreatedUpdatedBy {
				err = AutoCreatedUpdatedByUpdate(ctx, in)
				if err != nil {
					return nil, err
				}
			}

			if h.cacheEvict {
				err = CleanCache[gdb.HookUpdateInput](ctx, in)
				if err != nil {
					return nil, err
				}
			}

			result, err = in.Next(ctx)
			return
		},

		Delete: func(ctx context.Context, in *gdb.HookDeleteInput) (result sql.Result, err error) {
			if h.cacheEvict {
				err = CleanCache[gdb.HookDeleteInput](ctx, in)
				if err != nil {
					return nil, err
				}
			}

			result, err = in.Next(ctx)
			return
		},
	}
}
