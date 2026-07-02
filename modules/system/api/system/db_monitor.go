package system

import (
	"devinggo/modules/system/model"
	"devinggo/modules/system/model/res"

	"github.com/gogf/gf/v2/frame/g"
)

type DbMonitorGroupsReq struct {
	g.Meta `path:"/dbMonitor/groups" method:"get" tags:"数据库监控" summary:"获取数据库监控分组." x-permission:"system:dbMonitor:groups"`
	model.AuthorHeader
}

type DbMonitorGroupsRes struct {
	g.Meta `mime:"application/json"`
	Data   []res.DbMonitorGroup `json:"data"`
}

type DbMonitorInfoReq struct {
	g.Meta `path:"/dbMonitor/monitor" method:"get" tags:"数据库监控" summary:"获取数据库监控数据." x-permission:"system:dbMonitor:monitor"`
	model.AuthorHeader
	GroupName string `json:"groupName" dc:"数据库分组" v:"required#数据库分组不能为空"`
}

type DbMonitorInfoRes struct {
	g.Meta `mime:"application/json"`
	Data   res.DbMonitorPayload `json:"data"`
}
