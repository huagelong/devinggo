// Package user_agent
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package user_agent

import (
	"devinggo/modules/system/pkg/utils/request"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/mileusna/useragent"
)

type UserAgent struct {
	Browser string
	Os      string
	Device  string
}

func GetUserAgent(ctx context.Context) (userAgent UserAgent) {
	r := request.GetHttpRequest(ctx)
	if r == nil {
		g.Log().Warningf(ctx, "ctx not http request")
		return
	}
	ua := useragent.Parse(r.UserAgent())
	userAgent.Browser = ua.Name
	userAgent.Os = ua.OS + " " + ua.OSVersion
	userAgent.Device = ua.Device
	return
}
