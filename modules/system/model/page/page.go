// Package page
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package page

type ReqPageFunc interface {
	GetPage() int
	GetPageSize() int
}

const (
	DefaultPage     = 1
	DefaultPageSize = 15
)

// PageReq 分页请求
type PageReq struct {
	Page     int `json:"page" example:"1" d:"1" v:"min:1#页码最小值不能低于1"  dc:"当前页码"`
	PageSize int `json:"pageSize" example:"10" d:"15" dc:"每页数量"`
}

// PageRes 分页响应
type PageRes struct {
	PageInfo PageInfo `json:"pageInfo"`
}

// GetPage 获取当前页码
func (req *PageReq) GetPage() int {
	return req.Page
}

// GetPageSize 获取每页数量
func (req *PageReq) GetPageSize() int {
	return req.PageSize
}

type PageInfo struct {
	PageSize   int `json:"pageSize" example:"10" d:"10" v:"min:1|max:200#每页数量最小值不能低于1|最大值不能大于200" dc:"每页数量"`
	Page       int `json:"currentPage" example:"1" d:"1" v:"min:1#页码最小值不能低于1"  dc:"当前页码"`
	PageCount  int `json:"totalPage" example:"0"  dc:"分页个数"`
	TotalCount int `json:"total" example:"0" dc:"数据总行数"`
}

// Pack 打包分页数据
func (res *PageRes) Pack(req ReqPageFunc, totalCount int) {
	res.PageInfo.TotalCount = totalCount
	res.PageInfo.PageCount = CalPageCount(totalCount, req.GetPageSize())
	res.PageInfo.Page = req.GetPage()
	res.PageInfo.PageSize = req.GetPageSize()
}

func CalPageCount(totalCount int, pageSize int) int {
	return (totalCount + pageSize - 1) / pageSize
}

// CalPage 计算分页偏移量
func CalPage(page, pageSize int) (newPage, newPageSize int, offset int) {
	if page <= 0 {
		newPage = DefaultPage
	} else {
		newPage = page
	}
	if pageSize <= 0 {
		newPageSize = DefaultPageSize
	} else {
		newPageSize = pageSize
	}

	offset = (newPage - 1) * newPageSize
	return
}
