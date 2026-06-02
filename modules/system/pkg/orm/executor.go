// Package orm
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package orm

import (
	"context"

	"devinggo/modules/system/pkg/utils"

	"github.com/gogf/gf/v2/database/gdb"
)

// Executor 查询执行器，提供便捷的查询执行方法
type Executor[T any] struct {
	model *gdb.Model
}

// NewExecutor 创建新的查询执行器
func NewExecutor[T any](m *gdb.Model) *Executor[T] {
	return &Executor[T]{model: m}
}

// One 查询单条数据
func (e *Executor[T]) One(ctx context.Context) (result T, err error) {
	err = e.model.Ctx(ctx).Scan(&result)
	if utils.IsError(err) {
		return result, err
	}
	return result, nil
}

// All 查询所有数据
func (e *Executor[T]) All(ctx context.Context) (results []T, err error) {
	err = e.model.Ctx(ctx).Scan(&results)
	if utils.IsError(err) {
		return nil, err
	}
	return results, nil
}

// AllWithCount 查询所有数据并返回总数
func (e *Executor[T]) AllWithCount(ctx context.Context) (results []T, total int, err error) {
	err = e.model.Ctx(ctx).ScanAndCount(&results, &total, false)
	if utils.IsError(err) {
		return nil, 0, err
	}
	return results, total, nil
}

// PagedList 分页查询数据
func (e *Executor[T]) PagedList(ctx context.Context, page, pageSize int) (results []T, total int, err error) {
	err = e.model.Ctx(ctx).Page(page, pageSize).ScanAndCount(&results, &total, false)
	if utils.IsError(err) {
		return nil, 0, err
	}
	return results, total, nil
}

// Count 统计数量
func (e *Executor[T]) Count(ctx context.Context) (count int, err error) {
	count, err = e.model.Ctx(ctx).Count()
	if utils.IsError(err) {
		return 0, err
	}
	return count, nil
}

// Exists 判断数据是否存在
func (e *Executor[T]) Exists(ctx context.Context) (exists bool, err error) {
	count, err := e.Count(ctx)
	return count > 0, err
}
