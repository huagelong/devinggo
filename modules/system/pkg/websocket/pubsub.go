// Package websocket
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package websocket

import (
	"context"

	"devinggo/modules/system/pkg/redispubsub"
	"devinggo/modules/system/pkg/utils"
	"devinggo/modules/system/pkg/websocket/glob"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/util/gconv"
)

type sPubSub struct {
	PubSub *redispubsub.PubSub
}

func NewPubSub() *sPubSub {
	pubsub := redispubsub.New(redispubsub.WithRedisGroup("websocket"), redispubsub.WithLoggerName("ws"))
	//defer pubsub.Close()
	return &sPubSub{PubSub: pubsub}
}

func (s *sPubSub) SubscribeMessage(ctx context.Context, serverName string) (err error) {
	err = s.PubSub.Subscribe(ctx, serverName)
	if err != nil {
		return
	}
	utils.SafeGo(ctx, func(ctx context.Context) {
		func() {
			for {
				select {
				case msg := <-s.PubSub.Messages():
					j := gjson.New(msg.Payload)
					// terminate connection（只包含 socketId 的消息）
					if j.Contains("socketId") && !j.Contains("pusherResponse") {
						glob.WithWsLog().Debug(ctx, "SubscribeMessage terminate connection:", j.String())
						var msgData *TopicTerminateConnection
						if err := j.Scan(&msgData); err == nil {
							TerminateLocalClient(msgData.SocketID)
						} else {
							glob.WithWsLog().Warning(ctx, "TopicTerminateConnection parse error:", err)
						}
					}
					// send socket id（包含 pusherResponse 的消息）
					if j.Contains("pusherResponse") {
						glob.WithWsLog().Debug(ctx, "SubscribeMessage socketId:", j.String())
						var msgData *ClientIdWResponse
						if err := j.Scan(&msgData); err == nil {
							socketId := gconv.String(j.Get("socketId"))
							SendToSocketID(socketId, msgData.PusherResponse)
						} else {
							glob.WithWsLog().Warning(ctx, "ClientIdWResponse parse error:", err)
						}
					}
					// send channel
					if j.Contains("topic") {
						glob.WithWsLog().Debug(ctx, "SubscribeMessage channel:", j.String())
						var msgData *TopicWResponse
						if err := j.Scan(&msgData); err == nil {
							SendToChannelWithExclude(msgData.Topic, msgData.PusherResponse, msgData.ExcludeSocketID)
						} else {
							glob.WithWsLog().Warning(ctx, "TopicWResponse parse error:", err)
						}
					}
					// send Broadcast
					if j.Contains("broadcast") {
						glob.WithWsLog().Debug(ctx, "SubscribeMessage broadcast:", j.String())
						var msgData *BroadcastWResponse
						if err := j.Scan(&msgData); err == nil {
							SendToAll(msgData.PusherResponse)
						} else {
							glob.WithWsLog().Warning(ctx, "broadcastWResponse parse error:", err)
						}
					}
				case <-s.PubSub.Unsubscribe():
					glob.WithWsLog().Debug(ctx, "SubscribeMessage unsubscribe")
					return
				}
			}
		}()
	})
	return
}

func (s *sPubSub) PublishMessage(ctx context.Context, serverName string, msg string) error {
	return s.PubSub.Publish(ctx, serverName, msg)
}
