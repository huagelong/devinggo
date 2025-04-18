// Package upload
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package upload

import (
	"context"
	"devinggo/modules/system/model/res"
	"devinggo/modules/system/pkg/utils"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/grand"
	"mime"
	"net/http"
	"strconv"
	"strings"
)

func SaveNetworkImage(ctx context.Context, storageMode int, url string, randomName bool) (*res.SystemUploadFileRes, error) {

	runtimePath := GetRuntimePath()
	tmpPath := gstr.TrimRight(runtimePath, "/") + "/network"

	r, err := g.Client().Get(ctx, url)
	if err != nil {
		return nil, err
	}
	defer r.Close()

	originalName := gfile.Basename(url)
	ext := gfile.Ext(originalName)
	fileName := originalName
	if randomName {
		fileName = strings.ToLower(strconv.FormatInt(gtime.TimestampNano(), 36) + grand.S(6))
		fileName = fileName + ext
	}
	tmpFilePath := tmpPath + "/" + fileName
	err = gfile.PutBytes(tmpFilePath, r.ReadAll())
	if err != nil {
		return nil, err
	}
	ext, size, contentType, err := getImageInfo(tmpFilePath)
	if err != nil {
		return nil, err
	}

	resourceType := GetResourceType(contentType)
	dateDirName := gtime.Now().Format("Ymd")
	storagePath := gconv.String(resourceType) + "/" + dateDirName
	saveUrl := GetUploadUrlPath(ctx, resourceType, dateDirName, fileName)
	newPath := GetUploadFilePath(ctx, resourceType, dateDirName)
	newFileName := newPath + "/" + fileName
	if gfile.Exists(newFileName) {
		gfile.RemoveFile(tmpFilePath)
	} else {
		if err = gfile.Rename(tmpFilePath, newFileName); err != nil {
			return nil, err
		}
	}
	// 计算文件md5值
	md5, err := utils.FileMd5(newFileName)
	if err != nil {
		return nil, err
	}
	return &res.SystemUploadFileRes{
		StorageMode: storageMode,
		OriginName:  originalName,
		ObjectName:  fileName,
		Hash:        md5,
		MimeType:    resourceType,
		StoragePath: storagePath,
		Suffix:      ext,
		SizeByte:    size,
		SizeInfo:    formatSize(size * 1024),
		LocalPath:   newFileName,
		Url:         saveUrl,
	}, nil
}

func getImageInfo(filename string) (string, int64, string, error) {
	if !gfile.Exists(filename) {
		return "", 0, "", gerror.New("文件不存在")
	}
	dataByte := make([]byte, 4096)
	file, _ := gfile.Open(filename)
	defer file.Close()
	_, err := file.Read(dataByte)
	if err != nil {
		return "", 0, "", err
	}
	// 获取MIME类型
	mimeType := http.DetectContentType(dataByte)
	ext := mimeTypeToExtension(mimeType)
	// 获取文件扩展名
	size := gfile.Size(filename)
	return ext, size, mimeType, nil
}

// mimeTypeToExtension 根据MIME类型获取文件扩展名
func mimeTypeToExtension(mimeType string) string {
	exts, err := mime.ExtensionsByType(mimeType)
	if err != nil || len(exts) == 0 {
		return ""
	}
	return exts[0]
}
