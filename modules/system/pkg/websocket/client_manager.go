// Package websocket
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package websocket

import (
	"context"
	"sync"

	"devinggo/modules/system/pkg/websocket/glob"

	"github.com/gogf/gf/v2/os/gcron"
	"github.com/gogf/gf/v2/os/gtime"
)

// ClientManager 客户端管理（Pusher协议）
type ClientManager struct {
	Clients           map[string]*Client      // 全部的连接（key为SocketID）
	ClientsLock       sync.RWMutex            // 读写锁
	Connect           chan *Client            // 连接连接处理
	Disconnect        chan *Client            // 断开连接处理程序
	Broadcast         chan *PusherResponse    // 广播 向全部成员发送数据
	SocketIdBroadcast chan *ClientIdWResponse // 广播 向某个客户端发送数据
	ChannelBroadcast  chan *TopicWResponse    // 广播 向某个频道成员发送数据
	ServerName        string
}

func NewClientManager() (clientManager *ClientManager) {
	clientManager = &ClientManager{
		Clients:           make(map[string]*Client),
		Connect:           make(chan *Client, 1000),
		Disconnect:        make(chan *Client, 1000),
		Broadcast:         make(chan *PusherResponse, 1000),
		SocketIdBroadcast: make(chan *ClientIdWResponse, 1000),
		ChannelBroadcast:  make(chan *TopicWResponse, 1000),
	}
	return
}
func (manager *ClientManager) SetServerName(serverName string) {
	manager.ServerName = serverName
}

func (manager *ClientManager) GetClient(socketId string) (client *Client) {
	manager.ClientsLock.RLock()
	defer manager.ClientsLock.RUnlock()
	if v, ok := manager.Clients[socketId]; ok {
		client = v
	}
	return
}

// GetClientBySocketID 根据SocketID获取客户端（别名）
func (manager *ClientManager) GetClientBySocketID(socketId string) (client *Client) {
	return manager.GetClient(socketId)
}

// InClient 客户端是否存在
func (manager *ClientManager) InClient(client *Client) (ok bool) {
	manager.ClientsLock.RLock()
	defer manager.ClientsLock.RUnlock()
	_, ok = manager.Clients[client.SocketID]
	return
}

// GetClients 获取所有客户端
func (manager *ClientManager) GetClients() (clients map[string]*Client) {
	clients = make(map[string]*Client)
	manager.ClientsRange(func(socketId string, client *Client) (result bool) {
		clients[socketId] = client
		return true
	})
	return
}

// ClientsRange 遍历
func (manager *ClientManager) ClientsRange(f func(socketId string, client *Client) (result bool)) {
	manager.ClientsLock.RLock()
	defer manager.ClientsLock.RUnlock()
	for key, value := range manager.Clients {
		result := f(key, value)
		if result == false {
			return
		}
	}
}

// GetClientsLen 获取客户端总数
func (manager *ClientManager) GetClientsLen() (clientsLen int) {
	clientsLen = len(manager.Clients)
	return
}

// AddClients 添加客户端
func (manager *ClientManager) AddClients(client *Client) {
	manager.ClientsLock.Lock()
	defer manager.ClientsLock.Unlock()
	manager.Clients[client.SocketID] = client
}

// DelClients 删除客户端
func (manager *ClientManager) DelClients(ctx context.Context, client *Client) {
	manager.ClientsLock.Lock()
	defer manager.ClientsLock.Unlock()
	if _, ok := manager.Clients[client.SocketID]; ok {
		delete(manager.Clients, client.SocketID)
		ClearSocketId4Redis(ctx, client.SocketID)
	}
}

// EventConnect 用户建立连接事件
func (manager *ClientManager) EventConnect(ctx context.Context, client *Client) {
	manager.AddClients(client)
	glob.WithWsLog().Infof(ctx, "Client connected: socket_id=%s, addr=%s", client.SocketID, client.Addr)
}

// EventDisconnect 用户断开连接事件
func (manager *ClientManager) EventDisconnect(ctx context.Context, client *Client) {
	manager.DelClients(ctx, client)
	glob.WithWsLog().Infof(ctx, "Client disconnected: socket_id=%s", client.SocketID)
}

// ClearTimeoutConnections 定时清理超时连接
func (manager *ClientManager) clearTimeoutConnections(ctx context.Context) {
	currentTime := int64(gtime.Now().Unix())
	clients := clientManager.GetClients()
	for _, client := range clients {
		if client.IsHeartbeatTimeout(currentTime) {
			glob.WithWsLog().Debug(ctx, "Heart beat timeout , close connect ", client.Addr, client.LoginTime, client.HeartbeatTime)
			_ = client.Socket.Close()
			manager.DelClients(ctx, client)
		}
	}
}

