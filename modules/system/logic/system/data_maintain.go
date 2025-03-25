// Package system
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE
package system

import (
	"context"
	"devinggo/modules/system/logic/base"
	"devinggo/modules/system/model"
	"devinggo/modules/system/model/req"
	"devinggo/modules/system/model/res"
	"devinggo/modules/system/myerror"
	"devinggo/modules/system/pkg/utils/slice"
	"devinggo/modules/system/service"
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
	switch dbType {
	case "mysql":
		rs, err = s.getMysqlAllTableStatus(ctx, db)
		if err != nil {
			return
		}
		return
	case "pgsql":
		rs, err = s.getPgsqlAllTableStatus(ctx, db)
		if err != nil {
			return
		}
		return
	default:
		err = myerror.ValidationFailed(ctx, "暂不支持该数据库类型")
		return
	}
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

func (s *sDataMaintain) getPgsqlAllTableStatus(ctx context.Context, db gdb.DB) (rs []*res.DataMaintain, err error) {
	query := `
		SELECT 
			tc.table_name as "Name",
			pg_total_relation_size(quote_ident(tc.table_name)) as "Data_length",
			obj_description(quote_ident(tc.table_name)::regclass::oid, 'pg_class') as "Comment",
			to_char(greatest(
				COALESCE(last_vacuum, '1970-01-01'),
				COALESCE(last_autovacuum, '1970-01-01')
			), 'YYYY-MM-DD HH24:MI:SS') as "Update_time",
			'InnoDB' as "Engine",
			'utf8mb4_general_ci' as "Collation",
			0 as "Data_free"
		FROM 
			information_schema.tables tc
			LEFT JOIN pg_stat_user_tables st ON tc.table_name = st.relname
		WHERE 
			tc.table_schema = 'public'
			AND tc.table_type = 'BASE TABLE'
		ORDER BY 
			tc.table_name`

	tablesInfo, err := db.GetAll(ctx, query)
	if err != nil {
		return
	}

	err = gconv.Structs(tablesInfo, &rs)
	if err != nil {
		return
	}
	return
}
