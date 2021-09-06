package render

import (
	"bbs/model"
	"bbs/model/constants"
	"bbs/service"

	"github.com/mlogclub/simple"
)

func BuildSimpleTopics(topics []model.Topic, currentUser *model.User) []model.TopicResponse {
	if len(topics) == 0 {
		return make([]model.TopicResponse, 0)
	}
	// 如果用户登录了网站，则返回用户对文章是否点了喜欢 liked
	var likedTopicIds []int64
	if currentUser != nil {
		var topicIds []int64
		for _, topic := range topics {
			topicIds = append(topicIds, topic.Id)
		}
		likedTopicIds = service.UserLikeService.IsLiked(currentUser.Id, constants.EntityTopic, topicIds)
	}

	var response []model.TopicResponse
	for _, topic := range topics {
		var liked = simple.Contains(topic.Id, likedTopicIds)
		var item = BuildSimpleTopic(&topic)
		item.Liked = liked
		response = append(response, *item)
	}
	return response
}

func BuildSimpleTopic(topic *model.Topic) *model.TopicResponse {
	// topic.Id // topic的作者id
	authorInfo := BuildUserDefaultIfNull(topic.UserId)
	// 查找topic对应的node的详细信息
	nodeDetail := service.TopicNodeService.GetNodeById(topic.NodeId)
	// 查找topic的标签s
	topicTags := service.TopicTagService.GetTagsByTopicId(topic.Id, 1)
	return &model.TopicResponse{
		TopicId:         topic.Id,
		Type:            topic.Type,
		User:            authorInfo,
		Node:            BuildNode(nodeDetail),
		Tags:            BuildTags(topicTags),
		Title:           topic.Title,
		Summary:         "",
		Content:         topic.Content,
		ImageList:       []model.ImageInfo{},
		LastCommentTime: topic.LastCommentTime,
		ViewCount:       topic.ViewCount,
		CommentCount:    topic.CommentCount,
		LikeCount:       topic.LikeCount,
		Liked:           false,
		CreateTime:      topic.CreateTime,
		Recommend:       false,
		RecommendTime:   topic.RecommendTime,
	}
}
