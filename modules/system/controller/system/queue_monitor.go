package system

import (
	"context"

	api "devinggo/modules/system/api/system"
	"devinggo/modules/system/controller/base"
	"devinggo/modules/system/service"
)

var (
	QueueMonitorController = queueMonitorController{}
)

type queueMonitorController struct {
	base.BaseController
}

func (c *queueMonitorController) Queues(ctx context.Context, in *api.QueueMonitorQueuesReq) (out *api.QueueMonitorQueuesRes, err error) {
	out = &api.QueueMonitorQueuesRes{}
	out.Data, err = service.QueueMonitor().ListQueues(ctx)
	return
}

func (c *queueMonitorController) Overview(ctx context.Context, in *api.QueueMonitorOverviewReq) (out *api.QueueMonitorOverviewRes, err error) {
	out = &api.QueueMonitorOverviewRes{}
	data, err := service.QueueMonitor().GetOverview(ctx, in.Queue)
	if err != nil {
		return nil, err
	}
	if data != nil {
		out.Data = *data
	}
	return
}

func (c *queueMonitorController) Tasks(ctx context.Context, in *api.QueueMonitorTasksReq) (out *api.QueueMonitorTasksRes, err error) {
	out = &api.QueueMonitorTasksRes{}
	out.Items, out.PageInfo.TotalCount, err = service.QueueMonitor().ListTasks(ctx, in.Queue, in.State, &in.PageReq)
	if err != nil {
		return nil, err
	}
	out.PageRes.Pack(&in.PageReq, out.PageInfo.TotalCount)
	return
}

func (c *queueMonitorController) SchedulerEntries(ctx context.Context, in *api.QueueMonitorSchedulerEntriesReq) (out *api.QueueMonitorSchedulerEntriesRes, err error) {
	out = &api.QueueMonitorSchedulerEntriesRes{}
	out.Data, err = service.QueueMonitor().ListSchedulerEntries(ctx)
	return
}
