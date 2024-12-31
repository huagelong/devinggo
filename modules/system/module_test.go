// Package system
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE
package system

import (
	_ "devinggo/internal/logic"
	_ "devinggo/internal/packed"
	_ "devinggo/modules/_/logic"
	"devinggo/modules/system/pkg/cache"
	_ "devinggo/modules/system/pkg/orm/driver"
	"context"
	"fmt"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	"github.com/gogf/gf/v2/frame/g"
	"testing"
	"time"
)

// 测试生成代码
func TestGenerateCode(t *testing.T) {
	ctx := context.Background()
	g.Log().Info(ctx, "TestGenerateCode")
	//t.Log("TestGenerateCode")
	//_, err := service.SettingGenerateTables().GenerateCode(ctx, []uint64{1})
	//if err != nil {
	//	t.Error(err)
	//}
	//localPath := gfile.MainPkgPath() + "/resource/public/uploads/image/20241220/d6ge8urn0800s3tvja.jpg"
	//storagePath := "image/20241220/d6ge8urn0800s3tvja.jpg"
	//err := upload.PutFromFile(ctx, localPath, storagePath)
	//if err != nil {
	//	t.Error(err)
	//}
	cache, err := cache.NewTagCache(ctx, g.Redis("cache"))
	if err != nil {
		t.Error(err)
	}
	err = cache.Set(ctx, "item1", "value1", []string{"tag1", "tag2"}, 3600*time.Second)
	if err != nil {
		t.Error(err)
	}
	val, err := cache.Get(context.Background(), "item1")
	if err != nil {
		t.Error(err)
	} else {
		fmt.Println("Got value:", val)
	}

	//err = cache.InvalidateTags(context.Background(), []string{"tag1"})
	//if err != nil {
	//	t.Error(err)
	//}
	//
	////// 再次获取缓存项，应该返回错误
	//_, err = cache.Get(context.Background(), "item1")
	//if err != nil {
	//	t.Error(err)
	//}

	t.Log("GenerateCode Success")
}
