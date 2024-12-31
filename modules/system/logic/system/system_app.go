// Package system
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package system

import (
	"devinggo/internal/dao"
	"devinggo/internal/logic/base"
	"devinggo/internal/model"
	"devinggo/internal/model/do"
	"devinggo/internal/model/entity"
	"devinggo/modules/system/consts"
	"devinggo/modules/system/model/req"
	"devinggo/modules/system/model/res"
	"devinggo/modules/system/myerror"
	"devinggo/modules/system/pkg/contexts"
	"devinggo/modules/system/pkg/hook"
	"devinggo/modules/system/pkg/orm"
	"devinggo/modules/system/pkg/utils"
	"devinggo/modules/system/service"
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"net/url"
	"sort"
	"strings"
)

type sSystemApp struct {
	base.BaseService
}

func init() {
	service.RegisterSystemApp(NewSystemApp())
}

func NewSystemApp() *sSystemApp {
	return &sSystemApp{}
}

func (s *sSystemApp) Model(ctx context.Context) *gdb.Model {
	return dao.SystemApp.Ctx(ctx).Hook(hook.Bind()).Cache(orm.SetCacheOption(ctx))
}

func (s *sSystemApp) GetAppId(ctx context.Context) (string, error) {
	randomBytes := make([]byte, 5)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(randomBytes), nil
}

func (s *sSystemApp) GetAppSecret(ctx context.Context) (string, error) {
	// 生成32个字节的随机数
	randomBytes := make([]byte, 32)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}
	// 将随机字节转换为十六进制字符串
	hexStr := hex.EncodeToString(randomBytes)

	// 将十六进制字符串编码为base64
	base64Str := base64.StdEncoding.EncodeToString([]byte(hexStr))

	return base64Str, nil
}

func (s *sSystemApp) BindApp(ctx context.Context, Id uint64, ApiIds []uint64) (err error) {
	_, err = service.SystemAppApi().Model(ctx).Where("app_id", Id).Delete()
	if utils.IsError(err) {
		return err
	}
	for _, apiId := range ApiIds {
		_, err = service.SystemAppApi().Model(ctx).Insert(g.Map{
			"app_id": Id,
			"api_id": apiId,
		})
		if utils.IsError(err) {
			return err
		}
	}
	return
}

func (s *sSystemApp) GetPageListForSearch(ctx context.Context, req *model.PageListReq, in *req.SystemAppSearch) (rs []*res.SystemApp, total int, err error) {
	m := s.handleSearch(ctx, in)
	var entity []*entity.SystemApp
	err = orm.GetPageList(m, req).ScanAndCount(&entity, &total, false)
	if utils.IsError(err) {
		return nil, 0, err
	}
	rs = make([]*res.SystemApp, 0)
	if !g.IsEmpty(entity) {
		if err = gconv.Structs(entity, &rs); err != nil {
			return nil, 0, err
		}
	}
	return
}

func (s *sSystemApp) GetApiList(ctx context.Context, id uint64) (out []uint64, err error) {
	result, err := service.SystemAppApi().Model(ctx).Fields("api_id").Where("app_id", id).Array()
	if utils.IsError(err) {
		return nil, err
	}
	if !g.IsEmpty(result) {
		out = gconv.SliceUint64(result)
	}
	return
}

func (s *sSystemApp) handleSearch(ctx context.Context, in *req.SystemAppSearch) (m *gdb.Model) {
	m = s.Model(ctx)
	if !g.IsEmpty(in.AppName) {
		m = m.Where("app_name", in.AppName)
	}

	if !g.IsEmpty(in.AppId) {
		m = m.Where("app_id", in.AppId)
	}

	if !g.IsEmpty(in.GroupId) {
		m = m.Where("group_id", in.GroupId)
	}

	if !g.IsEmpty(in.Status) {
		m = m.Where("status", in.Status)
	}
	return
}

