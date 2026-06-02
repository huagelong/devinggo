// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"database/sql"

	"devinggo/internal/model/entity"
	"devinggo/modules/system/model"
	"devinggo/modules/system/model/req"
	"devinggo/modules/system/model/res"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"
)

type (
	// ICodeGen defines the interface for code generation operations.
	ICodeGen interface {
		// Model returns the database Model for code generation operations.
		Model(ctx context.Context) *gdb.Model
		// GetPageList retrieves a paginated list of code generation tables.
		GetPageList(ctx context.Context, req *model.PageListReq, in *req.CodeGenSearch) (rs []*res.CodeGenTable, total int, err error)
		// Delete removes code generation records by IDs.
		Delete(ctx context.Context, ids []int64) (err error)
		// Update modifies code generation configuration.
		Update(ctx context.Context, in *req.CodeGenUpdate, userId int64) (err error)
		// LoadTable loads database tables for code generation.
		LoadTable(ctx context.Context, in *req.CodeGenLoadTable, userId int64) (err error)
		// SyncTable synchronizes table structure for code generation.
		SyncTable(ctx context.Context, id int64, userId int64) (err error)
		// GenerateCode generates source code from table definitions.
		GenerateCode(ctx context.Context, ids string) (fileBytes []byte, err error)
		// PreviewCode previews generated code for a table.
		PreviewCode(ctx context.Context, id int64) (preview []res.CodeGenPreview, err error)
		// ReadTable retrieves table information for code generation.
		ReadTable(ctx context.Context, id int64) (tableInfo res.CodeGenReadTable, err error)
		// ListSourceTables lists tables from a data source.
		ListSourceTables(ctx context.Context, source string) (tables []res.CodeGenSourceTable, err error)
	}
	// IDashboard defines the interface for dashboard statistics operations.
	IDashboard interface {
		// GetStatistics retrieves dashboard statistics data.
		GetStatistics(ctx context.Context) (statistics map[string]interface{}, err error)
		// GetLoginChart retrieves login chart data for the specified number of days.
		GetLoginChart(ctx context.Context, days int) (chartData map[string]interface{}, err error)
	}
	// IDataMaintain defines the interface for data maintenance operations.
	IDataMaintain interface {
		// GetPageListForSearch retrieves a paginated list of data maintenance records with search criteria.
		GetPageListForSearch(ctx context.Context, req *model.PageListReq, in *req.DataMaintainSearch) (rs []*res.DataMaintain, total int, err error)
		// GetColumnList retrieves column information for a table.
		GetColumnList(ctx context.Context, source string, tableName string) (rs map[string]*gdb.TableField, err error)
		// GetAllTableStatus retrieves status for all tables in a group.
		GetAllTableStatus(ctx context.Context, groupName string) (rs []*res.DataMaintain, err error)
	}
	// ILogin defines the interface for user authentication operations.
	ILogin interface {
		// Model returns the database Model for login operations.
		Model(ctx context.Context) *gdb.Model
		// Login authenticates a user with username and password.
		Login(ctx context.Context, username string, password string) (token string, expire int64, err error)
	}
	// ISettingConfig defines the interface for configuration management operations.
	ISettingConfig interface {
		// Model returns the database Model for configuration operations.
		Model(ctx context.Context) *gdb.Model
		// GetConfigByKey retrieves a configuration value by key.
		GetConfigByKey(ctx context.Context, key string, groupKey ...string) (rs string, err error)
		// GetList retrieves a list of configurations with search criteria.
		GetList(ctx context.Context, in *req.SettingConfigSearch) (out []*res.SettingConfig, err error)
		// SaveConfig creates a new configuration.
		SaveConfig(ctx context.Context, data *req.SettingConfigSave) (id int64, err error)
		// UpdateConfig modifies an existing configuration.
		UpdateConfig(ctx context.Context, data *req.SettingConfigUpdate) (err error)
		// DeleteConfig removes configurations by keys.
		DeleteConfig(ctx context.Context, ids []string) (err error)
	}
	// ISettingConfigGroup defines the interface for configuration group management operations.
	ISettingConfigGroup interface {
		// Model returns the database Model for configuration group operations.
		Model(ctx context.Context) *gdb.Model
		// GetList retrieves a list of configuration groups.
		GetList(ctx context.Context) (out []*res.SettingConfigGroup, err error)
		// SaveConfigGroup creates a new configuration group.
		SaveConfigGroup(ctx context.Context, data *req.SettingConfigGroupSave) (id int64, err error)
		// UpdateConfigGroup modifies an existing configuration group.
		UpdateConfigGroup(ctx context.Context, data *req.SettingConfigGroupUpdate) (err error)
		// DeleteConfigGroup removes a configuration group by ID.
		DeleteConfigGroup(ctx context.Context, id int64) (err error)
	}
	// ISettingCrontab defines the interface for scheduled task management operations.
	ISettingCrontab interface {
		// Model returns the database Model for scheduled task operations.
		Model(ctx context.Context) *gdb.Model
		// GetValidateCron retrieves validated cron expressions.
		GetValidateCron(ctx context.Context) (rs []*res.SettingCrontabOne, err error)
		// GetPageList retrieves a paginated list of scheduled tasks.
		GetPageList(ctx context.Context, req *model.PageListReq, in *req.SettingCrontabSearch) (rs []*res.SettingCrontab, total int, err error)
		// Save creates a new scheduled task.
		Save(ctx context.Context, in *req.SettingCrontabSave) (id int64, err error)
		// GetById retrieves a scheduled task by ID.
		GetById(ctx context.Context, id int64) (res *res.SettingCrontab, err error)
		// Run executes a scheduled task immediately.
		Run(ctx context.Context, id int64) (err error)
		// Update modifies an existing scheduled task.
		Update(ctx context.Context, in *req.SettingCrontabUpdate) (err error)
		// Delete removes scheduled tasks by IDs.
		Delete(ctx context.Context, ids []int64) (err error)
		// ChangeStatus changes the status of a scheduled task.
		ChangeStatus(ctx context.Context, id int64, status int) (err error)
	}
	// ISettingCrontabLog defines the interface for scheduled task log operations.
	ISettingCrontabLog interface {
		// Model returns the database Model for scheduled task log operations.
		Model(ctx context.Context) *gdb.Model
		// GetPageList retrieves a paginated list of scheduled task logs.
		GetPageList(ctx context.Context, req *model.PageListReq, in *req.SettingCrontabLogSearch) (rs []*res.SettingCrontabLog, total int, err error)
		// Delete removes scheduled task logs by IDs.
		Delete(ctx context.Context, ids []int64) (err error)
		// AddLog adds a log entry for a scheduled task execution.
		AddLog(ctx context.Context, id int64, status int, exceptionInfo string, startTime *gtime.Time, endTime *gtime.Time, output string) (err error)
	}
	// ISystemApi defines the interface for API management operations.
	ISystemApi interface {
		// Model returns the database Model for API operations.
		Model(ctx context.Context) *gdb.Model
		// GetPageListForSearch retrieves a paginated list of APIs with search criteria.
		GetPageListForSearch(ctx context.Context, req *model.PageListReq, in *req.SystemApiSearch) (rs []*res.SystemApi, total int, err error)
		// GetList retrieves a list of APIs with search criteria.
		GetList(ctx context.Context, in *req.SystemApiSearch) (out []*res.SystemApi, err error)
		// Save creates a new API.
		Save(ctx context.Context, in *req.SystemApiSave) (id int64, err error)
		// GetById retrieves an API by ID.
		GetById(ctx context.Context, id int64) (res *res.SystemApi, err error)
		// Update modifies an existing API.
		Update(ctx context.Context, in *req.SystemApiUpdate) (err error)
		// Delete soft deletes APIs by IDs.
		Delete(ctx context.Context, ids []int64) (err error)
		// RealDelete permanently removes APIs by IDs.
		RealDelete(ctx context.Context, ids []int64) (err error)
		// Recovery restores soft-deleted APIs.
		Recovery(ctx context.Context, ids []int64) (err error)
		// ChangeStatus changes the status of an API.
		ChangeStatus(ctx context.Context, id int64, status int) (err error)
	}
	// ISystemApiGroup defines the interface for API group management operations.
	ISystemApiGroup interface {
		// Model returns the database Model for API group operations.
		Model(ctx context.Context) *gdb.Model
		// GetPageListForSearch retrieves a paginated list of API groups with search criteria.
		GetPageListForSearch(ctx context.Context, req *model.PageListReq, in *req.SystemApiGroupSearch) (rs []*res.SystemApiGroup, total int, err error)
		// GetList retrieves a list of API groups with search criteria.
		GetList(ctx context.Context, in *req.SystemApiGroupSearch) (out []*res.SystemApiGroup, err error)
		// Save creates a new API group.
		Save(ctx context.Context, in *req.SystemApiGroupSave) (id int64, err error)
		// GetById retrieves an API group by ID.
		GetById(ctx context.Context, id int64) (res *res.SystemApiGroup, err error)
		// Update modifies an existing API group.
		Update(ctx context.Context, in *req.SystemApiGroupUpdate) (err error)
		// Delete soft deletes API groups by IDs.
		Delete(ctx context.Context, ids []int64) (err error)
		// RealDelete permanently removes API groups by IDs.
		RealDelete(ctx context.Context, ids []int64) (err error)
		// Recovery restores soft-deleted API groups.
		Recovery(ctx context.Context, ids []int64) (err error)
		// ChangeStatus changes the status of an API group.
		ChangeStatus(ctx context.Context, id int64, status int) (err error)
	}
	// ISystemApiLog defines the interface for API log operations.
	ISystemApiLog interface {
		// Model returns the database Model for API log operations.
		Model(ctx context.Context) *gdb.Model
		// GetPageListForSearch retrieves a paginated list of API logs with search criteria.
		GetPageListForSearch(ctx context.Context, req *model.PageListReq, in *req.SystemApiLogSearch) (rs []*res.SystemApiLog, total int, err error)
		// Push pushes API logs to the system.
		Push(ctx context.Context)
		// DeleteApiLog removes API logs by IDs.
		DeleteApiLog(ctx context.Context, ids []int64) (err error)
	}
	// ISystemApp defines the interface for application management operations.
	ISystemApp interface {
		// Model returns the database Model for application operations.
		Model(ctx context.Context) *gdb.Model
		// GetAppId generates a new application ID.
		GetAppId(ctx context.Context) (string, error)
		// GetAppSecret generates a new application secret.
		GetAppSecret(ctx context.Context) (string, error)
		// BindApp binds APIs to an application.
		BindApp(ctx context.Context, Id int64, ApiIds []int64) (err error)
		// GetPageListForSearch retrieves a paginated list of applications with search criteria.
		GetPageListForSearch(ctx context.Context, req *model.PageListReq, in *req.SystemAppSearch) (rs []*res.SystemApp, total int, err error)
		// GetApiList retrieves API IDs bound to an application.
		GetApiList(ctx context.Context, id int64) (out []int64, err error)
		// Save creates a new application.
		Save(ctx context.Context, in *req.SystemAppSave, userId int64) (id int64, err error)
		// GetById retrieves an application by ID.
		GetById(ctx context.Context, id int64) (res *res.SystemApp, err error)
		// Update modifies an existing application.
		Update(ctx context.Context, in *req.SystemAppUpdate) (err error)
		// Delete soft deletes applications by IDs.
		Delete(ctx context.Context, ids []int64) (err error)
		// RealDelete permanently removes applications by IDs.
		RealDelete(ctx context.Context, ids []int64) (err error)
		// Recovery restores soft-deleted applications.
		Recovery(ctx context.Context, ids []int64) (err error)
		// ChangeStatus changes the status of an application.
		ChangeStatus(ctx context.Context, id int64, status int) (err error)
		// Verify verifies an API request.
		Verify(r *ghttp.Request) (bool, error)
		// GetAccessToken retrieves an access token for API authentication.
		GetAccessToken(ctx context.Context) (token string, exp int64, err error)
		// GetSignature generates a signature for API authentication.
		GetSignature(appSecret string, params map[string]interface{}) string
		// VerifyPre performs pre-verification for an API request.
		VerifyPre(ctx context.Context, appId string, apiId int64) (check bool, app *entity.SystemApp, err error)
		// VerifyEasyMode performs simplified verification for an API request.
		VerifyEasyMode(ctx context.Context, appId string, apiId int64) (check bool, err error)
	}
	// ISystemAppApi defines the interface for application API binding operations.
	ISystemAppApi interface {
		// Model returns the database Model for application API operations.
		Model(ctx context.Context) *gdb.Model
	}
	// ISystemAppGroup defines the interface for application group management operations.
	ISystemAppGroup interface {
		// Model returns the database Model for application group operations.
		Model(ctx context.Context) *gdb.Model
		// GetPageListForSearch retrieves a paginated list of application groups with search criteria.
		GetPageListForSearch(ctx context.Context, req *model.PageListReq, in *req.SystemAppGroupSearch) (rs []*res.SystemAppGroup, total int, err error)
		// GetList retrieves a list of application groups with search criteria.
		GetList(ctx context.Context, in *req.SystemAppGroupSearch) (out []*res.SystemAppGroup, err error)
		// Save creates a new application group.
		Save(ctx context.Context, in *req.SystemAppGroupSave) (id int64, err error)
		// GetById retrieves an application group by ID.
		GetById(ctx context.Context, id int64) (res *res.SystemAppGroup, err error)
		// Update modifies an existing application group.
		Update(ctx context.Context, in *req.SystemAppGroupUpdate) (err error)
		// Delete soft deletes application groups by IDs.
		Delete(ctx context.Context, ids []int64) (err error)
		// RealDelete permanently removes application groups by IDs.
		RealDelete(ctx context.Context, ids []int64) (err error)
		// Recovery restores soft-deleted application groups.
		Recovery(ctx context.Context, ids []int64) (err error)
		// ChangeStatus changes the status of an application group.
		ChangeStatus(ctx context.Context, id int64, status int) (err error)
	}
	// ISystemDept defines the interface for department management operations.
	ISystemDept interface {
		// Model returns the database Model for department operations.
		Model(ctx context.Context) *gdb.Model
		// GetSelectTree retrieves a department tree for selection.
		GetSelectTree(ctx context.Context, userId int64) (tree []*res.SystemDeptTree, err error)
		// GetList retrieves a list of departments with search criteria.
		GetList(ctx context.Context, in *req.SystemDeptSearch) (out []*res.SystemDept, err error)
		// GetListTreeList retrieves a department list tree with search criteria.
		GetListTreeList(ctx context.Context, in *req.SystemDeptSearch) (tree []*res.SystemListDeptTree, err error)
		// GetRecycleTreeList retrieves a department tree from recycle bin.
		GetRecycleTreeList(ctx context.Context, in *req.SystemDeptSearch) (tree []*res.SystemListDeptTree, err error)
		// GetTreeList retrieves a department tree with search criteria.
		GetTreeList(ctx context.Context, in *req.SystemDeptSearch) (tree []*res.SystemDeptTree, err error)
		// Save creates a new department.
		Save(ctx context.Context, in *req.SystemDeptSave) (id int64, err error)
		// AddLeader adds a leader to a department.
		AddLeader(ctx context.Context, in *req.SystemDeptAddLeader) (err error)
		// DelLeader removes leaders from a department.
		DelLeader(ctx context.Context, id int64, userIds []int64) (err error)
		// Update modifies an existing department.
		Update(ctx context.Context, in *req.SystemDeptSave) (err error)
		// Delete soft deletes departments by IDs.
		Delete(ctx context.Context, ids []int64) (err error)
		// RealDelete permanently removes departments by IDs.
		RealDelete(ctx context.Context, ids []int64) (err error)
		// Recovery restores soft-deleted departments.
		Recovery(ctx context.Context, ids []int64) (err error)
		// ChangeStatus changes the status of a department.
		ChangeStatus(ctx context.Context, id int64, status int) (err error)
		// NumberOperation performs a number operation on a department field.
		NumberOperation(ctx context.Context, id int64, numberName string, numberValue int) (err error)
	}
	// ISystemDeptLeader defines the interface for department leader operations.
	ISystemDeptLeader interface {
		// Model returns the database Model for department leader operations.
		Model(ctx context.Context) *gdb.Model
		// GetPageList retrieves a paginated list of department leaders.
		GetPageList(ctx context.Context, req *model.PageListReq, search *req.SystemDeptLeaderSearch) (res []*res.SystemDeptLeaderInfo, total int, err error)
	}
	// ISystemDictData defines the interface for dictionary data operations.
	ISystemDictData interface {
		// Model returns the database Model for dictionary data operations.
		Model(ctx context.Context) *gdb.Model
		// GetList retrieves a list of dictionary data with search criteria.
		GetList(ctx context.Context, listReq *model.ListReq, in *req.SystemDictDataSearch) (out []*res.SystemDictData, err error)
		// GetPageList retrieves a paginated list of dictionary data with search criteria.
		GetPageList(ctx context.Context, req *model.PageListReq, in *req.SystemDictDataSearch) (rs []*res.SystemDictDataFull, total int, err error)
		// Save creates new dictionary data.
		Save(ctx context.Context, in *req.SystemDictDataSave) (id int64, err error)
		// GetById retrieves dictionary data by ID.
		GetById(ctx context.Context, id int64) (res *res.SystemDictDataFull, err error)
		// Update modifies existing dictionary data.
		Update(ctx context.Context, in *req.SystemDictDataUpdate) (err error)
		// Delete soft deletes dictionary data by IDs.
		Delete(ctx context.Context, ids []int64) (err error)
		// RealDelete permanently removes dictionary data by IDs.
		RealDelete(ctx context.Context, ids []int64) (err error)
		// Recovery restores soft-deleted dictionary data.
		Recovery(ctx context.Context, ids []int64) (err error)
		// ChangeStatus changes the status of dictionary data.
		ChangeStatus(ctx context.Context, id int64, status int) (err error)
		// NumberOperation performs a number operation on a dictionary data field.
		NumberOperation(ctx context.Context, id int64, numberName string, numberValue int) (err error)
	}
	// ISystemDictType defines the interface for dictionary type operations.
	ISystemDictType interface {
		// Model returns the database Model for dictionary type operations.
		Model(ctx context.Context) *gdb.Model
		// GetPageList retrieves a paginated list of dictionary types with search criteria.
		GetPageList(ctx context.Context, req *model.PageListReq, in *req.SystemDictTypeSearch) (rs []*res.SystemDictType, total int, err error)
		// GetList retrieves a list of dictionary types with search criteria.
		GetList(ctx context.Context, listReq *model.ListReq, in *req.SystemDictTypeSearch) (out []*res.SystemDictType, err error)
		// Save creates a new dictionary type.
		Save(ctx context.Context, in *req.SystemDictTypeSave) (id int64, err error)
		// Update modifies an existing dictionary type.
		Update(ctx context.Context, in *req.SystemDictTypeUpdate) (err error)
		// Delete soft deletes dictionary types by IDs.
		Delete(ctx context.Context, ids []int64) (err error)
		// RealDelete permanently removes dictionary types by IDs.
		RealDelete(ctx context.Context, ids []int64) (err error)
		// GetById retrieves a dictionary type by ID (provided by GenericService).
		GetById(ctx context.Context, id int64) (res *res.SystemDictType, err error)
		// ChangeStatus changes the status of a dictionary type (provided by GenericService).
		ChangeStatus(ctx context.Context, id int64, status int) (err error)
		// Recovery restores soft-deleted dictionary types (provided by GenericService).
		Recovery(ctx context.Context, ids []int64) (err error)
	}
	// ISystemLoginLog defines the interface for login log operations.
	ISystemLoginLog interface {
		// Model returns the database Model for login log operations.
		Model(ctx context.Context) *gdb.Model
		// GetPageList retrieves a paginated list of login logs by username.
		GetPageList(ctx context.Context, req *model.PageListReq, username string) (res []*res.SystemLoginLog, total int, err error)
		// Push pushes a login log entry.
		Push(ctx context.Context, username string, err error)
		// GetPageListForSearch retrieves a paginated list of login logs with search criteria.
		GetPageListForSearch(ctx context.Context, req *model.PageListReq, in *req.SystemLoginLogSearch) (rs []*res.SystemLoginLog, total int, err error)
		// DeleteLoginLog removes login logs by IDs.
		DeleteLoginLog(ctx context.Context, ids []int64) (err error)
	}
	// ISystemMenu defines the interface for menu management operations.
	ISystemMenu interface {
		// Model returns the database Model for menu operations.
		Model(ctx context.Context) *gdb.Model
		// GetRoutersByIds retrieves router configurations by menu IDs.
		GetRoutersByIds(ctx context.Context, menuIds []int64) (routes []*res.Router, err error)
		// GetSuperAdminRouters retrieves router configurations for super admin.
		GetSuperAdminRouters(ctx context.Context) (routes []*res.Router, err error)
		// GetMenuCode retrieves menu codes by menu IDs.
		GetMenuCode(ctx context.Context, menuIds []int64) (menuCodes []string, err error)
		// GetMenuByPermission retrieves a menu by permission code.
		GetMenuByPermission(ctx context.Context, permission string, menuIds ...[]int64) (systemMenuEntity *entity.SystemMenu, err error)
		// GetTreeList retrieves a menu tree with search criteria.
		GetTreeList(ctx context.Context, in *req.SystemMenuSearch) (tree []*res.SystemMenuTree, err error)
		// GetRecycleTreeList retrieves a menu tree from recycle bin.
		GetRecycleTreeList(ctx context.Context, in *req.SystemMenuSearch) (tree []*res.SystemMenuTree, err error)
		// GetSelectTree retrieves a menu tree for selection.
		GetSelectTree(ctx context.Context, userId int64, onlyMenu bool, scope bool) (routes []*res.SystemDeptSelectTree, err error)
		// Save creates a new menu.
		Save(ctx context.Context, in *req.SystemMenuSave) (id int64, err error)
		// Update modifies an existing menu.
		Update(ctx context.Context, in *req.SystemMenuSave) (err error)
		// Delete soft deletes menus by IDs.
		Delete(ctx context.Context, ids []int64) (err error)
		// RealDelete permanently removes menus by IDs.
		RealDelete(ctx context.Context, ids []int64) (err error)
		// Recovery restores soft-deleted menus.
		Recovery(ctx context.Context, ids []int64) (err error)
		// ChangeStatus changes the status of a menu.
		ChangeStatus(ctx context.Context, id int64, status int) (err error)
		// NumberOperation performs a number operation on a menu field.
		NumberOperation(ctx context.Context, id int64, numberName string, numberValue int) (err error)
	}
	// ISystemModules defines the interface for module management operations.
	ISystemModules interface {
		// Model returns the database Model for module operations.
		Model(ctx context.Context) *gdb.Model
		// GetList retrieves a list of modules with search criteria.
		GetList(ctx context.Context, inReq *model.ListReq, in *req.SystemModulesSearch) (out []*res.SystemModules, err error)
		// GetPageList retrieves a paginated list of modules with search criteria.
		GetPageList(ctx context.Context, req *model.PageListReq, in *req.SystemModulesSearch) (rs []*res.SystemModules, total int, err error)
		// Save creates a new module.
		Save(ctx context.Context, in *req.SystemModulesSave) (id int64, err error)
		// GetById retrieves a module by ID.
		GetById(ctx context.Context, id int64) (res *res.SystemModules, err error)
		// Update modifies an existing module.
		Update(ctx context.Context, in *req.SystemModulesUpdate) (err error)
		// Delete soft deletes modules by IDs.
		Delete(ctx context.Context, ids []int64) (err error)
		// RealDelete permanently removes modules by IDs.
		RealDelete(ctx context.Context, ids []int64) (err error)
		// Recovery restores soft-deleted modules.
		Recovery(ctx context.Context, ids []int64) (err error)
		// ChangeStatus changes the status of a module.
		ChangeStatus(ctx context.Context, id int64, status int) (err error)
	}
	// ISystemNotice defines the interface for notice management operations.
	ISystemNotice interface {
		// Model returns the database Model for notice operations.
		Model(ctx context.Context) *gdb.Model
		// GetPageList retrieves a paginated list of notices.
		GetPageList(ctx context.Context, req *model.PageListReq) (res []*res.SystemNotice, total int, err error)
		// GetPageListForSearch retrieves a paginated list of notices with search criteria.
		GetPageListForSearch(ctx context.Context, req *model.PageListReq, in *req.SystemNoticeSearch) (rs []*res.SystemNotice, total int, err error)
		// Save creates a new notice.
		Save(ctx context.Context, in *req.SystemNoticeSave, userId int64) (id int64, err error)
		// Update modifies an existing notice.
		Update(ctx context.Context, in *req.SystemNoticeUpdate) (err error)
		// Delete soft deletes notices by IDs.
		Delete(ctx context.Context, ids []int64) (err error)
		// RealDelete permanently removes notices by IDs.
		RealDelete(ctx context.Context, ids []int64) (err error)
		// GetById retrieves a notice by ID (provided by GenericService).
		GetById(ctx context.Context, id int64) (res *res.SystemNotice, err error)
		// Recovery restores soft-deleted notices (provided by GenericService).
		Recovery(ctx context.Context, ids []int64) (err error)
	}
	// ISystemOperLog defines the interface for operation log operations.
	ISystemOperLog interface {
		// Model returns the database Model for operation log operations.
		Model(ctx context.Context) *gdb.Model
		// GetPageList retrieves a paginated list of operation logs by username.
		GetPageList(ctx context.Context, req *model.PageListReq, username string) (res []*res.SystemOperLog, total int, err error)
		// Push pushes an operation log entry.
		Push(ctx context.Context)
		// GetPageListForSearch retrieves a paginated list of operation logs with search criteria.
		GetPageListForSearch(ctx context.Context, req *model.PageListReq, in *req.SystemOperLogSearch) (rs []*res.SystemOperLog, total int, err error)
		// DeleteOperLog removes operation logs by IDs.
		DeleteOperLog(ctx context.Context, ids []int64) (err error)
	}
	// ISystemPost defines the interface for post management operations.
	ISystemPost interface {
		// Model returns the database Model for post operations.
		Model(ctx context.Context) *gdb.Model
		// GetList retrieves a list of posts with search criteria.
		GetList(ctx context.Context, in *req.SystemPostSearch) (out []*res.SystemPost, err error)
		// GetPageList retrieves a paginated list of posts with search criteria.
		GetPageList(ctx context.Context, req *model.PageListReq, in *req.SystemPostSearch) (rs []*res.SystemPost, total int, err error)
		// Save creates a new post.
		Save(ctx context.Context, in *req.SystemPostSave) (id int64, err error)
		// Update modifies an existing post.
		Update(ctx context.Context, in *req.SystemPostSave) (err error)
		// GetById retrieves a post by ID (provided by GenericService).
		GetById(ctx context.Context, id int64) (res *res.SystemPost, err error)
		// ChangeStatus changes the status of a post (provided by GenericService).
		ChangeStatus(ctx context.Context, id int64, status int) (err error)
		// Recovery restores soft-deleted posts (provided by GenericService).
		Recovery(ctx context.Context, ids []int64) (err error)
		// Delete soft deletes posts by IDs (provided by GenericService).
		Delete(ctx context.Context, ids []int64) (err error)
		// RealDelete permanently removes posts by IDs (provided by GenericService).
		RealDelete(ctx context.Context, ids []int64) (err error)
		// NumberOperation performs a number operation on a post field (provided by GenericService).
		NumberOperation(ctx context.Context, id int64, numberName string, numberValue int) (err error)
	}
	// ISystemQueueMessage defines the interface for queue message operations.
	ISystemQueueMessage interface {
		// Model returns the database Model for queue message operations.
		Model(ctx context.Context) *gdb.Model
		// GetReceiveUserPageList retrieves a paginated list of users who received a message.
		GetReceiveUserPageList(ctx context.Context, req *model.PageListReq, messageId int64) (rs []*res.MessageReceiveUser, total int, err error)
		// GetPageList retrieves a paginated list of queue messages.
		GetPageList(ctx context.Context, req *model.PageListReq, userId int64, params *req.SystemQueueMessageSearch) (rs []*res.SystemQueueMessage, total int, err error)
		// DeletesRelated removes message associations by IDs.
		DeletesRelated(ctx context.Context, ids []int64, userId int64) (err error)
		// SendMessage sends a queue message.
		SendMessage(ctx context.Context, sendReq *req.SystemQueueMessagesSend, sendUserId int64, contentType string) (err error, messageId int64)
	}
	// ISystemQueueMessageReceive defines the interface for queue message receive operations.
	ISystemQueueMessageReceive interface {
		// Model returns the database Model for queue message receive operations.
		Model(ctx context.Context) *gdb.Model
		// UpdateReadStatus updates the read status of messages.
		UpdateReadStatus(ctx context.Context, ids []int64, userId int64, value int) (err error)
	}
	// ISystemRole defines the interface for role management operations.
	ISystemRole interface {
		// Model returns the database Model for role operations.
		Model(ctx context.Context) *gdb.Model
		// GetByIds retrieves roles by their IDs.
		GetByIds(ctx context.Context, ids []int64) (res []*entity.SystemRole, err error)
		// Verify checks if the current user has permission to access the requested resource.
		Verify(r *ghttp.Request) bool
		// GetList retrieves a list of roles with search criteria.
		GetList(ctx context.Context, in *req.SystemRoleSearch, filterAdminRole bool) (out []*res.SystemRole, err error)
		// GetPageList retrieves a paginated list of roles with search criteria.
		GetPageList(ctx context.Context, req *model.PageListReq, in *req.SystemRoleSearch, filterAdminRole bool) (rs []*res.SystemRole, total int, err error)
		// Save creates a new role.
		Save(ctx context.Context, in *req.SystemRoleSave) (id int64, err error)
		// GetSuperAdminId retrieves the ID of the super admin role.
		GetSuperAdminId(ctx context.Context) (id int64, err error)
		// Update modifies an existing role.
		Update(ctx context.Context, in *req.SystemRoleSave) (err error)
		// Delete soft deletes roles by IDs.
		Delete(ctx context.Context, ids []int64) (err error)
		// RealDelete permanently removes roles by IDs.
		RealDelete(ctx context.Context, ids []int64) (err error)
		// Recovery restores soft-deleted roles.
		Recovery(ctx context.Context, ids []int64) (err error)
		// GetMenuByRoleIds retrieves menu associations for the given role IDs.
		GetMenuByRoleIds(ctx context.Context, ids []int64) (out []*res.SystemRoleMenus, err error)
		// GetDeptByRole retrieves department associations for the given role IDs.
		GetDeptByRole(ctx context.Context, ids []int64) (out []*res.SystemRoleDepts, err error)
		// GetById retrieves a role by ID (provided by GenericService).
		GetById(ctx context.Context, id int64) (res *res.SystemRole, err error)
		// ChangeStatus changes the status of a role (provided by GenericService).
		ChangeStatus(ctx context.Context, id int64, status int) (err error)
		// NumberOperation performs a number operation on a role field (provided by GenericService).
		NumberOperation(ctx context.Context, id int64, numberName string, numberValue int) (err error)
	}
	// ISystemRoleDept defines the interface for role-department binding operations.
	ISystemRoleDept interface {
		// Model returns the database Model for role-department operations.
		Model(ctx context.Context) *gdb.Model
	}
	// ISystemRoleMenu defines the interface for role-menu binding operations.
	ISystemRoleMenu interface {
		// Model returns the database Model for role-menu operations.
		Model(ctx context.Context) *gdb.Model
		// GetMenuIdsByRoleIds retrieves menu IDs by role IDs.
		GetMenuIdsByRoleIds(ctx context.Context, roleIds []int64) (rmenuIds []int64, err error)
	}
	// ISystemUploadfile defines the interface for file upload operations.
	ISystemUploadfile interface {
		// Model returns the database Model for file upload operations.
		Model(ctx context.Context) *gdb.Model
		// GetPageList retrieves a paginated list of uploaded files.
		GetPageList(ctx context.Context, in *model.PageListReq, params *req.SystemUploadFileSearch) (out []*res.SystemUploadFile, total int, err error)
		// SaveDb saves file information to the database.
		SaveDb(ctx context.Context, in *res.SystemUploadFileRes, createdBy int64) (rs int64, err error)
		// GetByHash retrieves a file by its hash.
		GetByHash(ctx context.Context, hash string) (rs *res.SystemUploadFileRes, err error)
		// GetFileInfoById retrieves file information by ID.
		GetFileInfoById(ctx context.Context, id int64) (rs *res.SystemUploadFile, err error)
		// GetFileInfoByHash retrieves file information by hash.
		GetFileInfoByHash(ctx context.Context, hash string) (rs *res.SystemUploadFile, err error)
		// Delete soft deletes files by IDs.
		Delete(ctx context.Context, ids []int64) (err error)
		// RealDelete permanently removes files by IDs.
		RealDelete(ctx context.Context, ids []int64) (err error)
		// Recovery restores soft-deleted files.
		Recovery(ctx context.Context, ids []int64) (err error)
	}
	// ISystemUser defines the interface for user management operations.
	ISystemUser interface {
		// Model returns the database Model for user operations.
		Model(ctx context.Context) *gdb.Model
		// GetPageList retrieves a paginated list of users.
		GetPageList(ctx context.Context, req *model.PageListReq) (res []*res.SystemUser, total int, err error)
		// GetPageListForSearch retrieves a paginated list of users with search criteria.
		GetPageListForSearch(ctx context.Context, req *model.PageListReq, in *req.SystemUserSearch) (res []*res.SystemUser, total int, err error)
		// GetOnlineUserPageListForSearch retrieves a paginated list of online users with search criteria.
		GetOnlineUserPageListForSearch(ctx context.Context, req *model.PageListReq, in *req.SystemUserSearch) (res []*res.SystemUser, total int, err error)
		// GetExportList retrieves a list of users for export.
		GetExportList(ctx context.Context, req *model.ListReq, in *req.SystemUserSearch) (res []*res.SystemUserExport, err error)
		// GetSupserAdminId returns the super admin user ID.
		GetSupserAdminId(ctx context.Context) int64
		// ExistsByUsername checks if a user with the given username exists.
		ExistsByUsername(ctx context.Context, username string) (rs bool, err error)
		// GetInfoById retrieves basic user information by user ID.
		GetInfoById(ctx context.Context, userId int64) (systemUser *res.SystemUser, err error)
		// GetInfoByIds retrieves basic user information for multiple user IDs.
		GetInfoByIds(ctx context.Context, userIds []int64) (systemUser []*res.SystemUser, err error)
		// GetInfo retrieves comprehensive user information including roles, permissions, and routers.
		GetInfo(ctx context.Context, userId int64) (systemUserInfo *res.SystemUserInfo, err error)
		// IsSuperAdmin checks if the user has super administrator privileges.
		IsSuperAdmin(ctx context.Context, userId int64) (isSuperAdmin bool, err error)
		// GetRoles retrieves the role IDs assigned to a user.
		GetRoles(ctx context.Context, userId int64) (roles []int64, err error)
		// GetDepts retrieves the department IDs assigned to a user.
		GetDepts(ctx context.Context, userId int64) (depts []int64, err error)
		// Update modifies user information with optional user ID override.
		Update(ctx context.Context, req *req.SystemUser, userId ...int64) (rs sql.Result, err error)
		// SetHomePage sets the home page dashboard for a user.
		SetHomePage(ctx context.Context, id int64, dashboard string) (out sql.Result, err error)
		// InitUserPassword initializes or resets a user's password with hashing.
		InitUserPassword(ctx context.Context, id int64, password string) (out sql.Result, err error)
		// UpdateSimple performs a simplified user update including roles, departments, and posts.
		UpdateSimple(ctx context.Context, in *req.SystemUserUpdate) (out sql.Result, err error)
		// Save creates a new system user with associated roles, departments, and posts.
		Save(ctx context.Context, in *req.SystemUserSave) (id int64, err error)
		// GetFullInfoById retrieves complete user information including associated roles, posts, and departments.
		GetFullInfoById(ctx context.Context, id int64) (out *res.SystemUserFullInfo, err error)
		// Delete soft deletes system users, excluding the super admin.
		Delete(ctx context.Context, ids []int64) (err error)
		// RealDelete permanently removes system users and their associated data, excluding the super admin.
		RealDelete(ctx context.Context, ids []int64) (err error)
		// Recovery restores soft-deleted system users.
		Recovery(ctx context.Context, ids []int64) (err error)
		// ChangeStatus changes the status of a system user.
		ChangeStatus(ctx context.Context, id int64, status int) (err error)
	}
	// ISystemUserDept defines the interface for user-department binding operations.
	ISystemUserDept interface {
		// Model returns the database Model for user-department operations.
		Model(ctx context.Context) *gdb.Model
		// GetDeptIdsByUserId retrieves department IDs by user ID.
		GetDeptIdsByUserId(ctx context.Context, userId int64) (deptIds []int64, err error)
	}
	// ISystemUserPost defines the interface for user-post binding operations.
	ISystemUserPost interface {
		// Model returns the database Model for user-post operations.
		Model(ctx context.Context) *gdb.Model
		// GetPostIdsByUserId retrieves post IDs by user ID.
		GetPostIdsByUserId(ctx context.Context, userId int64) (postIds []int64, err error)
	}
	// ISystemUserRole defines the interface for user-role binding operations.
	ISystemUserRole interface {
		// Model returns the database Model for user-role operations.
		Model(ctx context.Context) *gdb.Model
		// GetRoleIdsByUserId retrieves role IDs by user ID.
		GetRoleIdsByUserId(ctx context.Context, userId int64) (roleIds []int64, err error)
	}
)

