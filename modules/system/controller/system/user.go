// Package system
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package system

import (
	"context"
	"strings"

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
	"devinggo/modules/system/pkg/utils"
	"devinggo/modules/system/pkg/utils/config"
	"devinggo/modules/system/pkg/utils/request"
	"devinggo/modules/system/pkg/utils/secure"
	"devinggo/modules/system/service"

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

// GetInfo retrieves the current user's information including roles and permissions.
func (c *userController) GetInfo(ctx context.Context, in *system.GetInfoReq) (out *system.GetInfoRes, err error) {
	userId := c.UserId
	systemUserInfo, err := service.SystemUser().GetInfo(ctx, userId)
	if utils.IsError(err) {
		return
	}
	if g.IsEmpty(systemUserInfo) {
		return nil, nil
	}
	err = gconv.Struct(systemUserInfo, &out)
	if utils.IsError(err) {
		return
	}
	return
}

// UpdateInfo updates the current user's profile information.
func (c *userController) UpdateInfo(ctx context.Context, in *system.UpdateInfoReq) (out *system.UpdateInfoRes, err error) {
	var systemUser *req.SystemUser
	err = gconv.Struct(in, &systemUser)
	if utils.IsError(err) {
		return
	}
	_, err = service.SystemUser().Update(ctx, systemUser, c.UserId)
	return
}

// ModifyPassword changes the current user's password after verifying the old password.
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
	if utils.IsError(err) {
		return
	}
	_, err = service.SystemUser().Model(ctx).Data(g.Map{dao.SystemUser.Columns().Password: passwordHash}).Where(dao.SystemUser.Columns().Id, userId).Update()
	if utils.IsError(err) {
		return
	}
	err = service.Token().Logout(request.GetHttpRequest(ctx))

	return
}

// IndexUser retrieves a paginated list of users with their associated departments, roles, and posts.
func (c *userController) IndexUser(ctx context.Context, in *system.IndexUserReq) (out *system.IndexUserRes, err error) {
	out = &system.IndexUserRes{}
	items, totalCount, err := service.SystemUser().GetPageListForSearch(ctx, &in.PageListReq, &in.SystemUserSearch)
	if utils.IsError(err) {
		return
	}

	if !g.IsEmpty(items) {
		for _, item := range items {
			//获取用户所属部门
			item.DeptIds, _ = service.SystemUserDept().GetDeptIdsByUserId(ctx, item.Id)
			//获取用户所属角色
			item.RoleIds, _ = service.SystemUserRole().GetRoleIdsByUserId(ctx, item.Id)
			//获取用户所属岗位
			item.PostIds, _ = service.SystemUserPost().GetPostIdsByUserId(ctx, item.Id)
			if err = c.fillUserRelatedNames(ctx, item); utils.IsError(err) {
				return
			}

			out.Items = append(out.Items, *item)
		}
	} else {
		out.Items = make([]res.SystemUser, 0)
	}
	out.PageRes.Pack(in, totalCount)
	return
}