func (s *sSystemApp) Save(ctx context.Context, in *req.SystemAppSave, userId uint64) (id uint64, err error) {
	saveData := do.SystemApp{
		GroupId:     in.GroupId,
		AppSecret:   in.AppSecret,
		AppName:     in.AppName,
		Status:      in.Status,
		AppId:       in.AppId,
		Description: in.Description,
		Remark:      in.Remark,
	}
	rs, err := s.Model(ctx).Data(saveData).Save()
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

func (s *sSystemApp) GetById(ctx context.Context, id uint64) (res *res.SystemApp, err error) {
	err = s.Model(ctx).Where("id", id).Scan(&res)
	if utils.IsError(err) {
		return
	}
	return
}

func (s *sSystemApp) Update(ctx context.Context, in *req.SystemAppUpdate) (err error) {
	updateData := do.SystemApp{
		GroupId:     in.GroupId,
		AppSecret:   in.AppSecret,
		AppName:     in.AppName,
		Status:      in.Status,
		AppId:       in.AppId,
		Description: in.Description,
		Remark:      in.Remark,
	}
	_, err = s.Model(ctx).Data(updateData).Where("id", in.Id).Update()
	if utils.IsError(err) {
		return
	}
	return
}

func (s *sSystemApp) Delete(ctx context.Context, ids []uint64) (err error) {
	_, err = s.Model(ctx).WhereIn("id", ids).Delete()
	if utils.IsError(err) {
		return err
	}
	return
}

func (s *sSystemApp) RealDelete(ctx context.Context, ids []uint64) (err error) {
	var res []*res.SystemApp
	err = s.Model(ctx).Unscoped().WhereIn("id", ids).Scan(&res)
	if utils.IsError(err) {
		return
	}
	return
}

func (s *sSystemApp) Recovery(ctx context.Context, ids []uint64) (err error) {
	_, err = s.Model(ctx).Unscoped().WhereIn("id", ids).Update(g.Map{"deleted_at": nil})
	if utils.IsError(err) {
		return err
	}
	return
}

func (s *sSystemApp) ChangeStatus(ctx context.Context, id uint64, status int) (err error) {
	_, err = s.Model(ctx).Data(g.Map{"status": status}).Where("id", id).Update()
	if utils.IsError(err) {
		return err
	}
	return
}

func (s *sSystemApp) Verify(r *ghttp.Request) (bool, error) {
	ctx := r.Context()
	permission := contexts.New().GetPermission(ctx)
	if g.IsEmpty(permission) {
		return false, myerror.ValidationFailed(ctx, "权限标识未定义")
	}

	var api *entity.SystemApi
	err := service.SystemApi().Model(ctx).Where("access_name", permission).Scan(&api)
	if utils.IsError(err) {
		return false, err
	}
	authMode := api.AuthMode
	if authMode == 1 { //简易模式
		signatureRs := r.Get("signature")
		if g.IsEmpty(signatureRs) {
			return false, myerror.MissingParameter(ctx, "参数不存在: signature")
		}
		signature := signatureRs.String()
		appId := contexts.New().GetAppId(ctx)
		check, err := s.VerifyEasyMode(ctx, appId, signature, api.Id)
		if err != nil {
			return false, err
		}
		if !check {
			return false, myerror.ValidationFailed(ctx, "接口鉴权失败")
		}
		return true, nil
	} else if authMode == 2 { //复杂模式
		token := service.Token().GetAuthorization(r)
		if g.IsEmpty(token) {
			return false, myerror.MissingParameter(ctx, "参数不存在: token")
		}
		check, err := s.verifyNormalMode(ctx, token, api.Id)
		if err != nil {
			return false, err
		}
		if !check {
			return false, myerror.ValidationFailed(ctx, "接口鉴权失败")
		}
		return true, nil
	} else {
		return false, myerror.ValidationFailed(ctx, "接口鉴权失败")
	}
}

// getAccessToken 获取access_token
func (s *sSystemApp) GetAccessToken(ctx context.Context, params map[string]interface{}) (token string, exp int64, err error) {
	now := gtime.Now()
	if appIdRs, ok := params["app_id"]; !ok || g.IsEmpty(appIdRs) {
		return "", now.Unix(), myerror.MissingParameter(ctx, "参数不存在：app_id")
	}

	if signatureRs, ok := params["signature"]; !ok || g.IsEmpty(signatureRs) {
		return "", now.Unix(), myerror.MissingParameter(ctx, "参数不存在：signature")
	}

	appId := gconv.String(params["app_id"])
	signature := gconv.String(params["signature"])

	var app *entity.SystemApp
	err = s.Model(ctx).Where("app_id", appId).Scan(&app)
	if utils.IsError(err) {
		return
	}

	if g.IsEmpty(app) {
		err = myerror.ValidationFailed(ctx, "应用不存在")
		return
	}

	newSignature := s.GetSignature(app.AppSecret, params)
	g.Log().Info(ctx, "newSignature:", newSignature)
	if newSignature != signature {
		err = myerror.ValidationFailed(ctx, "签名验证失败")
		return
	}
	delete(params, "signature")
	scene := consts.ApiScene + "_" + appId
	token, exp, err = service.Token().GetToken(ctx, scene, params)
	return
}

// getSignature 获取签名
func (s *sSystemApp) GetSignature(appSecret string, params map[string]interface{}) string {
	// 删除签名参数
	delete(params, "signature")
	// 准备数据
	data := url.Values{}
	data.Set("app_secret", appSecret)

	for key, value := range params {
		data.Set(key, gconv.String(value))
	}
	// 对数据进行排序
	var keys []string
	for k := range data {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		return strings.Compare(keys[i], keys[j]) == -1 // 降序排序
	})
	var sortedData strings.Builder
	for _, k := range keys {
		if sortedData.Len() > 0 {
			sortedData.WriteString("&")
		}
		sortedData.WriteString(url.QueryEscape(k))
		sortedData.WriteString("=")
		sortedData.WriteString(url.QueryEscape(data.Get(k)))
	}
	hash, _ := gmd5.EncryptString(sortedData.String())
	return hash
}

