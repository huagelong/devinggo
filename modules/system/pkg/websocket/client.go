// Package websocket
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package websocket

import (
	"context"
	"encoding/json"
	"runtime/debug"
	"sync"
	"time"

	"devinggo/modules/system/pkg/utils"
	"devinggo/modules/system/pkg/websocket/glob"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gorilla/websocket"
)

const (
	// 用户连接超时时间（150秒 = 120 + 30，v8.3.0推荐）
	heartbeatExpirationTime = 150
)

// Client 客户端连接（Pusher协议）
type Client struct {
	Addr           string                 // 客户端地址
	SocketID       string                 // 连接唯一标识（Pusher格式：{serverName}.{uniqueId}）
	Socket         *websocket.Conn        // 用户连接
	Send           chan *PusherResponse   // 待发送的数据（Pusher格式）
	SendClose      bool                   // 发送是否关闭
	FirstTime      int64                  // 首次连接事件
	HeartbeatTime  int64                  // 用户上次心跳时间
	LoginTime      int64                  // 登录时间 登录以后才有
	ServerName     string                 // 服务器名称
	Channels       []string               // 订阅的频道列表（改名自topics）
	UserID         string                 // Presence Channel用户ID
	UserInfo       map[string]interface{} // Presence Channel用户信息
	LastEventTime  time.Time              // 最后事件时间（用于速率限制）
	EventCount     int                    // 1秒内事件计数
	IsDisconnected bool                   // 断线标记（Grace Period用）
	SessionID      string                 // 会话ID（用于认证）
	mu             sync.RWMutex           // 保护Channels等字段的读写锁
}

// NewClient 初始化（Pusher协议）
func NewClient(addr string, socketID string, socket *websocket.Conn, firstTime int64) (client *Client) {
	client = &Client{
		Addr:           addr,
		SocketID:       socketID,
		Socket:         socket,
		Send:           make(chan *PusherResponse, 100),
		SendClose:      false,
		FirstTime:      firstTime,
		HeartbeatTime:  firstTime,
		Channels:       make([]string, 0),
		UserInfo:       make(map[string]interface{}),
		LastEventTime:  time.Now(),
		EventCount:     0,
		IsDisconnected: false,
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
		glob.WithWsLog().Debug(ctx, "read conn close SocketID:", c.SocketID)
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
		glob.WithWsLog().Debug(ctx, "write conn close SocketID:", c.SocketID)
		c.close(ctx)
	}()
	for message := range c.Send {
		// glob.WithWsLog().Debug(ctx, "response:", message)

		err := c.Socket.WriteJSON(message)
		if err != nil {
			glob.WithWsLog().Warning(ctx, "WriteJSON error:", err)
		}

		// 发送完成后释放对象回池（性能优化）
		ReleasePusherResponse(message)
	}
}

// SendMsg 发送Pusher响应消息
func (c *Client) SendMsg(msg *PusherResponse) error {
	if c == nil || c.SendClose {
		return nil
	}
	defer func() {
		if r := recover(); r != nil {
			glob.WithWsLog().Warning(gctx.GetInitCtx(), "SendMsg error:", string(debug.Stack()), r)
		}
	}()

	c.Send <- msg
	return nil
}

