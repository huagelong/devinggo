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
	"devinggo/modules/system/logic/base"
	"devinggo/modules/system/model"
	"devinggo/modules/system/model/req"
	"devinggo/modules/system/model/res"
	"devinggo/modules/system/pkg/handler"
	"devinggo/modules/system/pkg/hook"
	"devinggo/modules/system/pkg/orm"
	"devinggo/modules/system/pkg/utils"
	"devinggo/modules/system/pkg/utils/slice"
	"devinggo/modules/system/service"
	"fmt"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

type sSystemDept struct {
	base.BaseService
}

func init() {
	service.RegisterSystemDept(NewSystemDept())
}

func NewSystemDept() *sSystemDept {
	return &sSystemDept{}
}

func (s *sSystemDept) Model(ctx context.Context) *gdb.Model {
	return dao.SystemDept.Ctx(ctx).Hook(hook.Bind()).Cache(orm.SetCacheOption(ctx)).OnConflict("id")
}

func (s *sSystemDept) GetSelectTree(ctx context.Context, userId int64) (tree []*res.SystemDeptTree, err error) {
	systemDeptEntity := []entity.SystemDept{}
	err = s.Model(ctx).Handler(handler.FilterAuth).Where(dao.SystemMenu.Columns().Status, 1).Order("parent_id, sort desc").Scan(&systemDeptEntity)
	if utils.IsError(err) {
		return
	}

	if !g.IsEmpty(systemDeptEntity) {
		systemDeptEntity2 := []entity.SystemDept{}
		deptIds := make([]int64, 0)
		result, err := service.SystemUserDept().Model(ctx).Fields(dao.SystemUserDept.Columns().DeptId).Where(dao.SystemUserDept.Columns().UserId, userId).Array()
		if utils.IsError(err) {
			return nil, err
		}
		if !g.IsEmpty(result) {
			deptIds = gconv.SliceInt64(result)
			err = s.Model(ctx).Where(dao.SystemDept.Columns().Status, 1).WhereIn(dao.SystemDept.Columns().Id, deptIds).Order("parent_id, sort desc").Scan(&systemDeptEntity2)
			if utils.IsError(err) {
				return nil, err
			}
		}
		if !g.IsEmpty(systemDeptEntity2) {
			systemDeptEntity = utils.MergeAndDeduplicateWithFunc[entity.SystemDept](s.compareFunc, systemDeptEntity, systemDeptEntity2)
		}
	}

	if g.IsEmpty(systemDeptEntity) {
		return
	}
	tree = s.treeList(systemDeptEntity)
	return
}

func (s *sSystemDept) compareFunc(item entity.SystemDept) string {
	return gconv.String(item.Id)
}

func (s *sSystemDept) treeList(nodes []entity.SystemDept) (tree []*res.SystemDeptTree) {
	type itemTree map[int64]*res.SystemDeptTree
	itemList := make(itemTree)
	for _, systemDeptEntity := range nodes {
		var item res.SystemDeptTree
		item.Id = systemDeptEntity.Id
		item.ParentId = systemDeptEntity.ParentId
		item.Value = systemDeptEntity.Id
		item.Label = systemDeptEntity.Name
		item.Children = make([]*res.SystemDeptTree, 0)
		if !g.IsEmpty(itemList[item.ParentId]) {
			itemList[item.ParentId].Children = append(itemList[item.ParentId].Children, &item)
		} else {
			tree = append(tree, &item)
		}
		itemList[systemDeptEntity.Id] = &item
	}
	return
}

