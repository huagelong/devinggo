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
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

var groupKey = "websocket"

// 删除心跳数据
func RemoveClientIdHeartbeatTime4Redis(ctx context.Context, clientId string) (err error) {
	if g.IsEmpty(clientId) {
		return
	}
	key := "ClientId2HeartbeatTime"
	_, err = g.Redis(groupKey).Do(ctx, "HDEL", key, clientId)
	if err != nil {
		glob.WithWsLog().Warning(ctx, "ClientId2HeartbeatTime HDEL error:", err)
		return
	}
	return
}

// 更新心跳数据
func UpdateClientIdHeartbeatTime4Redis(ctx context.Context, clientId string, currentTime uint64) (err error) {
	if g.IsEmpty(clientId) {
		return
	}
	key := "ClientId2HeartbeatTime"
	_, err = g.Redis(groupKey).HSet(ctx, key, g.Map{clientId: currentTime})
	if err != nil {
		glob.WithWsLog().Warning(ctx, "ClientId2HeartbeatTime HSET error:", err)
		return
	}
	return
}

// 清理心跳过期数据,清除所有客户端数据
func ClearExpire4Redis(ctx context.Context) (err error) {
	lockKey := "ClearExpire4Redis"
	rs, err := g.Redis(groupKey).SetNX(ctx, lockKey, 1)
	if err != nil {
		glob.WithWsLog().Warning(ctx, "ClearExpire4Redis SetNX error:", err)
		return
	}
	if !rs {
		return
	}
	g.Redis(groupKey).Expire(ctx, lockKey, 3600)
	key := "ClientId2HeartbeatTime"
	value, err := g.Redis(groupKey).HGetAll(ctx, key)
	if err != nil {
		glob.WithWsLog().Warning(ctx, "ClientId2HeartbeatTime HGETALL error:", err)
		g.Redis(groupKey).Del(ctx, lockKey)
		return
	}
	for clientId, currentTime := range value.Map() {
		now := uint(gtime.Now().Unix())
		currentTimeInt := gconv.Uint(currentTime)
		glob.WithWsLog().Debug(ctx, "ClearExpire4Redis:", clientId)
		if heartbeatExpirationTime+currentTimeInt <= now {
			ClearClientId4Redis(ctx, clientId)
		}
	}
	g.Redis(groupKey).Del(ctx, lockKey)
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
	key := "ClientId2ServerName:" + clientId
	_, err = g.Redis(groupKey).Del(ctx, key)
	if err != nil {
		glob.WithWsLog().Warning(ctx, "DeleteServerNameByClientId error:", err)
	}
	return
}

// 获取客户端订阅数据
func GetServerNameByClientId4Redis(ctx context.Context, clientId string) string {
	key := "ClientId2ServerName:" + clientId
	serverName, err := g.Redis(groupKey).Get(ctx, key)

	if err != nil {
		glob.WithWsLog().Warning(ctx, "GetServerNameByClientId4Redis error:", err)
		return ""
	}
	return gconv.String(serverName)
}

// 添加客户端订阅数据,并确认在那个服务器上
func AddServerNameClientId4Redis(ctx context.Context, clientId string, serverName string) (err error) {
	key := "ClientId2ServerName:" + clientId
	g.Redis(groupKey).Set(ctx, key, serverName)
	serverNameKey := "ServerNames"
	_, err = g.Redis(groupKey).Do(ctx, "SADD", serverNameKey, serverName)
	if err != nil {
		glob.WithWsLog().Warning(ctx, "ServerNames SADD error:", err)
	}
	return
}

// 获取所有服务器名称
func GetAllServerNames(ctx context.Context) []string {

	serverNameKey := "ServerNames"
	ls, err := g.Redis(groupKey).Do(ctx, "SMEMBERS", serverNameKey)
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
	g.Redis(groupKey).Do(ctx, "MULTI")
	keyTopics := "Topics"
	_, err = g.Redis(groupKey).Do(ctx, "SADD", keyTopics, topic)
	if err != nil {
		glob.WithWsLog().Warning(ctx, "Topics SADD error:", err)
		g.Redis(groupKey).Do(ctx, "DISCARD")
		return
	}

	key := "Topic2ClientId:" + topic
	_, err = g.Redis(groupKey).Do(ctx, "SADD", key, clientId)
	if err != nil {
		glob.WithWsLog().Warning(ctx, "Topic2ClientId SADD error:", err)
		g.Redis(groupKey).Do(ctx, "DISCARD")
		return
	}

	keyCLient2Topic := "ClientId2Topic:" + clientId
	_, err = g.Redis(groupKey).Do(ctx, "SADD", keyCLient2Topic, topic)
	if err != nil {
		glob.WithWsLog().Warning(ctx, "ClientId2Topic SADD error:", err)
		g.Redis(groupKey).Do(ctx, "DISCARD")
		return
	}

	g.Redis(groupKey).Do(ctx, "EXEC")

	keyServername := "Topic2ServerName:" + topic
	serverName := GetServerNameByClientId4Redis(ctx, clientId)

	if !g.IsEmpty(serverName) {
		_, err = g.Redis(groupKey).Do(ctx, "SADD", keyServername, serverName)
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
	key := "Topic2ClientId:" + topic
	_, err = g.Redis(groupKey).Do(ctx, "SREM", key, clientId)
	if err != nil {
		glob.WithWsLog().Warning(ctx, "Topic2ClientId SREM error:", err)
		return
	}
	keyServername := "Topic2ServerName:" + topic
	serverName := GetServerNameByClientId4Redis(ctx, clientId)
	if !g.IsEmpty(serverName) {
		_, err = g.Redis(groupKey).Do(ctx, "SREM", keyServername, serverName)
		if err != nil {
			glob.WithWsLog().Warning(ctx, "Topic2ServerName SREM error:", err)
			return
		}
	}

	keyCLient2Topic := "ClientId2Topic:" + clientId
	_, err = g.Redis(groupKey).Do(ctx, "SREM", keyCLient2Topic, topic)
	if err != nil {
		glob.WithWsLog().Warning(ctx, "ClientId2Topic SADD error:", err)
		return
	}

	keyTopic2ClientId := "Topic2ClientId:" + topic
	count, err := g.Redis(groupKey).Do(ctx, "SCARD", keyTopic2ClientId)
	if err != nil {
		glob.WithWsLog().Warning(ctx, "Topic2ClientId SCARD error:", err)
		return
	}

	if gconv.Int(count) == 0 {
		keyTopics := "Topics"
		_, err = g.Redis(groupKey).Do(ctx, "SREM", keyTopics, topic)
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

	keyServername := "Topic2ServerName:" + topic
	ls, err := g.Redis(groupKey).Do(ctx, "SMEMBERS", keyServername)
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

	key := "ClientId2Topic:" + clientId
	ls, err := g.Redis(groupKey).Do(ctx, "SMEMBERS", key)
	if err != nil {
		glob.WithWsLog().Warning(ctx, "ClientId2Topic SMEMBERS error:", err)
		return nil
	}
	return gconv.Strings(ls)
}

// 获取所有主题
func GetAllTopics(ctx context.Context) []string {

	keyTopics := "Topics"
	ls, err := g.Redis(groupKey).Do(ctx, "SMEMBERS", keyTopics)
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
	keyTopics := "Topics"
	ls, err := g.Redis(groupKey).Do(ctx, "SISMEMBER", keyTopics, topic)
	if err != nil {
		glob.WithWsLog().Warning(ctx, "Topics SMEMBERS error:", err)
		return false
	}
	return gconv.Int(ls) == 1
}
