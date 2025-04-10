// Package system
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package system

import (
	"context"
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
	return dao.SystemDeptLeader.Ctx(ctx).Hook(hook.Bind()).Cache(orm.SetCacheOption(ctx)).OnConflict("dept_id", "user_id")
}

func (s *sSystemDeptLeader) GetPageList(ctx context.Context, req *model.PageListReq, search *req.SystemDeptLeaderSearch) (res []*res.SystemDeptLeaderInfo, total int, err error) {
	m := service.SystemUser().Model(ctx).Fields("system_user.*", "system_dept_leader.created_at as leader_add_time")
	m = m.InnerJoinOnFields("system_dept_leader", "id", "=", "user_id")
	m = m.Where("system_dept_leader.dept_id =?", search.DeptId)
	if !g.IsEmpty(search.Username) {
		m = m.Where("system_user.username like ?", "%"+search.Username+"%")
	}
	if !g.IsEmpty(search.Nickname) {
		m = m.Where("system_user.nickname like ?", "%"+search.Nickname+"%")
	}

	if !g.IsEmpty(search.Status) {
		m = m.Where("system_user.status = ?", search.Status)
	}
	err = orm.GetPageList(m, req).ScanAndCount(&res, &total, false)
	if utils.IsError(err) {
		return nil, 0, err
	}
	return
}
