// Package websocket
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package websocket

import (
	"context"
	"devinggo/modules/system/pkg/utils"
	"devinggo/modules/system/pkg/websocket/glob"
	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gorilla/websocket"
	"runtime/debug"
)

const (
	// 用户连接超时时间
	heartbeatExpirationTime = 6 * 60
)

// Client 客户端连接
type Client struct {
	Addr          string          // 客户端地址
	ID            string          // 连接唯一标识
	Socket        *websocket.Conn // 用户连接
	Send          chan *WResponse // 待发送的数据
	SendClose     bool            // 发送是否关闭
	FirstTime     int64           // 首次连接事件
	HeartbeatTime int64           // 用户上次心跳时间
	LoginTime     int64           // 登录时间 登录以后才有
	ServerName    string
	topics        garray.StrArray // 标签
}

// NewClient 初始化
func NewClient(addr string, clientId string, socket *websocket.Conn, firstTime int64) (client *Client) {
	client = &Client{
		Addr:          addr,
		ID:            clientId,
		Socket:        socket,
		Send:          make(chan *WResponse, 100),
		SendClose:     false,
		FirstTime:     firstTime,
		HeartbeatTime: firstTime,
	}
	return
}

// 读取客户端数据
func (c *Client) read(ctx context.Context) {

	defer func() {
		if r := recover(); r != nil {
			glob.WithWsLog().Warning(ctx, "read error:", string(debug.Stack()), r)
		}
	}()

	defer func() {
		glob.WithWsLog().Debug(ctx, "read conn close ID:", c.ID)
		c.close(ctx)
	}()

	for {
		_, message, err := c.Socket.ReadMessage()
		if err != nil {
			glob.WithWsLog().Warning(ctx, "ReadMessage error:", err)
			return
		}
		if !g.IsEmpty(message) {
			ProcessData(ctx, c, message)
		}
	}
}

// 向客户端写数据
func (c *Client) write(ctx context.Context) {
	defer func() {
		if r := recover(); r != nil {
			glob.WithWsLog().Warning(ctx, "write error:", string(debug.Stack()), r)
		}
	}()
	defer func() {
		glob.WithWsLog().Debug(ctx, "write conn close ID:", c.ID)
		c.close(ctx)
	}()
	for {
		select {
		case message, ok := <-c.Send:
			if !ok {
				// 发送数据错误 关闭连接
				return
			}
			glob.WithWsLog().Debug(ctx, "response:", message)

			utils.SafeGo(ctx, func(ctx context.Context) {
				c.updateMsgLog(ctx, message)
			})
			err := c.Socket.WriteJSON(message)
			if err != nil {
				glob.WithWsLog().Warning(ctx, "WriteJSON error:", err)
			}
		}
	}
}

// SendMsg 发送数据
func (c *Client) SendMsg(ctx context.Context, msg *WResponse) {
	if c == nil || c.SendClose {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			glob.WithWsLog().Warning(ctx, "SendMsg error:", string(debug.Stack()), r)
		}
	}()

	c.Send <- msg
}

func (c *Client) Success(ctx context.Context, req *Request, data ...map[string]interface{}) {
	dataRs := g.Map{}
	if !g.IsEmpty(data) {
		dataRs = data[0]
	}
	c.SendMsg(ctx, &WResponse{
		Event:     req.Event,
		Code:      200,
		Data:      dataRs,
		RequestId: req.RequestId,
	})
}

func (c *Client) Fail(ctx context.Context, req *Request, message string, data ...map[string]interface{}) {
	dataRs := g.Map{}
	if !g.IsEmpty(data) {
		dataRs = data[0]
	}
	c.SendMsg(ctx, &WResponse{
		Event:     req.Event,
		Message:   message,
		Code:      500,
		Data:      dataRs,
		RequestId: req.RequestId,
	})
}

func (c *Client) ResponseSuccess(ctx context.Context, event string, requestId string, data ...map[string]interface{}) {
	dataRs := g.Map{}
	if !g.IsEmpty(data) {
		dataRs = data[0]
	}
	c.SendMsg(ctx, &WResponse{
		Event:     event,
		Code:      200,
		Data:      dataRs,
		RequestId: requestId,
		CallBack:  1,
	})
}

func (c *Client) ResponseFail(ctx context.Context, event string, requestId string, message string, data ...map[string]interface{}) {
	dataRs := g.Map{}
	if !g.IsEmpty(data) {
		dataRs = data[0]
	}
	c.SendMsg(ctx, &WResponse{
		Event:     event,
		Message:   message,
		Code:      500,
		Data:      dataRs,
		RequestId: requestId,
		CallBack:  1,
	})
}

// Heartbeat 心跳更新
func (c *Client) Heartbeat(currentTime int64) {
	c.HeartbeatTime = currentTime
	return
}

// IsHeartbeatTimeout 心跳是否超时
func (c *Client) IsHeartbeatTimeout(currentTime int64) (timeout bool) {
	if c.HeartbeatTime+heartbeatExpirationTime <= currentTime {
		timeout = true
	}
	return
}

// 关闭客户端
func (c *Client) close(ctx context.Context) {
	if c.SendClose {
		return
	}
	//删除客户端数据
	utils.SafeGo(gctx.GetInitCtx(), func(ctx context.Context) {
		ClearClientId4Redis(ctx, c.ID)
	})
	c.SendClose = true
	clientManager.Disconnect <- c
	c.Socket.Close()
	close(c.Send)
}
