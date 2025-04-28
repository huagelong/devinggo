// Package system
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package system

import (
	"context"
	"database/sql"
	"devinggo/internal/dao"
	"devinggo/internal/model/do"
	"devinggo/modules/system/consts"
	"devinggo/modules/system/logic/base"
	"devinggo/modules/system/model"
	"devinggo/modules/system/model/req"
	"devinggo/modules/system/model/res"
	"devinggo/modules/system/myerror"
	"devinggo/modules/system/pkg/hook"
	"devinggo/modules/system/pkg/orm"
	"devinggo/modules/system/pkg/utils"
	"devinggo/modules/system/pkg/utils/config"
	"devinggo/modules/system/pkg/utils/request"
	"devinggo/modules/system/pkg/utils/secure"
	"devinggo/modules/system/pkg/utils/slice"
	"devinggo/modules/system/service"
	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

type sSystemUser struct {
	base.BaseService
}

func init() {
	service.RegisterSystemUser(NewSystemUser())
}

func NewSystemUser() *sSystemUser {
	return &sSystemUser{}
}

func (s *sSystemUser) Model(ctx context.Context) *gdb.Model {
	return dao.SystemUser.Ctx(ctx).Hook(hook.Bind()).Cache(orm.SetCacheOption(ctx)).OnConflict("id")
}

func (s *sSystemUser) GetPageList(ctx context.Context, req *model.PageListReq) (res []*res.SystemUser, total int, err error) {
	err = orm.GetPageList(s.Model(ctx), req).ScanAndCount(&res, &total, false)
	if utils.IsError(err) {
		return nil, 0, err
	}
	return
}

func (s *sSystemUser) GetPageListForSearch(ctx context.Context, req *model.PageListReq, in *req.SystemUserSearch) (res []*res.SystemUser, total int, err error) {
	m := s.handleUserSearch(ctx, in)
	err = orm.GetPageList(m, req).ScanAndCount(&res, &total, false)
	if utils.IsError(err) {
		return nil, 0, err
	}
	return
}

func (s *sSystemUser) GetOnlineUserPageListForSearch(ctx context.Context, req *model.PageListReq, in *req.SystemUserSearch) (res []*res.SystemUser, total int, err error) {
	m := s.handleUserSearch(ctx, in)
	r := request.GetHttpRequest(ctx)
	userApps, err := service.Token().GetAllUserIds(r)
	if err != nil {
		return nil, 0, err
	}
	if g.IsEmpty(userApps) {
		return nil, 0, err
	}
	userIds := make([]int64, 0)
	userAppMap := make(map[int64]string)
	for _, userApp := range userApps {
		userIds = append(userIds, userApp.UserId)
		userAppMap[userApp.UserId] = userApp.AppId
	}
	m = m.WhereIn("id", userIds)
	err = orm.GetPageList(m, req).ScanAndCount(&res, &total, false)
	if utils.IsError(err) {
		return nil, 0, err
	}
	if !g.IsEmpty(res) {
		for _, user := range res {
			user.AppId = userAppMap[user.Id]
		}
	}
	return
}

func (s *sSystemUser) GetExportList(ctx context.Context, req *model.ListReq, in *req.SystemUserSearch) (res []*res.SystemUserExport, err error) {
	m := s.handleUserSearch(ctx, in)
	err = orm.GetList(m, req).Scan(&res)
	if utils.IsError(err) {
		return
	}
	return
}

func (s *sSystemUser) GetSupserAdminId(ctx context.Context) int64 {
	return config.GetConfigint64(ctx, "settings.superAdminId", 1)
}

func (s *sSystemUser) ExistsByUsername(ctx context.Context, username string) (rs bool, err error) {
	count, err := s.Model(ctx).Where(dao.SystemUser.Columns().Username, username).Count()
	if utils.IsError(err) {
		return false, err
	}
	return count > 0, err
}

func (s *sSystemUser) handleUserSearch(ctx context.Context, in *req.SystemUserSearch) (m *gdb.Model) {

	m = s.Model(ctx)
	if !g.IsEmpty(in.Status) {
		m = m.Where(dao.SystemUser.Table()+".status", in.Status)
	}

	if !g.IsEmpty(in.Phone) {
		m = m.Where(dao.SystemUser.Table()+".phone", in.Phone)
	}

	if !g.IsEmpty(in.Username) {
		m = m.Where(dao.SystemUser.Table()+".username like ? ", "%"+in.Username+"%")
	}

	if !g.IsEmpty(in.Nickname) {
		m = m.Where(dao.SystemUser.Table()+".nickname like ? ", "%"+in.Nickname+"%")
	}

	if !g.IsEmpty(in.Username) && in.FilterSuperAdmin {
		supserAdminId := s.GetSupserAdminId(ctx)
		m = m.WhereNot(dao.SystemUser.Table()+".id", supserAdminId)
	}
	if !g.IsEmpty(in.CreatedAt) {
		if len(in.CreatedAt) > 0 {
			m = m.WhereGTE(dao.SystemUser.Table()+".created_at", in.CreatedAt[0]+" 00:00:00")
		}
		if len(in.CreatedAt) > 1 {
			m = m.WhereLTE(dao.SystemUser.Table()+".created_at", in.CreatedAt[1]+"23:59:59")
		}
	}

	if !g.IsEmpty(in.UserIds) {
		m = m.WhereIn(dao.SystemUser.Table()+".id", in.UserIds)
	}

	if !g.IsEmpty(in.RoleId) {
		m = m.LeftJoinOnFields(dao.SystemUserRole.Table(), dao.SystemUser.Columns().Id, "=", dao.SystemUserRole.Columns().UserId).Where(dao.SystemUserRole.Table()+".role_id", in.RoleId)
	}

	if !g.IsEmpty(in.PostId) {
		m = m.LeftJoinOnFields(dao.SystemUserPost.Table(), dao.SystemUser.Columns().Id, "=", dao.SystemUserPost.Columns().UserId).Where(dao.SystemUserPost.Table()+".post_id", in.RoleId)
	}

	if !g.IsEmpty(in.DeptId) {
		result, err := service.SystemDept().Model(ctx).Fields(dao.SystemDept.Columns().Id).Where(dao.SystemDept.Columns().Id, in.DeptId).WhereOr("level like  ? ", "%,"+gconv.String(in.DeptId)+",%").Array()
		if !utils.IsError(err) {
			if !g.IsEmpty(result) {
				newDeptIds := gconv.SliceInt64(result)
				m = m.LeftJoinOnFields(dao.SystemUserDept.Table(), dao.SystemUser.Columns().Id, "=", dao.SystemUserDept.Columns().UserId).WhereIn(dao.SystemUserDept.Columns().DeptId, newDeptIds)
			}
		}
	}

	return
}

func (s *sSystemUser) GetInfoById(ctx context.Context, userId int64) (systemUser *res.SystemUser, err error) {
	err = s.Model(ctx).Where(dao.SystemUser.Columns().Id, userId).Scan(&systemUser)
	if utils.IsError(err) {
		return
	}
	return
}

func (s *sSystemUser) GetInfoByIds(ctx context.Context, userIds []int64) (systemUser []*res.SystemUser, err error) {
	err = s.Model(ctx).WhereIn(dao.SystemUser.Columns().Id, userIds).Scan(&systemUser)
	if utils.IsError(err) {
		return
	}
	return
}

func (s *sSystemUser) GetInfo(ctx context.Context, userId int64) (systemUserInfo *res.SystemUserInfo, err error) {
	systemUser := res.SystemUser{}
	systemUserInfo = &res.SystemUserInfo{}
	err = s.Model(ctx).Where(dao.SystemUser.Columns().Id, userId).Scan(&systemUser)
	if utils.IsError(err) {
		return
	}

	if g.IsEmpty(systemUser) {
		return nil, nil
	}

	systemUserInfo.User = systemUser
	isSuperAdmin, err := s.IsSuperAdmin(ctx, userId)
	if err != nil {
		return
	}
	if isSuperAdmin {
		systemUserInfo.Roles = gconv.Strings(garray.NewArray(true).Append("superAdmin"))
		systemUserInfo.Codes = gconv.Strings(garray.NewArray(true).Append("*"))
		superAdminRouters, err := service.SystemMenu().GetSuperAdminRouters(ctx)
		if err != nil {
			return systemUserInfo, err
		}
		systemUserInfo.Routers = superAdminRouters
	} else {
		roleCodes := make([]string, 0)
		menuCodes := make([]string, 0)
		routers := make([]*res.Router, 0)
		roleIds, err := service.SystemUser().GetRoles(ctx, userId)
		if err != nil {
			return systemUserInfo, err
		}
		if !g.IsEmpty(roleIds) {
			systemRoles, err := service.SystemRole().GetByIds(ctx, roleIds)
			if err != nil {
				return systemUserInfo, err
			}

			if !g.IsEmpty(systemRoles) {
				for _, role := range systemRoles {
					roleCodes = append(roleCodes, role.Code)
				}
			}
			menuIds, err := service.SystemRoleMenu().GetMenuIdsByRoleIds(ctx, roleIds)
			if err != nil {
				return systemUserInfo, err
			}
			menuCodes, err = service.SystemMenu().GetMenuCode(ctx, menuIds)
			if err != nil {
				return systemUserInfo, err
			}

			routers, err = service.SystemMenu().GetRoutersByIds(ctx, menuIds)
			if err != nil {
				return systemUserInfo, err
			}
		}
		systemUserInfo.Roles = roleCodes
		systemUserInfo.Codes = menuCodes
		systemUserInfo.Routers = routers
	}
	return
}

func (s *sSystemUser) IsSuperAdmin(ctx context.Context, userId int64) (isSuperAdmin bool, err error) {
	roleIds, err := service.SystemUser().GetRoles(ctx, userId)
	if err != nil {
		return false, err
	}
	if !g.IsEmpty(roleIds) {
		roles, err := service.SystemRole().GetByIds(ctx, roleIds)
		if err != nil {
			return false, err
		}
		if !g.IsEmpty(roles) {
			for _, role := range roles {
				if role.Code == consts.SuperRoleKey {
					return true, nil
				}
			}
		}
	}
	return false, nil
}

func (s *sSystemUser) GetRoles(ctx context.Context, userId int64) (roles []int64, err error) {
	result, err := service.SystemUserRole().Model(ctx).Fields(dao.SystemUserRole.Columns().RoleId).Where(dao.SystemUserRole.Columns().UserId, userId).Array()
	if utils.IsError(err) {
		return
	}

	if g.IsEmpty(result) {
		return
	}

	roles = gconv.SliceInt64(result)
	return
}

func (s *sSystemUser) GetDepts(ctx context.Context, userId int64) (depts []int64, err error) {
	result, err := service.SystemUserDept().Model(ctx).Fields(dao.SystemUserDept.Columns().DeptId).Where(dao.SystemUserDept.Columns().UserId, userId).Array()
	if utils.IsError(err) {
		return
	}

	if g.IsEmpty(result) {
		return
	}

	depts = gconv.SliceInt64(result)
	return
}

func (s *sSystemUser) Update(ctx context.Context, req *req.SystemUser, userId ...int64) (rs sql.Result, err error) {
	var systemUser *do.SystemUser
	if err = gconv.Struct(req, &systemUser); err != nil {
		return
	}
	if g.IsEmpty(req.Id) {
		if len(userId) > 0 {
			systemUser.Id = userId[0]
		} else {
			g.Log().Warning(ctx, "system user update failed, user id is empty")
			return rs, myerror.ValidationFailed(ctx, "用户id为空")
		}
	}

	rs, err = s.Model(ctx).OmitEmptyData().Data(systemUser).Where(dao.SystemUser.Columns().Id, systemUser.Id).Update()
	if utils.IsError(err) {
		return
	}
	return
}

func (s *sSystemUser) SetHomePage(ctx context.Context, id int64, dashboard string) (out sql.Result, err error) {
	systemUser := &do.SystemUser{
		Dashboard: dashboard,
	}
	out, err = s.Model(ctx).Data(systemUser).Where(dao.SystemUser.Columns().Id, id).Update()
	if utils.IsError(err) {
		return
	}
	return
}

func (s *sSystemUser) InitUserPassword(ctx context.Context, id int64, password string) (out sql.Result, err error) {
	password, err = secure.PasswordHash(password)
	if err != nil {
		return
	}
	systemUser := &do.SystemUser{
		Password: password,
	}
	out, err = s.Model(ctx).Data(systemUser).Where(dao.SystemUser.Columns().Id, id).Update()
	if utils.IsError(err) {
		return
	}
	return
}

func (s *sSystemUser) UpdateSimple(ctx context.Context, in *req.SystemUserUpdate) (out sql.Result, err error) {
	if g.IsEmpty(in.Id) {
		err = myerror.MissingParameter(ctx, "用户id为空")
		return
	}

	var systemUser *do.SystemUser
	if err = gconv.Struct(in, &systemUser); err != nil {
		return
	}

	out, err = s.Model(ctx).OmitEmptyData().Data(systemUser).Where(dao.SystemUser.Columns().Id, in.Id).Update()
	if utils.IsError(err) {
		return
	}
	id := in.Id
	if !g.IsEmpty(in.RoleIds) {
		service.SystemUserRole().Model(ctx).Where(dao.SystemUserRole.Columns().UserId, id).Delete()
		for _, roleId := range in.RoleIds {
			_, err = service.SystemUserRole().Model(ctx).Data(do.SystemUserRole{
				RoleId: roleId,
				UserId: id,
			}).Save()
		}
	}

	if !g.IsEmpty(in.DeptIds) {
		service.SystemUserDept().Model(ctx).Where(dao.SystemUserDept.Columns().UserId, id).Delete()
		for _, deptId := range in.DeptIds {
			_, err = service.SystemUserDept().Model(ctx).Data(do.SystemUserDept{
				UserId: id,
				DeptId: deptId,
			}).Save()
		}
	}

	if !g.IsEmpty(in.PostIds) {
		service.SystemUserPost().Model(ctx).Where(dao.SystemUserPost.Columns().UserId, id).Delete()
		for _, postId := range in.PostIds {
			_, err = service.SystemUserPost().Model(ctx).Data(do.SystemUserPost{
				UserId: id,
				PostId: postId,
			}).Save()
		}
	} else {
		service.SystemUserPost().Model(ctx).Where(dao.SystemUserPost.Columns().UserId, id).Delete()
	}

	return
}

func (s *sSystemUser) Save(ctx context.Context, in *req.SystemUserSave) (id int64, err error) {
	userNameExists, err := service.SystemUser().ExistsByUsername(ctx, in.Username)
	if err != nil {
		return
	}
	if userNameExists {
		err = myerror.ValidationFailed(ctx, "用户名已存在")
		return
	}
	if !g.IsEmpty(in.RoleIds) {
		supserAdminId := s.GetSupserAdminId(ctx)
		if slice.Contains(in.RoleIds, supserAdminId) {
			in.RoleIds = slice.Remove(in.RoleIds, supserAdminId)
		}
	}

	var systemUser *do.SystemUser
	if err = gconv.Struct(in, &systemUser); err != nil {
		return
	}

	if !g.IsEmpty(in.Password) {
		newPassword, err := secure.PasswordHash(in.Password)
		if err != nil {
			return 0, err
		}
		systemUser.Password = newPassword
	}
	rs, err := s.Model(ctx).Data(systemUser).Insert()
	if utils.IsError(err) {
		return
	}

	tmpId, err := rs.LastInsertId()
	if err != nil {
		return
	}
	id = gconv.Int64(tmpId)

	if !g.IsEmpty(in.RoleIds) {
		for _, roleId := range in.RoleIds {
			_, err = service.SystemUserRole().Model(ctx).Data(do.SystemUserRole{
				RoleId: roleId,
				UserId: id,
			}).Save()
		}
	}

	if !g.IsEmpty(in.DeptIds) {
		for _, deptId := range in.DeptIds {
			_, err = service.SystemUserDept().Model(ctx).Data(do.SystemUserDept{
				UserId: id,
				DeptId: deptId,
			}).Save()
		}
	}

	if !g.IsEmpty(in.PostIds) {
		for _, postId := range in.PostIds {
			_, err = service.SystemUserPost().Model(ctx).Data(do.SystemUserPost{
				UserId: id,
				PostId: postId,
			}).Save()
		}
	}

	return
}

func (s *sSystemUser) GetFullInfoById(ctx context.Context, id int64) (out *res.SystemUserFullInfo, err error) {
	err = s.Model(ctx).Where("id", id).Scan(&out)
	if utils.IsError(err) {
		return
	}
	if !g.IsEmpty(out) {
		var roleList = make([]*res.SystemRole, 0)
		err = service.SystemRole().Model(ctx).InnerJoinOnFields(dao.SystemUserRole.Table(), dao.SystemRole.Columns().Id, "=", dao.SystemUserRole.Columns().RoleId).Where(dao.SystemRole.Columns().Status, 1).WherePrefix(dao.SystemUserRole.Table(), dao.SystemUserRole.Columns().UserId, id).Scan(&roleList)
		if utils.IsError(err) {
			return
		}
		out.RoleList = roleList

		var postList = make([]*res.SystemPost, 0)
		err = service.SystemPost().Model(ctx).InnerJoinOnFields(dao.SystemUserPost.Table(), dao.SystemPost.Columns().Id, "=", dao.SystemUserPost.Columns().PostId).Where(dao.SystemPost.Columns().Status, 1).WherePrefix(dao.SystemUserPost.Table(), dao.SystemUserPost.Columns().UserId, id).Scan(&postList)
		if utils.IsError(err) {
			return
		}
		out.PostList = postList

		var deptList = make([]*res.SystemDept, 0)
		err = service.SystemDept().Model(ctx).InnerJoinOnFields(dao.SystemUserDept.Table(), dao.SystemDept.Columns().Id, "=", dao.SystemUserDept.Columns().DeptId).Where(dao.SystemDept.Columns().Status, 1).WherePrefix(dao.SystemUserDept.Table(), dao.SystemUserDept.Columns().UserId, id).Scan(&deptList)
		if utils.IsError(err) {
			return
		}
		out.DeptList = deptList
	}

	return
}

func (s *sSystemUser) Delete(ctx context.Context, ids []int64) (err error) {
	superAdminId := s.GetSupserAdminId(ctx)
	if err != nil {
		return
	}
	newIds := slice.Remove(ids, superAdminId)
	if !g.IsEmpty(newIds) {
		_, err = s.Model(ctx).WhereIn("id", ids).Delete()
		if utils.IsError(err) {
			return err
		}
	}
	return
}

func (s *sSystemUser) RealDelete(ctx context.Context, ids []int64) (err error) {
	superAdminId := s.GetSupserAdminId(ctx)
	if err != nil {
		return
	}
	newIds := slice.Remove(ids, superAdminId)
	for _, id := range newIds {
		s.Model(ctx).Unscoped().Where("id", id).Delete()
		service.SystemUserPost().Model(ctx).Where("user_id", id).Delete()
		service.SystemUserDept().Model(ctx).Where("user_id", id).Delete()
		service.SystemUserRole().Model(ctx).Where("user_id", id).Delete()
	}
	return
}

func (s *sSystemUser) Recovery(ctx context.Context, ids []int64) (err error) {
	_, err = s.Model(ctx).Unscoped().WhereIn("id", ids).Update(g.Map{"deleted_at": nil})
	if utils.IsError(err) {
		return err
	}
	return
}

func (s *sSystemUser) ChangeStatus(ctx context.Context, id int64, status int) (err error) {
	_, err = s.Model(ctx).Data(g.Map{"status": status}).Where("id", id).Update()
	if utils.IsError(err) {
		return err
	}
	return
}