// SendPusherEvent 发送Pusher事件（⚠️ 自动序列化data为JSON字符串）
func (c *Client) SendPusherEvent(event, channel string, data interface{}) error {
	var dataStr string
	if str, ok := data.(string); ok {
		dataStr = str
	} else {
		jsonBytes, err := json.Marshal(data)
		if err != nil {
			// 🔍 如果序列化失败，记录错误
			glob.WithWsLog().Errorf(gctx.GetInitCtx(), "🔍 [SendPusherEvent] JSON序列化失败: %v, data类型: %T", err, data)
			return err
		}
		dataStr = string(jsonBytes)
	}

	// 🔍 调试：打印发送的数据
	if event == EventSubscriptionSucceeded && channel != "" {
		glob.WithWsLog().Debugf(gctx.GetInitCtx(), "🔍 [SendPusherEvent] event=%s, channel=%s", event, channel)
		glob.WithWsLog().Debugf(gctx.GetInitCtx(), "🔍 [SendPusherEvent] dataStr长度=%d, 前200字符=%s", len(dataStr), dataStr[:min(200, len(dataStr))])
	}

	// 使用sync.Pool优化内存分配
	response := AcquirePusherResponse()
	response.Event = event
	response.Channel = channel
	response.Data = dataStr

	return c.SendMsg(response)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// SendError 发送错误消息（pusher:error）
func (c *Client) SendError(message string, code int) error {
	errData := ErrorData{
		Message: message,
		Code:    code,
	}
	return c.SendPusherEvent(EventError, "", errData)
}

// SendSubscriptionError 发送订阅错误（pusher:subscription_error，v8.3.0格式）
func (c *Client) SendSubscriptionError(channel, errorType, message string, status int) error {
	errData := SubscriptionErrorData{
		Type:   errorType,
		Error:  message,
		Status: status,
	}
	return c.SendPusherEvent(EventSubscriptionError, channel, errData)
}

// Heartbeat 心跳更新
func (c *Client) Heartbeat(currentTime int64) {
	c.HeartbeatTime = currentTime
}

// IsHeartbeatTimeout 心跳是否超时
func (c *Client) IsHeartbeatTimeout(currentTime int64) (timeout bool) {
	if c.HeartbeatTime+heartbeatExpirationTime <= currentTime {
		timeout = true
	}
	return
}

// AddChannel 添加频道订阅
func (c *Client) AddChannel(channel string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	// 检查是否已订阅
	for _, ch := range c.Channels {
		if ch == channel {
			return
		}
	}
	c.Channels = append(c.Channels, channel)
}

// RemoveChannel 移除频道订阅
func (c *Client) RemoveChannel(channel string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	for i, ch := range c.Channels {
		if ch == channel {
			c.Channels = append(c.Channels[:i], c.Channels[i+1:]...)
			return
		}
	}
}

// HasChannel 检查是否订阅了指定频道
func (c *Client) HasChannel(channel string) bool {
	c.mu.RLock()
	defer c.mu.RUnlock()

	for _, ch := range c.Channels {
		if ch == channel {
			return true
		}
	}
	return false
}

// GetChannels 获取所有订阅的频道（线程安全）
func (c *Client) GetChannels() []string {
	c.mu.RLock()
	defer c.mu.RUnlock()

	channels := make([]string, len(c.Channels))
	copy(channels, c.Channels)
	return channels
}

// 关闭客户端
func (c *Client) close(ctx context.Context) {
	if c.SendClose {
		return
	}

	c.IsDisconnected = true

	//删除客户端数据
	utils.SafeGo(gctx.GetInitCtx(), func(ctx context.Context) {
		// 立即清理非Presence订阅
		for _, channel := range c.GetChannels() {
			if !IsPresenceChannel(channel) {
				LeaveChannel4Redis(ctx, channel, c.SocketID)
			}
		}

		// ⚠️ Presence Channel延迟清理（30秒 Grace Period）
		presenceChannels := make([]string, 0)
		for _, channel := range c.GetChannels() {
			if IsPresenceChannel(channel) {
				presenceChannels = append(presenceChannels, channel)
				// 标记断线
				MarkPresenceDisconnect4Redis(ctx, c.SocketID)
			}
		}

		// 30秒后检查是否重连
		time.AfterFunc(30*time.Second, func() {
			newCtx := gctx.GetInitCtx()
			if c.IsDisconnected && !IsPresenceDisconnected4Redis(newCtx, c.SocketID) {
				// 已重连，不清理
				return
			}

			if c.IsDisconnected { // 30秒内未重连
				for _, channel := range presenceChannels {
					if c.UserID != "" {
						// 移除Presence成员
						RemovePresenceMember4Redis(newCtx, channel, c.UserID)
						// 广播member_removed事件
						BroadcastMemberRemoved(newCtx, channel, c.UserID)
						glob.WithWsLog().Debugf(newCtx, "Grace period expired: removed member %s from %s", c.UserID, channel)
					}
					LeaveChannel4Redis(newCtx, channel, c.SocketID)
				}
				// 清除断线标记
				ClearPresenceDisconnect4Redis(newCtx, c.SocketID)
			}
		})

		// 清理速率限制器
		GetRateLimiter().RemoveBucket(c.SocketID)

		ClearSocketId4Redis(ctx, c.SocketID)
	})

	// 关闭发送通道和WebSocket连接
	close(c.Send)
	c.SendClose = true
	c.Socket.Close()

	// 从客户端管理器中移除
	clientManager.DelClients(ctx, c)
}
