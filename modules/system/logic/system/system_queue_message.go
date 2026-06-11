// Package system
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package system

import (
	"context"
	"encoding/json"
	"fmt"

	"devinggo/internal/dao"
	"devinggo/internal/model/do"
	"devinggo/internal/model/entity"
	"devinggo/modules/system/logic/base"
	"devinggo/modules/system/model"
	"devinggo/modules/system/model/req"
	"devinggo/modules/system/model/res"
	"devinggo/modules/system/pkg/orm"
	"devinggo/modules/system/pkg/utils"
	"devinggo/modules/system/pkg/websocket"
	"devinggo/modules/system/service"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

type sSystemQueueMessage struct {
	base.BaseService
}

func init() {
	service.RegisterSystemQueueMessage(NewSystemQueueMessage())
}

func NewSystemQueueMessage() *sSystemQueueMessage {
	return &sSystemQueueMessage{}
}

func (s *sSystemQueueMessage) Model(ctx context.Context) *gdb.Model {
	return dao.SystemQueueMessage.Ctx(ctx).OnConflict("id")
}

func (s *sSystemQueueMessage) GetReceiveUserPageList(ctx context.Context, req *model.PageListReq, messageId int64) (rs []*res.MessageReceiveUser, total int, err error) {
	m := service.SystemUser().Model(ctx).Fields(
		fmt.Sprintf(`"%s"."read_status" as read_status_int`, dao.SystemQueueMessageReceive.Table()),
		fmt.Sprintf(`"%s"."username"`, dao.SystemUser.Table()),
		fmt.Sprintf(`"%s"."nickname"`, dao.SystemUser.Table()),
	).InnerJoinOnFields(dao.SystemQueueMessageReceive.Table(), "id", "=", "user_id")
	m = m.WherePrefix(dao.SystemQueueMessageReceive.Table(), dao.SystemQueueMessageReceive.Columns().MessageId, messageId)
	req.OrderBy = fmt.Sprintf(`"%s"."%s"`, dao.SystemUser.Table(), dao.SystemUser.Columns().CreatedAt)
	err = orm.NewQuery(m).WithPageListReq(req).ScanAndCount(&rs, &total)
	if utils.IsError(err) {
		return
	}
	return
}

func (s *sSystemQueueMessage) GetPageList(ctx context.Context, req *model.PageListReq, userId int64, params *req.SystemQueueMessageSearch) (rs []*res.SystemQueueMessage, total int, err error) {
	readStatus := params.ReadStatus
	contentType := params.ContentType
	title := params.Title
	createdAtArr := params.CreatedAt

	readStatusInt := 0
	if readStatus != "all" {
		readStatusInt = gconv.Int(readStatus)
	}
	m := service.SystemQueueMessageReceive().Model(ctx).InnerJoinOnFields(dao.SystemQueueMessage.Table(), "message_id", "=", "id")

	if !g.IsEmpty(contentType) && contentType != "all" {
		m = m.WherePrefix(dao.SystemQueueMessage.Table(), dao.SystemQueueMessage.Columns().ContentType, contentType)
	}
	if !g.IsEmpty(title) {
		m = m.WherePrefixLike(dao.SystemQueueMessage.Table(), dao.SystemQueueMessage.Columns().Title, "%"+title+"%")
	}

	if !g.IsEmpty(createdAtArr) {
		if len(createdAtArr) > 0 {
			m = m.WherePrefixGTE(dao.SystemQueueMessage.Table(), dao.SystemQueueMessage.Columns().CreatedAt, createdAtArr[0]+" 00:00:00")
		}

		if len(createdAtArr) > 1 {
			m = m.WherePrefixLTE(dao.SystemQueueMessage.Table(), dao.SystemQueueMessage.Columns().CreatedAt, createdAtArr[1]+"23:59:59")
		}
	}

	m = m.WherePrefix(dao.SystemQueueMessageReceive.Table(), dao.SystemQueueMessageReceive.Columns().UserId, userId)
	if readStatusInt != 0 {
		m = m.WherePrefix(dao.SystemQueueMessageReceive.Table(), dao.SystemQueueMessageReceive.Columns().ReadStatus, readStatusInt)
	}

	req.OrderBy = fmt.Sprintf(`"%s"."%s"`, dao.SystemQueueMessageReceive.Table(), dao.SystemQueueMessageReceive.Columns().MessageId)
	var receiveRes []*entity.SystemQueueMessageReceive
	err = orm.NewQuery(m).WithPageListReq(req).ScanAndCount(&receiveRes, &total)
	if utils.IsError(err) {
		return nil, 0, err
	}
	rs = make([]*res.SystemQueueMessage, 0)
	if !g.IsEmpty(receiveRes) {
		for _, v := range receiveRes {
			systemQueueMessageTmp := &res.SystemQueueMessage{}
			newUserId := v.UserId
			messageId := v.MessageId

			errorm := s.Model(ctx).Where("id", messageId).Scan(systemQueueMessageTmp)
			if utils.IsError(errorm) {
				g.Log().Error(ctx, errorm)
				continue
			}
			// 设置阅读状态
			systemQueueMessageTmp.ReadStatus = v.ReadStatus

			userInfo, errorm := service.SystemUser().GetInfoById(ctx, newUserId)
			if utils.IsError(errorm) {
				g.Log().Error(ctx, errorm)
				continue
			}
			if !g.IsEmpty(userInfo) {
				var userInfoTmp *model.UserRelate
				if errconv := gconv.Struct(userInfo, &userInfoTmp); errconv != nil {
					g.Log().Error(ctx, errconv)
					continue
				} else {
					systemQueueMessageTmp.SendUser = *userInfoTmp
				}
			}
			rs = append(rs, systemQueueMessageTmp)
		}
	}

	return
}

func (s *sSystemQueueMessage) DeletesRelated(ctx context.Context, ids []int64, userId int64) (err error) {
	_, err = service.SystemQueueMessageReceive().Model(ctx).Where(dao.SystemQueueMessageReceive.Columns().UserId, userId).WhereIn(dao.SystemQueueMessageReceive.Columns().MessageId, ids).Delete()
	if utils.IsError(err) {
		return err
	}
	return
}

const (
	wsChannelAdminUser   = "private-adminuser-%d"
	wsEventNotification  = "notification:new"
	wsBatchSize          = 100
	wsMaxConcurrentSends = 20
)

func (s *sSystemQueueMessage) SendMessage(ctx context.Context, sendReq *req.SystemQueueMessagesSend, sendUserId int64, contentType string) (err error, messageId int64) {
	data := do.SystemQueueMessage{
		ContentType: contentType,
		Title:       sendReq.Title,
		Content:     sendReq.Content,
		SendBy:      sendUserId,
		CreatedBy:   sendUserId,
	}
	rs, err := s.Model(ctx).Data(data).Insert()

	if utils.IsError(err) {
		return
	}

	messageIdTmp, err := rs.LastInsertId()
	if utils.IsError(err) {
		return
	}
	messageId = int64(messageIdTmp)
	//异步处理
	utils.SafeGo(ctx, func(ctx context.Context) {
		if !g.IsEmpty(sendReq.Users) {
			// 限制并发发送数量，防止对 WebSocket 服务造成瞬时压力
			sem := make(chan struct{}, wsMaxConcurrentSends)
			for _, v := range sendReq.Users {
				sem <- struct{}{}
				go func(userId int64) {
					defer func() { <-sem }()
					receiveData := do.SystemQueueMessageReceive{
						MessageId: messageId,
						UserId:    userId,
					}
					_, _ = service.SystemQueueMessageReceive().Model(ctx).Data(receiveData).Insert()
					s.sendWs(ctx, userId, messageId)
				}(v)
			}
			// 等待所有 goroutine 完成
			for i := 0; i < cap(sem); i++ {
				sem <- struct{}{}
			}
		} else {
			//获取所有有效用户，循环插入
			s.saveAllUserMessageReceive(ctx, messageId)
		}
	})
	return
}

func (s *sSystemQueueMessage) sendWs(ctx context.Context, userId int64, messageId int64) {
	message := &entity.SystemQueueMessage{}
	err := s.Model(ctx).Where("id", messageId).Scan(message)
	if utils.IsError(err) || g.IsEmpty(message.Id) {
		g.Log().Error(ctx, "sendWs: query message failed:", err)
		return
	}

	notificationData := g.Map{
		"id":           message.Id,
		"title":        message.Title,
		"content":      message.Content,
		"content_type": message.ContentType,
		"send_by":      message.SendBy,
	}
	if message.CreatedAt != nil {
		notificationData["created_at"] = message.CreatedAt.Format("Y-m-d H:i:s")
	}

	sendUserInfo, err := service.SystemUser().GetInfoById(ctx, message.SendBy)
	if err == nil && !g.IsEmpty(sendUserInfo) {
		notificationData["send_user"] = g.Map{
			"id":       sendUserInfo.Id,
			"nickname": sendUserInfo.Nickname,
			"username": sendUserInfo.Username,
		}
	}

	dataJson, err := json.Marshal(notificationData)
	if err != nil {
		g.Log().Error(ctx, "sendWs: marshal notification data failed:", err)
		return
	}

	channel := fmt.Sprintf(wsChannelAdminUser, userId)
	pusherResponse := &websocket.PusherResponse{
		Event:   wsEventNotification,
		Channel: channel,
		Data:    string(dataJson),
	}

	websocket.SendToChannel(channel, pusherResponse)

	topicMsg := &websocket.TopicWResponse{
		Topic:          channel,
		PusherResponse: pusherResponse,
	}
	if err := websocket.PublishChannelMessage(ctx, channel, topicMsg); err != nil {
		g.Log().Error(ctx, "sendWs: cross-server publish failed:", err)
	}
}

func (s *sSystemQueueMessage) saveAllUserMessageReceive(ctx context.Context, messageId int64) (err error) {
	page := 1
	sem := make(chan struct{}, wsMaxConcurrentSends)

	for {
		var userList []*res.SystemUserSimple
		m := service.SystemUser().Model(ctx).Where(dao.SystemUser.Columns().Status, 1).OrderDesc("id")
		err = m.Page(page, wsBatchSize).Scan(&userList)
		if utils.IsError(err) {
			g.Log().Error(ctx, "saveAllUserMessageReceive: query users failed:", err)
			return
		}
		if g.IsEmpty(userList) {
			break
		}

		for _, v := range userList {
			sem <- struct{}{}
			go func(userId int64) {
				defer func() { <-sem }()
				receiveData := do.SystemQueueMessageReceive{
					MessageId: messageId,
					UserId:    userId,
				}
				_, _ = service.SystemQueueMessageReceive().Model(ctx).Data(receiveData).Insert()
				s.sendWs(ctx, userId, messageId)
			}(v.Id)
		}

		page++
	}

	// 等待所有 goroutine 完成
	for i := 0; i < cap(sem); i++ {
		sem <- struct{}{}
	}
	return
}
