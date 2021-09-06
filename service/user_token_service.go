package service

import (
	"bbs/cache"
	"bbs/model"
	"bbs/model/constants"
	"bbs/repositories"
	"time"

	"github.com/kataras/iris/v12"
	"github.com/mlogclub/simple"
	"github.com/mlogclub/simple/date"
)

var UserTokenService *userTokenService = &userTokenService{}

type userTokenService struct{}

// 生成token - 跟用户id关联（token包含过期时间和状态）
func (s *userTokenService) Genertate(userId int64) (string, error) {
	token := simple.UUID()
	tokenExpireDays := 7 // token过期时间：默认7天
	expireAt := time.Now().Add(time.Hour * 24 * time.Duration(tokenExpireDays))
	userToken := &model.UserToken{
		Token:      token,
		UserId:     userId,
		ExpiredAt:  date.Timestamp(expireAt),
		Status:     constants.StatusOk,
		CreateTime: date.NowTimestamp(),
	}
	// 往数据库里面写入token信息
	err := repositories.UserTokenRepository.Create(simple.DB(), userToken)
	if err != nil {
		return "", err
	}
	return token, nil
}

// 从请求上下文获取当前登录用户
func (s *userTokenService) GetCurrentUserId(ctx iris.Context) int64 {
	user := s.GetCurrent(ctx)
	if user != nil {
		return user.Id
	}
	return -1
}

// 获取当前登录用户相关信息 model.User
func (s *userTokenService) GetCurrent(ctx iris.Context) *model.User {
	token := s.GetUserToken(ctx)
	// 判断token是否存在：缓存 > 数据库
	userToken := cache.UserToken.Get(token) // 缓存怎么做？
	// userToken 不存在或者 userToken 是删除状态
	if userToken == nil || userToken.Status == constants.StatusDeleted {
		return nil
	}
	// 判断是否已过期： expireAt - 当前时间戳 <= 0
	if userToken.ExpiredAt-date.NowTimestamp() <= 0 {
		return nil
	}
	// token 正常，返回用户信息 model.User
	user := cache.UserCache.Get(userToken.UserId)
	if user == nil || user.Status != constants.StatusOk {
		return nil
	}
	return user
}

// 从请求中获取token
func (s *userTokenService) GetUserToken(ctx iris.Context) string {
	// 从传递参数里面拿，没有就从request header里面拿
	userToken := ctx.FormValue("userToken")
	if len(userToken) > 0 {
		return userToken
	}
	return ctx.GetHeader("X-user-Token")
}
