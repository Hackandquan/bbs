package render

import (
	"bbs/model"
	"bbs/model/constants"
	"bbs/service"
	"strings"
)

func BuildUserDefaultIfNull(id int64) *model.UserInfo {
	/* user := cache.UserCache.Get(id)
	if user == nil {

	} */
	user := service.UserService.GetUserById(id)
	return BuildUser(user)
}

func BuildUser(user *model.User) *model.UserInfo {
	if user == nil {
		return nil
	}
	roles := strings.Split(user.Roles, ",")
	ret := &model.UserInfo{
		Id:                   user.Id,
		Username:             user.Username.String,
		Email:                user.Email.String,
		EmailVerified:        user.EmailVerified,
		Nickname:             user.Nickname,
		Avatar:               user.Avatar,
		SmallAvatar:          "", // 这里可以是oss处理过的图片路径
		BackgroundImage:      user.BackgroundImage,
		SmallBackgroundImage: "",
		Type:                 user.Type,
		Roles:                roles,
		HomePage:             user.HomePage,
		Description:          user.Description,
		Score:                user.Score,
		TopicCount:           user.TopicCount,
		CommentCount:         user.CommentCount,
		PasswordSet:          len(user.Password) > 0,
		Forbidden:            user.IsForbidden(),
		Status:               user.Status,
		CreateTime:           user.CreateTime,
	}
	// 判断用户是否已经被删除或者被禁言，是的话则修改用户信息
	if ret.Status == constants.StatusDeleted || user.IsForbidden() {
		ret.Forbidden = true
		ret.HomePage = ""
		ret.Description = ""
		ret.Email = ""
		ret.Nickname = "黑名单用户"
		ret.Score = 0
	}
	if len(ret.Description) == 0 {
		ret.Description = "这家伙很懒，什么都没留下"
	}
	return ret
}
