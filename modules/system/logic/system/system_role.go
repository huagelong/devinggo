// Package system
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package system

import (
	"context"
	"devinggo/internal/dao"
	"devinggo/internal/model/do"
	"devinggo/internal/model/entity"
	"devinggo/modules/system/consts"
	"devinggo/modules/system/logic/base"
	"devinggo/modules/system/model"
	"devinggo/modules/system/model/req"
	"devinggo/modules/system/model/res"
	"devinggo/modules/system/myerror"
	"devinggo/modules/system/pkg/contexts"
	"devinggo/modules/system/pkg/handler"
	"devinggo/modules/system/pkg/hook"
	"devinggo/modules/system/pkg/orm"
	"devinggo/modules/system/pkg/utils"
	"devinggo/modules/system/pkg/utils/slice"
	"devinggo/modules/system/service"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
)

type sSystemRole struct {
	base.BaseService
}

func init() {
	service.RegisterSystemRole(NewSystemRole())
}

func NewSystemRole() *sSystemRole {
	return &sSystemRole{}
}

func (s *sSystemRole) Model(ctx context.Context) *gdb.Model {
	return dao.SystemRole.Ctx(ctx).Hook(hook.Bind()).Cache(orm.SetCacheOption(ctx)).OnConflict("id")
}

func (s *sSystemRole) GetByIds(ctx context.Context, ids []int64) (res []*entity.SystemRole, err error) {
	err = s.Model(ctx).WhereIn(dao.SystemRole.Columns().Id, ids).Scan(&res)
	return
}

func (s *sSystemRole) Verify(r *ghttp.Request) bool {
	ctx := r.GetCtx()
	var (
		userId = contexts.New().GetUserId(ctx)
		err    error
	)

	if g.IsEmpty(userId) {
		g.Log().Debug(ctx, "system Verify user = nil")
		return false
	}

	isSuperAdmin, err := service.SystemUser().IsSuperAdmin(ctx, userId)
	if err != nil {
		g.Log().Debug(ctx, "get isSuperAdmin error", err)
		return false
	}
	if isSuperAdmin {
		return true
	}
	permission := contexts.New().GetPermission(ctx)
	if g.IsEmpty(permission) {
		g.Log().Debug(ctx, "permission is nil")
		return false
	}
	roleIds, err := service.SystemUser().GetRoles(ctx, userId)
	if err != nil {
		g.Log().Debug(ctx, "GetRoles error", err)
		return false
	}
	menuIds, err := service.SystemRoleMenu().GetMenuIdsByRoleIds(ctx, roleIds)
	if err != nil {
		g.Log().Debug(ctx, "menuIds error", err)
		return false
	}
	var systemMenuEntity *entity.SystemMenu
	systemMenuEntity, err = service.SystemMenu().GetMenuByPermission(ctx, permission, menuIds)
	if err != nil {
		g.Log().Debug(ctx, "systemMenuEntity error", err)
		return false
	}
	//g.Log().Debug(ctx, "systemMenuEntity:", systemMenuEntity)
	if !g.IsEmpty(systemMenuEntity) {
		return true
	}

	return false
}

func (s *sSystemRole) handleRoleSearch(ctx context.Context, in *req.SystemRoleSearch, filterAdminRole bool) (m *gdb.Model) {

	m = s.Model(ctx)

	if !g.IsEmpty(in.Code) {
		m = m.Where("code", in.Code)
	}

	if !g.IsEmpty(in.Status) {
		m = m.Where("status", in.Status)
	}

	if !g.IsEmpty(in.Name) {
		m = m.Where("name like ? ", "%"+in.Name+"%")
	}
	if filterAdminRole {
		m = m.WhereNot("code", consts.SuperRoleKey)
	}
	if !g.IsEmpty(in.CreatedAt) {
		if len(in.CreatedAt) > 0 {
			m = m.WhereGTE("created_at", in.CreatedAt[0]+" 00:00:00")
		}
		if len(in.CreatedAt) > 1 {
			m = m.WhereLTE("created_at", in.CreatedAt[1]+"23:59:59")
		}
	}
	return
}

