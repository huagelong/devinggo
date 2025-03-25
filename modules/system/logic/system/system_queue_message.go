// Package system
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package system

import (
	"context"
	"devinggo/internal/dao"
	"devinggo/internal/model/do"
	"devinggo/internal/model/entity"
	"devinggo/modules/system/logic/base"
	model2 "devinggo/modules/system/model"
	"devinggo/modules/system/model/req"
	"devinggo/modules/system/model/res"
	"devinggo/modules/system/pkg/orm"
	"devinggo/modules/system/pkg/utils"
	websocket2 "devinggo/modules/system/pkg/websocket"
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

func (s *sSystemQueueMessage) GetReceiveUserPageList(ctx context.Context, req *model2.PageListReq, messageId int64) (rs []*res.MessageReceiveUser, total int, err error) {
	m := service.SystemUser().Model(ctx).Fields(dao.SystemQueueMessageReceive.Table()+".read_status as read_status_int", dao.SystemUser.Table()+".username", dao.SystemUser.Table()+".nickname").InnerJoinOnFields(dao.SystemQueueMessageReceive.Table(), "id", "=", "user_id")
	m = m.Where(dao.SystemQueueMessageReceive.Table()+".message_id", messageId)
	m = m.OrderDesc(dao.SystemUser.Table() + ".created_at")
	err = orm.GetPageList(m, req).ScanAndCount(&rs, &total, false)
	if utils.IsError(err) {
		return
	}
	return
}

func (s *sSystemQueueMessage) GetPageList(ctx context.Context, req *model2.PageListReq, userId int64, params *req.SystemQueueMessageSearch) (rs []*res.SystemQueueMessage, total int, err error) {
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
		m = m.Where(dao.SystemQueueMessage.Table()+".content_type", contentType)
	}
	if !g.IsEmpty(title) {
		m = m.WhereLike(dao.SystemQueueMessage.Table()+".title", "%"+title+"%")
	}

	if !g.IsEmpty(createdAtArr) {
		if len(createdAtArr) > 0 {
			m = m.WhereGTE(dao.SystemQueueMessage.Table()+".created_at", createdAtArr[0]+" 00:00:00")
		}

		if len(createdAtArr) > 1 {
			m = m.WhereLTE(dao.SystemQueueMessage.Table()+".created_at", createdAtArr[1]+"23:59:59")
		}
	}

	m = m.Where(dao.SystemQueueMessageReceive.Table()+".user_id", userId)
	if readStatusInt != 0 {
		m = m.Where(dao.SystemQueueMessageReceive.Table()+".read_status", readStatusInt)
	}

	m = m.OrderDesc("message_id")
	var receiveRes []*entity.SystemQueueMessageReceive
	err = orm.GetPageList(m, req).ScanAndCount(&receiveRes, &total, false)
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

			userInfo, errorm := service.SystemUser().GetInfoById(ctx, newUserId)
			if utils.IsError(errorm) {
				g.Log().Error(ctx, errorm)
				continue
			}
			if !g.IsEmpty(userInfo) {
				var userInfoTmp *model2.UserRelate
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
	if err != nil {
		return
	}
	messageId = int64(messageIdTmp)
	//异步处理
	utils.SafeGo(ctx, func(ctx context.Context) {
		if !g.IsEmpty(sendReq.Users) {
			for _, v := range sendReq.Users {
				receiveData := do.SystemQueueMessageReceive{
					MessageId: messageId,
					UserId:    v,
				}
				service.SystemQueueMessageReceive().Model(ctx).Data(receiveData).Insert()
				s.sendWs(ctx, v)
			}
		} else {
			//获取所有有效用户，循环插入
			s.saveAllUserMessageReceive(ctx, messageId, 1)
		}
	})
	return
}

func (s *sSystemQueueMessage) sendWs(ctx context.Context, userId int64) {
	pageReq := &model2.PageListReq{
		OrderBy:   "created_at",
		OrderType: "desc",
	}
	pageReq.Page = 1
	pageReq.PageSize = 5
	search := &req.SystemQueueMessageSearch{
		ReadStatus: "1",
	}
	rs, _, err := service.SystemQueueMessage().GetPageList(ctx, pageReq, userId, search)
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}
	toId := gconv.String(userId)
	clientIdWResponse := &websocket2.ClientIdWResponse{
		ID: toId,
		WResponse: &websocket2.WResponse{
			BindEvent: "ev_new_message",
			Event:     websocket2.IdMessage,
			Data:      rs,
			Code:      200,
			RequestId: "0",
		},
	}
	websocket2.PublishIdMessage(ctx, toId, clientIdWResponse)
}

func (s *sSystemQueueMessage) saveAllUserMessageReceive(ctx context.Context, messageId int64, page int) (err error) {
	var userList []*res.SystemUserSimple
	pageSize := 100
	m := service.SystemUser().Model(ctx).Where(dao.SystemUser.Columns().Status, 1).OrderDesc("id")
	err = m.Page(page, pageSize).Scan(&userList)
	if utils.IsError(err) {
		return
	}
	if g.IsEmpty(userList) {
		return
	}
	for _, v := range userList {
		receiveData := do.SystemQueueMessageReceive{
			MessageId: messageId,
			UserId:    v.Id,
		}
		service.SystemQueueMessageReceive().Model(ctx).Data(receiveData).Insert()
	}
	s.saveAllUserMessageReceive(ctx, messageId, page+1)
	return
}