var (
	localCodeGen                   ICodeGen
	localDashboard                 IDashboard
	localDataMaintain              IDataMaintain
	localLogin                     ILogin
	localSettingConfig             ISettingConfig
	localSettingConfigGroup        ISettingConfigGroup
	localSettingCrontab            ISettingCrontab
	localSettingCrontabLog         ISettingCrontabLog
	localSystemApi                 ISystemApi
	localSystemApiGroup            ISystemApiGroup
	localSystemApiLog              ISystemApiLog
	localSystemApp                 ISystemApp
	localSystemAppApi              ISystemAppApi
	localSystemAppGroup            ISystemAppGroup
	localSystemDept                ISystemDept
	localSystemDeptLeader          ISystemDeptLeader
	localSystemDictData            ISystemDictData
	localSystemDictType            ISystemDictType
	localSystemLoginLog            ISystemLoginLog
	localSystemMenu                ISystemMenu
	localSystemModules             ISystemModules
	localSystemNotice              ISystemNotice
	localSystemOperLog             ISystemOperLog
	localSystemPost                ISystemPost
	localSystemQueueMessage        ISystemQueueMessage
	localSystemQueueMessageReceive ISystemQueueMessageReceive
	localSystemRole                ISystemRole
	localSystemRoleDept            ISystemRoleDept
	localSystemRoleMenu            ISystemRoleMenu
	localSystemUploadfile          ISystemUploadfile
	localSystemUser                ISystemUser
	localSystemUserDept            ISystemUserDept
	localSystemUserPost            ISystemUserPost
	localSystemUserRole            ISystemUserRole
)

