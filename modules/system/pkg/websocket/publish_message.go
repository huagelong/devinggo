// Package websocket
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package websocket

import (
	"context"

	"devinggo/modules/system/pkg/websocket/glob"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
)

// PublishSocketIdMessage 发布消息到指定SocketID（跨服务器）
func PublishSocketIdMessage(ctx context.Context, socketId string, msg *ClientIdWResponse) (err error) {
	glob.WithWsLog().Debug(ctx, "PublishSocketIdMessage:", msg)
	toClient := clientManager.GetClientBySocketID(socketId)
	if !g.IsEmpty(toClient) {
		SendToSocketID(socketId, msg.PusherResponse)
		return
	}
	serverName := GetServerNameBySocketId4Redis(ctx, socketId)
	j := gjson.NewWithTag(msg, "tag")
	if msg, err := j.ToJsonString(); err == nil {
		NewPubSub().PublishMessage(ctx, serverName, msg)
	} else {
		glob.WithWsLog().Warning(ctx, "SendMsg json encode error:", err)
	}
	return
}

// PublishChannelMessage 发布消息到频道（跨服务器）
func PublishChannelMessage(ctx context.Context, channel string, msg *TopicWResponse) (err error) {
	glob.WithWsLog().Debug(ctx, "PublishChannelMessage:", msg)
	serverNames := GetAllServerNameByChannel(ctx, channel)
	if g.IsEmpty(serverNames) {
		return
	}
	j := gjson.NewWithTag(msg, "tag")
	if msg, err := j.ToJsonString(); err == nil {
		for _, serverName := range serverNames {
			NewPubSub().PublishMessage(ctx, serverName, msg)
		}
	} else {
		glob.WithWsLog().Warning(ctx, "SendMsg json encode error:", err)
	}
	return
}

// PublishBroadcastMessage 发布广播消息（跨服务器）
func PublishBroadcastMessage(ctx context.Context, msg *BroadcastWResponse) (err error) {
	glob.WithWsLog().Debug(ctx, "PublishBroadcastMessage:", msg)
	serverNames := GetAllServerNames(ctx)
	if g.IsEmpty(serverNames) {
		return
	}
	j := gjson.NewWithTag(msg, "tag")
	if msg, err := j.ToJsonString(); err == nil {
		for _, serverName := range serverNames {
			NewPubSub().PublishMessage(ctx, serverName, msg)
		}
	} else {
		glob.WithWsLog().Warning(ctx, "SendMsg json encode error:", err)
	}
	return
}

// PublishTerminateConnectionMessage 发布终止连接消息（跨服务器）
func PublishTerminateConnectionMessage(ctx context.Context, socketId string, msg *TopicTerminateConnection) (err error) {
	glob.WithWsLog().Debug(ctx, "PublishTerminateConnectionMessage:", msg)
	serverName := GetServerNameBySocketId4Redis(ctx, socketId)
	if serverName == "" {
		glob.WithWsLog().Warning(ctx, "TerminateConnection: socket_id not found in Redis", socketId)
		return
	}

	// 只发送给目标服务器
	j := gjson.NewWithTag(msg, "tag")
	if msgStr, err := j.ToJsonString(); err == nil {
		NewPubSub().PublishMessage(ctx, serverName, msgStr)
	} else {
		glob.WithWsLog().Warning(ctx, "TerminateConnection json encode error:", err)
	}
	return
}
