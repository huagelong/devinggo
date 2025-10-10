// Package system
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package system

import (
	"devinggo/modules/system/model"
	"github.com/gogf/gf/v2/frame/g"
)

// GetDashboardStatisticsReq 获取仪表板统计数据请求
type GetDashboardStatisticsReq struct {
	g.Meta `path:"/dashboard/statistics" method:"get" tags:"仪表板" summary:"获取仪表板统计数据." x-exceptAuth:"true" x-permission:"system:dashboard:statistics"`
	model.AuthorHeader
}

// GetDashboardStatisticsRes 获取仪表板统计数据响应
type GetDashboardStatisticsRes struct {
	g.Meta `mime:"application/json"`
	Data   DashboardStatistics `json:"data" dc:"仪表板统计数据"`
}

// DashboardStatistics 仪表板统计数据结构
type DashboardStatistics struct {
	UserStats       StatItem `json:"userStats" dc:"用户统计"`
	AttachmentStats StatItem `json:"attachmentStats" dc:"附件统计"`
	LoginStats      StatItem `json:"loginStats" dc:"登录统计"`
	OperationStats  StatItem `json:"operationStats" dc:"操作统计"`
}

// StatItem 统计项数据结构
type StatItem struct {
	Total int64 `json:"total" dc:"总数"`
	New   int64 `json:"new" dc:"新增数"`
}

// GetDashboardLoginChartReq 获取仪表板登录图表数据请求
type GetDashboardLoginChartReq struct {
	g.Meta `path:"/dashboard/loginChart" method:"get" tags:"仪表板" summary:"获取仪表板登录图表数据." x-exceptAuth:"true" x-permission:"system:dashboard:loginChart"`
	model.AuthorHeader
	Days int `json:"days" dc:"天数" d:"10" v:"min:1|max:30"`
}

// GetDashboardLoginChartRes 获取仪表板登录图表数据响应
type GetDashboardLoginChartRes struct {
	g.Meta `mime:"application/json"`
	Data   DashboardLoginChart `json:"data" dc:"登录图表数据"`
}

// DashboardLoginChart 登录图表数据结构
type DashboardLoginChart struct {
	XAxis      []string `json:"xAxis" dc:"日期轴"`
	ChartsData []int64  `json:"chartsData" dc:"登录次数数据"`
}