func CodeGen() ICodeGen {
	if localCodeGen == nil {
		panic("implement not found for interface ICodeGen, forgot register?")
	}
	return localCodeGen
}

func RegisterCodeGen(i ICodeGen) {
	localCodeGen = i
}

func Dashboard() IDashboard {
	if localDashboard == nil {
		panic("implement not found for interface IDashboard, forgot register?")
	}
	return localDashboard
}

func RegisterDashboard(i IDashboard) {
	localDashboard = i
}

func DataMaintain() IDataMaintain {
	if localDataMaintain == nil {
		panic("implement not found for interface IDataMaintain, forgot register?")
	}
	return localDataMaintain
}

func RegisterDataMaintain(i IDataMaintain) {
	localDataMaintain = i
}

func Login() ILogin {
	if localLogin == nil {
		panic("implement not found for interface ILogin, forgot register?")
	}
	return localLogin
}

func RegisterLogin(i ILogin) {
	localLogin = i
}

func SettingConfig() ISettingConfig {
	if localSettingConfig == nil {
		panic("implement not found for interface ISettingConfig, forgot register?")
	}
	return localSettingConfig
}

func RegisterSettingConfig(i ISettingConfig) {
	localSettingConfig = i
}

func SettingConfigGroup() ISettingConfigGroup {
	if localSettingConfigGroup == nil {
		panic("implement not found for interface ISettingConfigGroup, forgot register?")
	}
	return localSettingConfigGroup
}

