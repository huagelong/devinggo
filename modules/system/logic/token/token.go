// Package token
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package token

import (
	"context"
	"devinggo/modules/system/model"
	"devinggo/modules/system/model/res"
	"devinggo/modules/system/myerror"
	"devinggo/modules/system/pkg/contexts"
	"devinggo/modules/system/pkg/redis"
	"devinggo/modules/system/pkg/utils"
	"devinggo/modules/system/pkg/utils/config"
	"devinggo/modules/system/service"
	"fmt"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type sToken struct {
}

func init() {
	service.RegisterToken(New())
}

func New() *sToken {
	return &sToken{}
}

type Token struct {
	ExpireAt     int64 `json:"exp"` // token过期时间
	RefreshAt    int64 `json:"ra"`  // 刷新时间
	RefreshCount int64 `json:"rc"`  // 刷新次数
}

const (
	CacheToken       = "token"      // 登录token
	CacheTokenBind   = "token_bind" // 登录用户身份绑定
	defaultSecretKey = "devinggohello123"
)

type Claims struct {
	*model.NormalIdentity
	jwt.RegisteredClaims
}

func (s *sToken) GetToken(ctx context.Context, scene string, claimsData map[string]interface{}) (string, int64, error) {
	now := gtime.Now()
	expiresConfig := config.GetConfigInt64(ctx, "token.expires", 604800)
	secretKey := config.GetConfigString(ctx, "token.secretKey", defaultSecretKey)
	normalIdentity := &model.NormalIdentity{}
	normalIdentity.Data = claimsData
	normalIdentity.Scene = scene
	normalIdentity.ExpiresAt = now.Unix() + expiresConfig
	claims := Claims{
		normalIdentity,
		jwt.RegisteredClaims{},
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(secretKey))
	if err != nil {
		return "", now.Unix(), err
	}
	return token, normalIdentity.ExpiresAt, nil
}

func (s *sToken) ParseToken(ctx context.Context, token string) (*model.NormalIdentity, error) {
	secretKey := config.GetConfigString(ctx, "token.secretKey", defaultSecretKey)
	tokenObj, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil {
		g.Log().Debugf(ctx, "parseToken err:%+v", err)
		return nil, err
	}

	if !tokenObj.Valid {
		return nil, myerror.ValidationFailed(ctx, "token验证失败")
	}

	claims, ok := tokenObj.Claims.(*Claims)
	if !ok {
		return nil, myerror.ValidationFailed(ctx, "token验证失败")
	}
	return claims.NormalIdentity, nil
}

func (s *sToken) Refresh(r *ghttp.Request) (string, int64, error) {
	var (
		ctx   = r.Context()
		token = s.GetAuthorization(r)
	)

	if token == "" {
		err := myerror.ValidationFailed(ctx, "token验证失败")
		return "", 0, err
	}

	claims, err := s.ParseToken(ctx, token)
	if err != nil {
		g.Log().Debugf(ctx, "logout parseToken err:%+v", err)
		err := myerror.ValidationFailed(ctx, "token验证失败")
		return "", 0, err
	}

	data := claims.Data
	scene := claims.Scene
	return s.GetToken(ctx, scene, data)
}

func (s *sToken) GenerateUserToken(ctx context.Context, scene, appId string, user *model.Identity) (string, int64, error) {
	now := gtime.Now()
	userMap := gconv.Map(user)
	token, exp, err := s.GetToken(ctx, scene, userMap)
	if err != nil {
		return "", now.Unix(), err
	}
	expiresConfig := config.GetConfigInt64(ctx, "token.expires", 604800)
	var (
		// 认证key
		authKey = s.GetAuthKey(token)
		// 登录token
		tokenKey = s.GetTokenKey(appId, authKey)
		// 身份绑定
		bindKey = s.GetBindKey(appId, user.Id)
		// 有效时长
		duration = time.Second * gconv.Duration(expiresConfig)
	)

	tokenStruct := &Token{
		ExpireAt:     exp,
		RefreshAt:    now.Unix(),
		RefreshCount: 0,
	}

	if err = redis.GetCacheClient().Set(ctx, tokenKey, tokenStruct, duration); err != nil {
		return "", now.Unix(), err
	}

	if err = redis.GetCacheClient().Set(ctx, bindKey, tokenKey, duration); err != nil {
		return "", now.Unix(), err
	}
	return token, tokenStruct.ExpireAt, nil
}

func (s *sToken) Logout(r *ghttp.Request) (err error) {
	var (
		ctx   = r.Context()
		token = s.GetAuthorization(r)
	)

	if token == "" {
		err = myerror.ValidationFailed(ctx, "token验证失败")
		return
	}

	claims, err := s.ParseToken(ctx, token)
	if err != nil {
		g.Log().Debugf(ctx, "logout parseToken err:%+v", err)
		err = myerror.ValidationFailed(ctx, "token验证失败")
		return
	}

	data := claims.Data
	var user *model.Identity
	err = gconv.Scan(data, &user)
	if err != nil {
		return err
	}

	var (
		// 认证key
		authKey = s.GetAuthKey(token)
		// 登录token
		tokenKey = s.GetTokenKey(contexts.New().GetModule(ctx), authKey)
		// 身份绑定
		bindKey = s.GetBindKey(contexts.New().GetModule(ctx), user.Id)
	)

	// 删除token
	if _, err = redis.GetCacheClient().Remove(ctx, tokenKey); err != nil {
		return
	}
	multiLogin := config.GetConfigBool(ctx, "token.multiLogin", true)
	if !multiLogin {
		if _, err = redis.GetCacheClient().Remove(ctx, bindKey); err != nil {
			return
		}
	}
	return
}

// 强退用户
func (s *sToken) Kick(r *ghttp.Request, userId int64, appId string) (err error) {
	ctx := r.Context()
	bindKey := s.GetBindKey(appId, userId)
	tokenKeyArr, err := redis.GetCacheClient().Get(ctx, bindKey)
	if err != nil {
		return
	}
	if !tokenKeyArr.IsEmpty() {
		if _, err = redis.GetCacheClient().Remove(ctx, tokenKeyArr.String()); err != nil {
			return
		}

		if _, err = redis.GetCacheClient().Remove(ctx, bindKey); err != nil {
			return
		}
	}

	return
}

// 强退所有app用户
func (s *sToken) KickAll(r *ghttp.Request, userId int64) (err error) {
	ctx := r.Context()
	match := fmt.Sprintf("%v:*", "token_bind")
	iterator := uint64(0)
	for {
		iterator, keys, err := redis.GetRedis().Scan(ctx, iterator, gredis.ScanOption{
			Match: match,
			Count: 100,
		})

		if err != nil {
			g.Log().Warning(ctx, "getAllUserIds redis.GetRedis().Do(SCAN) error:", err)
			break
		}

		if g.IsEmpty(keys) {
			break
		}
		dataSlice := gconv.SliceStr(keys)
		for _, value := range dataSlice {
			tmp := gstr.Split(value, ":")
			userIdTmp := gconv.Int64(tmp[2])
			appId := tmp[1]
			if userIdTmp == userId {
				s.Kick(r, userIdTmp, appId)
			}
		}

		if iterator == 0 {
			break
		}
	}

	return
}

func (s *sToken) GetAllUserIds(r *ghttp.Request) (userApps []res.SystemUserApp, err error) {
	ctx := r.Context()
	match := fmt.Sprintf("%v:*", "token_bind")
	iterator := uint64(0)
	userApps = make([]res.SystemUserApp, 0)
	for {
		iterator, keys, err := redis.GetRedis().Scan(ctx, iterator, gredis.ScanOption{
			Match: match,
			Count: 100,
		})

		if err != nil {
			g.Log().Warning(ctx, "getAllUserIds redis.GetRedis().Do(SCAN) error:", err)
			break
		}

		if g.IsEmpty(keys) {
			break
		}
		dataSlice := gconv.SliceStr(keys)
		for _, value := range dataSlice {
			tmp := gstr.Split(value, ":")
			userApps = append(userApps, res.SystemUserApp{
				AppId:  tmp[1],
				UserId: gconv.Int64(tmp[2]),
			})
		}

		if iterator == 0 {
			break
		}
	}
	return
}

// ParseLoginUser 解析登录用户信息
func (s *sToken) ParseLoginUser(r *ghttp.Request, appId string) (*model.Identity, error) {
	var (
		ctx   = r.Context()
		token = s.GetAuthorization(r)
	)

	if token == "" {
		return nil, myerror.ValidationFailed(ctx, "token不能为空")
	}

	claims, err := s.ParseToken(ctx, token)
	if err != nil {
		return nil, err
	}

	if g.IsEmpty(claims) {
		return nil, myerror.ValidationFailed(ctx, "token验证失败")
	}
	var user *model.Identity
	data := claims.Data
	err = gconv.Scan(data, &user)
	if err != nil {
		return nil, err
	}

	if g.IsEmpty(appId) {
		appId = user.AppId
	}

	var (
		// 认证key
		authKey = s.GetAuthKey(token)
		// 登录token
		tokenKey = s.GetTokenKey(appId, authKey)
		// 身份绑定
		bindKey = s.GetBindKey(appId, user.Id)
	)

	// 检查token是否存在
	tk, err := redis.GetCacheClient().Get(ctx, tokenKey)
	if err != nil {
		return nil, err
	}

	if tk.IsEmpty() {
		err = myerror.ValidationFailed(ctx, "token验证失败")
		return nil, err
	}

	var tokenStruct *Token
	if err = tk.Scan(&tokenStruct); err != nil {
		g.Log().Debugf(ctx, "token scan err:%+v", err)
		err = myerror.ValidationFailed(ctx, "token验证失败")
		return nil, err
	}

	if tokenStruct == nil {
		g.Log().Debugf(ctx, "token = nil")
		err = myerror.ValidationFailed(ctx, "token验证失败")
		return nil, err
	}

	now := gtime.Now()
	if tokenStruct.ExpireAt < now.Unix() {
		g.Log().Debugf(ctx, "token expired.")
		err = myerror.ValidationFailed(ctx, "token验证失败")
		return nil, err
	}

	// 是否允许多端登录
	multiLogin := config.GetConfigBool(ctx, "token.multiLogin", true)
	if !multiLogin {
		origin, err := redis.GetCacheClient().Get(ctx, bindKey)
		if err != nil {
			return nil, err
		}

		if origin == nil || origin.IsEmpty() {
			return nil, myerror.ValidationFailed(ctx, "token验证失败")
		}

		if tokenKey != origin.String() {
			return nil, myerror.ValidationFailed(ctx, "token验证失败")
		}
	}

	// 自动刷新token有效期
	refreshToken := func() {
		// 未开启自动刷新
		autoRefresh := config.GetConfigBool(ctx, "token.autoRefresh", true)
		if !autoRefresh {
			return
		}

		// 刷新次数已达上限
		maxRefreshTimes := config.GetConfigInt64(ctx, "token.maxRefreshTimes", -1)
		if maxRefreshTimes != -1 && tokenStruct.RefreshCount >= maxRefreshTimes {
			return
		}

		// 未达到刷新间隔
		refreshInterval := config.GetConfigInt64(ctx, "token.refreshInterval", 86400)
		if gtime.New(tokenStruct.RefreshAt).Unix()+refreshInterval > now.Unix() {
			return
		}

		// 刷新有效期
		expiresConfig := config.GetConfigInt64(ctx, "token.expires", 604800)
		tokenStruct.ExpireAt = now.Unix() + expiresConfig
		tokenStruct.RefreshAt = now.Unix()
		tokenStruct.RefreshCount += 1

		duration := time.Second * gconv.Duration(expiresConfig)

		if err = redis.GetCacheClient().Set(ctx, tokenKey, token, duration); err != nil {
			return
		}

		if err = redis.GetCacheClient().Set(ctx, bindKey, tokenKey, duration); err != nil {
			return
		}
	}

	utils.SafeGo(ctx, func(ctx context.Context) {
		refreshToken()
	})

	return user, nil
}

func (s *sToken) GetAuthorization(r *ghttp.Request) string {
	// 默认从请求头获取
	var authorization = r.Header.Get("Authorization")

	// 如果请求头不存在则从get参数获取
	if authorization == "" {
		return r.Get("authorization").String()
	}
	return gstr.Replace(authorization, "Bearer ", "")
}

// GetAuthKey 认证key
func (s *sToken) GetAuthKey(token string) string {
	return gmd5.MustEncryptString("devinggo:auth:" + token)
}

// GetTokenKey 令牌缓存key
func (s *sToken) GetTokenKey(appId, authKey string) string {
	return fmt.Sprintf("%v:%v:%v", CacheToken, appId, authKey)
}

// GetBindKey 令牌身份绑定key
func (s *sToken) GetBindKey(appId string, userId int64) string {
	return fmt.Sprintf("%v:%v:%v", CacheTokenBind, appId, userId)
}
