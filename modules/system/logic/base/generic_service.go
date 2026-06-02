// Package base
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package base

import (
	"context"

	"devinggo/modules/system/pkg/utils"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

// GenericService 通用 Service 基类，使用泛型消除重复代码
// T 为返回结果类型（如 *res.SystemPost）
type GenericService[T any] struct {
	BaseService
	// ModelFn 返回该 service 的 Model 方法
	ModelFn func(ctx context.Context) *gdb.Model
}

// Model 返回数据库 Model
func (s *GenericService[T]) Model(ctx context.Context) *gdb.Model {
	if s.ModelFn != nil {
		return s.ModelFn(ctx)
	}
	return nil
}

// GetById 根据 ID 获取单条记录
// 消除 15+ 个 service 中的重复实现
func (s *GenericService[T]) GetById(ctx context.Context, id int64) (res *T, err error) {
	err = s.Model(ctx).Where("id", id).Scan(&res)
	return
}

// ChangeStatus 修改状态
func (s *GenericService[T]) ChangeStatus(ctx context.Context, id int64, status int) (err error) {
	_, err = s.Model(ctx).Data(g.Map{"status": status}).Where("id", id).Update()
	return
}

// Recovery 恢复（软删除恢复）
func (s *GenericService[T]) Recovery(ctx context.Context, ids []int64) (err error) {
	_, err = s.Model(ctx).Unscoped().WhereIn("id", ids).Update(g.Map{"deleted_at": nil})
	return
}

// Delete 删除（软删除）
func (s *GenericService[T]) Delete(ctx context.Context, ids []int64) (err error) {
	_, err = s.Model(ctx).WhereIn("id", ids).Delete()
	return
}

// RealDelete 真删除
func (s *GenericService[T]) RealDelete(ctx context.Context, ids []int64) (err error) {
	_, err = s.Model(ctx).Unscoped().WhereIn("id", ids).Delete()
	return
}

// NumberOperation 字段数值操作
func (s *GenericService[T]) NumberOperation(ctx context.Context, id int64, numberName string, numberValue int) (err error) {
	_, err = s.Model(ctx).Data(g.Map{numberName: numberValue}).Where("id", id).Update()
	return
}

// GetPageList 通用分页查询
// entities 为实体类型（如 *entity.SystemPost）
// 返回转换后的结果类型 T
func (s *GenericService[T]) GetPageList(
	ctx context.Context,
	searchModel *gdb.Model,
	entities interface{},
) (results []*T, total int, err error) {
	err = searchModel.ScanAndCount(entities, &total, false)
	if utils.IsError(err) {
		return nil, 0, err
	}
	results = make([]*T, 0)
	if !g.IsEmpty(entities) {
		if err = gconv.Structs(entities, &results); utils.IsError(err) {
			return nil, 0, err
		}
	}
	return
}

// GetList 通用列表查询（不分页）
func (s *GenericService[T]) GetList(ctx context.Context, searchModel *gdb.Model) (results []*T, err error) {
	err = searchModel.Scan(&results)
	return
}
