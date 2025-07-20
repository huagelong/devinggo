// Package req
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package res

import "github.com/gogf/gf/v2/os/gtime"

type CmsTopic struct {

 Id  int64 `json:"id"  description:"ID" `
 
 CreatedAt  *gtime.Time `json:"created_at"  description:"创建时间" `
 
 Title  string `json:"title"  description:"标题" `
 
 Content  string `json:"content"  description:"内容" `
 
 Status  int `json:"status"  description:"状态 (1正常 2停用)" `
 
}

type CmsTopicExcel struct {

    Title  string `json:"title"  v:"required"  description:"标题"  excelName:"标题" excelIndex:"4"  `

    Content  string `json:"content"  v:"required"  description:"内容"  excelName:"内容" excelIndex:"5"  `

    Status  int `json:"status"  v:"required"  description:"状态 (1正常 2停用)"  excelName:"状态 (1正常 2停用)" excelIndex:"6"  `

    Remark  string `json:"remark"  v:"required"  description:"备注"  excelName:"备注" excelIndex:"7"  `

}
