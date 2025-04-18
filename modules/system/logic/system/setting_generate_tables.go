// Package system
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE
package system

import (
	"context"
	"devinggo/internal/dao"
	"devinggo/internal/model/do"
	"devinggo/internal/model/entity"
	"devinggo/modules/system/logic/base"
	"devinggo/modules/system/model"
	"devinggo/modules/system/model/req"
	"devinggo/modules/system/model/res"
	"devinggo/modules/system/myerror"
	"devinggo/modules/system/pkg/hook"
	"devinggo/modules/system/pkg/orm"
	"devinggo/modules/system/pkg/utils"
	"devinggo/modules/system/pkg/utils/idgen"
	"devinggo/modules/system/pkg/utils/slice"
	"devinggo/modules/system/service"
	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gview"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"regexp"
	"strings"
)

type sSettingGenerateTables struct {
	base.BaseService
}

func init() {
	service.RegisterSettingGenerateTables(NewSystemSettingGenerateTables())
}

func NewSystemSettingGenerateTables() *sSettingGenerateTables {
	return &sSettingGenerateTables{}
}

func (s *sSettingGenerateTables) Model(ctx context.Context) *gdb.Model {
	return dao.SettingGenerateTables.Ctx(ctx).Hook(hook.Bind()).Cache(orm.SetCacheOption(ctx)).OnConflict("id")
}

func (s *sSettingGenerateTables) GetPageListForSearch(ctx context.Context, req *model.PageListReq, in *req.SettingGenerateTablesSearch) (rs []*res.SettingGenerateTables, total int, err error) {
	m := s.handleSearch(ctx, in)
	var entity []*entity.SettingGenerateTables
	err = orm.GetPageList(m, req).ScanAndCount(&entity, &total, false)
	if utils.IsError(err) {
		return nil, 0, err
	}
	rs = make([]*res.SettingGenerateTables, 0)
	if !g.IsEmpty(entity) {
		if err = gconv.Structs(entity, &rs); err != nil {
			return nil, 0, err
		}
	}
	return
}

func (s *sSettingGenerateTables) handleSearch(ctx context.Context, in *req.SettingGenerateTablesSearch) (m *gdb.Model) {
	m = s.Model(ctx)
	if !g.IsEmpty(in.TableName) {
		m = m.Where("table_name", in.TableName)
	}
	return
}

func (s *sSettingGenerateTables) LoadTable(ctx context.Context, in *req.LoadTable) (err error) {
	source := in.Source
	names := in.Names
	if g.IsEmpty(names) {
		err = myerror.ValidationFailed(ctx, "加载数据表不能为空")
		return
	}

	for _, loadTableNames := range names {
		insertData := &do.SettingGenerateTables{
			TableName:    loadTableNames.Name,
			TableComment: loadTableNames.Comment,
			MenuName:     loadTableNames.Comment,
			Source:       source,
			Type:         "single",
		}
		result, err := s.Model(ctx).Insert(insertData)
		if utils.IsError(err) {
			return err
		}
		insertId, err := result.LastInsertId()
		if err != nil {
			return err
		}

		columnList, err := service.DataMaintain().GetColumnList(ctx, source, loadTableNames.Name)
		if err != nil {
			return err
		}

		for _, column := range columnList {
			isPk := 1
			if !g.IsEmpty(column.Key) {
				isPk = 2
			}
			isRequired := 1
			if !g.IsEmpty(column.Null) {
				isRequired = 2
			}
			iData := &do.SettingGenerateColumns{
				TableId:       insertId,
				ColumnName:    column.Name,
				ColumnComment: column.Comment,
				ColumnType:    column.Type,
				IsPk:          isPk,
				IsRequired:    isRequired,
				QueryType:     "eq",
				ViewType:      "text",
				Sort:          len(columnList) - column.Index,
				Extra:         column.Extra,
			}
			service.SettingGenerateColumns().Model(ctx).Insert(iData)
		}

	}
	return
}

func (s *sSettingGenerateTables) GetById(ctx context.Context, id int64) (res *res.SettingGenerateTables, err error) {
	err = s.Model(ctx).Where("id", id).Scan(&res)
	if utils.IsError(err) {
		return nil, err
	}
	return
}

func (s *sSettingGenerateTables) Delete(ctx context.Context, ids []int64) (err error) {
	for _, id := range ids {
		_, err = s.Model(ctx).Unscoped().Where("id", id).Delete()
		if utils.IsError(err) {
			return err
		}
		_, err = service.SettingGenerateColumns().Model(ctx).Unscoped().Where("table_id", id).Delete()
		if utils.IsError(err) {
			return err
		}
	}
	return
}

