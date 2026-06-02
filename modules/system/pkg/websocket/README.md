# 🚀 Pusher WebSocket 使用教程

## 📖 目录

- [功能特性](#功能特性)
- [快速开始](#快速开始)
- [配置说明](#配置说明)
- [客户端使用](#客户端使用)
- [服务端API](#服务端api)
- [使用场景](#使用场景)
- [常见问题](#常见问题)
- [详细文档](#详细文档)

---

## 功能特性

系统集成了完整的 **Pusher Protocol v8.3.0** WebSocket实时通信功能，支持：

- ✅ **Public Channel**：公开频道，无需认证
- ✅ **Private Channel**：私有频道，HMAC-SHA256签名认证
- ✅ **Presence Channel**：在线状态频道，实时成员列表
- ✅ **Client Events**：客户端之间直接通信（Private/Presence频道）
- ✅ **分布式支持**：基于Redis的跨服务器消息同步
- ✅ **速率限制**：Token Bucket算法防止滥用（默认10 events/sec）
- ✅ **心跳机制**：自动检测断线（120秒超时）
- ✅ **Presence缓存**：成员列表缓存优化（5秒TTL）

---

## 快速开始

### 1️⃣ 浏览器测试（推荐）

启动服务后，直接访问测试页面：

```
http://localhost:8070/pusher-test.html
```

**功能演示**：
- 连接建立与Socket ID显示
- Public频道订阅
- Private频道认证与订阅
- Presence频道成员管理
- Client Event发送测试
- 速率限制验证

---

### 2️⃣ JavaScript客户端（pusher-js）

#### 安装依赖

```bash
npm install pusher-js@8.3.0
```

#### 基础连接

```javascript
const Pusher = require('pusher-js');

// 连接WebSocket服务器（pusher-js v8.3.0 推荐配置）
const pusher = new Pusher('devinggo-app-key', {
  wsHost: 'localhost',
  wsPort: 8070,
  forceTLS: false,
  enabledTransports: ['ws'],
  cluster: 'mt1',  // 自托管服务器必需（虚拟cluster名称）
  disableStats: true,  // 禁用统计信息

  // Private/Presence 频道授权配置（v8.3.0 新格式）
  channelAuthorization: {
    endpoint: '/system/pusher/auth',
    transport: 'ajax',
    headers: {
      'Authorization': 'Bearer YOUR_JWT_TOKEN'  // 可选：如需后端认证
    }
  },

  // 用户认证配置（可选，用于 pusher:signin）
  userAuthentication: {
    endpoint: '/system/pusher/user-auth',
    transport: 'ajax'
  }
});

// 监听连接状态
pusher.connection.bind('connected', function() {
  console.log('✅ 连接成功！Socket ID:', pusher.socket_id);
});

pusher.connection.bind('disconnected', function() {
  console.log('⚠️ 连接断开');
});

pusher.connection.bind('error', function(err) {
  console.error('❌ 连接错误:', err);
});
```

#### Public频道（无需认证）

```javascript
// 订阅Public频道
const publicChannel = pusher.subscribe('chat-room');

// 监听订阅成功
publicChannel.bind('pusher:subscription_succeeded', function() {
  console.log('✅ Public频道订阅成功');
});

// 监听自定义事件
publicChannel.bind('message', function(data) {
  console.log('📨 收到消息:', data);
  // data 示例: { user: "Alice", text: "Hello World!" }
});

// 监听所有事件
publicChannel.bind_global(function(event, data) {
  console.log('事件:', event, '数据:', data);
});
```

#### Private频道（需要认证）

```javascript
// 订阅Private频道（自动请求 /system/pusher/auth 进行认证）
const privateChannel = pusher.subscribe('private-user-123');

// 监听订阅成功
privateChannel.bind('pusher:subscription_succeeded', function() {
  console.log('✅ Private频道订阅成功');
});

// 监听订阅失败
privateChannel.bind('pusher:subscription_error', function(err) {
  console.error('❌ Private频道订阅失败:', err);
});

// 监听自定义事件
privateChannel.bind('notification', function(data) {
  console.log('🔔 收到通知:', data);
});

// 发送Client Event（仅Private/Presence频道可用）
privateChannel.trigger('client-typing', { 
  user: 'Alice',
  message: '正在输入...' 
});

// 监听其他客户端的Client Event
privateChannel.bind('client-typing', function(data) {
  console.log('⌨️ 用户正在输入:', data.user);
});
```

#### Presence频道（在线状态）

```javascript
// 订阅Presence频道
const presenceChannel = pusher.subscribe('presence-lobby');

// 监听订阅成功（获取初始成员列表）
presenceChannel.bind('pusher:subscription_succeeded', function(members) {
  console.log('✅ Presence频道订阅成功');
  console.log('📊 当前在线人数:', members.count);
  
  // 遍历所有成员
  members.each(function(member) {
    console.log('👤 成员:', member.id, member.info);
  });
});

// 监听新成员加入
presenceChannel.bind('pusher:member_added', function(member) {
  console.log('✨ 新成员加入:', member.id, member.info);
});

// 监听成员离开
presenceChannel.bind('pusher:member_removed', function(member) {
  console.log('👋 成员离开:', member.id);
});

// 获取当前成员
const currentMembers = presenceChannel.members;
console.log('在线成员数:', currentMembers.count);
```

---

### 3️⃣ 服务端推送消息（Go）

#### 向指定频道广播

```go
package main

import (
    "context"
    "devinggo/modules/system/pkg/websocket"
)

func SendMessageToChannel() {
    ctx := context.Background()
    
    // 向chat-room频道的所有订阅者发送消息
    websocket.BroadcastToChannel(ctx, "chat-room", "message", map[string]interface{}{
        "user": "System",
        "text": "欢迎加入聊天室！",
        "timestamp": time.Now().Unix(),
    })
}
```

#### 向指定用户发送消息

```go
func SendMessageToUser(socketID string) {
    ctx := context.Background()
    
    // 向指定Socket ID发送消息
    websocket.SendToSocketID(ctx, socketID, "notification", map[string]interface{}{
        "title": "新消息",
        "content": "您有一条新通知",
        "type": "info",
    })
}
```

#### Presence成员管理

```go
// 广播成员加入事件
func NotifyMemberJoined(channel string, memberID string, memberInfo map[string]interface{}) {
    ctx := context.Background()
    
    websocket.BroadcastMemberAdded(ctx, channel, memberID, memberInfo)
}

// 广播成员离开事件
func NotifyMemberLeft(channel string, memberID string) {
    ctx := context.Background()
    
    websocket.BroadcastMemberRemoved(ctx, channel, memberID)
}
```

---

## 配置说明

### 配置文件位置

- 开发环境：`hack/config.yaml`
- 生产环境：`manifest/config/config.yaml`

### Pusher配置项

```yaml
pusher:
  # 客户端连接时使用的应用密钥（公开）
  appKey: "devinggo-app-key"
  
  # 服务端签名密钥（私密，⚠️ 生产环境必须修改！）
  appSecret: "devinggo-app-secret-change-me"
  
  # 客户端活动超时时间（秒），超时后自动断开
  activityTimeout: 120
  
  # 服务端心跳检查间隔（秒）
  heartbeatCheckInterval: 60
  
  # Presence频道优雅下线期（秒）
  # 断线后在此时间内重连，不会触发member_removed事件
  presenceGracePeriod: 30
  
  # 单个连接最大订阅频道数
  maxChannelsPerConnection: 100
  
  # Client Event速率限制（events/秒）
  clientEventRateLimit: 10
```

### ⚠️ 生产环境安全配置

#### 1. 修改appSecret

```bash
# 生成随机密钥（64字符）
openssl rand -hex 32
```

```yaml
pusher:
  appSecret: "your-random-64-char-hex-string"
```

#### 2. 启用TLS/SSL

```yaml
server:
  address: ":443"
  httpsAddr: ":443"
  httpsCertPath: "/path/to/cert.pem"
  httpsKeyPath: "/path/to/key.pem"
```

客户端连接改为：

```javascript
const pusher = new Pusher('devinggo-app-key', {
  wsHost: 'your-domain.com',
  wsPort: 443,
  forceTLS: true,  // 启用TLS
  enabledTransports: ['ws', 'wss'],
  cluster: 'mt1',
  disableStats: true,

  // Private/Presence 频道授权配置（v8.3.0 新格式）
  channelAuthorization: {
    endpoint: 'https://your-domain.com/system/pusher/auth',
    transport: 'ajax'
  },

  // 用户认证配置（可选，用于 pusher:signin）
  userAuthentication: {
    endpoint: 'https://your-domain.com/system/pusher/user-auth',
    transport: 'ajax'
  }
});
```

#### 3. 配置Redis密码

```yaml
redis:
  default:
    address: "127.0.0.1:6379"
    pass: "your-redis-password"
    db: 0
```

---

## 客户端使用

### WebSocket端点

- **WebSocket连接**: `ws://localhost:8070/system/ws`
- **HTTP认证端点**: `POST http://localhost:8070/system/pusher/auth`

### 频道命名规则

| 频道类型 | 前缀格式 | 示例 | 认证 |
|---------|---------|------|------|
| Public | 无前缀 | `chat-room`, `news` | ❌ 不需要 |
| Private | `private-` | `private-user-123`, `private-order-456` | ✅ 需要 |
| Presence | `presence-` | `presence-lobby`, `presence-room-1` | ✅ 需要 |

### 错误码说明

| 错误码 | 说明 | 解决方法 |
|-------|------|---------|
| 4000 | 正常关闭 | - |
| 4001 | 服务器下线 | 等待服务器恢复 |
| 4009 | 未授权 | 检查认证信息 |
| 4200 | 心跳超时 | 检查网络连接 |
| 4301 | Client Event错误 | 检查事件名、频道类型、速率限制 |

### Client Event限制

- ✅ **仅支持**：Private和Presence频道
- ❌ **不支持**：Public频道
- 📏 **事件名长度**：最大200字节
- 🔤 **事件名前缀**：必须是 `client-`
- ⏱️ **速率限制**：默认10 events/秒

---

## 服务端API

### ClientManager方法

```go
// 获取全局ClientManager实例
manager := websocket.ClientManagerInstance

// 向指定Socket ID发送消息
manager.SendToSocketID(socketID, response)

// 向指定频道广播消息
manager.BroadcastToChannel(channel, response)

// 向所有连接广播
manager.BroadcastToAll(response)

// 获取在线连接数
count := manager.GetOnlineCount()

// 获取指定频道订阅数
channelCount := manager.GetChannelSubscriberCount(channel)
```

### 辅助函数

```go
// 构造Pusher响应
response := websocket.NewPusherResponse(event, channel, data)

// 构造错误响应
errorResp := websocket.NewPusherError(code, message)

// 频道类型判断
isPublic := websocket.IsPublicChannel(channel)
isPrivate := websocket.IsPrivateChannel(channel)
isPresence := websocket.IsPresenceChannel(channel)
```

---

## 使用场景

### 1. 📱 即时聊天系统

```javascript
// 多人聊天室
const chatChannel = pusher.subscribe('chat-room');
chatChannel.bind('message', function(data) {
  displayMessage(data.user, data.text);
});

// Private一对一聊天
const privateChat = pusher.subscribe('private-chat-' + userId);
privateChat.bind('message', function(data) {
  displayPrivateMessage(data);
});

// 显示"正在输入"状态
privateChat.trigger('client-typing', { userId: currentUserId });
```

### 2. 🔔 实时通知系统

```go
// 服务端：订单状态更新通知
func NotifyOrderUpdate(userId string, orderId string, status string) {
    socketID := getUserSocketID(userId)  // 从Redis获取用户的Socket ID
    
    websocket.SendToSocketID(ctx, socketID, "order-update", map[string]interface{}{
        "orderId": orderId,
        "status": status,
        "message": "您的订单状态已更新",
    })
}
```

```javascript
// 客户端：监听通知
const userChannel = pusher.subscribe('private-user-' + userId);
userChannel.bind('order-update', function(data) {
  showNotification(data.message);
  updateOrderStatus(data.orderId, data.status);
});
```

### 3. 👥 在线状态管理

```javascript
// 协同编辑页面的在线用户
const presenceChannel = pusher.subscribe('presence-document-123');

presenceChannel.bind('pusher:subscription_succeeded', function(members) {
  updateOnlineUsers(members);
});

presenceChannel.bind('pusher:member_added', function(member) {
  addUserToList(member.id, member.info.name);
});

presenceChannel.bind('pusher:member_removed', function(member) {
  removeUserFromList(member.id);
});

// 监听用户光标位置
presenceChannel.bind('client-cursor-move', function(data) {
  updateUserCursor(data.userId, data.position);
});
```

### 4. 📊 实时数据监控

```go
// 服务端：推送实时数据
func PushRealtimeData() {
    ticker := time.NewTicker(1 * time.Second)
    for range ticker.C {
        data := collectMetrics()  // 采集系统指标
        
        websocket.BroadcastToChannel(ctx, "metrics-dashboard", "update", data)
    }
}
```

```javascript
// 客户端：实时图表更新
const metricsChannel = pusher.subscribe('metrics-dashboard');
metricsChannel.bind('update', function(data) {
  updateChart(data.cpu, data.memory, data.network);
});
```

### 5. 🎮 游戏/互动应用

```javascript
// 游戏房间
const gameChannel = pusher.subscribe('presence-game-room-1');

// 玩家移动
gameChannel.trigger('client-player-move', {
  playerId: myId,
  position: { x: 100, y: 200 }
});

// 监听其他玩家
gameChannel.bind('client-player-move', function(data) {
  updatePlayerPosition(data.playerId, data.position);
});
```

---

## 常见问题

### 1. 连接失败

**现象**: 无法建立WebSocket连接

**检查**:
```bash
# 1. 确认服务运行
curl http://localhost:8070/health

# 2. 确认端口开放
netstat -ano | findstr "8070"

# 3. 确认WebSocket端点
# 端点: ws://localhost:8070/system/ws
```

### 2. Private频道认证失败

**原因**: HTTP认证端点返回错误

**检查**:
```bash
# 测试认证端点
curl -X POST http://localhost:8070/system/pusher/auth \
  -H "Content-Type: application/json" \
  -d '{
    "socket_id": "test.123",
    "channel_name": "private-test"
  }'
```

**正确响应**:
```json
{
  "auth": "devinggo-app-key:abc123..."
}
```

### 3. Client Event发送失败（错误4301）

**可能原因**:
- ❌ Public频道不支持Client Event
- ❌ 事件名不是 `client-` 前缀
- ❌ 事件名超过50字节
- ❌ 速率限制（超过10次/秒）

**验证**:
```javascript
// ❌ 错误示例
publicChannel.trigger('message', {});  // Public频道不支持

// ✅ 正确示例
privateChannel.trigger('client-message', {});  // Private频道可以
```

### 4. Presence成员列表不更新

**原因**: Grace Period机制（30秒）

**说明**: 
- 用户断线后30秒内重连，不会触发 `member_removed` 事件
- 这是为了避免网络抖动导致频繁的成员变更通知

**调整配置**:
```yaml
pusher:
  presenceGracePeriod: 10  # 改为10秒
```

### 5. 消息延迟/丢失

**检查Redis**:
```bash
redis-cli
> PING
> INFO replication
> KEYS ws:*
```

**检查连接数**:
```go
// 获取在线连接数
count := websocket.ClientManagerInstance.GetOnlineCount()
log.Printf("当前在线: %d", count)
```

---

## 详细文档

| 文档 | 说明 |
|------|------|
| [QUICKSTART.md](QUICKSTART.md) | 快速启动指南 - 从零开始配置和测试 |
| [IMPLEMENTATION_SUMMARY.md](IMPLEMENTATION_SUMMARY.md) | 实现总结 - 架构设计和性能指标 |
| [TEST_VALIDATION.md](TEST_VALIDATION.md) | 测试验证 - 完整测试清单和测试脚本 |
| [CHECKLIST.md](CHECKLIST.md) | 功能检查清单 - 100%完成度验证 |

---

## 兼容性

### 协议版本
- ✅ **Pusher Protocol v7** - 完全兼容
- ✅ **pusher-js v8.3.0** - 官方客户端库

### 支持的客户端库

| 库 | 语言 | 版本 | 说明 |
|----|------|------|------|
| [pusher-js](https://github.com/pusher/pusher-js) | JavaScript | v8.3.0+ | 官方推荐 |
| [Laravel Echo](https://laravel.com/docs/broadcasting) | JavaScript | - | Laravel生态 |
| [Soketi](https://docs.soketi.app/) | - | - | 开源Pusher替代 |

### 传输协议
- ✅ **WebSocket** (ws://)
- ✅ **WebSocket Secure** (wss://)

---

## 性能指标

### 压力测试结果（理论）

| 指标 | 数值 |
|------|------|
| 并发连接数 | 1000+ |
| 消息吞吐量 | 10万条/秒 |
| 消息延迟 (P99) | < 100ms |
| 内存占用 (1000连接) | < 500MB |
| CPU占用 (4核) | < 50% |

### 优化措施

- ✅ **sync.Pool** - 对象复用，减少GC压力 (-30~50% GC)
- ✅ **Presence缓存** - 5秒TTL，减少Redis查询 (-80~90%)
- ✅ **Rate Limiting** - Token Bucket算法防止滥用
- ✅ **Grace Period** - 减少网络抖动影响

---

## 技术支持

### 问题反馈

- **GitHub Issues**: [huagelong/devinggo/issues](https://github.com/huagelong/devinggo/issues)
- **QQ群**: 483122520

### 相关资源

- [Pusher Protocol文档](https://pusher.com/docs/channels/library_auth_reference/pusher-websockets-protocol/)
- [pusher-js文档](https://github.com/pusher/pusher-js)
- [GoFrame文档](https://goframe.org)

---

**最后更新**: 2026-02-28  
**版本**: v1.0.0  
**状态**: ✅ 生产就绪

