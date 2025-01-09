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
					//send client id
					if j.Contains("id") {
						glob.WithWsLog().Debug(ctx, "SubscribeMessage client:", j.String())
						var clientWresponse *ClientIdWResponse
						if err := j.Scan(&clientWresponse); err == nil {
							clientId := gconv.String(j.Get("id"))
							SendToClientID(clientId, clientWresponse.WResponse)
						} else {
							glob.WithWsLog().Warning(ctx, "ClientIdWResponse parse error:", err)
						}
					}
					// send topic
					if j.Contains("topic") {
						glob.WithWsLog().Debug(ctx, "SubscribeMessage topic:", j.String())
						var topicWresponse *TopicWResponse
						if err := j.Scan(&topicWresponse); err == nil {
							topic := gconv.String(j.Get("topic"))
							SendToTopic(topic, topicWresponse.WResponse)
						} else {
							glob.WithWsLog().Warning(ctx, "TopicWResponse parse error:", err)
						}
					}
					// send Broadcast
					if j.Contains("broadcast") {
						glob.WithWsLog().Debug(ctx, "SubscribeMessage broadcast:", j.String())
						var broadcastWResponse *BroadcastWResponse
						if err := j.Scan(&broadcastWResponse); err == nil {
							SendToAll(broadcastWResponse.WResponse)
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
