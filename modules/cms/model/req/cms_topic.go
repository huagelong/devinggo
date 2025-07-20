// Package req
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package req

type CmsTopicSearch struct {

  Id  int64 `json:"id" description:"ID" `
 
  CreatedAt  []string `json:"created_at" description:"创建时间" `
 
  Title  string `json:"title" description:"标题" `
 
  Status  int `json:"status" description:"状态 (1正常 2停用)" `
 
}

type CmsTopicSave struct {

  Title  string `json:"title"  v:"required"  description:"标题" `
 
  Content  string `json:"content"  v:"required"  description:"内容" `
 
  Status  int `json:"status"  v:"required"  description:"状态 (1正常 2停用)" `
 
  Remark  string `json:"remark"  v:"required"  description:"备注" `
 
}

type CmsTopicUpdate struct {

  Id  int64 `json:"id"  description:"ID" `
 
  Title  string `json:"title"  v:"required"  description:"标题" `
 
  Content  string `json:"content"  v:"required"  description:"内容" `
 
  Status  int `json:"status"  v:"required"  description:"状态 (1正常 2停用)" `
 
  Remark  string `json:"remark"  v:"required"  description:"备注" `
 
}
