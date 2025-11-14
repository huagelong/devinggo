// Package system
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package system

import (
	"context"
	"devinggo/internal/dao"
	"devinggo/internal/model/entity"
	"devinggo/modules/system/logic/base"
	"devinggo/modules/system/pkg/hook"
	"devinggo/modules/system/pkg/orm"
	"devinggo/modules/system/pkg/utils"
	"devinggo/modules/system/service"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

type sSystemUserPost struct {
	base.BaseService
}

func init() {
	service.RegisterSystemUserPost(NewSystemUserPost())
}

func NewSystemUserPost() *sSystemUserPost {
	return &sSystemUserPost{}
}

func (s *sSystemUserPost) Model(ctx context.Context) *gdb.Model {
	return dao.SystemUserPost.Ctx(ctx).Hook(hook.Bind()).Cache(orm.SetCacheOption(ctx)).OnConflict("user_id", "post_id")
}

func (s *sSystemUserPost) GetPostIdsByUserId(ctx context.Context, userId int64) (postIds []int64, err error) {
	var list []*entity.SystemUserPost
	err = s.Model(ctx).Fields(dao.SystemUserPost.Columns().PostId).Where(dao.SystemUserPost.Columns().UserId, userId).Scan(&list)
	if utils.IsError(err) {
		return
	}
	if g.IsEmpty(list) {
		return
	}
	for _, item := range list {
		postIds = append(postIds, item.PostId)
	}
	return
}