func RegisterSettingConfigGroup(i ISettingConfigGroup) {
	localSettingConfigGroup = i
}

func SettingCrontab() ISettingCrontab {
	if localSettingCrontab == nil {
		panic("implement not found for interface ISettingCrontab, forgot register?")
	}
	return localSettingCrontab
}

func RegisterSettingCrontab(i ISettingCrontab) {
	localSettingCrontab = i
}

func SettingCrontabLog() ISettingCrontabLog {
	if localSettingCrontabLog == nil {
		panic("implement not found for interface ISettingCrontabLog, forgot register?")
	}
	return localSettingCrontabLog
}

func RegisterSettingCrontabLog(i ISettingCrontabLog) {
	localSettingCrontabLog = i
}

func SystemApi() ISystemApi {
	if localSystemApi == nil {
		panic("implement not found for interface ISystemApi, forgot register?")
	}
	return localSystemApi
}

func RegisterSystemApi(i ISystemApi) {
	localSystemApi = i
}

func SystemApiGroup() ISystemApiGroup {
	if localSystemApiGroup == nil {
		panic("implement not found for interface ISystemApiGroup, forgot register?")
	}
	return localSystemApiGroup
}

func RegisterSystemApiGroup(i ISystemApiGroup) {
	localSystemApiGroup = i
}