func (s *sSettingGenerateTables) SyncCode(ctx context.Context, id int64) (err error) {
	var entity *entity.SettingGenerateTables
	err = s.Model(ctx).Where("id", id).Scan(&entity)
	if utils.IsError(err) {
		return err
	}
	if g.IsEmpty(entity) {
		return myerror.ValidationFailed(ctx, "数据不存在")
	}

	source := entity.Source
	tableName := entity.TableName
	result, err := service.SettingGenerateColumns().Model(ctx).Unscoped().Fields("id").Where("table_id", id).Array()
	if utils.IsError(err) {
		return
	}

	if g.IsEmpty(result) {
		return
	}

	columnIds := gconv.SliceInt64(result)

	_, err = service.SettingGenerateColumns().Model(ctx).WhereIn("id", columnIds).Delete()

	if utils.IsError(err) {
		return
	}

	columnList, err := service.DataMaintain().GetColumnList(ctx, source, tableName)
	if err != nil {
		return err
	}

	for _, column := range columnList {
		isPk := 1
		if !g.IsEmpty(column.Key) {
			isPk = 2
		}
		isRequired := 1
		if !g.IsEmpty(column.Null) {
			isRequired = 2
		}
		iData := &do.SettingGenerateColumns{
			TableId:       id,
			ColumnName:    column.Name,
			ColumnComment: column.Comment,
			ColumnType:    column.Type,
			IsPk:          isPk,
			IsRequired:    isRequired,
			QueryType:     "eq",
			ViewType:      "text",
			Sort:          len(columnList) - column.Index,
			Extra:         column.Extra,
		}
		service.SettingGenerateColumns().Model(ctx).Insert(iData)
	}

	return
}

func (s *sSettingGenerateTables) UpdateTableAndColumns(ctx context.Context, in *req.TableAndColumnsUpdate) (err error) {

	updateData := &do.SettingGenerateTables{
		GenerateType:  in.GenerateType,
		BuildMenu:     in.BuildMenu,
		MenuName:      in.MenuName,
		ModuleName:    in.ModuleName,
		TableComment:  in.TableComment,
		TableName:     in.TableName,
		Type:          in.Type,
		ComponentType: in.ComponentType,
		PackageName:   in.PackageName,
		Options:       in.Options,
		Remark:        in.Remark,
	}

	if g.IsEmpty(in.BelongMenuId) {
		updateData.BelongMenuId = 0
	} else {
		updateData.BelongMenuId = in.BelongMenuId
	}

	if !g.IsEmpty(in.GenerateMenus) {
		updateData.GenerateMenus = gstr.Join(in.GenerateMenus, ",")
	}
	_, err = s.Model(ctx).Data(updateData).OmitEmptyData().Where("id", in.Id).Update()
	if utils.IsError(err) {
		return
	}

	if !g.IsEmpty(in.Columns) {
		for _, column := range in.Columns {
			isInsert := 1
			if column.IsInsert {
				isInsert = 2
			}

			isEdit := 1
			if column.IsEdit {
				isEdit = 2
			}

			isList := 1
			if column.IsList {
				isList = 2
			}

			isQuery := 1
			if column.IsQuery {
				isQuery = 2
			}

			isRequired := 1
			if column.IsRequired {
				isRequired = 2
			}

			isSort := 1
			if column.IsSort {
				isSort = 2
			}

			updateColumnData := &do.SettingGenerateColumns{
				IsInsert:      isInsert,
				IsEdit:        isEdit,
				IsList:        isList,
				IsQuery:       isQuery,
				IsSort:        isSort,
				IsRequired:    isRequired,
				AllowRoles:    column.AllowRoles,
				ColumnComment: column.ColumnComment,
				ColumnName:    column.ColumnName,
				ColumnType:    column.ColumnType,
				DictType:      column.DictType,
				Extra:         column.Extra,
				IsPk:          column.IsPk,
				Options:       column.Options,
				QueryType:     column.QueryType,
				Remark:        column.Remark,
				Sort:          column.Sort,
				ViewType:      column.ViewType,
			}
			_, err = service.SettingGenerateColumns().Model(ctx).Data(updateColumnData).OmitEmptyData().Where("id", column.Id).Update()
			if utils.IsError(err) {
				return
			}
		}
	}
	return
}

func (s *sSettingGenerateTables) GenerateCode(ctx context.Context, ids []int64) (filePath string, err error) {
	generatePath := utils.GetRootPath() + "/resource/runtime/generate"
	err = s.generateCodeInit(ctx, generatePath)
	if err != nil {
		return
	}
	//codePath := generatePath + "/code"
	codePath := utils.GetTmpDir() + "/code"
	for _, id := range ids {
		err = s.generateOneCode(ctx, codePath, id)
		if err != nil {
			return
		}
	}
	randStr := gconv.String(idgen.NextId(ctx))
	fileName := "devinggo_" + randStr + ".zip"
	filePath, err = s.packageCodeFile(ctx, codePath, generatePath+"/"+fileName)
	if err != nil {
		return
	}
	return
}

