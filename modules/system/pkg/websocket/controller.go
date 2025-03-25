// Package websocket
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package websocket

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

func IdMessageController(ctx context.Context, client *Client, req *Request) {
	if _, ok := req.Data["toSessionId"]; !ok {
		client.ResponseFail(ctx, IdMessage, req.RequestId, "toSessionId miss")
	}
	toId := gconv.String(req.Data["toSessionId"])
	if g.IsEmpty(clientManager.GetClient(toId)) {
		err := PublishIdMessage(ctx, toId, &ClientIdWResponse{
			ID: toId,
			WResponse: &WResponse{
				BindEvent: req.BindEvent,
				Event:     IdMessage,
				Code:      200,
				Data:      req.Data,
				RequestId: req.RequestId,
			},
		})
		if err != nil {
			client.ResponseFail(ctx, IdMessage, req.RequestId, err.Error())
			return
		}
	} else {
		if g.IsEmpty(GetServerNameByClientId4Redis(ctx, toId)) {
			client.ResponseFail(ctx, IdMessage, req.RequestId, "client miss")
			return
		} else {
			SendToClientID(toId, &WResponse{
				BindEvent: req.BindEvent,
				Event:     IdMessage,
				Data:      req.Data,
				Code:      200,
				RequestId: req.RequestId,
			})
		}
	}
	client.ResponseSuccess(ctx, IdMessage, req.RequestId)
}

func PublishController(ctx context.Context, client *Client, req *Request) {
	if _, ok := req.Data["topic"]; !ok {
		client.ResponseFail(ctx, Publish, req.RequestId, "topic miss")
	}
	topic := gconv.String(req.Data["topic"])
	if isTopicExist(ctx, topic) {
		err := PublishTopicMessage(ctx, topic, &TopicWResponse{
			Topic: topic,
			WResponse: &WResponse{
				Event:     Publish,
				Data:      req.Data,
				RequestId: req.RequestId,
				Code:      200,
			},
		})
		if err != nil {
			client.ResponseFail(ctx, Publish, req.RequestId, err.Error())
			return
		}
		client.ResponseSuccess(ctx, Publish, req.RequestId)
	} else {
		client.ResponseFail(ctx, Publish, req.RequestId, "tag miss")
	}

}

func BroadcastMessageController(ctx context.Context, client *Client, req *Request) {
	err := PublishBroadcastMessage(ctx, &BroadcastWResponse{
		Broadcast: "1",
		WResponse: &WResponse{
			BindEvent: req.BindEvent,
			Event:     BroadcastMessage,
			Data:      req.Data,
			RequestId: req.RequestId,
			Code:      200,
		},
	})
	if err != nil {
		client.ResponseFail(ctx, BroadcastMessage, req.RequestId, err.Error())
		return
	}
	client.ResponseSuccess(ctx, BroadcastMessage, req.RequestId)
}

// SubscribeController join topic
func SubscribeController(ctx context.Context, client *Client, req *Request) {
	if _, ok := req.Data["topic"]; !ok {
		client.ResponseFail(ctx, Subscribe, req.RequestId, "topic miss")
	}

	topic := gconv.String(req.Data["topic"])

	if !client.topics.Contains(topic) {
		client.topics.Append(topic)
	}

	err := JoinTopic4Redis(ctx, client.ID, topic)
	if err != nil {
		client.ResponseFail(ctx, Subscribe, req.RequestId, err.Error())
		return
	}
	client.ResponseSuccess(ctx, Subscribe, req.RequestId)
}

// UnsubscribeController 退出
func UnsubscribeController(ctx context.Context, client *Client, req *Request) {
	if _, ok := req.Data["topic"]; !ok {
		client.ResponseFail(ctx, Unsubscribe, req.RequestId, "topic miss")
	}
	client.ResponseSuccess(ctx, Unsubscribe, req.RequestId)
	topic := gconv.String(req.Data["topic"])
	if client.topics.Contains(topic) {
		client.topics.RemoveValue(topic)
	}
	QuitTopic4Redis(ctx, client.ID, topic)
}

func PingController(ctx context.Context, client *Client, req *Request) {
	currentTime := int64(gtime.Now().Unix())
	client.Heartbeat(currentTime)
	UpdateClientIdHeartbeatTime4Redis(ctx, client.ID, currentTime)
	client.ResponseSuccess(ctx, Ping, req.RequestId)
}

func PongController(ctx context.Context, client *Client, req *Request) {
	currentTime := int64(gtime.Now().Unix())
	client.Heartbeat(currentTime)
	UpdateClientIdHeartbeatTime4Redis(ctx, client.ID, currentTime)
	client.ResponseSuccess(ctx, Pong, req.RequestId)
}

func CloseController(ctx context.Context, client *Client, req *Request) {
	client.ResponseSuccess(ctx, Close, req.RequestId)
	client.close(ctx)
}