func SystemApiLog() ISystemApiLog {
	if localSystemApiLog == nil {
		panic("implement not found for interface ISystemApiLog, forgot register?")
	}
	return localSystemApiLog
}

func RegisterSystemApiLog(i ISystemApiLog) {
	localSystemApiLog = i
}

func SystemApp() ISystemApp {
	if localSystemApp == nil {
		panic("implement not found for interface ISystemApp, forgot register?")
	}
	return localSystemApp
}

func RegisterSystemApp(i ISystemApp) {
	localSystemApp = i
}

func SystemAppApi() ISystemAppApi {
	if localSystemAppApi == nil {
		panic("implement not found for interface ISystemAppApi, forgot register?")
	}
	return localSystemAppApi
}

func RegisterSystemAppApi(i ISystemAppApi) {
	localSystemAppApi = i
}

func SystemAppGroup() ISystemAppGroup {
	if localSystemAppGroup == nil {
		panic("implement not found for interface ISystemAppGroup, forgot register?")
	}
	return localSystemAppGroup
}

func RegisterSystemAppGroup(i ISystemAppGroup) {
	localSystemAppGroup = i
}

func SystemDept() ISystemDept {
	if localSystemDept == nil {
		panic("implement not found for interface ISystemDept, forgot register?")
	}
	return localSystemDept
}

