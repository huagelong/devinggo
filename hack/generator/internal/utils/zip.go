package utils

import (
	"archive/zip"
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/gogf/gf/v2/frame/g"
)

// ZipDirectory 压缩目录到zip文件
// srcDir: 源目录路径
// dstZip: 目标zip文件路径
func ZipDirectory(ctx context.Context, srcDir, dstZip string) error {
	// 创建zip文件
	zipFile, err := os.Create(dstZip)
	if err != nil {
		return fmt.Errorf("创建zip文件失败: %w", err)
	}
	defer zipFile.Close()

	// 创建zip.Writer
	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	// 遍历源目录
	err = filepath.Walk(srcDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 跳过根目录本身
		if path == srcDir {
			return nil
		}

		// 获取相对路径
		relPath, err := filepath.Rel(srcDir, path)
		if err != nil {
			return err
		}

		// 标准化路径分隔符（使用/）
		relPath = filepath.ToSlash(relPath)

		// 创建zip文件头
		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}
		header.Name = relPath
		header.Method = zip.Deflate

		// 如果是目录，确保名称以/结尾
		if info.IsDir() {
			header.Name += "/"
		}

		// 创建文件写入器
		writer, err := zipWriter.CreateHeader(header)
		if err != nil {
			return err
		}

		// 如果是文件，写入内容
		if !info.IsDir() {
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()

			_, err = io.Copy(writer, file)
			if err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return fmt.Errorf("压缩目录失败: %w", err)
	}

	g.Log().Infof(ctx, "成功压缩: %s -> %s", srcDir, dstZip)
	return nil
}

// UnzipFile 解压zip文件到目标目录
// srcZip: 源zip文件路径
// dstDir: 目标解压目录
func UnzipFile(srcZip, dstDir string) error {
	// 打开zip文件
	reader, err := zip.OpenReader(srcZip)
	if err != nil {
		return fmt.Errorf("打开zip文件失败: %w", err)
	}
	defer reader.Close()

	// 确保目标目录存在
	if err := os.MkdirAll(dstDir, 0755); err != nil {
		return fmt.Errorf("创建目标目录失败: %w", err)
	}

	// 遍历zip文件中的所有文件
	for _, file := range reader.File {
		// 安全检查：防止路径穿越攻击
		if strings.Contains(file.Name, "..") {
			return fmt.Errorf("检测到不安全的路径: %s", file.Name)
		}

		// 构建目标路径
		dstPath := filepath.Join(dstDir, file.Name)

		// 如果是目录
		if file.FileInfo().IsDir() {
			if err := os.MkdirAll(dstPath, file.Mode()); err != nil {
				return fmt.Errorf("创建目录失败: %w", err)
			}
			continue
		}

		// 确保父目录存在
		if err := os.MkdirAll(filepath.Dir(dstPath), 0755); err != nil {
			return fmt.Errorf("创建父目录失败: %w", err)
		}

		// 打开zip中的文件
		srcFile, err := file.Open()
		if err != nil {
			return fmt.Errorf("打开文件失败: %w", err)
		}

		// 创建目标文件
		dstFile, err := os.OpenFile(dstPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
		if err != nil {
			srcFile.Close()
			return fmt.Errorf("创建文件失败: %w", err)
		}

		// 复制内容
		_, err = io.Copy(dstFile, srcFile)
		srcFile.Close()
		dstFile.Close()

		if err != nil {
			return fmt.Errorf("写入文件失败: %w", err)
		}
	}

	return nil
}
