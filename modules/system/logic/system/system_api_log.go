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
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

type sSystemApiLog struct {
	base.BaseService
}

func init() {
	service.RegisterSystemApiLog(NewSystemApiLog())
}

func NewSystemApiLog() *sSystemApiLog {
	return &sSystemApiLog{}
}

func (s *sSystemApiLog) Model(ctx context.Context) *gdb.Model {
	return dao.SystemApiLog.Ctx(ctx).OnConflict("id")
}

func (s *sSystemApiLog) handleSearch(ctx context.Context, in *req.SystemApiLogSearch) (m *gdb.Model) {

	m = s.Model(ctx)

	if !g.IsEmpty(in.ApiName) {
		m = m.Where("api_name like ? ", "%"+in.ApiName+"%")
	}

	if !g.IsEmpty(in.Ip) {
		m = m.Where("ip like ? ", "%"+in.Ip+"%")
	}

	if !g.IsEmpty(in.AccessName) {
		m = m.Where("access_name like ? ", "%"+in.AccessName+"%")
	}
	if !g.IsEmpty(in.AccessTime) {
		if len(in.AccessTime) > 0 {
			m = m.WhereGTE("access_time", in.AccessTime[0]+" 00:00:00")
		}
		if len(in.AccessTime) > 1 {
			m = m.WhereLTE("access_time", in.AccessTime[1]+"23:59:59")
		}
	}
	return
}

func (s *sSystemApiLog) GetPageListForSearch(ctx context.Context, req *model.PageListReq, in *req.SystemApiLogSearch) (rs []*res.SystemApiLog, total int, err error) {
	m := s.handleSearch(ctx, in)
	var entity []*entity.SystemApiLog
	err = orm.GetPageList(m, req).ScanAndCount(&entity, &total, false)
	if utils.IsError(err) {
		return nil, 0, err
	}
	rs = make([]*res.SystemApiLog, 0)
	if !g.IsEmpty(entity) {
		if err = gconv.Structs(entity, &rs); err != nil {
			return nil, 0, err
		}
	}
	return
}

func (s *sSystemApiLog) Push(ctx context.Context) {
	r := request.GetHttpRequest(ctx)
	if r == nil {
		g.Log().Warningf(ctx, "ctx not http request")
		return
	}

	if r.GetError() == nil && r.Response.BufferLength() > 0 {
		return
	}

	permission := contexts.New().GetPermission(ctx)
	var entity *entity.SystemApi
	if !g.IsEmpty(permission) {
		err := service.SystemApi().Model(ctx).Where("access_name", permission).Scan(&entity)
		if utils.IsError(err) {
			return
		}
	}

	if g.IsEmpty(entity) {
		return
	}

	clientIp := location.GetClientIp(r)
	ipData, localErr := location.GetLocation(ctx, clientIp)
	if localErr != nil {
		g.Log().Debugf(ctx, "location.GetLocation clientIp:%v, err:%+v", clientIp, localErr)
		return
	}
	area := "本地"
	if !g.IsEmpty(ipData) {
		area = ipData.Area
	}

	res, bizCode := response.ResponseHandler(r)
	resJson := response.Json(r, bizCode, res)
	postData := contexts.New().GetRequestBody(ctx)

	systemApiLog := &do.SystemApiLog{
		ApiId:        entity.Id,         // 用户名
		ApiName:      entity.Name,       // 请求方式
		AccessName:   entity.AccessName, // 请求路由
		Ip:           clientIp,          // 请求IP地址
		IpLocation:   area,              // IP所属地
		RequestData:  postData,          // 请求数据
		ResponseCode: bizCode.Code(),    // 响应状态码
		ResponseData: resJson,           // 响应数据
		AccessTime:   gtime.Now(),       // 请求时间
		Remark:       "",                // 备注
	}
	s.Model(ctx).Data(systemApiLog).Insert()
}
