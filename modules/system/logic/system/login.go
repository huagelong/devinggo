// Package system
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package system

import (
	"context"
	"devinggo/internal/dao"
	"devinggo/internal/model/entity"
	consts2 "devinggo/modules/system/consts"
	"devinggo/modules/system/logic/base"
	"devinggo/modules/system/model"
	"devinggo/modules/system/myerror"
	"devinggo/modules/system/pkg/contexts"
	"devinggo/modules/system/pkg/utils"
	"devinggo/modules/system/pkg/utils/config"
	"devinggo/modules/system/pkg/utils/location"
	"devinggo/modules/system/pkg/utils/request"
	"devinggo/modules/system/pkg/utils/secure"
	"devinggo/modules/system/service"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type sLogin struct {
	base.BaseService
}

func init() {
	service.RegisterLogin(NewLogin())
}

func NewLogin() *sLogin {
	return &sLogin{}
}

func (s *sLogin) Model(ctx context.Context) *gdb.Model {
	return dao.SystemUser.Ctx(ctx).OnConflict("id")
}

func (s *sLogin) Login(ctx context.Context, username, password string) (token string, expire int64, err error) {
	//记录日志
	defer func() {
		service.SystemLoginLog().Push(ctx, username, err)
	}()
	userInfo := &entity.SystemUser{}
	err = s.Model(ctx).Where(dao.SystemUser.Columns().Username, username).Scan(userInfo)
	if utils.IsError(err) {
		return
	}

	if g.IsNil(userInfo) {
		err = myerror.ValidationFailed(ctx, "用户不存在")
		return
	}
	// AES解密密码
	aesKey := config.GetConfigString(ctx, "settings.aesKey")
	if aesKey == "" {
		err = myerror.ValidationFailed(ctx, "系统加密配置异常")
		return
	}

	decryptedPass, err := secure.AESDecrypt(password, aesKey)
	if err != nil {
		g.Log().Errorf(ctx, "AES解密失败: %v", err)
		err = myerror.ValidationFailed(ctx, "密码解析失败")
		return
	}

	if !secure.PasswordVerify(decryptedPass, userInfo.Password) {
		err = myerror.ValidationFailed(ctx, "用户或者密码错误")
		return
	}

	status := userInfo.Status
	if !((status == consts2.UserNormal) || (status == consts2.UserBan && userInfo.Id == 1)) {
		err = myerror.ValidationFailed(ctx, "没有权限登录")
		return
	}

	userType := userInfo.UserType
	if userType != consts2.TypeSysUser {
		err = myerror.ValidationFailed(ctx, "没有权限登录")
		return
	}

	roleIds, _ := service.SystemUser().GetRoles(ctx, userInfo.Id)

	deptIds, _ := service.SystemUser().GetDepts(ctx, userInfo.Id)
	appId := contexts.New().GetAppId(ctx)
	token, expire, err = service.Token().GenerateUserToken(ctx, consts2.AdminScene, appId, &model.Identity{
		Id:       userInfo.Id,
		AppId:    appId,
		RoleIds:  roleIds,
		DeptIds:  deptIds,
		Username: userInfo.Username,
	})

	if err != nil {
		return
	}

	r := request.GetHttpRequest(ctx)
	if r == nil {
		g.Log().Warningf(ctx, "ctx not http request")
		return
	}
	//更新登录信息
	clientIp := location.GetClientIp(r)
	loginTime := gtime.Now()
	s.Model(ctx).Data(g.Map{dao.SystemUser.Columns().LoginIp: clientIp, dao.SystemUser.Columns().LoginTime: loginTime}).Where(dao.SystemUser.Columns().Id, userInfo.Id).Update()
	return
}