func (s *sSettingGenerateTables) packageCodeFile(ctx context.Context, codePath string, zipFilePath string) (filePath string, err error) {
	// 打包目录下面所有文件
	if !gfile.Exists(zipFilePath) {
		err = gfile.RemoveAll(zipFilePath)
		if err != nil {
			return
		}
	}
	err = utils.ZipDirectory(ctx, codePath, zipFilePath)
	if err != nil {
		return
	}
	// 删除目录
	err = gfile.RemoveAll(codePath)
	if err != nil {
		return
	}
	return zipFilePath, nil
}

func (s *sSettingGenerateTables) generateCodeInit(ctx context.Context, dirPath string) (err error) {
	// 检查目录是否存在
	if !gfile.Exists(dirPath) {
		// 如果目录不存在，则创建
		err = gfile.Mkdir(dirPath)
		if err != nil {
			return
		}
	} else {
		// 如果目录存在，则删除目录中的所有内容
		err = gfile.RemoveAll(dirPath)
		if err != nil {
			return
		}
		// 重新创建目录
		err = gfile.Mkdir(dirPath)
		if err != nil {
			return
		}
	}
	return
}

func (s *sSettingGenerateTables) getOneCode(ctx context.Context, id int64) (rs *gmap.StrStrMap, tables *entity.SettingGenerateTables, columns []*entity.SettingGenerateColumns, view *gview.View, err error) {
	err = s.Model(ctx).Where("id", id).Scan(&tables)
	if utils.IsError(err) {
		return
	}
	if g.IsEmpty(tables.ModuleName) {
		tables.ModuleName = "system"
	}

	if g.IsEmpty(tables.Namespace) {
		tables.Namespace = ""
	}

	if g.IsEmpty(tables.BelongMenuId) {
		tables.BelongMenuId = 1000
	}

	if g.IsEmpty(tables.PackageName) {
		tables.PackageName = ""
	}

	if g.IsEmpty(tables.GenerateMenus) {
		tables.GenerateMenus = ""
	}

	err = service.SettingGenerateColumns().Model(ctx).Where("table_id", id).OrderDesc("sort").Scan(&columns)
	if utils.IsError(err) {
		return
	}
	rs = gmap.NewStrStrMap()
	view = gview.New()
	view.SetConfigWithMap(g.Map{
		"Paths":      []string{"resource/generate"},
		"Delimiters": []string{"{%", "%}"},
	})
	view.BindFunc("contains", s.sliceContains)
	view.BindFunc("caseCamel", s.caseCamel)
	view.BindFunc("parseColumnType", s.parseColumnType)
	view.BindFunc("getModelColumnType", s.getModelColumnType)
	apiCode, err := s.generateApi(ctx, view, tables, columns)
	if err != nil {
		return
	}
	rs.Set("apiCode", apiCode)
	modelReqCode, err := s.generateModelReq(ctx, view, tables, columns)
	if err != nil {
		return
	}
	rs.Set("modelReqCode", modelReqCode)
	modelResCode, err := s.generateModelRes(ctx, view, tables, columns)
	if err != nil {
		return
	}
	rs.Set("modelResCode", modelResCode)
	logicCode, err := s.generateLogic(ctx, view, tables, columns)
	if err != nil {
		return
	}
	rs.Set("logicCode", logicCode)
	controllerCode, err := s.generateController(ctx, view, tables, columns)
	if err != nil {
		return
	}
	rs.Set("controllerCode", controllerCode)
	sqlCode, err := s.generateSql(ctx, view, tables, columns)
	if err != nil {
		return
	}
	rs.Set("sqlCode", sqlCode)
	vueCode, err := s.generateVue(ctx, view, tables, columns)
	if err != nil {
		return
	}
	rs.Set("vueCode", vueCode)
	jsApiCode, err := s.generateJsApi(ctx, view, tables, columns)
	if err != nil {
		return
	}
	rs.Set("jsApiCode", jsApiCode)
	return
}

func (s *sSettingGenerateTables) Preview(ctx context.Context, id int64) (rs []res.PreviewTable, err error) {
	rs = make([]res.PreviewTable, 0)
	coders, tables, _, _, err := s.getOneCode(ctx, id)
	if err != nil {
		return
	}
	tableCaseCamelLowerName := gstr.CaseCamelLower(tables.TableName)
	apiCode := coders.Get("apiCode")
	rs = append(rs, res.PreviewTable{TabName: "api/" + tables.PackageName + "/" + tables.TableName + ".go", Code: apiCode, Lang: "go", Name: "api"})
	modelReqCode := coders.Get("modelReqCode")
	rs = append(rs, res.PreviewTable{TabName: "model/req/" + tables.TableName + ".go", Code: modelReqCode, Lang: "go", Name: "req"})
	modelResCode := coders.Get("modelResCode")
	rs = append(rs, res.PreviewTable{TabName: "model/res/" + tables.TableName + ".go", Code: modelResCode, Lang: "go", Name: "res"})
	logicCode := coders.Get("logicCode")
	rs = append(rs, res.PreviewTable{TabName: "logic/" + tables.PackageName + "/" + tables.TableName + ".go", Code: logicCode, Lang: "go", Name: "logic"})
	controllerCode := coders.Get("controllerCode")
	rs = append(rs, res.PreviewTable{TabName: "controller/" + tables.PackageName + "/" + tables.TableName + ".go", Code: controllerCode, Lang: "go", Name: "controller"})
	sqlCode := coders.Get("sqlCode")
	rs = append(rs, res.PreviewTable{TabName: tables.TableName + ".sql", Code: sqlCode, Lang: "sql", Name: "sql"})
	vueCode := coders.Get("vueCode")
	rs = append(rs, res.PreviewTable{TabName: tableCaseCamelLowerName + "/index.vue", Code: vueCode, Lang: "vue", Name: "vue"})
	jsApiCode := coders.Get("jsApiCode")
	rs = append(rs, res.PreviewTable{TabName: tableCaseCamelLowerName + ".js", Code: jsApiCode, Lang: "js", Name: "js"})
	return
}

