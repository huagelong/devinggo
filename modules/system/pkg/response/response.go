// Package response
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package response

import (
	"net/http"

	"devinggo/modules/system/codes"
	"devinggo/modules/system/pkg/contexts"
	"devinggo/modules/system/pkg/i18n"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"
)

type Response struct {
	RequestId  string      `json:"requestId"`
	Path       string      `json:"path"`
	Success    bool        `json:"success"`
	Message    string      `json:"message"`
	Code       int         `json:"code"`
	Data       interface{} `json:"data"`
	TakeUpTime int64       `json:"takeUpTime"`
}

// ============================== 底层方法 ==============================
// 以下方法为底层实现，建议使用下方的便捷方法（Success, Fail, Ok, Error 等）

// Redirect 重定向
func Redirect(r *ghttp.Request, location string, code ...int) {
	r.Response.RedirectTo(location, code...)
}

// Json 构建响应对象（不直接输出）
// 注意：通常情况下推荐使用便捷方法（Success, Fail 等），此方法用于特殊场景如日志记录
func Json(r *ghttp.Request, bizCode gcode.Code, responseData interface{}) (jsonData Response) {
	var (
		msg string
	)
	ctx := r.GetCtx()
	bizCode = codes.NewCode(ctx, bizCode)
	msg = bizCode.Message()
	// 清空响应
	r.Response.ClearBuffer()
	if r.Response.Status != http.StatusNotFound &&
		r.Response.Status != http.StatusUnauthorized &&
		r.Response.Status != http.StatusForbidden &&
		r.Response.Status != http.StatusBadRequest &&
		r.Response.Status != http.StatusInternalServerError {
		r.Response.WriteHeader(http.StatusOK)
	}
	path := r.Request.URL.Path
	success := false
	if bizCode.Code() == 0 {
		success = true
	}

	if g.IsNil(responseData) {
		responseData = g.Map{}
	}

	// 请求耗时
	takeUpTime := contexts.GetTakeUpTime(r.GetCtx())

	jsonData = Response{
		Code:       bizCode.Code(),
		Message:    msg,
		Success:    success,
		Path:       path,
		RequestId:  gctx.CtxId(r.Context()),
		Data:       responseData,
		TakeUpTime: takeUpTime,
	}
	return
}

// JsonExit 输出 JSON 响应并退出
// 注意：通常情况下推荐使用便捷方法（Success, Fail, Ok, Error 等），此方法用于需要自定义 code 的场景
func JsonExit(r *ghttp.Request, code gcode.Code, data ...interface{}) {
	responseData := interface{}(nil)
	if len(data) > 0 {
		responseData = data[0]
	}
	jsonData := Json(r, code, responseData)
	r.Response.WriteJsonExit(jsonData)
}

// ResponseHandler 响应处理器（用于中间件）
func ResponseHandler(r *ghttp.Request) (res interface{}, bizCode gcode.Code) {
	//ctx := r.Context()
	var (
		err = r.GetError()
	)
	res = r.GetHandlerResponse()
	if err != nil {
		defaultErr := err.Error()
		//g.Log().Debug(ctx, "responseHandler err:", defaultErr)
		bizCode = gerror.Code(err)
		res = g.Map{}
		if !g.IsEmpty(defaultErr) {
			bizCode = gcode.New(bizCode.Code(), defaultErr, nil)
		}
	} else {
		if r.Response.Status == http.StatusOK { //200
			bizCode = gcode.CodeOK
		} else if r.Response.Status == http.StatusNotFound { //404
			bizCode = gcode.CodeNotFound
		} else if r.Response.Status == http.StatusUnauthorized { //401
			bizCode = codes.CodeNotLogged
		} else if r.Response.Status == http.StatusForbidden { //403
			bizCode = codes.CodeForbidden
		} else if r.Response.Status == http.StatusBadRequest { //400
			bizCode = gcode.CodeInvalidRequest
		} else { //500
			bizCode = gcode.CodeInternalError
		}
	}
	return
}

// ============================== 便捷方法 ==============================

// Success 成功响应，返回数据
func Success(r *ghttp.Request, data ...interface{}) {
	responseData := interface{}(nil)
	if len(data) > 0 {
		responseData = data[0]
	}
	JsonExit(r, gcode.CodeOK, responseData)
}

// SuccessWithMessage 成功响应，自定义消息
func SuccessWithMessage(r *ghttp.Request, message string, data ...interface{}) {
	responseData := interface{}(nil)
	if len(data) > 0 {
		responseData = data[0]
	}
	code := gcode.New(gcode.CodeOK.Code(), message, nil)
	JsonExit(r, code, responseData)
}

// Ok Success 的别名，更符合 HTTP 语义
func Ok(r *ghttp.Request, data ...interface{}) {
	Success(r, data...)
}

// OkWithMessage SuccessWithMessage 的别名
func OkWithMessage(r *ghttp.Request, message string, data ...interface{}) {
	SuccessWithMessage(r, message, data...)
}

// Fail 失败响应，自定义错误消息
func Fail(r *ghttp.Request, message string) {
	code := gcode.New(gcode.CodeInternalError.Code(), message, nil)
	JsonExit(r, code, nil)
}

// FailWithCode 失败响应，自定义错误码和消息
func FailWithCode(r *ghttp.Request, code int, message string) {
	bizCode := gcode.New(code, message, nil)
	JsonExit(r, bizCode, nil)
}

// FailWithData 失败响应，带数据
func FailWithData(r *ghttp.Request, message string, data interface{}) {
	code := gcode.New(gcode.CodeInternalError.Code(), message, nil)
	JsonExit(r, code, data)
}

// Error Fail 的别名，更直观的错误响应
func Error(r *ghttp.Request, message string) {
	Fail(r, message)
}

// ErrorWithCode FailWithCode 的别名
func ErrorWithCode(r *ghttp.Request, code int, message string) {
	FailWithCode(r, code, message)
}

// BadRequest 参数错误响应（400）
func BadRequest(r *ghttp.Request, message ...string) {
	ctx := r.GetCtx()
	msg := i18n.T(ctx, "请求参数错误")
	if len(message) > 0 {
		msg = message[0]
	}
	code := gcode.New(gcode.CodeInvalidParameter.Code(), msg, nil)
	JsonExit(r, code, nil)
}

// Unauthorized 未授权响应（401）
func Unauthorized(r *ghttp.Request, message ...string) {
	ctx := r.GetCtx()
	msg := i18n.T(ctx, "未授权")
	if len(message) > 0 {
		msg = message[0]
	}
	code := gcode.New(codes.CodeNotLogged.Code(), msg, nil)
	JsonExit(r, code, nil)
}

// Forbidden 无权限响应（403）
func Forbidden(r *ghttp.Request, message ...string) {
	ctx := r.GetCtx()
	msg := i18n.T(ctx, "无权限访问")
	if len(message) > 0 {
		msg = message[0]
	}
	code := gcode.New(codes.CodeForbidden.Code(), msg, nil)
	JsonExit(r, code, nil)
}

// NotFound 资源未找到响应（404）
func NotFound(r *ghttp.Request, message ...string) {
	ctx := r.GetCtx()
	msg := i18n.T(ctx, "资源未找到")
	if len(message) > 0 {
		msg = message[0]
	}
	code := gcode.New(gcode.CodeNotFound.Code(), msg, nil)
	JsonExit(r, code, nil)
}
