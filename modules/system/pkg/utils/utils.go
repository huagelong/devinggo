// Package utils
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package utils

import (
	"archive/zip"
	"context"
	"crypto/md5"
	"database/sql"
	"fmt"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/glog"
	"io"
	"log"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"

	"github.com/gogf/gf/v2/text/gstr"
)

func IsError(err error) bool {
	if err != nil && err != sql.ErrNoRows {
		return true
	} else {
		return false
	}
}

// 获取query参数
func GetQueryMap(rawQuery string) (map[string]interface{}, error) {
	queryMap, err := gstr.Parse(rawQuery)
	return queryMap, err
}

// 获取db名称
func GetConnectDbName(dsn string) (string, error) {
	// 正则表达式匹配 protocol(address) 部分
	re := regexp.MustCompile(`^(\w+):(.*)@(\w+)\(([^)]+)\)\/(\w+)`)
	matches := re.FindStringSubmatch(dsn)

	fmt.Printf("URL: %+v\n", matches)

	if len(matches) < 6 {
		return "", fmt.Errorf("invalid DSN format")
	}

	return matches[5], nil
}

func SafeGo(ctx context.Context, f func(ctx context.Context), lv ...int) {
	g.Go(ctx, f, func(ctx context.Context, err error) {
		var level = glog.LEVEL_ERRO
		if len(lv) > 0 {
			level = lv[0]
		}
		Logf(level, ctx, "SafeGo exec failed:%+v", err)
	})
}

func Logf(level int, ctx context.Context, format string, v ...interface{}) {
	switch level {
	case glog.LEVEL_DEBU:
		g.Log().Debugf(ctx, format, v...)
	case glog.LEVEL_INFO:
		g.Log().Infof(ctx, format, v...)
	case glog.LEVEL_NOTI:
		g.Log().Noticef(ctx, format, v...)
	case glog.LEVEL_WARN:
		g.Log().Warningf(ctx, format, v...)
	case glog.LEVEL_ERRO:
		g.Log().Errorf(ctx, format, v...)
	case glog.LEVEL_CRIT:
		g.Log().Criticalf(ctx, format, v...)
	case glog.LEVEL_PANI:
		g.Log().Panicf(ctx, format, v...)
	case glog.LEVEL_FATA:
		g.Log().Fatalf(ctx, format, v...)
	default:
		g.Log().Errorf(ctx, format, v...)
	}
}

// 获取module  system，api，websocket 等
func GetModule(path string) (module string) {
	slice := strings.Split(path, "/")
	if len(slice) < 2 {
		module = "system"
		return
	}

	if slice[1] == "" {
		module = "system"
		return
	}
	return slice[1]
}

