package service

import "github.com/mlogclub/simple"

var UserLikeService *userLikeService = newUserLikeService()

func newUserLikeService() *userLikeService {
	return &userLikeService{}
}

type userLikeService struct{}

// 判断用户在entity类型资源中有没有ids这些资源，返回liked的ids
func (s *userLikeService) IsLiked(userId int64, entityType string, ids []int64) (likedIds []int64) {
	// SELECT entity_id from user_like where entity=? AND user_id=? AND entity_id in ?
	simple.DB().Exec("SELECT entity_id FROM user_like user_id=? AND entity_type=? AND entity_id in ?", userId, entityType, ids).Find(&likedIds)
	return
}
