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
	"devinggo/modules/system/service"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type sDashboard struct {
	base.BaseService
}

func init() {
	service.RegisterDashboard(NewDashboard())
}

func NewDashboard() *sDashboard {
	return &sDashboard{}
}

// GetStatistics 获取仪表板统计数据
func (s *sDashboard) GetStatistics(ctx context.Context) (statistics map[string]interface{}, err error) {
	// 获取当前时间和今天开始时间
	now := gtime.Now()
	todayStart := gtime.NewFromStr(now.Format("Y-m-d 00:00:00"))

	// 用户统计
	userTotal, err := dao.SystemUser.Ctx(ctx).Count()
	if err != nil {
		return nil, err
	}
	userNew, err := dao.SystemUser.Ctx(ctx).Where("created_at >= ?", todayStart).Count()
	if err != nil {
		return nil, err
	}

	// 附件统计
	attachmentTotal, err := dao.SystemUploadfile.Ctx(ctx).Count()
	if err != nil {
		return nil, err
	}
	attachmentNew, err := dao.SystemUploadfile.Ctx(ctx).Where("created_at >= ?", todayStart).Count()
	if err != nil {
		return nil, err
	}

	// 登录统计
	loginTotal, err := dao.SystemLoginLog.Ctx(ctx).Count()
	if err != nil {
		return nil, err
	}
	loginNew, err := dao.SystemLoginLog.Ctx(ctx).Where("login_time >= ?", todayStart).Count()
	if err != nil {
		return nil, err
	}

	// 操作统计
	operationTotal, err := dao.SystemOperLog.Ctx(ctx).Count()
	if err != nil {
		return nil, err
	}
	operationNew, err := dao.SystemOperLog.Ctx(ctx).Where("created_at >= ?", todayStart).Count()
	if err != nil {
		return nil, err
	}

	statistics = map[string]interface{}{
		"userStats": map[string]interface{}{
			"total": userTotal,
			"new":   userNew,
		},
		"attachmentStats": map[string]interface{}{
			"total": attachmentTotal,
			"new":   attachmentNew,
		},
		"loginStats": map[string]interface{}{
			"total": loginTotal,
			"new":   loginNew,
		},
		"operationStats": map[string]interface{}{
			"total": operationTotal,
			"new":   operationNew,
		},
	}

	return statistics, nil
}

// GetLoginChart 获取登录图表数据
func (s *sDashboard) GetLoginChart(ctx context.Context, days int) (chartData map[string]interface{}, err error) {
	if days <= 0 || days > 30 {
		days = 10
	}

	// 计算日期范围
	endDate := gtime.Now()
	startDate := endDate.AddDate(0, 0, -days+1)

	// 初始化日期数组和数据数组
	xAxis := make([]string, 0, days)
	chartsData := make([]int64, 0, days)

	// 生成日期范围
	for i := 0; i < days; i++ {
		date := startDate.AddDate(0, 0, i)
		dateStr := date.Format("Y-m-d")  // 使用 gtime 的格式字符串
		xAxis = append(xAxis, dateStr)
	}

	// 查询每天的登录次数
	for _, dateStr := range xAxis {
		dayStart := gtime.NewFromStr(dateStr + " 00:00:00")
		dayEnd := gtime.NewFromStr(dateStr + " 23:59:59")

		count, err := dao.SystemLoginLog.Ctx(ctx).
			Where("login_time >= ? AND login_time <= ?", dayStart, dayEnd).
			Count()
		if err != nil {
			g.Log().Error(ctx, "查询登录日志失败:", err)
			count = 0
		}

		chartsData = append(chartsData, int64(count))
	}

	chartData = map[string]interface{}{
		"xAxis":      xAxis,
		"chartsData": chartsData,
	}

	return chartData, nil
}