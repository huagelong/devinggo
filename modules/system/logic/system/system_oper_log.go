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
	"devinggo/modules/system/pkg/contexts"
	"devinggo/modules/system/pkg/orm"
	"devinggo/modules/system/pkg/response"
	"devinggo/modules/system/pkg/utils"
	"devinggo/modules/system/pkg/utils/location"
	"devinggo/modules/system/pkg/utils/request"
	"devinggo/modules/system/service"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

type sSystemOperLog struct {
	base.BaseService
}

func init() {
	service.RegisterSystemOperLog(NewSystemOperLog())
}

func NewSystemOperLog() *sSystemOperLog {
	return &sSystemOperLog{}
}

func (s *sSystemOperLog) Model(ctx context.Context) *gdb.Model {
	return dao.SystemOperLog.Ctx(ctx).OnConflict("id")
}

func (s *sSystemOperLog) GetPageList(ctx context.Context, req *model.PageListReq, username string) (res []*res.SystemOperLog, total int, err error) {
	err = orm.GetPageList(s.Model(ctx), req, g.Map{"username": username}).ScanAndCount(&res, &total, false)
	if utils.IsError(err) {
		return nil, 0, err
	}
	return
}

func (s *sSystemOperLog) Push(ctx context.Context) {
	r := request.GetHttpRequest(ctx)
	if r == nil {
		g.Log().Warningf(ctx, "ctx not http request")
		return
	}

	if r.GetError() == nil && r.Response.BufferLength() > 0 {
		return
	}

	userId := contexts.New().GetUserId(ctx)
	if g.IsEmpty(userId) {
		return
	}
	userInfo, err := service.SystemUser().GetInfoById(ctx, userId)
	if err != nil {
		return
	}

	serviceName := ""
	permission := contexts.New().GetPermission(ctx)
	if !g.IsEmpty(permission) {
		var systemMenuEntity *entity.SystemMenu
		systemMenuEntity, err = service.SystemMenu().GetMenuByPermission(ctx, permission)
		if err == nil {
			if !g.IsEmpty(systemMenuEntity) {
				serviceName = systemMenuEntity.Name
			}
		}
	}

	if g.IsEmpty(serviceName) {
		return
	}

	clientIp := location.GetClientIp(r)
	ipData, localErr := location.GetLocation(ctx, clientIp)
	if localErr != nil {
		g.Log().Debugf(ctx, "location.GetLocation clientIp:%v, err:%+v", clientIp, localErr)
	}
	area := "本地"
	if !g.IsEmpty(ipData) {
		area = ipData.Area
	}

	res, bizCode := response.ResponseHandler(r)
	resJson := response.Json(r, bizCode, res)
	postData := contexts.New().GetRequestBody(ctx)

	systemOperLog := &do.SystemOperLog{
		Username:     userInfo.Username,  // 用户名
		Method:       r.Method,           // 请求方式
		Router:       r.URL.RequestURI(), // 请求路由
		ServiceName:  serviceName,        // 业务名称
		Ip:           clientIp,           // 请求IP地址
		IpLocation:   area,               // IP所属地
		RequestData:  postData,           // 请求数据
		ResponseCode: bizCode.Code(),     // 响应状态码
		ResponseData: resJson,            // 响应数据
		CreatedBy:    userId,             // 创建者
		UpdatedBy:    userId,             // 更新者
		Remark:       "",                 // 备注
	}
	s.Model(ctx).Data(systemOperLog).Insert()
}

func (s *sSystemOperLog) handleSearch(ctx context.Context, in *req.SystemOperLogSearch) (m *gdb.Model) {

	m = s.Model(ctx)

	if !g.IsEmpty(in.Ip) {
		m = m.Where("ip", in.Ip)
	}

	if !g.IsEmpty(in.Status) {
		m = m.Where("status", in.Status)
	}

	if !g.IsEmpty(in.ServiceName) {
		m = m.Where("service_name like ? ", "%"+in.ServiceName+"%")
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

func (s *sSystemOperLog) GetPageListForSearch(ctx context.Context, req *model.PageListReq, in *req.SystemOperLogSearch) (rs []*res.SystemOperLog, total int, err error) {
	m := s.handleSearch(ctx, in)
	var entity []*entity.SystemOperLog
	err = orm.GetPageList(m, req).ScanAndCount(&entity, &total, false)
	if utils.IsError(err) {
		return nil, 0, err
	}
	rs = make([]*res.SystemOperLog, 0)
	if !g.IsEmpty(entity) {
		if err = gconv.Structs(entity, &rs); err != nil {
			return nil, 0, err
		}
	}
	return
}
