// Package system
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package system

import (
	"context"
	"devinggo/modules/system/api/system"
	"devinggo/modules/system/controller/base"
	"devinggo/modules/system/service"
	"github.com/gogf/gf/v2/util/gconv"
)

var (
	DashboardController = dashboardController{}
)

type dashboardController struct {
	base.BaseController
}

// GetStatistics 获取仪表板统计数据
func (c *dashboardController) GetStatistics(ctx context.Context, req *system.GetDashboardStatisticsReq) (rs *system.GetDashboardStatisticsRes, err error) {
	rs = &system.GetDashboardStatisticsRes{}
	
	statistics, err := service.Dashboard().GetStatistics(ctx)
	if err != nil {
		return nil, err
	}

	// 转换数据结构
	var dashboardStats system.DashboardStatistics
	err = gconv.Struct(statistics, &dashboardStats)
	if err != nil {
		return nil, err
	}

	rs.Data = dashboardStats
	return rs, nil
}

// GetLoginChart 获取仪表板登录图表数据
func (c *dashboardController) GetLoginChart(ctx context.Context, req *system.GetDashboardLoginChartReq) (rs *system.GetDashboardLoginChartRes, err error) {
	rs = &system.GetDashboardLoginChartRes{}
	
	days := req.Days
	if days <= 0 || days > 30 {
		days = 10
	}

	chartData, err := service.Dashboard().GetLoginChart(ctx, days)
	if err != nil {
		return nil, err
	}

	// 转换数据结构
	var loginChart system.DashboardLoginChart
	err = gconv.Struct(chartData, &loginChart)
	if err != nil {
		return nil, err
	}

	rs.Data = loginChart
	return rs, nil
}