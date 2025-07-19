package user

import (
	"context"
	"devinggo/api"
	"devinggo/internal/service"
)

var (
	UserController = userController{}
)

type userController struct {
}

func (c *userController) Health(ctx context.Context, in *api.GetOneReq) (out *api.GetOneRes, err error) {
	out = &api.GetOneRes{}
	out.User, err = service.User().GetOne(ctx, in.Id)
	return
}
