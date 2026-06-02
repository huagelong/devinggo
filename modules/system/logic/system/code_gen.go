// Package system
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package system

import (
	"archive/zip"
	"bytes"
	"context"

	"devinggo/internal/dao"
	"devinggo/internal/model/entity"
	"devinggo/modules/system/logic/base"
	"devinggo/modules/system/model"
	"devinggo/modules/system/model/req"
	"devinggo/modules/system/model/res"
	"devinggo/modules/system/pkg/handler"
	"devinggo/modules/system/pkg/hook"
	"devinggo/modules/system/pkg/orm"
	"devinggo/modules/system/pkg/utils"
	"devinggo/modules/system/service"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

type sCodeGen struct {
	base.GenericService[res.CodeGenTable]
}

func init() {
	service.RegisterCodeGen(NewCodeGen())
}

func NewCodeGen() *sCodeGen {
	s := &sCodeGen{}
	s.GenericService = base.GenericService[res.CodeGenTable]{
		ModelFn: func(ctx context.Context) *gdb.Model {
			return dao.CodeGenTables.Ctx(ctx).Hook(hook.Default()).Cache(orm.SetCacheOption(ctx)).OnConflict("id")
		},
	}
	return s
}

// Model 返回数据库 Model
func (s *sCodeGen) Model(ctx context.Context) *gdb.Model {
	return s.GenericService.Model(ctx)
}

// GetPageList 获取分页列表
func (s *sCodeGen) GetPageList(ctx context.Context, req *model.PageListReq, in *req.CodeGenSearch) (rs []*res.CodeGenTable, total int, err error) {
	m := s.handleSearch(ctx, in).Handler(handler.FilterAuth)
	var entity []*entity.CodeGenTables
	err = orm.NewQuery(m).WithPageListReq(req).ScanAndCount(&entity, &total)
	if utils.IsError(err) {
		return nil, 0, err
	}
	rs = make([]*res.CodeGenTable, 0)
	if !g.IsEmpty(entity) {
		if err = gconv.Structs(entity, &rs); utils.IsError(err) {
			return nil, 0, err
		}
	}
	return
}

// handleSearch 处理搜索条件
func (s *sCodeGen) handleSearch(ctx context.Context, in *req.CodeGenSearch) *gdb.Model {
	m := s.Model(ctx)
	if !g.IsEmpty(in) {
		if in.TableName != "" {
			m = m.Where("table_name like ?", "%"+in.TableName+"%")
		}
		if in.Type != "" {
			m = m.Where("type", in.Type)
		}
	}
	return m.WhereNotNull("deleted_at").OrderDesc("id")
}

// Delete 删除记录
func (s *sCodeGen) Delete(ctx context.Context, ids []int64) (err error) {
	_, err = s.Model(ctx).Where("id in (?)", ids).Data(g.Map{
		"deleted_at": gtime.Now(),
	}).Update()
	return
}

// Update 更新配置
func (s *sCodeGen) Update(ctx context.Context, in *req.CodeGenUpdate, userId int64) (err error) {
	// 更新主表
	mainData := g.Map{
		"table_comment":  in.TableComment,
		"remark":         in.Remark,
		"module_name":    in.ModuleName,
		"belong_menu_id": in.BelongMenuId,
		"type":           in.Type,
		"menu_name":      in.MenuName,
		"component_type": in.ComponentType,
		"tpl_type":       in.TplType,
		"tree_id":        in.TreeId,
		"tree_parent_id": in.TreeParentId,
		"tree_name":      in.TreeName,
		"tag_id":         in.TagId,
		"tag_name":       in.TagName,
		"tag_view_name":  in.TagViewName,
		"updated_by":     userId,
		"updated_at":     gtime.Now(),
	}
	_, err = s.Model(ctx).Where("id", in.Id).Data(mainData).Update()
	if utils.IsError(err) {
		return
	}

	// 更新字段配置
	if !g.IsEmpty(in.Fields) {
		// 删除旧字段
		_, err = dao.CodeGenFields.Ctx(ctx).Where("table_id", in.Id).Delete()
		if utils.IsError(err) {
			return
		}
		// 插入新字段
		for _, field := range in.Fields {
			_, err = dao.CodeGenFields.Ctx(ctx).Data(g.Map{
				"table_id":       in.Id,
				"column_name":    field.ColumnName,
				"column_comment": field.ColumnComment,
				"column_type":    field.ColumnType,
				"sort":           field.Sort,
				"is_required":    field.IsRequired,
				"is_insert":      field.IsInsert,
				"is_edit":        field.IsEdit,
				"is_list":        field.IsList,
				"is_query":       field.IsQuery,
				"is_sort":        field.IsSort,
				"query_type":     field.QueryType,
				"view_type":      field.ViewType,
				"dict_type":      field.DictType,
				"allow_roles":    field.AllowRoles,
			}).Insert()
			if utils.IsError(err) {
				return
			}
		}
	}

	// 更新按钮配置
	if !g.IsEmpty(in.MenuButtons) {
		// 删除旧按钮
		_, err = dao.CodeGenButtons.Ctx(ctx).Where("table_id", in.Id).Delete()
		if utils.IsError(err) {
			return
		}
		// 倒序插入新按钮
		buttons := in.MenuButtons
		for i := len(buttons) - 1; i >= 0; i-- {
			code := buttons[i]
			_, err = dao.CodeGenButtons.Ctx(ctx).Data(g.Map{
				"table_id":    in.Id,
				"button_code": code,
				"button_name": getButtonName(code),
				"is_show":     1,
				"sort":        len(buttons) - i,
			}).Insert()
			if utils.IsError(err) {
				return
			}
		}
	}

	return
}

// LoadTable 装载数据表
func (s *sCodeGen) LoadTable(ctx context.Context, in *req.CodeGenLoadTable, userId int64) (err error) {
	for _, item := range in.Names {
		// 检查是否已存在
		var existing int
		existing, err = s.Model(ctx).Where("table_name", item.Name).Count()
		if utils.IsError(err) {
			return
		}
		if existing > 0 {
			continue
		}

		// 插入主表记录
		var id int64
		id, err = s.Model(ctx).Data(g.Map{
			"table_name":     item.Name,
			"table_comment":  item.Comment,
			"type":           "single",
			"component_type": 1,
			"tpl_type":       "default",
			"created_by":     userId,
			"created_at":     gtime.Now(),
			"status":         1,
			"sort":           0,
		}).InsertAndGetId()
		if utils.IsError(err) {
			return
		}

		// 从数据库读取表结构
		columns, err := s.getTableColumns(ctx, item.Name)
		if utils.IsError(err) {
			continue
		}

		// 插入字段配置
		for i, col := range columns {
			_, err = dao.CodeGenFields.Ctx(ctx).Data(g.Map{
				"table_id":       id,
				"column_name":    col.ColumnName,
				"column_comment": col.ColumnComment,
				"column_type":    col.ColumnType,
				"data_type":      col.DataType,
				"is_nullable":    col.IsNullable,
				"sort":           i + 1,
				"is_required":    2,
				"is_insert":      2,
				"is_edit":        2,
				"is_list":        2,
				"is_query":       2,
				"is_sort":        2,
				"view_type":      s.getDefaultViewType(col.DataType),
			}).Insert()
			if utils.IsError(err) {
				continue
			}
		}
	}
	return
}

// getTableColumns 获取表的列信息
func (s *sCodeGen) getTableColumns(ctx context.Context, tableName string) ([]res.CodeGenColumn, error) {
	db := g.DB()
	sql := `SELECT column_name, column_comment, column_type, data_type, is_nullable
		FROM information_schema.columns
		WHERE table_name = ? AND table_schema = current_schema()`
	resultSet, err := db.GetCore().Query(ctx, sql, tableName)
	if utils.IsError(err) {
		return nil, err
	}
	result := make([]res.CodeGenColumn, 0, len(resultSet))
	for _, row := range resultSet {
		col := res.CodeGenColumn{
			ColumnName:    row["column_name"].String(),
			ColumnComment: row["column_comment"].String(),
			ColumnType:    row["column_type"].String(),
			DataType:      row["data_type"].String(),
			IsNullable:    row["is_nullable"].String(),
		}
		result = append(result, col)
	}
	return result, nil
}

// getDefaultViewType 根据数据类型获取默认的视图类型
func (s *sCodeGen) getDefaultViewType(dataType string) string {
	switch {
	case gstr.Contains(dataType, "int"), gstr.Contains(dataType, "bigint"), gstr.Contains(dataType, "smallint"):
		return "inputNumber"
	case gstr.Contains(dataType, "decimal"), gstr.Contains(dataType, "float"), gstr.Contains(dataType, "double"):
		return "inputNumber"
	case gstr.Contains(dataType, "text"), gstr.Contains(dataType, "varchar"):
		return "input"
	case gstr.Contains(dataType, "date"), gstr.Contains(dataType, "time"), gstr.Contains(dataType, "timestamp"):
		return "date"
	case gstr.Contains(dataType, "json"):
		return "input"
	default:
		return "input"
	}
}

// SyncTable 同步数据表结构
func (s *sCodeGen) SyncTable(ctx context.Context, id int64, userId int64) (err error) {
	var table entity.CodeGenTables
	err = s.Model(ctx).Where("id", id).Scan(&table)
	if utils.IsError(err) {
		return
	}
	if g.IsEmpty(table) {
		return nil
	}

	// 删除旧字段
	_, err = dao.CodeGenFields.Ctx(ctx).Where("table_id", id).Delete()
	if utils.IsError(err) {
		return
	}

	// 重新获取字段
	columns, err := s.getTableColumns(ctx, table.TableName)
	if utils.IsError(err) {
		return
	}

	// 插入新字段
	for i, col := range columns {
		_, err = dao.CodeGenFields.Ctx(ctx).Data(g.Map{
			"table_id":       id,
			"column_name":    col.ColumnName,
			"column_comment": col.ColumnComment,
			"column_type":    col.ColumnType,
			"data_type":      col.DataType,
			"is_nullable":    col.IsNullable,
			"sort":           i + 1,
			"is_required":    2,
			"is_insert":      2,
			"is_edit":        2,
			"is_list":        2,
			"is_query":       2,
			"is_sort":        2,
			"view_type":      s.getDefaultViewType(col.DataType),
		}).Insert()
		if utils.IsError(err) {
			continue
		}
	}

	// 更新同步时间
	_, err = s.Model(ctx).Where("id", id).Data(g.Map{
		"updated_by": userId,
		"updated_at": gtime.Now(),
	}).Update()
	return
}

// GenerateCode 生成代码
func (s *sCodeGen) GenerateCode(ctx context.Context, ids string) (fileBytes []byte, err error) {
	// 解析 IDs
	idArr := gconv.Ints(gstr.Split(ids, ","))
	if len(idArr) == 0 {
		return nil, nil
	}

	// 获取表配置
	var tables []entity.CodeGenTables
	err = s.Model(ctx).Where("id in (?)", idArr).Scan(&tables)
	if utils.IsError(err) {
		return nil, err
	}

	// 创建 ZIP 文件
	buf := new(bytes.Buffer)
	zipWriter := zip.NewWriter(buf)

	for _, table := range tables {
		// 获取字段配置
		var fields []entity.CodeGenFields
		err = dao.CodeGenFields.Ctx(ctx).Where("table_id", table.Id).OrderAsc("sort").Scan(&fields)
		if utils.IsError(err) {
			continue
		}

		// 生成 Entity 文件
		entityCode := s.generateEntityCode(&table, fields)
		err = s.addFileToZip(zipWriter, table.TableName+"/entity/"+table.TableName+".go", []byte(entityCode))
		if utils.IsError(err) {
			continue
		}

		// 生成 API 文件
		apiCode := s.generateApiCode(&table, fields)
		err = s.addFileToZip(zipWriter, table.TableName+"/api/"+table.TableName+".go", []byte(apiCode))
		if utils.IsError(err) {
			continue
		}

		// 生成 Controller 文件
		controllerCode := s.generateControllerCode(&table, fields)
		err = s.addFileToZip(zipWriter, table.TableName+"/controller/"+table.TableName+".go", []byte(controllerCode))
		if utils.IsError(err) {
			continue
		}

		// 生成 Logic 文件
		logicCode := s.generateLogicCode(&table, fields)
		err = s.addFileToZip(zipWriter, table.TableName+"/logic/"+table.TableName+".go", []byte(logicCode))
		if utils.IsError(err) {
			continue
		}

		// 生成 Model 文件
		modelCode := s.generateModelCode(&table, fields)
		err = s.addFileToZip(zipWriter, table.TableName+"/model/"+table.TableName+".go", []byte(modelCode))
		if utils.IsError(err) {
			continue
		}
	}

	// 关闭 ZIP writer
	err = zipWriter.Close()
	if utils.IsError(err) {
		return nil, err
	}

	fileBytes = buf.Bytes()
	return
}

// addFileToZip 添加文件到 ZIP
func (s *sCodeGen) addFileToZip(zipWriter *zip.Writer, filename string, content []byte) error {
	w, err := zipWriter.Create(filename)
	if utils.IsError(err) {
		return err
	}
	_, err = w.Write(content)
	return err
}

// getGoType 将数据库类型转换为 Go 类型
func (s *sCodeGen) getGoType(columnType string) string {
	switch {
	case gstr.Contains(columnType, "int"):
		return "int64"
	case gstr.Contains(columnType, "decimal"), gstr.Contains(columnType, "float"), gstr.Contains(columnType, "double"):
		return "float64"
	case gstr.Contains(columnType, "bool"):
		return "bool"
	default:
		return "string"
	}
}

// PreviewCode 预览代码
func (s *sCodeGen) PreviewCode(ctx context.Context, id int64) (preview []res.CodeGenPreview, err error) {
	var table entity.CodeGenTables
	err = s.Model(ctx).Where("id", id).Scan(&table)
	if utils.IsError(err) {
		return
	}

	// 获取字段配置
	var fields []entity.CodeGenFields
	err = dao.CodeGenFields.Ctx(ctx).Where("table_id", id).OrderAsc("sort").Scan(&fields)
	if utils.IsError(err) {
		return
	}

	// 生成代码预览
	preview = append(preview, res.CodeGenPreview{
		Name:    table.TableName + "_entity.go",
		TabName: "Entity",
		Code:    s.generateEntityCode(&table, fields),
		Lang:    "go",
	})

	preview = append(preview, res.CodeGenPreview{
		Name:    table.TableName + "_api.go",
		TabName: "API",
		Code:    s.generateApiCode(&table, fields),
		Lang:    "go",
	})

	preview = append(preview, res.CodeGenPreview{
		Name:    table.TableName + "_controller.go",
		TabName: "Controller",
		Code:    s.generateControllerCode(&table, fields),
		Lang:    "go",
	})

	preview = append(preview, res.CodeGenPreview{
		Name:    table.TableName + "_logic.go",
		TabName: "Logic",
		Code:    s.generateLogicCode(&table, fields),
		Lang:    "go",
	})

	preview = append(preview, res.CodeGenPreview{
		Name:    table.TableName + "_model.go",
		TabName: "Model",
		Code:    s.generateModelCode(&table, fields),
		Lang:    "go",
	})

	return
}

// generateEntityCode 生成 Entity 代码
func (s *sCodeGen) generateEntityCode(table *entity.CodeGenTables, fields []entity.CodeGenFields) string {
	var buf bytes.Buffer
	buf.WriteString("package entity\n\n")
	buf.WriteString("import \"github.com/gogf/gf/v2/os/gtime\"\n\n")
	buf.WriteString("type " + gstr.CaseCamel(table.TableName) + " struct {\n")
	buf.WriteString("\tId        int64       `json:\"id\" orm:\"id\"`\n")
	for _, f := range fields {
		if f.ColumnName == "id" || f.ColumnName == "created_at" || f.ColumnName == "updated_at" || f.ColumnName == "deleted_at" || f.ColumnName == "created_by" || f.ColumnName == "updated_by" {
			continue
		}
		fieldType := s.getGoType(f.ColumnType)
		jsonName := gstr.ToLower(f.ColumnName)
		buf.WriteString("\t" + gstr.CaseCamel(f.ColumnName) + " " + fieldType + " `json:\"" + jsonName + "\" orm:\"" + f.ColumnName + "\"`\n")
	}
	buf.WriteString("\tCreatedAt *gtime.Time `json:\"createdAt\" orm:\"created_at\"`\n")
	buf.WriteString("\tUpdatedAt *gtime.Time `json:\"updatedAt\" orm:\"updated_at\"`\n")
	buf.WriteString("}\n")
	return buf.String()
}

// generateApiCode 生成 API 代码
func (s *sCodeGen) generateApiCode(table *entity.CodeGenTables, fields []entity.CodeGenFields) string {
	var buf bytes.Buffer
	entityName := gstr.CaseCamel(table.TableName)
	varName := gstr.CaseCamelLower(table.TableName)

	buf.WriteString("package system\n\n")
	buf.WriteString("type Index" + entityName + "Req struct {\n")
	buf.WriteString("\tg.Meta `path:\"/" + varName + "/index\" method:\"get\" tags:\"" + table.MenuName + "\" summary:\"获取列表\"`\n")
	buf.WriteString("\tmodel.PageListReq\n")
	buf.WriteString("}\n\n")

	buf.WriteString("type Index" + entityName + "Res struct {\n")
	buf.WriteString("\tpage.PageRes\n")
	buf.WriteString("\tItems []res." + entityName + " `json:\"items\"`\n")
	buf.WriteString("}\n\n")

	buf.WriteString("type Read" + entityName + "Req struct {\n")
	buf.WriteString("\tId int64 `json:\"id\" v:\"required|min:1#Id不能为空\"`\n")
	buf.WriteString("}\n\n")

	buf.WriteString("type Read" + entityName + "Res struct {\n")
	buf.WriteString("\tData res." + entityName + " `json:\"data\"`\n")
	buf.WriteString("}\n")

	return buf.String()
}

// generateControllerCode 生成 Controller 代码
func (s *sCodeGen) generateControllerCode(table *entity.CodeGenTables, fields []entity.CodeGenFields) string {
	var buf bytes.Buffer
	entityName := gstr.CaseCamel(table.TableName)

	buf.WriteString("func (c *c" + entityName + ") Index(ctx context.Context, req *system.Index" + entityName + "Req) (res *system.Index" + entityName + "Res, err error) {\n")
	buf.WriteString("\tres = &system.Index" + entityName + "Res{}\n")
	buf.WriteString("\tres.Items, _, err = service." + entityName + "().GetPageList(ctx, &req.PageListReq)\n")
	buf.WriteString("\treturn\n")
	buf.WriteString("}\n")

	return buf.String()
}

// generateLogicCode 生成 Logic 代码
func (s *sCodeGen) generateLogicCode(table *entity.CodeGenTables, fields []entity.CodeGenFields) string {
	var buf bytes.Buffer
	entityName := gstr.CaseCamel(table.TableName)

	buf.WriteString("func (s *s" + entityName + ") GetPageList(ctx context.Context, req *model.PageListReq) (rs []*res." + entityName + ", total int, err error) {\n")
	buf.WriteString("\t// TODO: implement\n")
	buf.WriteString("\treturn\n")
	buf.WriteString("}\n")

	return buf.String()
}

// ReadTable 读取表信息
func (s *sCodeGen) ReadTable(ctx context.Context, id int64) (tableInfo res.CodeGenReadTable, err error) {
	var table entity.CodeGenTables
	err = s.Model(ctx).Where("id", id).Scan(&table)
	if utils.IsError(err) {
		return
	}

	tableInfo.TableName = table.TableName
	tableInfo.TableComment = table.TableComment

	// 获取字段
	var fields []entity.CodeGenFields
	err = dao.CodeGenFields.Ctx(ctx).Where("table_id", id).OrderAsc("sort").Scan(&fields)
	if utils.IsError(err) {
		return
	}

	for _, f := range fields {
		tableInfo.Columns = append(tableInfo.Columns, res.CodeGenColumn{
			ColumnName:    f.ColumnName,
			ColumnComment: f.ColumnComment,
			ColumnType:    f.ColumnType,
			IsNullable:    f.IsNullable,
			DataType:      f.DataType,
		})
	}

	return
}

// ListSourceTables 获取数据源表列表
func (s *sCodeGen) ListSourceTables(ctx context.Context, source string) (tables []res.CodeGenSourceTable, err error) {
	db := g.DB()
	sql := `SELECT table_name, table_comment
		FROM information_schema.tables
		WHERE table_schema = current_schema() AND table_type = 'BASE TABLE'
		AND table_name NOT LIKE 'code_gen_%'
		AND table_name NOT LIKE 'system_%'
		ORDER BY table_name`
	rows, err := db.GetCore().Query(ctx, sql)
	if utils.IsError(err) {
		return nil, err
	}
	for _, row := range rows {
		t := res.CodeGenSourceTable{
			Name:    row["table_name"].String(),
			Comment: row["table_comment"].String(),
		}
		tables = append(tables, t)
	}
	return
}

// generateModelCode 生成 Model 代码
func (s *sCodeGen) generateModelCode(table *entity.CodeGenTables, fields []entity.CodeGenFields) string {
	var buf bytes.Buffer
	entityName := gstr.CaseCamel(table.TableName)

	buf.WriteString("package model\n\n")
	buf.WriteString("type " + entityName + "Req struct {\n")
	buf.WriteString("}\n\n")

	buf.WriteString("type " + entityName + "Res struct {\n")
	buf.WriteString("}\n")

	return buf.String()
}

// getButtonName 获取按钮名称
func getButtonName(code string) string {
	buttonNames := map[string]string{
		"save":   "保存",
		"update": "更新",
		"read":   "查看",
		"delete": "删除",
		"export": "导出",
		"import": "导入",
		"query":  "查询",
		"reset":  "重置",
	}
	if name, ok := buttonNames[code]; ok {
		return name
	}
	return code
}
