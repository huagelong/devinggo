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
	"devinggo/modules/system/api/system"
	"devinggo/modules/system/consts"
	"devinggo/modules/system/controller/base"
	"devinggo/modules/system/model/req"
	"devinggo/modules/system/model/res"
	"devinggo/modules/system/myerror"
	"devinggo/modules/system/pkg/cache"
	"devinggo/modules/system/pkg/excel"
	"devinggo/modules/system/pkg/orm"
	"devinggo/modules/system/pkg/utils"
	"devinggo/modules/system/pkg/utils/config"
	"devinggo/modules/system/pkg/utils/request"
	"devinggo/modules/system/pkg/utils/secure"
	"devinggo/modules/system/service"
	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/gmode"
)

var (
	UserController = userController{}
)

type userController struct {
	base.BaseController
}

func (c *userController) GetInfo(ctx context.Context, in *system.GetInfoReq) (out *system.GetInfoRes, err error) {
	userId := c.UserId
	systemUserInfo, err := service.SystemUser().GetInfo(ctx, userId)
	if err != nil {
		return
	}
	if g.IsEmpty(systemUserInfo) {
		return nil, nil
	}
	err = gconv.Struct(systemUserInfo, &out)
	if err != nil {
		return
	}
	return
}

func (c *userController) UpdateInfo(ctx context.Context, in *system.UpdateInfoReq) (out *system.UpdateInfoRes, err error) {
	var systemUser *req.SystemUser
	err = gconv.Struct(in, &systemUser)
	if err != nil {
		return
	}
	_, err = service.SystemUser().Update(ctx, systemUser, c.UserId)
	return
}

func (c *userController) ModifyPassword(ctx context.Context, in *system.ModifyPasswordReq) (out *system.ModifyPasswordRes, err error) {

	newPassword := in.NewPassword
	oldPassword := in.OldPassword
	newPasswordConfirmation := in.NewPasswordConfirmation
	userId := c.UserId

	if newPassword != newPasswordConfirmation {
		return nil, myerror.ValidationFailed(ctx, "新密码与确认密码不一致")
	}
	var userInfo *entity.SystemUser
	err = service.SystemUser().Model(ctx).Where(dao.SystemUser.Columns().Id, userId).Scan(&userInfo)
	if utils.IsError(err) {
		return
	}

	if g.IsEmpty(userInfo) {
		err = myerror.ValidationFailed(ctx, "用户不存在")
		return
	}

	isSuperAdmin, _ := service.SystemUser().IsSuperAdmin(ctx, userId)
	if isSuperAdmin && (gmode.Mode() == gmode.DEVELOP) {
		err = myerror.ValidationFailed(ctx, "超级管理员开发环境不允许修改密码")
		return
	}

	if !secure.PasswordVerify(oldPassword, userInfo.Password) {
		err = myerror.ValidationFailed(ctx, "旧密码错误")
		return
	}
	passwordHash, err := secure.PasswordHash(newPassword)
	if err != nil {
		return
	}
	_, err = service.SystemUser().Model(ctx).Data(g.Map{dao.SystemUser.Columns().Password: passwordHash}).Where(dao.SystemUser.Columns().Id, userId).Update()
	if utils.IsError(err) {
		return
	}
	err = service.Token().Logout(request.GetHttpRequest(ctx))

	return
}

func (c *userController) IndexUser(ctx context.Context, in *system.IndexUserReq) (out *system.IndexUserRes, err error) {
	out = &system.IndexUserRes{}
	items, totalCount, err := service.SystemUser().GetPageListForSearch(ctx, &in.PageListReq, &in.SystemUserSearch)
	if err != nil {
		return
	}

	if !g.IsEmpty(items) {
		for _, item := range items {
			out.Items = append(out.Items, *item)
		}
	} else {
		out.Items = make([]res.SystemUser, 0)
	}
	out.PageRes.Pack(in, totalCount)
	return
}
func (c *userController) IndexOnlineUser(ctx context.Context, in *system.IndexOnlineUserReq) (out *system.IndexOnlineUserRes, err error) {
	out = &system.IndexOnlineUserRes{}
	items, totalCount, err := service.SystemUser().GetOnlineUserPageListForSearch(ctx, &in.PageListReq, &in.SystemUserSearch)
	if err != nil {
		return
	}

	if !g.IsEmpty(items) {
		for _, item := range items {
			out.Items = append(out.Items, *item)
		}
	} else {
		out.Items = make([]res.SystemUser, 0)
	}
	out.PageRes.Pack(in, totalCount)
	return
}

