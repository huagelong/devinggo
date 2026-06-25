// Package upload
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package upload

import (
	"mime"
	"net/http"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gfile"
)

// getFileInfo 获取文件信息
func getFileInfo(filename string) (string, int64, string, error) {
	if !gfile.Exists(filename) {
		return "", 0, "", gerror.New("文件不存在")
	}

	dataByte := make([]byte, 4096)
	file, err := gfile.Open(filename)
	if err != nil {
		return "", 0, "", gerror.Wrap(err, "打开文件失败")
	}
	defer file.Close()

	if _, err := file.Read(dataByte); err != nil {
		return "", 0, "", gerror.Wrap(err, "读取文件失败")
	}

	// 获取MIME类型
	mimeType := http.DetectContentType(dataByte)
	ext := mimeTypeToExtension(mimeType)
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