func (s *sSettingGenerateTables) generateOneCode(ctx context.Context, codePath string, id int64) (err error) {
	coders, tables, _, _, err := s.getOneCode(ctx, id)
	if err != nil {
		return
	}
	tableCaseCamelLowerName := gstr.CaseCamelLower(tables.TableName)
	apiCode := coders.Get("apiCode")
	err = gfile.PutContents(codePath+"/"+tables.ModuleName+"/api/"+tables.PackageName+"/"+tables.TableName+".go", apiCode)
	if err != nil {
		return
	}
	modelReqCode := coders.Get("modelReqCode")
	err = gfile.PutContents(codePath+"/"+tables.ModuleName+"/model/req/"+tables.TableName+".go", modelReqCode)
	if err != nil {
		return
	}
	modelResCode := coders.Get("modelResCode")
	err = gfile.PutContents(codePath+"/"+tables.ModuleName+"/model/res/"+tables.TableName+".go", modelResCode)
	if err != nil {
		return
	}
	logicCode := coders.Get("logicCode")
	err = gfile.PutContents(codePath+"/"+tables.ModuleName+"/logic/"+tables.PackageName+"/"+tables.TableName+".go", logicCode)
	if err != nil {
		return
	}
	controllerCode := coders.Get("controllerCode")
	err = gfile.PutContents(codePath+"/"+tables.ModuleName+"/controller/"+tables.PackageName+"/"+tables.TableName+".go", controllerCode)
	if err != nil {
		return
	}
	sqlCode := coders.Get("sqlCode")
	err = gfile.PutContents(codePath+"/sql/"+tables.TableName+".sql", sqlCode)
	if err != nil {
		return
	}
	vueCode := coders.Get("vueCode")
	err = gfile.PutContents(codePath+"/vue/views/"+tables.ModuleName+"/"+tableCaseCamelLowerName+"/index.vue", vueCode)
	if err != nil {
		return
	}
	jsApiCode := coders.Get("jsApiCode")
	err = gfile.PutContents(codePath+"/vue/api/"+tables.ModuleName+"/"+tableCaseCamelLowerName+".js", jsApiCode)
	if err != nil {
		return
	}
	return
}

func (s *sSettingGenerateTables) generateApi(ctx context.Context, view *gview.View, tables *entity.SettingGenerateTables, columns []*entity.SettingGenerateColumns) (code string, err error) {
	tableCaseCamelName := gstr.CaseCamel(tables.TableName)
	tableCaseCamelLowerName := gstr.CaseCamelLower(tables.TableName)
	generateMenus := gstr.Split(tables.GenerateMenus, ",")
	code, err = view.Parse(context.TODO(), "api/main.html", g.Map{"table": tables, "tableCaseCamelName": tableCaseCamelName, "tableCaseCamelLowerName": tableCaseCamelLowerName, "generateMenus": generateMenus})
	if err != nil {
		return
	}
	code = s.removeExtraBlankLines(code)
	return
}

func (s *sSettingGenerateTables) generateModelReq(ctx context.Context, view *gview.View, tables *entity.SettingGenerateTables, columns []*entity.SettingGenerateColumns) (code string, err error) {
	tableCaseCamelName := gstr.CaseCamel(tables.TableName)
	tableCaseCamelLowerName := gstr.CaseCamelLower(tables.TableName)
	hasGtime := s.hasGtime(ctx, columns, []string{"isInsert", "isEdit", "isQuery"})
	code, err = view.Parse(context.TODO(), "model/req.html", g.Map{"table": tables, "tableCaseCamelName": tableCaseCamelName, "tableCaseCamelLowerName": tableCaseCamelLowerName, "columns": columns, "hasGtime": hasGtime})
	if err != nil {
		return
	}
	code = s.removeExtraBlankLines(code)
	return
}

