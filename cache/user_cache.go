package cache

import (
	"bbs/model"
	"bbs/model/constants"

	"github.com/goburrow/cache"
)

var UserCache = newUserCache()

type userCache struct {
	cache cache.LoadingCache
}

func newUserCache() *userCache {
	return &userCache{
		cache: cache.NewLoadingCache(func(key cache.Key) (value cache.Value, e error) {

		}),
	}
}

func (c *userCache) Get(userId int64) *model.User {
	if userId == constants.NilInt64 {
		return nil
	}
	user, err := c.cache.Get(userId)
	if err != nil {
		return nil
	}
	if user == nil {
		return nil
	}
	return user.(*model.User)
}
