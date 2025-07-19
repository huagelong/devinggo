package api

import (
	"devinggo/internal/model/entity"
	"github.com/gogf/gf/v2/frame/g"
)

type GetOneReq struct {
	g.Meta `path:"/user/{id}" method:"get" tags:"User" x-exceptLogin:"true" x-exceptAuth:"true" summary:"Get one user"`
	Id     int64 `v:"required" dc:"user id"`
}
type GetOneRes struct {
	*entity.User `dc:"user"`
}
