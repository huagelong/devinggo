package system

import (
	"context"

	api "devinggo/modules/system/api/system"
	"devinggo/modules/system/controller/base"
	"devinggo/modules/system/service"
)

var (
	PusherMonitorController = pusherMonitorController{}
)

type pusherMonitorController struct {
	base.BaseController
}

func (c *pusherMonitorController) Apps(ctx context.Context, in *api.PusherMonitorAppsReq) (out *api.PusherMonitorAppsRes, err error) {
	out = &api.PusherMonitorAppsRes{}
	out.Data, err = service.PusherMonitor().ListApps(ctx)
	return
}

func (c *pusherMonitorController) Overview(ctx context.Context, in *api.PusherMonitorOverviewReq) (out *api.PusherMonitorOverviewRes, err error) {
	out = &api.PusherMonitorOverviewRes{}
	data, err := service.PusherMonitor().GetOverview(ctx, in.AppId)
	if err != nil {
		return nil, err
	}
	if data != nil {
		out.Data = *data
	}
	return
}

func (c *pusherMonitorController) Channels(ctx context.Context, in *api.PusherMonitorChannelsReq) (out *api.PusherMonitorChannelsRes, err error) {
	out = &api.PusherMonitorChannelsRes{}
	out.Items, out.PageInfo.TotalCount, err = service.PusherMonitor().ListChannels(ctx, in.AppId, in.Keyword, in.ChannelType, &in.PageReq)
	if err != nil {
		return nil, err
	}
	out.PageRes.Pack(&in.PageReq, out.PageInfo.TotalCount)
	return
}

func (c *pusherMonitorController) ChannelDetail(ctx context.Context, in *api.PusherMonitorChannelDetailReq) (out *api.PusherMonitorChannelDetailRes, err error) {
	out = &api.PusherMonitorChannelDetailRes{}
	data, err := service.PusherMonitor().GetChannelDetail(ctx, in.AppId, in.ChannelName)
	if err != nil {
		return nil, err
	}
	if data != nil {
		out.Data = *data
	}
	return
}

func (c *pusherMonitorController) Connections(ctx context.Context, in *api.PusherMonitorConnectionsReq) (out *api.PusherMonitorConnectionsRes, err error) {
	out = &api.PusherMonitorConnectionsRes{}
	out.Items, out.PageInfo.TotalCount, err = service.PusherMonitor().ListConnections(ctx, in.AppId, in.SocketId, in.UserId, in.Channel, &in.PageReq)
	if err != nil {
		return nil, err
	}
	out.PageRes.Pack(&in.PageReq, out.PageInfo.TotalCount)
	return
}

func (c *pusherMonitorController) TerminateConnections(ctx context.Context, in *api.PusherMonitorTerminateConnectionsReq) (out *api.PusherMonitorTerminateConnectionsRes, err error) {
	out = &api.PusherMonitorTerminateConnectionsRes{}
	data, err := service.PusherMonitor().TerminateConnections(ctx, in.AppId, in.SocketIds)
	if err != nil {
		return nil, err
	}
	if data != nil {
		out.Data = *data
	}
	return
}

func (c *pusherMonitorController) TerminateUsers(ctx context.Context, in *api.PusherMonitorTerminateUsersReq) (out *api.PusherMonitorTerminateUsersRes, err error) {
	out = &api.PusherMonitorTerminateUsersRes{}
	data, err := service.PusherMonitor().TerminateUsers(ctx, in.AppId, in.UserIds)
	if err != nil {
		return nil, err
	}
	if data != nil {
		out.Data = *data
	}
	return
}

func (c *pusherMonitorController) DebugEvent(ctx context.Context, in *api.PusherMonitorDebugEventReq) (out *api.PusherMonitorDebugEventRes, err error) {
	out = &api.PusherMonitorDebugEventRes{}
	data, err := service.PusherMonitor().DebugEvent(ctx, in.AppId, in.Name, in.Channels, in.Data, in.SocketId)
	if err != nil {
		return nil, err
	}
	if data != nil {
		out.Data = *data
	}
	return
}
