package repositories

import (
	"bbs/model"

	"gorm.io/gorm"
)

var UserTokenRepository *userTokenRepository = new(userTokenRepository)

type userTokenRepository struct{}

// 新建userToken
func (r *userTokenRepository) Create(db *gorm.DB, t *model.UserToken) error {
	return db.Create(t).Error
}

// 跟token从user_token表获取该token完整的信息
func (r *userTokenRepository) GetByToken(db *gorm.DB, token string) *model.UserToken {
	if len(token) == 0 {
		return nil
	}
	return r.Take(db, "token=?", token)
}

func (r *userTokenRepository) Take(db *gorm.DB, where ...interface{}) *model.UserToken {
	ret := &model.UserToken{}
	if err := db.Take(ret, where...).Error; err != nil {
		return nil
	}
	return ret
}