func RegisterSystemDept(i ISystemDept) {
	localSystemDept = i
}

func SystemDeptLeader() ISystemDeptLeader {
	if localSystemDeptLeader == nil {
		panic("implement not found for interface ISystemDeptLeader, forgot register?")
	}
	return localSystemDeptLeader
}

func RegisterSystemDeptLeader(i ISystemDeptLeader) {
	localSystemDeptLeader = i
}

func SystemDictData() ISystemDictData {
	if localSystemDictData == nil {
		panic("implement not found for interface ISystemDictData, forgot register?")
	}
	return localSystemDictData
}

func RegisterSystemDictData(i ISystemDictData) {
	localSystemDictData = i
}

func SystemDictType() ISystemDictType {
	if localSystemDictType == nil {
		panic("implement not found for interface ISystemDictType, forgot register?")
	}
	return localSystemDictType
}

func RegisterSystemDictType(i ISystemDictType) {
	localSystemDictType = i
}

func SystemLoginLog() ISystemLoginLog {
	if localSystemLoginLog == nil {
		panic("implement not found for interface ISystemLoginLog, forgot register?")
	}
	return localSystemLoginLog
}

func RegisterSystemLoginLog(i ISystemLoginLog) {
	localSystemLoginLog = i
}

func SystemMenu() ISystemMenu {
	if localSystemMenu == nil {
		panic("implement not found for interface ISystemMenu, forgot register?")
	}
	return localSystemMenu
}

