// Package websocket
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package websocket

import (
	"context"
	"devinggo/modules/system/model"
	"devinggo/modules/system/pkg/utils"
	"devinggo/modules/system/pkg/websocket/glob"
	"devinggo/modules/system/service"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gorilla/websocket"
	"net/http"
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
	glob.WithWsLog().Debugf(ctx, "Connected!currentTime:%d", currentTime)
	conn, err := GetConnection(r)
	if err != nil {
		glob.WithWsLog().Errorf(ctx, "ws Upgrade error:%v", err)
	}
	serverName := clientManager.ServerName
	sessionId := r.GetCtxVar(SESSION_ID_KEY)
	if err != nil && g.IsEmpty(sessionId) {
		conn.WriteJSON(&WResponse{
			Event:     Connected,
			Message:   "sessionId miss",
			Code:      500,
			RequestId: "0",
		})
		conn.Close()
		return
	}
	client := NewClient(conn.RemoteAddr().String(), gconv.String(sessionId), conn, currentTime)
	AddServerNameClientId4Redis(ctx, client.ID, serverName)
	UpdateClientIdHeartbeatTime4Redis(ctx, client.ID, currentTime)
	client.ServerName = serverName
	utils.SafeGo(ctx, func(ctx context.Context) {
		client.read(ctx)
	})
	utils.SafeGo(ctx, func(ctx context.Context) {
		client.write(ctx)
	})
	// 用户连接事件
	clientManager.Connect <- client
}
