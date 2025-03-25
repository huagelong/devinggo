// Package websocket
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package websocket

import (
	"context"
	"devinggo/modules/system/pkg/websocket/glob"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcron"
	"github.com/gogf/gf/v2/os/gtime"
	"sync"
)

// ClientManager 客户端管理
type ClientManager struct {
	Clients           map[string]*Client      // 全部的连接
	ClientsLock       sync.RWMutex            // 读写锁
	Connect           chan *Client            // 连接连接处理
	Disconnect        chan *Client            // 断开连接处理程序
	Broadcast         chan *WResponse         // 广播 向全部成员发送数据
	ClientIdBroadcast chan *ClientIdWResponse // 广播 向某个客户端发送数据
	TopicBroadcast    chan *TopicWResponse    // 广播 向某个标签成员发送数据
	ServerName        string
}

func NewClientManager() (clientManager *ClientManager) {
	clientManager = &ClientManager{
		Clients:           make(map[string]*Client),
		Connect:           make(chan *Client, 1000),
		Disconnect:        make(chan *Client, 1000),
		Broadcast:         make(chan *WResponse, 1000),
		ClientIdBroadcast: make(chan *ClientIdWResponse, 1000),
		TopicBroadcast:    make(chan *TopicWResponse, 1000),
	}
	return
}
func (manager *ClientManager) SetServerName(serverName string) {
	manager.ServerName = serverName
}

func (manager *ClientManager) GetClient(clientId string) (client *Client) {
	manager.ClientsLock.RLock()
	defer manager.ClientsLock.RUnlock()
	if v, ok := manager.Clients[clientId]; ok {
		client = v
	}
	return
}

// InClient 客户端是否存在
func (manager *ClientManager) InClient(client *Client) (ok bool) {
	manager.ClientsLock.RLock()
	defer manager.ClientsLock.RUnlock()
	_, ok = manager.Clients[client.ID]
	return
}

// GetClients 获取所有客户端
func (manager *ClientManager) GetClients() (clients map[string]*Client) {
	clients = make(map[string]*Client)
	manager.ClientsRange(func(clientId string, client *Client) (result bool) {
		clients[clientId] = client
		return true
	})
	return
}

// ClientsRange 遍历
func (manager *ClientManager) ClientsRange(f func(clientId string, client *Client) (result bool)) {
	manager.ClientsLock.RLock()
	defer manager.ClientsLock.RUnlock()
	for key, value := range manager.Clients {
		result := f(key, value)
		if result == false {
			return
		}
	}
	return
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
	manager.Clients[client.ID] = client
}

// DelClients 删除客户端
func (manager *ClientManager) DelClients(ctx context.Context, client *Client) {
	manager.ClientsLock.Lock()
	defer manager.ClientsLock.Unlock()
	if _, ok := manager.Clients[client.ID]; ok {
		delete(manager.Clients, client.ID)
		ClearClientId4Redis(ctx, client.ID)
	}
}

// EventConnect 用户建立连接事件
func (manager *ClientManager) EventConnect(ctx context.Context, client *Client) {
	manager.AddClients(client)

	client.ResponseSuccess(ctx, Connected, "0", g.Map{
		"sessionId": client.ID,
	})
}

// EventDisconnect 用户断开连接事件
func (manager *ClientManager) EventDisconnect(ctx context.Context, client *Client) {
	//更新离开时间 todo
	manager.DelClients(ctx, client)
	client.ResponseSuccess(ctx, Close, "0", g.Map{
		"sessionId": client.ID,
	})
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
	//定时任务，发送心跳包
	_, _ = gcron.Add(ctx, "0 0 */1 * * *", func(ctx context.Context) {
		res := &WResponse{
			Event:     PingAll,
			Code:      200,
			RequestId: "0",
		}
		SendToAll(res)
	})

	//定时清理
	_, _ = gcron.Add(ctx, "0 30 */1 * * *", func(ctx context.Context) {
		ClearExpire4Redis(ctx)
	})
	// 定时任务，清理超时连接
	_, _ = gcron.Add(ctx, "* */1 * * * *", func(ctx context.Context) {
		manager.clearTimeoutConnections(ctx)
	})

}

func (manager *ClientManager) EventBroadcast(ctx context.Context, response *WResponse) {
	clients := manager.GetClients()
	for _, conn := range clients {
		conn.SendMsg(ctx, response)
	}
}

func (manager *ClientManager) EventTopicBroadcast(ctx context.Context, response *TopicWResponse) {
	clients := manager.GetClients()
	for _, conn := range clients {
		if conn.topics.Contains(response.Topic) {
			conn.SendMsg(ctx, response.WResponse)
		}
	}
}

func (manager *ClientManager) EventClientIdBroadcast(ctx context.Context, response *ClientIdWResponse) {
	clients := manager.GetClients()
	for _, conn := range clients {
		if conn.ID == response.ID {
			conn.SendMsg(ctx, response.WResponse)
		}
	}
}

// 管道处理程序
func (manager *ClientManager) start(ctx context.Context) {
	for {
		select {
		case conn := <-manager.Connect:
			// 建立连接事件
			glob.WithWsLog().Debug(ctx, "EventConnect:", "conn.id:", conn.ID)
			manager.EventConnect(ctx, conn)
		case conn := <-manager.Disconnect:
			// 断开连接事件
			glob.WithWsLog().Debug(ctx, "EventDisconnect:", "conn.id:", conn.ID)
			manager.EventDisconnect(ctx, conn)
		case response := <-manager.Broadcast:
			// 全部客户端广播事件
			glob.WithWsLog().Debug(ctx, "EventBroadcast:", response)
			manager.EventBroadcast(ctx, response)
		case response := <-manager.TopicBroadcast:
			// 标签广播事件
			glob.WithWsLog().Debug(ctx, "EventTopicBroadcast:", response)
			manager.EventTopicBroadcast(ctx, response)
		case response := <-manager.ClientIdBroadcast:
			// 单个客户端广播事件
			glob.WithWsLog().Debug(ctx, "EventClientIdBroadcast:", response)
			manager.EventClientIdBroadcast(ctx, response)
		}

	}
}

// SendToAll 发送全部客户端
func SendToAll(response *WResponse) {
	clientManager.Broadcast <- response
}

// SendToClientID  发送单个客户端
func SendToClientID(id string, response *WResponse) {
	clientRes := &ClientIdWResponse{
		ID:        id,
		WResponse: response,
	}
	clientManager.ClientIdBroadcast <- clientRes
}

// SendToTopic 发送某个标签
func SendToTopic(topic string, response *WResponse) {
	topicRes := &TopicWResponse{
		Topic:     topic,
		WResponse: response,
	}
	clientManager.TopicBroadcast <- topicRes
}