func (c *userController) KickUser(ctx context.Context, in *system.KickUserReq) (out *system.KickUserRes, err error) {
	out = &system.KickUserRes{}
	r := request.GetHttpRequest(ctx)
	err = service.Token().Kick(r, in.Id, in.AppId)
	if err != nil {
		return
	}
	return
}

func (c *userController) RecycleUser(ctx context.Context, in *system.RecycleUserReq) (out *system.RecycleUserRes, err error) {
	out = &system.RecycleUserRes{}
	in.Recycle = true
	in.FilterAuth = true
	items, totalCount, err := service.SystemUser().GetPageListForSearch(ctx, &in.PageListReq, &in.SystemUserSearch)
	if err != nil {
		return
	}

	if !g.IsEmpty(items) {
		for _, item := range items {
			out.Items = append(out.Items, *item)
		}
	} else {
		out.Items = make([]res.SystemUser, 0)
	}
	out.PageRes.Pack(in, totalCount)
	return
}

func (c *userController) SaveUser(ctx context.Context, in *system.SaveUserReq) (out *system.SaveUserRes, err error) {
	out = &system.SaveUserRes{}
	id, err := service.SystemUser().Save(ctx, &in.SystemUserSave)
	if err != nil {
		return
	}
	out.Id = id
	return
}

func (c *userController) ReadUser(ctx context.Context, in *system.ReadUserReq) (out *system.ReadUserRes, err error) {
	out = &system.ReadUserRes{}
	info, err := service.SystemUser().GetFullInfoById(ctx, in.Id)
	if err != nil {
		return
	}
	out.Data = *info
	return
	return
}

func (c *userController) ClearCache(ctx context.Context, in *system.ClearCacheReq) (out *system.ClearCacheRes, err error) {
	utils.SafeGo(ctx, func(ctx context.Context) {
		cache.RemoveByTag(ctx, consts.USER_CACHE_TAG+gconv.String(in.Id))
	})
	return
}