func (s *sSettingGenerateTables) generateModelRes(ctx context.Context, view *gview.View, tables *entity.SettingGenerateTables, columns []*entity.SettingGenerateColumns) (code string, err error) {
	tableCaseCamelName := gstr.CaseCamel(tables.TableName)
	tableCaseCamelLowerName := gstr.CaseCamelLower(tables.TableName)
	generateMenus := gstr.Split(tables.GenerateMenus, ",")
	hasGtime := s.hasGtime(ctx, columns, []string{"isList"})
	code, err = view.Parse(context.TODO(), "model/res.html", g.Map{"table": tables, "tableCaseCamelName": tableCaseCamelName, "tableCaseCamelLowerName": tableCaseCamelLowerName, "columns": columns, "hasGtime": hasGtime, "generateMenus": generateMenus})
	if err != nil {
		return
	}
	code = s.removeExtraBlankLines(code)
	return
}

func (s *sSettingGenerateTables) generateController(ctx context.Context, view *gview.View, tables *entity.SettingGenerateTables, columns []*entity.SettingGenerateColumns) (code string, err error) {
	tableCaseCamelName := gstr.CaseCamel(tables.TableName)
	tableCaseCamelLowerName := gstr.CaseCamelLower(tables.TableName)
	hasGtime := s.hasGtime(ctx, columns, []string{"isList"})
	generateMenus := gstr.Split(tables.GenerateMenus, ",")
	hasExportOrImport := false
	if s.sliceContains(generateMenus, "export") || s.sliceContains(generateMenus, "import") {
		hasExportOrImport = true
	}
	code, err = view.Parse(context.TODO(), "controller/main.html", g.Map{"table": tables, "generateMenus": generateMenus, "tableCaseCamelName": tableCaseCamelName, "tableCaseCamelLowerName": tableCaseCamelLowerName, "columns": columns, "hasGtime": hasGtime, "hasExportOrImport": hasExportOrImport})
	if err != nil {
		return
	}
	code = s.removeExtraBlankLines(code)
	return
}

func (s *sSettingGenerateTables) generateLogic(ctx context.Context, view *gview.View, tables *entity.SettingGenerateTables, columns []*entity.SettingGenerateColumns) (code string, err error) {
	tableCaseCamelName := gstr.CaseCamel(tables.TableName)
	tableCaseCamelLowerName := gstr.CaseCamelLower(tables.TableName)
	hasGtime := s.hasGtime(ctx, columns, []string{"isList"})
	generateMenus := gstr.Split(tables.GenerateMenus, ",")
	hasExportOrImport := false
	if s.sliceContains(generateMenus, "export") || s.sliceContains(generateMenus, "import") {
		hasExportOrImport = true
	}
	code, err = view.Parse(context.TODO(), "logic/main.html", g.Map{"table": tables, "generateMenus": generateMenus, "tableCaseCamelName": tableCaseCamelName, "tableCaseCamelLowerName": tableCaseCamelLowerName, "columns": columns, "hasGtime": hasGtime, "hasExportOrImport": hasExportOrImport})
	if err != nil {
		return
	}
	code = s.removeExtraBlankLines(code)
	return
}

func (s *sSettingGenerateTables) generateSql(ctx context.Context, view *gview.View, tables *entity.SettingGenerateTables, columns []*entity.SettingGenerateColumns) (code string, err error) {
	tableCaseCamelName := gstr.CaseCamel(tables.TableName)
	tableCaseCamelLowerName := gstr.CaseCamelLower(tables.TableName)
	generateMenus := gstr.Split(tables.GenerateMenus, ",")
	menuTableName := dao.SystemMenu.Table()
	var menu *entity.SystemMenu
	err = service.SystemMenu().Model(ctx).Where("id", tables.BelongMenuId).Scan(&menu)
	if err != nil {
		return
	}
	menu.Level = menu.Level + gconv.String(tables.BelongMenuId) + ","
	adminId := service.SystemUser().GetSupserAdminId(ctx)
	code, err = view.Parse(context.TODO(), "sql/main.html", g.Map{"table": tables, "adminId": adminId, "menu": menu, "generateMenus": generateMenus, "tableCaseCamelName": tableCaseCamelName, "tableCaseCamelLowerName": tableCaseCamelLowerName, "menuTableName": menuTableName})
	if err != nil {
		return
	}
	code = s.removeExtraBlankLines(code)
	return
}

func (s *sSettingGenerateTables) generateVue(ctx context.Context, view *gview.View, tables *entity.SettingGenerateTables, columns []*entity.SettingGenerateColumns) (code string, err error) {
	tableCaseCamelName := gstr.CaseCamel(tables.TableName)
	tableCaseCamelLowerName := gstr.CaseCamelLower(tables.TableName)
	generateMenus := gstr.Split(tables.GenerateMenus, ",")
	adminId := service.SystemUser().GetSupserAdminId(ctx)
	columnsView, err := s.getColumns(ctx, columns)
	if err != nil {
		return
	}
	OptionsView, err := s.getOptions(ctx, tables, columns)
	if err != nil {
		return
	}
	authCode := tables.PackageName + ":" + tableCaseCamelLowerName
	code, err = view.Parse(context.TODO(), "vue/main.html", g.Map{"table": tables, "authCode": authCode, "adminId": adminId, "OptionsView": OptionsView, "columnsView": columnsView, "generateMenus": generateMenus, "tableCaseCamelName": tableCaseCamelName, "tableCaseCamelLowerName": tableCaseCamelLowerName})
	if err != nil {
		return
	}
	code = s.removeExtraBlankLines(code)
	return
}

