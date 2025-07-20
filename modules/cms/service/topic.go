// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"devinggo/modules/cms/model/req"
	"devinggo/modules/cms/model/res"
	"devinggo/modules/system/model"

	"github.com/gogf/gf/v2/database/gdb"
)

type (
	ICmsTopic interface {
		Model(ctx context.Context) *gdb.Model
		GetList(ctx context.Context, inReq *model.ListReq, in *req.CmsTopicSearch) (out []*res.CmsTopic, err error)
		GetPageList(ctx context.Context, req *model.PageListReq, in *req.CmsTopicSearch) (rs []*res.CmsTopic, total int, err error)
		Save(ctx context.Context, in *req.CmsTopicSave) (id uint64, err error)
		GetById(ctx context.Context, id uint64) (res *res.CmsTopic, err error)
		Update(ctx context.Context, in *req.CmsTopicUpdate) (err error)
		Delete(ctx context.Context, ids []uint64) (err error)
		RealDelete(ctx context.Context, ids []uint64) (err error)
		Recovery(ctx context.Context, ids []uint64) (err error)
		ChangeStatus(ctx context.Context, id uint64, status int) (err error)
		GetExportList(ctx context.Context, req *model.ListReq, in *req.CmsTopicSearch) (res []*res.CmsTopicExcel, err error)
	}
)

var (
	localCmsTopic ICmsTopic
)

func CmsTopic() ICmsTopic {
	if localCmsTopic == nil {
		panic("implement not found for interface ICmsTopic, forgot register?")
	}
	return localCmsTopic
}

func RegisterCmsTopic(i ICmsTopic) {
	localCmsTopic = i
}