// IndexOnlineUser retrieves a paginated list of currently online users.
func (c *userController) IndexOnlineUser(ctx context.Context, in *system.IndexOnlineUserReq) (out *system.IndexOnlineUserRes, err error) {
	out = &system.IndexOnlineUserRes{}
	items, totalCount, err := service.SystemUser().GetOnlineUserPageListForSearch(ctx, &in.PageListReq, &in.SystemUserSearch)
	if utils.IsError(err) {
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

func (c *userController) fillUserRelatedNames(ctx context.Context, item *res.SystemUser) (err error) {
	if !g.IsEmpty(item.DeptIds) {
		var deptList []*entity.SystemDept
		err = service.SystemDept().Model(ctx).Fields(dao.SystemDept.Columns().Name).WhereIn(dao.SystemDept.Columns().Id, item.DeptIds).Scan(&deptList)
		if utils.IsError(err) {
			return
		}
		if !g.IsEmpty(deptList) {
			deptNames := make([]string, 0, len(deptList))
			for _, dept := range deptList {
				deptNames = append(deptNames, dept.Name)
			}
			item.DeptName = strings.Join(deptNames, ",")
		}
	}

	if !g.IsEmpty(item.RoleIds) {
		var roleList []*entity.SystemRole
		roleList, err = service.SystemRole().GetByIds(ctx, item.RoleIds)
		if utils.IsError(err) {
			return
		}
		if !g.IsEmpty(roleList) {
			roleNames := make([]string, 0, len(roleList))
			for _, role := range roleList {
				roleNames = append(roleNames, role.Name)
			}
			item.RoleName = strings.Join(roleNames, ",")
		}
	}

	if !g.IsEmpty(item.PostIds) {
		var postList []*entity.SystemPost
		err = service.SystemPost().Model(ctx).Fields(dao.SystemPost.Columns().Name).WhereIn(dao.SystemPost.Columns().Id, item.PostIds).Scan(&postList)
		if utils.IsError(err) {
			return
		}
		if !g.IsEmpty(postList) {
			postNames := make([]string, 0, len(postList))
			for _, post := range postList {
				postNames = append(postNames, post.Name)
			}
			item.PostName = strings.Join(postNames, ",")
		}
	}

	return
}

// KickUser forcibly logs out a user from the system.
func (c *userController) KickUser(ctx context.Context, in *system.KickUserReq) (out *system.KickUserRes, err error) {
	out = &system.KickUserRes{}
	r := request.GetHttpRequest(ctx)
	err = service.Token().Kick(r, in.Id, in.AppId)
	if utils.IsError(err) {
		return
	}
	return
}

// RecycleUser retrieves a paginated list of soft-deleted users for recovery purposes.
func (c *userController) RecycleUser(ctx context.Context, in *system.RecycleUserReq) (out *system.RecycleUserRes, err error) {
	out = &system.RecycleUserRes{}
	in.Recycle = true
	in.FilterAuth = true
	items, totalCount, err := service.SystemUser().GetPageListForSearch(ctx, &in.PageListReq, &in.SystemUserSearch)
	if utils.IsError(err) {
		return
	}

	if !g.IsEmpty(items) {
		for _, item := range items {
			item.DeptIds, _ = service.SystemUserDept().GetDeptIdsByUserId(ctx, item.Id)
			item.RoleIds, _ = service.SystemUserRole().GetRoleIdsByUserId(ctx, item.Id)
			item.PostIds, _ = service.SystemUserPost().GetPostIdsByUserId(ctx, item.Id)
			if err = c.fillUserRelatedNames(ctx, item); utils.IsError(err) {
				return
			}
			out.Items = append(out.Items, *item)
		}
	} else {
		out.Items = make([]res.SystemUser, 0)
	}
	out.PageRes.Pack(in, totalCount)
	return
}

// SaveUser creates a new user in the system.
func (c *userController) SaveUser(ctx context.Context, in *system.SaveUserReq) (out *system.SaveUserRes, err error) {
	out = &system.SaveUserRes{}
	id, err := service.SystemUser().Save(ctx, &in.SystemUserSave)
	if utils.IsError(err) {
		return
	}
	out.Id = id
	return
}

// ReadUser retrieves detailed information about a specific user.
func (c *userController) ReadUser(ctx context.Context, in *system.ReadUserReq) (out *system.ReadUserRes, err error) {
	out = &system.ReadUserRes{}
	info, err := service.SystemUser().GetFullInfoById(ctx, in.Id)
	if utils.IsError(err) {
		return
	}
	out.Data = *info
	return
}

// ClearCache clears the cache for a specific user.
func (c *userController) ClearCache(ctx context.Context, in *system.ClearCacheReq) (out *system.ClearCacheRes, err error) {
	utils.SafeGo(ctx, func(ctx context.Context) {
		cache.RemoveByTag(ctx, consts.USER_CACHE_TAG+gconv.String(in.Id))
	})
	return
}

// Export exports user data to an Excel file.
func (c *userController) Export(ctx context.Context, in *system.ExportReq) (out *system.ExportRes, err error) {
	var (
		fileName  = "用户列表"
		sheetName = "Sheet1"
	)
	exports, err := service.SystemUser().GetExportList(ctx, &in.ListReq, &in.SystemUserSearch)
	if utils.IsError(err) {
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
	if utils.IsError(err) {
		return
	}
	return
}

// Import imports user data from an Excel file.
func (c *userController) Import(ctx context.Context, in *system.ImportReq) (out *system.ImportRes, err error) {
	tmpPath := utils.GetTmpDir()
	fileName, err := in.File.Save(tmpPath, true)
	if utils.IsError(err) {
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
	if utils.IsError(err) {
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
					if utils.IsError(err) {
						return err
					}
				}
				return
			})
			if utils.IsError(err) {
				return
			}
		} else {
			err = myerror.ValidationFailed(ctx, "没有数据!")
		}
	}
	return
}