func RegisterSystemMenu(i ISystemMenu) {
	localSystemMenu = i
}

func SystemModules() ISystemModules {
	if localSystemModules == nil {
		panic("implement not found for interface ISystemModules, forgot register?")
	}
	return localSystemModules
}

func RegisterSystemModules(i ISystemModules) {
	localSystemModules = i
}

func SystemNotice() ISystemNotice {
	if localSystemNotice == nil {
		panic("implement not found for interface ISystemNotice, forgot register?")
	}
	return localSystemNotice
}

func RegisterSystemNotice(i ISystemNotice) {
	localSystemNotice = i
}

func SystemOperLog() ISystemOperLog {
	if localSystemOperLog == nil {
		panic("implement not found for interface ISystemOperLog, forgot register?")
	}
	return localSystemOperLog
}

func RegisterSystemOperLog(i ISystemOperLog) {
	localSystemOperLog = i
}

func SystemPost() ISystemPost {
	if localSystemPost == nil {
		panic("implement not found for interface ISystemPost, forgot register?")
	}
	return localSystemPost
}

func RegisterSystemPost(i ISystemPost) {
	localSystemPost = i
}

func SystemQueueMessage() ISystemQueueMessage {
	if localSystemQueueMessage == nil {
		panic("implement not found for interface ISystemQueueMessage, forgot register?")
	}
	return localSystemQueueMessage
}

