// Package topic
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package controller

import (
"devinggo/modules/system/controller/base"
"devinggo/modules/cms/api"
"devinggo/modules/cms/model/req"
"devinggo/modules/cms/model/res"
"devinggo/modules/system/pkg/orm"
"devinggo/modules/system/pkg/utils/request"
"devinggo/modules/cms/service"
"context"
"github.com/gogf/gf/v2/container/gmap"
"github.com/gogf/gf/v2/frame/g"

"devinggo/modules/system/pkg/excel"

"devinggo/internal/dao"
"github.com/gogf/gf/v2/database/gdb"
"devinggo/internal/model/do"
"devinggo/modules/system/pkg/utils"
"devinggo/modules/system/myerror"
"github.com/gogf/gf/v2/util/gconv"

)

var (
    CmsTopicController = cmsTopicController{}
)

type cmsTopicController struct {
    base.BaseController
}

func (c *cmsTopicController) Index(ctx context.Context, in *api.IndexCmsTopicReq) (out *api.IndexCmsTopicRes, err error) {
    out = &api.IndexCmsTopicRes{}
    items, totalCount, err := service.CmsTopic().GetPageList(ctx, &in.PageListReq, &in.CmsTopicSearch)
    if err != nil {
        return
    }

    if !g.IsEmpty(items) {
        for _, item := range items {
            out.Items = append(out.Items, *item)
        }
    } else {
        out.Items = make([]res.CmsTopic, 0)
    }
   out.PageRes.Pack(in, totalCount)
   return
}

func (c *cmsTopicController) Recycle(ctx context.Context, in *api.RecycleCmsTopicReq) (out *api.RecycleCmsTopicRes, err error) {
    out = &api.RecycleCmsTopicRes{}
    pageListReq := &in.PageListReq
    pageListReq.Recycle = true
    items, totalCount, err := service.CmsTopic().GetPageList(ctx, pageListReq, &in.CmsTopicSearch)
    if err != nil {
        return
    }

    if !g.IsEmpty(items) {
        for _, item := range items {
            out.Items = append(out.Items, *item)
        }
    } else {
        out.Items = make([]res.CmsTopic, 0)
    }
  out.PageRes.Pack(in, totalCount)
    return
}

func (c *cmsTopicController) List(ctx context.Context, in *api.ListCmsTopicReq) (out *api.ListCmsTopicRes, err error) {
    out = &api.ListCmsTopicRes{}
    rs, err := service.CmsTopic().GetList(ctx, &in.ListReq, &req.CmsTopicSearch{})
    if err != nil {
        return
    }

    if !g.IsEmpty(rs) {
        for _, v := range rs {
            out.Data = append(out.Data, *v)
        }
    } else {
        out.Data = make([]res.CmsTopic, 0)
    }
    return
}

func (c *cmsTopicController) Save(ctx context.Context, in *api.SaveCmsTopicReq) (out *api.SaveCmsTopicRes, err error) {
    out = &api.SaveCmsTopicRes{}
    id, err := service.CmsTopic().Save(ctx, &in.CmsTopicSave)
    if err != nil {
        return
    }
    out.Id = id
   return
}

func (c *cmsTopicController) Read(ctx context.Context, in *api.ReadCmsTopicReq) (out *api.ReadCmsTopicRes, err error) {
    out = &api.ReadCmsTopicRes{}
    info, err := service.CmsTopic().GetById(ctx, in.Id)
    if err != nil {
        return
    }
    out.Data = *info
    return
}

func (c *cmsTopicController) Update(ctx context.Context, in *api.UpdateCmsTopicReq) (out *api.UpdateCmsTopicRes, err error) {
    out = &api.UpdateCmsTopicRes{}
    err = service.CmsTopic().Update(ctx, &in.CmsTopicUpdate)
    if err != nil {
        return
    }
    return
}

func (c *cmsTopicController) Delete(ctx context.Context, in *api.DeleteCmsTopicReq) (out *api.DeleteCmsTopicRes, err error) {
    out = &api.DeleteCmsTopicRes{}
    err = service.CmsTopic().Delete(ctx, in.Ids)
    if err != nil {
        return
    }
    return
}
func (c *cmsTopicController) RealDelete(ctx context.Context, in *api.RealDeleteCmsTopicReq) (out *api.RealDeleteCmsTopicRes, err error) {
    out = &api.RealDeleteCmsTopicRes{}
    err = service.CmsTopic().RealDelete(ctx, in.Ids)
    if err != nil {
        return
    }
    return
}

