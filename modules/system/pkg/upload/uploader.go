// Package upload
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package upload

import (
	"context"
	"path/filepath"
	"strconv"
	"strings"

	"devinggo/modules/system/model/res"
	"devinggo/modules/system/pkg/utils"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/grand"
)

// Uploader 文件上传器
type Uploader struct {
	ctx          context.Context
	storageMode  int
	randomName   bool
	validateType bool
	customName   string
}

// NewUploader 创建上传器
func NewUploader(ctx context.Context) *Uploader {
	return &Uploader{
		ctx:          ctx,
		storageMode:  StorageModeLocal,
		randomName:   true,
		validateType: true,
	}
}

// SetStorageMode 设置存储模式
func (u *Uploader) SetStorageMode(mode int) *Uploader {
	u.storageMode = mode
	return u
}

// SetRandomName 设置是否使用随机文件名
func (u *Uploader) SetRandomName(random bool) *Uploader {
	u.randomName = random
	return u
}

// SetCustomName 设置自定义文件名
func (u *Uploader) SetCustomName(name string) *Uploader {
	u.customName = name
	return u
}

// SetValidateType 设置是否验证文件类型
func (u *Uploader) SetValidateType(validate bool) *Uploader {
	u.validateType = validate
	return u
}

// UseCloudStorage 使用云存储
func (u *Uploader) UseCloudStorage() *Uploader {
	u.storageMode = StorageModeCloud
	return u
}

// UseLocalStorage 使用本地存储
func (u *Uploader) UseLocalStorage() *Uploader {
	u.storageMode = StorageModeLocal
	return u
}

// UploadFile 上传文件
func (u *Uploader) UploadFile(file *ghttp.UploadFile) (*res.SystemUploadFileRes, error) {
	if file == nil {
		return nil, gerror.New("文件不能为空")
	}

	// 验证文件类型
	if u.validateType {
		if err := NewValidator(u.ctx).ValidateFile(file); err != nil {
			return nil, err
		}
	}

	// 设置文件名
	if !g.IsEmpty(u.customName) {
		file.Filename = u.customName
	}

	// 获取文件信息
	contentType := file.FileHeader.Header.Get("Content-Type")
	resourceType := GetResourceType(contentType)
	dateDirName := gtime.Now().Format(DateFormat)

	// 保存文件
	uploadPath, err := GetUploadFilePath(u.ctx, resourceType, dateDirName)
	if err != nil {
		return nil, err
	}
	fileName, err := file.Save(uploadPath, u.randomName)
	if err != nil {
		return nil, gerror.Wrap(err, "保存文件失败")
	}

	// 构建文件信息
	return u.buildFileResult(file, fileName, resourceType, dateDirName, uploadPath)
}

// UploadImage 上传图片
func (u *Uploader) UploadImage(file *ghttp.UploadFile) (*res.SystemUploadFileRes, error) {
	if file == nil {
		return nil, gerror.New("图片文件不能为空")
	}

	// 验证图片类型
	if u.validateType {
		if err := NewValidator(u.ctx).ValidateImage(file); err != nil {
			return nil, err
		}
	}

	return u.UploadFile(file)
}

