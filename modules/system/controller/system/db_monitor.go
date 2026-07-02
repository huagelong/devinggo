package system

import (
	"context"

	"devinggo/modules/system/api/system"
	"devinggo/modules/system/controller/base"
	"devinggo/modules/system/service"
)

var (
	DbMonitorController = dbMonitorController{}
)

type dbMonitorController struct {
	base.BaseController
}

func (c *dbMonitorController) Groups(ctx context.Context, in *system.DbMonitorGroupsReq) (out *system.DbMonitorGroupsRes, err error) {
	out = &system.DbMonitorGroupsRes{}
	out.Data, err = service.DbMonitor().ListGroups(ctx)
	return
}

func (c *dbMonitorController) Monitor(ctx context.Context, in *system.DbMonitorInfoReq) (out *system.DbMonitorInfoRes, err error) {
	out = &system.DbMonitorInfoRes{}
	data, err := service.DbMonitor().GetMonitor(ctx, in.GroupName)
	if err != nil {
		return nil, err
	}
	if data != nil {
		out.Data = *data
	}
	return
}
