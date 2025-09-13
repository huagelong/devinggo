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

type HookOptions struct {
	AutoCreatedUpdatedBy *bool
	CacheEvict           *bool
	UserRelate           *bool
	Params               interface{}
}

func Bind(optionsFirst ...*HookOptions) gdb.HookHandler {
	defaultAutoCreatedUpdatedBy := true
	defaultCacheEvict := true
	defaultUserRelate := true
	var options = HookOptions{
		AutoCreatedUpdatedBy: &defaultAutoCreatedUpdatedBy,
		CacheEvict:           &defaultCacheEvict,
		UserRelate:           &defaultUserRelate,
	}

	if !g.IsEmpty(optionsFirst) {
		optionsFirstTmp := optionsFirst[0]
		if !g.IsNil(optionsFirstTmp.AutoCreatedUpdatedBy) {
			options.AutoCreatedUpdatedBy = optionsFirstTmp.AutoCreatedUpdatedBy
		}
		if !g.IsNil(optionsFirstTmp.CacheEvict) {
			options.CacheEvict = optionsFirstTmp.CacheEvict
		}
		if !g.IsNil(optionsFirstTmp.UserRelate) {
			options.UserRelate = optionsFirstTmp.UserRelate
		}

		if !g.IsNil(optionsFirstTmp.Params) {
			options.Params = optionsFirstTmp.Params
		}
	}

	return gdb.HookHandler{
		Select: func(ctx context.Context, in *gdb.HookSelectInput) (result gdb.Result, err error) {
			result, err = in.Next(ctx)
			if err != nil {
				return result, err
			}
			if *options.UserRelate && !g.IsEmpty(options.Params) {
				return UserRelate(ctx, result, options.Params.([]string))
			}
			return
		},
		Insert: func(ctx context.Context, in *gdb.HookInsertInput) (result sql.Result, err error) {
			if *options.AutoCreatedUpdatedBy {
				err = AutoCreatedUpdatedByInsert(ctx, in)
				if err != nil {
					return nil, err
				}
			}

			if *options.CacheEvict {
				err = CleanCache[gdb.HookInsertInput](ctx, in)
				if err != nil {
					return nil, err
				}
			}
			result, err = in.Next(ctx)
			if err != nil {
				g.Log().Debug(ctx, "in:", in)
				g.Log().Debug(ctx, "Insert:", err)
			}
			return
		},

		Update: func(ctx context.Context, in *gdb.HookUpdateInput) (result sql.Result, err error) {
			if *options.AutoCreatedUpdatedBy {
				err = AutoCreatedUpdatedByUpdatefunc(ctx, in)
				if err != nil {
					return nil, err
				}
			}

			if *options.CacheEvict {
				err = CleanCache[gdb.HookUpdateInput](ctx, in)
				if err != nil {
					return nil, err
				}
			}

			result, err = in.Next(ctx)
			return
		},

		Delete: func(ctx context.Context, in *gdb.HookDeleteInput) (result sql.Result, err error) {

			if *options.CacheEvict {
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
