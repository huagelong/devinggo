// Package response
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package response

import (
	"devinggo/modules/system/codes"
	"devinggo/modules/system/pkg/contexts"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"
	"net/http"
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

func JsonError(r *ghttp.Request, code gcode.Code, text ...string) {
	var textData string
	if len(text) > 0 {
		textData = text[0]
	}
	r.SetError(gerror.NewCode(code, textData))
}

// Redirect 重定向
func Redirect(r *ghttp.Request, location string, code ...int) {
	r.Response.RedirectTo(location, code...)
}

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
	takeUpTime := contexts.New().GetTakeUpTime(r.GetCtx())

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

func JsonExit(r *ghttp.Request, code gcode.Code, data ...interface{}) {
	responseData := interface{}(nil)
	if len(data) > 0 {
		responseData = data[0]
	}
	jsonData := Json(r, code, responseData)
	r.Response.WriteJsonExit(jsonData)
}

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