func (s *sSystemRole) GetList(ctx context.Context, in *req.SystemRoleSearch, filterAdminRole bool) (out []*res.SystemRole, err error) {
	inReq := &model.ListReq{
		OrderBy:   "sort",
		OrderType: "desc",
	}
	m := s.handleRoleSearch(ctx, in, filterAdminRole)
	m = orm.GetList(m, inReq)
	err = m.Scan(&out)
	if utils.IsError(err) {
		return
	}
	return
}

func (s *sSystemRole) GetPageList(ctx context.Context, req *model.PageListReq, in *req.SystemRoleSearch, filterAdminRole bool) (rs []*res.SystemRole, total int, err error) {
	m := s.handleRoleSearch(ctx, in, filterAdminRole).Handler(handler.FilterAuth)
	var postEntity []*entity.SystemRole
	err = orm.GetPageList(m, req).ScanAndCount(&postEntity, &total, false)
	if utils.IsError(err) {
		return nil, 0, err
	}
	rs = make([]*res.SystemRole, 0)
	if !g.IsEmpty(postEntity) {
		if err = gconv.Structs(postEntity, &rs); err != nil {
			return nil, 0, err
		}
	}
	return
}

func (s *sSystemRole) Save(ctx context.Context, in *req.SystemRoleSave) (id int64, err error) {
	if s.checkRoleCode(ctx, in.Code) {
		return 0, myerror.ValidationFailed(ctx, "角色标识已存在")
	}
	saveData := do.SystemRole{
		Name:   in.Name,
		Sort:   in.Sort,
		Status: in.Status,
		Code:   in.Code,
		Remark: in.Remark,
	}
	rs, err := s.Model(ctx).Data(saveData).Insert()
	if utils.IsError(err) {
		return
	}
	tmpId, err := rs.LastInsertId()
	if err != nil {
		return
	}
	id = gconv.Int64(tmpId)

	superAdminId, _ := s.GetSuperAdminId(ctx)
	if id == superAdminId {
		return
	}
	if !g.IsEmpty(in.MenuIds) {
		for _, menuId := range in.MenuIds {
			_, err = service.SystemRoleMenu().Model(ctx).Data(do.SystemRoleMenu{
				RoleId: id,
				MenuId: menuId,
			}).Save()
		}
	}

	if !g.IsEmpty(in.DeptIds) {
		for _, deptId := range in.DeptIds {
			_, err = service.SystemRoleDept().Model(ctx).Data(do.SystemRoleDept{
				RoleId: id,
				DeptId: deptId,
			}).Save()
		}
	}
	return
}

func (s *sSystemRole) checkRoleCode(ctx context.Context, code string) bool {
	count, err := s.Model(ctx).Where("code", code).Count()
	if utils.IsError(err) {
		return true
	}
	if count > 0 {
		return true
	}
	return false
}

func (s *sSystemRole) GetSuperAdminId(ctx context.Context) (id int64, err error) {
	var role *entity.SystemRole
	err = s.Model(ctx).Where("code", consts.SuperRoleKey).Scan(&role)
	if utils.IsError(err) {
		return
	}
	id = role.Id
	return
}

func (s *sSystemRole) GetById(ctx context.Context, id int64) (res *res.SystemRole, err error) {
	err = s.Model(ctx).Where("id", id).Scan(&res)
	if utils.IsError(err) {
		return
	}
	return
}

func (s *sSystemRole) Update(ctx context.Context, in *req.SystemRoleSave) (err error) {
	updateData := do.SystemRole{
		Name:      in.Name,
		DataScope: in.DataScope,
		Sort:      in.Sort,
		Status:    in.Status,
		Code:      in.Code,
		Remark:    in.Remark,
	}
	_, err = s.Model(ctx).Data(updateData).OmitEmptyData().Where("id", in.Id).Update()
	if utils.IsError(err) {
		return
	}
	id := in.Id

	superAdminId, _ := s.GetSuperAdminId(ctx)
	if id == superAdminId {
		return
	}
	if !g.IsEmpty(in.MenuIds) {
		service.SystemRoleMenu().Model(ctx).Where("role_id", id).Delete()
		for _, menuId := range in.MenuIds {
			_, err = service.SystemRoleMenu().Model(ctx).Data(do.SystemRoleMenu{
				RoleId: id,
				MenuId: menuId,
			}).Save()
		}
	}

	if !g.IsEmpty(in.DeptIds) {
		service.SystemRoleDept().Model(ctx).Where("role_id", id).Delete()
		for _, deptId := range in.DeptIds {
			_, err = service.SystemRoleDept().Model(ctx).Data(do.SystemRoleDept{
				RoleId: id,
				DeptId: deptId,
			}).Save()
		}
	}

	return
}

