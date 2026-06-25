// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"devinggo/modules/system/model"
	"devinggo/modules/system/model/res"

	"github.com/gogf/gf/v2/net/ghttp"
)

type (
	IToken interface {
		GetToken(ctx context.Context, scene string, claimsData map[string]interface{}) (string, int64, error)
		ParseToken(ctx context.Context, token string) (*model.NormalIdentity, error)
		Refresh(r *ghttp.Request) (string, int64, error)
		GenerateUserToken(ctx context.Context, scene string, appId string, user *model.Identity) (string, int64, error)
		Logout(r *ghttp.Request) (err error)
		// 强退用户
		Kick(r *ghttp.Request, userId int64, appId string) (err error)
		// 强退所有app用户
		KickAll(r *ghttp.Request, userId int64) (err error)
		GetAllUserIds(r *ghttp.Request) (userApps []res.SystemUserApp, err error)
		// ParseLoginUser 解析登录用户信息
		ParseLoginUser(r *ghttp.Request, appId string) (*model.Identity, error)
		GetAuthorization(r *ghttp.Request) string
		// GetAuthKey 认证key
		GetAuthKey(token string) string
		// GetTokenKey 令牌缓存key
		GetTokenKey(appId string, authKey string) string
		// GetBindKey 令牌身份绑定key
		GetBindKey(appId string, userId int64) string
	}
)

var (
	localToken IToken
)

func Token() IToken {
	if localToken == nil {
		panic("implement not found for interface IToken, forgot register?")
	}
	return localToken
}

func RegisterToken(i IToken) {
	localToken = i
}
