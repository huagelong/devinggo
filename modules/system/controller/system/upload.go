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
	"devinggo/modules/system/model/req"
	upload2 "devinggo/modules/system/pkg/upload"
	"devinggo/modules/system/pkg/utils"
	"devinggo/modules/system/pkg/utils/request"
	"devinggo/modules/system/service"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

var (
	UploadController = uploadController{}
)

type uploadController struct {
	base.BaseController
}

func (c *uploadController) UploadFile(ctx context.Context, in *system.UploadFileReq) (out *system.UploadFileRes, err error) {
	//扩展验证
	err = upload2.CheckFileMineType(ctx, in.File)
	if err != nil {
		return
	}
	// hash验证 计算文件md5值
	md5Hash, err := upload2.CalcFileMd5(in.File)
	if err != nil {
		return
	}

	fileInfo, err := service.SystemUploadfile().GetByHash(ctx, md5Hash)
	if err != nil {
		return
	}
	out = &system.UploadFileRes{}
	if !g.IsEmpty(fileInfo) {
		out.Data = *fileInfo
		return
	}
	//上传文件
	storageMode, err := service.SettingConfig().GetConfigByKey(ctx, "upload_mode", "upload_config")
	if err != nil {
		return nil, err
	}

	storageModeInt := gconv.Int(storageMode)
	inUpload := &req.FileUploadInput{
		File:        in.File,
		RandomName:  true,
		StorageMode: storageModeInt,
	}
	upload, err := upload2.Upload(ctx, inUpload)
	out.Data = *upload

	_, err = service.SystemUploadfile().SaveDb(ctx, upload, c.UserId)
	if err != nil {
		return nil, err
	}
	return
}

func (c *uploadController) UploadImage(ctx context.Context, in *system.UploadImageReq) (out *system.UploadImageRes, err error) {

	//扩展验证
	err = upload2.CheckImageMineType(ctx, in.File)
	if err != nil {
		return
	}
	// hash验证 计算文件md5值
	md5Hash, err := upload2.CalcFileMd5(in.File)
	if err != nil {
		return
	}

	fileInfo, err := service.SystemUploadfile().GetByHash(ctx, md5Hash)
	if err != nil {
		return
	}
	out = &system.UploadImageRes{}
	if !g.IsEmpty(fileInfo) {
		out.Data = *fileInfo
		return
	}
	//上传文件
	storageMode, err := service.SettingConfig().GetConfigByKey(ctx, "upload_mode", "upload_config")
	if err != nil {
		return nil, err
	}

	storageModeInt := gconv.Int(storageMode)
	inUpload := &req.FileUploadInput{
		File:        in.File,
		RandomName:  true,
		StorageMode: storageModeInt,
	}
	upload, err := upload2.Upload(ctx, inUpload)
	if err != nil {
		return nil, err
	}

	if g.IsEmpty(upload) {
		return nil, err
	}

	out.Data = *upload

	_, err = service.SystemUploadfile().SaveDb(ctx, upload, c.UserId)
	if err != nil {
		return nil, err
	}
	return
}

func (c *uploadController) ChunkUpload(ctx context.Context, in *system.ChunkUploadReq) (out *system.ChunkUploadRes, err error) {
	//扩展验证
	err = upload2.CheckChunkFileMineType(ctx, in.Ext)
	if err != nil {
		return
	}
	// hash验证 计算文件md5值
	md5Hash, err := upload2.CalcFileMd5(in.File)
	if err != nil {
		return
	}

	fileInfo, err := service.SystemUploadfile().GetByHash(ctx, md5Hash)
	if err != nil {
		return
	}
	out = &system.ChunkUploadRes{}
	if !g.IsEmpty(fileInfo) {
		out.Data = *fileInfo
		return
	}
	//上传文件
	storageMode, err := service.SettingConfig().GetConfigByKey(ctx, "upload_mode", "upload_config")
	if err != nil {
		return nil, err
	}

	storageModeInt := gconv.Int(storageMode)

	inUpload := &req.ChunkUploadInput{
		File:        in.File,
		Total:       in.Total,
		Index:       in.Index,
		Hash:        in.Hash,
		Ext:         in.Ext,
		Type:        in.Type,
		Name:        in.Name,
		Size:        in.Size,
		StorageMode: storageModeInt,
		RandomName:  true,
	}
	upload, err := upload2.ChunkUpload(ctx, inUpload)
	if err != nil {
		return nil, err
	}
	if g.IsNil(upload) {
		//return ['chunk' => $data['index'], 'code' => 201, 'status' => 'success'];
		out.Data.Chunk = in.Index
	} else {
		out.Data = *upload

		_, err = service.SystemUploadfile().SaveDb(ctx, upload, c.UserId)
		if err != nil {
			return nil, err
		}

	}
	return
}