func (s *sSystemApp) checkAppHasBindApi(ctx context.Context, appId, apiId uint64) (bool, error) {
	count, err := service.SystemAppApi().Model(ctx).Where("app_id", appId).Where("api_id", apiId).Count()
	if utils.IsError(err) {
		return false, err
	}
	if count > 0 {
		return true, nil
	}
	return false, nil
}

// 复杂模式
func (s *sSystemApp) verifyNormalMode(ctx context.Context, token string, apiId uint64) (check bool, err error) {
	identity, err := service.Token().ParseToken(ctx, token)
	if err != nil {
		return false, err
	}

	if g.IsEmpty(identity) {
		return false, myerror.ValidationFailed(ctx, "接口鉴权失败")
	}

	scene := identity.Scene
	data := identity.Data
	appIdTmp, ok := data["app_id"]
	if !ok || g.IsEmpty(appIdTmp) {
		return false, myerror.ValidationFailed(ctx, "接口鉴权失败")
	}
	appId := gconv.String(appIdTmp)
	newScene := consts.ApiScene + "_" + appId
	if scene != newScene {
		return false, myerror.ValidationFailed(ctx, "接口鉴权失败")
	}

	check, _, err = s.VerifyPre(ctx, appId, apiId)
	if err != nil {
		return false, err
	}

	if !check {
		err = myerror.ValidationFailed(ctx, "接口鉴权失败")
		return false, err
	}

	return true, nil
}

func (s *sSystemApp) VerifyPre(ctx context.Context, appId string, apiId uint64) (check bool, app *entity.SystemApp, err error) {
	err = s.Model(ctx).Where("app_id", appId).Scan(&app)
	if utils.IsError(err) {
		return false, app, err
	}

	if g.IsEmpty(app) {
		err = myerror.ValidationFailed(ctx, "应用不存在")
		return false, app, err
	}

	if app.Status != 1 {
		err = myerror.ValidationFailed(ctx, "应用未激活")
		return false, app, err
	}

	check, err = s.checkAppHasBindApi(ctx, app.Id, apiId)
	if err != nil {
		err = myerror.ValidationFailed(ctx, "接口未绑定")
		return false, app, err
	}

	if !check {
		err = myerror.ValidationFailed(ctx, "接口鉴权失败")
		return false, app, err
	}
	return true, app, nil
}

// 简单模式
func (s *sSystemApp) VerifyEasyMode(ctx context.Context, appId string, signature string, apiId uint64) (check bool, err error) {
	check, app, err := s.VerifyPre(ctx, appId, apiId)
	if err != nil {
		return false, err
	}

	if !check {
		err = myerror.ValidationFailed(ctx, "接口鉴权失败")
		return false, err
	}

	md5Str, err := gmd5.EncryptString(appId + app.AppSecret)
	if err != nil {
		err = myerror.ValidationFailed(ctx, "接口鉴权失败")
		return false, err
	}
	g.Log().Info(ctx, "md5Str:", md5Str)
	if signature != md5Str {
		err = myerror.ValidationFailed(ctx, "接口鉴权失败")
		return false, err
	}

	return true, nil
}