func (s *sSettingGenerateTables) generateJsApi(ctx context.Context, view *gview.View, tables *entity.SettingGenerateTables, columns []*entity.SettingGenerateColumns) (code string, err error) {
	tableCaseCamelName := gstr.CaseCamel(tables.TableName)
	tableCaseCamelLowerName := gstr.CaseCamelLower(tables.TableName)
	generateMenus := gstr.Split(tables.GenerateMenus, ",")
	adminId := service.SystemUser().GetSupserAdminId(ctx)
	columnsView, err := s.getColumns(ctx, columns)
	if err != nil {
		return
	}
	OptionsView, err := s.getOptions(ctx, tables, columns)
	if err != nil {
		return
	}

	requestRoute := gstr.ToLower(tables.ModuleName) + "/" + tableCaseCamelLowerName

	authCode := tables.PackageName + ":" + tableCaseCamelLowerName
	code, err = view.Parse(context.TODO(), "vue/jsApi.html", g.Map{"table": tables, "requestRoute": requestRoute, "authCode": authCode, "adminId": adminId, "OptionsView": OptionsView, "columnsView": columnsView, "generateMenus": generateMenus, "tableCaseCamelName": tableCaseCamelName, "tableCaseCamelLowerName": tableCaseCamelLowerName})
	if err != nil {
		return
	}
	code = s.removeExtraBlankLines(code)
	return
}

func (s *sSettingGenerateTables) getColumns(ctx context.Context, columns []*entity.SettingGenerateColumns) (string, error) {
	options := make([]*gmap.StrAnyMap, 0)
	for _, column := range columns {
		tmp := gmap.NewStrAnyMap()
		tmp.Set("title", column.ColumnComment)
		tmp.Set("dataIndex", column.ColumnName)
		tmp.Set("formType", s.getViewType(column.ViewType))
		if column.IsQuery == 2 {
			tmp.Set("search", true)
		} else {
			tmp.Set("search", false)
		}
		if column.IsInsert == 2 {
			tmp.Set("addDisplay", true)
		} else {
			tmp.Set("addDisplay", false)
		}
		if column.IsEdit == 2 {
			tmp.Set("editDisplay", true)
		} else {
			tmp.Set("editDisplay", false)
		}
		if column.IsList == 2 {
			tmp.Set("hide", false)
		} else {
			tmp.Set("hide", true)
		}
		if column.IsRequired == 2 {
			tmp.Set("commonRules", g.Map{"required": true, "message": "请输入" + column.ColumnComment})
		} else {
			tmp.Set("commonRules", g.Map{"required": false, "message": "请输入" + column.ColumnComment})
		}
		if column.IsSort == 2 {
			tmp.Set("sortable", g.Map{"sortDirections": []string{"ascend", "descend"}, "sorter": true})
		} else {
			tmp.Set("sortable", g.Map{})
		}

		if !g.IsEmpty(column.Options) {
			j, err := gjson.DecodeToJson(column.Options)
			if err != nil {
				return "", err
			}
			collection := g.Map{}
			if !g.IsEmpty(j.Get("collection")) {
				collection = j.Get("collection").Map()
			}
			// collection 与 tmp 合并
			tmp.Merge(gmap.NewStrAnyMapFrom(j.Map()))
			if (column.ViewType == "select" || column.ViewType == "radio" || column.ViewType == "checkbox" || column.ViewType == "transfer") && !g.IsEmpty(collection) {
				tmp.Set("dict", g.Map{"data": collection, "translation": true})
			}

			if column.ViewType == "date" && j.Get("mode").String() == "date" {
				tmp.Remove("mode")
				if !g.IsEmpty(j.Get("range")) {
					tmp.Set("formType", "range")
					tmp.Remove("range")
				}
			}
			tmp.Remove("collection")
		}

		if !g.IsEmpty(column.DictType) {
			tmp.Set("dict", g.Map{"name": column.DictType, "props": g.Map{"label": "title", "value": "key"}, "translation": true})
		}

		if column.ViewType == "password" {
			tmp.Set("type", "password")
		}
		options = append(options, tmp)
	}
	if g.IsEmpty(options) {
		return "", nil
	} else {
		optionsJsonSTr, err := s.jsonFormat(options, false)
		if err != nil {
			return "", err
		}
		return "const columns = reactive(" + optionsJsonSTr + ")", nil
	}
}