// file md5
func FileMd5(filePath string) (string, error) {
	f, err := gfile.Open(filePath)
	if err != nil {
		err = gerror.Wrapf(err, `os.Open failed for "%s"`, filePath)
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

// 根目录
func GetRootPath() string {
	// 如果是go run则返回temp目录 go build 则返回当前目录
	//dir := getCurrentAbPathByExecutable()
	//tempDir := GetTmpDir()
	//
	//// 如果是临时目录执行 从Caller中获取
	//if strings.Contains(dir, tempDir) || tempDir == "." {
	//	dir = getCurrentAbPathByCaller()
	//}
	//g.Log().Info(context.Background(), "project dir: ", dir)
	//
	//return dir
	selfPath := gfile.SelfPath()
	tempDir := GetTmpDir()
	dir := gfile.SelfDir()
	// 判断是否是通过go run运行或位于临时目录
	if strings.Contains(selfPath, tempDir) || strings.Contains(selfPath, "go-build") {
		// 返回当前工作目录（项目根目录）
		dir = gfile.Pwd()
	}
	g.Log().Info(context.Background(), "project dir: ", dir)
	// 返回可执行文件所在目录
	return dir
}

// 获取系统临时目录，兼容go run
func GetTmpDir() string {
	dir := os.Getenv("TEMP")
	if dir == "" {
		dir = os.Getenv("TMP")
	}
	res, _ := filepath.EvalSymlinks(dir)
	return res
}

// 获取当前执行文件绝对路径
func getCurrentAbPathByExecutable() string {
	exePath, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	res, _ := filepath.EvalSymlinks(filepath.Dir(exePath))
	return res
}

// 获取当前执行文件绝对路径（go run）
func getCurrentAbPathByCaller() string {
	var abPath string
	_, filename, _, ok := runtime.Caller(0)
	if ok {
		abPath = path.Dir(path.Dir(filename))
	}
	return abPath
}

// 合并多个列表，并去重，使用自定义的比较函数
func MergeAndDeduplicateWithFunc[T any](compareFunc func(T) string, lists ...[]T) []T {
	var result []T
	seen := make(map[string]bool)
	for _, list := range lists {
		for _, item := range list {
			key := compareFunc(item)
			if _, ok := seen[key]; !ok {
				seen[key] = true
				result = append(result, item)
			}
		}
	}
	return result
}

// level 替换
func ReplaceSubstr(s, oldSubstr, newSubstr string) string {
	return strings.ReplaceAll(s, oldSubstr, newSubstr)
}

// ZipDirectory 压缩目录为 ZIP 文件
func ZipDirectory(ctx context.Context, source, target string) error {
	// 创建目标ZIP文件
	zipFile, err := os.Create(target)
	if err != nil {
		return err
	}
	defer zipFile.Close()

	// 创建ZIP写入器
	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	// 用于追踪已添加路径（处理大小写不敏感文件系统）
	seenPaths := make(map[string]struct{})
	basePath := filepath.Clean(source)

	// 递归遍历目录
	return filepath.Walk(basePath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 跳过根目录自身
		if path == basePath {
			return nil
		}

		// 获取相对路径
		relPath, err := filepath.Rel(basePath, path)
		if err != nil {
			return err
		}

		// 统一使用斜杠路径格式
		zipPath := filepath.ToSlash(relPath)

		// 处理空目录（需要单独创建目录条目）
		if info.IsDir() {
			zipPath += "/"
		}

		// 检查路径唯一性（考虑大小写不敏感系统）
		lowerPath := strings.ToLower(zipPath)
		if _, exists := seenPaths[lowerPath]; exists {
			return nil
		}
		seenPaths[lowerPath] = struct{}{}

		// 创建文件头
		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		// 设置压缩方法
		header.Method = zip.Deflate

		// 修正路径和名称
		header.Name = zipPath
		if info.IsDir() {
			header.Name += "/"
		}

		// 写入文件头
		writer, err := zipWriter.CreateHeader(header)
		if err != nil {
			return err
		}

		// 如果是目录，不需要写入内容
		if info.IsDir() {
			return nil
		}

		// 打开源文件
		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()

		// 复制文件内容到ZIP
		_, err = io.Copy(writer, file)
		return err
	})
}

func GetDbType() string {
	dbConfig := g.DB().GetConfig()
	link := dbConfig.Link
	dbType := "mysql" // 默认为MySQL
	if g.IsEmpty(link) {
		link = dbConfig.Type
	}
	// 判断数据库类型
	if strings.HasPrefix(link, "postgres:") || strings.HasPrefix(link, "postgresql:") || strings.HasPrefix(link, "pgsql:") {
		dbType = "postgres"
	}
	return dbType
}

// UnzipFile 解压ZIP文件到指定目录
func UnzipFile(zipPath string, destPath string) error {
	// 打开ZIP文件
	reader, err := zip.OpenReader(zipPath)
	if err != nil {
		return gerror.Wrapf(err, "打开ZIP文件失败: %s", zipPath)
	}
	defer reader.Close()

	// 确保目标目录存在
	if err := os.MkdirAll(destPath, 0755); err != nil {
		return gerror.Wrapf(err, "创建目标目录失败: %s", destPath)
	}

	// 用于存储已创建的目录
	createdDirs := make(map[string]bool)

	// 遍历ZIP文件中的所有文件和目录
	for _, file := range reader.File {
		// 构建目标路径
		path := filepath.Join(destPath, file.Name)

		// 检查路径穿越漏洞
		if !strings.HasPrefix(path, filepath.Clean(destPath)+string(os.PathSeparator)) {
			return gerror.Newf("非法的文件路径: %s", file.Name)
		}

		// 如果是目录
		if file.FileInfo().IsDir() {
			if err := os.MkdirAll(path, 0755); err != nil {
				return gerror.Wrapf(err, "创建目录失败: %s", path)
			}
			createdDirs[path] = true
			continue
		}

		// 确保父目录存在
		dir := filepath.Dir(path)
		if !createdDirs[dir] {
			if err := os.MkdirAll(dir, 0755); err != nil {
				return gerror.Wrapf(err, "创建父目录失败: %s", dir)
			}
			createdDirs[dir] = true
		}

		// 创建目标文件
		dstFile, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
		if err != nil {
			return gerror.Wrapf(err, "创建文件失败: %s", path)
		}

		// 打开ZIP中的源文件
		srcFile, err := file.Open()
		if err != nil {
			dstFile.Close()
			return gerror.Wrapf(err, "打开ZIP中的文件失败: %s", file.Name)
		}

		// 复制文件内容
		_, err = io.Copy(dstFile, srcFile)
		srcFile.Close()
		dstFile.Close()

		if err != nil {
			return gerror.Wrapf(err, "复制文件内容失败: %s", file.Name)
		}
	}

	return nil
}
