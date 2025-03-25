// Package websocket
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package websocket

import (
	"context"
	"devinggo/modules/system/pkg/websocket/glob"
	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

// Redis key 常量定义
const (
	KeyClientIdHeartbeatTime = "ClientId2HeartbeatTime" // 客户端心跳时间
	KeyClearExpireLock       = "ClearExpire4Redis"      // 清理过期数据的锁
	KeyClientId2ServerName   = "ClientId2ServerName:"   // 客户端对应的服务器名称
	KeyServerNames           = "ServerNames"            // 所有服务器名称集合
	KeyTopics                = "Topics"                 // 所有主题集合
	KeyTopic2ClientId        = "Topic2ClientId:"        // 主题对应的客户端集合
	KeyClientId2Topic        = "ClientId2Topic:"        // 客户端对应的主题集合
	KeyTopic2ServerName      = "Topic2ServerName:"      // 主题对应的服务器名称集合
)

func getRedisClient() *gredis.Redis {
	return g.Redis("websocket")
}

// 删除心跳数据
func RemoveClientIdHeartbeatTime4Redis(ctx context.Context, clientId string) (err error) {
	if g.IsEmpty(clientId) {
		return
	}
	_, err = getRedisClient().Do(ctx, "HDEL", KeyClientIdHeartbeatTime, clientId)
	if err != nil {
		glob.WithWsLog().Warning(ctx, "ClientId2HeartbeatTime HDEL error:", err)
		return
	}
	return
}

// 更新心跳数据
func UpdateClientIdHeartbeatTime4Redis(ctx context.Context, clientId string, currentTime int64) (err error) {
	if g.IsEmpty(clientId) {
		return
	}
	_, err = getRedisClient().HSet(ctx, KeyClientIdHeartbeatTime, g.Map{clientId: currentTime})
	if err != nil {
		glob.WithWsLog().Warning(ctx, "ClientId2HeartbeatTime HSET error:", err)
		return
	}
	return
}

// 清理心跳过期数据,清除所有客户端数据
func ClearExpire4Redis(ctx context.Context) (err error) {
	rs, err := getRedisClient().SetNX(ctx, KeyClearExpireLock, 1)
	if err != nil {
		glob.WithWsLog().Warning(ctx, "ClearExpire4Redis SetNX error:", err)
		return
	}
	if !rs {
		return
	}
	getRedisClient().Expire(ctx, KeyClearExpireLock, 3600)
	value, err := getRedisClient().HGetAll(ctx, KeyClientIdHeartbeatTime)
	if err != nil {
		glob.WithWsLog().Warning(ctx, "ClientId2HeartbeatTime HGETALL error:", err)
		getRedisClient().Del(ctx, KeyClearExpireLock)
		return
	}
	for clientId, currentTime := range value.Map() {
		now := int(gtime.Now().Unix())
		currentTimeInt := gconv.Int(currentTime)
		glob.WithWsLog().Debug(ctx, "ClearExpire4Redis:", clientId)
		if heartbeatExpirationTime+currentTimeInt <= now {
			ClearClientId4Redis(ctx, clientId)
		}
	}
	getRedisClient().Del(ctx, KeyClearExpireLock)
	return
}

// 清除所有客户端数据，包含心跳数据，订阅数据，全局数据
func ClearClientId4Redis(ctx context.Context, clientId string) (err error) {
	err = RemoveClientIdHeartbeatTime4Redis(ctx, clientId)
	for _, topic := range GetAllTopicByClientId(ctx, clientId) {
		QuitTopic4Redis(ctx, clientId, topic)
	}
	err = DeleteServerNameByClientId4Redis(ctx, clientId)
	return
}

// 删除客户端订阅数据
func DeleteServerNameByClientId4Redis(ctx context.Context, clientId string) (err error) {
	key := KeyClientId2ServerName + clientId
	_, err = getRedisClient().Del(ctx, key)
	if err != nil {
		glob.WithWsLog().Warning(ctx, "DeleteServerNameByClientId error:", err)
	}
	return
}

// 获取客户端订阅数据
func GetServerNameByClientId4Redis(ctx context.Context, clientId string) string {
	key := KeyClientId2ServerName + clientId
	serverName, err := getRedisClient().Get(ctx, key)

	if err != nil {
		glob.WithWsLog().Warning(ctx, "GetServerNameByClientId4Redis error:", err)
		return ""
	}
	return gconv.String(serverName)
}

// 添加客户端订阅数据,并确认在那个服务器上
func AddServerNameClientId4Redis(ctx context.Context, clientId string, serverName string) (err error) {
	key := KeyClientId2ServerName + clientId
	getRedisClient().Set(ctx, key, serverName)
	_, err = getRedisClient().Do(ctx, "SADD", KeyServerNames, serverName)
	if err != nil {
		glob.WithWsLog().Warning(ctx, "ServerNames SADD error:", err)
	}
	return
}

// 获取所有服务器名称
func GetAllServerNames(ctx context.Context) []string {
	ls, err := getRedisClient().Do(ctx, "SMEMBERS", KeyServerNames)
	if err != nil {
		glob.WithWsLog().Warning(ctx, "ServerNames error:", err)
		return nil
	}
	return gconv.Strings(ls)
}

// 加入主题
func JoinTopic4Redis(ctx context.Context, clientId string, topic string) (err error) {
	if g.IsEmpty(topic) {
		return
	}
	getRedisClient().Do(ctx, "MULTI")
	_, err = getRedisClient().Do(ctx, "SADD", KeyTopics, topic)
	if err != nil {
		glob.WithWsLog().Warning(ctx, "Topics SADD error:", err)
		getRedisClient().Do(ctx, "DISCARD")
		return
	}

	key := KeyTopic2ClientId + topic
	_, err = getRedisClient().Do(ctx, "SADD", key, clientId)
	if err != nil {
		glob.WithWsLog().Warning(ctx, "Topic2ClientId SADD error:", err)
		getRedisClient().Do(ctx, "DISCARD")
		return
	}

	keyCLient2Topic := KeyClientId2Topic + clientId
	_, err = getRedisClient().Do(ctx, "SADD", keyCLient2Topic, topic)
	if err != nil {
		glob.WithWsLog().Warning(ctx, "ClientId2Topic SADD error:", err)
		getRedisClient().Do(ctx, "DISCARD")
		return
	}

	getRedisClient().Do(ctx, "EXEC")

	keyServername := KeyTopic2ServerName + topic
	serverName := GetServerNameByClientId4Redis(ctx, clientId)

	if !g.IsEmpty(serverName) {
		_, err = getRedisClient().Do(ctx, "SADD", keyServername, serverName)
		if err != nil {
			glob.WithWsLog().Warning(ctx, "Topic2ServerName SADD error:", err)
			return
		}
	}
	return
}

// 退出主题
func QuitTopic4Redis(ctx context.Context, clientId string, topic string) (err error) {
	if g.IsEmpty(topic) {
		return
	}
	key := KeyTopic2ClientId + topic
	_, err = getRedisClient().Do(ctx, "SREM", key, clientId)
	if err != nil {
		glob.WithWsLog().Warning(ctx, "Topic2ClientId SREM error:", err)
		return
	}

	keyServername := KeyTopic2ServerName + topic
	serverName := GetServerNameByClientId4Redis(ctx, clientId)
	if !g.IsEmpty(serverName) {
		_, err = getRedisClient().Do(ctx, "SREM", keyServername, serverName)
		if err != nil {
			glob.WithWsLog().Warning(ctx, "Topic2ServerName SREM error:", err)
			return
		}
	}

	keyCLient2Topic := KeyClientId2Topic + clientId
	_, err = getRedisClient().Do(ctx, "SREM", keyCLient2Topic, topic)
	if err != nil {
		glob.WithWsLog().Warning(ctx, "ClientId2Topic SADD error:", err)
		return
	}

	keyTopic2ClientId := KeyTopic2ClientId + topic
	count, err := getRedisClient().Do(ctx, "SCARD", keyTopic2ClientId)
	if err != nil {
		glob.WithWsLog().Warning(ctx, "Topic2ClientId SCARD error:", err)
		return
	}

	if gconv.Int(count) == 0 {
		_, err = getRedisClient().Do(ctx, "SREM", KeyTopics, topic)
		if err != nil {
			glob.WithWsLog().Warning(ctx, "Topics SREM error:", err)
			return
		}
	}
	return
}

// 获取主题的服务器名称
func GetAllServerNameByTopic(ctx context.Context, topic string) []string {
	if g.IsEmpty(topic) {
		return nil
	}

	keyServername := KeyTopic2ServerName + topic
	ls, err := getRedisClient().Do(ctx, "SMEMBERS", keyServername)
	if err != nil {
		glob.WithWsLog().Warning(ctx, "Topic2ServerName error:", err)
		return nil
	}
	return gconv.Strings(ls)
}

// 获取客户端订阅的所有主题
func GetAllTopicByClientId(ctx context.Context, clientId string) []string {
	if g.IsEmpty(clientId) {
		return nil
	}

	key := KeyClientId2Topic + clientId
	ls, err := getRedisClient().Do(ctx, "SMEMBERS", key)
	if err != nil {
		glob.WithWsLog().Warning(ctx, "ClientId2Topic SMEMBERS error:", err)
		return nil
	}
	return gconv.Strings(ls)
}

// 获取所有主题
func GetAllTopics(ctx context.Context) []string {
	ls, err := getRedisClient().Do(ctx, "SMEMBERS", KeyTopics)
	if err != nil {
		glob.WithWsLog().Warning(ctx, "Topics SMEMBERS error:", err)
		return nil
	}
	return gconv.Strings(ls)
}

// 判断主题是否存在
func isTopicExist(ctx context.Context, topic string) bool {
	if g.IsEmpty(topic) {
		return false
	}
	ls, err := getRedisClient().Do(ctx, "SISMEMBER", KeyTopics, topic)
	if err != nil {
		glob.WithWsLog().Warning(ctx, "Topics SMEMBERS error:", err)
		return false
	}
	return gconv.Int(ls) == 1
}