func (c *uploadController) SaveNetworkImage(ctx context.Context, in *system.SaveNetworkImageReq) (out *system.SaveNetworkImageRes, err error) {
	//url验证
	if !(gstr.HasPrefix(in.Url, "http") || gstr.HasPrefix(in.Url, "https")) {
		return nil, gerror.New("图片地址请以 http 或 https 开头")
	}

	out = &system.SaveNetworkImageRes{}
	storageMode, err := service.SettingConfig().GetConfigByKey(ctx, "upload_mode", "upload_config")
	if err != nil {
		return nil, err
	}

	storageModeInt := gconv.Int(storageMode)
	upload, err := upload2.SaveNetworkImage(ctx, storageModeInt, in.Url, true)
	out.Data = *upload

	_, err = service.SystemUploadfile().SaveDb(ctx, upload, c.UserId)
	if err != nil {
		return nil, err
	}
	return
}

func (c *uploadController) GetFileInfoById(ctx context.Context, in *system.GetFileInfoByIdReq) (out *system.GetFileInfoByIdRes, err error) {
	out = &system.GetFileInfoByIdRes{}
	fileInfo, err := service.SystemUploadfile().GetFileInfoById(ctx, in.Id)
	if err != nil {
		return nil, err
	}
	if g.IsEmpty(fileInfo) {
		return 
	}
	out.Data = *fileInfo
	return
}

func (c *uploadController) GetFileInfoByHash(ctx context.Context, in *system.GetFileInfoByHashReq) (out *system.GetFileInfoByHashRes, err error) {
	out = &system.GetFileInfoByHashRes{}
	fileInfo, err := service.SystemUploadfile().GetFileInfoByHash(ctx, in.Hash)
	if err != nil {
		return nil, err
	}
	if g.IsEmpty(fileInfo) {
		return 
	}
	out.Data = *fileInfo
	return
}

func (c *uploadController) DownloadById(ctx context.Context, in *system.DownloadByIdReq) (out *system.DownloadByIdRes, err error) {
	fileInfo, err := service.SystemUploadfile().GetFileInfoById(ctx, in.Id)
	if err != nil {
		return nil, err
	}
	if g.IsEmpty(fileInfo) {
		return nil, gerror.New("文件不存在")
	}
	r := request.GetHttpRequest(ctx)
	r.Response.ServeFileDownload(utils.GetRootPath()+"/resource/public/"+fileInfo.Url, fileInfo.OriginName)
	return
}

func (c *uploadController) DownloadByHash(ctx context.Context, in *system.DownloadByHashReq) (out *system.DownloadByHashRes, err error) {
	fileInfo, err := service.SystemUploadfile().GetFileInfoByHash(ctx, in.Hash)
	if err != nil {
		return nil, err
	}
	if g.IsEmpty(fileInfo) {
		return nil, gerror.New("文件不存在")
	}
	r := request.GetHttpRequest(ctx)
	r.Response.ServeFileDownload(utils.GetRootPath()+"/resource/public/"+fileInfo.Url, fileInfo.OriginName)
	return
}

func (c *uploadController) ShowFile(ctx context.Context, in *system.ShowFileReq) (out *system.ShowFileRes, err error) {
	fileInfo, err := service.SystemUploadfile().GetFileInfoByHash(ctx, in.Hash)
	if err != nil {
		return nil, err
	}
	if g.IsEmpty(fileInfo) {
		return nil, gerror.New("文件不存在")
	}
	r := request.GetHttpRequest(ctx)
	r.Response.ServeFile(utils.GetRootPath() + "/resource/public/" + fileInfo.Url)
	return
}