func (s *sSystemDept) handleSearch(ctx context.Context, in *req.SystemDeptSearch) (m *gdb.Model) {
	m = s.Model(ctx)

	if !g.IsEmpty(in.Status) {
		m = m.Where("status", in.Status)
	}

	if !g.IsEmpty(in.Leader) {
		m = m.Where("leader", in.Leader)
	}

	if !g.IsEmpty(in.Phone) {
		m = m.Where("phone", in.Phone)
	}

	if !g.IsEmpty(in.Name) {
		m = m.Where("name like ? ", "%"+in.Name+"%")
	}

	if !g.IsEmpty(in.Level) {
		m = m.Where("level like ? ", "%,"+in.Level+",%")
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

func (s *sSystemDept) GetListTreeList(ctx context.Context, in *req.SystemDeptSearch) (tree []*res.SystemListDeptTree, err error) {
	inReq := &model.ListReq{
		Recycle: in.Recycle,
	}
	m := s.handleSearch(ctx, in)
	m = orm.GetList(m, inReq)
	systemDeptEntity := []entity.SystemDept{}
	err = m.Order("parent_id, sort desc").Scan(&systemDeptEntity)
	if utils.IsError(err) {
		return
	}
	if g.IsEmpty(systemDeptEntity) {
		return
	}
	tree = s.treeList2(systemDeptEntity)
	return
}

func (s *sSystemDept) GetRecycleTreeList(ctx context.Context, in *req.SystemDeptSearch) (tree []*res.SystemListDeptTree, err error) {
	inReq := &model.ListReq{
		Recycle: in.Recycle,
	}
	m := s.handleSearch(ctx, in)
	m = orm.GetList(m, inReq)
	systemDeptEntity := []entity.SystemDept{}
	err = m.Order("parent_id, sort desc").Scan(&systemDeptEntity)
	if utils.IsError(err) {
		return
	}
	if g.IsEmpty(systemDeptEntity) {
		return
	}
	tree = s.treeList2(systemDeptEntity)
	return
}

func (s *sSystemDept) treeList2(nodes []entity.SystemDept) (tree []*res.SystemListDeptTree) {
	type itemTree map[int64]*res.SystemListDeptTree
	itemList := make(itemTree)
	for _, systemDeptEntity := range nodes {
		var item res.SystemListDeptTree
		item.Id = systemDeptEntity.Id
		item.ParentId = systemDeptEntity.ParentId
		item.Name = systemDeptEntity.Name
		item.Sort = systemDeptEntity.Sort
		item.Status = systemDeptEntity.Status
		item.Leader = systemDeptEntity.Leader
		item.Phone = systemDeptEntity.Phone
		item.CreatedAt = systemDeptEntity.CreatedAt
		item.Children = make([]*res.SystemListDeptTree, 0)
		if !g.IsEmpty(itemList[item.ParentId]) {
			itemList[item.ParentId].Children = append(itemList[item.ParentId].Children, &item)
		} else {
			tree = append(tree, &item)
		}
		itemList[systemDeptEntity.Id] = &item
	}
	return
}

func (s *sSystemDept) GetTreeList(ctx context.Context, in *req.SystemDeptSearch) (tree []*res.SystemDeptTree, err error) {
	inReq := &model.ListReq{
		Recycle: in.Recycle,
	}
	m := s.handleSearch(ctx, in)
	m = orm.GetList(m, inReq)
	systemDeptEntity := []entity.SystemDept{}
	err = m.Order("parent_id, sort desc").Scan(&systemDeptEntity)
	if utils.IsError(err) {
		return
	}
	if g.IsEmpty(systemDeptEntity) {
		return
	}
	tree = s.treeList(systemDeptEntity)
	return
}

func (s *sSystemDept) handleData(ctx context.Context, data *req.SystemDeptSave) (dataRes *req.SystemDeptSave, err error) {

	if g.IsEmpty(data.ParentId) {
		data.ParentId = 0
	}

	if !g.IsEmpty(data.Id) && (data.Id == data.ParentId) {
		return nil, gerror.New("id cannot be equal to parent_id")
	}

	var level string
	if data.ParentId == 0 {
		level = ",0,"
	} else {
		var parentDept *entity.SystemDept
		err = s.Model(ctx).Where(dao.SystemDept.Columns().Id, data.ParentId).Scan(&parentDept)
		if utils.IsError(err) {
			return nil, err
		}
		if !g.IsEmpty(parentDept) {
			level = fmt.Sprintf("%s%d,", parentDept.Level, data.ParentId)
		} else {
			return nil, gerror.New("parent dept not found")
		}
	}
	data.Level = level
	dataRes = data
	return
}

func (s *sSystemDept) Save(ctx context.Context, in *req.SystemDeptSave) (id int64, err error) {
	data, err := s.handleData(ctx, in)
	if err != nil {
		return
	}
	saveData := do.SystemDept{
		Name:     data.Name,
		ParentId: data.ParentId,
		Level:    data.Level,
		Sort:     data.Sort,
		Status:   data.Status,
		Leader:   data.Leader,
		Phone:    data.Phone,
		Remark:   data.Remark,
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
	return
}

func (s *sSystemDept) AddLeader(ctx context.Context, in *req.SystemDeptAddLeader) (err error) {
	users := in.Users
	deptId := in.Id
	for _, user := range users {
		count, err := service.SystemDeptLeader().Model(ctx).Where("dept_id", deptId).Where("user_id", user.Id).Count()
		if utils.IsError(err) {
			return err
		}
		if count == 0 {
			saveData := do.SystemDeptLeader{
				DeptId:    deptId,
				UserId:    user.Id,
				Username:  user.Username,
				CreatedAt: gtime.Now(),
			}
			_, err = service.SystemDeptLeader().Model(ctx).Data(saveData).Insert()
			if utils.IsError(err) {
				return err
			}
		}
	}
	return
}

func (s *sSystemDept) DelLeader(ctx context.Context, id int64, userIds []int64) (err error) {
	_, err = service.SystemDeptLeader().Model(ctx).Where("dept_id", id).WhereIn("user_id", userIds).Delete()
	if utils.IsError(err) {
		return err
	}
	return
}

func (s *sSystemDept) Update(ctx context.Context, in *req.SystemDeptSave) (err error) {
	oldLevel := in.Level
	data, err := s.handleData(ctx, in)
	if err != nil {
		return
	}
	saveData := do.SystemDept{
		Name:     data.Name,
		ParentId: data.ParentId,
		Level:    data.Level,
		Sort:     data.Sort,
		Status:   data.Status,
		Leader:   data.Leader,
		Phone:    data.Phone,
		Remark:   data.Remark,
	}
	_, err = s.Model(ctx).Where(dao.SystemDept.Columns().Id, data.Id).Data(saveData).Update()
	if utils.IsError(err) {
		return err
	}
	if !s.checkChildrenExists(ctx, data.Id) {
		return
	}

	var dept []*entity.SystemDept

	err = s.Model(ctx).Unscoped().WhereLike("level", "%,"+gconv.String(data.Id)+",%").Scan(&dept)
	if utils.IsError(err) {
		return err
	}
	if !g.IsEmpty(dept) {
		for _, item := range dept {
			newLevel := utils.ReplaceSubstr(item.Level, oldLevel, data.Level)
			_, err = s.Model(ctx).Unscoped().Where(dao.SystemDept.Columns().Id, item.Id).Data(do.SystemDept{
				Level: newLevel,
			}).Update()
			if utils.IsError(err) {
				return err
			}
		}
	}
	return
}

func (s *sSystemDept) checkChildrenExists(ctx context.Context, id int64) bool {
	count, err := s.Model(ctx).Unscoped().Where("parent_id", id).Count()
	if utils.IsError(err) {
		return false
	}
	if count > 0 {
		return true
	}
	return false
}

func (s *sSystemDept) checkChildrenUnscopedAllExists(ctx context.Context, id int64, ids []int64) bool {
	var depts []*entity.SystemDept
	err := s.Model(ctx).Unscoped().Where("parent_id", id).Scan(&depts)
	if utils.IsError(err) {
		return false
	}
	hasAllChildrenExists := true
	if !g.IsEmpty(depts) {
		for _, item := range depts {
			if !slice.Contains(ids, item.Id) {
				hasAllChildrenExists = false
			}
		}
	}
	return hasAllChildrenExists
}

func (s *sSystemDept) checkChildrenAllExists(ctx context.Context, id int64, ids []int64) bool {
	var depts []*entity.SystemDept
	err := s.Model(ctx).Where("parent_id", id).Scan(&depts)
	if utils.IsError(err) {
		return false
	}
	hasAllChildrenExists := true
	if !g.IsEmpty(depts) {
		for _, item := range depts {
			if !slice.Contains(ids, item.Id) {
				hasAllChildrenExists = false
			}
		}
	}
	return hasAllChildrenExists
}

func (s *sSystemDept) Delete(ctx context.Context, ids []int64) (names []string, err error) {
	ctuIds := make([]int64, 0)
	for _, id := range ids {
		if s.checkChildrenAllExists(ctx, id, ids) {
			_, err = s.Model(ctx).Where("id", id).Delete()
			if utils.IsError(err) {
				return nil, err
			}
		} else {
			ctuIds = append(ctuIds, id)
		}
	}

	if !g.IsEmpty(ctuIds) {
		result, err := s.Model(ctx).Fields("name").WhereIn("id", ctuIds).Array()
		if utils.IsError(err) {
			return nil, err
		}
		if !g.IsEmpty(result) {
			names = gconv.SliceStr(result)
		}
	}
	return
}

func (s *sSystemDept) RealDelete(ctx context.Context, ids []int64) (names []string, err error) {
	ctuIds := make([]int64, 0)
	for _, id := range ids {
		if s.checkChildrenUnscopedAllExists(ctx, id, ids) {
			_, err = s.Model(ctx).Unscoped().Where("id", id).Delete()
			if utils.IsError(err) {
				return nil, err
			}
		} else {
			ctuIds = append(ctuIds, id)
		}
	}

	if !g.IsEmpty(ctuIds) {
		result, err := s.Model(ctx).Unscoped().Fields("name").WhereIn("id", ctuIds).Array()
		if utils.IsError(err) {
			return nil, err
		}
		if !g.IsEmpty(result) {
			names = gconv.SliceStr(result)
		}
	}

	return
}

func (s *sSystemDept) Recovery(ctx context.Context, ids []int64) (err error) {
	_, err = s.Model(ctx).Unscoped().WhereIn("id", ids).Update(g.Map{"deleted_at": nil})
	if utils.IsError(err) {
		return err
	}
	return
}

func (s *sSystemDept) ChangeStatus(ctx context.Context, id int64, status int) (err error) {
	_, err = s.Model(ctx).Data(g.Map{"status": status}).Where("id", id).Update()
	if utils.IsError(err) {
		return err
	}
	return
}

func (s *sSystemDept) NumberOperation(ctx context.Context, id int64, numberName string, numberValue int) (err error) {
	_, err = s.Model(ctx).Data(g.Map{numberName: numberValue}).Where("id", id).Update()
	if utils.IsError(err) {
		return err
	}
	return
}
