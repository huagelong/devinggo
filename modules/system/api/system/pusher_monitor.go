package system

import (
	"devinggo/modules/system/model"
	"devinggo/modules/system/model/page"
	"devinggo/modules/system/model/res"

	"github.com/gogf/gf/v2/frame/g"
)

type PusherMonitorAppsReq struct {
	g.Meta `path:"/pusherMonitor/apps" method:"get" tags:"Pusher监控" summary:"获取 Pusher App 列表." x-permission:"system:pusherMonitor:apps"`
	model.AuthorHeader
}

type PusherMonitorAppsRes struct {
	g.Meta `mime:"application/json"`
	Data   []res.PusherMonitorAppItem `json:"data"`
}

type PusherMonitorOverviewReq struct {
	g.Meta `path:"/pusherMonitor/apps/{appId}/overview" method:"get" tags:"Pusher监控" summary:"获取 Pusher App 概览." x-permission:"system:pusherMonitor:overview"`
	model.AuthorHeader
	AppId string `json:"appId" in:"path" v:"required#appId不能为空"`
}

type PusherMonitorOverviewRes struct {
	g.Meta `mime:"application/json"`
	Data   res.PusherMonitorOverview `json:"data"`
}

type PusherMonitorChannelsReq struct {
	g.Meta `path:"/pusherMonitor/apps/{appId}/channels" method:"get" tags:"Pusher监控" summary:"获取频道列表." x-permission:"system:pusherMonitor:channels"`
	model.AuthorHeader
	page.PageReq
	AppId       string `json:"appId" in:"path" v:"required#appId不能为空"`
	Keyword     string `json:"keyword"`
	ChannelType string `json:"channelType"`
}

type PusherMonitorChannelsRes struct {
	g.Meta `mime:"application/json"`
	page.PageRes
	Items []res.PusherMonitorChannelItem `json:"items"`
}

type PusherMonitorChannelDetailReq struct {
	g.Meta `path:"/pusherMonitor/apps/{appId}/channels/{channelName}/detail" method:"get" tags:"Pusher监控" summary:"获取频道详情." x-permission:"system:pusherMonitor:channels"`
	model.AuthorHeader
	AppId       string `json:"appId" in:"path" v:"required#appId不能为空"`
	ChannelName string `json:"channelName" in:"path" v:"required#channelName不能为空"`
}

type PusherMonitorChannelDetailRes struct {
	g.Meta `mime:"application/json"`
	Data   res.PusherMonitorChannelDetail `json:"data"`
}

type PusherMonitorConnectionsReq struct {
	g.Meta `path:"/pusherMonitor/apps/{appId}/connections" method:"get" tags:"Pusher监控" summary:"获取连接列表." x-permission:"system:pusherMonitor:connections"`
	model.AuthorHeader
	page.PageReq
	AppId    string `json:"appId" in:"path" v:"required#appId不能为空"`
	SocketId string `json:"socketId"`
	UserId   string `json:"userId"`
	Channel  string `json:"channel"`
}

type PusherMonitorConnectionsRes struct {
	g.Meta `mime:"application/json"`
	page.PageRes
	Items []res.PusherMonitorConnectionItem `json:"items"`
}

type PusherMonitorTerminateConnectionsReq struct {
	g.Meta `path:"/pusherMonitor/apps/{appId}/connections/terminate-batch" method:"post" tags:"Pusher监控" summary:"按 socket_id 批量断连." x-permission:"system:pusherMonitor:terminate"`
	model.AuthorHeader
	AppId     string   `json:"appId" in:"path" v:"required#appId不能为空"`
	SocketIds []string `json:"socketIds" v:"required|min-length:1#socketIds不能为空"`
}

type PusherMonitorTerminateConnectionsRes struct {
	g.Meta `mime:"application/json"`
	Data   res.PusherMonitorTerminateResult `json:"data"`
}

type PusherMonitorTerminateUsersReq struct {
	g.Meta `path:"/pusherMonitor/apps/{appId}/users/terminate-batch" method:"post" tags:"Pusher监控" summary:"按 user_id 批量断连." x-permission:"system:pusherMonitor:terminate"`
	model.AuthorHeader
	AppId   string   `json:"appId" in:"path" v:"required#appId不能为空"`
	UserIds []string `json:"userIds" v:"required|min-length:1#userIds不能为空"`
}

type PusherMonitorTerminateUsersRes struct {
	g.Meta `mime:"application/json"`
	Data   res.PusherMonitorTerminateResult `json:"data"`
}

type PusherMonitorDebugEventReq struct {
	g.Meta `path:"/pusherMonitor/apps/{appId}/events/debug" method:"post" tags:"Pusher监控" summary:"调试触发事件." x-permission:"system:pusherMonitor:debugEvent"`
	model.AuthorHeader
	AppId    string   `json:"appId" in:"path" v:"required#appId不能为空"`
	Name     string   `json:"name" v:"required#事件名不能为空"`
	Channels []string `json:"channels" v:"required|min-length:1#channels不能为空"`
	Data     string   `json:"data" v:"required#事件数据不能为空"`
	SocketId string   `json:"socketId"`
}

type PusherMonitorDebugEventRes struct {
	g.Meta `mime:"application/json"`
	Data   res.PusherMonitorDebugEventResult `json:"data"`
}
