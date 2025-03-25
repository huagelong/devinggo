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
	"devinggo/modules/system/consts"
	"devinggo/modules/system/logic/base"
	"devinggo/modules/system/model"
	"devinggo/modules/system/model/req"
	"devinggo/modules/system/model/res"
	"devinggo/modules/system/pkg/hook"
	"devinggo/modules/system/pkg/orm"
	"devinggo/modules/system/pkg/utils"
	"devinggo/modules/system/service"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

type sSystemNotice struct {
	base.BaseService
}

func init() {
	service.RegisterSystemNotice(NewSystemNotice())
}

func NewSystemNotice() *sSystemNotice {
	return &sSystemNotice{}
}

func (s *sSystemNotice) Model(ctx context.Context) *gdb.Model {
	return dao.SystemNotice.Ctx(ctx).Hook(hook.Bind()).Cache(orm.SetCacheOption(ctx)).OnConflict("id")
}

func (s *sSystemNotice) GetPageList(ctx context.Context, req *model.PageListReq) (res []*res.SystemNotice, total int, err error) {
	err = orm.GetPageList(s.Model(ctx), req).ScanAndCount(&res, &total, false)
	if utils.IsError(err) {
		return nil, 0, err
	}
	return
}

func (s *sSystemNotice) GetPageListForSearch(ctx context.Context, req *model.PageListReq, in *req.SystemNoticeSearch) (rs []*res.SystemNotice, total int, err error) {
	m := s.handleSearch(ctx, in)
	var entity []*entity.SystemNotice
	err = orm.GetPageList(m, req).ScanAndCount(&entity, &total, false)
	if utils.IsError(err) {
		return nil, 0, err
	}
	rs = make([]*res.SystemNotice, 0)
	if !g.IsEmpty(entity) {
		if err = gconv.Structs(entity, &rs); err != nil {
			return nil, 0, err
		}
		for _, v := range entity {
			for _, u := range rs {
				if v.Id == u.Id {
					if !g.IsEmpty(v.ReceiveUsers) {
						receiveUsersStr := gstr.Split(v.ReceiveUsers, ",")
						u.Users = gconv.Int64s(receiveUsersStr)
					}
				}
			}
		}
	}
	return
}

func (s *sSystemNotice) handleSearch(ctx context.Context, in *req.SystemNoticeSearch) (m *gdb.Model) {
	m = s.Model(ctx)
	if !g.IsEmpty(in.Title) {
		m = m.Where("title", in.Title)
	}

	if !g.IsEmpty(in.Type) {
		m = m.Where("type", in.Type)
	}
	if !g.IsEmpty(in.CreatedAt) {
		if len(in.CreatedAt) > 0 {
			m = m.WhereGTE("created_at", in.CreatedAt[0]+" 00:00:00")
		}
		if len(in.CreatedAt) > 1 {
			m = m.WhereLTE("created_at", in.CreatedAt[1]+"23:59:59")
		}
	}
	return
}

func (s *sSystemNotice) Save(ctx context.Context, in *req.SystemNoticeSave, userId int64) (id int64, err error) {
	contentType := ""
	if in.Type == 1 {
		contentType = consts.TYPE_NOTICE
	} else {
		contentType = consts.TYPE_ANNOUNCE
	}

	sendReq := &req.SystemQueueMessagesSend{
		Title:   in.Title,
		Users:   in.Users,
		Content: in.Content,
	}

	err, messageId := service.SystemQueueMessage().SendMessage(ctx, sendReq, userId, contentType)
	if err != nil {
		return
	}
	saveData := do.SystemNotice{
		Title:     in.Title,
		Type:      in.Type,
		Content:   in.Content,
		MessageId: messageId,
		Remark:    in.Remark,
	}
	if !g.IsEmpty(in.Users) {
		saveData.ReceiveUsers = gstr.Join(gconv.Strings(in.Users), ",")
	}

	rs, err := s.Model(ctx).Data(saveData).Insert()
	if utils.IsError(err) {
		return
	}
	tmpId, err := rs.LastInsertId()
	if err != nil {
		return
	}
	id = gconv.Int64(tmpId)
	return
}

func (s *sSystemNotice) GetById(ctx context.Context, id int64) (res *res.SystemNotice, err error) {
	err = s.Model(ctx).Where("id", id).Scan(&res)
	if utils.IsError(err) {
		return
	}
	return
}

func (s *sSystemNotice) Update(ctx context.Context, in *req.SystemNoticeUpdate) (err error) {
	updateData := do.SystemNotice{
		Title:   in.Title,
		Type:    in.Type,
		Content: in.Content,
		Remark:  in.Remark,
	}
	_, err = s.Model(ctx).Data(updateData).Where("id", in.Id).Update()
	if utils.IsError(err) {
		return
	}
	info, err := s.GetById(ctx, in.Id)
	if err != nil {
		return
	}
	service.SystemQueueMessage().Model(ctx).Where("id", info.MessageId).Update(g.Map{"content": in.Content, "title": in.Title})
	return
}

func (s *sSystemNotice) Delete(ctx context.Context, ids []int64) (err error) {
	_, err = s.Model(ctx).WhereIn("id", ids).Delete()
	if utils.IsError(err) {
		return err
	}
	return
}

func (s *sSystemNotice) RealDelete(ctx context.Context, ids []int64) (err error) {
	var res []*res.SystemNotice
	err = s.Model(ctx).Unscoped().WhereIn("id", ids).Scan(&res)
	if utils.IsError(err) {
		return
	}
	if g.IsEmpty(res) {
		return
	}
	for _, v := range res {
		_, err = service.SystemQueueMessage().Model(ctx).Unscoped().Where("id", v.MessageId).Delete()
		if utils.IsError(err) {
			return
		}
		_, err = service.SystemQueueMessageReceive().Model(ctx).Unscoped().Where("message_id", v.MessageId).Delete()
		if utils.IsError(err) {
			return
		}
	}
	_, err = s.Model(ctx).Unscoped().WhereIn("id", ids).Delete()
	if utils.IsError(err) {
		return
	}
	return
}

func (s *sSystemNotice) Recovery(ctx context.Context, ids []int64) (err error) {
	_, err = s.Model(ctx).Unscoped().WhereIn("id", ids).Update(g.Map{"deleted_at": nil})
	if utils.IsError(err) {
		return err
	}
	return
}
