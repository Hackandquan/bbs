package cache

import (
	"bbs/model"
	"bbs/repositories"
	"errors"
	"time"

	"github.com/goburrow/cache"
	"github.com/mlogclub/simple"
)

// 使用https://github.com/goburrow/cache这个包实现缓存

/**
作用：将登录用户的token信息缓存到内存中，加快获取用户token信息的速度
*/
var UserToken = newUserTokeCache()

type userTokenCache struct {
	cache cache.LoadingCache
}

func newUserTokeCache() *userTokenCache {
	return &userTokenCache{
		// 第一个函数的作用是，当key对应的缓存数据不存在的时候，就会将函数的返回值作为key的缓存
		cache: cache.NewLoadingCache(func(key cache.Key) (value cache.Value, e error) {
			// key 就是传入的key
			token := key.(string)
			value = repositories.UserTokenRepository.GetByToken(simple.DB(), token)
			if value == nil {
				e = errors.New("数据不存在")
			}
			return
		}, cache.WithMaximumSize(1000), cache.WithExpireAfterAccess(60*time.Minute)),
	}
}

func (c *userTokenCache) Get(token string) *model.UserToken {
	if len(token) == 0 {
		return nil
	}
	val, err := c.cache.Get(token)
	if err != nil {
		return nil
	}
	if val == nil {
		return nil
	}
	return val.(*model.UserToken)
}
