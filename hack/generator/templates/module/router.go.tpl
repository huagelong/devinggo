// Package {{.moduleName}}
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package {{.moduleName}}

import (
	"devinggo/modules/{{.moduleName}}/controller/{{.moduleName}}"

	"github.com/gogf/gf/v2/net/ghttp"
)

func BindController(group *ghttp.RouterGroup) {
	group.Bind(
		{{.moduleName}}.TestController,
	)
}
