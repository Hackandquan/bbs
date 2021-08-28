package service

import (
	"bbs/model"

	"github.com/mlogclub/simple"
)

var UserService = newUserService()

func newUserService() *userService {
	return &userService{}
}

type userService struct{}

func (s *userService) GetUserById(userId int64) *model.User {
	ret := &model.User{}
	sql := `SELECT * FROM t_user where id=?`
	if err := simple.DB().Exec(sql, userId, '1').First(ret).Error; err != nil {
		return nil
	}
	return ret
}
