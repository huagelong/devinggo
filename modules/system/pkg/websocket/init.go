// Package websocket
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package websocket

import (
	"context"
	"crypto/rand"
	"fmt"
	"math/big"
	"net/http"
	"time"

	"devinggo/modules/system/model"
	"devinggo/modules/system/pkg/utils"
	"devinggo/modules/system/pkg/websocket/glob"
	"devinggo/modules/system/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gorilla/websocket"
)

var (
	clientManager = NewClientManager() // 管理者
	//connInstance   *websocket.Conn
	//once           sync.Once
	SESSION_ID_KEY = "WS_SESSION_ID"
)

func StartWebSocket(ctx context.Context, serverName string) {
	glob.WithWsLog().Debug(ctx, "start：WebSocket")
	clientManager.SetServerName(serverName)
	utils.SafeGo(ctx, func(ctx context.Context) {
		clientManager.start(ctx)
	})
	utils.SafeGo(ctx, func(ctx context.Context) {
		clientManager.cronJob(ctx)
	})
	utils.SafeGo(ctx, func(ctx context.Context) {
		NewPubSub().SubscribeMessage(ctx, serverName)
	})
}

func parseSessionId(r *ghttp.Request) (sessionId string, err error) {
	ctx := r.GetCtx()
	sessionIdTmp := r.GetQuery("sessionId")
	token := r.GetQuery("token")
	if g.IsEmpty(sessionIdTmp) {

		if g.IsEmpty(token) {
			return "", nil
		}

		claims, err := service.Token().ParseToken(ctx, token.String())
		if err != nil {
			return "", err
		}
		data := claims.Data
		if g.IsEmpty(data) {
			return "", nil
		} else {
			var user *model.Identity
			data := claims.Data
			err = gconv.Scan(data, &user)
			if err != nil {
				return "", err
			}
			if g.IsEmpty(user) {
				return "", nil
			} else {
				return gconv.String(user.Id), nil
			}
		}
	} else {
		return sessionIdTmp.String(), nil
	}
}

func GetConnection(r *ghttp.Request) (conn *websocket.Conn, err error) {
	var upGrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			glob.WithWsLog().Debug(r.Context(), "r.Host:", r.Host)
			return true
		},
	}
	ctx := r.GetCtx()
	conn, err = upGrader.Upgrade(r.Response.Writer, r.Request, nil)
	if err != nil {
		glob.WithWsLog().Errorf(ctx, "ws Upgrade error:%v", err)
		return
	}
	return
}

func WsPage(r *ghttp.Request) {
	ctx := r.GetCtx()
	currentTime := int64(gtime.Now().Unix())

	// ⚠️ v8.3.0要求：验证协议版本
	// Pusher WebSocket 协议版本，默认为 7
	protocol := r.GetQuery("protocol").String()
	if protocol == "" {
		// 兼容没有指定 protocol 参数的客户端，默认使用协议版本 7
		protocol = "7"
		glob.WithWsLog().Debugf(ctx, "No protocol specified, defaulting to version 7")
	}
	if protocol != "7" {
		glob.WithWsLog().Warning(ctx, "Unsupported protocol version:", protocol)
		r.Response.WriteStatus(400)
		r.Response.WriteJson(g.Map{"error": "Unsupported protocol version"})
		return
	}

	glob.WithWsLog().Debugf(ctx, "Pusher connection request, protocol=%s, currentTime:%d", protocol, currentTime)

	conn, err := GetConnection(r)
	if err != nil {
		glob.WithWsLog().Errorf(ctx, "ws Upgrade error:%v", err)
		return
	}

	serverName := clientManager.ServerName
	sessionId := r.GetCtxVar(SESSION_ID_KEY)

	// 生成socket_id（v8.3.0格式：{serverName}.{timestamp}{random}）
	socketID := fmt.Sprintf("%s.%d%05d",
		serverName,
		time.Now().Unix(),
		safeRandomInt(100000))

	client := NewClient(conn.RemoteAddr().String(), socketID, conn, currentTime)
	client.SessionID = gconv.String(sessionId)
	client.ServerName = serverName

	// 保存客户端到Redis
	AddServerNameSocketId4Redis(ctx, client.SocketID, serverName)
	UpdateSocketIdHeartbeatTime4Redis(ctx, client.SocketID, currentTime)

	// 发送connection_established事件（⚠️ activity_timeout改为120秒）
	establishedData := ConnectionEstablishedData{
		SocketID:        socketID,
		ActivityTimeout: 120, // v8.3.0推荐值
	}

	err = client.SendPusherEvent(EventConnectionEstablished, "", establishedData)
	if err != nil {
		glob.WithWsLog().Errorf(ctx, "SendPusherEvent error:%v", err)
	}

	// 启动读写协程
	utils.SafeGo(ctx, func(ctx context.Context) {
		client.read(ctx)
	})
	utils.SafeGo(ctx, func(ctx context.Context) {
		client.write(ctx)
	})

	// 用户连接事件
	clientManager.Connect <- client

	glob.WithWsLog().Infof(ctx, "Pusher client connected: socket_id=%s", socketID)
}

// safeRandomInt returns a cryptographically secure random integer in [0, max).
func safeRandomInt(max int64) int64 {
	n, err := rand.Int(rand.Reader, big.NewInt(max))
	if err != nil {
		return 0
	}
	return n.Int64()
}
