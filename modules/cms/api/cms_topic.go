// Package topic
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package api

import (
    "devinggo/modules/system/model"
    "devinggo/modules/system/model/page"
    "devinggo/modules/cms/model/req"
    "devinggo/modules/cms/model/res"
    "github.com/gogf/gf/v2/frame/g"

    "github.com/gogf/gf/v2/net/ghttp"

)

type IndexCmsTopicReq struct {
    g.Meta `path:"/cmsTopic/index" method:"get" tags:"主题管理" summary:"分页列表" x-permission:"cms:cmsTopic:index" `
    model.AuthorHeader
    model.PageListReq
    req.CmsTopicSearch
}

type IndexCmsTopicRes struct {
    g.Meta `mime:"application/json"`
    page.PageRes
    Items []res.CmsTopic `json:"items"  dc:"list" `
}

type ListCmsTopicReq struct {
    g.Meta `path:"/cmsTopic/list" method:"get" tags:"主题管理" summary:"列表" x-permission:"cms:cmsTopic:list" `
    model.AuthorHeader
    model.ListReq
    req.CmsTopicSearch
}

type ListCmsTopicRes struct {
    g.Meta `mime:"application/json"`
    Data   []res.CmsTopic `json:"data"  dc:"list" `
}

type SaveCmsTopicReq struct {
    g.Meta `path:"/cmsTopic/save" method:"post" tags:"主题管理" summary:"新增" x-permission:"cms:cmsTopic:save"`
    model.AuthorHeader
    req.CmsTopicSave
}

type SaveCmsTopicRes struct {
    g.Meta `mime:"application/json"`
    Id     uint64 `json:"id" dc:"id"`
}

type ReadCmsTopicReq struct {
    g.Meta `path:"/cmsTopic/read/{Id}" method:"get" tags:"主题管理" summary:"获取单个信息" x-permission:"cms:cmsTopic:read"`
    model.AuthorHeader
    Id uint64 `json:"id" dc:"主题管理 id" v:"required|min:1#Id不能为空"`
}

type ReadCmsTopicRes struct {
    g.Meta `mime:"application/json"`
    Data   res.CmsTopic `json:"data" dc:"信息数据"`
}

type UpdateCmsTopicReq struct {
    g.Meta `path:"/cmsTopic/update/{Id}" method:"put" tags:"主题管理" summary:"更新" x-permission:"cms:cmsTopic:update"`
    model.AuthorHeader
    req.CmsTopicUpdate
}

type UpdateCmsTopicRes struct {
    g.Meta `mime:"application/json"`
}

type DeleteCmsTopicReq struct {
    g.Meta `path:"/cmsTopic/delete" method:"delete" tags:"主题管理" summary:"删除" x-permission:"cms:cmsTopic:delete"`
    model.AuthorHeader
    Ids []uint64 `json:"ids" dc:"ids" v:"min-length:1#Id不能为空"`
}

type DeleteCmsTopicRes struct {
    g.Meta `mime:"application/json"`
}

type RecycleCmsTopicReq struct {
    g.Meta `path:"/cmsTopic/recycle" method:"get" tags:"主题管理" summary:"回收站列表" x-permission:"cms:cmsTopic:recycle" `
    model.AuthorHeader
    model.PageListReq
    req.CmsTopicSearch
}

type RecycleCmsTopicRes struct {
    g.Meta `mime:"application/json"`
    page.PageRes
    Items []res.CmsTopic `json:"items"  dc:"list" `
}

type RealDeleteCmsTopicReq struct {
    g.Meta `path:"/cmsTopic/realDelete" method:"delete" tags:"主题管理" summary:"单个或批量真实删除 （清空回收站）" x-permission:"cms:cmsTopic:realDelete"`
    model.AuthorHeader
    Ids []uint64 `json:"ids" dc:"ids" v:"min-length:1#Id不能为空"`
}

type RealDeleteCmsTopicRes struct {
    g.Meta `mime:"application/json"`
}

type RecoveryCmsTopicReq struct {
        g.Meta `path:"/cmsTopic/recovery" method:"put" tags:"主题管理" summary:"单个或批量恢复在回收站的" x-permission:"cms:cmsTopic:recovery"`
        model.AuthorHeader
        Ids []uint64 `json:"ids" dc:"ids" v:"min-length:1#Id不能为空"`
}

type RecoveryCmsTopicRes struct {
        g.Meta `mime:"application/json"`
}

type ExportCmsTopicReq struct {
    g.Meta `path:"/cmsTopic/export" method:"post" tags:"主题管理" summary:"导出" x-permission:"cms:cmsTopic:export"`
    model.AuthorHeader
    model.ListReq
    req.CmsTopicSearch
}

type ExportCmsTopicRes struct {
    g.Meta `mime:"application/json"`
}

type ImportCmsTopicReq struct {
    g.Meta `path:"/cmsTopic/import" method:"post" mime:"multipart/form-data" tags:"主题管理" summary:"导入" x-permission:"cms:cmsTopic:import"`
    model.AuthorHeader
    File *ghttp.UploadFile `json:"file" type:"file"  dc:"pls upload file"`
}

type ImportCmsTopicRes struct {
    g.Meta `mime:"application/json"`
}

type DownloadTemplateCmsTopicReq struct {
    g.Meta `path:"/cmsTopic/downloadTemplate" method:"post,get" tags:"主题管理" summary:"下载导入模板." x-exceptAuth:"true" x-permission:"cms:cmsTopic:downloadTemplate"`
    model.AuthorHeader
}

type DownloadTemplateCmsTopicRes struct {
    g.Meta `mime:"application/json"`
}

type ChangeStatusCmsTopicReq struct {
    g.Meta `path:"/cmsTopic/changeStatus" method:"put" tags:"主题管理" summary:"更改状态" x-permission:"cms:cmsTopic:changeStatus"`
    model.AuthorHeader
    Id     uint64 `json:"id" dc:"ids" v:"min:1#Id不能为空"`
    Status int    `json:"status" dc:"status" v:"min:1#状态不能为空"`
}

type ChangeStatusCmsTopicRes struct {
    g.Meta `mime:"application/json"`
}

type RemoteCmsTopicReq struct {
    g.Meta `path:"/cmsTopic/remote" method:"post" tags:"主题管理" summary:"远程万能通用列表接口" x-exceptAuth:"true" x-permission:"cms:cmsTopic:remote"`
    model.AuthorHeader
    model.PageListReq
}

type RemoteCmsTopicRes struct {
    g.Meta `mime:"application/json"`
    page.PageRes
    Items []res.CmsTopic `json:"items"  dc:"list" `
    Data  []res.CmsTopic `json:"data"  dc:"list" `
}
