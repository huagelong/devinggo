// Package upload
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package upload

import (
	"io"
	"os"
	"strconv"

	"github.com/gogf/gf/v2/errors/gerror"
)

// combineChunks 合并分片文件
func combineChunks(totalChunks int64, outputFilePath string, tempDir string, hash string) error {
	// 创建最终文件
	outputFile, err := os.Create(outputFilePath)
	if err != nil {
		return gerror.Wrap(err, "创建输出文件失败")
	}
	defer outputFile.Close()

	// 合并所有分片
	for i := 1; i <= int(totalChunks); i++ {
		chunkFileName := hash + "_" + strconv.Itoa(i) + ChunkExtension
		chunkFilePath := tempDir + "/" + chunkFileName

		if err := copyChunkToOutput(chunkFilePath, outputFile); err != nil {
			return err
		}
	}

	// 清理临时目录
	_ = os.RemoveAll(tempDir)
	return nil
}

// copyChunkToOutput 复制分片到输出文件
func copyChunkToOutput(chunkFilePath string, outputFile *os.File) error {
	chunkFile, err := os.Open(chunkFilePath)
	if err != nil {
		return gerror.Wrapf(err, "打开分片文件失败: %s", chunkFilePath)
	}
	defer chunkFile.Close()

	if _, err := io.Copy(outputFile, chunkFile); err != nil {
		return gerror.Wrap(err, "复制分片数据失败")
	}

	return nil
}