func (s *sSystemRole) Delete(ctx context.Context, ids []int64) (err error) {
	superAdminId, err := s.GetSuperAdminId(ctx)
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

func (s *sSystemRole) RealDelete(ctx context.Context, ids []int64) (err error) {
	superAdminId, err := s.GetSuperAdminId(ctx)
	if err != nil {
		return
	}
	for _, id := range ids {
		if id == superAdminId {
			continue
		}
		s.Model(ctx).Unscoped().Where("id", id).Delete()
		service.SystemRoleMenu().Model(ctx).Where("role_id", id).Delete()
		service.SystemRoleDept().Model(ctx).Where("role_id", id).Delete()
		service.SystemUserRole().Model(ctx).Where("role_id", id).Delete()
	}
	return
}

func (s *sSystemRole) Recovery(ctx context.Context, ids []int64) (err error) {
	_, err = s.Model(ctx).Unscoped().WhereIn("id", ids).Update(g.Map{"deleted_at": nil})
	if utils.IsError(err) {
		return err
	}
	return
}

func (s *sSystemRole) ChangeStatus(ctx context.Context, id int64, status int) (err error) {
	_, err = s.Model(ctx).Data(g.Map{"status": status}).Where("id", id).Update()
	if utils.IsError(err) {
		return err
	}
	return
}

func (s *sSystemRole) NumberOperation(ctx context.Context, id int64, numberName string, numberValue int) (err error) {
	_, err = s.Model(ctx).Data(g.Map{numberName: numberValue}).Where("id", id).Update()
	if utils.IsError(err) {
		return err
	}
	return
}

func (s *sSystemRole) GetMenuByRoleIds(ctx context.Context, ids []int64) (out []*res.SystemRoleMenus, err error) {
	for _, id := range ids {
		systemRoleMenus := &res.SystemRoleMenus{}
		var roleMenuEntity []*entity.SystemRoleMenu
		err = service.SystemRoleMenu().Model(ctx).Where("role_id", id).Scan(&roleMenuEntity)
		if utils.IsError(err) {
			return
		}

		systemRoleMenus.Id = id
		systemRoleMenus.Menus = make([]res.MenuIdsArr, 0)
		if !g.IsEmpty(roleMenuEntity) {
			for _, roleMenu := range roleMenuEntity {
				menuIdsArr := &res.MenuIdsArr{}
				menuIdsArr.Id = roleMenu.MenuId
				menuIdsArr.Pivot.MenuId = roleMenu.MenuId
				menuIdsArr.Pivot.RoleId = roleMenu.RoleId
				systemRoleMenus.Menus = append(systemRoleMenus.Menus, *menuIdsArr)
			}
		}
		out = append(out, systemRoleMenus)
	}

	return
}

func (s *sSystemRole) GetDeptByRole(ctx context.Context, ids []int64) (out []*res.SystemRoleDepts, err error) {
	for _, id := range ids {
		systemRoleDepts := &res.SystemRoleDepts{}
		var roleDeptEntity []*entity.SystemRoleDept
		err = service.SystemRoleDept().Model(ctx).Where("role_id", id).Scan(&roleDeptEntity)
		if utils.IsError(err) {
			return
		}

		systemRoleDepts.Id = id
		systemRoleDepts.Depts = make([]res.DeptIdsArr, 0)
		if g.IsEmpty(roleDeptEntity) {
			continue
		}
		for _, roleMenu := range roleDeptEntity {
			deptIdsArr := &res.DeptIdsArr{}
			deptIdsArr.Id = roleMenu.DeptId
			deptIdsArr.Pivot.DeptId = roleMenu.DeptId
			deptIdsArr.Pivot.RoleId = roleMenu.RoleId
			systemRoleDepts.Depts = append(systemRoleDepts.Depts, *deptIdsArr)
		}
		out = append(out, systemRoleDepts)
	}

	return
}
