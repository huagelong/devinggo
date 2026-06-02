// Package worker
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package worker

import (
	// 导入worker和cron包，触发它们的init自动注册
	_ "devinggo/modules/system/worker/cron"
	_ "devinggo/modules/system/worker/server"
)