// SaveFromURL 从URL保存文件
func (u *Uploader) SaveFromURL(url string) (result *res.SystemUploadFileRes, err error) {
	if g.IsEmpty(url) {
		return nil, gerror.New("URL不能为空")
	}

	// 下载文件到临时目录
	runtimePath, err := GetRuntimePath()
	if err != nil {
		return nil, err
	}
	tmpPath := filepath.Join(runtimePath, PathNetwork)

	if dirErr := ensureDir(tmpPath); dirErr != nil {
		return nil, dirErr
	}

	// 下载文件
	r, getErr := g.Client().Get(u.ctx, url)
	if getErr != nil {
		return nil, gerror.Wrap(getErr, "下载文件失败")
	}
	defer r.Close()

	// 生成文件名
	originalName := gfile.Basename(url)
	ext := gfile.Ext(originalName)
	fileName := u.generateFileName(originalName, ext)

	// 保存临时文件
	tmpFilePath := filepath.Join(tmpPath, fileName)
	if putErr := gfile.PutBytes(tmpFilePath, r.ReadAll()); putErr != nil {
		return nil, gerror.Wrap(putErr, "保存临时文件失败")
	}

	// 获取文件信息
	fileExt, size, contentType, infoErr := getFileInfo(tmpFilePath)
	if infoErr != nil {
		return nil, infoErr
	}

	// 移动到最终目录
	resourceType := GetResourceType(contentType)
	dateDirName := gtime.Now().Format(DateFormat)
	finalPath, err := GetUploadFilePath(u.ctx, resourceType, dateDirName)
	if err != nil {
		return nil, err
	}
	finalFilePath := filepath.Join(finalPath, fileName)

	// 如果文件已存在，删除临时文件
	if gfile.Exists(finalFilePath) {
		gfile.RemoveFile(tmpFilePath)
	} else {
		if mvErr := gfile.Rename(tmpFilePath, finalFilePath); mvErr != nil {
			return nil, gerror.Wrap(mvErr, "移动文件失败")
		}
	}

	// 计算MD5
	md5Hash, err := utils.FileMd5(finalFilePath)
	if err != nil {
		return nil, gerror.Wrap(err, "计算MD5失败")
	}

	// 上传到云存储
	if err := u.uploadToCloud(finalFilePath, resourceType, dateDirName, fileName); err != nil {
		return nil, err
	}

	// 构建结果
	return &res.SystemUploadFileRes{
		StorageMode: u.storageMode,
		OriginName:  originalName,
		ObjectName:  fileName,
		Hash:        md5Hash,
		MimeType:    resourceType,
		StoragePath: filepath.Join(resourceType, dateDirName),
		Suffix:      fileExt,
		SizeByte:    size,
		SizeInfo:    FormatSize(size * ByteSize),
		LocalPath:   finalFilePath,
		Url:         GetUploadUrlPath(u.ctx, resourceType, dateDirName, fileName),
	}, nil
}

// buildFileResult 构建文件上传结果
func (u *Uploader) buildFileResult(file *ghttp.UploadFile, fileName, resourceType, dateDirName, uploadPath string) (*res.SystemUploadFileRes, error) {
	originalName := file.Filename
	storagePath := filepath.Join(resourceType, dateDirName)
	url := GetUploadUrlPath(u.ctx, resourceType, dateDirName, fileName)
	localPath := filepath.Join(uploadPath, fileName)

	// 计算文件MD5值
	md5Hash, err := CalcFileMd5(file)
	if err != nil {
		return nil, gerror.Wrap(err, "计算MD5失败")
	}

	// 上传到云存储
	if err := u.uploadToCloud(localPath, resourceType, dateDirName, fileName); err != nil {
		return nil, err
	}

	return &res.SystemUploadFileRes{
		StorageMode: u.storageMode,
		OriginName:  originalName,
		ObjectName:  fileName,
		Hash:        md5Hash,
		MimeType:    resourceType,
		StoragePath: storagePath,
		Suffix:      Ext(fileName),
		SizeByte:    file.Size,
		SizeInfo:    FormatSize(file.Size * ByteSize),
		LocalPath:   localPath,
		Url:         url,
	}, nil
}

// uploadToCloud 上传到云存储
func (u *Uploader) uploadToCloud(localPath, resourceType, dateDirName, fileName string) error {
	if !IsLocalUpload(u.ctx) {
		url := GetUploadUrlPath(u.ctx, resourceType, dateDirName, fileName)
		fullPath := filepath.Join(utils.GetRootPath(), localPath)
		if err := PutFromFile(u.ctx, fullPath, url); err != nil {
			return gerror.Wrap(err, "上传到云存储失败")
		}
	}
	return nil
}

