// Package req
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package req

type SystemQueueMessageSearch struct {
	ReadStatus  string   `json:"read_status" dc:"read status"`
	ContentType string   `json:"content_type" dc:"content type"`
	Title       string   `json:"title" dc:"title"`
	CreatedAt   []string `json:"created_at" dc:"created at"`
}

type SystemQueueMessagesSend struct {
	Title   string  `json:"title" v:"required|length:2,50#请输入标题|标题长度为2~50位" dc:"标题"`
	Users   []int64 `json:"users" v:"required#请输入选择用户" dc:"用户"`
	Content string  `json:"content" v:"required#请输入内容" dc:"内容"`
}
