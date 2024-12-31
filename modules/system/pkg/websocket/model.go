// Package websocket
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package websocket

import "github.com/gogf/gf/v2/frame/g"

// 当前输入对象
type Request struct {
	BindEvent string `json:"be"` //绑定的事件名称
	Event     string `json:"e"`  //事件名称
	Data      g.Map  `json:"d"`  //数据
	RequestId string `json:"r"`
}

// WResponse 输出对象
type WResponse struct {
	BindEvent string      `json:"be"` //绑定的事件名称
	Event     string      `json:"e"`  //事件名称
	Data      interface{} `json:"d"`  //数据
	RequestId string      `json:"r"`
	Code      int         `json:"c"`
	Message   string      `json:"m"`
	CallBack  int         `json:"cb"` //是否回调
}

type TopicWResponse struct {
	Topic     string     `json:"topic"`
	WResponse *WResponse `json:"wResponse"`
}

type ClientIdWResponse struct {
	ID        string     `json:"id"`
	WResponse *WResponse `json:"wResponse"`
}

type BroadcastWResponse struct {
	Broadcast string     `json:"broadcast"`
	WResponse *WResponse `json:"wResponse"`
}
