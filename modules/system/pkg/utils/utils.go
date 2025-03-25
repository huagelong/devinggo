// Package utils
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package utils

import (
	"context"
	"crypto/md5"
	"database/sql"
	"fmt"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/mholt/archives"
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
	dir := getCurrentAbPathByExecutable()
	tempDir := GetTmpDir()

	// 如果是临时目录执行 从Caller中获取
	if strings.Contains(dir, tempDir) || tempDir == "." {
		dir = getCurrentAbPathByCaller()
	}
	g.Log().Info(context.Background(), "project dir: ", dir)

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

	list, err := gfile.ScanDirFunc(source, "*", true, func(path string) string {
		return path
	})
	if err != nil {
		return err
	}
	if !g.IsEmpty(list) {
		zipPath := map[string]string{}
		for _, path := range list {
			relPath, err := filepath.Rel(source, path)
			if err != nil {
				return err
			}
			relPath = gstr.Replace(relPath, "\\", "/")
			zipPath[path] = relPath
		}
		files, err := archives.FilesFromDisk(ctx, &archives.FromDiskOptions{FollowSymlinks: false, ClearAttributes: false}, zipPath)
		if err != nil {
			return err
		}
		out, err := os.Create(target)
		if err != nil {
			return err
		}
		defer out.Close()
		format := archives.Zip{
			SelectiveCompression: true,
			ContinueOnError:      false,
		}
		err = format.Archive(ctx, out, files)
		if err != nil {
			return err
		}
	}

	return nil
}

func GetDbType() string {
	dbConfig := g.DB().GetConfig()
	link := dbConfig.Link
	dbType := "mysql" // 默认为MySQL

	// 判断数据库类型
	if strings.HasPrefix(link, "postgres:") || strings.HasPrefix(link, "postgresql:") || strings.HasPrefix(link, "pgsql:") {
		dbType = "postgres"
	}
	return dbType
}
