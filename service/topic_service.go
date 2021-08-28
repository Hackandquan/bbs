package service

import (
	"bbs/model"
	"bbs/model/constants"

	"github.com/mlogclub/simple"
)

var TopicService = newTopicService()

func newTopicService() *topicService {
	return &topicService{}
}

type topicService struct{}

func (s *topicService) GetTopics(nodeId, cursor int64, recommend bool) (topics []model.Topic, nextCursor int64, hasMore bool) {
	limit := 20 // 每页20条
	cnd := simple.NewSqlCnd()
	if nodeId > 0 {
		cnd.Eq("node_id", nodeId)
	}
	if cursor > 0 { // 分页用的
		cnd.Lt("last_comment_time", cursor)
	}

	if recommend {
		cnd.Eq("recommend", true)
	}

	cnd.Eq("status", constants.StatusOk).Desc("last_comment_time").Limit(limit)
	simple.DB().Find(&topics, cnd)
	if len(topics) > 0 {
		nextCursor = topics[len(topics)-1].LastCommentTime
		hasMore = len(topics) >= limit
	} else {
		nextCursor = cursor // hasMore 默认false
	}
	return
}
