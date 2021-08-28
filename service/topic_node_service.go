package service

import (
	"bbs/model"
	"fmt"

	"github.com/mlogclub/simple"
)

var TopicNodeService = newTopicNodeService()

func newTopicNodeService() *topicNodeService {
	return &topicNodeService{}
}

type topicNodeService struct{}

func (s *topicNodeService) GetNodeById(topicId int64) *model.TopicNode {
	sql := `SELECT * FROM t_topic_node WHERE topic_id=? AND status=1`
	ret := &model.TopicNode{}
	if err := simple.DB().Exec(sql).First(&ret); err != nil {
		fmt.Println("TopicNodeService.GteNodeById Error:", err)
	}
	return ret
}
