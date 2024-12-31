// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"devinggo/internal/model"
	"devinggo/internal/model/entity"
	"devinggo/modules/system/model/req"
	"devinggo/modules/system/model/res"
	"context"
	"database/sql"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/net/ghttp"
)

type (
	IDataMaintain interface {
		GetPageListForSearch(ctx context.Context, req *model.PageListReq, in *req.DataMaintainSearch) (rs []*res.DataMaintain, total int, err error)
		GetColumnList(ctx context.Context, source, tableName string) (rs map[string]*gdb.TableField, err error)
		GetAllTableStatus(ctx context.Context, groupName string) (rs []*res.DataMaintain, err error)
	}
	ILogin interface {
		Model(ctx context.Context) *gdb.Model
		Login(ctx context.Context, username, password string) (token string, expire int64, err error)
	}
	ISettingConfig interface {
		Model(ctx context.Context) *gdb.Model
		GetConfigByKey(ctx context.Context, key string, groupKey ...string) (rs string, err error)
		GetList(ctx context.Context, in *req.SettingConfigSearch) (out []*res.SettingConfig, err error)
		SaveConfig(ctx context.Context, data *req.SettingConfigSave) (id uint64, err error)
		UpdateConfig(ctx context.Context, data *req.SettingConfigUpdate) (err error)
		DeleteConfig(ctx context.Context, ids []string) (err error)
	}
	ISettingConfigGroup interface {
		Model(ctx context.Context) *gdb.Model
		GetList(ctx context.Context) (out []*res.SettingConfigGroup, err error)
		SaveConfigGroup(ctx context.Context, data *req.SettingConfigGroupSave) (id uint64, err error)
		UpdateConfigGroup(ctx context.Context, data *req.SettingConfigGroupUpdate) (err error)
		DeleteConfigGroup(ctx context.Context, id uint64) (err error)
	}
	ISettingCrontab interface {
		Model(ctx context.Context) *gdb.Model
		GetValidateCron(ctx context.Context) (rs []*res.SettingCrontabOne, err error)
		GetPageList(ctx context.Context, req *model.PageListReq, in *req.SettingCrontabSearch) (rs []*res.SettingCrontab, total int, err error)
		Save(ctx context.Context, in *req.SettingCrontabSave) (id uint64, err error)
		GetById(ctx context.Context, id uint64) (res *res.SettingCrontab, err error)
		Run(ctx context.Context, id uint64) (err error)
		Update(ctx context.Context, in *req.SettingCrontabUpdate) (err error)
		Delete(ctx context.Context, ids []uint64) (err error)
		ChangeStatus(ctx context.Context, id uint64, status int) (err error)
	}
	ISettingCrontabLog interface {
		Model(ctx context.Context) *gdb.Model
		GetPageList(ctx context.Context, req *model.PageListReq, in *req.SettingCrontabLogSearch) (rs []*res.SettingCrontabLog, total int, err error)
		Delete(ctx context.Context, ids []uint64) (err error)
		AddLog(ctx context.Context, id uint64, status int, exceptionInfo string) (err error)
	}
	ISettingGenerateColumns interface {
		Model(ctx context.Context) *gdb.Model
		GetList(ctx context.Context, in *req.SettingGenerateColumnsSearch) (out []*res.SettingGenerateColumns, err error)
	}
	ISettingGenerateTables interface {
		Model(ctx context.Context) *gdb.Model
		GetPageListForSearch(ctx context.Context, req *model.PageListReq, in *req.SettingGenerateTablesSearch) (rs []*res.SettingGenerateTables, total int, err error)
		LoadTable(ctx context.Context, in *req.LoadTable) (err error)
		GetById(ctx context.Context, id uint64) (res *res.SettingGenerateTables, err error)
		Delete(ctx context.Context, ids []uint64) (err error)
		SyncCode(ctx context.Context, id uint64) (err error)
		UpdateTableAndColumns(ctx context.Context, in *req.TableAndColumnsUpdate) (err error)
		GenerateCode(ctx context.Context, ids []uint64) (filePath string, err error)
		Preview(ctx context.Context, id uint64) (rs []res.PreviewTable, err error)
	}
	ISystemApi interface {
		Model(ctx context.Context) *gdb.Model
		GetPageListForSearch(ctx context.Context, req *model.PageListReq, in *req.SystemApiSearch) (rs []*res.SystemApi, total int, err error)
		GetList(ctx context.Context, in *req.SystemApiSearch) (out []*res.SystemApi, err error)
		Save(ctx context.Context, in *req.SystemApiSave) (id uint64, err error)
		GetById(ctx context.Context, id uint64) (res *res.SystemApi, err error)
		Update(ctx context.Context, in *req.SystemApiUpdate) (err error)
		Delete(ctx context.Context, ids []uint64) (err error)
		RealDelete(ctx context.Context, ids []uint64) (err error)
		Recovery(ctx context.Context, ids []uint64) (err error)
		ChangeStatus(ctx context.Context, id uint64, status int) (err error)
	}
	ISystemApiGroup interface {
		Model(ctx context.Context) *gdb.Model
		GetPageListForSearch(ctx context.Context, req *model.PageListReq, in *req.SystemApiGroupSearch) (rs []*res.SystemApiGroup, total int, err error)
		GetList(ctx context.Context, in *req.SystemApiGroupSearch) (out []*res.SystemApiGroup, err error)
		Save(ctx context.Context, in *req.SystemApiGroupSave) (id uint64, err error)
		GetById(ctx context.Context, id uint64) (res *res.SystemApiGroup, err error)
		Update(ctx context.Context, in *req.SystemApiGroupUpdate) (err error)
		Delete(ctx context.Context, ids []uint64) (err error)
		RealDelete(ctx context.Context, ids []uint64) (err error)
		Recovery(ctx context.Context, ids []uint64) (err error)
		ChangeStatus(ctx context.Context, id uint64, status int) (err error)
	}
	ISystemApiLog interface {
		Model(ctx context.Context) *gdb.Model
		GetPageListForSearch(ctx context.Context, req *model.PageListReq, in *req.SystemApiLogSearch) (rs []*res.SystemApiLog, total int, err error)
		Push(ctx context.Context)
	}
	ISystemApp interface {
		Model(ctx context.Context) *gdb.Model
		GetAppId(ctx context.Context) (string, error)
		GetAppSecret(ctx context.Context) (string, error)
		BindApp(ctx context.Context, Id uint64, ApiIds []uint64) (err error)
		GetPageListForSearch(ctx context.Context, req *model.PageListReq, in *req.SystemAppSearch) (rs []*res.SystemApp, total int, err error)
		GetApiList(ctx context.Context, id uint64) (out []uint64, err error)
		Save(ctx context.Context, in *req.SystemAppSave, userId uint64) (id uint64, err error)
		GetById(ctx context.Context, id uint64) (res *res.SystemApp, err error)
		Update(ctx context.Context, in *req.SystemAppUpdate) (err error)
		Delete(ctx context.Context, ids []uint64) (err error)
		RealDelete(ctx context.Context, ids []uint64) (err error)
		Recovery(ctx context.Context, ids []uint64) (err error)
		ChangeStatus(ctx context.Context, id uint64, status int) (err error)
		Verify(r *ghttp.Request) (bool, error)
		// getAccessToken 获取access_token
		GetAccessToken(ctx context.Context, params map[string]interface{}) (token string, exp int64, err error)
		// getSignature 获取签名
		GetSignature(appSecret string, params map[string]interface{}) string
		VerifyPre(ctx context.Context, appId string, apiId uint64) (check bool, app *entity.SystemApp, err error)
		// 简单模式
		VerifyEasyMode(ctx context.Context, appId string, signature string, apiId uint64) (check bool, err error)
	}
	ISystemAppApi interface {
		Model(ctx context.Context) *gdb.Model
	}
	ISystemAppGroup interface {
		Model(ctx context.Context) *gdb.Model
		GetPageListForSearch(ctx context.Context, req *model.PageListReq, in *req.SystemAppGroupSearch) (rs []*res.SystemAppGroup, total int, err error)
		GetList(ctx context.Context, in *req.SystemAppGroupSearch) (out []*res.SystemAppGroup, err error)
		Save(ctx context.Context, in *req.SystemAppGroupSave) (id uint64, err error)
		GetById(ctx context.Context, id uint64) (res *res.SystemAppGroup, err error)
		Update(ctx context.Context, in *req.SystemAppGroupUpdate) (err error)
		Delete(ctx context.Context, ids []uint64) (err error)
		RealDelete(ctx context.Context, ids []uint64) (err error)
		Recovery(ctx context.Context, ids []uint64) (err error)
		ChangeStatus(ctx context.Context, id uint64, status int) (err error)
	}
	ISystemDept interface {
		Model(ctx context.Context) *gdb.Model
		GetSelectTree(ctx context.Context, userId uint64) (tree []*res.SystemDeptTree, err error)
		GetListTreeList(ctx context.Context, in *req.SystemDeptSearch) (tree []*res.SystemListDeptTree, err error)
		GetRecycleTreeList(ctx context.Context, in *req.SystemDeptSearch) (tree []*res.SystemListDeptTree, err error)
		GetTreeList(ctx context.Context, in *req.SystemDeptSearch) (tree []*res.SystemDeptTree, err error)
		Save(ctx context.Context, in *req.SystemDeptSave) (id uint64, err error)
		AddLeader(ctx context.Context, in *req.SystemDeptAddLeader) (err error)
		DelLeader(ctx context.Context, id uint64, userIds []uint64) (err error)
		Update(ctx context.Context, in *req.SystemDeptSave) (err error)
		Delete(ctx context.Context, ids []uint64) (names []string, err error)
		RealDelete(ctx context.Context, ids []uint64) (names []string, err error)
		Recovery(ctx context.Context, ids []uint64) (err error)
		ChangeStatus(ctx context.Context, id uint64, status int) (err error)
		NumberOperation(ctx context.Context, id uint64, numberName string, numberValue int) (err error)
	}
	ISystemDeptLeader interface {
		Model(ctx context.Context) *gdb.Model
		GetPageList(ctx context.Context, req *model.PageListReq, search *req.SystemDeptLeaderSearch) (res []*res.SystemDeptLeaderInfo, total int, err error)
	}
	ISystemDictData interface {
		Model(ctx context.Context) *gdb.Model
		GetList(ctx context.Context, listReq *model.ListReq, in *req.SystemDictDataSearch) (out []*res.SystemDictData, err error)
		GetPageList(ctx context.Context, req *model.PageListReq, in *req.SystemDictDataSearch) (rs []*res.SystemDictDataFull, total int, err error)
		Save(ctx context.Context, in *req.SystemDictDataSave) (id uint64, err error)
		GetById(ctx context.Context, id uint64) (res *res.SystemDictDataFull, err error)
		Update(ctx context.Context, in *req.SystemDictDataUpdate) (err error)
		Delete(ctx context.Context, ids []uint64) (err error)
		RealDelete(ctx context.Context, ids []uint64) (err error)
		Recovery(ctx context.Context, ids []uint64) (err error)
		ChangeStatus(ctx context.Context, id uint64, status int) (err error)
		NumberOperation(ctx context.Context, id uint64, numberName string, numberValue int) (err error)
	}
	ISystemDictType interface {
		Model(ctx context.Context) *gdb.Model
		GetPageList(ctx context.Context, req *model.PageListReq, in *req.SystemDictTypeSearch) (rs []*res.SystemDictType, total int, err error)
		GetList(ctx context.Context, listReq *model.ListReq, in *req.SystemDictTypeSearch) (out []*res.SystemDictType, err error)
		Save(ctx context.Context, in *req.SystemDictTypeSave) (id uint64, err error)
		GetById(ctx context.Context, id uint64) (res *res.SystemDictType, err error)
		Update(ctx context.Context, in *req.SystemDictTypeUpdate) (err error)
		Delete(ctx context.Context, ids []uint64) (err error)
		RealDelete(ctx context.Context, ids []uint64) (err error)
		Recovery(ctx context.Context, ids []uint64) (err error)
		ChangeStatus(ctx context.Context, id uint64, status int) (err error)
	}
	ISystemLoginLog interface {
		Model(ctx context.Context) *gdb.Model
		GetPageList(ctx context.Context, req *model.PageListReq, username string) (res []*res.SystemLoginLog, total int, err error)
		Push(ctx context.Context, username string, err error)
		GetPageListForSearch(ctx context.Context, req *model.PageListReq, in *req.SystemLoginLogSearch) (rs []*res.SystemLoginLog, total int, err error)
	}
	ISystemMenu interface {
		Model(ctx context.Context) *gdb.Model
		GetRoutersByIds(ctx context.Context, menuIds []uint64) (routes []*res.Router, err error)
		GetSuperAdminRouters(ctx context.Context) (routes []*res.Router, err error)
		GetMenuCode(ctx context.Context, menuIds []uint64) (menuCodes []string, err error)
		GetMenuByPermission(ctx context.Context, permission string, menuIds ...[]uint64) (systemMenuEntity *entity.SystemMenu, err error)
		GetTreeList(ctx context.Context, in *req.SystemMenuSearch) (tree []*res.SystemMenuTree, err error)
		GetRecycleTreeList(ctx context.Context, in *req.SystemMenuSearch) (tree []*res.SystemMenuTree, err error)
		GetSelectTree(ctx context.Context, userId uint64, onlyMenu, scope bool) (routes []*res.SystemDeptSelectTree, err error)
		Save(ctx context.Context, in *req.SystemMenuSave) (id uint64, err error)
		Update(ctx context.Context, in *req.SystemMenuSave) (err error)
		Delete(ctx context.Context, ids []uint64) (names []string, err error)
		RealDelete(ctx context.Context, ids []uint64) (names []string, err error)
		Recovery(ctx context.Context, ids []uint64) (err error)
		ChangeStatus(ctx context.Context, id uint64, status int) (err error)
		NumberOperation(ctx context.Context, id uint64, numberName string, numberValue int) (err error)
	}
	ISystemModules interface {
		Model(ctx context.Context) *gdb.Model
		GetList(ctx context.Context, inReq *model.ListReq, in *req.SystemModulesSearch) (out []*res.SystemModules, err error)
		GetPageList(ctx context.Context, req *model.PageListReq, in *req.SystemModulesSearch) (rs []*res.SystemModules, total int, err error)
		Save(ctx context.Context, in *req.SystemModulesSave) (id uint64, err error)
		GetById(ctx context.Context, id uint64) (res *res.SystemModules, err error)
		Update(ctx context.Context, in *req.SystemModulesUpdate) (err error)
		Delete(ctx context.Context, ids []uint64) (err error)
		RealDelete(ctx context.Context, ids []uint64) (err error)
		Recovery(ctx context.Context, ids []uint64) (err error)
		ChangeStatus(ctx context.Context, id uint64, status int) (err error)
	}
	ISystemNotice interface {
		Model(ctx context.Context) *gdb.Model
		GetPageList(ctx context.Context, req *model.PageListReq) (res []*res.SystemNotice, total int, err error)
		GetPageListForSearch(ctx context.Context, req *model.PageListReq, in *req.SystemNoticeSearch) (rs []*res.SystemNotice, total int, err error)
		Save(ctx context.Context, in *req.SystemNoticeSave, userId uint64) (id uint64, err error)
		GetById(ctx context.Context, id uint64) (res *res.SystemNotice, err error)
		Update(ctx context.Context, in *req.SystemNoticeUpdate) (err error)
		Delete(ctx context.Context, ids []uint64) (err error)
		RealDelete(ctx context.Context, ids []uint64) (err error)
		Recovery(ctx context.Context, ids []uint64) (err error)
	}
	ISystemOperLog interface {
		Model(ctx context.Context) *gdb.Model
		GetPageList(ctx context.Context, req *model.PageListReq, username string) (res []*res.SystemOperLog, total int, err error)
		Push(ctx context.Context)
		GetPageListForSearch(ctx context.Context, req *model.PageListReq, in *req.SystemOperLogSearch) (rs []*res.SystemOperLog, total int, err error)
	}
	ISystemPost interface {
		Model(ctx context.Context) *gdb.Model
		GetList(ctx context.Context, in *req.SystemPostSearch) (out []*res.SystemPost, err error)
		GetPageList(ctx context.Context, req *model.PageListReq, in *req.SystemPostSearch) (rs []*res.SystemPost, total int, err error)
		Save(ctx context.Context, in *req.SystemPostSave) (id uint64, err error)
		GetById(ctx context.Context, id uint64) (res *res.SystemPost, err error)
		Update(ctx context.Context, in *req.SystemPostSave) (err error)
		Delete(ctx context.Context, ids []uint64) (err error)
		RealDelete(ctx context.Context, ids []uint64) (err error)
		Recovery(ctx context.Context, ids []uint64) (err error)
		ChangeStatus(ctx context.Context, id uint64, status int) (err error)
		NumberOperation(ctx context.Context, id uint64, numberName string, numberValue int) (err error)
	}
	ISystemQueueMessage interface {
		Model(ctx context.Context) *gdb.Model
		GetReceiveUserPageList(ctx context.Context, req *model.PageListReq, messageId uint64) (rs []*res.MessageReceiveUser, total int, err error)
		GetPageList(ctx context.Context, req *model.PageListReq, userId uint64, params *req.SystemQueueMessageSearch) (rs []*res.SystemQueueMessage, total int, err error)
		DeletesRelated(ctx context.Context, ids []uint64, userId uint64) (err error)
		SendMessage(ctx context.Context, sendReq *req.SystemQueueMessagesSend, sendUserId uint64, contentType string) (err error, messageId uint64)
	}
	ISystemQueueMessageReceive interface {
		Model(ctx context.Context) *gdb.Model
		UpdateReadStatus(ctx context.Context, ids []uint64, userId uint64, value int) (err error)
	}
	ISystemRole interface {
		Model(ctx context.Context) *gdb.Model
		GetByIds(ctx context.Context, ids []uint64) (res []*entity.SystemRole, err error)
		Verify(r *ghttp.Request) bool
		GetList(ctx context.Context, in *req.SystemRoleSearch, filterAdminRole bool) (out []*res.SystemRole, err error)
		GetPageList(ctx context.Context, req *model.PageListReq, in *req.SystemRoleSearch, filterAdminRole bool) (rs []*res.SystemRole, total int, err error)
		Save(ctx context.Context, in *req.SystemRoleSave) (id uint64, err error)
		GetSuperAdminId(ctx context.Context) (id uint64, err error)
		GetById(ctx context.Context, id uint64) (res *res.SystemRole, err error)
		Update(ctx context.Context, in *req.SystemRoleSave) (err error)
		Delete(ctx context.Context, ids []uint64) (err error)
		RealDelete(ctx context.Context, ids []uint64) (err error)
		Recovery(ctx context.Context, ids []uint64) (err error)
		ChangeStatus(ctx context.Context, id uint64, status int) (err error)
		NumberOperation(ctx context.Context, id uint64, numberName string, numberValue int) (err error)
		GetMenuByRoleIds(ctx context.Context, ids []uint64) (out []*res.SystemRoleMenus, err error)
		GetDeptByRole(ctx context.Context, ids []uint64) (out []*res.SystemRoleDepts, err error)
	}
	ISystemRoleDept interface {
		Model(ctx context.Context) *gdb.Model
	}
	ISystemRoleMenu interface {
		Model(ctx context.Context) *gdb.Model
		GetMenuIdsByRoleIds(ctx context.Context, roleIds []uint64) (rmenuIds []uint64, err error)
	}
	ISystemUploadfile interface {
		Model(ctx context.Context) *gdb.Model
		GetPageList(ctx context.Context, in *model.PageListReq, params *req.SystemUploadFileSearch) (out []*res.SystemUploadFile, total int, err error)
		SaveDb(ctx context.Context, in *res.SystemUploadFileRes, createdBy uint64) (rs uint64, err error)
		GetByHash(ctx context.Context, hash string) (rs *res.SystemUploadFileRes, err error)
		GetFileInfoById(ctx context.Context, id uint64) (rs *res.SystemUploadFile, err error)
		GetFileInfoByHash(ctx context.Context, hash string) (rs *res.SystemUploadFile, err error)
		Delete(ctx context.Context, ids []uint64) (err error)
		RealDelete(ctx context.Context, ids []uint64) (err error)
		Recovery(ctx context.Context, ids []uint64) (err error)
	}
	ISystemUser interface {
		Model(ctx context.Context) *gdb.Model
		GetPageList(ctx context.Context, req *model.PageListReq) (res []*res.SystemUser, total int, err error)
		GetPageListForSearch(ctx context.Context, req *model.PageListReq, in *req.SystemUserSearch) (res []*res.SystemUser, total int, err error)
		GetOnlineUserPageListForSearch(ctx context.Context, req *model.PageListReq, in *req.SystemUserSearch) (res []*res.SystemUser, total int, err error)
		GetExportList(ctx context.Context, req *model.ListReq, in *req.SystemUserSearch) (res []*res.SystemUserExport, err error)
		GetSupserAdminId(ctx context.Context) uint64
		ExistsByUsername(ctx context.Context, username string) (rs bool, err error)
		GetInfoById(ctx context.Context, userId uint64) (systemUser *res.SystemUser, err error)
		GetInfoByIds(ctx context.Context, userIds []uint64) (systemUser []*res.SystemUser, err error)
		GetInfo(ctx context.Context, userId uint64) (systemUserInfo *res.SystemUserInfo, err error)
		IsSuperAdmin(ctx context.Context, userId uint64) (isSuperAdmin bool, err error)
		GetRoles(ctx context.Context, userId uint64) (roles []uint64, err error)
		GetDepts(ctx context.Context, userId uint64) (depts []uint64, err error)
		Update(ctx context.Context, req *req.SystemUser, userId ...uint64) (rs sql.Result, err error)
		SetHomePage(ctx context.Context, id uint64, dashboard string) (out sql.Result, err error)
		InitUserPassword(ctx context.Context, id uint64, password string) (out sql.Result, err error)
		UpdateSimple(ctx context.Context, in *req.SystemUserUpdate) (out sql.Result, err error)
		Save(ctx context.Context, in *req.SystemUserSave) (id uint64, err error)
		GetFullInfoById(ctx context.Context, id uint64) (out *res.SystemUserFullInfo, err error)
		Delete(ctx context.Context, ids []uint64) (err error)
		RealDelete(ctx context.Context, ids []uint64) (err error)
		Recovery(ctx context.Context, ids []uint64) (err error)
		ChangeStatus(ctx context.Context, id uint64, status int) (err error)
	}
	ISystemUserDept interface {
		Model(ctx context.Context) *gdb.Model
	}
	ISystemUserPost interface {
		Model(ctx context.Context) *gdb.Model
	}
	ISystemUserRole interface {
		Model(ctx context.Context) *gdb.Model
	}
)

var (
	localDataMaintain              IDataMaintain
	localLogin                     ILogin
	localSettingConfig             ISettingConfig
	localSettingConfigGroup        ISettingConfigGroup
	localSettingCrontab            ISettingCrontab
	localSettingCrontabLog         ISettingCrontabLog
	localSettingGenerateColumns    ISettingGenerateColumns
	localSettingGenerateTables     ISettingGenerateTables
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

func SettingGenerateColumns() ISettingGenerateColumns {
	if localSettingGenerateColumns == nil {
		panic("implement not found for interface ISettingGenerateColumns, forgot register?")
	}
	return localSettingGenerateColumns
}

func RegisterSettingGenerateColumns(i ISettingGenerateColumns) {
	localSettingGenerateColumns = i
}

func SettingGenerateTables() ISettingGenerateTables {
	if localSettingGenerateTables == nil {
		panic("implement not found for interface ISettingGenerateTables, forgot register?")
	}
	return localSettingGenerateTables
}

func RegisterSettingGenerateTables(i ISettingGenerateTables) {
	localSettingGenerateTables = i
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