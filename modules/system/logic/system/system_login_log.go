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
	"devinggo/modules/system/model"
	"devinggo/modules/system/model/req"
	"devinggo/modules/system/model/res"
	"devinggo/modules/system/pkg/orm"
	"devinggo/modules/system/pkg/utils"
	"devinggo/modules/system/pkg/utils/location"
	"devinggo/modules/system/pkg/utils/request"
	"devinggo/modules/system/pkg/utils/user_agent"
	"devinggo/modules/system/service"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

type sSystemLoginLog struct {
	base.BaseService
}

func init() {
	service.RegisterSystemLoginLog(NewSystemLoginLog())
}

func NewSystemLoginLog() *sSystemLoginLog {
	return &sSystemLoginLog{}
}

func (s *sSystemLoginLog) Model(ctx context.Context) *gdb.Model {
	return dao.SystemLoginLog.Ctx(ctx).OnConflict("id")
}

func (s *sSystemLoginLog) GetPageList(ctx context.Context, req *model.PageListReq, username string) (res []*res.SystemLoginLog, total int, err error) {
	err = orm.GetPageList(s.Model(ctx), req, g.Map{"username": username}).ScanAndCount(&res, &total, false)
	if utils.IsError(err) {
		return nil, 0, err
	}
	return
}

func (s *sSystemLoginLog) Push(ctx context.Context, username string, err error) {

	if g.IsEmpty(username) {
		return
	}

	r := request.GetHttpRequest(ctx)
	if r == nil {
		g.Log().Warningf(ctx, "ctx not http request")
		return
	}

	//g.Log().Debug(ctx, "err:", err)

	clientIp := location.GetClientIp(r)
	ipData, localErr := location.GetLocation(ctx, clientIp)
	area := "本地"
	if localErr != nil {
		g.Log().Debugf(ctx, "location.GetLocation clientIp:%v, err:%+v", clientIp, localErr)
	} else {
		if !g.IsEmpty(ipData) {
			area = ipData.Area
		}
	}

	loginStatus := 1
	message := "登录成功"
	if err != nil {
		loginStatus = 2
		message = err.Error()
	}

	userAgent := user_agent.GetUserAgent(ctx)

	systemLoginLog := &do.SystemLoginLog{
		Username:   username,
		Ip:         clientIp,          // 登录IP地址
		IpLocation: area,              // IP所属地
		Os:         userAgent.Os,      // 操作系统
		Browser:    userAgent.Browser, // 浏览器
		Status:     loginStatus,       // 登录状态 (1成功 2失败)
		Message:    message,           // 提示消息
		LoginTime:  gtime.Now(),       // 登录时间
		Remark:     "",                // 备注
	}
	s.Model(ctx).Data(systemLoginLog).Insert()
}

func (s *sSystemLoginLog) handleSearch(ctx context.Context, in *req.SystemLoginLogSearch) (m *gdb.Model) {

	m = s.Model(ctx)

	if !g.IsEmpty(in.Ip) {
		m = m.Where("ip", in.Ip)
	}

	if !g.IsEmpty(in.Status) {
		m = m.Where("status", in.Status)
	}

	if !g.IsEmpty(in.Username) {
		m = m.Where("username like ? ", "%"+in.Username+"%")
	}
	if !g.IsEmpty(in.LoginTime) {
		if len(in.LoginTime) > 0 {
			m = m.WhereGTE("login_time", in.LoginTime[0]+" 00:00:00")
		}
		if len(in.LoginTime) > 1 {
			m = m.WhereLTE("login_time", in.LoginTime[1]+"23:59:59")
		}
	}
	return
}

func (s *sSystemLoginLog) GetPageListForSearch(ctx context.Context, req *model.PageListReq, in *req.SystemLoginLogSearch) (rs []*res.SystemLoginLog, total int, err error) {
	m := s.handleSearch(ctx, in)
	var entity []*entity.SystemLoginLog
	err = orm.GetPageList(m, req).ScanAndCount(&entity, &total, false)
	if utils.IsError(err) {
		return nil, 0, err
	}
	rs = make([]*res.SystemLoginLog, 0)
	if !g.IsEmpty(entity) {
		if err = gconv.Structs(entity, &rs); err != nil {
			return nil, 0, err
		}
	}
	return
}
