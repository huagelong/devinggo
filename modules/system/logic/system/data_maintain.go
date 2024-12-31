// Package system
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE
package system

import (
	"devinggo/internal/logic/base"
	"devinggo/internal/model"
	"devinggo/modules/system/model/req"
	"devinggo/modules/system/model/res"
	"devinggo/modules/system/myerror"
	"devinggo/modules/system/pkg/utils/slice"
	"devinggo/modules/system/service"
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"strings"
)

type sDataMaintain struct {
	base.BaseService
}

func init() {
	service.RegisterDataMaintain(NewSystemDataMaintain())
}

func NewSystemDataMaintain() *sDataMaintain {
	return &sDataMaintain{}
}

func (s *sDataMaintain) GetPageListForSearch(ctx context.Context, req *model.PageListReq, in *req.DataMaintainSearch) (rs []*res.DataMaintain, total int, err error) {
	allList, err := s.GetAllTableStatus(ctx, in.GroupName)
	if err != nil {
		return
	}

	if g.IsEmpty(allList) {
		return
	}
	if !g.IsEmpty(in.Name) {
		allListNew := make([]*res.DataMaintain, 0)
		for _, v := range allList {
			if v.Name == in.Name {
				allListNew = append(allListNew, v)
			}
		}
		allList = allListNew
	}

	rs, err = slice.Paginate[*res.DataMaintain](allList, req.PageSize, req.Page)
	if err != nil {
		return
	}
	total = len(allList)
	return
}

func (s *sDataMaintain) GetColumnList(ctx context.Context, source, tableName string) (rs map[string]*gdb.TableField, err error) {
	db := g.DB(source)
	rs, err = db.TableFields(ctx, tableName)
	if err != nil {
		return
	}
	return
}

func (s *sDataMaintain) GetAllTableStatus(ctx context.Context, groupName string) (rs []*res.DataMaintain, err error) {
	if g.IsEmpty(groupName) {
		groupName = "default"
	}
	db := g.DB(groupName)
	if db == nil {
		err = myerror.ValidationFailed(ctx, "数据库组不存在")
		return
	}
	dbType := strings.ToLower(db.GetConfig().Type)
	if dbType == "mysql" {
		rs, err = s.getMysqlAllTableStatus(ctx, db)
		if err != nil {
			return
		}
		return
	} else {
		err = myerror.ValidationFailed(ctx, "暂不支持该数据库类型")
		return
	}
	return
}

func (s *sDataMaintain) getMysqlAllTableStatus(ctx context.Context, db gdb.DB) (rs []*res.DataMaintain, err error) {
	tablesInfo, err := db.GetAll(ctx, "SHOW TABLE STATUS")
	if err != nil {
		return
	}
	//g.Log().Info(ctx, "tablesInfo:", tablesInfo)

	err = gconv.Structs(tablesInfo, &rs)
	if err != nil {
		return
	}
	//g.Log().Info(ctx, "rs:", rs)
	return
}