func (s *sSettingGenerateTables) getOptions(ctx context.Context, tables *entity.SettingGenerateTables, columns []*entity.SettingGenerateColumns) (string, error) {
	tableCaseCamelLowerName := gstr.CaseCamelLower(tables.TableName)
	generateMenus := gstr.Split(tables.GenerateMenus, ",")

	options := gmap.NewStrAnyMap()
	options.Set("id", "'"+tables.TableName+"'")
	options.Set("rowSelection", g.Map{"showCheckedAll": true})
	options.Set("pk", "'"+s.getPk(columns)+"'")
	options.Set("operationColumn", false)
	options.Set("operationColumnWidth", 160)
	options.Set("formOption", g.Map{"viewType": "'" + s.getComponentType(tables) + "'", "width": 600})

	jOptions := gjson.New(tables.Options)

	if s.getComponentType(tables) == "tag" {
		tagId := tables.TableName
		if !g.IsEmpty(jOptions.Get("tag_id")) {
			tagId = jOptions.Get("tag_id").String()
		}
		tagName := tables.TableComment
		if !g.IsEmpty(jOptions.Get("tag_name")) {
			tagName = jOptions.Get("tag_name").String()
		}
		titleDataIndex := s.getPk(columns)
		if !g.IsEmpty(jOptions.Get("tag_view_name")) {
			titleDataIndex = jOptions.Get("tag_view_name").String()
		}
		options.Set("formOption", g.Map{
			"tagId":          "'" + tagId + "'",
			"tagName":        "'" + tagName + "'",
			"titleDataIndex": "'" + titleDataIndex + "'",
		})
	}
	authCode := tables.PackageName + ":" + tableCaseCamelLowerName
	options.Set("api", tableCaseCamelLowerName+".getPageList")
	if s.sliceContains(generateMenus, "recycle") {
		options.Set("recycleApi", tableCaseCamelLowerName+".getPageRecycleList")
	}
	if s.sliceContains(generateMenus, "save") {
		options.Set("add", g.Map{"show": true, "api": tableCaseCamelLowerName + ".save", "auth": []string{authCode + ":save"}})
	}
	if s.sliceContains(generateMenus, "update") {
		options.Set("operationColumn", true)
		options.Set("edit", g.Map{"show": true, "api": tableCaseCamelLowerName + ".update", "auth": []string{authCode + ":update"}})
	}
	if s.sliceContains(generateMenus, "delete") {
		options.Set("operationColumn", true)
		options.Set("delete", g.Map{"show": true, "api": tableCaseCamelLowerName + ".deletes", "auth": []string{authCode + ":delete"}})
		if s.sliceContains(generateMenus, "recycle") {
			options.Set("delete", g.Map{"show": true, "api": tableCaseCamelLowerName + ".deletes", "auth": []string{authCode + ":delete"}, "realApi": tableCaseCamelLowerName + ".realDeletes", "realAuth": []string{authCode + ":realDelete"}})
			options.Set("recovery", g.Map{"show": true, "api": tableCaseCamelLowerName + ".recoverys", "auth": []string{authCode + ":recovery"}})
		}
	}
	requestRoute := gstr.ToLower(tables.ModuleName) + "/" + tableCaseCamelLowerName
	if s.sliceContains(generateMenus, "import") {
		options.Set("import", g.Map{"show": true, "url": "'" + requestRoute + "/import'", "templateUrl": "'" + requestRoute + "/downloadTemplate'", "auth": []string{authCode + ":import"}})
	}
	if s.sliceContains(generateMenus, "export") {
		options.Set("export", g.Map{"show": true, "url": "'" + requestRoute + "/export'", "auth": []string{authCode + ":export"}})
	}
	optionsJsonSTr, err := s.jsonFormat(options, true)
	if err != nil {
		return "", err
	}
	return "const options = reactive(" + optionsJsonSTr + ")", nil
}

func (s *sSettingGenerateTables) getComponentType(tables *entity.SettingGenerateTables) string {
	switch tables.ComponentType {
	case 1:
		return "modal"
	case 2:
		return "drawer"
	default:
		return "tag"
	}
}

func (s *sSettingGenerateTables) getPk(columns []*entity.SettingGenerateColumns) string {
	for _, column := range columns {
		if column.IsPk == 2 {
			return column.ColumnName
		}
	}
	return ""
}

func (s *sSettingGenerateTables) jsonFormat(data interface{}, removeValueQuotes bool) (string, error) {
	// 将数据转换为 JSON 格式，并使用缩进
	jsonData := gjson.New(data)

	// 将 JSON 字符串转换为可操作的字符串
	jsonString, err := jsonData.ToJsonIndentString()
	if err != nil {
		return "", err
	}
	// 替换 "true" 和 "false" 为 true 和 false
	jsonString = gstr.Replace(jsonString, `"true"`, "true")
	jsonString = gstr.Replace(jsonString, `"false"`, "false")

	// 删除反斜杠
	jsonString = gstr.Replace(jsonString, `\\`, "")

	// 删除键的引号
	re := regexp.MustCompile(`(\s+)"(.+)":`)
	jsonString = re.ReplaceAllString(jsonString, `$1$2:`)

	// 如果 removeValueQuotes 为 true，则删除值的引号
	if removeValueQuotes {
		re = regexp.MustCompile(`:\s+"([^"]+)"`)
		jsonString = re.ReplaceAllString(jsonString, `: $1`)
	}

	return jsonString, nil
}

