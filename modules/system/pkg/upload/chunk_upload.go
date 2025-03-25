// Package upload
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package upload

import (
	"context"
	"devinggo/modules/system/model/req"
	"devinggo/modules/system/model/res"
	"devinggo/modules/system/myerror"
	"devinggo/modules/system/pkg/utils"
	"devinggo/modules/system/service"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/grand"
	"io"
	"os"
	"strconv"
	"strings"
)

func ChunkUpload(ctx context.Context, in *req.ChunkUploadInput) (*res.SystemUploadFileRes, error) {
	runtimePath := GetRuntimePath()
	path := gstr.TrimRight(runtimePath, "/") + "/chunk"
	in.File.Filename = in.Hash + "_" + gconv.String(in.Index) + ".chunk"
	_, err := in.File.Save(path, false)
	if err != nil {
		return nil, err
	}

	// 分片上完，合并分片
	if in.Index == in.Total {
		fileName := in.Name
		if in.RandomName {
			fileName = strings.ToLower(strconv.FormatInt(gtime.TimestampNano(), 36) + grand.S(6))
			fileName = fileName + "." + in.Ext
		}
		contentType := in.Type
		resourceType := GetResourceType(contentType)
		dateDirName := gtime.Now().Format("Ymd")
		tmpPath := GetUploadFilePath(ctx, resourceType, dateDirName)
		newFileName := tmpPath + "/" + fileName
		err = combineChunks(in.Total, newFileName, path, in.Hash)
		if err != nil {
			return nil, err
		}
		originalName := in.Name
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
			Suffix:      in.Ext,
			SizeByte:    in.File.Size,
			SizeInfo:    formatSize(in.File.Size * 1024),
			LocalPath:   localPath,
			Url:         url,
		}, nil
	}
	return nil, nil
}

func combineChunks(totalChunks int64, outputFilePath string, tempDir string, hash string) error {
	// 创建最终文件
	outputFile, err := os.Create(outputFilePath)
	if err != nil {
		return err
	}
	for i := 1; i <= gconv.Int(totalChunks); i++ {
		chunkfileName := hash + "_" + gconv.String(i) + ".chunk"
		chunkFilePath := tempDir + "/" + chunkfileName
		// 打开分片文件
		chunkFile, err := os.Open(chunkFilePath)
		if err != nil {
			return err
		}
		// 从分片文件中读取数据并写入最终文件
		_, err = io.Copy(outputFile, chunkFile)
		if err != nil {
			err2 := chunkFile.Close()
			if err2 != nil {
				return err2
			}
			return err
		}
		err = chunkFile.Close()
		if err != nil {
			return err
		}
	}
	err = outputFile.Close()
	if err != nil {
		return err
	}
	_ = os.RemoveAll(tempDir)

	return nil
}

func CheckChunkFileMineType(ctx context.Context, ext string) (err error) {
	configMineType, err := service.SettingConfig().GetConfigByKey(ctx, "upload_allow_file", "upload_config")
	if err != nil {
		return
	}
	allowMineType := gstr.Split(configMineType, ",")
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
