// Package glob
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package glob

import (
	"strings"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/glog"
)

// 支持的日志级别（参考 Pusher Server Library Reference Specification）
// https://pusher.com/docs/channels/library_auth_reference/logging
const (
	LogLevelDebug   = "debug"
	LogLevelInfo    = "info"
	LogLevelWarning = "warning"
	LogLevelError   = "error"
	LogLevelNone    = "none"
	LogLevelOff     = "off" // Pusher 标准的 OFF 级别
)

// WithWsLog 获取 WebSocket 日志器
// 支持通过 pusher.logLevel 配置项设置日志级别
//
// 配置示例：
//
//	pusher:
//	  logLevel: "info"  // debug, info, warning, error, none
func WithWsLog() *glog.Logger {
	ctx := gctx.GetInitCtx()
	logLevel := g.Cfg().MustGet(ctx, "pusher.logLevel", LogLevelInfo).String()
	return g.Log("ws").Level(parseLogLevel(logLevel))
}

// parseLogLevel 将 Pusher 风格的日志级别转换为 GoFrame 级别
// Pusher: debug, info, warning, error, none
// GoFrame: LEVEL_ALL, LEVEL_INFO, LEVEL_WARN, LEVEL_ERRO, LEVEL_NONE
func parseLogLevel(level string) int {
	switch strings.ToLower(strings.TrimSpace(level)) {
	case "debug":
		return glog.LEVEL_ALL
	case "info":
		return glog.LEVEL_INFO | glog.LEVEL_NOTI | glog.LEVEL_WARN | glog.LEVEL_ERRO | glog.LEVEL_CRIT
	case "warning", "warn":
		return glog.LEVEL_WARN | glog.LEVEL_ERRO | glog.LEVEL_CRIT
	case "error":
		return glog.LEVEL_ERRO | glog.LEVEL_CRIT
	case "none", "off":
		return glog.LEVEL_NONE
	default:
		// 默认 info 级别
		return glog.LEVEL_INFO | glog.LEVEL_NOTI | glog.LEVEL_WARN | glog.LEVEL_ERRO | glog.LEVEL_CRIT
	}
}
