// Package system
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package system

import (
	"context"
	"fmt"

	"devinggo/internal/dao"
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
)

type sSystemDeptLeader struct {
	base.BaseService
}

func init() {
	service.RegisterSystemDeptLeader(NewSystemDeptLeader())
}

func NewSystemDeptLeader() *sSystemDeptLeader {
	return &sSystemDeptLeader{}
}

func (s *sSystemDeptLeader) Model(ctx context.Context) *gdb.Model {
	return dao.SystemDeptLeader.Ctx(ctx).Hook(hook.Default()).Cache(orm.SetCacheOption(ctx)).OnConflict("dept_id", "user_id")
}

func (s *sSystemDeptLeader) GetPageList(ctx context.Context, req *model.PageListReq, search *req.SystemDeptLeaderSearch) (res []*res.SystemDeptLeaderInfo, total int, err error) {
	m := service.SystemUser().Model(ctx).Fields(
		fmt.Sprintf(`"%s".*`, dao.SystemUser.Table()),
		fmt.Sprintf(`"%s"."created_at" as leader_add_time`, dao.SystemDeptLeader.Table()),
	)
	m = m.InnerJoinOnFields(
		dao.SystemDeptLeader.Table(),
		dao.SystemUser.Columns().Id,
		"=",
		dao.SystemDeptLeader.Columns().UserId,
	)
	m = m.Where(dao.SystemDeptLeader.Table()+"."+dao.SystemDeptLeader.Columns().DeptId+" = ?", search.DeptId)
	if !g.IsEmpty(search.Username) {
		m = m.Where(dao.SystemUser.Table()+"."+dao.SystemUser.Columns().Username+" like ?", "%"+search.Username+"%")
	}
	if !g.IsEmpty(search.Nickname) {
		m = m.Where(dao.SystemUser.Table()+"."+dao.SystemUser.Columns().Nickname+" like ?", "%"+search.Nickname+"%")
	}

	if !g.IsEmpty(search.Status) {
		m = m.Where(dao.SystemUser.Table()+"."+dao.SystemUser.Columns().Status+" = ?", search.Status)
	}
	req.OrderBy = fmt.Sprintf(`"%s"."%s"`, dao.SystemUser.Table(), dao.SystemUser.Columns().Id)
	err = orm.NewQuery(m).WithPageListReq(req).ScanAndCount(&res, &total)
	if utils.IsError(err) {
		return nil, 0, err
	}
	return
}
