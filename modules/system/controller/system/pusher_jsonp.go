// Package system
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package system

import (
	"encoding/json"
	"regexp"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

var jsonpCallbackPattern = regexp.MustCompile(`^[a-zA-Z_$][0-9a-zA-Z_$\.]*$`)

func writeJSONOrJSONP(r *ghttp.Request, payload interface{}, status int) {
	r.Response.Status = status

	callback := r.Get("callback").String()
	if callback == "" {
		callback = r.Get("cb").String()
	}

	if callback == "" {
		r.Response.WriteJson(payload)
		r.ExitAll()
		return
	}

	if !jsonpCallbackPattern.MatchString(callback) {
		r.Response.Status = 400
		r.Response.WriteJson(g.Map{
			"error": "Invalid JSONP callback parameter",
		})
		r.ExitAll()
		return
	}

	jsonBytes, err := json.Marshal(payload)
	if err != nil {
		r.Response.Status = 500
		r.Response.WriteJson(g.Map{
			"error": "Failed to marshal response",
		})
		r.ExitAll()
		return
	}

	r.Response.Header().Set("Content-Type", "application/javascript; charset=utf-8")
	r.Response.Writef("%s(%s);", callback, string(jsonBytes))
	r.ExitAll()
}
