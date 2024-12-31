// Package websocket
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package websocket

import (
	"devinggo/modules/system/pkg/websocket/glob"
	"context"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
)

func PublishIdMessage(ctx context.Context, clientId string, msg *ClientIdWResponse) (err error) {
	glob.WithWsLog().Debug(ctx, "PublishIdMessage:", msg)
	toClient := clientManager.GetClient(clientId)
	if !g.IsEmpty(toClient) {
		SendToClientID(clientId, msg.WResponse)
		return
	}
	serverName := GetServerNameByClientId4Redis(ctx, clientId)
	j := gjson.NewWithTag(msg, "tag")
	if msg, err := j.ToJsonString(); err == nil {
		NewPubSub().PublishMessage(ctx, serverName, msg)
	} else {
		glob.WithWsLog().Warning(ctx, "SendMsg json encode error:", err)
	}
	return
}

func PublishTopicMessage(ctx context.Context, topic string, msg *TopicWResponse) (err error) {
	glob.WithWsLog().Debug(ctx, "PublishTopicMessage:", msg)
	serverNames := GetAllServerNameByTopic(ctx, topic)
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
