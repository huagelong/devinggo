package user

import (
	"context"
	"devinggo/internal/dao"
	"devinggo/internal/model/entity"
	"devinggo/internal/service"
	"devinggo/modules/system/logic/base"
)

type sUser struct {
	base.BaseService
}

func init() {
	service.RegisterUser(NewUser())
}

func NewUser() *sUser {
	return &sUser{}
}

func (s *sUser) GetOne(ctx context.Context, id int64) (res *entity.User, err error) {
	res = &entity.User{}
	err = dao.User.Ctx(ctx).Where(dao.User.Columns().Id, id).Scan(res)
	return
}
