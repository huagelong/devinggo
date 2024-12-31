// Package system
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package system

import (
	"devinggo/internal/dao"
	"devinggo/internal/logic/base"
	"devinggo/internal/model/do"
	"devinggo/modules/system/service"
	"context"
	"github.com/gogf/gf/v2/database/gdb"
)

type sSystemQueueMessageReceive struct {
	base.BaseService
}

func init() {
	service.RegisterSystemQueueMessageReceive(NewSystemQueueMessageReceive())
}

func NewSystemQueueMessageReceive() *sSystemQueueMessageReceive {
	return &sSystemQueueMessageReceive{}
}

func (s *sSystemQueueMessageReceive) Model(ctx context.Context) *gdb.Model {
	return dao.SystemQueueMessageReceive.Ctx(ctx)
}

func (s *sSystemQueueMessageReceive) UpdateReadStatus(ctx context.Context, ids []uint64, userId uint64, value int) (err error) {
	data := &do.SystemQueueMessageReceive{
		ReadStatus: value,
	}
	_, err = s.Model(ctx).Data(data).Where(dao.SystemQueueMessageReceive.Columns().ReadStatus+" !=? ", value).WhereIn(dao.SystemQueueMessageReceive.Columns().MessageId, ids).Where(dao.SystemQueueMessageReceive.Columns().UserId, userId).Update()
	return err
}