func (s *sSettingGenerateTables) getViewType(viewType string) string {
	viewTypes := map[string]string{
		"text":           "input",
		"password":       "input-password",
		"textarea":       "textarea",
		"inputNumber":    "input-number",
		"inputTag":       "input-tag",
		"mention":        "mention",
		"switch":         "switch",
		"slider":         "slider",
		"select":         "select",
		"radio":          "radio",
		"checkbox":       "checkbox",
		"treeSelect":     "tree-select",
		"date":           "date",
		"time":           "time",
		"rate":           "rate",
		"cascader":       "cascader",
		"transfer":       "transfer",
		"selectUser":     "user-select",
		"userInfo":       "user-info",
		"cityLinkage":    "city-linkage",
		"icon":           "icon-picker",
		"formGroup":      "form-group",
		"upload":         "upload",
		"selectResource": "resource",
		"editor":         "editor",
		"wangEditor":     "wang-editor",
		"codeEditor":     "code-editor",
	}
	if _, ok := viewTypes[viewType]; ok {
		return viewTypes[viewType]
	} else {
		return "input"
	}
}

func (s *sSettingGenerateTables) hasGtime(ctx context.Context, columns []*entity.SettingGenerateColumns, dataTypes []string) bool {
	for _, dataType := range dataTypes {
		for _, column := range columns {
			if column.IsInsert == 2 && (dataType == "isInsert") {
				parseColumnType := s.parseColumnType(*column)
				if parseColumnType == "*gtime.Time" {
					return true
				}
			} else if column.IsEdit == 2 && (dataType == "isEdit") {
				parseColumnType := s.parseColumnType(*column)
				if parseColumnType == "*gtime.Time" {
					return true
				}
			} else if column.IsList == 2 && (dataType == "isList") {
				parseColumnType := s.getModelColumnType(*column)
				if parseColumnType == "*gtime.Time" {
					return true
				}
			} else if column.IsQuery == 2 && (dataType == "isQuery") {
				parseColumnType := s.parseColumnType(*column)
				if parseColumnType == "*gtime.Time" {
					return true
				}
			}
		}
	}
	return false
}

func (s *sSettingGenerateTables) removeExtraBlankLines(input string) string {
	// 按换行符分割字符串
	lines := gstr.Split(input, "\n")

	// 用于存储处理后的行
	var resultLines []string

	// 用于记录连续空白行的数量
	blankCount := 0

	for _, line := range lines {
		// 去除行首尾的空白字符
		trimmedLine := gstr.Trim(line)

		if trimmedLine == "" {
			// 如果是空白行，增加空白行计数
			blankCount++
		} else {
			// 如果不是空白行，重置空白行计数
			blankCount = 0
		}

		// 如果空白行数量小于等于 2，保留该行
		if blankCount < 2 {
			resultLines = append(resultLines, line)
		}
	}

	// 将处理后的行重新组合成字符串
	return gstr.Join(resultLines, "\n")
}

func (s *sSettingGenerateTables) sliceContains(arr []string, value string) bool {
	return slice.Contains(arr, value)
}

func (s *sSettingGenerateTables) caseCamel(value string) string {
	return gstr.CaseCamel(value)
}

func (s *sSettingGenerateTables) parseDataType(value string) (dataType string, limit string) {
	parts := strings.Split(value, "(")
	if len(parts) > 1 {
		dataType = parts[0]
		limit = strings.TrimRight(parts[1], ")")
	} else {
		dataType = value
		limit = ""
	}
	return
}

func (s *sSettingGenerateTables) parseColumnType(entity entity.SettingGenerateColumns) string {
	queryType := entity.QueryType
	if queryType == "like" {
		return "string"
	} else if queryType == "between" {
		return "[]string"
	} else if queryType == "in" {
		return "[]string"
	} else if queryType == "notin" {
		return "[]string"
	} else {
		return s.getModelColumnType(entity)
	}
}

func (s *sSettingGenerateTables) getModelColumnType(entity entity.SettingGenerateColumns) string {
	value := entity.ColumnType
	dataType, _ := s.parseDataType(value)
	//fmt.Println(value, "-", dataType)
	switch dataType {
	case "decimal":
		return "float64"
	case "date":
		return "*gtime.Time"
	case "time":
		return "*gtime.Time"
	case "year":
		return "int"
	case "blob", "tinyblob", "mediumblob", "longblob":
		return "[]byte"
	case "enum":
		return "string"
	case "set":
		return "string"
	case "float":
		return "float32"
	case "double":
		return "float64"
	case "bigint":
		return "int64"
	case "int", "tinyint", "smallint", "mediumint":
		return "int"
	case "datetime", "timestamp":
		return "*gtime.Time"
	case "bool":
		return "int"
	case "text", "varchar", "char", "longtext", "mediumtext", "tinytext":
		return "string"
	default:
		return "string"
	}
}