func (c *userController) Export(ctx context.Context, in *system.ExportReq) (out *system.ExportRes, err error) {
	var (
		fileName  = "用户列表"
		sheetName = "Sheet1"
	)
	exports, err := service.SystemUser().GetExportList(ctx, &in.ListReq, &in.SystemUserSearch)
	if err != nil {
		return
	}
	//创建导出对象
	export := excel.NewExcelExport(sheetName, res.SystemUserExport{})
	//销毁对象
	defer export.Close()
	newExports := []res.SystemUserExport{}
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

func (c *userController) Import(ctx context.Context, in *system.ImportReq) (out *system.ImportRes, err error) {
	tmpPath := utils.GetTmpDir()
	fileName, err := in.File.Save(tmpPath, true)
	if err != nil {
		return nil, err
	}
	localPath := tmpPath + "/" + fileName
	var result []res.SystemUserExport
	//创建导入对象
	importFile := excel.NewExcelImportFile(localPath, res.SystemUserExport{})
	//对象销毁
	defer importFile.Close()

	//数据填充
	err = importFile.ImportDataToStruct(&result).Error()
	//数据显示
	if err != nil {
		return nil, err
	} else {
		if !g.IsEmpty(result) {
			err = dao.SystemUser.Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
				for _, item := range result {
					initPassword := config.GetConfigString(ctx, "settings.initPassword", "123456")
					_, err = service.SystemUser().Save(ctx, &req.SystemUserSave{
						Username: item.Username,
						Nickname: item.Nickname,
						Password: initPassword,
						UserType: "100",
						Phone:    item.Phone,
						Status:   item.Status,
					})
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

func (c *userController) DownloadTemplate(ctx context.Context, in *system.DownloadTemplateReq) (out *system.DownloadTemplateRes, err error) {
	var (
		fileName  = "模板下载"
		sheetName = "Sheet1"
		exports   = make([]res.SystemUserExport, 0)
	)
	//创建导出对象
	export := excel.NewExcelExport(sheetName, res.SystemUserExport{})
	//销毁对象
	defer export.Close()
	err = export.ExportSmallExcelByStruct(exports).Download(ctx, fileName).Error()
	if err != nil {
		return
	}
	return
}

func (c *userController) SetHomePage(ctx context.Context, in *system.SetHomePageReq) (out *system.SetHomePageRes, err error) {
	out = &system.SetHomePageRes{}
	_, err = service.SystemUser().SetHomePage(ctx, in.Id, in.Dashboard)
	if err != nil {
		return
	}
	return
	return
}

func (c *userController) Update(ctx context.Context, in *system.UpdateUserReq) (out *system.UpdateUserRes, err error) {
	out = &system.UpdateUserRes{}
	_, err = service.SystemUser().UpdateSimple(ctx, &in.SystemUserUpdate)
	if err != nil {
		return
	}
	return
}

func (c *userController) InitUserPassword(ctx context.Context, in *system.InitUserPasswordReq) (out *system.InitUserPasswordRes, err error) {
	out = &system.InitUserPasswordRes{}
	password := config.GetConfigString(ctx, "settings.initPassword", "123456")
	_, err = service.SystemUser().InitUserPassword(ctx, in.Id, password)
	if err != nil {
		return
	}
	return
}

func (c *userController) DeleteUser(ctx context.Context, in *system.DeleteUserReq) (out *system.DeleteUserRes, err error) {
	out = &system.DeleteUserRes{}
	err = service.SystemUser().Delete(ctx, in.Ids)
	if err != nil {
		return
	}
	return
}

func (c *userController) RealDeleteUser(ctx context.Context, in *system.RealDeleteUserReq) (out *system.RealDeleteUserRes, err error) {
	out = &system.RealDeleteUserRes{}
	err = service.SystemUser().RealDelete(ctx, in.Ids)
	if err != nil {
		return
	}
	return
}

func (c *userController) RecoveryUser(ctx context.Context, in *system.RecoveryUserReq) (out *system.RecoveryUserRes, err error) {
	out = &system.RecoveryUserRes{}
	err = service.SystemUser().Recovery(ctx, in.Ids)
	if err != nil {
		return
	}
	return
}

func (c *userController) ChangeStatusUser(ctx context.Context, in *system.ChangeStatusUserReq) (out *system.ChangeStatusUserRes, err error) {
	out = &system.ChangeStatusUserRes{}
	err = service.SystemUser().ChangeStatus(ctx, in.Id, in.Status)
	if err != nil {
		return
	}
	if in.Status == 2 {
		r := request.GetHttpRequest(ctx)
		err = service.Token().KickAll(r, in.Id)
		if err != nil {
			return
		}
	}

	return
}

func (c *userController) Remote(ctx context.Context, in *system.RemoteUserReq) (out *system.RemoteUserRes, err error) {
	out = &system.RemoteUserRes{}
	r := request.GetHttpRequest(ctx)

	params := gmap.NewStrAnyMapFrom(r.GetMap())
	m := service.SystemUser().Model(ctx)
	var rs res.SystemUser
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
			out.Items = make([]res.SystemUser, 0)
		}
		out.PageRes.Pack(in, totalCount)
	} else {
		if !g.IsEmpty(items) {
			out.Data = items
		} else {
			out.Data = make([]res.SystemUser, 0)
		}
	}
	return
}
