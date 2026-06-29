// Package upload
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package upload

import (
	"context"
	"crypto/md5"
	"fmt"
	"io"
	"path"

	"devinggo/modules/system/model/req"
	"devinggo/modules/system/model/res"
	"devinggo/modules/system/pkg/utils"
	"devinggo/modules/system/pkg/utils/config"
	"devinggo/modules/system/service"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

// GetResourceType 根据MIME类型获取资源类型
func GetResourceType(mimeType string) string {
	mimeType = gstr.ToLower(mimeType)

	switch {
	case gstr.Contains(mimeType, "image"):
		return ResourceTypeImage
	case gstr.Contains(mimeType, "text"),
		gstr.Contains(mimeType, "pdf"),
		gstr.Contains(mimeType, "rtf"),
		gstr.Contains(mimeType, "vnd"),
		gstr.Contains(mimeType, "msword"):
		return ResourceTypeText
	case gstr.Contains(mimeType, "audio"):
		return ResourceTypeAudio
	case gstr.Contains(mimeType, "video"):
		return ResourceTypeVideo
	case gstr.Contains(mimeType, "zip"),
		gstr.Contains(mimeType, "7z"),
		gstr.Contains(mimeType, "tar"),
		gstr.Contains(mimeType, "rar"):
		return ResourceTypeZip
	default:
		return ResourceTypeOther
	}
}

func Upload(ctx context.Context, in *req.FileUploadInput) (*res.SystemUploadFileRes, error) {
	if !g.IsEmpty(in.Name) {
		in.File.Filename = in.Name
	}
	contentType := in.File.FileHeader.Header.Get("Content-Type")
	resourceType := GetResourceType(contentType)
	dateDirName := gtime.Now().Format(DateFormat)
	tmpPath, err := GetUploadFilePath(ctx, resourceType, dateDirName)
	if err != nil {
		return nil, err
	}
	fileName, err := in.File.Save(tmpPath, in.RandomName)
	if err != nil {
		return nil, gerror.Wrap(err, "保存文件失败")
	}
	originalName := in.File.Filename
	storagePath := gconv.String(resourceType) + "/" + dateDirName
	url := GetUploadUrlPath(ctx, resourceType, dateDirName, fileName)
	localPath := tmpPath + "/" + fileName
	// 计算文件md5值
	md5, err := CalcFileMd5(in.File)
	if err != nil {
		return nil, gerror.Wrap(err, "计算MD5失败")
	}

	// 上传到云存储
	if !IsLocalUpload(ctx) {
		fullPath := utils.GetRootPath() + "/" + localPath
		if err = PutFromFile(ctx, fullPath, url); err != nil {
			return nil, gerror.Wrap(err, "上传到云存储失败")
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
		SizeInfo:    FormatSize(in.File.Size * ByteSize),
		LocalPath:   localPath,
		Url:         url,
	}, nil
}

// Ext 获取文件后缀
func Ext(baseName string) string {
	return gstr.ToLower(gstr.StrEx(path.Ext(baseName), "."))
}

// FormatSize 格式化文件大小
func FormatSize(size int64) string {
	units := []string{"B", "KB", "MB", "GB", "TB", "PB"}
	floatSize := float64(size)

	for i := 0; i < len(units)-1; i++ {
		if floatSize < ByteSize {
			return fmt.Sprintf("%.2f %s", floatSize, units[i])
		}
		floatSize /= ByteSize
	}

	return fmt.Sprintf("%.2f %s", floatSize, units[len(units)-1])
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

// GetUploadFilePath 获取上传文件的完整路径
func GetUploadFilePath(ctx context.Context, resourceType, dateDirName string) (string, error) {
	uploadPath, err := GetUploadPath(ctx)
	if err != nil {
		return "", err
	}
	tmpPath := gstr.TrimRight(uploadPath, "/") + "/" + resourceType + "/" + dateDirName
	if err := ensureDir(tmpPath); err != nil {
		return "", err
	}
	return tmpPath, nil
}

// GetUploadUrlPath 获取文件访问URL路径
func GetUploadUrlPath(ctx context.Context, resourceType, dateDirName, fileName string) string {
	uploadPath := config.GetConfigString(ctx, "upload.dir", DefaultUpload)
	filePath := "/" + gstr.TrimRight(uploadPath, "/") + "/" + resourceType + "/" + dateDirName + "/" + fileName

	uploadDomain, err := service.SettingConfig().GetConfigByKey(ctx, ConfigKeyUploadDomain, ConfigGroupUpload)
	if err != nil || g.IsEmpty(uploadDomain) {
		return filePath
	}

	return gstr.TrimRight(uploadDomain, "/") + filePath
}

// GetUploadPath 获取上传基础路径
func GetUploadPath(ctx context.Context) (string, error) {
	uploadPath := config.GetConfigString(ctx, "upload.dir", DefaultUpload)
	tmpPath := PathResource + "/" + PathPublic + "/" + uploadPath
	if err := ensureDir(tmpPath); err != nil {
		return "", err
	}
	return tmpPath, nil
}

// GetRuntimePath 获取运行时路径
func GetRuntimePath() (string, error) {
	tmpPath := PathResource + "/" + PathRuntime
	if err := ensureDir(tmpPath); err != nil {
		return "", err
	}
	return tmpPath, nil
}