// WebsocketPing 定时任务
func (manager *ClientManager) cronJob(ctx context.Context) {
	//定时清理
	_, _ = gcron.Add(ctx, "0 30 */1 * * *", func(ctx context.Context) {
		ClearExpire4Redis(ctx)
	})
	// 定时任务，清理超时连接
	_, _ = gcron.Add(ctx, "* */1 * * * *", func(ctx context.Context) {
		manager.clearTimeoutConnections(ctx)
	})
}

func (manager *ClientManager) EventBroadcast(ctx context.Context, response *PusherResponse) {
	clients := manager.GetClients()
	for _, conn := range clients {
		conn.SendMsg(response)
	}
}

func (manager *ClientManager) EventChannelBroadcast(ctx context.Context, response *TopicWResponse) {
	clients := manager.GetClients()
	glob.WithWsLog().Debugf(ctx, "EventChannelBroadcast: channel=%s, event=%s, total_clients=%d",
		response.Topic, response.PusherResponse.Event, len(clients))

	sentCount := 0
	for _, conn := range clients {
		if response.ExcludeSocketID != "" && conn.SocketID == response.ExcludeSocketID {
			continue
		}
		if conn.HasChannel(response.Topic) {
			glob.WithWsLog().Debugf(ctx, "Sending to client: socket_id=%s, channel=%s", conn.SocketID, response.Topic)
			conn.SendMsg(response.PusherResponse)
			sentCount++
		}
	}

	glob.WithWsLog().Debugf(ctx, "EventChannelBroadcast completed: sent to %d clients", sentCount)
}

func (manager *ClientManager) EventSocketIdBroadcast(ctx context.Context, response *ClientIdWResponse) {
	clients := manager.GetClients()
	for _, conn := range clients {
		if conn.SocketID == response.SocketID {
			conn.SendMsg(response.PusherResponse)
		}
	}
}

// 管道处理程序
func (manager *ClientManager) start(ctx context.Context) {
	for {
		select {
		case conn := <-manager.Connect:
			// 建立连接事件
			glob.WithWsLog().Debug(ctx, "EventConnect:", "conn.socketId:", conn.SocketID)
			manager.EventConnect(ctx, conn)
		case conn := <-manager.Disconnect:
			// 断开连接事件
			glob.WithWsLog().Debug(ctx, "EventDisconnect:", "conn.socketId:", conn.SocketID)
			manager.EventDisconnect(ctx, conn)
		case response := <-manager.Broadcast:
			// 全部客户端广播事件
			glob.WithWsLog().Debug(ctx, "EventBroadcast:", response)
			manager.EventBroadcast(ctx, response)
		case response := <-manager.ChannelBroadcast:
			// 频道广播事件
			glob.WithWsLog().Debug(ctx, "EventChannelBroadcast:", response)
			manager.EventChannelBroadcast(ctx, response)
		case response := <-manager.SocketIdBroadcast:
			// 单个客户端广播事件
			glob.WithWsLog().Debug(ctx, "EventSocketIdBroadcast:", response)
			manager.EventSocketIdBroadcast(ctx, response)
		}

	}
}

// SendToAll 发送全部客户端
func SendToAll(response *PusherResponse) {
	clientManager.Broadcast <- response
}

// SendToSocketID 发送单个客户端
func SendToSocketID(socketId string, response *PusherResponse) {
	clientRes := &ClientIdWResponse{
		SocketID:       socketId,
		PusherResponse: response,
	}
	clientManager.SocketIdBroadcast <- clientRes
}

// SendToChannel 发送某个频道
func SendToChannel(channel string, response *PusherResponse) {
	SendToChannelWithExclude(channel, response, "")
}

// SendToChannelWithExclude 发送某个频道（可排除指定socket_id）
func SendToChannelWithExclude(channel string, response *PusherResponse, excludeSocketID string) {
	channelRes := &TopicWResponse{
		Topic:           channel,
		ExcludeSocketID: excludeSocketID,
		PusherResponse:  response,
	}
	clientManager.ChannelBroadcast <- channelRes
}

// GetClientManager 获取客户端管理器（用于外部访问）
func GetClientManager() *ClientManager {
	return clientManager
}

// GetServerName 获取当前服务器名称
func GetServerName() string {
	return clientManager.ServerName
}

// GetLocalClient 根据socket_id获取本地客户端
func GetLocalClient(socketId string) *Client {
	return clientManager.GetClient(socketId)
}

// TerminateLocalClient 终止本地客户端连接
func TerminateLocalClient(socketId string) bool {
	if client := clientManager.GetClient(socketId); client != nil {
		// 发送关闭帧并关闭连接
		client.SendClose = true
		_ = client.Socket.Close()
		return true
	}
	return false
}