func (c *cmsTopicController) Recovery(ctx context.Context, in *api.RecoveryCmsTopicReq) (out *api.RecoveryCmsTopicRes, err error) {
    out = &api.RecoveryCmsTopicRes{}
    err = service.CmsTopic().Recovery(ctx, in.Ids)
    if err != nil {
        return
    }
    return
}

func (c *cmsTopicController) ChangeStatus(ctx context.Context, in *api.ChangeStatusCmsTopicReq) (out *api.ChangeStatusCmsTopicRes, err error) {
    out = &api.ChangeStatusCmsTopicRes{}
    err = service.CmsTopic().ChangeStatus(ctx, in.Id, in.Status)
    if err != nil {
        return
    }
    return
}

func (c *cmsTopicController) Export(ctx context.Context, in *api.ExportCmsTopicReq) (out *api.ExportCmsTopicRes, err error) {
    var (
        fileName  = "cmsTopic"
        sheetName = "Sheet1"
    )
    exports, err := service.CmsTopic().GetExportList(ctx, &in.ListReq, &in.CmsTopicSearch)
    if err != nil {
        return
    }
    //创建导出对象
    export := excel.NewExcelExport(sheetName, res.CmsTopicExcel{})
    //销毁对象
    defer export.Close()
    newExports := []res.CmsTopicExcel{}
    if !g.IsEmpty(exports) {
        for _, item := range exports {
            newExports = append(newExports, *item)
        }
    }
    err = export.ExportSmallExcelByStruct(newExports).Download(ctx, fileName).Error()
    if err != nil {
        return
    }
    return
}

func (c *cmsTopicController) Import(ctx context.Context, in *api.ImportCmsTopicReq) (out *api.ImportCmsTopicRes, err error) {
    tmpPath := utils.GetTmpDir()
    fileName, err := in.File.Save(tmpPath, true)
    if err != nil {
        return nil, err
    }
    localPath := tmpPath + "/" + fileName
    var result []res.CmsTopicExcel
    importFile := excel.NewExcelImportFile(localPath, res.CmsTopicExcel{})
    defer importFile.Close()

    err = importFile.ImportDataToStruct(&result).Error()
    if err != nil {
        return nil, err
    } else {
        if !g.IsEmpty(result) {
            err = dao.CmsTopic.Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
                for _, item := range result {
                    var saveData *do.CmsTopic
                    if err = gconv.Struct(item, &saveData); err != nil {
                        return
                    }
                    _, err = service.CmsTopic().Model(ctx).OmitEmptyData().Data(saveData).Save()
                    if err != nil {
                        return err
                    }
                }
                return
            })
            if err != nil {
                return
            }
        } else {
            err = myerror.ValidationFailed(ctx, "没有数据!")
        }
    }
    return
}

func (c *cmsTopicController) DownloadTemplate(ctx context.Context, in *api.DownloadTemplateCmsTopicReq) (out *api.DownloadTemplateCmsTopicRes, err error) {
    var (
        fileName  = "cmsTopic_template"
        sheetName = "Sheet1"
        exports   = make([]res.CmsTopicExcel, 0)
    )
    export := excel.NewExcelExport(sheetName, res.CmsTopicExcel{})
    defer export.Close()
    err = export.ExportSmallExcelByStruct(exports).Download(ctx, fileName).Error()
    if err != nil {
        return
    }
    return
}

func (c *cmsTopicController) Remote(ctx context.Context, in *api.RemoteCmsTopicReq) (out *api.RemoteCmsTopicRes, err error) {
    out = &api.RemoteCmsTopicRes{}
    r := request.GetHttpRequest(ctx)
    params := gmap.NewStrAnyMapFrom(r.GetMap())
    m := service.CmsTopic().Model(ctx)
    var rs res.CmsTopic
    remote := orm.NewRemote(m, rs)
    openPage := params.GetVar("openPage")
    items, totalCount, err := remote.GetRemote(ctx, params)
    if err != nil {
        return
    }
    if !g.IsEmpty(openPage) && openPage.Bool() {
        if !g.IsEmpty(items) {
            for _, item := range items {
                out.Items = append(out.Items, item)
            }
        } else {
            out.Items = make([]res.CmsTopic, 0)
        }
        out.PageRes.Pack(in, totalCount)
    } else {
        if !g.IsEmpty(items) {
            out.Data = items
        } else {
            out.Data = make([]res.CmsTopic, 0)
        }
    }
    return
}