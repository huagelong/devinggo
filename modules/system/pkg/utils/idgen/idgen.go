// Package idgen
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package idgen

import (
	"context"
	"devinggo/modules/system/pkg/utils/config"

	"github.com/gogf/gf/v2/util/gconv"
	"github.com/yitter/idgenerator-go/idgen"
)

func NextId(ctx context.Context) int64 {
	workerIdInt := config.GetConfigInt(ctx, "settings.snowflake.workerId")
	workerId := gconv.Uint16(workerIdInt)
	var options = idgen.NewIdGeneratorOptions(workerId)
	options.WorkerIdBitLength = 10
	idgen.SetIdGenerator(options)
	return idgen.NextId()
}