func RegisterSystemQueueMessage(i ISystemQueueMessage) {
	localSystemQueueMessage = i
}

func SystemQueueMessageReceive() ISystemQueueMessageReceive {
	if localSystemQueueMessageReceive == nil {
		panic("implement not found for interface ISystemQueueMessageReceive, forgot register?")
	}
	return localSystemQueueMessageReceive
}

func RegisterSystemQueueMessageReceive(i ISystemQueueMessageReceive) {
	localSystemQueueMessageReceive = i
}

func SystemRole() ISystemRole {
	if localSystemRole == nil {
		panic("implement not found for interface ISystemRole, forgot register?")
	}
	return localSystemRole
}

func RegisterSystemRole(i ISystemRole) {
	localSystemRole = i
}

func SystemRoleDept() ISystemRoleDept {
	if localSystemRoleDept == nil {
		panic("implement not found for interface ISystemRoleDept, forgot register?")
	}
	return localSystemRoleDept
}

func RegisterSystemRoleDept(i ISystemRoleDept) {
	localSystemRoleDept = i
}

func SystemRoleMenu() ISystemRoleMenu {
	if localSystemRoleMenu == nil {
		panic("implement not found for interface ISystemRoleMenu, forgot register?")
	}
	return localSystemRoleMenu
}

func RegisterSystemRoleMenu(i ISystemRoleMenu) {
	localSystemRoleMenu = i
}

func SystemUploadfile() ISystemUploadfile {
	if localSystemUploadfile == nil {
		panic("implement not found for interface ISystemUploadfile, forgot register?")
	}
	return localSystemUploadfile
}

func RegisterSystemUploadfile(i ISystemUploadfile) {
	localSystemUploadfile = i
}

func SystemUser() ISystemUser {
	if localSystemUser == nil {
		panic("implement not found for interface ISystemUser, forgot register?")
	}
	return localSystemUser
}

func RegisterSystemUser(i ISystemUser) {
	localSystemUser = i
}

func SystemUserDept() ISystemUserDept {
	if localSystemUserDept == nil {
		panic("implement not found for interface ISystemUserDept, forgot register?")
	}
	return localSystemUserDept
}

func RegisterSystemUserDept(i ISystemUserDept) {
	localSystemUserDept = i
}

func SystemUserPost() ISystemUserPost {
	if localSystemUserPost == nil {
		panic("implement not found for interface ISystemUserPost, forgot register?")
	}
	return localSystemUserPost
}

func RegisterSystemUserPost(i ISystemUserPost) {
	localSystemUserPost = i
}

func SystemUserRole() ISystemUserRole {
	if localSystemUserRole == nil {
		panic("implement not found for interface ISystemUserRole, forgot register?")
	}
	return localSystemUserRole
}

func RegisterSystemUserRole(i ISystemUserRole) {
	localSystemUserRole = i
}
