// Package upload
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package upload

import (
	"devinggo/modules/system/model/req"
	"devinggo/modules/system/model/res"
	"devinggo/modules/system/myerror"
	"devinggo/modules/system/pkg/utils"
	"devinggo/modules/system/pkg/utils/config"
	"devinggo/modules/system/service"
	"context"
	"crypto/md5"
	"fmt"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"io"
	"path"
)

func GetResourceType(mimeType string) string {
	mimeType = gstr.ToLower(mimeType)
	if gstr.Contains(mimeType, "image") {
		return "image"
	} else if gstr.Contains(mimeType, "text") ||
		gstr.Contains(mimeType, "pdf") ||
		gstr.Contains(mimeType, "rtf") ||
		gstr.Contains(mimeType, "vnd") ||
		gstr.Contains(mimeType, "msword") {
		return "text"
	} else if gstr.Contains(mimeType, "audio") {
		return "audio"
	} else if gstr.Contains(mimeType, "video") {
		return "video"
	} else if gstr.Contains(mimeType, "zip") ||
		gstr.Contains(mimeType, "7z") ||
		gstr.Contains(mimeType, "tar") ||
		gstr.Contains(mimeType, "rar") {
		return "zip"
	} else {
		return "other"
	}
}

func Upload(ctx context.Context, in *req.FileUploadInput) (*res.SystemUploadFileRes, error) {
	if !g.IsEmpty(in.Name) {
		in.File.Filename = in.Name
	}
	contentType := in.File.FileHeader.Header.Get("Content-Type")
	resourceType := GetResourceType(contentType)
	dateDirName := gtime.Now().Format("Ymd")
	tmpPath := GetUploadFilePath(ctx, resourceType, dateDirName)
	fileName, err := in.File.Save(tmpPath, in.RandomName)
	if err != nil {
		return nil, err
	}
	originalName := in.File.Filename
	storagePath := gconv.String(resourceType) + "/" + dateDirName
	url := GetUploadUrlPath(ctx, resourceType, dateDirName, fileName)
	localPath := tmpPath + "/" + fileName
	// 计算文件md5值
	md5, err := CalcFileMd5(in.File)
	if err != nil {
		return nil, err
	}

	if !IsLocalUpload(ctx) {
		err = PutFromFile(ctx, utils.GetRootPath()+"/"+localPath, url)
		if err != nil {
			return nil, err
		}
	}

	return &res.SystemUploadFileRes{
		StorageMode: in.StorageMode,
		OriginName:  originalName,
		ObjectName:  fileName,
		Hash:        md5,
		MimeType:    resourceType,
		StoragePath: storagePath,
		Suffix:      Ext(fileName),
		SizeByte:    in.File.Size,
		SizeInfo:    formatSize(in.File.Size * 1024),
		LocalPath:   localPath,
		Url:         url,
	}, nil
}

// Ext 获取文件后缀
func Ext(baseName string) string {
	return gstr.ToLower(gstr.StrEx(path.Ext(baseName), "."))
}

func formatSize(size int64) string {
	units := []string{"B", "KB", "MB", "GB", "TB", "PB"}
	index := 0
	for i := 0; size >= 1024 && i < 5; i++ {
		size /= 1024
		index = i
	}
	return fmt.Sprintf("%.2f %s", float64(size), units[index])
}

// CalcFileMd5 计算文件md5值
func CalcFileMd5(file *ghttp.UploadFile) (string, error) {
	f, err := file.Open()
	if err != nil {
		err = gerror.Wrapf(err, `os.Open failed for name "%s"`, file.Filename)
		return "", err
	}
	defer f.Close()
	h := md5.New()
	_, err = io.Copy(h, f)
	if err != nil {
		err = gerror.Wrap(err, `io.Copy failed`)
		return "", err
	}
	return fmt.Sprintf("%x", h.Sum(nil)), nil
}

// UploadFileByte 获取上传文件的byte
func UploadFileByte(file *ghttp.UploadFile) ([]byte, error) {
	open, err := file.Open()
	if err != nil {
		return nil, err
	}
	return io.ReadAll(open)
}

func CheckFileMineType(ctx context.Context, file *ghttp.UploadFile) (err error) {
	configMineType, err := service.SettingConfig().GetConfigByKey(ctx, "upload_allow_file", "upload_config")
	if err != nil {
		return
	}
	allowMineType := gstr.Split(configMineType, ",")
	ext := Ext(file.Filename)
	if !gstr.InArray(allowMineType, ext) {
		uploadAllowImage, err := service.SettingConfig().GetConfigByKey(ctx, "upload_allow_image", "upload_config")
		if err != nil {
			return err
		}
		allowImageMineType := gstr.Split(uploadAllowImage, ",")
		if !gstr.InArray(allowImageMineType, ext) {
			return myerror.ValidationFailed(ctx, "不允许上传此类型文件")
		}
	}
	return
}

func CheckImageMineType(ctx context.Context, file *ghttp.UploadFile) (err error) {
	configMineType, err := service.SettingConfig().GetConfigByKey(ctx, "upload_allow_image", "upload_config")
	if err != nil {
		return
	}
	allowMineType := gstr.Split(configMineType, ",")
	ext := Ext(file.Filename)
	if !gstr.InArray(allowMineType, ext) {
		return myerror.ValidationFailed(ctx, "不允许上传此类型图片")
	}
	return
}

func GetUploadFilePath(ctx context.Context, resourceType, dateDirName string) string {
	uploadPath := GetUploadPath(ctx)
	tmpPath := gstr.TrimRight(uploadPath, "/") + "/" + resourceType + "/" + dateDirName
	if !gfile.Exists(tmpPath) {
		if err := gfile.Mkdir(tmpPath); err != nil {
			panic(err)
		}
	}
	return tmpPath
}

func GetUploadUrlPath(ctx context.Context, resourceType, dateDirName, fileName string) string {
	uploadPath := config.GetConfigString(ctx, "upload.dir", "uploads")
	tmpPath := "/" + gstr.TrimRight(uploadPath, "/") + "/" + resourceType + "/" + dateDirName + "/" + fileName
	return tmpPath
}

func GetUploadPath(ctx context.Context) string {
	uploadPath := config.GetConfigString(ctx, "upload.dir", "uploads")
	tmpPath := getResourcePath() + "/public/" + uploadPath
	if !gfile.Exists(tmpPath) {
		if err := gfile.Mkdir(tmpPath); err != nil {
			panic(err)
		}
	}
	return tmpPath
}

func getResourcePath() string {
	return "resource"
}

func GetRuntimePath() string {
	tmpPath := getResourcePath() + "/runtime"
	if !gfile.Exists(tmpPath) {
		if err := gfile.Mkdir(tmpPath); err != nil {
			panic(err)
		}
	}
	return tmpPath
}
