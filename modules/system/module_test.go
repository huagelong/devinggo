// Package system
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE
package system

import (
	"context"
	_ "devinggo/internal/logic"
	_ "devinggo/internal/packed"
	_ "devinggo/modules/_/logic"
	"devinggo/modules/system/pkg/utils"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	"github.com/gogf/gf/v2/frame/g"
	"testing"
)

// 测试生成代码
func TestGenerateCode(t *testing.T) {
	ctx := context.Background()
	g.Log().Info(ctx, "TestGenerateCode")
	//t.Log("TestGenerateCode")
	g.Log().Info(ctx, utils.GetRootPath())
	t.Log("GenerateCode Success")
}