// UploadChunk 上传文件分片
func (u *Uploader) UploadChunk(file *ghttp.UploadFile, index, total int64, hash, ext, fileType, fileName string) (*res.SystemUploadFileRes, error) {
	if file == nil {
		return nil, gerror.New("分片文件不能为空")
	}

	// 验证文件扩展名
	if u.validateType {
		if err := NewValidator(u.ctx).ValidateExtension(ext, true); err != nil {
			return nil, err
		}
	}

	// 保存分片
	runtimePath, err := GetRuntimePath()
	if err != nil {
		return nil, err
	}
	chunkPath := filepath.Join(runtimePath, PathChunk)
	if err := ensureDir(chunkPath); err != nil {
		return nil, err
	}

	// 设置分片文件名
	file.Filename = hash + "_" + strconv.FormatInt(index, 10) + ChunkExtension
	if _, err := file.Save(chunkPath, false); err != nil {
		return nil, gerror.Wrap(err, "保存分片文件失败")
	}

	// 如果是最后一个分片，合并所有分片
	if index == total {
		return u.mergeAndUploadChunks(hash, total, ext, fileType, fileName, chunkPath)
	}

	return nil, nil
}

// mergeAndUploadChunks 合并并上传分片
func (u *Uploader) mergeAndUploadChunks(hash string, total int64, ext, fileType, originalName, chunkPath string) (result *res.SystemUploadFileRes, err error) {
	// 生成最终文件名
	finalFileName := originalName
	if u.randomName {
		finalFileName = strings.ToLower(strconv.FormatInt(gtime.TimestampNano(), 36) + grand.S(6))
		finalFileName = finalFileName + "." + ext
	}

	// 获取资源类型和路径
	resourceType := GetResourceType(fileType)
	dateDirName := gtime.Now().Format(DateFormat)
	uploadPath, upErr := GetUploadFilePath(u.ctx, resourceType, dateDirName)
	if upErr != nil {
		return nil, upErr
	}
	finalFilePath := filepath.Join(uploadPath, finalFileName)

	// 合并分片
	if combErr := combineChunks(total, finalFilePath, chunkPath, hash); combErr != nil {
		return nil, gerror.Wrap(combErr, "合并分片文件失败")
	}

	// 计算MD5
	md5Hash, err := utils.FileMd5(finalFilePath)
	if err != nil {
		return nil, gerror.Wrap(err, "计算MD5失败")
	}

	// 获取文件大小
	fileSize := gfile.Size(finalFilePath)

	// 上传到云存储
	if cloudErr := u.uploadToCloud(finalFilePath, resourceType, dateDirName, finalFileName); cloudErr != nil {
		return nil, cloudErr
	}

	// 构建结果
	return &res.SystemUploadFileRes{
		StorageMode: u.storageMode,
		OriginName:  originalName,
		ObjectName:  finalFileName,
		Hash:        md5Hash,
		MimeType:    resourceType,
		StoragePath: filepath.Join(resourceType, dateDirName),
		Suffix:      ext,
		SizeByte:    fileSize,
		SizeInfo:    FormatSize(fileSize * ByteSize),
		LocalPath:   finalFilePath,
		Url:         GetUploadUrlPath(u.ctx, resourceType, dateDirName, finalFileName),
	}, nil
}

// generateFileName 生成文件名
func (u *Uploader) generateFileName(originalName, ext string) string {
	if !u.randomName {
		return originalName
	}
	return strings.ToLower(strconv.FormatInt(gtime.TimestampNano(), 36)+grand.S(6)) + ext
}

// ensureDir 确保目录存在
func ensureDir(dirPath string) error {
	if !gfile.Exists(dirPath) {
		if err := gfile.Mkdir(dirPath); err != nil {
			return gerror.Wrapf(err, "创建目录失败: %s", dirPath)
		}
	}
	return nil
}