// DownloadTemplate downloads an Excel template for user import.
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
	if utils.IsError(err) {
		return
	}
	return
}

// SetHomePage sets the home page dashboard for a user.
func (c *userController) SetHomePage(ctx context.Context, in *system.SetHomePageReq) (out *system.SetHomePageRes, err error) {
	out = &system.SetHomePageRes{}
	_, err = service.SystemUser().SetHomePage(ctx, in.Id, in.Dashboard)
	if utils.IsError(err) {
		return
	}
	return
}

// Update modifies user information including roles, departments, and posts.
func (c *userController) Update(ctx context.Context, in *system.UpdateUserReq) (out *system.UpdateUserRes, err error) {
	out = &system.UpdateUserRes{}
	_, err = service.SystemUser().UpdateSimple(ctx, &in.SystemUserUpdate)
	if utils.IsError(err) {
		return
	}
	return
}

// InitUserPassword resets a user's password to the initial default password.
func (c *userController) InitUserPassword(ctx context.Context, in *system.InitUserPasswordReq) (out *system.InitUserPasswordRes, err error) {
	out = &system.InitUserPasswordRes{}
	password := config.GetConfigString(ctx, "settings.initPassword", "123456")
	_, err = service.SystemUser().InitUserPassword(ctx, in.Id, password)
	if utils.IsError(err) {
		return
	}
	return
}

// DeleteUser soft deletes users from the system.
func (c *userController) DeleteUser(ctx context.Context, in *system.DeleteUserReq) (out *system.DeleteUserRes, err error) {
	out = &system.DeleteUserRes{}
	err = service.SystemUser().Delete(ctx, in.Ids)
	if utils.IsError(err) {
		return
	}
	return
}

// RealDeleteUser permanently removes users from the system.
func (c *userController) RealDeleteUser(ctx context.Context, in *system.RealDeleteUserReq) (out *system.RealDeleteUserRes, err error) {
	out = &system.RealDeleteUserRes{}
	err = service.SystemUser().RealDelete(ctx, in.Ids)
	if utils.IsError(err) {
		return
	}
	return
}

// RecoveryUser restores soft-deleted users.
func (c *userController) RecoveryUser(ctx context.Context, in *system.RecoveryUserReq) (out *system.RecoveryUserRes, err error) {
	out = &system.RecoveryUserRes{}
	err = service.SystemUser().Recovery(ctx, in.Ids)
	if utils.IsError(err) {
		return
	}
	return
}

// ChangeStatusUser changes the status of a user and kicks them if disabled.
func (c *userController) ChangeStatusUser(ctx context.Context, in *system.ChangeStatusUserReq) (out *system.ChangeStatusUserRes, err error) {
	out = &system.ChangeStatusUserRes{}
	err = service.SystemUser().ChangeStatus(ctx, in.Id, in.Status)
	if utils.IsError(err) {
		return
	}
	if in.Status == 2 {
		r := request.GetHttpRequest(ctx)
		err = service.Token().KickAll(r, in.Id)
		if utils.IsError(err) {
			return
		}
	}

	return
}
