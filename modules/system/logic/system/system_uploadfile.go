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
	"devinggo/modules/system/pkg/hook"
	"devinggo/modules/system/pkg/orm"
	"devinggo/modules/system/pkg/utils"
	"devinggo/modules/system/service"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

type sSystemUploadfile struct {
	base.BaseService
}

func init() {
	service.RegisterSystemUploadfile(NewSystemUploadfile())
}

func NewSystemUploadfile() *sSystemUploadfile {
	return &sSystemUploadfile{}
}

func (s *sSystemUploadfile) Model(ctx context.Context) *gdb.Model {
	return dao.SystemUploadfile.Ctx(ctx).Hook(hook.Bind()).Cache(orm.SetCacheOption(ctx)).OnConflict("id")
}

func (s *sSystemUploadfile) GetPageList(ctx context.Context, in *model.PageListReq, params *req.SystemUploadFileSearch) (out []*res.SystemUploadFile, total int, err error) {
	m := s.Model(ctx)

	if !g.IsEmpty(params.StorageMode) {
		m = m.Where("storage_mode", params.StorageMode)
	}

	if !g.IsEmpty(params.OriginName) {
		m = m.WhereLike("origin_name", "%"+params.OriginName+"%")
	}

	if !g.IsEmpty(params.MimeType) {
		m = m.Where("mime_type", params.MimeType)
	}

	if !g.IsEmpty(params.StoragePath) {
		m = m.WhereLike("storage_path", params.StoragePath+"%")
	}

	if !g.IsEmpty(params.MinDate) && !g.IsEmpty(params.MaxDate) {
		m = m.WhereBetween("created_at", params.MinDate+" 00:00:00", params.MaxDate+" 23:59:59")
	}

	err = orm.GetPageList(m, in).ScanAndCount(&out, &total, false)
	if utils.IsError(err) {
		return nil, 0, err
	}
	return
}

func (s *sSystemUploadfile) SaveDb(ctx context.Context, in *res.SystemUploadFileRes, createdBy int64) (rs int64, err error) {
	data := do.SystemUploadfile{
		StorageMode: in.StorageMode,
		OriginName:  in.OriginName,
		ObjectName:  in.ObjectName,
		Hash:        in.Hash,
		MimeType:    in.MimeType,
		StoragePath: in.StoragePath,
		Suffix:      in.Suffix,
		SizeByte:    in.SizeByte,
		SizeInfo:    in.SizeInfo,
		Url:         in.Url,
		CreatedBy:   createdBy,
	}
	result, err := s.Model(ctx).Data(data).Insert()
	if utils.IsError(err) {
		return
	}
	rsTmp, err := result.LastInsertId()
	if err != nil {
		return
	}
	rs = gconv.Int64(rsTmp)
	return
}

func (s *sSystemUploadfile) GetByHash(ctx context.Context, hash string) (rs *res.SystemUploadFileRes, err error) {
	var fileInfo *entity.SystemUploadfile
	err = s.Model(ctx).Where("hash", hash).Scan(&fileInfo)
	if utils.IsError(err) {
		return
	}
	if g.IsEmpty(fileInfo) {
		s.Model(ctx).Unscoped().Where("hash", hash).Delete()
		return
	}

	if err := gconv.Struct(fileInfo, &rs); err != nil {
		return nil, err
	}
	return
}

func (s *sSystemUploadfile) GetFileInfoById(ctx context.Context, id int64) (rs *res.SystemUploadFile, err error) {
	var fileInfo *entity.SystemUploadfile
	err = s.Model(ctx).Where("id", id).Scan(&fileInfo)
	if utils.IsError(err) {
		return
	}
	if err := gconv.Struct(fileInfo, &rs); err != nil {
		return nil, err
	}
	return
}

func (s *sSystemUploadfile) GetFileInfoByHash(ctx context.Context, hash string) (rs *res.SystemUploadFile, err error) {
	var fileInfo *entity.SystemUploadfile
	err = s.Model(ctx).Where("hash", hash).Scan(&fileInfo)
	if utils.IsError(err) {
		return
	}
	if g.IsEmpty(fileInfo) {
		s.Model(ctx).Unscoped().Where("hash", hash).Delete()
		return
	}

	if err := gconv.Struct(fileInfo, &rs); err != nil {
		return nil, err
	}
	return
}

func (s *sSystemUploadfile) Delete(ctx context.Context, ids []int64) (err error) {
	_, err = s.Model(ctx).WhereIn("id", ids).Delete()
	if utils.IsError(err) {
		return err
	}
	return
}

func (s *sSystemUploadfile) RealDelete(ctx context.Context, ids []int64) (err error) {
	_, err = s.Model(ctx).Unscoped().WhereIn("id", ids).Delete()
	if utils.IsError(err) {
		return
	}
	//todo 真实删除 网盘 处理
	return
}

func (s *sSystemUploadfile) Recovery(ctx context.Context, ids []int64) (err error) {
	_, err = s.Model(ctx).Unscoped().WhereIn("id", ids).Update(g.Map{"deleted_at": nil})
	if utils.IsError(err) {
		return err
	}
	return
}
