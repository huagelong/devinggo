// Package base
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package base

import (
	"context"
	"testing"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/stretchr/testify/assert"
)

type TestEntity struct {
	ID   int64
	Name string
}

func TestGenericServiceModel(t *testing.T) {
	ctx := context.Background()

	t.Run("with ModelFn", func(t *testing.T) {
		called := false
		service := &GenericService[TestEntity]{
			ModelFn: func(ctx context.Context) *gdb.Model {
				called = true
				return nil
			},
		}

		_ = service.Model(ctx)
		assert.True(t, called)
	})

	t.Run("without ModelFn", func(t *testing.T) {
		service := &GenericService[TestEntity]{}

		model := service.Model(ctx)
		assert.Nil(t, model)
	})
}

func TestGenericServiceStructure(t *testing.T) {
	service := &GenericService[TestEntity]{}

	assert.NotNil(t, service)
	assert.Nil(t, service.ModelFn)
}

func TestGenericServiceWithModelFn(t *testing.T) {
	ctx := context.Background()

	service := &GenericService[TestEntity]{
		ModelFn: func(ctx context.Context) *gdb.Model {
			return nil
		},
	}

	model := service.Model(ctx)
	assert.Nil(t, model)
}
