// Package topic
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package topic

import (
    "devinggo/internal/dao"
    "devinggo/modules/system/logic/base"
    "devinggo/modules/system/model"
    "devinggo/internal/model/do"
    "devinggo/internal/model/entity"
    "devinggo/modules/cms/model/req"
    "devinggo/modules/cms/model/res"
    "devinggo/modules/system/pkg/handler"
    "devinggo/modules/system/pkg/hook"
    "devinggo/modules/system/pkg/orm"
    "devinggo/modules/system/pkg/utils"
    "devinggo/modules/cms/service"
    "context"
    "github.com/gogf/gf/v2/database/gdb"
    "github.com/gogf/gf/v2/frame/g"
    "github.com/gogf/gf/v2/util/gconv"
)

type sCmsTopic struct {
    base.BaseService
}

func init() {
    service.RegisterCmsTopic(NewCmsTopic())
}

func NewCmsTopic() *sCmsTopic {
    return &sCmsTopic{}
}

func (s *sCmsTopic) Model(ctx context.Context) *gdb.Model {
    return dao.CmsTopic.Ctx(ctx).Hook(hook.Bind()).Cache(orm.SetCacheOption(ctx))
}

func (s *sCmsTopic) handleSearch(ctx context.Context, in *req.CmsTopicSearch) (m *gdb.Model) {
    m = s.Model(ctx)

            if !g.IsEmpty(in.Id) {
                m = m.Where("id", in.Id)
            }
        
            if !g.IsEmpty(in.CreatedAt) {
                if len(in.CreatedAt) > 0 {
                    m = m.WhereGTE("created_at", in.CreatedAt[0]+" 00:00:00")
                }
                if len(in.CreatedAt) > 1 {
                    m = m.WhereLTE("created_at", in.CreatedAt[1]+"23:59:59")
                }
            }
        
            if !g.IsEmpty(in.Title) {
            m = m.Where("title like ? ", "%"+in.Title+"%")
            }
        
            if !g.IsEmpty(in.Status) {
                m = m.Where("status", in.Status)
            }
        
    return
}

func (s *sCmsTopic) GetList(ctx context.Context,inReq *model.ListReq, in *req.CmsTopicSearch) (out []*res.CmsTopic, err error) {
    m := s.handleSearch(ctx, in).Handler(handler.FilterAuth)
    m = orm.GetList(m, inReq)
    err = m.Scan(&out)
    if utils.IsError(err) {
        return
    }
    return
}

func (s *sCmsTopic) GetPageList(ctx context.Context, req *model.PageListReq, in *req.CmsTopicSearch) (rs []*res.CmsTopic, total int, err error) {
    m := s.handleSearch(ctx, in).Handler(handler.FilterAuth)
    var entity []*entity.CmsTopic
    err = orm.GetPageList(m, req).ScanAndCount(&entity, &total, false)
    if utils.IsError(err) {
        return nil, 0, err
    }
    rs = make([]*res.CmsTopic, 0)
    if !g.IsEmpty(entity) {
        if err = gconv.Structs(entity, &rs); err != nil {
            return nil, 0, err
        }
    }
    return
}

func (s *sCmsTopic) Save(ctx context.Context, in *req.CmsTopicSave) (id uint64, err error) {
    var saveData *do.CmsTopic
    if err = gconv.Struct(in, &saveData); err != nil {
        return
    }
    rs, err := s.Model(ctx).OmitEmptyData().Data(saveData).Save()
    if utils.IsError(err) {
        return
    }
    tmpId, err := rs.LastInsertId()
    if err != nil {
        return
    }
    id = gconv.Uint64(tmpId)
    return
}

func (s *sCmsTopic) GetById(ctx context.Context, id uint64) (res *res.CmsTopic, err error) {
    err = s.Model(ctx).Where("id", id).Scan(&res)
    if utils.IsError(err) {
        return
    }
    return
}

func (s *sCmsTopic) Update(ctx context.Context, in *req.CmsTopicUpdate) (err error) {
    var updateData *do.CmsTopic
    if err = gconv.Struct(in, &updateData); err != nil {
        return
    }
    _, err = s.Model(ctx).OmitEmptyData().Data(updateData).Where("id", in.Id).Update()
    if utils.IsError(err) {
        return
    }
    return
}

func (s *sCmsTopic) Delete(ctx context.Context, ids []uint64) (err error) {
    _, err = s.Model(ctx).WhereIn("id", ids).Delete()
    if utils.IsError(err) {
        return err
    }
    return
}

func (s *sCmsTopic) RealDelete(ctx context.Context, ids []uint64) (err error) {
    _, err = s.Model(ctx).Unscoped().WhereIn("id", ids).Delete()
    if utils.IsError(err) {
        return
    }
    return
}

func (s *sCmsTopic) Recovery(ctx context.Context, ids []uint64) (err error) {
    _, err = s.Model(ctx).Unscoped().WhereIn("id", ids).Update(g.Map{"deleted_at": nil})
    if utils.IsError(err) {
        return err
    }
    return
}

func (s *sCmsTopic) ChangeStatus(ctx context.Context, id uint64, status int) (err error) {
    _, err = s.Model(ctx).Data(g.Map{"status": status}).Where("id", id).Update()
    if utils.IsError(err) {
        return err
    }
    return
}

func (s *sCmsTopic) GetExportList(ctx context.Context, req *model.ListReq, in *req.CmsTopicSearch) (res []*res.CmsTopicExcel, err error) {
    m := s.handleSearch(ctx, in)
    err = orm.GetList(m, req).Scan(&res)
    if utils.IsError(err) {
        return
    }
    return
}
